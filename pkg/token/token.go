package token

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"minepin-backend/config"
	"minepin-backend/model"
	"time"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")
)

// Context is the context of the JSON web token.
type Context struct {
	Role model.UserType // 用户角色
	UUID string         // 用户 UUID
	//nbf  int64          // JWT Token 生效时间
	//iat  int64          // JWT Token 签发时间
	//exp  int64          // JWT Token 过期时间
}

type Claims struct {
	URole model.UserType `json:"u_role"` // 用户角色
	UID   uint64         `json:"uid"`    // 用户 ID
	UName string         `json:"u_name"`
	jwt.StandardClaims
}

func Sign(c Claims) (accessTokenString string, refreshTokenString string, err error) {
	// The token content.
	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		URole: c.URole,
		UID:   c.UID,
		UName: c.UName,
		StandardClaims: jwt.StandardClaims{
			Audience:  c.UName,
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	})
	accessTokenString, err = aToken.SignedString([]byte(config.GetMinePinJwtAccessSecret()))
	if err != nil {
		return "", "", err
	}
	accessTokenString = base64.URLEncoding.EncodeToString([]byte(accessTokenString))

	rToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		URole: c.URole,
		UID:   c.UID,
		UName: c.UName,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	})
	refreshTokenString, err = rToken.SignedString([]byte(config.GetMinePinJwtRefreshSecret()))
	if err != nil {
		return "", "", err
	}
	refreshTokenString = base64.URLEncoding.EncodeToString([]byte(refreshTokenString))

	return
}

//func SignWithRefresh(c Context) (accessTokenString string, err error) {
//	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"uuid": c.UUID,
//		"role": c.Role,
//		"nbf":  time.Now().Unix(),                       // JWT Token 生效时间
//		"iat":  time.Now().Unix(),                       // JWT Token 签发时间
//		"exp":  time.Now().Add(time.Minute * 15).Unix(), // JWT Token 过期时间
//	})
//	accessTokenString, err = aToken.SignedString([]byte(config.GetMinePinJwtAccessSecret()))
//	if err != nil {
//		return "", err
//	}
//	accessTokenString = base64.URLEncoding.EncodeToString([]byte(accessTokenString))
//	return
//}

// secretFunc 验证密钥格式
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

// Parse 使用指定的 secret 验证 token ，有效则
// 返回上下文。
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(tokenString, secretFunc(secret))

	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.UUID = claims["uuid"].(string)
		ctx.Role = model.UserType(claims["role"].(float64))

		return ctx, nil
	} else {
		return ctx, err
	}
}

// ParseRequest 从 HTTP 请求头获取 token
// 并将其传递给 Parse 函数以验证 token 有消息。
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")
	secret := config.GetMinePinJwtAccessSecret()

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	fmt.Sscanf(header, "Bearer %s", &t)
	bt, _ := base64.URLEncoding.DecodeString(t)

	return Parse(string(bt), secret)
}

// ParseRefreshTokenRequest 从 HTTP 请求头获取 refresh token
// 并将其传递给 Parse 函数以验证 token 合法性。
func ParseRefreshTokenRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")
	secret := config.GetMinePinJwtRefreshSecret()

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	fmt.Sscanf(header, "Bearer %s", &t)
	bt, _ := base64.URLEncoding.DecodeString(t)

	return Parse(string(bt), secret)
}
