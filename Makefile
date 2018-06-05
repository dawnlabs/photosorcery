all: build_default

build_default: deps xgo
	xgo --targets=windows/amd64,darwin/amd64,linux/amd64 .

build_all: deps xgo
	xgo .

build_native: deps
	go build

deps:
	go get

xgo: xgo_docker
	go get github.com/karalabe/xgo

xgo_docker: verify_docker_running	
	docker pull karalabe/xgo-latest	

verify_docker_running:
	docker info > /dev/null

clean:
	rm photosorcery
	rm -f *amd64*
