PROJECT_NAME := "composed"
APP_NAME := "composed"
PKG := "./"

.PHONY: download set-env run remove

download:
	mkdir -p ../task_repo
	git clone git@gitlab.com:greenteam1/task_repo.git ../task_repo
	mkdir -p ../executioner
	git clone git@gitlab.com:greenteam1/executioner.git ../executioner

set-env:
	cp ../executioner/env.example ../executioner/.env
	cp ../task_repo/env.example ../task_repo/.env
	cp ./env.example ./.env

run:
	docker-compose up

remove:
	docker-compose down -v --remove-orphans