Golang DB Connection Pool Playground
===
## Config
Configuration using Environment Variables
```
DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=test_db
DB_USERNAME=root
DB_PASSWORD=
DB_CONNECTION_MAX=100
```

## Docker
```bash
$ docker pull yohang/golang-db
$ docker run -it -p 8080:8080 yohang/golang-db
```