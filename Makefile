
build-linux:
	(export GOOS=linux; go build -o json2hcl-linux.bin)

build-macos:
	(export GOOS=darwin; go build -o json2hcl-darwin.bin)


build:	build-linux	build-macos