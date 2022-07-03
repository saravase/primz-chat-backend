package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config interface {
	AppPort() int
	DbPort() int
	Dsn() string
	DbName() string
	AuthBaseURL() string
	OrgBaseURL() string
	ChannelBaseURL() string
	MessageBaseURL() string
	ChatBaseURL() string
	HandlerTimeout() int64
	IdTokenExp() int64
	RefreshTokenExp() int64
	PrivKeyFile() string
	PubKeyFile() string
	RefreshSecret() string
}

type config struct {
	appPort         int
	dbUser          string
	dbPass          string
	dbHost          string
	dbPort          int
	dbName          string
	dsn             string
	authBaseURL     string
	orgBaseURL      string
	channelBaseURL  string
	messageBaseURL  string
	chatBaseURL     string
	handlerTimeout  int
	idTokenExp      int
	refreshSecret   string
	privKeyFile     string
	pubKeyFile      string
	refreshTokenExp int
}

func NewConfig() Config {
	var cfg config
	var err error
	cfg.appPort, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatalln("Error on load env var:", err.Error())
	}
	cfg.dbUser = os.Getenv("DATABASE_USER")
	cfg.dbPass = os.Getenv("DATABASE_PASS")
	cfg.dbHost = os.Getenv("DATABASE_HOST")
	cfg.dbName = os.Getenv("DATABASE_NAME")
	cfg.dbPort, err = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatalln("Error on load env var:", err.Error())
	}
	cfg.dsn = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cfg.dbUser, cfg.dbPass, cfg.dbHost, cfg.dbPort, cfg.dbName)
	cfg.authBaseURL = os.Getenv("AUTH_BASE_URL")
	cfg.orgBaseURL = os.Getenv("ORG_BASE_URL")
	cfg.channelBaseURL = os.Getenv("CHANNEL_BASE_URL")
	cfg.messageBaseURL = os.Getenv("MESSAGE_BASE_URL")
	cfg.chatBaseURL = os.Getenv("CHAT_BASE_URL")
	cfg.handlerTimeout, err = strconv.Atoi(os.Getenv("HANDLER_TIMEOUT"))
	if err != nil {
		log.Fatalln("Error on load env var:", err.Error())
	}
	cfg.idTokenExp, err = strconv.Atoi(os.Getenv("ID_TOKEN_EXP"))
	if err != nil {
		log.Fatalln("Error on load env var:", err.Error())
	}
	cfg.refreshTokenExp, err = strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXP"))
	if err != nil {
		log.Fatalln("Error on load env var:", err.Error())
	}
	cfg.refreshSecret = os.Getenv("REFRESH_SECRET")
	cfg.privKeyFile = os.Getenv("PRIV_KEY_FILE")
	cfg.pubKeyFile = os.Getenv("PUB_KEY_FILE")
	return &cfg
}

func (c *config) Dsn() string {
	return c.dsn
}

func (c *config) DbName() string {
	return c.dbName
}

func (c *config) DbPort() int {
	return c.dbPort
}

func (c *config) AuthBaseURL() string {
	return c.authBaseURL
}

func (c *config) OrgBaseURL() string {
	return c.orgBaseURL
}

func (c *config) ChannelBaseURL() string {
	return c.channelBaseURL
}

func (c *config) MessageBaseURL() string {
	return c.messageBaseURL
}

func (c *config) ChatBaseURL() string {
	return c.chatBaseURL
}

func (c *config) HandlerTimeout() int64 {
	return int64(c.handlerTimeout)
}

func (c *config) IdTokenExp() int64 {
	return int64(c.idTokenExp)
}

func (c *config) RefreshTokenExp() int64 {
	return int64(c.refreshTokenExp)
}

func (c *config) RefreshSecret() string {
	return c.refreshSecret
}

func (c *config) PrivKeyFile() string {
	return c.privKeyFile
}

func (c *config) PubKeyFile() string {
	return c.pubKeyFile
}

func (c *config) AppPort() int {
	return c.appPort
}
