package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/carlosfmreis/go-sample-micro-service/user"
	_ "github.com/go-sql-driver/mysql"
)

const port = ":8080"

const dataBaseHost = "localhost"
const dataBasePort = "3306"
const dataBaseUser = "root"
const dataBasePassword = "root"
const dataBaseName = "test"
const dbsource = dataBaseUser + ":" + dataBasePassword + "@tcp(" + dataBaseHost + ":" + dataBasePort + ")/" + dataBaseName

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	db, err := sql.Open("mysql", dbsource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	repository := user.NewRepository(db)

	service := user.NewService(repository)

	endpoints := user.MakeEndpoints(service)

	context := context.Background()
	handler := user.NewServer(context, endpoints)

	log.Println("Users Micro-Service listening on", port)

	setupCloseHandler()

	err = http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatal(err)
	}
}

func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Fatal("Users Micro-Service stopped")
		os.Exit(0)
	}()
}
