package util

import (
	"encoding/json"
	"encoding/xml"
)

type AtomString struct {
	Data XmlMap
}

func NewAtomString() (atomString *AtomString) {
	atomString = new(AtomString)
	atomString.Data = make(XmlMap)
	return
}

func (a *AtomString) ToXml() string {
	x, _ := xml.MarshalIndent(a.Data, "", " ")
	return string(x)
}

func (a *AtomString) ToUrl() string {
	if a.Data == nil {
		return ""
	}
	return a.Data.Encode()
}
func (a *AtomString) ToJson() string {
	v, _ := json.Marshal(a.Data)
	return string(v)
}

func (a *AtomString) FromXml(xml string, key string) (err error) {
	a.Data, err = XmlToXmlMap(xml)
	return
}

// func (a *AtomString) FromXml(xmlStr string, key string) (err error) {
// 	var data XmlMap
// 	data, err = XmlToXmlMap(xmlStr)
// 	if err != nil {
// 		return
// 	}
// 	inData, has := data["xml"]
// 	if has {
// 		for k, v := range inData {
// 			a.Data[k] = v
// 		}
// 	}
// 	// if has {
// 	// 	a.Data = v.(XmlMap)
// 	// }
// 	return
// }

func (a *AtomString) SetValue(key string, value string) {
	a.Data[key] = value
}

func (a *AtomString) GetValue(key string) string {
	return a.Data[key]
}
func (a *AtomString) RemoveKey(key string) {
	delete(a.Data, key)
}

func (a *AtomString) IsSet(key string) bool {
	_, ok := a.Data[key]
	return ok
}
