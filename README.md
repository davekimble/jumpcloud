# Password Hash Server

This utility starts an HTTP server that listens on the specified port (defaults to 8080).

When it receives a POST with a password, it pauses the specified number of seconds (defaults to 5), then returns the SHA512 hash of hat password encoded in base64.

## Usage

To use default values described above, run it like this:
```shell
> go run test
```
To modify either the port or pause length, run with the flags "port" and/or "pause".

For example, to run such that the server listens on port **8088**, run it like this:
```shell
> go run test -port=8088
```
To run with a pause of **3** seconds, run it like this:
```shell
> go run test -pause=3
```

Also, if you'd like to show debug information as the server is running, set the debug flag to true like this:
```shell
> go run test -debug=true
```



Once the server is running, send a password to it to get hashed by running this from another shell:
```shell
> curl http://localhost:8088?password=angryMonkey
```

or this:
```shell
> curl --data "password=angryMonkey" http://localhost:8080
```

