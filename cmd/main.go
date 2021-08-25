package main

import (
	"ToDoListApp"
	"ToDoListApp/pkg/handler"
	"ToDoListApp/pkg/repository"
	"ToDoListApp/pkg/servies"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err !=nil {
		log.Fatalf("error initializing configs", err.Error())
	}
	if err := godotenv.Load(); err != nil{
		log.Fatalf("error loading env variables: %s", err.Error())
	}


	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err!=nil{
		log.Fatalf("failed initialization db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := servies.NewService(repos)
	handlers := handler.NewHandler(services)

	srv:= new(ToDoListApp.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error", err.Error())
	}
}

func initConfig() error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}