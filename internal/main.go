package internal

import (
	"log"

	"github.com/gin-gonic/gin"
	authHandler "github.com/hosnibounechada/go-api/internal/auth/handlers"
	authRoutes "github.com/hosnibounechada/go-api/internal/auth/routes"
	database "github.com/hosnibounechada/go-api/internal/db"
	productHandler "github.com/hosnibounechada/go-api/internal/product/handlers"
	productRoutes "github.com/hosnibounechada/go-api/internal/product/routes"
	"github.com/hosnibounechada/go-api/pkg/util"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func NewApp() (*App, error) {

	postgresConnectionStr := util.BuildConnectionStr()

	db, err := database.InitDatabase(postgresConnectionStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	router := gin.Default()

	app := &App{
		Router: router,
		DB:     db,
	}

	v1 := app.Router.Group("/v1")

	authRoutes.SetupAuthRoutes(v1, *authHandler.NewAuthHandler())
	authRoutes.SetupDevicesRoutes(v1, *authHandler.NewDeviceHandler())
	productRoutes.SetupProductsRoutes(v1, *productHandler.NewProductHandler())

	return app, nil
}

func (app *App) Run() {
	app.Router.Run(":8080")
}
