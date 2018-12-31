build:
	go build

all:
	GOOS=darwin GOARCH=amd64 go build -o dist/rebalance_tool.darwin.amd64
	GOOS=linux GOARCH=amd64 go build -o dist/rebalance_tool.linux.amd64
	GOOS=windows go get -u github.com/spf13/cobra
	GOOS=windows GOARCH=amd64 go build  -o dist/rebalance_tool.windows.amd64.exe

