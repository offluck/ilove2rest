CURRDIR=$(shell pwd)
BINDIR=${CURRDIR}/bin
PACKAGE=github.com/offluck/ilove2rest/cmd/app

build: bindir
	go build -o ${BINDIR}/userapp ${PACKAGE}

test:
	go test ./...

run:
	go run ${PACKAGE} -config config/dev.yaml

bindir:
	mkdir -p ${BINDIR}

recreate-database:
	docker-compose down
	rm -rf pgdata
	docker-compose up database --build -d
