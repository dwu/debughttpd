## debughttpd

Simple logging HTTP server for debugging purposes

```shell
Usage of ./debughttpd:
  -a string
    	server address, format: '<hostname|ip-address>:<port>' (default ":8080")
  -c string
    	content to return to the client; given as string value on the command line
  -cf string
    	filename of the content to return to the client
  -ct string
    	content type of the response (default "text/plain")
```

## Usage example

```shell
$ ./debughttpd -a localhost:8888 -c '{"foo":"foobar"}' -ct "application/json"
>> Listening on localhost:8888
>> Serving content: {"foo":"bar"}, Content-Type: application/json
>> Request URL: /test | Timestamp: 2023-09-03 16:19:18 | Method: POST
POST /test HTTP/1.1
Host: localhost:8888
Accept: */*
Content-Length: 18
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/8.2.1

message=helloworld
```
