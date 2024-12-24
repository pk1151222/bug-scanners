package waf_detector

import (
	"net/http"
	"strings"
)

func DetectWAF(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "Error fetching URL"
	}
	defer resp.Body.Close()

	wafHeaders := []string{"X-WAF", "X-EdgeProtect", "X-Akamai-Edge"}
	for _, header := range wafHeaders {
		if value := resp.Header.Get(header); value != "" {
			return fmt.Sprintf("WAF Detected: %s", value)
		}
	}
	return "No WAF detected"
}
