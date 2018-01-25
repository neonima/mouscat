package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/neonima/mouscat/pkg/notifier/terminal"
	"github.com/neonima/mouscat/pkg/scanner"
)

func main() {
	flag.Parse()
	var args = flag.Args()
	if len(args) == 0 {
		log.Fatal("please a string to search")
		os.Exit(1)
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for _, file := range args[1:] {
		f, err := os.Open(file)
		defer f.Close()
		if err != nil {
			log.Panic(err)
		}
		go func() {
			log.Fatal(readFile(f, args[0]))
		}()
	}
	go func() {
		log.Fatal(readFile(os.Stdin, args[0]))
		defer os.Stdin.Close()
	}()
	<-sigs
	log.Println("mouscat exited...")
	os.Exit(0)

}

func readFile(file *os.File, query string) error {
	term := terminal.Notify
	s := scanner.New(query)
	err := s.Listen(file, term)
	return err
}
