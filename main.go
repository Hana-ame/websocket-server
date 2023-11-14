package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	router := gin.Default()
	router.GET("/ws/", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer conn.Close()
		for {
			typ, data, e := conn.ReadMessage()
			if e != nil {
				err = e
				break
			}
			conn.WriteMessage(typ, data)
		}
		if err != nil {
			fmt.Println(err.Error())
			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		}
	})
	router.Run("127.114.5.14:8080")
}
