package util

import (
    "fmt"
    "regexp"
    "strings"
    "net/url"
)

func Clean(message string) string {
   return strings.TrimSpace(strings.Split(message, "aspira:")[1])
}

func RemoveHtml(message string) string {
    pattern := regexp.MustCompile("<[^<>]*>")
    return pattern.ReplaceAllString(message, "")
}

func AskEndpoint(question string) string {
    params := url.Values{}
    params.Add("server", "0.0.0.0:8085")
    params.Add("charset_post", "utf-8")
    params.Add("pure", "1")
    params.Add("js", "0")
    params.Add("tst", "1")
    params.Add("msg", question)

    endpoint, _ := url.Parse("http://www.ed.conpet.gov.br/mod_perl/bot_gateway.cgi?")
    endpoint.RawQuery = params.Encode()

    return endpoint.String()
}

func ResponseEndpoint() string {
    endpoint := "https://.slack.com/services/hooks/incoming-webhook?token="
    return endpoint
}

func ResponsePostBody(message string) url.Values {
    body := fmt.Sprintf(`{
        "channel": "#hacker",
        "username": "aspira",
        "text": "%s",
        "icon_url": "http://sinistersaladmusikal.files.wordpress.com/2011/01/100_9002.jpg?w=50"
    }`, RemoveHtml(message))

    return url.Values { "payload": {body} }
}
