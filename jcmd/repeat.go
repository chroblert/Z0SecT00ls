package jcmd

import (
	"github.com/chroblert/jgoutils/jlog"
	"github.com/chroblert/jgoutils/jnet/jhttp"
	"github.com/spf13/cobra"
)

var repeatCommand = &cobra.Command{
	Use:   "repeat",
	Short: "对http数据包进行重放",
	Long:  "对http数据包进行重放",
	Run: func(cmd *cobra.Command, args []string) {
		if isVerbose {
			jlog.SetLevel(jlog.DEBUG)
		} else {
			jlog.SetLevel(jlog.INFO)
		}

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
	},
}

//var headers = make([]string,0)
var (
	reqFile     string
	isUseSSL    bool
	reqMethod   string
	reqUrl      string
	reqHeaders  []string
	reqDataStr  string
	proxy       string
	repeatCount int
)

func init() {

	repeatCommand.Flags().StringVarP(&reqFile, "reqFile", "r", "req.txt", "request file")
	repeatCommand.Flags().BoolVarP(&isUseSSL, "use-ssl", "", false, "if is use https,then --use-ssl")
	repeatCommand.Flags().StringVarP(&reqMethod, "reqMethod", "m", "", "request method ,only GET or POST")
	repeatCommand.Flags().StringVarP(&reqUrl, "", "u", "", "request url with query string")
	repeatCommand.Flags().StringSliceVarP(&reqHeaders, "header", "H", []string{}, "set request header.")
	repeatCommand.Flags().StringVarP(&reqDataStr, "reqData", "d", "", "request body, u need set header manual")
	repeatCommand.Flags().StringVarP(&proxy, "proxy", "", "", "proxy (default value is \" \")")
	//repeatCommand.Flags().IntVarP(&repeatCount,"count","c",1,"repeat count")
	RootCmd.AddCommand(repeatCommand)
}
