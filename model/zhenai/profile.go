package zhenai

import "encoding/json"

type Profile struct {
	Name          string
	Gender        string
	Age           string
	Height        string
	Weight        string
	Income        string
	Marriage      string
	Education     string
	Occupation    string
	Hukou         string
	constellation string
	House         string
	Car           string
	Avatar        string
}

func Json2Obj(o interface{}) (Profile, error) {
	var r Profile
	bytes, err := json.Marshal(o)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(bytes, &r)
	return r, err
}

type UserInfo struct {
	Url    string
	Id     string
	Name   string
	Gender string
	Avatar string
}
