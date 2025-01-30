

docs: format
	find pkg -type d -exec bash -c "ls {}/*.go && godocdown -o ./{}/doc.md ./{}" \;

format:
	find . -name '*.go' -exec gofumpt -w -s -extra {} \;