package middleware

import (
	"Initial/dao"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWT struct {
	SigningKey []byte
}

type JwtConfig struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // 缓冲时间
}

type CustomClaims struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

type LoginResponse struct {
	User      dao.User `json:"user"`
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expiresAt"`
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			FailResponse(c, 400, "not login")
			c.Abort()
			return
		}
		j := NewJwt()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				FailResponse(c, 400, " token expired")
				c.Abort()
				return
			}
			FailResponse(c, 400, err.Error())
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

func NewJwt() *JWT {
	return &JWT{
		[]byte(JwtConfig{}.SigningKey),
	}
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}
