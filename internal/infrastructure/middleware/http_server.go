package middleware

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go-admin/configs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitHttpServer(handler http.Handler) {
	srv := &http.Server{
		Addr:    configs.C.Web.Addr,
		Handler: handler,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Infof("listen: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of ShutdownTime seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Infof("Shutting down server...")

	// The context is used to inform the server it has ShutdownTime seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), configs.C.Web.ShutdownTime*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Infof("Server forced to shutdown:", err)
	}
	logrus.Infof("Server exited")
}
