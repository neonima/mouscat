package scanner

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//Scanner is a Scanner that scan any input
type Scanner struct {
	Query string
	//buffleng int
}

//New returns a new scanner type
func New(query string) Scanner {
	return Scanner{
		Query: query,
	}
}

//Notify is callback method that defines what to Notify
type Notify func(data []byte) (err error)

//Listen a stream and notify if occurencies are found
func (s *Scanner) Listen(reader io.Reader, notifier ...Notify) error {
	p := bufio.NewReader(reader)
	for {
		n, _, _ := p.ReadLine()
		ok := strings.Contains(string(n), s.Query)
		if ok {
			for _, notif := range notifier {
				err := notif(n)
				if err != nil {
					return err
				}
			}
			colorized(n)
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
