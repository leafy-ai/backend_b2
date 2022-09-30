package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var blogManager BlogManager = BlogManager{}

// Returns Preset Routes to be used in the main function
func GetRoutes(g *gin.Engine) {
	blogManager.new()
	protected := g.Group("/blogs")
	protected.Use(JwtAuthMiddleware())
	protected.GET("/user", getAllBlogs)
	protected.POST("/create", createBlog)
	g.GET("/", base)
	g.GET("/blogs/all", getAllBlogs)
	// g.GET(one)
}

func base(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func getAllBlogs(c *gin.Context) {
	blogs, _ := blogManager.GetAllBlogs()
	c.JSON(http.StatusOK, blogs)
}

func createBlog(c *gin.Context) {
	user_id, err := ExtractTokenID(c)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
	}

	blog := Blog{}
	user := &User{}
	err = user.getById(user_id)
	blog.Createdby = *user
	c.BindJSON(&blog)
	err = blogManager.CreateBlog(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": blog})
	}
}
