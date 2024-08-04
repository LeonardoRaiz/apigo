package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type SearchResult struct {
	Items []struct {
		Link string `json:"link"`
	} `json:"items"`
}

func googleSearch(searchTerm string) (*SearchResult, error) {
	apiKey := os.Getenv("API_KEY")
	cseID := os.Getenv("CSE_ID")
	endpoint := "https://www.googleapis.com/customsearch/v1"
	params := url.Values{}
	params.Add("q", searchTerm)
	params.Add("key", apiKey)
	params.Add("cx", cseID)

	resp, err := http.Get(fmt.Sprintf("%s?%s", endpoint, params.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func extractDomains(searchResult *SearchResult) []string {
	domainSet := make(map[string]struct{})
	for _, item := range searchResult.Items {
		link := item.Link
		if link != "" {
			domain := strings.Split(link, "/")[2]
			domainSet[domain] = struct{}{}
		}
	}

	domains := make([]string, 0, len(domainSet))
	for domain := range domainSet {
		domains = append(domains, domain)
	}
	return domains
}
