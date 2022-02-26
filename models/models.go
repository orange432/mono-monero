package models

import "gorm.io/gorm"

type TemplateData struct {
	Strings     map[string]string
	User        UserDetails
	Token       string // CSRF Token
	Captcha     string
	CaptchaHash string
	Error       string
}

type User struct {
	gorm.Model
	ID             uint   `gorm:"primaryKey;autoIncrement"`
	Username       string `gorm:"unique"`
	Password       string
	PGPKey         string
	SecretPhrase   string
	Balance        float32
	Address        string
	WalletPassword string
}

type UserDetails struct {
	Username string
	Balance  float32
}
