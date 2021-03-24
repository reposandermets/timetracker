package server

import (
	"github.com/gin-gonic/gin"
	"github.com/reposandermets/timetracker/controller"
	"github.com/reposandermets/timetracker/repository"
	"github.com/reposandermets/timetracker/service"
)

var (
	sessionRepository repository.SessionRepository = repository.NewSessionRepository()
	sessionService    service.SessionService       = service.New(sessionRepository)
	sessionController controller.SessionController = controller.New(sessionService)
)

func Run() {
	defer sessionRepository.CloseDB()
	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/session", func(ctx *gin.Context) {
			sessions, err := sessionController.Find(ctx)
			if err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(200, sessions)
		})

		apiRoutes.POST("/session", func(ctx *gin.Context) {
			session, err := sessionController.Save(ctx)
			if err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(200, session)
		})

		apiRoutes.PUT("/session/:id", func(ctx *gin.Context) {
			session, err := sessionController.Update(ctx)
			if err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(200, session)
		})
	}

	server.Run("0.0.0.0:8080")
}
