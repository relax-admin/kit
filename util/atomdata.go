package util

import (
	"encoding/json"
	"encoding/xml"
)

type AtomData struct {
	Data map[string]interface{}
}

func NewAtomData() *AtomData {
	AtomData := AtomData{}
	AtomData.Data = make(map[string]interface{})
	return &AtomData
}

func (a *AtomData) ToXml() string {
	if len(a.Data) == 0 {
		return ""
	}
	if x, err := xml.MarshalIndent(a.Data, "", " "); err == nil {
		return string(x)
	}
	return ""
}

func (a *AtomData) ToUrl() string {
	var m = make(map[string]string, 0)
	for k, v := range a.Data {
		m[k] = ToString(v)
	}
	return Encode(&m)
}

func (a *AtomData) ToJson() string {
	v, _ := json.Marshal(a.Data)
	return string(v)
}

func (a *AtomData) FromXml(xmlStr string, key string) (err error) {
	var data map[string]interface{}
	data, err = XmlToMap(xmlStr)
	if err != nil {
		return
	}
	v, has := data["xml"]
	if has {
		a.Data = v.(map[string]interface{})
	}
	return
}

// func (a *AtomData) FromXml(xmlStr string, key string) (err error) {
// 	a.Data, err = XmlToMap(xmlStr)
// 	return
// }

func (a *AtomData) SetValue(key string, value interface{}) {
	a.Data[key] = value
}

func (a *AtomData) GetValue(key string) interface{} {
	return a.Data[key]
}
func (a *AtomData) RemoveKey(key string) {
	delete(a.Data, key)
}

func (a *AtomData) IsSet(key string) bool {
	_, ok := a.Data[key]
	return ok
}
