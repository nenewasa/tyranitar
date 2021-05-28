package libtyranitar

import (
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/fatih/color"
)

func Execshell() {
	color.Blue("Test shell")
	shell := ishell.New()
	shell.SetPrompt(">>")
	shell.AddCmd(&ishell.Cmd{
		Name: "test",
		Help: "test command",
		Func: func(c *ishell.Context) {
			c.Println("test function", strings.Join(c.Args, " "))
		},
	})
	shell.Run()
}
