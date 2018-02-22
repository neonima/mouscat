package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/neonima/mouscat/pkg/notifier/terminal"

	"github.com/Tkanos/gonfig"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/neonima/mouscat/pkg/scanner"
	log "github.com/sirupsen/logrus"
)

func init() {
	//TODO IF linux / Darwing
	// Creating config folder/file into home direction
	// into .mouscat/config
	dir, _ := homedir.Dir()
	mouscdir := fmt.Sprintf("%v/.mouscat/", dir)
	os.Mkdir(mouscdir, 0755)
	conf := fmt.Sprintf("%vconfig.json", mouscdir)
	if _, err := os.Stat(conf); os.IsNotExist(err) {
		ioutil.WriteFile(conf, []byte("{}"), 0755)
	}

}

func main() {
	conf := Configuration{}

	dir, _ := homedir.Dir()
	mous := fmt.Sprintf("%v/.mouscat/config.json", dir)
	err := gonfig.GetConf(mous, &conf)
	if err != nil {
		log.Warnf("Problem while opening the config file: %v", err)
	}

	offset := flag.Int("offset", 10, "Defines the number of line after the occurence that should be taken into consideration")
	pattern := flag.String("p", "", "The pattern to be matched")
	flag.Parse()
	var args = flag.Args()
	patterns := strings.Split(*pattern, " ")
	log.Debug(patterns)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	conf.overrideConf(patterns, *offset)
	if len(conf.Options.Queries) == 0 {
		log.Fatal("please a string to search")
		os.Exit(1)
	}
	for _, file := range args {
		f, err := os.Open(file)
		defer f.Close()
		if err != nil {
			log.Panic(err)
		}
		go func() {
			log.Fatal(readFile(f, conf.Options.Queries, *offset))
		}()
	}
	go func() {
		log.Fatal(readFile(os.Stdin, conf.Options.Queries, *offset))
		defer os.Stdin.Close()
	}()
	<-sigs

	log.Println("mouscat exited...")
	os.Exit(0)

}

func readFile(file *os.File, query []scanner.Query, offset int) error {
	term := terminal.Terminal{}
	s := scanner.New(query)
	err := s.Listen(file, offset, term)
	return err
}
