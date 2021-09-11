package lib

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/uzzeet/uzzeet-gateway/libs/helper"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"github.com/uzzeet/uzzeet-gateway/models"
	"os"
	"strings"
)

func ClaimToken(tokens []string) (response models.AuthorizationInfo, serr serror.SError) {
	secretKey := []byte("um_phrase")

	if tokens == nil {
		return response, serror.New("Token tidak ditemukan")
	}

	tokenString, err := parseToken(tokens[0])
	if err != nil {
		return response, serror.NewFromError(err)
	}

	if tokenString == os.Getenv("DEV_TOKEN") {
		response = models.AuthorizationInfo{
			UserID:         "1",
			Username:       "nsq",
			IsOrgAdmin:     1,
			IsActive:       0,
			OrganizationId: "nsq",
			AppId:          "nsq",
		}
	} else {
		decode, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, serror.NewFromError(fmt.Errorf("Unexpected signing method: %v", token.Header["alg"]))
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			return response, serror.NewFromError(err)
		}

		resToken := decode.Claims.(jwt.MapClaims)
		response = models.AuthorizationInfo{
			UserID:         fmt.Sprintf("%v", resToken["id"]),
			Username:       fmt.Sprintf("%v", resToken["username"]),
			IsOrgAdmin:     int(helper.StringToInt(helper.IntToString(int(resToken["isorgadmin"].(float64))), 0)),
			IsActive:       int(helper.StringToInt(helper.IntToString(int(resToken["isactive"].(float64))), 0)),
			OrganizationId: fmt.Sprintf("%v", resToken["organizationid"]),
			AppId:          fmt.Sprintf("%v", resToken["app"]),
			Exp:            int(helper.StringToInt(helper.IntToString(int(resToken["exp"].(float64))), 0)),
		}
	}

	return response, nil
}

func parseToken(source string) (token string, err error) {

	separator := " "
	valueSection := 1
	expectedTokenLength := 2

	if source == "" {
		return token, errors.New("Token tidak ditemukan")
	}

	tokens := strings.Split(source, separator)
	if len(tokens) != expectedTokenLength {
		return token, errors.New("Token tidak valid")
	}

	token = tokens[valueSection]
	return token, nil
}

func IsAdmin(ctx *gin.Context) (err error) {
	oaut, errx := ClaimToken(ctx.Request.Header["Authorization"])
	if errx != nil {
		return errx
	}

	level := oaut.IsOrgAdmin
	if level != 1 && level != 2 && level != 3 {
		return errors.New("Token access denied")
	}
	return errx
}
