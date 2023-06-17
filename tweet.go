package go_functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func getTwitterAPI() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_KEY_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	return api
}

func time2json(t time.Time) string {
	type timeStruct struct {
		Timestamp time.Time `json:"timestamp"`
	}
	b, _ := json.Marshal(timeStruct{t})
	return string(b)
}

func Tweet(writer http.ResponseWriter, request *http.Request) {
	api := getTwitterAPI()
	tweet, _ := api.PostTweet(time2json(time.Now()),nil)
	fmt.Println(tweet.Text)
}
