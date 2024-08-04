package services

import (
	"fmt"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
	}
}

func PerformSearch(terms []string, email string) error {
	allDomains := make(map[string]struct{})
	for _, term := range terms {
		results, err := googleSearch(term)
		if err != nil {
			return fmt.Errorf("erro ao buscar o termo %s: %v", term, err)
		}
		domains := extractDomains(results)
		for _, domain := range domains {
			allDomains[domain] = struct{}{}
		}
	}

	domainList := make([]string, 0, len(allDomains))
	for domain := range allDomains {
		domainList = append(domainList, domain)
	}

	err := saveResults(terms, email, domainList)
	if err != nil {
		return err
	}

	err = sendConfirmationEmail(email, terms, domainList)
	if err != nil {
		return err
	}

	return nil
}
