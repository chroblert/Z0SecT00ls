package http

import (
	"github.com/chroblert/jgoutils/jconfig"
	"github.com/chroblert/jgoutils/jlog"
	"github.com/chroblert/jgoutils/jnet/jhttp"
	"github.com/chroblert/jgoutils/jthirdutil/github.com/desertbit/grumble2"
)

var intrudeCommand = &grumble.Command{
	Name:      "intrude",
	FullPath: "basic/http/intrude",
	Aliases:   nil,
	Help:      "intrude http/https request ",
	LongHelp:  "",
	HelpGroup: "Basic",
	Usage:     "",
	Flags: func(f *grumble.Flags) {
		f.String("r", "reqFile", "req.txt", "request file")
		f.Bool("", "use-ssl", false, "if is use https,then --use-ssl")
		f.String("m", "reqMethod", "", "request method ,only GET or POST")
		f.String("u", "reqUrl", "", "request url with query string")
		//repeatCommand.Flags().StringSliceVarP(&reqHeaders,"header","H",[]string{},"set request header.")
		f.StringSlice("H", "header", []string{}, "set request header.")
		f.String("d", "reqDataStr", "", "request body, u need set header manual")
		f.String("", "proxy", "", "proxy (default value is \" \")")
		//f.Int("c", "count", 1, "repeat count")
		f.StringSlice("w","wordFile",[]string{}," set word file")
	},
	Run: func(c *grumble.Context) error {
		wordFiles := c.Flags.StringSlice("wordFile")
		reqFile := c.Flags.String("reqFile")
		if len(wordFiles) >0 && reqFile != ""{
			jhttpobj := jhttp.New()
			jhttpobj.InitWithFile(reqFile)
			for _,v := range wordFiles{
				jhttpobj.SetWordfiles(v)
			}
			jhttpobj.SetIsUseSSL(c.Flags.Bool("use-ssl"))
			jhttpobj.SetProxy(c.Flags.String("proxy"))
			jhttpobj.Intrude(false, func(statuscode int, headers map[string][]string, body []byte, err error) {
				jlog.Info("statuscode:",statuscode)
			})
		}
		return nil
	},
	Completer: nil,
}

func init() {
	//jcmd.App.AddCommand(repeatCommand)
	//jconfig.Get("app").(*grumble.App).AddCommand(repeatCommand)
	var tmpCommands []*grumble.Command
	if jconfig.Get("jcommands") == nil{
		tmpCommands = make([]*grumble.Command,0)
	}else{
		tmpCommands = jconfig.Get("jcommands").([]*grumble.Command)
	}
	tmpCommands = append(tmpCommands,intrudeCommand)
	jconfig.Set("jcommands",tmpCommands)
}