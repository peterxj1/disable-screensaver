# Disable screen saver on windows

## Build 
```cmd
go build -ldflags "-H=windowsgui"
```

## Run
Application will create a systray icon with gopher logo. It will release key f15 every 59 seconds to prevent screen saver activation
