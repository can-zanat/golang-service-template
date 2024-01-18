package main

import (
	"fmt"
	"main/config"
	"main/internal"
	store "main/persistent"
	"os"

	logger "github.com/can-zanat/gologger"
	_ "github.com/go-sql-driver/mysql"
)

const serverPort = ":80"

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	appConfig, err := config.New()
	if err != nil {
		return err
	}

	logger := logger.NewWithLogLevel("info")
	defer func() {
		err = logger.Sync()
		if err != nil {
			fmt.Println(err)
		}
	}()

	repository := store.NewMysqlStore(appConfig.Mysql)
	service := internal.NewService(repository)
	handler := internal.NewHandler(service)

	New(serverPort, handler, logger).Run()

	return nil
}
