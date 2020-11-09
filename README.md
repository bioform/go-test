# Build
```bash
GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -X 'main.Version=v2.0.0' -X 'github.com/bioform/go-test/build.User=$(id -u -n)' -X 'github.com/bioform/go-test/build.Time=$(LANG=en_us_88591; date)'" .
```

```
-s
	Omit the symbol table and debug information.
-w
	Omit the DWARF symbol table.
```

# Get symbols from builded file 
*if -s option wasn't provided during build*
```bash
go tool nm ./go-test | grep go-test
```

# Run
```bash
 docker run -it --rm bioform/go-test:first
```

 # Display docker images
```bash
 docker ps -a --format "{{.ID}} {{.Command}}" --no-trunc
```

 # Remove all STOPPED containers
```bash
 docker rm $(docker ps -a -q)
```
