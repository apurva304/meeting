package main

import (
	"context"
	"fmt"
	conf "gethelpnow/bin/config"
	repositories "gethelpnow/repositories/mongo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	// "github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	ser "gethelpnow/service"
)

var (
	ctx            = context.Background()
	configFileName = "config.json"
)

func main() {

	var logger *log.Logger
	{
		logger = log.New(os.Stderr, "", 0)
	}

	file, err := os.Open(configFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	config, err := conf.NewConfig(file)
	if err != nil {
		panic(err)
	}

	var dbUri string
	switch config.Mode {
	case conf.Dev:
		dbUri = fmt.Sprintf("mongodb://%s", config.Mongo.Address)
	case conf.Production:
		dbUri = fmt.Sprintf("mongodb://%s:%s@%s", config.Mongo.UserName, config.Mongo.Password, config.Mongo.Address)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)

	db := client.Database(config.Mongo.DBName, nil)

	repo := repositories.NewMeetingRepository(db)

	service := ser.NewService(repo)
	svc := ser.NewLoggingService(service, logger)

	r := mux.NewRouter()
	r.PathPrefix("/").Handler(ser.MakeHandler(svc))

	errs := make(chan error)

	httpAddr := fmt.Sprintf("localhost:%d", config.Port)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Println("transport", "HTTP", "addr", httpAddr)
		errs <- http.ListenAndServe(httpAddr, accessControl(r))
	}()

	logger.Println("exit", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			return
		}
		h.ServeHTTP(w, r)
	})
}
