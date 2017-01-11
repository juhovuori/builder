#!/bin/bash

commit=$(git rev-parse HEAD)
echo -n "{\"commit\":\"$commit\"}"
