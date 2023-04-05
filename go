#!/bin/sh

docker run --rm -it -v "$PWD":/usr/src/myapp -w /usr/src/myapp mygolang:latest go $@
