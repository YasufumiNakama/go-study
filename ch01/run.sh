# gofmt
gofmt -l -s -w .

# test
go env -w GO111MODULE=auto

# ex01
cd ./ex01
go run echo.go one two three
cd ../

# ex02
cd ./ex02
go run echo.go one two three
go test -run ""
cd ../

# ex03
cd ./ex03
go run echo.go one two three

