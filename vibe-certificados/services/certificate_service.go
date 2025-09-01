package services

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"
	"vibe-certificados/models"
	"vibe-certificados/storage"
)

// CertificateService handles certificate-related operations
type CertificateService struct {
	storage *storage.MemoryStorage
}

// NewCertificateService creates a new certificate service
func NewCertificateService(storage *storage.MemoryStorage) *CertificateService {
	return &CertificateService{
		storage: storage,
	}
}

// CreateCertificate creates a new certificate from a request
func (cs *CertificateService) CreateCertificate(req *models.CertificateRequest) (*models.Certificate, error) {
	// Parse completion date
	completionDate, err := time.Parse("2006-01-02", req.CompletionDate)
	if err != nil {
		return nil, errors.New("invalid completion_date format. Use YYYY-MM-DD")
	}

	// Use default template if not specified
	templateID := req.TemplateID
	if templateID == "" {
		templateID = "default"
	}

	// Verify template exists
	_, err = cs.storage.GetTemplate(templateID)
	if err != nil {
		return nil, errors.New("template not found: " + templateID)
	}

	// Create certificate
	cert := models.NewCertificate(
		req.Email,
		req.Name,
		req.Course,
		templateID,
		completionDate,
		req.Data,
	)

	// Save certificate
	err = cs.storage.SaveCertificate(cert)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

// GetCertificate retrieves a certificate by ID
func (cs *CertificateService) GetCertificate(id string) (*models.Certificate, error) {
	return cs.storage.GetCertificate(id)
}

// GetCertificatesByEmail retrieves all certificates for an email
func (cs *CertificateService) GetCertificatesByEmail(email string) ([]*models.Certificate, error) {
	return cs.storage.GetCertificatesByEmail(email)
}

// CreateCertificatesFromCSV creates multiple certificates from CSV data
func (cs *CertificateService) CreateCertificatesFromCSV(csvData io.Reader) (*models.BatchCertificateResponse, error) {
	response := &models.BatchCertificateResponse{
		CreatedIDs: make([]string, 0),
		Errors:     make([]string, 0),
	}

	reader := csv.NewReader(csvData)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("failed to parse CSV: " + err.Error())
	}

	if len(records) == 0 {
		return nil, errors.New("CSV file is empty")
	}

	// Assume first row is header
	headers := records[0]
	response.Total = len(records) - 1

	// Find required column indices
	emailIdx, nameIdx, courseIdx, dateIdx := -1, -1, -1, -1
	templateIdx := -1

	for i, header := range headers {
		switch strings.ToLower(strings.TrimSpace(header)) {
		case "email":
			emailIdx = i
		case "name":
			nameIdx = i
		case "course":
			courseIdx = i
		case "completion_date", "date":
			dateIdx = i
		case "template_id", "template":
			templateIdx = i
		}
	}

	if emailIdx == -1 || nameIdx == -1 || courseIdx == -1 || dateIdx == -1 {
		return nil, errors.New("CSV must contain email, name, course, and completion_date columns")
	}

	// Process each record
	for i := 1; i < len(records); i++ {
		record := records[i]
		rowNum := strconv.Itoa(i + 1)

		if len(record) <= emailIdx || len(record) <= nameIdx || len(record) <= courseIdx || len(record) <= dateIdx {
			response.Failed++
			response.Errors = append(response.Errors, "Row "+rowNum+": insufficient columns")
			continue
		}

		templateID := "default"
		if templateIdx >= 0 && len(record) > templateIdx && record[templateIdx] != "" {
			templateID = record[templateIdx]
		}

		req := &models.CertificateRequest{
			Email:          strings.TrimSpace(record[emailIdx]),
			Name:           strings.TrimSpace(record[nameIdx]),
			Course:         strings.TrimSpace(record[courseIdx]),
			CompletionDate: strings.TrimSpace(record[dateIdx]),
			TemplateID:     templateID,
		}

		cert, err := cs.CreateCertificate(req)
		if err != nil {
			response.Failed++
			response.Errors = append(response.Errors, "Row "+rowNum+": "+err.Error())
		} else {
			response.Success++
			response.CreatedIDs = append(response.CreatedIDs, cert.ID)
		}
	}

	return response, nil
}