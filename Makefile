SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

all: embed build

deps:
	go get github.com/astaxie/beego
	go get github.com/astaxie/beego
  go get github.com/fzzy/radix/redis
  go get github.com/barakmich/glog
  go get github.com/jmcvetta/neoism

embed:

build:

test:

install:

clean:

dpkg:

run:
