package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jasparmetzger-dev/agora/conf"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	type tc struct {
		Name       string
		Token      string
		ExpectCode int
	}

	validToken, _ := GenerateJWT("user123", conf.SECRET_KEY)

	cases := []tc{
		{Name: "valid token", Token: validToken, ExpectCode: 200},
		{Name: "invalid token", Token: "badtoken", ExpectCode: 401},
		{Name: "missing token", Token: "", ExpectCode: 401},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			ctx.Request = httptest.NewRequest(http.MethodGet, "/", nil)
			if c.Token != "" {
				ctx.Request.Header.Set("Authorization", "Bearer "+c.Token)
			}

			AuthMiddleware()(ctx)

			assert.Equal(t, c.ExpectCode, w.Code)
		})
	}
}
