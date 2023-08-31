export GOPATH="/code" # set GOPATH="C:\code" for windows
go doc github.com/EvgChu/keyboard
go doc greeting AllGreetings # information about function in package
go get github.com/EvgChu/keyboard
godoc -http=:6060
go test  hfg/chapter12/prose -v