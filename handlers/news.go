package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"newsbackend/models"

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

// FetchNews returns the fetched news in form of JSON
func (n *News) FetchNews(w http.ResponseWriter, r *http.Request) {

	var response models.News
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?country=in&apiKey=29d795ddf1c040778350321158922cd3")
	if err != nil {
		n.l.Error(err.Error())
	}
	defer resp.Body.Close()
	readBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		n.l.Error(err.Error())
	}
	err = json.Unmarshal([]byte(readBody), &response)
	if err != nil {
		n.l.Error(err.Error())
	}
	fmt.Fprintf(w,response.Articles[0].Description)

}
