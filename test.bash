#!/bin/bash

cd db && mockery -name=Db && cd -
cd record && cp normtest/record_test.go . && echo "test     " && go test testmock/record -covermode=count -coverprofile /tmp/r.cov && sleep 1 \
 	  && cp mocktest/record_test.go . && echo "mock test" && go test testmock/record -covermode=count -coverprofile /tmp/r.cov && cd - 
