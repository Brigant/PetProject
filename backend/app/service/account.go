package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/golang-jwt/jwt"
)

type AccountService struct {
	storage AccountStorage
}

func NewAccountService(storage AccountStorage) AccountService {
	return AccountService{storage: storage}
}

const (
	signingKey      = "sdFWlnxb13t&r"
	accessTokenTTL  = 15 * time.Minute
	refreshTokenTTL = 24 * time.Hour
)

var (
	errInvalidSigningMethod = errors.New("invalid signing metod")
	errWrongTokenClaimType  = errors.New("token claims are not of type *tokenClaims")
)

type UserInfo struct {
	UserID   string
	UserRole string
}

type Claims struct {
	jwt.StandardClaims
	Info UserInfo
}

// The function receives the account model and store it in the repository, after that returns account id
// of the new created account or an error if it occures.
func (a AccountService) CreateUser(account core.Account) (string, error) {
	id, err := a.storage.InsertAccount(account)
	if err != nil {
		return "", fmt.Errorf("service CreateUser get an error: %w", err)
	}

	return id, nil
}

func (a AccountService) GetUser() error {
	return nil
}

// The function GenerateToken represents bissness logic layer
// and  generate token.
// func (a AccountService) GenerateTokens(phone, password string) (string, string, error) {
// 	user, err := a.storage.SelectAccount(phone, a.sha256(password))
// 	if err != nil {
// 		return "", "", fmt.Errorf("error occures while GetUser: %w", err)
// 	}

// 	accessToken, err := a.generateAccessToken(user.ID, user.Role)
// 	if err != nil {
// 		return "", "", fmt.Errorf("cerror occures while generateAccessToken: %w", err)
// 	}

// 	refreshToken, err := a.generateRefreshToken(user.ID)
// 	if err != nil {
// 		return "", "", fmt.Errorf("error occures while generateRefreshToken: %w", err)
// 	}

// 	return accessToken, refreshToken, nil
// }

// The function returns user ID if accessToken is valid.
func (a AccountService) ParseToken(accesToken string) (string, string, error) {
	token, err := jwt.ParseWithClaims(accesToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errInvalidSigningMethod
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", "", fmt.Errorf("accessToken throws an error during parsing: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", "", errWrongTokenClaimType
	}

	return claims.Info.UserID, claims.Info.UserRole, nil
}

func (a AccountService) generateAccessToken(id, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Info: UserInfo{
			UserID:   id,
			UserRole: role,
		},
	})

	accessToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", fmt.Errorf("cannot get SignetString token: %w", err)
	}

	newClaims := &Claims{}

	jwtToken, err := jwt.ParseWithClaims(accessToken, newClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", fmt.Errorf("receives %v, error occurs while ParseWithClaims: %w", jwtToken, err)
	}

	return accessToken, nil
}

func (a AccountService) generateRefreshToken(accountID string) (string, error) {
	_ = accountID
	// token, err := a.storage .InsertRefreshToken(accountID, refreshTokenTTL)
	// if err != nil {
	// 	return "", fmt.Errorf("cannot generate refresh tokes cause of: %w", err)
	// }

	return "token", nil
}
