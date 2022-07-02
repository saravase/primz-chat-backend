package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/saravase/primz-chat-backend/config"
	"github.com/saravase/primz-chat-backend/db"
)

var (
	local bool
)

func init() {
	flag.BoolVar(&local, "local", true, "Run service in local")
	flag.Parse()
}

// @title           Primz Chat
// @version         1.0
// @description     This is chat application server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Saravanakumar Selvam
// @contact.url    http://www.swagger.io/support
// @contact.email  saravanakumar323py@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @tokenUrl https://example.com/oauth/token

// @host      localhost:5000
// @BasePath  /
func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error, While loading env variables. Reason : %v\n", err)
		}
	}

	cfg := config.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Error, While connecting mongodb. Reason : %v\n", err)
	}
	defer conn.Close()

	router, err := inject(cfg, conn.DB())

	if err != nil {
		log.Fatalf("Failure to inject data sources: %v\n", err)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.AppPort()),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error, While listening on port : %s\n", err)
		}
	}()
	log.Printf("Server listening on port: %s\n", srv.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server exiting")

}
