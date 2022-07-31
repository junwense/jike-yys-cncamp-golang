export tag=v1.0
export docker_registry = 13697916905

root:
	export ROOT=https://github.com/junwense/jike-yys-cncamp-golang/tree/myworkbranch/examples/module3/mytest

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t cncamp/httpserver:${tag} .

push: release
	echo "pushing cncamp/httpserver"
	docker push ${docker_registry}/httpserver:${tag}
