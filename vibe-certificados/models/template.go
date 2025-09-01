package models

// Template represents a certificate template
type Template struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	HTMLTemplate string          `json:"html_template"`
	Fields       []TemplateField `json:"fields"`
	CreatedAt    string          `json:"created_at"`
	UpdatedAt    string          `json:"updated_at"`
}

// TemplateField represents a field definition in a template
type TemplateField struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Required    bool        `json:"required"`
	Default     interface{} `json:"default,omitempty"`
	Description string      `json:"description,omitempty"`
}

// CertificateRequest represents a request to generate a certificate
type CertificateRequest struct {
	Email          string            `json:"email" binding:"required"`
	Name           string            `json:"name" binding:"required"`
	Course         string            `json:"course" binding:"required"`
	CompletionDate string            `json:"completion_date" binding:"required"`
	TemplateID     string            `json:"template_id"`
	Data           map[string]string `json:"data,omitempty"`
}

// BatchCertificateRequest represents the response for batch creation
type BatchCertificateResponse struct {
	Total      int      `json:"total"`
	Success    int      `json:"success"`
	Failed     int      `json:"failed"`
	Errors     []string `json:"errors,omitempty"`
	CreatedIDs []string `json:"created_ids"`
}