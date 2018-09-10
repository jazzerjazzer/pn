# PubNative

## How to build: 
- Clone this repository:
```
git clone git@github.com:jazzerjazzer/mb.git
```
- Install direnv: 
```
brew install direnv
```
- Change your working directory to `pn`
- Allow direnv: 
``` 
direnv allow
```
- Run the server (It will be running on port 1321): 
```
go run src/pn/main.go
```

## How to upload the CSV file: 
curl -X POST -F file=@ids-test.csv  http://localhost:1321/upload

## How to search an ID: 
curl http://localhost:1321/promotions/d018ef0b-dbd9-48f1-ac1a-eb4d90e57118 -i

## E2E tests: 
```
go test src/pn/integration_test.go
```

## Benchmarks (Mid 2012 Macbook Pro, 4GB RAM, Core i5 (I5-3210M))
~200.000 lines: 60ms
~1.000.000 lines: 270ms
~3.000.000 lines: 1.91s

## Known Issues: 
- When the server is shutdown, CSV file needs to be uploaded again.
- No unit tests are included.