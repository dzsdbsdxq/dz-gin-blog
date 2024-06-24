package initialize

import (
	"context"
	"fmt"
	"github.com/dzsdbsdxq/dz-gin-blog/app/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(engine *gin.Engine) {
	host := global.G_DZ_CONFIG.System.Host
	port := global.G_DZ_CONFIG.System.Port
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.G_DZ_LOG.Fatalf("listen: %s\n", err)
		}
	}()
	global.G_DZ_LOG.Info(fmt.Sprintf("Server is running at %s:%d", host, port))
	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.G_DZ_LOG.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.G_DZ_LOG.Fatal("Server forced to shutdown:", err)
	}

	global.G_DZ_LOG.Info("Server exiting!")
}
