# gofmt
gofmt -l -s -w .

# test
go env -w GO111MODULE=auto

# ex01
echo "===== run ex01 ====="
cd ./ex01
go run echo.go one two three
cd ../

# ex02
echo "===== run ex02 ====="
cd ./ex02
go run echo.go one two three
go test -run ""
cd ../

# ex03
echo "===== run ex03 ====="
cd ./ex03
go run echo.go one two three
cd ../

# ex04
echo "===== run ex04 ====="
cd ./ex04
go run dup2.go a.txt b.txt c.txt
cd ../

# ex05
echo "===== run ex05 ====="
cd ./ex05
go run lissajous.go > out.gif
cd ../
