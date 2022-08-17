include $(.env)
export

test:
	go test ./pkg/* -v