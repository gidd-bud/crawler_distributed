package model

import "encoding/json"

type Profile struct {
	MemberID	string
	Nickname	string
	//TrueName	string
	Age			int
	Gender		string
	Height		int
	Education	string
	WorkCity	string
	Marriage	string
	Salary		string
}

func (p *Profile)FromJsonObj(o interface{}) error {
	bytes, err := json.Marshal(o)
	if err == nil {
		err = json.Unmarshal(bytes, p)
	}
	return err
}