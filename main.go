package main
import(
	"github.com/chroblert/Z0SecT00ls/jcmd"
	"github.com/chroblert/Z0SecT00ls/jvendor/grumble"
)
func main(){
	//jcmd.RootCmd.Execute()
	//jcmd.Execute()
	grumble.Main(jcmd.App)
}