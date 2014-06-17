package aspira

import (
    "fmt"
    "./util"
    "strings"
    "net/url"
    "net/http"
    "io/ioutil"
)

func Ask(question string) string {
    serviceUrl := "http://www.ed.conpet.gov.br/mod_perl/bot_gateway.cgi?"
    serviceParams := "server=0.0.0.0:8085&charset_post=utf-8&pure=1&js=0&tst=1&msg="+question
    endpoint := serviceUrl+serviceParams

    response, err := http.Get(endpoint)
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
    endpoint := "https://.slack.com/services/hooks/incoming-webhook?token="
    body := fmt.Sprintf(`{
        "channel": "#general",
        "username": "aspira",
        "text": "%s",
        "icon_url": "http://sinistersaladmusikal.files.wordpress.com/2011/01/100_9002.jpg?w=50"
    }`, util.RemoveHtml(message))

    response, _  := http.PostForm(endpoint, url.Values{"payload": {body}})
    defer response.Body.Close()
}

