package main

import (
	"time"
	"fmt"
)

type AuthUser struct {
	ID UserID
	token string

	Email string `json:"email"`
	EmailVerified bool `json:"email_verified"`
	Name string `json:"name"`
	GivenName string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Locale string `json:"locale"`
	Picture string `json:"picture"`

	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Azp string `json:"azp"`
	Aud string `json:"aud"`
	Iat JSONTime `json:"iat"`
	Exp JSONTime `json:"exp"`
}

type User struct {
	ID UserID
	Email string `json:"email"`
	Name string `json:"name"`
	GivenName string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture string `json:"picture"`
	Admin bool `json:"admin"`

	Parts []PartID `json:"parts"`
	Comments []CommentID `json:"comments"`
	Media []MediaID `json:"comments"`
}

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Second())
	return []byte(stamp), nil
}

type UserID string

type Subsystem struct {
	ID SubsystemID `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Media []MediaID `json:"media"`
	Comments []CommentID `json:"comments"`
	Parts []PartID `json:"parts"`
}

type SubsystemID string

type Part struct {
	ID          PartID
	Author      UserID
	Name        string
	Description string
	Comment     []CommentID
	Subsystem   []SubsystemID
}

type PartID string

type Comment struct {
	ID CommentID `json:"id"`
	Time JSONTime `json:"created"`
	Author UserID `json:"author"`
	Content string `json:"content"`
	Media []MediaID `json:"media"`
}

type CommentID string

type Media struct {
	ID int `json:"id"`
	Author UserID `json:"author"`
	Base64 string `json:"base64"`
	hash []byte
}

type MediaID int