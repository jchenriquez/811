package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func subdomainVisits(cpdomains []string) []string {
	spaceRegex := regexp.MustCompile(` `)
	periodRegex := regexp.MustCompile(`\.`)
	domainDict := make(map[string] int)

	for _, domain := range cpdomains {
		countDomain := spaceRegex.Split(domain, -1)
		periodSplitted := periodRegex.Split(countDomain[1], -1)
		count, _ := strconv.Atoi(countDomain[0])
		var collected string

		for i := len(periodSplitted) -1; i >= 0; i-- {
			str := periodSplitted[i]

			if len(collected) == 0 {
				collected = str
			} else {
				collected = fmt.Sprintf("%s.%s", str, collected)
			}
			ct, in := domainDict[collected]

			if !in {
				ct = 0
			}

			domainDict[collected] = ct + count
		}
	}

	answer := make([]string, 0, len(domainDict))

	for key, count := range domainDict {
		answer = append(answer, fmt.Sprintf("%d %s", count, key))
	}
	return answer
}

func main() {
	fmt.Printf("visits %v", subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
