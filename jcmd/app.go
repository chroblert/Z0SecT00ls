package jcmd

import(
	"github.com/chroblert/Z0SecT00ls/jvendor/grumble"
	//"github.com/chroblert/jgoutils/jthirdutil/color"
	"github.com/fatih/color"
)

var App = grumble.New(&grumble.Config{
	Name:                  "foo",
	Description:           "An awesome foo bar",
	HistoryFile:           "/tmp/foo.hist",
	Prompt:                "foo » ",
	PromptColor:           color.New(color.FgGreen, color.Bold),
	HelpHeadlineColor:     color.New(color.FgGreen),
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,

	Flags: func(f *grumble.Flags) {
		f.String("d", "directory", "DEFAULT", "set an alternative root directory path")
		f.Bool("v", "verbose", false, "enable verbose mode")
	},
})

func init() {
	App.SetPrintASCIILogo(func(a *grumble.App) {
		a.Println("                   _   _     ")
		a.Println(" ___ ___ _ _ _____| |_| |___ ")
		a.Println("| . |  _| | |     | . | | -_|")
		a.Println("|_  |_| |___|_|_|_|___|_|___|")
		a.Println("|___|                        ")
		a.Println()
	})
}
