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

# ex03
echo "===== run ex03 ====="
cd ./ex03/popcount
go mod init popcount
go test -bench=. 
cd ../../

# ex04
echo "===== run ex04 ====="
cd ./ex04/popcount
go mod init popcount
go test -bench=. 
cd ../../
