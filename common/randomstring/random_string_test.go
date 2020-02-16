package randomstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	rds := NewRandomString()

	t.Run("GenerateToken", func(t *testing.T) {
		token, err := rds.GenerateToken()
		assert.Nil(t, err)
		assert.IsType(t, "string", token)
		assert.GreaterOrEqual(t, len(token), 32)
	})

	t.Run("NumberOfBytes - token", func(t *testing.T) {
		token, _ := rds.GenerateToken()

		num, err := rds.NumberOfBytes(token)
		assert.Nil(t, err)
		assert.Equal(t, 32, num)
	})

	t.Run("NumberOfBytes - invalid token", func(t *testing.T) {
		token := "token"

		num, err := rds.NumberOfBytes(token)
		assert.NotNil(t, err)
		assert.Less(t, num, 32)
	})
}
