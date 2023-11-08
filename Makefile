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

dev: recreate-database
	sleep 10
	make run

prod:
	docker-compose up --build -d

stop:
	docker-compose down

clean:
	rm -rf bin
	rm -rf pgdata

prune: stop clean
