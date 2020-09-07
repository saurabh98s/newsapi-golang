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
func (n *News) FetchNewsHeadlines(w http.ResponseWriter, r *http.Request) ([]models.Articles, error) {
	var response models.News

	fmt.Println("GET FetchNewsHeadlines")

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

	return response.Articles, nil

}

// FetchAndRenderSearchBar reads the form input and searches the matching query
func (n *News) FetchAndRenderSearchBar(w http.ResponseWriter, r *http.Request) {
	groupError := "[ERROR] FETCH SEARCH BAR"
	var searchResults *models.News


	searchData := r.FormValue("search-bar")
	n.l.Info("Got Data", searchData)
	
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?q=" + searchData + "&apiKey=29d795ddf1c040778350321158922cd3")
	if err != nil {
		n.l.Error(groupError + err.Error())
	}

	defer resp.Body.Close()

	readBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		n.l.Error(groupError + err.Error())
	}

	err = json.Unmarshal([]byte(readBody), &searchResults)

	if err != nil {
		n.l.Error(groupError + err.Error())
	}

	if searchResults.TotalResults < 5 {

		err = tpl.ExecuteTemplate(w, "404.html", "Not Enough newsworthy")
		if err != nil {
			n.l.Error(err)
		}
	} else {
		err = tpl.ExecuteTemplate(w, "search.html", searchResults.Articles[0:5])
		if err != nil {
			n.l.Error(err)
		}
	}

}
