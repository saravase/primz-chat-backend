package injector

import (
	"fmt"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/saravase/primz-chat-backend/repository"
	"github.com/saravase/primz-chat-backend/service"
)

func (i *Injector) TokenInjector() error {
	tokenRepository := repository.NewTokenRepository(i.DB)
	i.TokenRepository = tokenRepository
	privKeyFile := i.Cfg.PrivKeyFile()
	priv, err := ioutil.ReadFile(privKeyFile)

	if err != nil {
		return fmt.Errorf("could not read private key pem file: %w", err)
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)

	if err != nil {
		return fmt.Errorf("could not parse private key: %w", err)
	}

	pubKeyFile := i.Cfg.PubKeyFile()
	pub, err := ioutil.ReadFile(pubKeyFile)

	if err != nil {
		return fmt.Errorf("could not read public key pem file: %w", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)

	if err != nil {
		return fmt.Errorf("could not parse public key: %w", err)
	}

	refreshSecret := i.Cfg.RefreshSecret()

	idTokenExp := i.Cfg.IdTokenExp()
	refreshTokenExp := i.Cfg.RefreshTokenExp()

	tokenService := service.NewTokenService(&service.TSConfig{
		TokenRepository:       tokenRepository,
		PrivKey:               privKey,
		PubKey:                pubKey,
		RefreshSecret:         refreshSecret,
		IDExpirationSecs:      idTokenExp,
		RefreshExpirationSecs: refreshTokenExp,
	})
	i.TokenService = tokenService

	return nil
}
