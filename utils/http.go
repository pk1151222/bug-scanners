// internal/utils/http.go

package utils

import (
	"net/http"
)

// डोमेन के HTTP स्थिति कोड की जांच करने का फ़ंक्शन।
func CheckHTTPStatus(domain string) (int, error) {
	resp, err := http.Get("http://" + domain)
	if err != nil {
		return 0, fmt.Errorf("HTTP अनुरोध भेजने में त्रुटि: %v", err)
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
