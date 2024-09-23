package api

import (
	"encoding/json"
	"fmt"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/container"
	"strings"
)

type DomainCondition struct {
	Key         string `json:"key"`
	Conditional string `json:"conditional"`
	Value       string `json:"value"`
}

func ParseDomainConditions(domainStr string) ([]DomainCondition, error) {
	var conditions []DomainCondition
	replacer := strings.NewReplacer(
		`(`, "[",
		`)`, "]",
		`<`, "[",
		`>`, "]",
	)

	domainStr = replacer.Replace(domainStr)
	//domainStr = domainStr[1 : len(domainStr)-1]
	//
	//fmt.Println(ParseData(domainStr))
	err := json.Unmarshal([]byte(domainStr), &conditions)
	if err != nil {
		return nil, err
	}
	return conditions, nil
}

type BaseSearch struct {
	filters []map[string]interface{}
	domain  []DomainCondition
}

func NewBaseSearch(domain string) *BaseSearch {
	d, err := ParseDomainConditions(domain)
	if err != nil {
		container.Logger.Error("err =>", err)
		d = []DomainCondition{}
	}
	return &BaseSearch{domain: d}
}

func (b *BaseSearch) Filter() {
	fmt.Println("filter => ", b.domain)
	for k, v := range b.domain {
		fmt.Println(k, " => ", v)
	}
}

// ParseData parses the given string into a slice of strings
func ParseData(data string) ([]string, error) {
	// Remove outer brackets and parentheses
	data = strings.TrimPrefix(strings.TrimSuffix(data, ")"), "(")

	// Split by comma
	parts := strings.Split(data, ",")

	// Trim each part and collect in a new slice
	var result []string
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if len(part) > 0 && part[0] == '\'' && part[len(part)-1] == '\'' {
			result = append(result, part[1:len(part)-1])
		} else {
			return nil, fmt.Errorf("invalid format: %s", part)
		}
	}

	return result, nil
}
