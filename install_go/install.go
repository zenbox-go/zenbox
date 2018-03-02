package install_go

import (
	"context"
	"errors"
	"fmt"
	"os"
)

var exitSetup = errors.New("取消安装")

func Setup(ctx context.Context) {
	runStep := func(m step) {
		err := m(ctx)
		if err == exitSetup {
			fmt.Fprintln(os.Stdout, err)
			os.Exit(0)
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	}

	runStep(welcome)
	runStep(checkGoInstalled)
	runStep(setupGo)
	runStep(setGOPATH)
	runStep(setupDone)
}
