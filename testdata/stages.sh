#!/bin/bash
which curl >~/moe
echo curl $URL/v1/builds/$BUILD_ID?type=progress\&name=manual-build-stage>>~/moe

curl "$URL/v1/builds/$BUILD_ID?type=progress&name=manual-build-stage" -d ''
exit $?
