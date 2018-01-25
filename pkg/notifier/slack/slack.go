package slack

import "github.com/ashwanthkumar/slack-go-webhook"

type pslack struct {
	att     slack.Attachment
	pay     slack.Payload
	webhook string
}

//New return a new slack instance
func New(sa slack.Attachment, sp slack.Payload, webhook string) *pslack {
	t := pslack{
		att:     sa,
		pay:     sp,
		webhook: webhook,
	}
	return &t
}

//Notify implements a notifier for the scanner
func (s *pslack) Notify(data []byte) (err error) {
	slack.Send(s.webhook, "", s.pay)
	return nil
}
