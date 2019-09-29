# lets-go
Playing with GO

```
docker build -t lets-go .
```

```
docker run -it --rm --cpus=".5" --memory="256m" -p 8182:8182 lets-go
```

```
curl localhost:8182/api/user/1  #1 or 2
curl -X POST localhost:8182/api/log/2000
curl -X POST localhost:8182/api/log/2000000 #crash
```
