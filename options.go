package main

import (
	"github.com/neonima/mouscat/pkg/notifier"
	"github.com/neonima/mouscat/pkg/scanner"
)

type Configuration struct {
	BotName, LogoURL string
	Deamon           bool
	Options          struct {
		Offset  int
		Queries []scanner.Query
	}
	Notifiers []notifier.Notifier
}

func (c *Configuration) overrideConf(patterns []string, offset int) {
	if offset > 0 {
		c.Options.Offset = offset
	}

	if len(patterns) > 0 {
		c.Options.Queries = []scanner.Query{}
		for _, pattern := range patterns {
			c.Options.Queries = append(c.Options.Queries, scanner.Query{Pattern: pattern, Err: true})
		}
	}
}
