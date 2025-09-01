package services_test

import (
	"strings"
	"testing"
	"vibe-certificados/models"
	"vibe-certificados/services"
	"vibe-certificados/storage"
)

func TestCertificateService_CreateCertificate(t *testing.T) {
	// Setup
	memStorage := storage.NewMemoryStorage()
	_ = services.NewTemplateService(memStorage) // Initialize templates
	certService := services.NewCertificateService(memStorage)

	// Test valid request
	req := &models.CertificateRequest{
		Email:          "test@example.com",
		Name:           "João Silva",
		Course:         "Go Programming",
		CompletionDate: "2024-01-15",
		TemplateID:     "default",
	}

	cert, err := certService.CreateCertificate(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cert.Email != req.Email {
		t.Errorf("Expected email %s, got %s", req.Email, cert.Email)
	}
	if cert.Name != req.Name {
		t.Errorf("Expected name %s, got %s", req.Name, cert.Name)
	}
	if cert.Course != req.Course {
		t.Errorf("Expected course %s, got %s", req.Course, cert.Course)
	}
	if cert.TemplateID != req.TemplateID {
		t.Errorf("Expected template ID %s, got %s", req.TemplateID, cert.TemplateID)
	}

	// Test invalid date format
	invalidReq := &models.CertificateRequest{
		Email:          "test@example.com",
		Name:           "João Silva",
		Course:         "Go Programming",
		CompletionDate: "invalid-date",
		TemplateID:     "default",
	}

	_, err = certService.CreateCertificate(invalidReq)
	if err == nil {
		t.Error("Expected error for invalid date format")
	}
}

func TestCertificateService_GetCertificate(t *testing.T) {
	// Setup
	memStorage := storage.NewMemoryStorage()
	_ = services.NewTemplateService(memStorage) // Initialize templates
	certService := services.NewCertificateService(memStorage)

	// Create a certificate
	req := &models.CertificateRequest{
		Email:          "test@example.com",
		Name:           "João Silva",
		Course:         "Go Programming",
		CompletionDate: "2024-01-15",
	}

	cert, err := certService.CreateCertificate(req)
	if err != nil {
		t.Fatalf("Failed to create certificate: %v", err)
	}

	// Get the certificate
	retrieved, err := certService.GetCertificate(cert.ID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if retrieved.ID != cert.ID {
		t.Errorf("Expected ID %s, got %s", cert.ID, retrieved.ID)
	}

	// Test non-existent certificate
	_, err = certService.GetCertificate("non-existent")
	if err == nil {
		t.Error("Expected error for non-existent certificate")
	}
}

func TestCertificateService_GetCertificatesByEmail(t *testing.T) {
	// Setup
	memStorage := storage.NewMemoryStorage()
	_ = services.NewTemplateService(memStorage) // Initialize templates
	certService := services.NewCertificateService(memStorage)

	email := "test@example.com"

	// Create multiple certificates for the same email
	for i := 0; i < 3; i++ {
		req := &models.CertificateRequest{
			Email:          email,
			Name:           "João Silva",
			Course:         "Course " + string(rune('A'+i)),
			CompletionDate: "2024-01-15",
		}

		_, err := certService.CreateCertificate(req)
		if err != nil {
			t.Fatalf("Failed to create certificate %d: %v", i, err)
		}
	}

	// Get certificates by email
	certificates, err := certService.GetCertificatesByEmail(email)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(certificates) != 3 {
		t.Errorf("Expected 3 certificates, got %d", len(certificates))
	}

	// Test empty email
	emptyCertificates, err := certService.GetCertificatesByEmail("empty@example.com")
	if err != nil {
		t.Fatalf("Expected no error for empty email, got %v", err)
	}

	if len(emptyCertificates) != 0 {
		t.Errorf("Expected 0 certificates for empty email, got %d", len(emptyCertificates))
	}
}

func TestCertificateService_CreateCertificatesFromCSV(t *testing.T) {
	// Setup
	memStorage := storage.NewMemoryStorage()
	_ = services.NewTemplateService(memStorage) // Initialize templates
	certService := services.NewCertificateService(memStorage)

	// Test CSV data
	csvData := `email,name,course,completion_date
test1@example.com,João Silva,Go Programming,2024-01-15
test2@example.com,Maria Santos,Web Development,2024-01-20
test3@example.com,Pedro Oliveira,Database Design,2024-01-25`

	reader := strings.NewReader(csvData)

	response, err := certService.CreateCertificatesFromCSV(reader)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Total != 3 {
		t.Errorf("Expected total 3, got %d", response.Total)
	}
	if response.Success != 3 {
		t.Errorf("Expected success 3, got %d", response.Success)
	}
	if response.Failed != 0 {
		t.Errorf("Expected failed 0, got %d", response.Failed)
	}
	if len(response.CreatedIDs) != 3 {
		t.Errorf("Expected 3 created IDs, got %d", len(response.CreatedIDs))
	}

	// Test invalid CSV
	invalidCSV := `email,name
test@example.com,João Silva`

	invalidReader := strings.NewReader(invalidCSV)
	_, err = certService.CreateCertificatesFromCSV(invalidReader)
	if err == nil {
		t.Error("Expected error for invalid CSV format")
	}
}