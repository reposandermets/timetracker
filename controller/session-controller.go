package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/reposandermets/timetracker/entity"
	"github.com/reposandermets/timetracker/service"
)

type SessionController interface {
	Save(ctx *gin.Context) (entity.Session, error)
	Update(ctx *gin.Context) (entity.Session, error)
	FindAll() []entity.Session
}

type controller struct {
	service service.SessionService
}

func New(service service.SessionService) SessionController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Session {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) (entity.Session, error) {
	var session entity.Session
	err := ctx.ShouldBindJSON(&session)
	if err != nil {
		return session, err
	}
	c.service.Save(session)
	return session, nil
}

func (c *controller) Update(ctx *gin.Context) (entity.Session, error) {
	var session entity.Session
	err := ctx.ShouldBindJSON(&session)
	if err != nil {
		return session, err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return session, err
	}
	session.ID = id
	c.service.Update(session)
	return session, nil
}
