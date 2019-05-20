# goRestApi
A simple RestApi Client built with go

### Dependecies
- packr2 (https://github.com/gobuffalo/packr/tree/master/v2)
Has added to path to be able to use from terminal `export PATH="$PATH":"$GOBIN"`
where `GOBIN` is the path to go binaries folder

### run

```
# running localy 
packr2 run server.go

# build
packr2 build server.go

# run the compiled application:
./server`

# Compile to ARM7
env GOOS=linux GOARCH=arm GOARM=7 packr2 build
```
