package storage

import (
	"errors"
	"sync"
	"vibe-certificados/models"
)

// MemoryStorage provides in-memory storage for certificates and templates
type MemoryStorage struct {
	certificates map[string]*models.Certificate
	templates    map[string]*models.Template
	emailIndex   map[string][]string // email -> list of certificate IDs
	mutex        sync.RWMutex
}

// NewMemoryStorage creates a new in-memory storage instance
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		certificates: make(map[string]*models.Certificate),
		templates:    make(map[string]*models.Template),
		emailIndex:   make(map[string][]string),
	}
}

// SaveCertificate stores a certificate
func (ms *MemoryStorage) SaveCertificate(cert *models.Certificate) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	ms.certificates[cert.ID] = cert

	// Update email index
	if _, exists := ms.emailIndex[cert.Email]; !exists {
		ms.emailIndex[cert.Email] = make([]string, 0)
	}
	ms.emailIndex[cert.Email] = append(ms.emailIndex[cert.Email], cert.ID)

	return nil
}

// GetCertificate retrieves a certificate by ID
func (ms *MemoryStorage) GetCertificate(id string) (*models.Certificate, error) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	cert, exists := ms.certificates[id]
	if !exists {
		return nil, errors.New("certificate not found")
	}
	return cert, nil
}

// GetCertificatesByEmail retrieves all certificates for an email
func (ms *MemoryStorage) GetCertificatesByEmail(email string) ([]*models.Certificate, error) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	ids, exists := ms.emailIndex[email]
	if !exists {
		return []*models.Certificate{}, nil
	}

	certificates := make([]*models.Certificate, 0, len(ids))
	for _, id := range ids {
		if cert, exists := ms.certificates[id]; exists {
			certificates = append(certificates, cert)
		}
	}

	return certificates, nil
}

// SaveTemplate stores a template
func (ms *MemoryStorage) SaveTemplate(template *models.Template) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	ms.templates[template.ID] = template
	return nil
}

// GetTemplate retrieves a template by ID
func (ms *MemoryStorage) GetTemplate(id string) (*models.Template, error) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	template, exists := ms.templates[id]
	if !exists {
		return nil, errors.New("template not found")
	}
	return template, nil
}

// GetAllTemplates retrieves all templates
func (ms *MemoryStorage) GetAllTemplates() ([]*models.Template, error) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	templates := make([]*models.Template, 0, len(ms.templates))
	for _, template := range ms.templates {
		templates = append(templates, template)
	}
	return templates, nil
}

// DeleteTemplate removes a template
func (ms *MemoryStorage) DeleteTemplate(id string) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	if _, exists := ms.templates[id]; !exists {
		return errors.New("template not found")
	}
	delete(ms.templates, id)
	return nil
}