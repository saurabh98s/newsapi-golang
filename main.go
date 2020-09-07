package main

import (
	"fmt"
	"net/http"

	"newsbackend/handlers"
	"newsbackend/logger"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	handlers.Init()
	router := mux.NewRouter()
	news := handlers.NewInstanceNews(&logrus.Logger{})
	router.HandleFunc("/news",news.RenderMainPage)
	router.HandleFunc("/search",news.FetchAndRenderSearchBar)
	logger.Log.Info("[DEBUG] Starting Server on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
