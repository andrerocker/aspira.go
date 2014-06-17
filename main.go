package main

import (
    "fmt"
    "net/http"
    "./aspira"
    "./aspira/util"
    "github.com/go-martini/martini"
)

func handleMessage(request *http.Request) {
    user := request.FormValue("user_name")
    question := util.Clean(request.FormValue("text"))
    response := aspira.Ask(question)

    fmt.Println("user: "+user+" question: "+question+" response: "+response)
    aspira.Say("@"+user+": "+response)
}

func main() {
    martini := martini.Classic()
    martini.Get("/", handleMessage)
    martini.Post("/", handleMessage)
    http.ListenAndServe("0.0.0.0:8081", martini)
}
