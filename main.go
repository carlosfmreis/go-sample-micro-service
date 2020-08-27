package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/carlosfmreis/sample-microservice/user"
	_ "github.com/go-sql-driver/mysql"
)

const port = "8080"

const dataBaseHost = "localhost"
const dataBasePort = "3306"
const dataBaseUser = "root"
const dataBasePassword = "root"
const dataBaseName = "test"
const dbsource = dataBaseUser + ":" + dataBasePassword + "@tcp(" + dataBaseHost + ":" + dataBasePort + ")/" + dataBaseName

func main() {
	httpAddress := flag.String("http", ":"+port, "http listen address")

	var db *sql.DB
	{
		var err error
		db, err = sql.Open("mysql", dbsource)
		defer db.Close()
		if err != nil {
			os.Exit(-1)
		}

	}

	flag.Parse()
	context := context.Background()
	var service user.Service
	{
		repository := user.NewRepository(db)
		service = user.NewService(repository)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := user.MakeEndpoints(service)

	go func() {
		fmt.Println("Microservice listening on ", *httpAddress)
		handler := user.NewServer(context, endpoints)
		errs <- http.ListenAndServe(*httpAddress, handler)
	}()
}
