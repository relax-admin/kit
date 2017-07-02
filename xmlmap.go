package kit

import (
	"bytes"
	"encoding/xml"
	"net/url"
	"sort"
)

// StringMap is a map[string]string.
type XmlMap map[string]string

// StringMap marshals into XML.
func (x XmlMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "xml"
	tokens := []xml.Token{start}
	for key, value := range x {
		t := xml.StartElement{Name: xml.Name{"", key}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (x *XmlMap) Encode(isUrl ...bool) string {
	var escape bool
	if len(isUrl) == 0 {
		escape = true
	} else {
		escape = isUrl[0]
	}
	if x == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(*x))
	for k := range *x {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := (*x)[k]
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
