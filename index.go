package main

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// router.GET("/deposits", func(c *gin.Context) {
	// 	log.Printf("hogehogehogehogehogehogehogehogehogehoge")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"deposits": []string{"a", "b"},
	// 	})
	// })

	router.POST("/v1/order", func(c *gin.Context) {
		type Order struct {
			FlowerID           string `json:"flower_id"`
			DestinationAddress string `json:"destination_address"`
			DestinationName    string `json:"destination_name"`
		}
		var order Order
		c.BindJSON(&order)

		postToSlack(order.FlowerID, order.DestinationAddress, order.DestinationName)

		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}

func postToSlack(flowerID string, destinationAddress string, destinationName string) {
	jsonStr := `{"text":"*注文が入りました！:rose:*\n\n花：` + flowerID + `\nあて先：` + destinationAddress + `\nあて名：` + destinationName + `", "icon_emoji": ":rose:"}`

	http.Post("https://hooks.slack.com/services/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", "application/json",
		bytes.NewBuffer([]byte(jsonStr)),
	)
}
