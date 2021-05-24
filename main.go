package main

import (
	"syscall"
	"time"

	"github.com/getlantern/systray"
	"github.com/peterxj1/disable-screenssaver/icon"
)

var (
	kernel32                    = syscall.NewLazyDLL("kernel32.dll")
	procSetThreadExecutionState = kernel32.NewProc("SetThreadExecutionState")
)

const (
	title            = "Disable screensaver"
	EsSystemRequired = 0x00000001
)

func main() {

	systray.Run(onReady, nil)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle(title)
	systray.SetTooltip(title)
	quit := systray.AddMenuItem("Quit", "Quit")

	go func() {
		<-quit.ClickedCh
		systray.Quit()
	}()

	pulse := time.NewTicker(time.Second * 10)

	for {
		select {
		case <-pulse.C:
			procSetThreadExecutionState.Call(uintptr(EsSystemRequired))
		}
	}

}
