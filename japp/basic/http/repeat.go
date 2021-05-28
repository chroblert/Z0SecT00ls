package http

import (
	"github.com/chroblert/jgoutils/jconfig"
	"github.com/chroblert/jgoutils/jlog"
	"github.com/chroblert/jgoutils/jnet/jhttp"
	"github.com/chroblert/jgoutils/jthirdutil/github.com/desertbit/grumble2"
)

var repeatCommand = &grumble.Command{
	Name:      "repeat",
	FullPath: "basic/http/repeat",
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
		f.StringSlice("H", "header", []string{"xxx"}, "set request header.")
		f.String("d", "reqDataStr", "", "request body, u need set header manual")
		f.String("", "proxy", "", "proxy (default value is \" \")")
		f.Int("c", "count", 1, "repeat count")
	},
	Run: func(c *grumble.Context) error {
		jlog.Debug("Flag,", c.Flags.StringSlice("header"))
		reqMethod := c.Flags.String("reqMethod")
		reqFile := c.Flags.String("reqFile")
		reqUrl := c.Flags.String("reqUrl")
		reqDataStr := c.Flags.String("reqDataStr")
		proxy := c.Flags.String("proxy")
		repeatCount := c.Flags.Int("count")
		isUseSSL := c.Flags.Bool("use-ssl")
		// 执行命令
		if reqMethod != "" {
			jhttpobj := jhttp.New()
			jhttpobj.SetReqMethod(reqMethod)
			jhttpobj.SetURL(reqUrl)
			jhttpobj.SetReqData(reqDataStr)
			jhttpobj.SetProxy(proxy)
			jhttpobj.Repeat(repeatCount)
		} else {
			jhttpobj := jhttp.New()
			jhttpobj.InitWithFile(reqFile)
			jhttpobj.SetIsUseSSL(isUseSSL)
			jhttpobj.SetProxy(proxy)
			jhttpobj.Repeat(repeatCount)
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
	tmpCommands = append(tmpCommands,repeatCommand)
	jconfig.Set("jcommands",tmpCommands)
}