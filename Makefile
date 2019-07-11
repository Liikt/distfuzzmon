all: build-client build-server

build-client:
	go build client/

build-server:
	go build server/

build-test-env: build-server-arm build-client-arm
	ansible-playbook ansible/fuzzy.yml

build-client-arm:
	GOARCH=arm64 go build client/

build-server-arm:
	GOARCH=arm64 go build server/
