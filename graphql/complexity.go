package graphql

import "github.com/oreqizer/goiler/generated"

func Array(child, count int) int {
	return count * child
}

// Complexity adds array complexities to prevent resource hogging
// https://gqlgen.com/reference/complexity/
func Complexity(c *generated.Config) {
	// Example:
	// c.Complexity.Query.Posts = Array
}
