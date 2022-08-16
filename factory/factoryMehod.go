// Package factory DESC
package factory

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type jsonRuleConfigParserFactory struct {
}

func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

type yamlRuleConfigParseFactory struct{}

func (y yamlRuleConfigParseFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParseFactory{}
	}
	return nil
}
