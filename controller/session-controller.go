package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/reposandermets/timetracker/entity"
	"github.com/reposandermets/timetracker/service"
)

type SessionController interface {
	Save(ctx *gin.Context) (entity.Session, error)
	Update(ctx *gin.Context) (entity.Session, error)
}

type controller struct {
	service service.SessionService
}

func New(service service.SessionService) SessionController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) (entity.Session, error) {
	var session entity.Session
	err := ctx.ShouldBindJSON(&session)
	if err != nil {
		return session, err
	}
	if session.Status != "started" {
		return session, errors.New("status has to be 'started'")
	}
	s, err := c.service.Save(session)
	return s, err
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
	s, err := c.service.Update(session)
	return s, err
}
