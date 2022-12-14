package api

import (
	"papercrane/middleware"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db             *gorm.DB
	staticFilePath string
	Router         *gin.Engine
}

func NewServer(db *gorm.DB, dir string) *Server {
	server := &Server{db: db, staticFilePath: dir}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// combine gin router with react router
	fs := static.LocalFile(server.staticFilePath, true)
	r.Use(static.Serve("/", fs))
	r.Use(static.Serve("/edit", fs))
	r.Use(static.Serve("/view", fs))
	r.Use(middleware.CORSMiddleware())

	// customize route group
	storyGourp := r.Group("story")
	{
		storyGourp.POST("/create", server.CreateStory)
		storyGourp.GET("/query", server.QueryStories)
		storyGourp.GET("/view", server.ViewStory)
		storyGourp.POST("/update", server.UpdateStory)
		storyGourp.GET("/search", server.SearchStory)
	}

	server.Router = r
	return server
}
