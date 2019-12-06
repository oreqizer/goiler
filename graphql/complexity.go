package graphql

import "github.com/oreqizer/goiler/generated"

const DefaultListComplexity = 20

func Connection(child int, after *string, first *int, before *string, last *int) int {
	if first != nil {
		return child * *first
	}
	if last != nil {
		return child * *last
	}
	return child * DefaultListComplexity
}

// Complexity adds array complexities to prevent resource hogging
// https://gqlgen.com/reference/complexity/
func Complexity(c *generated.Config) {
	c.Complexity.Query.Accounts = Connection
}
