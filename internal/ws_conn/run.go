package ws_conn

import "net/http"

import "github.com/yuwe1/golangim/pkg/logger"

func StartWSServer(address string) {
	http.HandleFunc("/ws", ws)
	logger.Logger.Info("websocket serve start")
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
