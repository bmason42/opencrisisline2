#!/usr/bin/env bash
if [ -z "$GOPATH" ]
then
      mkdir -p /gomods
      export GOPATH=/gomods
fi
export PATH=$GOPATH/bin:$PATH
export TEST_SCRIPTS=$(pwd)/scripts
mkdir -p out
go test -v -coverpkg=github.com/bmason42/opencrisisline2/pkg/... -coverprofile=out/unit_coverage.out github.com/bmason42/opencrisisline2/pkg/...  2>&1 | go-junit-report > out/report_l1.xml

