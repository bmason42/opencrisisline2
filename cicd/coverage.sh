#!/usr/bin/env bash

gocovmerge out/unit_coverage.out out/l2_client_coverage.out out/l2_server_coverage.out > out/coverage.out

gocover-cobertura < out/coverage.out > out/coverage.xml