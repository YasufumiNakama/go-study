# go-study
## Docker
### Start
```
docker-compose up -d --build
```
### in container
```
docker-compose exec app /bin/sh
go run main.go
```
### from host machine
```
docker-compose exec app go run main.go
```
### End
```
docker-compose down --rmi all --volumes
```