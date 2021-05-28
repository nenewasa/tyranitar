package pupitar

import (
	"github.com/abiosoft/ishell"
	"github.com/fatih/color"

	"strings"
)

func ShellFunction() {

	shell := ishell.New()
	color.Red("GOSPLOIT V1")
	shell.SetPrompt("gosploit > ")

	shell.AddCmd(&ishell.Cmd{
		Name: "greet",
		Help: "greet user",
		Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
		},
	})

	shell.Run()

}
