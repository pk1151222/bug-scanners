// internal/subdomain/subdomain_enumerator.go

package subdomain

import (
	"fmt"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

// डोमेन के लिए सबडोमेन खोजने का फ़ंक्शन।
func EnumerateSubdomains(domain string) ([]string, error) {
	var subdomains []string
	options := runner.Options{Domain: domain}
	subfinderRunner, err := runner.NewRunner(&options)
	if err != nil {
		return nil, fmt.Errorf("subfinder रनर बनाने में त्रुटि: %v", err)
	}

	results := subfinderRunner.RunEnumeration()
	return results, nil
}
