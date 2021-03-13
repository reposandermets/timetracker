package service

import (
	"github.com/reposandermets/timetracker/entity"
	"github.com/reposandermets/timetracker/repository"
)

type SessionService interface {
	Save(entity.Session) entity.Session
	Update(session entity.Session) entity.Session
	FindAll() []entity.Session
}

type sessionService struct {
	sessionRepository repository.SessionRepository
}

func New(repo repository.SessionRepository) SessionService {
	return &sessionService{
		sessionRepository: repo,
	}
}

func (service *sessionService) Save(session entity.Session) entity.Session {
	service.sessionRepository.Save(session)
	return session
}

func (service *sessionService) Update(session entity.Session) entity.Session {
	service.sessionRepository.Update(session)
	return session
}

func (service *sessionService) FindAll() []entity.Session {
	return service.sessionRepository.FindAll()
}
