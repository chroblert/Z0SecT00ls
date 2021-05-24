package jcmd

import (
	"github.com/chroblert/Z0SecT00ls/jvendor/github.com/desertbit/grumble"
	"github.com/chroblert/jgoutils/jlog"
	"plugin"
)

var useCommand = &grumble.Command{
	Name:      "use",
	Aliases:   nil,
	Help:      "用于切换要使用的命令 ",
	LongHelp:  "",
	HelpGroup: "",
	Usage:     "",
	Args: func(a *grumble.Args) {
		a.String("commandPath","命令所在文件的路径")
	},
	Run: func(c *grumble.Context) error {
		jlog.Info(c.Args.String("commandPath"))
		// 判断插件是否存在
		// 使用plugin打开插件
		plug,err := plugin.Open(c.Args.String("commandPath"))
		if err != nil{
			jlog.Error(err)
			return err
		}
		Main,err := plug.Lookup("Main")
		if err != nil{
			jlog.Error(err)
			return err
		}
		Main.(func())()
		return nil
	},
	Completer: nil,
}

func init() {
	App.AddCommand(useCommand)
}

