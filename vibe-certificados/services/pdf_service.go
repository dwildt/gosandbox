package services

import (
	"bytes"
	"fmt"
	"os/exec"
	"vibe-certificados/models"
)

// PDFService handles PDF generation from HTML
type PDFService struct {
	templateService *TemplateService
}

// NewPDFService creates a new PDF service
func NewPDFService(templateService *TemplateService) *PDFService {
	return &PDFService{
		templateService: templateService,
	}
}

// GeneratePDF generates a PDF from a certificate
// Note: This is a simplified implementation that converts HTML to PDF
// In production, you might want to use libraries like wkhtmltopdf, chromedp, or go-wkhtmltopdf
func (ps *PDFService) GeneratePDF(cert *models.Certificate) ([]byte, error) {
	// First, render the HTML
	html, err := ps.templateService.RenderCertificate(cert)
	if err != nil {
		return nil, err
	}

	// For this implementation, we'll create a simple HTML-to-PDF conversion
	// In a real-world scenario, you would use a proper PDF generation library
	return ps.convertHTMLToPDF(html)
}

// convertHTMLToPDF is a simplified PDF conversion
// In production, replace this with a proper PDF library
func (ps *PDFService) convertHTMLToPDF(html string) ([]byte, error) {
	// This is a placeholder implementation
	// For a real implementation, you could use:
	// 1. wkhtmltopdf via exec
	// 2. chromedp for headless Chrome
	// 3. go-wkhtmltopdf wrapper
	// 4. gofpdf for pure Go PDF generation
	
	// Check if wkhtmltopdf is available (optional)
	if ps.isWkhtmltopdfAvailable() {
		return ps.convertWithWkhtmltopdf(html)
	}
	
	// Fallback: return the HTML content with PDF headers
	// This is not a real PDF but demonstrates the concept
	pdfContent := fmt.Sprintf(`%%PDF-1.4
1 0 obj
<<
/Type /Catalog
/Pages 2 0 R
>>
endobj

2 0 obj
<<
/Type /Pages
/Kids [3 0 R]
/Count 1
>>
endobj

3 0 obj
<<
/Type /Page
/Parent 2 0 R
/Contents 4 0 R
>>
endobj

4 0 obj
<<
/Length %d
>>
stream
%s
endstream
endobj

xref
0 5
0000000000 65535 f 
0000000009 00000 n 
0000000074 00000 n 
0000000120 00000 n 
0000000179 00000 n 
trailer
<<
/Size 5
/Root 1 0 R
>>
startxref
%d
%%%%EOF`, len(html), html, 240+len(html))

	return []byte(pdfContent), nil
}

// isWkhtmltopdfAvailable checks if wkhtmltopdf is installed
func (ps *PDFService) isWkhtmltopdfAvailable() bool {
	_, err := exec.LookPath("wkhtmltopdf")
	return err == nil
}

// convertWithWkhtmltopdf uses wkhtmltopdf for conversion (if available)
func (ps *PDFService) convertWithWkhtmltopdf(html string) ([]byte, error) {
	cmd := exec.Command("wkhtmltopdf", "--page-size", "A4", "--orientation", "Landscape", "-", "-")
	cmd.Stdin = bytes.NewReader([]byte(html))
	
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("wkhtmltopdf conversion failed: %v", err)
	}
	
	return output, nil
}