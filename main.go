package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"web_test/csvdecoder"
	"github.com/gin-gonic/gin"
	"github.com/braintree/manners"
)

func main() {
	router := setupRouter()
	manners.ListenAndServe("localhost:8080", router)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.LoadHTMLGlob("templates/*.html")

	persons := csvdecoder.Convert("ueba.csv")

	router.GET("/main", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", nil)
	})

	router.GET("/", func(c *gin.Context) {
		idsStr := c.QueryArray("id") 
		if idsStr[0] == "-1" {
			manners.Close()
			c.HTML(http.StatusOK, "main.html", gin.H{
				"content": "end",
			})
		}  else {
			output, err := get(persons,idsStr)
			if err !=nil {
				c.Status(400)
			}
			c.HTML(http.StatusOK, "main.html", gin.H{
				"content": output,
			})
		}
	})
	return router
}

func get(persons []csvdecoder.Person, ids []string) ([]string, error) {
	var output []string
	idsArray := strings.Split(ids[0], " ")
	for i := 0; i < len(persons); i++ {
		for _, id := range idsArray {
			id, err := strconv.Atoi(id)
			if err != nil {
				log.Println(err)
				return output, err
			}
			if persons[i].Id == id {
				data, err := json.Marshal(persons[i])
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			output = append(output, string(data)) 
			}
		}
	}
	return output, nil
}
