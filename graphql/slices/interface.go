package slices

// StringsToInterfaces creates a slice of interfaces from strings
func StringsToInterfaces(strs []string) []interface{} {
	its := make([]interface{}, 0, len(strs))
	for _, key := range strs {
		its = append(its, key)
	}
	return its
}
