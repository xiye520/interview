for t in 1 2 3 4; do GOMAXPROCS=4 go run maps.go $t; done