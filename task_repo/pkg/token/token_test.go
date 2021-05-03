package token

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/greenteam1/task_repo/pkg/models"
)

func TestCreateToken(t *testing.T) {
	user := models.User{
		ID:       102,
		Username: "John",
		Password: "Doe",
	}

	actual, err := Create(&user, "verysecretsaltfortesting")
	assert.NoError(t, err, "test failed")
	t.Run("Create", func(t *testing.T) {
		assert.NotEmpty(t, actual, "token string should not be empty")
		assert.NoError(t, err, "should be no errors")
	})
}

func TestParseToken(t *testing.T) {
	user := models.User{
		ID:       14,
		Username: "John",
		Password: "Doedoe",
	}

	secretKey := "verysecret"

	testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTQsImV4cCI6MjUyODAzNzY3M30.DBDmbwi20-l2ebxM5aj6PS09KQsVun3eqH-UFQZKUSE"
	actual, err := Parse(testToken, secretKey)
	assert.NoError(t, err, "test failed")
	t.Run("Parse", func(t *testing.T) {
		assert.Equal(t, user.ID, actual, "ID should be equal")
		assert.NoError(t, err, "should be no errors")
	})
}
