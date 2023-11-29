package auth

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hosnibounechada/go-api/internal/auth/handlers"
	"github.com/hosnibounechada/go-api/internal/auth/routes"
	database "github.com/hosnibounechada/go-api/internal/db"
	"github.com/hosnibounechada/go-api/pkg/util"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func NewApp() (*App, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	postgresConnectionStr := util.BuildConnectionStr()

	db, err := database.InitDatabase(postgresConnectionStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	router := gin.Default()

	app := &App{
		Router: router,
		DB:     db,
	}

	v1 := app.Router.Group("/v1")

	routes.SetupAuthRoutes(v1, *handlers.NewAuthHandler())
	routes.SetupDevicesRoutes(v1, *handlers.NewDeviceHandler())

	return app, nil
}

func (app *App) Run() {
	app.Router.Run(":8080")
}
