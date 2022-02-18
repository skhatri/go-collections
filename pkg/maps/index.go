package maps


func MapByStringKey(source interface{}) interface{} {
	switch source.(type) {
	case map[interface{}]interface{}:
		m := source.(map[interface{}]interface{})
		result := make(map[string]interface{}, len(m))
		for k, v := range m {
			result[k.(string)] = MapByStringKey(v)
		}
		return result

	case map[string]interface{}:
		target := make(map[string]interface{})

		for k, v := range source.(map[string]interface{}) {
			target[k] = MapByStringKey(v)
		}
		return target
	case []interface{}:
		items := make([]interface{}, 0)
		for _, interfaceItem := range source.([]interface{}) {
			items = append(items, MapByStringKey(interfaceItem))
		}
		return items
	default:
		return source
	}
}

