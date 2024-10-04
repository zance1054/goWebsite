package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./static")

	// Serve HTML templates
	router.LoadHTMLGlob("templates/*")

	// Store your job description as a variable
	jobDescription := `As a Software Architect, I design, implement, and optimize large-scale continuous integration and delivery (CI/CD) pipelines, ensuring robust and scalable software development lifecycles. My role is focused on enhancing developer productivity and maintaining high-quality software delivery standards through automated tools and containerization technologies.

I specialize in Kubernetes and Docker, architecting container orchestration solutions that enable efficient scaling and management of distributed applications. I develop and optimize end-to-end CI/CD pipelines using tools like Jenkins, Artifactory, and Helm, ensuring seamless deployment from development to production.

I lead teams of developers and DevOps engineers across different regions, ensuring alignment on architecture and delivery strategies, while focusing on performance optimization, automation, and cloud integration.`

	// Route for the HTML page with Bootstrap and the dynamic job description
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":       "Bootstrap Example",
			"description": jobDescription,
		})
	})

	// Route for JSON response
	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, JSON!",
		})
	})

	// Route for blog list page
	router.GET("/blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog_list.html", nil)
	})

	// Route for experience page
	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})

	// Route for about page
	router.GET("/experience", func(c *gin.Context) {
		c.HTML(http.StatusOK, "experience.html", nil)
	})

	// Start the server on port 8080
	router.Run(":8080")

	// Render individual blog posts
	router.GET("/blog_posts/*filepath", func(c *gin.Context) {
		filepath := c.Param("filepath")
		c.HTML(http.StatusOK, "blog_post.html", gin.H{
			"filepath": filepath,
		})
	})

	// Start the server on port 8080
	router.Run(":8080")

}
