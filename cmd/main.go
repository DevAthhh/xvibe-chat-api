package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/DevAthhh/xvibe-chat/initializers"
	"github.com/DevAthhh/xvibe-chat/internal/handler"
	ws "github.com/DevAthhh/xvibe-chat/internal/websocket"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDB()
	initializers.LoadConfig()

	env := viper.GetString("server.enviroment")
	if env == "developed" {
		gin.SetMode(gin.DebugMode)
	} else if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	host := viper.GetString("server.host")

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		srv := handler.NewServer(viper.GetString("server.api_port"), host)
		if err := srv.Start(); err != nil {
			log.Fatal(err)
		}
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		http.HandleFunc("/ws", ws.HandleWebsocket)
		if err := http.ListenAndServe(host+":"+viper.GetString("server.ws_port"), nil); err != nil {
			log.Fatal(err)
		}
	}(wg)

	wg.Wait()
}
