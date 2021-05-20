package jcmd

import (
	"github.com/chroblert/jgoutils/jlog"
)

func Execute(){

	if err := RootCmd.Execute(); err != nil {
		jlog.Fatal(err)
	}

}
