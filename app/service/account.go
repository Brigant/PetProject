package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Brigant/PetPorject/app/core"
	"github.com/Brigant/PetPorject/config"
	"github.com/golang-jwt/jwt"
)

type AccountService struct {
	storage AccountStorage
	cfg     config.Config
}

func NewAccountService(storage AccountStorage, cfg config.Config) AccountService {
	return AccountService{storage: storage, cfg: cfg}
}

var (
	errInvalidSigningMethod = errors.New("invalid signing metod")
	errWrongTokenClaimType  = errors.New("token claims are not of type *tokenClaims")
)

type ClinteSideInfo struct {
	AccountID    string
	Role         string
	RefreshToken string
	RequestHost  string
	UserAgent    string
	ClientIP     string
}

type Claims struct {
	jwt.StandardClaims
	Info ClinteSideInfo
}

// The function receives the account model and store it in the repository, after that returns account id
// of the new created account or an error if it occures.
func (a AccountService) CreateUser(account core.Account) (string, error) {
	account.Password = core.SHA256(account.Password, a.cfg.Salt)

	id, err := a.storage.InsertAccount(account)
	if err != nil {
		return "", fmt.Errorf("service CreateUser get an error: %w", err)
	}

	return id, nil
}

// The service implementation of login functionality.
func (a AccountService) Login(phone, password string, session core.Session) (core.TokenPair, error) {
	var tokenPair core.TokenPair

	account, err := a.storage.SelectAccountByPhone(phone)
	if err != nil {
		return core.TokenPair{}, fmt.Errorf("service Login got the error: %w", err)
	}

	if core.SHA256(password, a.cfg.Salt) != account.Password {
		return core.TokenPair{}, core.ErrWrongPassword
	}

	expired := time.Now().Add(a.cfg.RefreshTokenTTL)

	session.AccountID = account.ID
	session.Role = account.Role
	session.Expired = expired

	session, err = a.storage.InsertSession(session)
	if err != nil {
		return core.TokenPair{}, fmt.Errorf("error occures in service Loign: %w", err)
	}

	accesstoken, err := a.generateAccessToken(session)
	if err != nil {
		return core.TokenPair{}, fmt.Errorf("error occures in service Loign: %w", err)
	}

	tokenPair.AccessToken = accesstoken
	tokenPair.RefreshToken = session.RefreshToken

	return tokenPair, nil
}

// The function returns user ID if accessToken is valid.
func (a AccountService) ParseToken(accesToken string) (string, string, error) {
	token, err := jwt.ParseWithClaims(accesToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errInvalidSigningMethod
		}

		return []byte(a.cfg.SigningKey), nil
	})
	if err != nil {
		return "", "", fmt.Errorf("accessToken throws an error during parsing: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", "", errWrongTokenClaimType
	}

	return claims.Info.AccountID, claims.Info.Role, nil
}

func (a AccountService) RefreshTokenpair(session core.Session) (core.TokenPair, error) {
	sessionFromDB, err := a.storage.SelectSession(session)
	if err != nil {
		return core.TokenPair{}, fmt.Errorf("can't Select Session: %w", err)
	}

	if sessionFromDB.Expired.Unix() < time.Now().Unix() {
		return core.TokenPair{}, core.ErrRefreshTokenExpired
	}

	sessionFromDB.Expired = time.Now().Add(a.cfg.RefreshTokenTTL)

	err = a.storage.RefreshSession(sessionFromDB)
	if err != nil {
		return core.TokenPair{}, fmt.Errorf("storege can't refress this session: %w", err)
	}

	accessToken, err := a.generateAccessToken(sessionFromDB)
	if err != nil {
		return core.TokenPair{}, fmt.Errorf("error happened while generating Access Token: %w", err)
	}

	tokenPair := core.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: session.RefreshToken,
	}

	return tokenPair, nil
}

func (a AccountService) generateAccessToken(session core.Session) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(a.cfg.RefreshTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Info: ClinteSideInfo{
			AccountID:    session.AccountID,
			Role:         session.Role,
			RefreshToken: session.RefreshToken,
			RequestHost:  session.RequestHost,
			UserAgent:    session.UserAgent,
			ClientIP:     session.ClientIP,
		},
	})

	accessToken, err := token.SignedString([]byte(a.cfg.SigningKey))
	if err != nil {
		return "", fmt.Errorf("cannot get SignetString token: %w", err)
	}

	return accessToken, nil
}

func (a AccountService) Logout(acountID string) error {
	if err := a.storage.DeleteSesions(acountID); err != nil {
		return fmt.Errorf("can't delete the account sessions: %w", err)
	}

	return nil
}
