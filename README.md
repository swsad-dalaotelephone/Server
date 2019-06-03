# baobaozhuan Server
Go Server for baobaozhuan, use gin + gorm development framework 

## How to start
#### Reference
- gin : https://github.com/gin-gonic/gin
- gorm : https://github.com/jinzhu/gorm
- weapp sdk : https://github.com/medivhzhan/weapp 

#### Required 
- mysql 5.7
- redis 5.0.2

#### Automatically download dependent packages by go module (go >= 1.11)
    open go modules
    ```
    $ export GO111MODULE=on
    ```
#### Quick Start
- go get
    ```
    $ go get -u github.com/swsad-dalaotelephone/Server
    ```
- configure settings: modify config.json 
- run
    ```
    $ go install
    $ Server
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

