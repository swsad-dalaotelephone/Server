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
#### Start with docker
[official preference](https://yeasy.gitbooks.io/docker_practice/content/install/centos.html)

install docker
```
sudo yum install docker
```

test 
```
[root@VM_0_9_centos nginx]# docker --version
Docker version 1.13.1, build b2f74b2/1.13.1
[root@VM_0_9_centos nginx]# docker-compose --version
docker-compose version 1.18.0, build 8dd22a9

[root@VM_0_9_centos Server]# docker run hello-world
Unable to find image 'hello-world:latest' locally
Trying to pull repository docker.io/library/hello-world ... 
latest: Pulling from docker.io/library/hello-world
1b930d010525: Pull complete 
Digest: sha256:41a65640635299bab090f783209c1e3a3f11934cf7756b09cb2f1e02147c6ed8
Status: Downloaded newer image for docker.io/hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.
```

pull image as follows:
```
[root@VM_0_9_centos Server]# docker image ls
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
docker.io/mysql     5.7.25              98455b9624a9        3 months ago        372 MB
docker.io/golang    1.11.5              1454e2b3d01f        3 months ago        816 MB
docker.io/redis     5.0.2               5958914cc558        6 months ago        94.9 MB
docker.io/nginx     1.12.2              4037a5562b03        14 months ago       108 MB

$ docker pull xxx:version
```

configure settings: modify config.json 

create Dockfile as [file](./Dockerfile)

create docker-compose.yml as [file](./docker-compose.yml)

modify /etc/nginx/nginx.conf as [file](https://github.com/swsad-dalaotelephone/Nginx-conf/blob/master/nginx.conf)

finally run it
```
$ docker-compose up
```
## Project Structure
```
.
├── config              # configuration
├── config.json           
├── controllers           # controller
│   ├── ad
│   ├── resource
│   ├── task
│   └── user
├── database            # database connection
├── main.go             # main program
├── middlewares         # middleware for gin router
│   ├── auth
│   ├── cache
│   ├── logger
│   └── session
├── models              # models
│   ├── ad
│   ├── campus
│   ├── common
│   ├── school
│   ├── tag
│   ├── task
│   └── user
├── modules             # module of project
│   ├── gredis
│   ├── log
│   └── util
├── router              # router
└── storage             # storage
    ├── file
    ├── img
    └── logs
```

