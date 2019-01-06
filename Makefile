# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID=/tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## start: Start in development mode. Auto-starts when code changes.
start:
	~/Sites/go/bin/gin -i --port 4000 --appPort 8089 run main.go

## swagger-generate: Generate swagger.json from documentation.
swagger-generate:
	swagger generate spec -o swagger-api/v1/swagger.json && swagger validate swagger-api/v1/swagger.json

## swagger-start: Start swagger flavour serve.
swagger-flavour:
	swagger serve --flavor=swagger --port=6060 swagger-api/v1/swagger.json

## swagger-start: Start swagger flavour serve.
swagger-redoc:
	swagger serve --flavor=redoc --port=6060 swagger-api/v1/swagger.json

