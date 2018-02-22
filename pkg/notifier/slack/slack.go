package slack

import "github.com/ashwanthkumar/slack-go-webhook"

//Slack Represents a slack notifier
type Slack struct {
	att     slack.Attachment
	pay     slack.Payload
	isError bool
	webhook string
}

//New return a new slack instance
func New(sa slack.Attachment, sp slack.Payload, webhook string) *Slack {
	t := Slack{
		att:     sa,
		pay:     sp,
		webhook: webhook,
	}
	return &t
}

//Notify implements a notifier for the scanner
func (s Slack) Notify(data []byte) (err error) {
	slack.Send(s.webhook, "", s.pay)
	return nil
}
