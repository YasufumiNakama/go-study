# gofmt
gofmt -l -s -w .

# test
go env -w GO111MODULE=auto

# ex01
echo "===== run ex01 ====="
cd ./ex01/tmpconv
go test -run ""

# ex01
echo "===== run ex02 ====="
cd ./ex02/convs
go test -run ""
