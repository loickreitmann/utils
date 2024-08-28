#!/usr/bin/env bash
mkdir coverage
go test . -coverprofile=coverage/utils_test_coverage.out
go tool cover -html=coverage/utils_test_coverage.out  