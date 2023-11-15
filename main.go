package main

import (
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var ops atomic.Uint64
var syncMap = sync.Map{}

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
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer conn.Close()
		handler(conn)
	})
	router.POST("/upload", upload)
	router.Run("127.114.5.14:8080")
}

func handler(c *websocket.Conn) error {
	op := ops.Add(1)
	syncMap.Store(op, c)
	defer syncMap.Delete(op)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
	}()

	for {
		if typ, data, err := c.ReadMessage(); err != nil {
			c.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
			return err
		} else {
			handleMessage(c, typ, data)
		}
	}
}

func handleMessage(c *websocket.Conn, typ int, data []byte) {
	// echo
	c.WriteMessage(typ, data)
}
