# Golang APIRest - Hexagonal Arquitecture

## To use MySql database instead data mock in memory using docker

```
docker compose up -d
```
To access database using mysql client

```sh
# mysql -h <your-localhost-IP> -P 3306 -u root -p
mysql -h 192.168.18.46 -P 3306 -u root -p
```
pass: secret

## Run server


In src folder
```
go run cmd/main.go
```

then visit [http://localhost:3000](http://localhost:3000) in your browser


