package main

import (
	"log"
	"vibe-certificados/api"
	"vibe-certificados/services"
	"vibe-certificados/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize storage
	memoryStorage := storage.NewMemoryStorage()

	// Initialize services
	templateService := services.NewTemplateService(memoryStorage)
	certificateService := services.NewCertificateService(memoryStorage)
	pdfService := services.NewPDFService(templateService)

	// Initialize handlers
	handlers := api.NewHandlers(certificateService, templateService, pdfService)

	// Setup Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Setup routes
	api.SetupRoutes(r, handlers)

	// Add root endpoint with API documentation
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service":     "Vibe Certificados API",
			"version":     "1.0.0",
			"description": "API para geração de certificados em HTML e PDF",
			"endpoints": map[string]interface{}{
				"health": "/api/health",
				"certificates": map[string]string{
					"create":       "POST /api/certificates",
					"batch":        "POST /api/certificates/batch",
					"html":         "GET /api/certificates/{id}.html",
					"pdf":          "GET /api/certificates/{id}.pdf",
					"by_email":     "GET /api/certificates/by-email/{email}",
				},
				"templates": map[string]string{
					"list":   "GET /api/templates",
					"create": "POST /api/templates",
					"get":    "GET /api/templates/{id}",
					"update": "PUT /api/templates/{id}",
					"delete": "DELETE /api/templates/{id}",
				},
			},
			"documentation": "https://github.com/dwildt/gosandbox/tree/main/vibe-certificados",
		})
	})

	// Start server
	log.Println("Starting Vibe Certificados API on :8080")
	log.Println("API Documentation: http://localhost:8080")
	log.Println("Health Check: http://localhost:8080/api/health")
	
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}