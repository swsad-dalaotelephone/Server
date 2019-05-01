# baobaozhuan Server
Go Server for baobaozhuan, use gin + gorm development framework 

## How to start
#### Reference
- gin : https://github.com/gin-gonic/gin
- gorm : https://github.com/jinzhu/gorm
- weapp sdk : https://github.com/medivhzhan/weapp 

#### Automatically download dependent packages by go module (go >= 1.11)
1. open go modules
    ```
    $ export GO111MODULE=on
    ```
2. download module
    ```
    $ go mod tidy -v
    ```

#### configure settings
modify config.json

#### run
```
$ go build
$ ./baobaozhuan
```

## Project Structure
```
.
├── config          
├── config.json         configuration
├── controllers         controller
│   └── user
├── database            database connection
├── main.go             main program
├── middlewares         middleware for gin router
│   ├── auth
│   ├── cache
│   └── session
├── models              models
│   └── user
├── modules             module of project
│   ├── log
│   └── util
├── router              router
│   └── router.go
├── storage             storage
│   └── logs
│   └── upload
```

