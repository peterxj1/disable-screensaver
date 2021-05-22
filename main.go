package main

import (
	"time"

	"github.com/getlantern/systray"
	"github.com/micmonay/keybd_event"
	"github.com/peterxj1/disable-screenssaver/icon"
	"github.com/peterxj1/disable-screenssaver/win/mutex"
)

const (
	title = "Disable screensaver"
)

func main() {
	id, err := mutex.CreateMutex(title)
	if err != nil {
		return
	}
	defer mutex.ReleaseMutex(id)

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

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	kb.SetKeys(keybd_event.VK_F15)

	ticker := time.NewTicker(59 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				kb.Release()
			case <-quit.ClickedCh:
				ticker.Stop()
				return
			}
		}
	}()
}
