package notifier

import (
	"github.com/neonima/mouscat/pkg/notifier/slack"
)

//Notifier represents a pluggable notification tool
type Notifier struct {
	slack slack.Slack
}

type Notification interface {
	Notify(data []byte, isError bool) error
}
