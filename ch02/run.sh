# gofmt
gofmt -l -s -w .

# test
go env -w GO111MODULE=auto

# ex01
echo "===== run ex01 ====="
cd ./ex01/tmpconv
go test -run ""
cd ../../

# ex02
echo "===== run ex02 ====="
cd ./ex02/convs
go test -run ""
cd ../
go mod init ex02
go run main.go -t -l -w 100
go run main.go # 標準入力
cd ../
