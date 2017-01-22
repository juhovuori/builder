for pkg in $(go list ./...)
do
  go test -coverprofile="$GOPATH/src/$pkg/coverage.out" $pkg
done
