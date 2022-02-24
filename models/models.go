package models

type TemplateData struct {
	Strings map[string]string
	User    UserDetails
}

type UserDetails struct {
	Username string
	Balance  float32
}
