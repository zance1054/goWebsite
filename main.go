package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files for resources first
	router.Static("/resources", "./resources")

	// Serve blog posts directly, allowing access to the full path
	router.Static("/blog_posts", "./blog_posts")

	// Serve static files for CSS and JS
	router.Static("/static", "./static")

	// Serve HTML templates
	router.LoadHTMLGlob("templates/*")

	// Store your job description
	jobDescription := `As a Software Architect, I design, implement, and optimize large-scale continuous integration and delivery (CI/CD) pipelines, ensuring robust and scalable software development lifecycles. My role is focused on enhancing developer productivity and maintaining high-quality software delivery standards through automated tools and containerization technologies.

I specialize in Kubernetes and Docker, architecting container orchestration solutions that enable efficient scaling and management of distributed applications. I develop and optimize end-to-end CI/CD pipelines using tools like Jenkins, Artifactory, and Helm, ensuring seamless deployment from development to production.

I lead teams of developers and DevOps engineers across different regions, ensuring alignment on architecture and delivery strategies, while focusing on performance optimization, automation, and cloud integration.`

	// Route for the main HTML page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":       "Alexander Fielding",
			"description": jobDescription,
		})
	})

	// Route for JSON response
	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, JSON!"})
	})

	// Route for the blog list page
	router.GET("/blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog_list.html", nil)
	})

	// Route for individual blog posts
	router.GET("/blog/:year/:month/:filename", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		filename := c.Param("filename")
		filePath := filepath.Join("blog_posts", year, month, filename)

		// Serve the blog post directly as HTML
		c.HTML(http.StatusOK, filePath, nil)
	})

	// Route for about page
	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})

	// Route for experience page
	router.GET("/experience", func(c *gin.Context) {
		c.HTML(http.StatusOK, "experience.html", nil)
	})

	// Route to get gallery images
	router.GET("/api/getGalleryImages", getGalleryImages)

	// Start the server on port 8080
	router.Run(":8080")
}

// getGalleryImages serves the list of images in the gallery directory
func getGalleryImages(c *gin.Context) {
	galleryPath := "resources/gallery" // Adjust if necessary
	files, err := os.ReadDir(galleryPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read gallery directory"})
		return
	}

	var imageFiles []string
	for _, file := range files {
		if !file.IsDir() && (filepath.Ext(file.Name()) == ".jpg" || filepath.Ext(file.Name()) == ".jpeg" || filepath.Ext(file.Name()) == ".png" || filepath.Ext(file.Name()) == ".gif") {
			imageFiles = append(imageFiles, "/resources/gallery/"+file.Name()) // Include the path
		}
	}

	// Return the JSON response
	c.JSON(http.StatusOK, imageFiles)
}
