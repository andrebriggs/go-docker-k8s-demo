# GoLang App --> Docker --> k8s

A "hello world" app to demonstrate simple web server. Contains a Dockerfile that is based off of instructions [here](https://medium.com/travis-on-docker/how-to-dockerize-your-go-golang-app-542af15c27a2). 

We also will push the dockerized GoLang app to a Kubernetes cluster using an included yaml file.
## Docker Pull and Run
Docker version can be run with the commands
```
$ docker pull andrebriggs/goserver
$ docker run --rm -p 8080:8080  andrebriggs/goserver
```

Test that the app is running locally by running
```
$ curl http://localhost:8080
```

You should see 
```
Congratulations! Version 1.0 of your application is running on Kubernetes.
```

## Docker Build and Push
If you want to roll your own image use the following commands below, otherwise skip to the next section. 
```
$ docker run --rm -v "$PWD":/go/src/github.com/andrebriggs/goserver -w /go/src/github.com/andrebriggs/goserver iron/go:dev go build -o bin/myapp
$ docker build -t andrebriggs/goserver:latest .
```
Note that in the first line we __build__ "myapp" outside of the Dockerfile. The dependency of go at __iron/go:dev__ isn't a part of the final image so we save space. Be sure to replace __andrebriggs__ with your own DockerHub account name

FInally I pushed the image to DockerHub
```
$ docker push andrebriggs/goserver
```
## Kubernetes Setup

Once pushed I can declaratively set a Kubernetes environment (assuming one is created) to use this resource
```
$ kubectl create -f k8s
```
To verify the creation with can ask kubectl to describe the service with the command
```
$ kubectl describe svc mywebapp
```
The result will be similar to this
```
Name:                     mywebapp
Namespace:                default
Labels:                   app=mywebapp
Annotations:              kubectl.kubernetes.io/last-applied-configuration:
                            {"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"labels":{"app":"mywebapp"},"name":"mywebapp","namespace":"default"},"spe...
Selector:                 app=mywebapp
Type:                     LoadBalancer
IP:                       10.0.210.108
LoadBalancer Ingress:     137.116.69.143
Port:                     http  8080/TCP
TargetPort:               8080/TCP
NodePort:                 http  31855/TCP
Endpoints:                10.200.0.69:8080
Session Affinity:         None
External Traffic Policy:  Cluster
Events:
  Type    Reason                Age    From                Message
  ----    ------                ----   ----                -------
  Normal  EnsuringLoadBalancer  6m2s   service-controller  Ensuring load balancer
  Normal  EnsuredLoadBalancer   5m14s  service-controller  Ensured load balancer
```
Make a note of the external IP address (LoadBalancer Ingress) exposed by the service. Read more about this [here](https://kubernetes.io/docs/tutorials/stateless-application/expose-external-ip-address/).

You should be able to access the web app 

```
$ curl http://(LoadBalancer Ingress IP):8080
Congratulations! Version 1.0 of your application is running on Kubernetes.
```

## Update the Docker image 

Suppose you make changes to [server.go](server.go) file to update hardcoded version to __1.2__ and want to push new changes the Docker image tag as a __v1.2__. Do the following commands
```
$ docker run --rm -v "$PWD":/go/src/github.com/andrebriggs/goserver -w /go/src/github.com/andrebriggs/goserver iron/go:dev go build -o bin/myapp
$ docker build -t andrebriggs/goserver:v1.2 .
$ docker tag andrebriggs/goserver:latest andrebriggs/goserver:v1.2
$ docker push andrebriggs/goserver:v1.2
```

## Rollout a new deployment
Edit the manifest [YAML file](k8s/mywebapp-all-in-one.yaml) to alter the .spec.containers.image to be andrebriggs/goserver:v1.2
```
    ...
    spec:
      containers:
      - name: mywebapp
        image: andrebriggs/goserver:v1.2
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 8080
```

Also edit the .spec.replicas value to be 10

Save then run

```
$ kubectl apply -f k8s --record
```

Then immediately run the following to check the status of the deploy
```
kubectl rollout status deployment mywebapp-v1
```
Once done try hitting the endpoint again with curl and see your changes.
```
Congratulations! Version 1.2 of your application is running on Kubernetes.
```

## Kubernetes tear down
Run the commands
```
$ kubectl delete services mywebapp
$ kubectl delete deployment mywebapp
```