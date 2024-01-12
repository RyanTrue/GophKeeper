SHELL=bash
APP_VERSION=v1.0.0
.PHONY: install stop

install:
	docker-compose up --detach
	sleep 3
	go install -ldflags="-X 'github.com/RyanTrue/GophKeeper/cmd.version=$(APP_VERSION)' -X 'github.com/github.com/RyanTrue/GophKeeper/cmd.buildDate=$(shell date)'"

stop:
	docker-compose down
	docker image rm goph-keeper-server --force & docker image rm goph-keeper-migrate --force & docker image rm goph-keeper-server --force