package install_go

import (
	"context"
	"errors"
	"os"
	"zenbox/print"
)

const ContextKey = "ContexyKey"

var exitSetup = errors.New("取消安装")

func Setup(ctx context.Context) {
	c := context.WithValue(ctx, ContextKey, make(map[string]interface{}))

	runStep := func(m step) {
		err := m(c)
		if err == exitSetup {
			print.W(err)
			os.Exit(0)
		}

		if err != nil {
			print.E(err)
			os.Exit(2)
		}
	}

	runStep(welcome)
	runStep(checkGoInstalled)
	runStep(setupGo)
	runStep(setGOPATH)
	runStep(setupDone)
}
