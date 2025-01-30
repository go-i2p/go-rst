

docs: format
	find pkg -type d -exec bash -c "ls {}/*.go && godocdown -o ./{}/doc.md ./{}" \;

format:
	find . -name '*.go' -exec gofumpt -w -s -extra {} \;

build:
	go build -o go-rst .

test: build
	./go-rst -rst example/doc.rst -out example/example.html
	./go-rst -rst example/complexDoc.rst -out example/complexExample.html
	./go-rst -rst example/rstSpec.rst -out example/rstSpec.html
