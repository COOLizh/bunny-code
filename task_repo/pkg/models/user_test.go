package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCasesIsValid = []struct {
	in  User
	out bool
}{
	{User{Username: "test", Password: "test"}, false},
	{User{Username: "test", Password: "testsr"}, true},
	{User{Username: "", Password: ""}, false},
	{User{Username: "test", Password: ""}, false},
	{User{Username: "", Password: "testst"}, false},
	{User{Username: "abcdefghijklmnopqrstuvwxyz", Password: "abcdefghijklmnopqrstuvwxyz"}, true},
	{User{Username: "a  a", Password: "       "}, false},
}

func TestUser_IsValid(t *testing.T) {
	for _, v := range testCasesIsValid {
		actual := v.in.IsValid()
		if v.out {
			assert.Truef(t, actual, "Failed: %v", v.in)
		} else {
			assert.Falsef(t, actual, "Failed: %v", v.in)
		}
	}
}
