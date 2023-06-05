package owasp

import "github.com/daveshanley/vacuum/model"

// Copied from openapi test

func buildOwaspTestRuleAction(given, function, field string, functionOptions interface{}) model.Rule {
	return model.Rule{
		Given: given,
		Then: &model.RuleAction{
			Field:           field,
			Function:        function,
			FunctionOptions: functionOptions,
		},
	}
}

func buildOwaspTestContext(action *model.RuleAction, options map[string]string) model.RuleFunctionContext {
	return model.RuleFunctionContext{
		RuleAction: action,
		Options:    options,
	}
}
