test:
	go test -run=. -cover
	go1.10rc1 test -run=. -cover
	go1.9 test -run=. -cover

bench:
	go test -run=. -cover -bench=. -benchmem
	go1.10rc1 test -run=. -cover -bench=. -benchmem
	go1.9 test -run=. -cover -bench=. -benchmem

setup:
	go get golang.org/x/build/version/go1.10rc1 && go1.10rc1 download
	go get golang.org/x/build/version/go1.9 && go1.9 download

