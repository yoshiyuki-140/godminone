.PHONY: register login logout getTask getAllTask createTask set unset

register:
	go build
	./godminone register --name kuro --password pass

login:
	go build
	./godminone login --name kuro --password pass

logout:
	go build
	./godminone logout

getTask:
	go build
	./godminone getTask --id 2

getAllTask:
	go build
	./godminone getTask

createTask:
	go build
	./godminone createTask --task "buy a Milk"


set:
	go build
	./godminone set --file settings.json

unset:
	go build
	./godminone unset --file settings.json