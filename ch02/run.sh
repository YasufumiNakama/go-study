# gofmt
gofmt -l -s -w .

# test
go env -w GO111MODULE=auto

# ex01
echo "===== run ex01 ====="
cd ./ex01/tmpconv
go test -run ""
