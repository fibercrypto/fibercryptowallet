#!/bin/bash

#Set Script Name variabl
SCRIPT=`basename ${BASH_SOURCE[0]}`

MODE="emulator"
TIMEOUT="60m"
# run go test with -v flag
VERBOSE=""
# run go test with -run flag
RUN_TESTS=""
FAILFAST=""
NAME=""
AUTO_PRESS_BUTTONS=""

usage () {
  echo "Usage: $SCRIPT"
  echo "Optional command line arguments"
  echo "-m <string>  -- Testmode to run, EMULATOR or USB;"
  echo "-r <string>  -- Run test with -run flag"
  echo "-n <string>  -- Specific name for this test, affects coverage output files"
  echo "-v <boolean> -- Run test with -v flag"
  echo "-f <boolean> -- Run test with -failfast flag"
  echo "-a <boolean> -- Auto press buttons in emulator mode"
  exit 1
}

while getopts "h?m:r:n:vfa" args; do
case $args in
    h|\?)
        usage;
        exit;;
    m ) MODE=${OPTARG};;
    r ) RUN_TESTS="-run ${OPTARG}";;
    n ) NAME="-${OPTARG}";;
    v ) VERBOSE="-v";;
    f ) FAILFAST="-failfast";;
    a ) AUTO_PRESS_BUTTONS="1";;
  esac
done


set +e

HW_GO_INTEGRATION_TESTS=1 HW_GO_INTEGRATION_TEST_MODE=$MODE AUTO_PRESS_BUTTONS=$AUTO_PRESS_BUTTONS \
    go test -count=1 ./src/cli/integration/... $FAILFAST -timeout=$TIMEOUT $VERBOSE $RUN_TESTS

TEST_FAIL=$?

exit $TEST_FAIL
