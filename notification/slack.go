package notification

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PhotoManager/config"
)

const DefaultSlackTimeout = 5 * time.Second

var SlackClient SlackWebHookConfig

func NewSlackClient() {
	cfg := config.Parsed
	SlackClient = SlackWebHookConfig{
		WebHookUrl: cfg.PhotoManagerSlackWebHookURL,
		UserName:   cfg.PhotoManagerSlackUserName,
		Channel:    cfg.PhotoManagerSlackChannelName,
		TimeOut:    DefaultSlackTimeout,
	}
}

type SlackWebHookConfig struct {
	WebHookUrl string
	UserName   string
	Channel    string
	TimeOut    time.Duration
}

type SlackMessage struct {
	Channel     string       `json:"channel,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Attachment https://api.slack.com/reference/messaging/attachments#example
type Attachment struct {
	PreText MsgAction         `json:"pretext,omitempty"`
	Color   StatusColorAction `json:"color,omitempty"`
	Title   string            `json:"title,omitempty"`
	Footer  string            `json:"footer,omitempty"`
}

func (sc SlackWebHookConfig) NewSlackMessage(preText MsgAction) SlackMessage {
	attachment := Attachment{
		PreText: preText,
	}
	return SlackMessage{
		Channel:     sc.Channel,
		Attachments: []Attachment{attachment},
	}
}

func (sc SlackWebHookConfig) NewSlackMessageError(preText MsgAction, detail string) SlackMessage {
	attachment := Attachment{
		PreText: preText,
		Color:   StatusColorActionError,
		Title:   detail,
	}
	return SlackMessage{
		Channel:     sc.Channel,
		Attachments: []Attachment{attachment},
	}
}

func (sc SlackWebHookConfig) NewSlackMessageInfo(preText MsgAction, detail string) SlackMessage {
	attachment := Attachment{
		PreText: preText,
		Color:   StatusColorActionSuccess,
		Title:   detail,
	}
	return SlackMessage{
		Channel:     sc.Channel,
		Attachments: []Attachment{attachment},
	}
}

func (sc SlackWebHookConfig) SendMsg(sm SlackMessage) (err error) {
	return sc.sendHttpRequest(sm)
}

func (sc SlackWebHookConfig) sendHttpRequest(slackRequest SlackMessage) error {
	slackBody, _ := json.Marshal(slackRequest)
	req, errNewReq := http.NewRequest(
		http.MethodPost,
		sc.WebHookUrl,
		bytes.NewBuffer(slackBody),
	)
	if errNewReq != nil {
		return fmt.Errorf("err | creating Slack request: %v", errNewReq)
	}
	req.Header.Add("Content-Type", "application/json")
	if sc.TimeOut == 0 {
		sc.TimeOut = DefaultSlackTimeout
	}
	client := &http.Client{Timeout: sc.TimeOut}
	resp, errDo := client.Do(req)
	defer resp.Body.Close()
	if errDo != nil {
		return fmt.Errorf("err | sending Slack request: %v", errDo)
	}

	buf := new(bytes.Buffer)

	if _, errNewReq = buf.ReadFrom(resp.Body); errNewReq != nil {
		return errNewReq
	}
	// https://api.slack.com/messaging/webhooks#oauth_response
	if buf.String() != "ok" {
		return errors.New("err | non-ok response returned from Slack")
	}
	return nil
}
