build:
	go build

all:
	GOOS=darwin GOARCH=amd64 go build -o dist/rebalance_tool -installsuffix
	GOOS=linux GOARCH=amd64 go build -o dist/rebalance_tool -installsuffix
	GOOS=windows go get -u github.com/spf13/cobra
	GOOS=windows GOARCH=amd64 go build  -o dist/rebalance_tool -installsuffix

