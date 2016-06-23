DEBUG_FLAG = $(if $(DEBUG),-debug)
COMMIT = $$(git describe --always)

ORG := github.com/tcnksm
REPO := $(ORG)/ghr
REPO_PATH :=$(GOPATH)/src/$(REPO)

clean:
	rm $(GOPATH)/bin/ghr
	rm bin/ghr

deps:
	go get -d -t -v ./...
	go get golang.org/x/tools/cmd/cover

install: deps
	go install -ldflags "-X main.GitCommit=\"$(COMMIT)\""

test: deps
	go test -v -timeout=30s -parallel=4 ./...
	go test -race ./...
	go vet .

cover: deps
	go test $(TEST) -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

# build generate binary on './bin' directory.
build: deps
	go build -ldflags "-X main.GitCommit=\"$(COMMIT)\"" -o bin/ghr

# package runs compile.sh to run gox and zip them.
# Artifacts will be generated in './pkg' directory
package: deps
	@sh -c "'$(CURDIR)/scripts/package.sh'"
