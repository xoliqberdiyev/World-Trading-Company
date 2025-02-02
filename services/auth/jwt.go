package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	types_admin "github.com/XoliqberdiyevBehruz/wtc_backend/types/admin"
	"github.com/XoliqberdiyevBehruz/wtc_backend/utils"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "userId"

func CreateJWT(secrets []byte, userId string) (string, error) {
	expiration := time.Second * time.Duration(3600*24*7)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userId,
		"expiredAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secrets)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthWithJWT(HandleFunc http.HandlerFunc, store types_admin.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from user request
		tokenString := getTokenFronRequest(r)

		// validate token
		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenidet(w)
			return
		}
		if !token.Valid {
			log.Println("invlid token")
			permissionDenidet(w)
			return
		}

		// if we need to fetch userId from database
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["userID"].(string)

		// userId, err := strconv.Atoi(str)
		// if err != nil {
		// log.Println("failed to convert user Id to int: %v", err)
		// permissionDenidet(w)
		// return
		// }

		u, err := store.GetUserById(userId)
		if err != nil {
			log.Printf("failed to user get: %v", err)
			permissionDenidet(w)
			return
		}

		// set context "userId" to the user Id
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.Id)
		r = r.WithContext(ctx)

		HandleFunc(w, r)
	}
}

func getTokenFronRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	if tokenAuth != "" {
		return tokenAuth
	}
	return ""
}

func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexepted signing method: %v", t.Header["alg"])
		}
		return []byte("jwt_token_secrets"), nil
	})
}

func permissionDenidet(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("permission denidet"))
}

func GetUserIdFronContext(ctx context.Context) string {
	userId, ok := ctx.Value(UserKey).(string)
	if !ok {
		return ""
	}
	return userId
}
