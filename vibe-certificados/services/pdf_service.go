package services

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"vibe-certificados/models"
	"github.com/jung-kurt/gofpdf"
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

// GeneratePDF generates a PDF from a certificate using gofpdf
func (ps *PDFService) GeneratePDF(cert *models.Certificate) ([]byte, error) {
	// Create a new PDF in landscape orientation, A4 size
	pdf := gofpdf.New("L", "mm", "A4", "")
	
	// Add a page
	pdf.AddPage()
	
	// Set margins
	pdf.SetMargins(20, 20, 20)
	
	// Add Portuguese content with character conversion
	ps.addPortugueseCertificateContent(pdf, cert)
	
	// Check for errors
	if pdf.Error() != nil {
		return nil, fmt.Errorf("PDF generation error: %v", pdf.Error())
	}
	
	// Generate PDF as bytes
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, fmt.Errorf("failed to generate PDF: %v", err)
	}
	
	return buf.Bytes(), nil
}

// addCertificateContent adds the certificate content to the PDF
func (ps *PDFService) addCertificateContent(pdf *gofpdf.Fpdf, cert *models.Certificate) error {
	// Check for errors
	if pdf.Error() != nil {
		return fmt.Errorf("PDF error before content: %v", pdf.Error())
	}
	
	// Set font for the title
	pdf.SetFont("Arial", "B", 28)
	if pdf.Error() != nil {
		return fmt.Errorf("Failed to set title font: %v", pdf.Error())
	}
	
	// Add title
	pdf.CellFormat(0, 20, "CERTIFICADO DE CONCLUSAO", "", 1, "C", false, 0, "")
	if pdf.Error() != nil {
		return fmt.Errorf("Failed to add title: %v", pdf.Error())
	}
	pdf.Ln(10)
	
	// Set font for subtitle
	pdf.SetFont("Arial", "", 16)
	pdf.CellFormat(0, 10, "Certificate of Completion", "", 1, "C", false, 0, "")
	pdf.Ln(5)
	
	// Set font for student name (larger and bold)
	pdf.SetFont("Arial", "B", 24)
	pdf.CellFormat(0, 15, cert.Name, "", 1, "C", false, 0, "")
	pdf.Ln(10)
	
	// Set font for course details
	pdf.SetFont("Arial", "", 16)
	pdf.CellFormat(0, 10, "has successfully completed the course", "", 1, "C", false, 0, "")
	pdf.Ln(5)
	
	// Course name (bold)
	pdf.SetFont("Arial", "B", 20)
	pdf.CellFormat(0, 12, cert.Course, "", 1, "C", false, 0, "")
	pdf.Ln(15)
	
	// Completion date
	pdf.SetFont("Arial", "", 14)
	completionDate := cert.CompletionDate.Format("02/01/2006")
	pdf.CellFormat(0, 8, fmt.Sprintf("Completed on: %s", completionDate), "", 1, "C", false, 0, "")
	pdf.Ln(10)
	
	// Certificate ID
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 6, fmt.Sprintf("Certificate ID: %s", cert.ID), "", 1, "C", false, 0, "")
	
	// Final error check
	if pdf.Error() != nil {
		return fmt.Errorf("PDF error after content: %v", pdf.Error())
	}
	
	return nil
}

// addSimpleCertificateContent adds simplified certificate content
func (ps *PDFService) addSimpleCertificateContent(pdf *gofpdf.Fpdf, cert *models.Certificate) {
	// Title
	pdf.SetFont("Arial", "B", 30)
	pdf.SetY(40)
	pdf.CellFormat(0, 15, "CERTIFICATE OF COMPLETION", "", 1, "C", false, 0, "")
	
	// Student name
	pdf.SetY(80)
	pdf.SetFont("Arial", "B", 24)
	pdf.CellFormat(0, 12, cert.Name, "", 1, "C", false, 0, "")
	
	// Course
	pdf.SetY(110)
	pdf.SetFont("Arial", "", 18)
	pdf.CellFormat(0, 10, "has completed the course:", "", 1, "C", false, 0, "")
	
	pdf.SetY(130)
	pdf.SetFont("Arial", "B", 20)
	pdf.CellFormat(0, 10, cert.Course, "", 1, "C", false, 0, "")
	
	// Date
	pdf.SetY(160)
	pdf.SetFont("Arial", "", 14)
	dateStr := cert.CompletionDate.Format("January 02, 2006")
	pdf.CellFormat(0, 8, "Completed: "+dateStr, "", 1, "C", false, 0, "")
}

