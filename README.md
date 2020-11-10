# Meeting
This is a simple api for scheduling meeing.

# Test

To run the tests run the following command from `service` directory

```
$ go test -v
```

# Build
To build the project run the following command from the `bin` directory.

```
$ go build -o meeting main.go
```

# Run
To run the project run the `main.go` file in `bin` directory.
```
$ go run main.go
```
You can also run the binary generated in the build step

```
$ ./meeting
```

# Configuration

Configuration can be changed from `config.json` file in the `bin` directory