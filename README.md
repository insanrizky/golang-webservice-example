# Web Service Example using GoLang
Create own web server and mini service with [GoLang](http://golang.org)

## Usage
## 1. Import Database
Open `db` directory then `import` `golang.sql` into your database.

## 2. Install Dependencies
### github.com/go-sql-driver/mysql
it is used for connecting go with mysql database (driver)
```sh
$ go get github.com/go-sql-driver/mysql
```

### golang.org/x/crypto/bcrypt
it is used for encrypt password using bcrypt algorithm
```sh
$ go get golang.org/x/crypto/bcrypt
```

## 2. Install Project
```sh
$ go install github.com/yourusername/golang-webservice-example
```

## 3. Run Project
```sh
$ $GOPATH/bin/golang-webservice-example
```

Hope it useful!
Thanks

