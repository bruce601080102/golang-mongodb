package main

import (
	api "MongodbGO/pkg"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	var mydb = api.ConnectionDB()
	mydb.Login("root", "root")

	path, _ := os.Getwd()
	_ = godotenv.Load(path + "/conf/.env")
	CHANNELl_SECRET := os.Getenv("CHANNELl_SECRET")
	CHANNELlTOKEN := os.Getenv("CHANNELlTOKEN")
	bot, err := linebot.New(CHANNELl_SECRET, CHANNELlTOKEN)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()

	router.POST("/callback", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Print(err)
			}
			return
		}

		for _, event := range events {
			userID := event.Source.UserID
			var resmassage string
			result := api.User{}
			c := mydb.C("admin")
			argDB := api.Arg{
				Keys:   userID,
				C:      c,
				Result: result,
			}
			status := api.ReadDB(&argDB)

			switch {
			case status == 1:
				api.InsertDB(argDB)
				if err != nil {
					resmassage = fmt.Sprintf("存入失敗")
				} else {
					resmassage = fmt.Sprintf("已存入mongodb-來訪次數%d", 1)
				}
			case status == 0:
				times := argDB.Result.Times + 1
				fmt.Println(times)
				api.Update(times, argDB)
				resmassage = fmt.Sprintf("已存入mongodb-來訪次數%d", times)
			}

			fmt.Println("userID:", userID)
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					// resmassage := api.RecommandVtuber(message.Text)
					fmt.Println("--->", message.Text)
					if message.Text == "remove" {
						api.Remove(argDB)
						resmassage = "已刪除成功"
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(resmassage)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	router.Run(":8080")

}
