package security_headers

import (
	"net/http"
	"fmt"
)

func CheckSecurityHeaders(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, ok := resp.Header["X-Content-Type-Options"]; !ok {
		fmt.Println("Missing X-Content-Type-Options header")
	}
	// Check for more headers here
	return nil
}
