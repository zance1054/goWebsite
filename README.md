# goWebsite

You need to open docker desktop and sign in so that you can push the image

build command

```
docker build --no-cache -t zance/fielding.cloud:latest .
```

push image

docker push zance/fielding.cloud:latest

SSh to server, docker pull, then deploy

deploy command

``` 
docker run -d --name my-go-app --network proxynet -p 3021:8080 e628c2d9ff8a
```