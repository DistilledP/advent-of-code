WATCH_PERIOD=2
DEBUG_ON=DEBUG=1
GO_FLAGS=GOGC=12800 GOMAXPROCS=60
GO_RUN=${GO_FLAGS} go run
GORUN_DEBUG=${DEBUG_ON} ${GO_RUN}

run1:
	@${GO_FLAGS} go run ./src/bin/part01.go

run1-debug:
	@${GORUN_DEBUG} ./src/bin/part01.go

run2:
	@${GO_RUN} ./src/bin/part02.go

run2-debug:
	@${GORUN_DEBUG} ./src/bin/part02.go

run2-batch:
	@${GO_RUN} ./src/bin/part02_batch.go

run2-batch-debug:
	@${GORUN_DEBUG} ./src/bin/part02_batch.go

watch:
	@watch -n ${WATCH_PERIOD} go test -v ./src/lib

watch-debug:
	@${DEBUG_ON} watch -n ${WATCH_PERIOD} go test -v ./src/lib

test:
	@go test -v ./src/lib

test-debug:
	@${DEBUG_ON} go test -v ./src/lib
