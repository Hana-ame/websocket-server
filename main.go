package main

import (
	"encoding/json"
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
	// c.WriteMessage(typ, data)
	// netwrokMsg
	c2s := &C2S_Message{}
	s2c := &S2C_Message{NetworkMessageID: SMsgID(c2s.NetworkMessageID)}
	if err := json.Unmarshal(data, c2s); err != nil {
		// error
		s2c.Load(Unknown, err.Error())
		SendMsg(c, typ, s2c)
		return
	}
	switch c2s.NetworkMessageID {
	case C2S_Regist:
		// receive
		s2c := Regist(c2s.Payload)
		SendMsg(c, typ, s2c)
		return
	case C2S_Login:
		// receive
		s2c := Login(c2s.Payload)
		SendMsg(c, typ, s2c)
		return
	default:
		s2c.Load(Unknown, "NetworkMessageID not supported")
		SendMsg(c, typ, s2c)
	}

}
