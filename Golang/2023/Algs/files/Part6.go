package algs

import (
	"fmt"
	"strings"
)

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		locale = "en"
	}
	if strings.HasPrefix(domain, locale) {
		return domain
	} else {
		return locale + "." + domain
	}
}

func main() {
	fmt.Println(DomainForLocale("example.com", ""))
}
