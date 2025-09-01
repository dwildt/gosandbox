package services

import (
	"bytes"
	"errors"
	"html/template"
	"time"
	"vibe-certificados/models"
	"vibe-certificados/storage"
)

// TemplateService handles template-related operations
type TemplateService struct {
	storage *storage.MemoryStorage
}

// NewTemplateService creates a new template service
func NewTemplateService(storage *storage.MemoryStorage) *TemplateService {
	ts := &TemplateService{
		storage: storage,
	}
	
	// Initialize with default template
	ts.initializeDefaultTemplate()
	
	return ts
}

// initializeDefaultTemplate creates the default template
func (ts *TemplateService) initializeDefaultTemplate() {
	defaultTemplate := &models.Template{
		ID:   "default",
		Name: "Default Certificate Template",
		HTMLTemplate: `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Certificate</title>
    <style>
        body {
            font-family: 'Georgia', serif;
            margin: 0;
            padding: 40px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: #333;
        }
        .certificate {
            background: white;
            max-width: 800px;
            margin: 0 auto;
            padding: 60px;
            border-radius: 15px;
            box-shadow: 0 10px 30px rgba(0,0,0,0.3);
            text-align: center;
        }
        .header {
            border-bottom: 3px solid #764ba2;
            padding-bottom: 20px;
            margin-bottom: 40px;
        }
        .title {
            font-size: 48px;
            color: #764ba2;
            margin: 0;
            text-transform: uppercase;
            letter-spacing: 2px;
        }
        .subtitle {
            font-size: 18px;
            color: #666;
            margin: 10px 0 0 0;
        }
        .content {
            margin: 40px 0;
            line-height: 1.8;
        }
        .recipient {
            font-size: 36px;
            color: #333;
            margin: 20px 0;
            text-decoration: underline;
        }
        .course {
            font-size: 24px;
            color: #764ba2;
            font-weight: bold;
            margin: 20px 0;
        }
        .date {
            font-size: 16px;
            color: #666;
            margin-top: 40px;
        }
        .footer {
            margin-top: 60px;
            border-top: 2px solid #764ba2;
            padding-top: 20px;
        }
        .certificate-id {
            font-size: 12px;
            color: #999;
            margin-top: 20px;
        }
    </style>
</head>
<body>
    <div class="certificate">
        <div class="header">
            <h1 class="title">Certificado</h1>
            <p class="subtitle">Certificate of Completion</p>
        </div>
        
        <div class="content">
            <p>Certificamos que</p>
            <div class="recipient">{{.Name}}</div>
            <p>concluiu com êxito o curso</p>
            <div class="course">{{.Course}}</div>
            <p>demonstrando conhecimento e dedicação ao aprendizado.</p>
        </div>
        
        <div class="footer">
            <div class="date">
                Concluído em: {{.CompletionDate}}<br>
                Emitido em: {{.CreatedAt}}
            </div>
            <div class="certificate-id">
                ID do Certificado: {{.ID}}
            </div>
        </div>
    </div>
</body>
</html>`,
		Fields: []models.TemplateField{
			{Name: "name", Type: "string", Required: true, Description: "Nome do certificado"},
			{Name: "course", Type: "string", Required: true, Description: "Nome do curso"},
			{Name: "completion_date", Type: "date", Required: true, Description: "Data de conclusão (YYYY-MM-DD)"},
		},
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	
	ts.storage.SaveTemplate(defaultTemplate)
}

// GetTemplate retrieves a template by ID
func (ts *TemplateService) GetTemplate(id string) (*models.Template, error) {
	return ts.storage.GetTemplate(id)
}

// GetAllTemplates retrieves all templates
func (ts *TemplateService) GetAllTemplates() ([]*models.Template, error) {
	return ts.storage.GetAllTemplates()
}

// CreateTemplate creates a new template
func (ts *TemplateService) CreateTemplate(template *models.Template) error {
	template.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	template.UpdatedAt = template.CreatedAt
	return ts.storage.SaveTemplate(template)
}

// UpdateTemplate updates an existing template
func (ts *TemplateService) UpdateTemplate(template *models.Template) error {
	// Check if template exists
	_, err := ts.storage.GetTemplate(template.ID)
	if err != nil {
		return err
	}
	
	template.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	return ts.storage.SaveTemplate(template)
}

// DeleteTemplate removes a template
func (ts *TemplateService) DeleteTemplate(id string) error {
	// Don't allow deletion of default template
	if id == "default" {
		return errors.New("cannot delete default template")
	}
	return ts.storage.DeleteTemplate(id)
}

// RenderCertificate renders a certificate using its template
func (ts *TemplateService) RenderCertificate(cert *models.Certificate) (string, error) {
	tmpl, err := ts.storage.GetTemplate(cert.TemplateID)
	if err != nil {
		return "", err
	}

	// Parse template
	t, err := template.New("certificate").Parse(tmpl.HTMLTemplate)
	if err != nil {
		return "", err
	}

	// Render with certificate data
	var buf bytes.Buffer
	err = t.Execute(&buf, cert.GetAllData())
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}