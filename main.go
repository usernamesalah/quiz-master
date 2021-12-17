package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/usernamesalah/quiz-master/internal/config"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/usernamesalah/quiz-master/app/controllers"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/usernamesalah/quiz-master/docs"
)

// @title Api Documentation for quiz master
// @version 1.0.0
// @description API documentation for quiz master

// @contact.name Rezi Apriliansyah
// @contact.email reziapriliansyah@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {
	cfg := config.AppConfig

	log.Println("Initializing the database connection ...")
	logger := logger.Default.LogMode(logger.Error)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.DB)
	dbConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger:                 logger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	log.Println("Initializing the web server ...")
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Request().Header.Set("Cache-Control", "max-age:3600, public")
			return next(c)
		}
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.Pre(middleware.Rewrite(map[string]string{
		"/v1/*": "/$1",
	}))

	e.Validator = &requestValidator{}

	// Utility endpoints
	e.GET("/docs/index.html", echoSwagger.WrapHandler)
	e.GET("/docs/doc.json", echoSwagger.WrapHandler)
	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.GET("/ping", ping)

	// init controllers
	controllers.InitAll(e, dbConnection)

	// Start server
	s := &http.Server{
		Addr:         "0.0.0.0:" + cfg.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	e.Logger.Fatal(e.StartServer(s))
}

type requestValidator struct{}

func (rv *requestValidator) Validate(i interface{}) (err error) {
	_, err = govalidator.ValidateStruct(i)
	return
}

// ping write pong to http.ResponseWriter.
func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
