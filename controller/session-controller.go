package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/reposandermets/timetracker/entity"
	"github.com/reposandermets/timetracker/service"
)

type SessionController interface {
	Save(ctx *gin.Context) (entity.Session, error)
	Update(ctx *gin.Context) (entity.Session, error)
	Find(ctx *gin.Context) ([]entity.Session, error)
	FindUsers() ([]entity.User, error)
}

type controller struct {
	service service.SessionService
}

func New(service service.SessionService) SessionController {
	return &controller{
		service: service,
	}
}

type UserT struct {
	UserID string `form:"user_id" binding:"required"`
}

func (c *controller) FindUsers() ([]entity.User, error) {
	return c.service.FindUsers()
}

func (c *controller) Find(ctx *gin.Context) ([]entity.Session, error) {
	sessions := []entity.Session{}
	user := UserT{}
	if err := ctx.ShouldBindWith(&user, binding.Query); err != nil {
		return sessions, err
	}

	uuid, err := uuid.Parse(user.UserID)

	if err != nil {
		return sessions, err
	}

	return c.service.FindSessionsByUserId(uuid)
}

func (c *controller) Save(ctx *gin.Context) (entity.Session, error) {
	var session entity.Session
	err := ctx.ShouldBindJSON(&session)
	if err != nil {
		return session, err
	}
	if session.Status != "started" {
		return session, errors.New("status has to be started'")
	}
	return c.service.Save(session)
}

func (c *controller) Update(ctx *gin.Context) (entity.Session, error) {
	var session entity.Session
	err := ctx.ShouldBindJSON(&session)
	if err != nil {
		return session, err
	}

	uuid, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		return session, err
	}

	if session.Status != "started" && session.Status != "paused" && session.Status != "ended" {
		return session, errors.New("status should be one of [started, paused, ended]")
	}

	session.ID = uuid
	return c.service.Update(session)
}
