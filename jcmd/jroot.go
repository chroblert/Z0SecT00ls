package jcmd

import (
	"github.com/chroblert/jgoutils/jconfig"
	"github.com/chroblert/jgoutils/jlog"
	"github.com/spf13/cobra"
)

var cfgFile string
var isVerbose bool

var RootCmd = &cobra.Command{
	Use:"Z0SecT00ls",
	Short: "Zer0ne Sec T00ls",
	Long: "A tools that contains some attack tool....developing",
	Run:func(cmd *cobra.Command,args []string) {
		if isVerbose{
			jlog.SetLevel(jlog.DEBUG)
		}else{
			jlog.SetLevel(jlog.INFO)
		}
		jlog.Debug("初始化配置文件")
		jconfig.InitWithFile(cfgFile)
		jlog.Debug("配置文件加载成功")
		// 如果没有输入任何flag，则输出帮助信息
		if len(args) == 0 {
			cmd.Help()
			return
		}

	},
}


func init(){
	RootCmd.PersistentFlags().StringVar(&cfgFile,"config","conf/config.json","config file")
	RootCmd.PersistentFlags().BoolVarP(&isVerbose,"verbose","v",false,"verbose msg")
}