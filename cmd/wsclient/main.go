package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var patterns = []map[string]interface{}{
	{
		"action": "wakeup",
	},
	{
		"action": "set-cart",
		"data": []map[string]interface{}{
			{
				"code":     "010",
				"quantity": 1,
			},
		},
	},
	{
		"action": "done",
	},
	// // Paid
	// {
	// 	"action": "confirm",
	// 	"data": map[string]interface{}{
	// 		"payment_channel": "promptpay",
	// 	},
	// },
	{
		"action": "confirm",
		"data": map[string]interface{}{
			"payment_channel": "creditcard",
		},
	},
	// // Free
	// {
	// 	"action": "confirm",
	// 	"data": map[string]interface{}{
	// 		"payment_channel": "free",
	// 	},
	// },
	// {
	// 	"action": "request-otp",
	// 	"data": map[string]interface{}{
	// 		"mail": "ztrixack.th@gmail.com",
	// 	},
	// },
	// {
	// 	"action": "check-otp",
	// 	"data": map[string]interface{}{
	// 		"mail":      "ztrixack.th@gmail.com",
	// 		"reference": "DjHvId",
	// 		"otp":       "628702",
	// 	},
	// },
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	go func() {
		for i, pattern := range patterns {
			write(pattern, i, conn)

			time.Sleep(5 * time.Second)
		}
	}()

	for {
		read(conn)
		time.Sleep(100 * time.Millisecond)
	}
}

func write(pattern map[string]interface{}, i int, conn *websocket.Conn) {
	b2, _ := json.Marshal(pattern)
	fmt.Println("FRONTEND >", string(b2))

	err := conn.WriteJSON(pattern)
	if err != nil {
		log.Println(err)
	}
}

func read(conn *websocket.Conn) {
	var read map[string]interface{}
	err := conn.ReadJSON(&read)
	if err != nil {
		log.Println(err)
		return
	}

	b1, _ := json.Marshal(read)
	fmt.Println("_BACKEND >", string(b1))
}
