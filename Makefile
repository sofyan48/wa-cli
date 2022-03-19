.PHONY: build clean build-darwin build-linux install

build:
	go build -o bin/orn src/main.go

clean:
	rm -rf ./bin
release:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=$(ARCH) go build -a -o bin/orn.darwin-$(ARCH) src/main.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=$(ARCH) go build -a -o bin/orn.windows-$(ARCH) src/main.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/orn.linux-$(ARCH) src/main.go

build-darwin:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=$(ARCH) go build -a -o bin/orn.darwin-$(ARCH) src/main.go

build-linux:
	env CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/orn.linux-$(ARCH) src/main.go

install:
	cp bin/orn /usr/local/bin/
	chmod +x /usr/local/bin/orn


