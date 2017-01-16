#!/bin/bash

commit=$(git rev-parse HEAD)
ts=$(date -Is)
echo -n "{\"commit\":\"$commit\", \"build-time\":\"$ts\"}"
