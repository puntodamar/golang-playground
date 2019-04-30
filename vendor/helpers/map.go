package helpers

func CleanEmptyMap(m map[string]interface{}) map[string]interface{}{
	for k,v := range m {
		if _, ok := v.(string); ok {
			delete(m,k)
		}
	}

	return map[string]interface{}{}
}
