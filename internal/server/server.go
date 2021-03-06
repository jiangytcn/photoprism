package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism/internal/context"
	log "github.com/sirupsen/logrus"
)

// Start the REST API server using the configuration provided
func Start(ctx *context.Context) {
	if ctx.HttpServerMode() != "" {
		gin.SetMode(ctx.HttpServerMode())
	} else if ctx.Debug() == false {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// Set template directory
	router.LoadHTMLGlob(ctx.HttpTemplatesPath() + "/*")

	registerRoutes(router, ctx)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", ctx.HttpServerHost(), ctx.HttpServerPort()),
		Handler: router,
	}

	quit := make(chan os.Signal)

	/*
		TODO: Use context for graceful shutdown of all services and add HTTP / TiDB server tests

		See
		- https://github.com/gin-gonic/gin/blob/dfe37ea6f1b9127be4cff4822a1308b4349444e0/examples/graceful-shutdown/graceful-shutdown/server.go
		- https://stackoverflow.com/questions/45500836/close-multiple-goroutine-if-an-error-occurs-in-one-in-go
	*/

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Info("received interrupt signal - shutting down")

		ctx.Shutdown()

		if err := server.Close(); err != nil {
			log.Errorf("server close: %s", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Info("web server closed")
		} else {
			log.Errorf("web server closed unexpect: %s", err)
		}
	}

	log.Info("please come back another time")

	time.Sleep(2 * time.Second)
}
