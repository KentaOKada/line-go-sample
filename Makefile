TARGET         := main
RUN            ?= ''

ifneq ($(RUN), '')
	RUN_OPT = -run $(RUN)
endif

help:             ## このヘルプを表示します
	@awk -F: '/^[A-Za-z0-9_-]+:.*## / { sub(/.*## /, "", $$2); printf "make %-15s %s\n", $$1, $$2 }' Makefile

build:            ## アプリケーションをビルドします
	go build -o "$(TARGET)" "./cmd/$(TARGET)"

run: build        ## アプリケーションをビルドして実行します
	"./$(TARGET)"

update:            ## ライブラリのアップデート
	go mod download
	go mod tidy
