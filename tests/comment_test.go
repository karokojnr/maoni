//go:build e2e

package tests

import (
	"fmt"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("c2VjcmV0ignvbnNlY3JldCJ9"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func TestPostComment(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			setHeader("Authorization", "bearer "+createToken()).
			SetBody(map[string]interface{}{
				"slug":   "slug",
				"author": "author",
				"body":   "body",
			}).
			Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})

	t.Run("cannot post comment withut JWT", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().SetBody(`{"slug": "/", "author": "Test Author", "body": "hello world"}`).Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())
	})

}
