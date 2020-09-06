package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"newsbackend/models"

	"newsbackend/config"

	"github.com/sirupsen/logrus"
)

// News handler for fetching news
type News struct {
	l *logrus.Logger
}

// NewInstanceNews  returns a news handler with the given dependencies
func NewInstanceNews(l *logrus.Logger) *News {
	return &News{l}
}

// FetchNewsHeadlines returns the fetched news in form of JSON
func (n *News) FetchNewsHeadlines(w http.ResponseWriter, r *http.Request) ([]models.Articles,error){
	var response models.News

	n.l.Info("GET FetchNewsHeadlines")

	groupError := "[ERROR] FETCH NEWS"

	// TODO: fix reading from .env file
	apiKey := config.APIKey()

	if apiKey == " " {
		n.l.Error(groupError + " " + apiKey)
	}
	fmt.Println("Key Found: ", apiKey)

	
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?country=in&apiKey=29d795ddf1c040778350321158922cd3" + apiKey)
	if err != nil {
		n.l.Error(groupError + err.Error())
	}

	defer resp.Body.Close()
	readBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		n.l.Error(groupError + err.Error())
	}

	err = json.Unmarshal([]byte(readBody), &response)

	if err != nil {
		n.l.Error(groupError + err.Error())
	}

	// var arr []models.News

	return response.Articles,nil

}

// // FetchBreakingNews about
// func  (n *News) FetchBreakingNews(w http.ResponseWriter, r *http.Request){
// 	groupError:="[ERROR] Fetching Breaking News"
// 	var response *models.News

// 	resp:=http.Get("")

// }
