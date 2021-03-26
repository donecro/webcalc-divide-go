# webcalc-multiple-go

## Usage
* Accept requests on `0.0.0.0:8005` an write request headers and body to standard output:
```
go run ./server.go
```


### CI Test
* Unit Test -Router
```
go
```
* Unit Test -Calc
```
go test -v ./
```

## Docker
* Run
```
docker run -p 8004:8004 --rm webcalc-divide-go -hec
```

* Test
```
curl -d '{"x": "6", "y": "7"}' http://localhost:8004
```
