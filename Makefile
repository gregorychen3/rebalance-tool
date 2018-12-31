build:
	go build

all:
	GOOS=darwin GOARCH=amd64 go build -o rebalance_tool.${GOOS}.${GOARCH}
	GOOS=linux GOARCH=amd64 go build -o rebalance_tool.${GOOS}.${GOARCH}
	GOOS=windows go get -u github.com/spf13/cobra
	GOOS=windows GOARCH=amd64 go build  -o rebalance_tool.${GOOS}.${GOARCH}

