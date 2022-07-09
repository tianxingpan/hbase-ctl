# Author: tianxingpan
# Date: 2022/07/08

version := $(shell /bin/date "+%Y-%m-%d %H:%M")

.PHONY: build
build:
	go build -o hbase_ctl -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" main.go
	$(if $(shell command -v upx), upx gofy)


.PHONY: clean
clean:
	rm hbase_ctl*