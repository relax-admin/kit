package kit

//务必保证值为string
func ToMapString(param map[string]interface{}) (result map[string]string) {
	result = make(map[string]string, 0)
	for k, v := range param {
		result[k] = v.(string)
	}
	return
}

func ToMapInterface(param map[string]string) (result map[string]interface{}) {
	result = make(map[string]interface{}, 0)
	for k, v := range param {
		result[k] = v
	}
	return
}
