package port

import (
	"github.com/chroblert/jgoutils/jasync"
	"github.com/chroblert/jgoutils/jconfig"
	"github.com/chroblert/jgoutils/jnet/jparser"
	"github.com/chroblert/jgoutils/jnet/jtcp"
	"github.com/chroblert/jgoutils/jnet/jtcp/jcore"
	"github.com/chroblert/jgoutils/jthirdutil/github.com/desertbit/grumble2"
	"strconv"
	"sync"
)

var portScanCommand = &grumble.Command{
	Name:      "portscan",
	FullPath: "infocollect/port/portscan",
	Aliases:   nil,
	Help:      "port scan. only show open ports ",
	LongHelp:  "",
	HelpGroup: "InfoCollect",
	Usage:     "",
	Flags: func(f *grumble.Flags) {
		f.String("","ip","","ip.eg:ip,ip/24,ip1-ip2")
		f.String("","port","80","ports.eg:port1,port2-port3")
		f.Int("","rate",1000,"send packet number per second.default:5000,max:8000")
	},
	Run: func(c *grumble.Context) error {
		//t := c.Flags.String("ip")
		t := jparser.ParseIPStr(c.Flags.String("ip"))
		p := jparser.ParsePortStr(c.Flags.String("port"))
		rate := c.Flags.Int("rate")
		jtcpobj := jtcp.New()
		if jtcpobj != nil {
			jcore.ShowNetworks()
			//jtcpobj.SetNetwork(2)
			jasyncobj := jasync.New()
			var wg = new(sync.WaitGroup)
			for _, v := range t {
				//jlog.Info(v)
				for _, v2 := range p {
					wg.Add(1)
					go func(v string, v2 int) {
						jasyncobj.Add(v+":"+strconv.Itoa(v2), jtcpobj.SinglePortSYNScan, nil, v, uint16(v2), "test")
						wg.Done()
					}(v, v2)
				}
			}
			wg.Wait()
			//jlog.Debug("sleep 100 second")
			//time.Sleep(100*time.Second)
			jasyncobj.Run(rate)
			//jtcpobj.RecvScanRes()

			jasyncobj.Wait()
			//jtcpobj.Test()
			jasyncobj.Clean()
			//jtcpobj.Test()
			//time.Sleep(3*time.Second)
			jtcpobj.CloseHandle()
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
	tmpCommands = append(tmpCommands,portScanCommand)
	jconfig.Set("jcommands",tmpCommands)
}