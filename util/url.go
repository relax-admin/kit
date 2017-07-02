package util

import (
	"bytes"
	"net/url"
	"sort"
)

// func SortedKeys(m *map[string]string) []string {
// 	sortedKeys := make([]string, len(*m))
// 	i := 0
// 	for key := range *m {
// 		sortedKeys[i] = key
// 		i++
// 	}
// 	sort.Strings(sortedKeys)
// 	return sortedKeys
// }

// func SortedUrl(m *map[string]string) string {
// 	if len(*m) == 0 {
// 		return ""
// 	}
// 	sk := SortedKeys(m)
// 	var sortedData string
// 	for _, k := range sk {
// 		sortedData += "&" + k + "=" + (*m)[k]
// 	}
// 	return sortedData[1:]
// }

func JointedUrl(m *map[string]string) string {
	if len(*m) == 0 {
		return ""
	}
	var data string
	for k, v := range *m {
		data += "&" + k + "=" + v
	}
	return data[1:]
}

/*
escape:
false is for sign
true is for url
*/
func Encode(v *map[string]string, isSign ...bool) string {
	var escape bool
	if len(isSign) == 0 {
		escape = false
	} else {
		escape = isSign[0]
	}
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(*v))
	for k := range *v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := (*v)[k]
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		if escape {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(vs))
		} else {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(vs)
		}
	}
	return buf.String()
}

func EncodeObject(v *map[string]interface{}, isSign ...bool) string {
	var escape bool
	if len(isSign) == 0 {
		escape = false
	} else {
		escape = isSign[0]
	}
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(*v))
	for k := range *v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := (*v)[k]
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		if escape {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			if tValue, ok := vs.(string); ok {
				buf.WriteString(url.QueryEscape(tValue))
			} else if tValue, ok := vs.([]byte); ok {
				buf.Write(tValue)
			}
			//buf.Write(url.QueryEscape(vs))
		} else {
			buf.WriteString(k)
			buf.WriteByte('=')
			if tValue, ok := vs.(string); ok {
				buf.WriteString(tValue)
			} else if tValue, ok := vs.([]byte); ok {
				buf.Write(tValue)
			}
			//buf.WriteString(vs)
		}
	}
	return buf.String()
}
