package httputil

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SetupGracefulStop(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	Shutdown(srv)
}

func Shutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("[http server shutdown err:]", err)
	}

	select {
	case <-ctx.Done():
		log.Println("[http server exit timeout of 5 seconds.]")
	default:

	}

	log.Printf("[http server exited.]")
}
