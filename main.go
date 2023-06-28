package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {

	// Initial Bot & Whebhook

	bot := tbot.New(os.Getenv("API_KEY"), tbot.WithWebhook("https://b42b-2001-8f8-1471-2b9e-e26e-6eb8-7142-dfac.in.ngrok.io", ":8100"))
	c := bot.Client()
	fmt.Println("Server Running 8100")

	//Get API Data

	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin%2Cethereum%2Csolana%2Cripple%2Cfilecoin&vs_currencies=usd")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)

	}

	// Parsing- Unmarshal Data
	var objMap map[string]interface{}
	err = json.Unmarshal((body), &objMap)

	if err != nil {
		fmt.Println("Json Error Decode ")
	}

	// Bot Handler

	bot.HandleMessage("Price", func(m *tbot.Message) {

		Price1 := objMap["bitcoin"].(map[string]interface{})["usd"]
		str := fmt.Sprint("Today BTC Price is ", Price1, "$")
		c.SendMessage(m.Chat.ID, str)

		if err != nil {
			panic(err)
		}
		fmt.Println("Message Sent ")
	})
	log.Fatal(bot.Start())
}
