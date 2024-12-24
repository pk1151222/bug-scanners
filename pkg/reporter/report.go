package reporter

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"github.com/pk1151222/bug-scanners/internal/models"
)

// GeneratePDF generates a PDF report of the scan results
func GeneratePDF(scanResult models.ScanResult) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Bug Scanner Report")

	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("URL: %s", scanResult.URL))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Subdomains: %v", scanResult.Subdomains))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("CVEs: %v", scanResult.CVEs))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Security Headers: %v", scanResult.SecurityHeaders))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("SQL Vulnerabilities: %v", scanResult.SQLVulnerabilities))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("WAF Status: %s", scanResult.WAFStatus))

	// Add more content here...

	return pdf.OutputFileAndClose("scan_report.pdf")
}
