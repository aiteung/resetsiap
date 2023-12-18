package siappkg

import (
	"encoding/json"

	"github.com/imroc/req/v3"
)

func SendEmailTo(EmailURL, to, subject, body string, attachment ...Files) (res bool) {
	data := Email{
		From:        "ITeung Artificial Intellegence<iteung@ulbi.ac.id>",
		To:          to,
		Subject:     subject,
		Body:        body,
		Attachments: attachment,
	}

	resp := new(StatusEmail)

	jsonString, err := json.Marshal(data)
	if err != nil {
		return
	}

	client := CreateClientHTTP()
	ugent, err := client.SetUserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.120 Safari/537.36").
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Encoding", "gzip, deflate").
		SetHeader("Cache-Control", "max-age=0").
		SetHeader("Connection", "keep-alive").
		SetHeader("Accept-Language", "en-US,en;q=0.8,id;q=0.6").
		SetBody(jsonString).
		Post(EmailURL)

	if err != nil {
		return
	}

	err = json.Unmarshal(ugent.Bytes(), resp)
	if err != nil {
		return
	}

	res = resp.Status != ""

	return
}

func CreateClientHTTP() *req.Client {
	return req.
		C().
		SetJsonUnmarshal(json.Unmarshal).
		SetJsonMarshal(json.Marshal)
}
