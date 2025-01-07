package app

import (
	"net/http"
	"time"

	"github.com/fatjan/gogomanager/internal/config"
	"github.com/fatjan/gogomanager/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewServer(cfg *config.Config, db *sqlx.DB) *http.Server {
	router := gin.Default()
	handlers.SetupRouter(cfg, db, router)

	return &http.Server{
		Addr:         cfg.App.Port,
		Handler:      router,
		WriteTimeout: time.Second * 600,
		ReadTimeout:  time.Second * 600,
		IdleTimeout:  time.Second * 600,
	}
}
