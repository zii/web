package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zii/web"

	"github.com/gorilla/mux"
	"github.com/zii/web/cmd/example1/method"
	"github.com/zii/web/service"
)

var args struct {
	Port int
}

func healthHandler(_ http.ResponseWriter, _ *http.Request) {
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)

	flag.IntVar(&args.Port, "port", 8866, "server listening port")
	rand.Seed(time.Now().Unix())
}

func main() {
	listenAddr := fmt.Sprintf("0.0.0.0:%d", args.Port)
	router := mux.NewRouter()
	// 健康检查
	router.HandleFunc("/health", healthHandler)

	service.RegisterMethod(router, "/v1/hello", method.Hello)

	// Interrupt handler.
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// HTTP transport.
	go func() {
		log.Println("listening:", listenAddr)
		err := http.ListenAndServe(listenAddr, router)
		web.Raise(err)
	}()

	// Run!
	log.Println("exit:", <-errc)
}
