package aspira

import (
    "./util"
    "strings"
    "net/http"
    "io/ioutil"
)

func Ask(question string) string {
    response, err := http.Get(util.AskEndpoint(question))
    defer response.Body.Close()

    if err != nil {
        return "q fita"
    }

    message, err := ioutil.ReadAll(response.Body)

    if err != nil {
        return "olha a sonoridade!"
    }

    return strings.TrimSpace(string(message))
}

func Say(message string) {
    endpoint := util.ResponseEndpoint()
    response, _  := http.PostForm(endpoint, util.ResponsePostBody(message))
    defer response.Body.Close()
}
