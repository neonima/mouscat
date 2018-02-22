package scanner

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/neonima/mouscat/pkg/notifier"
	log "github.com/sirupsen/logrus"
)

//Scanner is a Scanner that scan any input
type Scanner struct {
	Queries []Query
}

type Query struct {
	Pattern string
	Err     bool `json:"Error"`
}

type occurence struct {
	occ []byte
	err bool
}

//New returns a new scanner type
func New(queries []Query) Scanner {
	return Scanner{
		Queries: queries,
	}
}

//Notify is callback method that defines what to Notify

//Listen a stream and notify if occurencies are found
func (s *Scanner) Listen(reader io.Reader, offset int, notifiers ...notifier.Notification) error {
	sight := make(chan occurence, offset)
	defer close(sight)
	log.Info(offset)
	p := bufio.NewReader(reader)

	go func(notifiers []notifier.Notification) {
		//TODO returns error
		select {
		case c := <-sight:
			for _, notif := range notifiers {
				notif.Notify(c.occ, c.err)
			}
		}

	}(notifiers)

	for {

		n, _, _ := p.ReadLine()
		//if err != nil {
		//	return err
		//}
		for _, query := range s.Queries {
			if query.Contains(string(n)) {
				sight <- occurence{n, query.Err}
			}

		}

		if len(n) != 0 {
			fmt.Println(string(n))
		}

	}
}

func colorized(b []byte) {
	c := string(b)
	b = []byte(fmt.Sprintf("%v%v%v", `\033[0;31m`, c, `\x1b[0m`))

}

//Contains is a wrapper around Query.Contains() that iterate over each
// queries

func (o *Query) Contains(search string) bool {
	return strings.Contains(search, o.Pattern)
}
