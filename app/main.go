package main

import (
	"fmt"
	"go-hexagonal/config"

	migration "go-hexagonal/modules/migration"

	emailQueue "go-hexagonal/modules/email"

	"os"

	"github.com/streadway/amqp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newDatabaseConnection(config *config.AppConfig) *gorm.DB {

	configDB := map[string]string{
		"DB_Username": os.Getenv("GOHEXAGONAL_DB_USERNAME"),
		"DB_Password": os.Getenv("GOHEXAGONAL_DB_PASSWORD"),
		"DB_Port":     os.Getenv("GOHEXAGONAL_DB_PORT"),
		"DB_Host":     os.Getenv("GOHEXAGONAL_DB_ADDRESS"),
		"DB_Name":     os.Getenv("GOHEXAGONAL_DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])

	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	migration.InitMigrate(db)

	return db
}

func newQueueConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return conn
}

func main() {
	//load config if available or set to default
	// config := config.GetConfig()

	//initialize database connection based on given config
	// dbConnection := newDatabaseConnection(config)

	queueConnection := newQueueConnection()

	// //initiate user repository
	// userRepo := userRepository.NewGormDBRepository(dbConnection)

	// //initiate user service
	// userService := userService.NewService(userRepo)

	// //initiate user controller
	// userController := userController.NewController(userService)

	emailQueue.NewQueueRepository(queueConnection)

	// emailService.NewService(emailQueue)

	//initiate pet repository
	// petRepo := petRepository.NewGormDBRepository(dbConnection)

	// //initiate pet service
	// petService := petService.NewService(petRepo)

	// //initiate pet controller
	// petController := petController.NewController(petService)

	// //initiate auth service
	// authService := authService.NewService(userService)

	// //initiate auth controller
	// authController := authController.NewController(authService)

	// //create echo http
	// e := echo.New()

	// //register API path and handler
	// api.RegisterPath(e, authController, userController, petController)

	// run server
	// go func() {
	// 	address := fmt.Sprintf("localhost:%d", config.AppPort)

	// 	if err := e.Start(address); err != nil {
	// 		log.Info("shutting down the server")
	// 	}
	// }()

	// // Wait for interrupt signal to gracefully shutdown the server with
	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit

	// // a timeout of 10 seconds to shutdown the server
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// if err := e.Shutdown(ctx); err != nil {
	// 	log.Fatal(err)
	// }
}
