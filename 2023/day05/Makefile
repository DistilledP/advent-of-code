WATCH_PERIOD=2

run1:
	go run ./src/bin/part01.go

run1-debug:
	DEBUG=1 go run ./src/bin/part01.go

run2:
	go run ./src/bin/part02.go

run2-debug:
	DEBUG=1 go run ./src/bin/part02.go

watch:
	watch -n ${WATCH_PERIOD} go test -v ./src/lib

watch-debug:
	DEBUG=1 watch -n ${WATCH_PERIOD} go test -v ./src/lib