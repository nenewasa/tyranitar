package libtyranitar

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/abiosoft/ishell"
)

func ExecShell() {
	color.Blue("Test shell")
	shell := ishell.New()
	shell.SetPrompt(">>")
	shell.AddCmd(&ishell.Cmd{
		Name: "test",
		Help: "test command"
		Func: func(c *ishell.Context) {
			c.Println("test function". stirngs.Join(c.Args, " "))
		},
	})
	shell.Run()
}
