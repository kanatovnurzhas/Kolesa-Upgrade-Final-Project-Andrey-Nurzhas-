package models

type Recipient struct {
	User string
}

func (r Recipient) Recipient() string {
	return r.User
}
