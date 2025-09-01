package api

import (
	"net/http"
	"strings"
	"vibe-certificados/models"
	"vibe-certificados/services"

	"github.com/gin-gonic/gin"
)

// Handlers contains all HTTP handlers
type Handlers struct {
	certificateService *services.CertificateService
	templateService    *services.TemplateService
	pdfService         *services.PDFService
}

// NewHandlers creates a new handlers instance
func NewHandlers(certService *services.CertificateService, templateService *services.TemplateService, pdfService *services.PDFService) *Handlers {
	return &Handlers{
		certificateService: certService,
		templateService:    templateService,
		pdfService:         pdfService,
	}
}

// CreateCertificate handles POST /api/certificates
func (h *Handlers) CreateCertificate(c *gin.Context) {
	var req models.CertificateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cert, err := h.certificateService.CreateCertificate(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cert)
}

// CreateCertificatesBatch handles POST /api/certificates/batch
func (h *Handlers) CreateCertificatesBatch(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CSV file is required"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	response, err := h.certificateService.CreateCertificatesFromCSV(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetCertificateByFormat handles both HTML and PDF export based on file extension
func (h *Handlers) GetCertificateByFormat(c *gin.Context) {
	idParam := c.Param("id")
	
	// Check if it ends with .html or .pdf
	if strings.HasSuffix(idParam, ".html") {
		id := strings.TrimSuffix(idParam, ".html")
		h.serveCertificateHTML(c, id)
	} else if strings.HasSuffix(idParam, ".pdf") {
		id := strings.TrimSuffix(idParam, ".pdf")
		h.serveCertificatePDF(c, id)
	} else {
		// Default to JSON response with certificate data
		h.serveCertificateJSON(c, idParam)
	}
}

// serveCertificateJSON returns certificate as JSON
func (h *Handlers) serveCertificateJSON(c *gin.Context, id string) {
	cert, err := h.certificateService.GetCertificate(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Certificate not found"})
		return
	}
	c.JSON(http.StatusOK, cert)
}

// serveCertificateHTML serves certificate as HTML
func (h *Handlers) serveCertificateHTML(c *gin.Context, id string) {
	cert, err := h.certificateService.GetCertificate(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Certificate not found"})
		return
	}

	html, err := h.templateService.RenderCertificate(cert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render certificate"})
		return
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, html)
}

// serveCertificatePDF serves certificate as PDF
func (h *Handlers) serveCertificatePDF(c *gin.Context, id string) {
	cert, err := h.certificateService.GetCertificate(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Certificate not found"})
		return
	}

	pdf, err := h.pdfService.GeneratePDF(cert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "inline; filename=certificate_"+cert.ID+".pdf")
	c.Data(http.StatusOK, "application/pdf", pdf)
}

// GetCertificatesByEmail handles GET /api/certificates/by-email/{email}
func (h *Handlers) GetCertificatesByEmail(c *gin.Context) {
	email := c.Param("email")

	certificates, err := h.certificateService.GetCertificatesByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email":        email,
		"count":        len(certificates),
		"certificates": certificates,
	})
}

// GetTemplates handles GET /api/templates
func (h *Handlers) GetTemplates(c *gin.Context) {
	templates, err := h.templateService.GetAllTemplates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, templates)
}

// GetTemplate handles GET /api/templates/{id}
func (h *Handlers) GetTemplate(c *gin.Context) {
	id := c.Param("id")

	template, err := h.templateService.GetTemplate(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Template not found"})
		return
	}

	c.JSON(http.StatusOK, template)
}

// CreateTemplate handles POST /api/templates
func (h *Handlers) CreateTemplate(c *gin.Context) {
	var template models.Template
	if err := c.ShouldBindJSON(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.templateService.CreateTemplate(&template)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, template)
}

// UpdateTemplate handles PUT /api/templates/{id}
func (h *Handlers) UpdateTemplate(c *gin.Context) {
	id := c.Param("id")

	var template models.Template
	if err := c.ShouldBindJSON(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	template.ID = id
	err := h.templateService.UpdateTemplate(&template)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, template)
}

// DeleteTemplate handles DELETE /api/templates/{id}
func (h *Handlers) DeleteTemplate(c *gin.Context) {
	id := c.Param("id")

	err := h.templateService.DeleteTemplate(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Template deleted successfully"})
}