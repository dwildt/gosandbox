package models_test

import (
	"testing"
	"time"
	"vibe-certificados/models"
)

func TestNewCertificate(t *testing.T) {
	email := "test@example.com"
	name := "João Silva"
	course := "Go Programming"
	templateID := "default"
	completionDate := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	additionalData := map[string]string{"instructor": "Prof. Silva"}

	cert := models.NewCertificate(email, name, course, templateID, completionDate, additionalData)

	// Test basic fields
	if cert.Email != email {
		t.Errorf("Expected email %s, got %s", email, cert.Email)
	}
	if cert.Name != name {
		t.Errorf("Expected name %s, got %s", name, cert.Name)
	}
	if cert.Course != course {
		t.Errorf("Expected course %s, got %s", course, cert.Course)
	}
	if cert.TemplateID != templateID {
		t.Errorf("Expected template ID %s, got %s", templateID, cert.TemplateID)
	}
	if !cert.CompletionDate.Equal(completionDate) {
		t.Errorf("Expected completion date %v, got %v", completionDate, cert.CompletionDate)
	}

	// Test UUID generation
	if cert.ID == "" {
		t.Error("Certificate ID should not be empty")
	}
	if len(cert.ID) != 36 { // UUID v4 format
		t.Errorf("Certificate ID should be 36 characters, got %d", len(cert.ID))
	}

	// Test additional data
	if cert.Data["instructor"] != "Prof. Silva" {
		t.Errorf("Expected instructor 'Prof. Silva', got %s", cert.Data["instructor"])
	}

	// Test created at
	if cert.CreatedAt.IsZero() {
		t.Error("CreatedAt should not be zero")
	}
}

func TestGetAllData(t *testing.T) {
	cert := models.NewCertificate(
		"test@example.com",
		"João Silva",
		"Go Programming",
		"default",
		time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
		map[string]string{"instructor": "Prof. Silva"},
	)

	data := cert.GetAllData()

	// Check standard fields
	if data["Email"] != "test@example.com" {
		t.Errorf("Expected email in data")
	}
	if data["Name"] != "João Silva" {
		t.Errorf("Expected name in data")
	}
	if data["Course"] != "Go Programming" {
		t.Errorf("Expected course in data")
	}
	if data["CompletionDate"] != "15/01/2024" {
		t.Errorf("Expected formatted completion date, got %s", data["CompletionDate"])
	}

	// Check custom data
	if data["instructor"] != "Prof. Silva" {
		t.Errorf("Expected instructor in data")
	}
}