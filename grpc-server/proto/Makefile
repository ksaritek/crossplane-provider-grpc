LOCAL_BIN := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))/.bin

format:
	buf format -w

clean:
	rm -rf gen

generate: format clean
	buf generate
	