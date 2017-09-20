#!/bin/bash
for i in `find . | grep go`; do gofmt -w -s $i; done
for i in  ass srt ttml wvtt mdvd; do cd $i; go test; cd ..; done
