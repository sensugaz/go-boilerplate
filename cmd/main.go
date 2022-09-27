package main

import (
	"log"

	"github.com/iamgoangle/gokub/env/required"
	"gitlab.bitkub.io/bo/gokub-boilerplate/cmd/repository"
	"gitlab.bitkub.io/bo/gokub-boilerplate/cmd/services"

	"github.com/go-playground/validator/v10"
	"gitlab.bitkub.io/bo/gokub-boilerplate/cmd/handler"

	"github.com/labstack/echo/v4"
)

var (
	appPort string
	appName string

	//mongodbHost     string
	//mongodbUsername string
	//mongodbPassword string
	//
	//redisHost     string
	//redisUsername string
	//redisPassword string
)

func initAppConfigs() {
	appPort = required.GetEnv("APP_PORT")
	appName = required.GetEnv("APP_NAME")

	log.Print("[initAppConfigs]: initial app config success") // DO NOT USING THIS LOGGER IN PRODUCTION
}

func initRedisConfigs() {

}

func initMongoConfigs() {

}

// ==============================================
// ===== DO NOT USE THIS CODE IN PRODUCTION =====
// ==============================================
func main() {
	initAppConfigs()
	initRedisConfigs()
	initMongoConfigs()

	log.Printf("== started %s app is works ==", appName)

	e := echo.New()
	e.Debug = false
	e.Validator = &handler.CustomValidator{
		Validator: validator.New(),
	}

	mgoRepo, err := repository.NewMongoDB(&repository.MongoConfigs{})
	if err != nil {
		panic(err)
	}

	userService, err := services.NewUserService(mgoRepo)
	if err != nil {
		panic(err)
	}

	h, err := handler.New(e, appPort, userService)
	if err != nil {
		panic(err)
	}

	err = h.RunServer()
	if err != nil {
		panic(err)
	}
}
