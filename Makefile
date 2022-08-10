GOPKG ?=	moul.io/asanaman
DOCKER_IMAGE ?=	moul/asanaman
GOBINS ?=	.
NPM_PACKAGES ?=	.

include rules.mk

generate: install
	GO111MODULE=off go get github.com/campoy/embedmd
	mkdir -p .tmp
	echo 'foo@bar:~$$ asanaman' > .tmp/usage.txt
	asanaman 2>&1 >> .tmp/usage.txt
	embedmd -w README.md
	rm -rf .tmp
.PHONY: generate
