# Mouscat

Mouscat will notify you if a string occur in a stream. You can pipe in the meantime you are reading multiple files. Awesome.

It acts like a `tail -f` from the STDinput and from given files.

The stache is everywhere...

## Motivations

Mouscat is a little app to make me learn golang <3 in order to have a better understanding about chans, io package and pluggable pattern. Also I wanted an easy to use application that can notify the user based on a string search with awesome notifictions. Currently supporting slack, terminal and discord

## How to use it

Pipe it with a command
`top | mouscat heisenberg`

Pipe it with a command and a file
`top | mouscat 123.123.45.2 access.log`

Pipe it with a command multiple files
`top | mouscat francis access.log auth.log`

Or just a single file
`top | mouscat malcom access.log auth.log`

## How to use the API

```golang
import "github.com/neonima/mouscat"
func main(){
    //set your plugins (callback)
    term := terminal.Notify

    //New returns a scanner type
    s := scanner.New("a Gopher in love")

    //Read your stream and call a plugin
    err := s.Listen(file, term)
}

```

To create a new notifier, you just need to follow this func signature `Notify func(data []byte) (err error)` and call it inside the `Listen()` func



## Next features

- Advanced query feature such as error detection based on signal and maybe regex

- gRPC client - server to handle notifications and log them (who said microservices quick reporting?)


## Improvment

Colorized detected line

Gracefully close streams, providing a Close() func

Better command line option  with documention without adding difficulty to use. Slack notification settings into a config file (avoiding long commad line). Same goes for discord (yep gamer here :D)

LICENSE MIT