// addPortugueseCertificateContent adds Portuguese certificate content with proper encoding
func (ps *PDFService) addPortugueseCertificateContent(pdf *gofpdf.Fpdf, cert *models.Certificate) {
	// Title - using special characters that work with gofpdf's CP1252 encoding
	pdf.SetFont("Arial", "B", 30)
	pdf.SetY(40)
	title := ps.toCP1252("CERTIFICADO DE CONCLUSÃO")
	pdf.CellFormat(0, 15, title, "", 1, "C", false, 0, "")
	
	// Subtitle
	pdf.SetY(70)
	pdf.SetFont("Arial", "", 16)
	subtitle := ps.toCP1252("Certificamos que")
	pdf.CellFormat(0, 10, subtitle, "", 1, "C", false, 0, "")
	
	// Student name (apply CP1252 conversion)
	pdf.SetY(95)
	pdf.SetFont("Arial", "B", 24)
	studentName := ps.toCP1252(cert.Name)
	pdf.CellFormat(0, 12, studentName, "", 1, "C", false, 0, "")
	
	// Course description
	pdf.SetY(125)
	pdf.SetFont("Arial", "", 18)
	courseDesc := ps.toCP1252("concluiu com êxito o curso")
	pdf.CellFormat(0, 10, courseDesc, "", 1, "C", false, 0, "")
	
	// Course name (apply CP1252 conversion)
	pdf.SetY(150)
	pdf.SetFont("Arial", "B", 20)
	courseName := ps.toCP1252(cert.Course)
	pdf.CellFormat(0, 10, courseName, "", 1, "C", false, 0, "")
	
	// Date
	pdf.SetY(175)
	pdf.SetFont("Arial", "", 14)
	dateStr := cert.CompletionDate.Format("02 de January de 2006")
	dateStr = ps.toCP1252("Concluído em: " + dateStr)
	pdf.CellFormat(0, 8, dateStr, "", 1, "C", false, 0, "")
}

// toCP1252 converts UTF-8 Portuguese characters to CP1252 encoding for gofpdf
func (ps *PDFService) toCP1252(text string) string {
	// gofpdf supports CP1252 encoding, which includes Portuguese characters
	// Map UTF-8 characters to their CP1252 equivalents
	result := strings.NewReplacer(
		// Portuguese accented characters to CP1252
		"ã", "\xe3", "á", "\xe1", "à", "\xe0", "â", "\xe2",
		"é", "\xe9", "ê", "\xea", "è", "\xe8",
		"í", "\xed", "ì", "\xec", "î", "\xee",
		"ó", "\xf3", "ô", "\xf4", "ò", "\xf2", "õ", "\xf5",
		"ú", "\xfa", "ù", "\xf9", "û", "\xfb",
		"ç", "\xe7",
		
		// Uppercase
		"Ã", "\xc3", "Á", "\xc1", "À", "\xc0", "Â", "\xc2",
		"É", "\xc9", "Ê", "\xca", "È", "\xc8",
		"Í", "\xcd", "Ì", "\xcc", "Î", "\xce",
		"Ó", "\xd3", "Ô", "\xd4", "Ò", "\xd2", "Õ", "\xd5",
		"Ú", "\xda", "Ù", "\xd9", "Û", "\xdb",
		"Ç", "\xc7",
	).Replace(text)
	
	return result
}

// convertWithWkhtmltopdf uses wkhtmltopdf for conversion (if available)
// This method is kept for backward compatibility but not used in the main flow
func (ps *PDFService) convertWithWkhtmltopdf(html string) ([]byte, error) {
	cmd := exec.Command("wkhtmltopdf", "--page-size", "A4", "--orientation", "Landscape", "-", "-")
	cmd.Stdin = bytes.NewReader([]byte(html))
	
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("wkhtmltopdf conversion failed: %v", err)
	}
	
	return output, nil
}