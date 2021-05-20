package jcmd

import (
	"github.com/chroblert/Z0SecT00ls/jvendor/github.com/desertbit/grumble"
	"github.com/chroblert/jgoutils/jlog"
)

var repeat2Command = &grumble.Command{
	Name:      "Repeat",
	Aliases:   nil,
	Help:      "repeat http/https request ",
	LongHelp:  "",
	HelpGroup: "",
	Usage:     "",
	Flags: func(f *grumble.Flags) {
		f.String("r", "reqFile", "req.txt", "request file")
		f.Bool("", "use-ssl", false, "if is use https,then --use-ssl")
		f.String("m", "reqMethod", "", "request method ,only GET or POST")
		f.String("u", "reqUrl", "", "request url with query string")
		//repeatCommand.Flags().StringSliceVarP(&reqHeaders,"header","H",[]string{},"set request header.")
		f.StringSlice("H", "header", []string{}, "set request header.")
		f.String("d", "reqData", "", "request body, u need set header manual")
		f.String("", "proxy", "", "proxy (default value is \" \")")
		repeatCommand.Flags().IntVarP(&repeatCount, "count", "c", 1, "repeat count")
	},
	Args: func(a *grumble.Args) {
		a.String("r", "request file")
	},
	Run: func(c *grumble.Context) error {
		//jlog.Debug(c.Args.String("r"))
		//jlog.Debug("Flag,",c.Flags.String("reqFile"))
		jlog.Debug("Flag,", c.Flags.StringSlice("header"))
		return nil
	},
	Completer: nil,
}

func init() {
	App.AddCommand(repeat2Command)
}
