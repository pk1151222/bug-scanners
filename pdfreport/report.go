package pdfreport

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

// GeneratePDF generates a PDF report of the subdomains and TLS details
func GeneratePDF(subdomains []string, tlsDetails *TLSDetails) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Bug Scanner Report")

	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)

	pdf.Cell(40, 10, fmt.Sprintf("Subdomains found:"))
	for _, subdomain := range subdomains {
		pdf.Cell(40, 10, subdomain)
		pdf.Ln(6)
	}

	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("TLS Details:"))
	pdf.Cell(40, 10, fmt.Sprintf("Supported Versions: %v", tlsDetails.SupportedVersions))
	pdf.Ln(6)
	pdf.Cell(40, 10, fmt.Sprintf("Cipher Suites: %v", tlsDetails.CipherSuites))

	return pdf.OutputFileAndClose("report.pdf")
}
