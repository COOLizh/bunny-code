// Package models describes entities
package models

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

// User - struct for testing
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserResponse - struct for testing
type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// Prepare ...
func (u *User) Prepare() error {
	if !u.IsValid() {
		return fmt.Errorf("couldn't validate user")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(u.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}

// IsValid checks if User's fields matches preassigned template
func (u *User) IsValid() bool {
	reString := "^[A-Za-z0-9]+(?:[ _-][A-Za-z0-9]+)*$"
	re := regexp.MustCompile(reString)

	return re.MatchString(u.Username) &&
		re.MatchString(u.Password) &&
		!(utf8.RuneCountInString(u.Password) < 6)
}

// PopulateFromRequest add fields from bytes to struct
func (u *User) PopulateFromRequest(requestBody io.Reader) (err error) {
	decoder := json.NewDecoder(requestBody)
	err = decoder.Decode(u)
	return
}

// PrepareResponse hidden password in response
func (u *User) PrepareResponse() (response UserResponse) {
	response.ID = u.ID
	response.Username = u.Username
	return
}
