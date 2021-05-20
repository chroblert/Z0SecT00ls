package jcmd

import (
	"github.com/chroblert/jgoutils/jlog"
	"github.com/chroblert/jgoutils/jnet/jhttp"
	"github.com/spf13/cobra"
)

var intrudeCommand = &cobra.Command{
	Use:"intrude",
	Short: "对http数据包进行暴破",
	Long: "对http数据包进行暴破",
	Run:func(cmd *cobra.Command,args []string){
		if isVerbose{
			jlog.SetLevel(jlog.DEBUG)
		}else{
			jlog.SetLevel(jlog.INFO)
		}
		// 执行命令
		if len(wordFiles) > 0 && reqFile != ""{
			jhttpobj := jhttp.New()
			jhttpobj.InitWithFile(reqFile)
			for _,v := range wordFiles{
				jhttpobj.SetWordfiles(v)
			}
			jhttpobj.SetIsUseSSL(isUseSSL)
			jhttpobj.SetProxy(proxy)
			jhttpobj.Intrude(isVerbose,func(statuscode int, headers map[string][]string, body []byte, err error) {
				//jlog.Info("\x1b[1A",statuscode, err)
				jlog.Info("",statuscode, err)

			})
		}
	},
}

//var headers = make([]string,0)
var(
	//reqFile string
	//isUseSSL bool
	//reqMethod string
	//reqUrl string
	//reqHeaders []string
	//reqDataStr string
	//proxy string
	//repeatCount int
	wordFiles []string
	)
func init(){

	intrudeCommand.Flags().StringVarP(&reqFile,"reqFile","r","req.txt","request file")
	intrudeCommand.MarkFlagRequired("reqFile")
	intrudeCommand.Flags().BoolVarP(&isUseSSL,"use-ssl","",false,"if is use https,then --use-ssl")
	intrudeCommand.Flags().StringVarP(&reqMethod,"reqMethod","m","","request method ,only GET or POST")
	//intrudeCommand.Flags().StringVarP(&reqUrl,"","u","","request url with query string")
	intrudeCommand.Flags().StringSliceVarP(&reqHeaders,"header","H",[]string{},"set request header.eg: -H Content-Type=application/json")
	//intrudeCommand.Flags().StringVarP(&reqDataStr,"reqData","d","","request body, u need set header manual")
	intrudeCommand.Flags().StringVarP(&proxy,"proxy","","","proxy (default value is \" \")")
	//intrudeCommand.Flags().IntVarP(&repeatCount,"count","c",1,"repeat count")
	// 设置字典文件
	intrudeCommand.Flags().StringSliceVarP(&wordFiles,"wordfile","w",[]string{},"word file")
	RootCmd.AddCommand(intrudeCommand)
}