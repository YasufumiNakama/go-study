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

# ex06
echo "===== run ex06 ====="
cd ./ex06
go run lissajous.go > out.gif
cd ../

# ex07
echo "===== run ex07 ====="
cd ./ex07
go run fetch.go http://www.gopl.io/
go run fetch.go http://bad.gopl.io/
cd ../

# ex08
echo "===== run ex08 ====="
cd ./ex08
go run fetch.go www.gopl.io/
go run fetch.go bad.gopl.io/
cd ../

# ex09
echo "===== run ex09 ====="
cd ./ex09
go run fetch.go http://www.gopl.io/
go run fetch.go http://bad.gopl.io/
cd ../

# ex10
echo "===== run ex10 ====="
cd ./ex10
rm *.txt
go run fetchall.go https://en.wikipedia.org/wiki/Baseball https://ja.wikipedia.org/wiki/野球
go run fetchall.go https://en.wikipedia.org/wiki/Baseball https://ja.wikipedia.org/wiki/野球
cd ../
