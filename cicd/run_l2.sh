#!/usr/bin/env bash
if [ -z "$GOPATH" ]
then
      mkdir -p /gomods
      export GOPATH=/gomods
fi
export PATH=$GOPATH/bin:$PATH
export TEST_SCRIPTS=$(pwd)/scripts
go get github.com/jstemmer/go-junit-report
mkdir -p out
go test -v -coverpkg=github.com/bmason42/opencrisisline2/pkg/... -coverprofile=out/l2_server_coverage.out app/servermain_test.go &
export SERVER_PID=$!
echo "Server PID $SERVER_PID"
sleep 5
echo "Start RUNNING L2 Tests"
go test -v -coverpkg=github.com/bmason42/opencrisisline2/pkg/... -coverprofile=out/l2_client_coverage.out github.com/bmason42/opencrisisline2/tests/l2/...  2>&1 | go-junit-report > out/report_l2.xml
echo "Finish RUNNING L2 Tests"
kill $SERVER_PID
echo "Dont care if kill fails, the image is halting"

