// internal/subdomain/subdomain_enumerator.go

package subdomain

import (
    "fmt"
    "github.com/pk1151222/bug-scanners/internal/utils"
    "github.com/projectdiscovery/subfinder/v2/pkg/subscraping"
)

func EnhancedScanSubdomains(domain string) {
    subdomains := utils.EnumerateSubdomains(domain)

    thirdPartySubdomains, err := subscraping.GetSubdomainsFromShodan(domain)
    if err != nil {
        fmt.Println("Error fetching subdomains from Shodan:", err)
    }

    subdomains = append(subdomains, thirdPartySubdomains...)

    fmt.Println("Discovered subdomains:")
    for _, subdomain := range subdomains {
        fmt.Println(subdomain)
    }
}
