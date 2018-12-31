build:
	GO111MODULE=on go build -o dist/rebalance_tool

all:
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o dist/rebalance_tool.darwin.amd64
	GGO111MODULE=on OOS=linux GOARCH=amd64 go build -o dist/rebalance_tool.linux.amd64
	GGO111MODULE=on OOS=windows go get -u github.com/spf13/cobra
	GGO111MODULE=on OOS=windows GOARCH=amd64 go build  -o dist/rebalance_tool.windows.amd64.exe

