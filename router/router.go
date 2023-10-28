package router

import (
	"blog/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//	creates a router as 'r' & sets it to use the gin framework's default settings.
	//This allows the router to use all of the default routes and middleware functions that are available in the gin framework.
	r := gin.Default()
	r.POST("/postblog", handler.PostBlog)
	r.GET("/getblogposts", handler.GetBlogPosts)
	r.GET("/getblogpost/:id", handler.GetBlogPost)
	r.PUT("/updateblogpost/:id", handler.UpdateBlogPost)
	r.DELETE("/deleteblogpost/:id", handler.DeleteBlogPost)
	return r
}
