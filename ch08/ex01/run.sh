go build -o clock2 clock.go
TZ=US/Eastern    ./clock2 -port 8010 &
TZ=Asia/Tokyo    ./clock2 -port 8020 &
TZ=Europe/London ./clock2 -port 8030 &
go run clockwall.go NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030