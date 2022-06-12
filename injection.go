package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/config"
	"github.com/saravase/primz-chat-backend/handler/auth"
	"github.com/saravase/primz-chat-backend/repository"
	"github.com/saravase/primz-chat-backend/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func inject(cfg config.Config, db *mongo.Database) (*gin.Engine, error) {
	log.Println("Injecting data sources")

	userRepository := repository.NewUserRepository(db)
	tokenRepository := repository.NewTokenRepository(db)

	userService := service.NewUserService(&service.USConfig{
		UserRepository: userRepository,
	})

	privKeyFile := cfg.PrivKeyFile()
	priv, err := ioutil.ReadFile(privKeyFile)

	if err != nil {
		return nil, fmt.Errorf("could not read private key pem file: %w", err)
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)

	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %w", err)
	}

	pubKeyFile := cfg.PubKeyFile()
	pub, err := ioutil.ReadFile(pubKeyFile)

	if err != nil {
		return nil, fmt.Errorf("could not read public key pem file: %w", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)

	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}

	refreshSecret := cfg.RefreshSecret()

	idTokenExp := cfg.IdTokenExp()
	refreshTokenExp := cfg.RefreshTokenExp()

	tokenService := service.NewTokenService(&service.TSConfig{
		TokenRepository:       tokenRepository,
		PrivKey:               privKey,
		PubKey:                pubKey,
		RefreshSecret:         refreshSecret,
		IDExpirationSecs:      idTokenExp,
		RefreshExpirationSecs: refreshTokenExp,
	})

	router := gin.Default()

	baseURL := cfg.AuthBaseURL()

	// read in HANDLER_TIMEOUT
	handlerTimeout := cfg.HandlerTimeout()

	auth.NewHandler(&auth.Config{
		R:               router,
		UserService:     userService,
		TokenService:    tokenService,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(handlerTimeout) * time.Second),
	})

	return router, nil
}
