#!/bin/bash
curl "$URL/v1/builds/$BUILD_ID?type=progress&name=manual-build-stage" -d ''
exit $?
