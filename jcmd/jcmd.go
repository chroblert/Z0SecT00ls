package jcmd

import (
	"github.com/chroblert/jgoutils/jlog"
	"github.com/chroblert/Z0SecT00ls/jvendor/grumble"
)

func Execute(){

	if err := RootCmd.Execute(); err != nil {
		jlog.Fatal(err)
	}

}

func Main(){
	grumble.New(&grumble.Config{
		Name:                  "",
		Description:           "",
		Flags:                 nil,
		HistoryFile:           "",
		HistoryLimit:          0,
		NoColor:               false,
		Prompt:                "",
		PromptColor:           nil,
		MultiPrompt:           "",
		MultiPromptColor:      nil,
		ASCIILogoColor:        nil,
		ErrorColor:            nil,
		HelpHeadlineUnderline: false,
		HelpSubCommands:       false,
		HelpHeadlineColor:     nil,
	})
}