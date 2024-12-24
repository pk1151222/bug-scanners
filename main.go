package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pk1151222/bug-scanners/cmd"
	"github.com/pk1151222/bug-scanners/internal/models"
	"github.com/pk1151222/bug-scanners/internal/scanner"
	"github.com/pk1151222/bug-scanners/internal/subdomain"
	"github.com/pk1151222/bug-scanners/internal/utils"
	"github.com/pk1151222/bug-scanners/pkg/cve_scanner"
	"github.com/pk1151222/bug-scanners/pkg/reporter"
	"github.com/pk1151222/bug-scanners/pkg/security_headers"
	"github.com/pk1151222/bug-scanners/pkg/sql_scanner"
	"github.com/pk1151222/bug-scanners/pkg/waf_detector"
)

func main() {
	// Initialize a Gin router for the web-based interface
	router := gin.Default()

	// Load HTML templates for dashboard
	router.LoadHTMLGlob("templates/*")

	// Root route: Welcome message
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Bug Scanners API. Use /scan/:url to start scanning.",
		})
	})

	// Scan route: Performs comprehensive scanning for a given URL
	router.GET("/scan/:url", func(c *gin.Context) {
		url := c.Param("url")

		// Subdomain enumeration
		fmt.Println("Starting subdomain enumeration for:", url)
		subdomains := subdomain.EnumerateSubdomains(url)

		// CVE scanning
		fmt.Println("Starting CVE scanning for:", url)
		cves := cve_scanner.ScanCVE(url)

		// Security headers analysis
		fmt.Println("Analyzing security headers for:", url)
		headers := security_headers.AnalyzeSecurityHeaders(url)

		// SQL Injection detection
		fmt.Println("Scanning for SQL Injection vulnerabilities...")
		sqlVulnerabilities, err := sql_scanner.ScanForSQLInjection(url)
		if err != nil {
			fmt.Println("Error during SQL Injection scan:", err)
		}

		// WAF detection
		fmt.Println("Detecting WAF for:", url)
		wafStatus := waf_detector.DetectWAF(url)

		// Compile scan results
		scanResult := models.ScanResult{
			URL:                 url,
			Subdomains:          subdomains,
			CVEs:                cves,
			SecurityHeaders:     headers,
			SQLVulnerabilities:  sqlVulnerabilities,
			WAFStatus:           wafStatus,
			HTTPVulnerabilities: []models.Vulnerability{}, // Placeholder
			SSLVulnerabilities:  []models.Vulnerability{}, // Placeholder
		}

		// Print results for logging
		fmt.Println("Scan Result:", scanResult)

		// Generate a PDF report
		if err := reporter.GeneratePDF(scanResult); err != nil {
			log.Printf("Failed to generate PDF report: %v", err)
		}

		// Return JSON response
		c.JSON(200, gin.H{
			"scan_result": scanResult,
		})
	})

	// Dashboard route: Simple HTML-based dashboard
	router.GET("/dashboard", func(c *gin.Context) {
		c.HTML(200, "dashboard.html", gin.H{
			"title": "Bug Scanner Dashboard",
			"features": []string{
				"Subdomain Enumeration",
				"CVE Scanning",
				"Security Header Analysis",
				"SQL Injection Detection",
				"WAF Detection",
			},
		})
	})

	// Start the web server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
