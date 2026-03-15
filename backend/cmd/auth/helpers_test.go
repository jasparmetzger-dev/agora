package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	type s struct {
		RawPassword string
		Name        string
	}
	data := []s{
		{RawPassword: "password123!", Name: "valid"},
		{RawPassword: "pwd", Name: "unsafe"},
	}
	for _, d := range data {
		hashed, err := HashPassword(d.RawPassword)
		assert.NoError(t, err, d.Name)
		assert.NotEmpty(t, hashed, d.Name)
	}
}
func TestCheckPasswordHash(t *testing.T) {
	type s struct {
		RawPassword string
		Name        string
	}
	data := []s{
		{RawPassword: "password123!", Name: "valid"},
		{RawPassword: "pwd", Name: "unsafe"},
	}
	for _, d := range data {
		hashed, _ := HashPassword(d.RawPassword)
		var got1 bool = CheckPasswordHash(d.RawPassword, hashed)
		var got2 bool = CheckPasswordHash(d.RawPassword+"a", hashed+"a")

		assert.True(t, got1, d.Name)
		assert.False(t, got2, d.Name)
	}
}

func TestJWTLogic(t *testing.T) {
	type s struct {
		userId      string
		SECRET_KEY  string
		tokenString string
		Name        string
	}
	data := []s{
		{userId: "1", SECRET_KEY: "testingisbnroiung", Name: "simple data", tokenString: "1"},
		{userId: "USBqwefdwjprg12p123eipdcv", SECRET_KEY: "SOJB23rowbndcaop23AIDH", tokenString: "OSJERowerbnonhowe234.wsdgi230o4ht.werogj", Name: "Invalid tokenString"},
	}

	for _, d := range data {
		token, err := GenerateJWT(d.userId, d.SECRET_KEY)
		assert.NoError(t, err, d.Name)
		assert.NotEmpty(t, token, d.Name)

		var fa bool = ValidateJWT(d.tokenString, d.SECRET_KEY, d.userId)
		var tr bool = ValidateJWT(token, d.SECRET_KEY, d.userId)
		assert.False(t, fa, d.Name)
		assert.True(t, tr, d.Name)
	}
}
