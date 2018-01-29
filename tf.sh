#!/bin/bash
for i in `find . | grep go`; do gofmt -w -s $i; done
for i in  ass srt ttml wvtt mdvd; do cd $i; go get -v -t -d ./...;  go test; cd ..; done
