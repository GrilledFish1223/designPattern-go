// Package factory DESC
package factory

import "fmt"

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct {
}

func (j jsonRuleConfigParser) Parse(data []byte) {
	fmt.Println("json parse")
}

type yamlRuleConfigParser struct {
}

func (y yamlRuleConfigParser) Parse(data []byte) {
	fmt.Println("yaml parse")
}

func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}
