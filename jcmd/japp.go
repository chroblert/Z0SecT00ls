package jcmd

import (
	_ "github.com/chroblert/Z0SecT00ls/japp"
	"github.com/chroblert/jgoutils/jconfig"
	"github.com/chroblert/jgoutils/jfile"
	"github.com/chroblert/jgoutils/jlog"
	"github.com/chroblert/jgoutils/jthirdutil/github.com/desertbit/grumble2"
	"github.com/fatih/color"
	"plugin"
)

var App = grumble.New(&grumble.Config{
	Name:                  "Z0SecT00ls",
	Description:           "a set of sec tools",
	HistoryFile:           "/tmp/foo.hist",
	Prompt:                "Z0SecT00ls » ",
	PromptColor:           color.New(color.FgGreen, color.Bold),
	HelpHeadlineColor:     color.New(color.FgGreen),
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,
	Flags: func(f *grumble.Flags) {
		f.String("d", "directory", "DEFAULT", "set an alternative root directory path")
		f.Bool("v", "verbose", false, "enable verbose mode")
		f.String("c", "config", "conf/config.json", "config file")
		f.Bool("p","use-plugin",false,"if use plugin ,then --use-plugin")
	},
	CurrentCommand: "app",
})

func init() {
	jconfig.Set("app",App)
	App.SetPrintASCIILogo(func(a *grumble.App) {
		jlog.Warn("=============================")
		jlog.Warn("        Z0SecT00ls           ")
		jlog.Warn("=============================")
	})
	App.OnInit(func(a *grumble.App, flags grumble.FlagMap) error {
		if flags.Bool("verbose") {
			jlog.SetLevel(jlog.DEBUG)
		} else {
			jlog.SetLevel(jlog.INFO)
		}
		// initial jconfig
		jlog.Debug("初始化配置文件")
		jconfig.InitWithFile(flags.String("config"))
		jlog.Debug("配置文件加载成功")
		// load commands
		if jconfig.Get("jcommands") != nil{
			for _,v := range jconfig.Get("jcommands").([]*grumble.Command) {
				a.AddCommand(v)
			}
		}
		// load plugins
		if flags.Bool("use-plugin"){
			// 枚举并加载所有的插件
			jlog.Debug("加载插件")
			filenames, _ := jfile.GetFilenamesByDir("jplugin/jhttp")
			for _, filename := range filenames {
				jlog.Debug(filename)
				plug, err := plugin.Open(filename)
				if err != nil {
					jlog.Fatal(err)
				}
				plugMain, err := plug.Lookup("Main")
				if err != nil {
					jlog.Fatal(err)
				}
				a.AddCommand(plugMain.(func() *grumble.Command)())
			}
		}

		return nil
	})

}
