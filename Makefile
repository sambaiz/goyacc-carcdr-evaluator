.PHONY: parser test

parser:
	cd parser && goyacc -p "expr" ../expr.y

test:
	go test ./...
