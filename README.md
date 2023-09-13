# testcontainer demo

This repo illustrates the testcontainer issue when starting too many containers at once - and a fix for that.

## 1. Run without the fix:

Start multiple terminals and run `go test -count 1 -v .`. Observe how tests fail because of OCI proc issues.

## 2. Run with the fix: 

1. Run ./fix-docker-issue.sh
2. Start multiple terminals and run `go test -count 1 -v .`. Observe how everything works as expected