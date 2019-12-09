package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/github-user/server/database/database"
)

//handle a request with method GET and path "/api/".
func filmsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		req.ParseForm()
		page := 1
		w.Write([]byte("{\"result\" : \n["))
		if req.Form["page"] != nil {
			page, _ = strconv.Atoi(req.Form["page"][0])
		}
		count := 0
		for i := 1; ; i++ {
			item := database.GetValue([]byte("films"), []byte(strconv.Itoa(i)))
			if len(item) != 0 {
				count++
				if count > pagelen*(page-1) {
					w.Write([]byte(item))
					if count >= pagelen*page || count >= database.GetBucketCount([]byte("films")) {
						break
					}
					w.Write([]byte(", \n"))
				}
			}
		}
		w.Write([]byte("]\n}"))
	}
}

func getFilmsById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	data := database.GetValue([]byte("films"), []byte(vars["id"]))
	w.Write([]byte(data))
}

func filmsPagesHandler(w http.ResponseWriter, req *http.Request) {
	data := database.GetBucketCount([]byte("films"))
	w.Write([]byte(strconv.Itoa(data)))
}
