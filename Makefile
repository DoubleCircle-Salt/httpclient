NAME := httpclient
PACKAGE_NAME := github.com/DoubleCircle-Salt/httpclient
VERSION := `git describe --dirty`
COMMIT := `git rev-parse HEAD`

GOBUILD = go build -tags "full" -trimpath -ldflags="-s -w" -o httpclient.a -buildmode=c-archive httpclient.go

normal: clean httpclient

clean:
	rm -rf httpclient.a
	rm -rf httpclient.h

httpclient:
	$(GOBUILD)

