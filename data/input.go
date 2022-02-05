package data

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Data struct {
	Content interface{}
}

func (d *Data) Templates() []string {
	if m, ok := d.Content.(map[string]interface{}); ok {
		templates := m["__templates__"]
		if t, ok := templates.([]interface{}); ok {
			var sList []string
			for _, s := range t {
				if str, ok := s.(string); ok {
					sList = append(sList, str)
				}
			}
			return sList
		}
		switch t := templates.(type) {
		case []string:
			return t
		default:
			return make([]string, 0)
		}
	}
	return make([]string, 0)
}

func New(file string) (Data, error) {
	var data Data
	if bytes, err := ioutil.ReadFile(file); err != nil {
		return data, err
	} else {
		if err := yaml.Unmarshal(bytes, &data.Content); err != nil {
			return data, err
		} else {
			return data, nil
		}
	}
}
