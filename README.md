# Go Server

A "hello world" app to demonstrate simple web server. Contains a Dockerfile that is based off of instructions [here](https://medium.com/travis-on-docker/how-to-dockerize-your-go-golang-app-542af15c27a2).

Docker version can be run with the commands
```
$ docker pull andrebriggs/goserver
$ docker run --rm -p 8080:8080  andrebriggs/goserver
```

The way I built the docker image was using the following commands
```
$ docker run --rm -v "$PWD":/go/src/github.com/andrebriggs/goserver -w /go/src/github.com/andrebriggs/goserver iron/go:dev go build -o bin/myapp
$ docker build -t andrebriggs/goserver:latest .
```
Note that in the first line we __build__ "myapp" outside of the Dockerfile. The dependency of go at __iron/go:dev__ isn't a part of the final image so we save space.

FInally I pushed the image to DockerHub
```
$ docker push andrebriggs/goserver
```