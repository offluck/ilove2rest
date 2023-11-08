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
	docker-compose -f dev.docker-compose.yaml up --build -d

dev: recreate-database
	echo "We need to wait 'till database is locked and loaded. Let me take a little nap, Z... Z... Z..."
	sleep 60
	make run

prod:
	docker-compose up --build -d

stop:
	docker-compose down

clean:
	rm -rf bin
	rm -rf pgdata

prune: stop clean
