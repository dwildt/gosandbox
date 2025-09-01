package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine, handlers *Handlers) {
	// API group
	api := r.Group("/api")

	// Certificate routes
	certificates := api.Group("/certificates")
	{
		certificates.POST("", handlers.CreateCertificate)
		certificates.POST("/batch", handlers.CreateCertificatesBatch)
		certificates.GET("/:id", handlers.GetCertificateByFormat) // Handle both .html and .pdf
		certificates.GET("/by-email/:email", handlers.GetCertificatesByEmail)
	}

	// Template routes
	templates := api.Group("/templates")
	{
		templates.GET("", handlers.GetTemplates)
		templates.POST("", handlers.CreateTemplate)
		templates.GET("/:id", handlers.GetTemplate)
		templates.PUT("/:id", handlers.UpdateTemplate)
		templates.DELETE("/:id", handlers.DeleteTemplate)
	}

	// Health check
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "vibe-certificados",
		})
	})
}