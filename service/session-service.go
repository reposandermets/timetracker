package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/reposandermets/timetracker/entity"
	"github.com/reposandermets/timetracker/repository"
)

type SessionService interface {
	Save(entity.Session) (entity.Session, error)
	Update(session entity.Session) (entity.Session, error)
}

type sessionService struct {
	sessionRepository repository.SessionRepository
}

func New(repo repository.SessionRepository) SessionService {
	return &sessionService{
		sessionRepository: repo,
	}
}

func (service *sessionService) Save(session entity.Session) (entity.Session, error) {
	userResult := service.sessionRepository.FindUserById(&entity.User{}, session.UserID)
	if errors.Is(userResult.Error, gorm.ErrRecordNotFound) {
		return session, errors.New("user not found")
	}

	sessionResult := service.sessionRepository.FindStartedSessionByUserId(&entity.Session{}, session.UserID)
	if sessionResult.RowsAffected == 1 {
		return session, errors.New("user has started session")
	}

	timeNow := time.Now()
	session.StartedAt = timeNow
	session.StopperAt = timeNow
	session.Seconds = 0
	session.ID = uuid.New()
	result := service.sessionRepository.Save(&session)
	return session, result.Error
}

func (service *sessionService) Update(session entity.Session) (entity.Session, error) {
	existingSession := entity.Session{}
	sessionResult := service.sessionRepository.FindSessionById(&existingSession, session.ID)

	if errors.Is(sessionResult.Error, gorm.ErrRecordNotFound) {
		return session, errors.New("session not found")
	}

	if existingSession.UserID != session.UserID {
		return session, errors.New("incorrect user")
	}

	if existingSession.Status == "ended" {
		return session, errors.New("session has been already ended")
	}

	if existingSession.Status == session.Status {
		return session, errors.New("session is already " + session.Status)
	}

	currentTime := time.Now()
	oldTime := existingSession.StopperAt
	diff := currentTime.Sub(oldTime)

	session.StartedAt = existingSession.StartedAt

	switch {

	case existingSession.Status == "started" && session.Status == "paused":
		session.Seconds = existingSession.Seconds + diff.Seconds()

	case existingSession.Status == "started" && session.Status == "ended":
		session.Seconds = existingSession.Seconds + diff.Seconds()
		session.EndedAt = time.Now()

	case existingSession.Status == "paused" && session.Status == "started":
		session.Seconds = existingSession.Seconds
		session.StopperAt = time.Now()

	case existingSession.Status == "paused" && session.Status == "ended":
		session.Seconds = existingSession.Seconds
		session.EndedAt = time.Now()

	default:
		return session, errors.New("unknown use case current: " + existingSession.Status + " requested: " + session.Status)

	}

	result := service.sessionRepository.Update(&session)
	return session, result.Error
}
