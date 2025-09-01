package models

import (
	"time"

	"github.com/google/uuid"
)

// Certificate represents a generated certificate
type Certificate struct {
	ID             string            `json:"id"`
	Email          string            `json:"email"`
	Name           string            `json:"name"`
	Course         string            `json:"course"`
	CompletionDate time.Time         `json:"completion_date"`
	TemplateID     string            `json:"template_id"`
	CreatedAt      time.Time         `json:"created_at"`
	Data           map[string]string `json:"data,omitempty"`
}

// NewCertificate creates a new certificate with a unique UUID
func NewCertificate(email, name, course, templateID string, completionDate time.Time, additionalData map[string]string) *Certificate {
	cert := &Certificate{
		ID:             uuid.New().String(),
		Email:          email,
		Name:           name,
		Course:         course,
		CompletionDate: completionDate,
		TemplateID:     templateID,
		CreatedAt:      time.Now(),
		Data:           make(map[string]string),
	}

	// Add additional data if provided
	if additionalData != nil {
		for k, v := range additionalData {
			cert.Data[k] = v
		}
	}

	return cert
}

// GetAllData returns all certificate data including standard fields
func (c *Certificate) GetAllData() map[string]interface{} {
	data := make(map[string]interface{})
	data["ID"] = c.ID
	data["Email"] = c.Email
	data["Name"] = c.Name
	data["Course"] = c.Course
	data["CompletionDate"] = c.CompletionDate.Format("02/01/2006")
	data["CreatedAt"] = c.CreatedAt.Format("02/01/2006 15:04:05")

	// Add custom data
	for k, v := range c.Data {
		data[k] = v
	}

	return data
}