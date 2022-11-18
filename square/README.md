# How to use this module?
To test this module, run the following command in a server machine.

To start the server

go run server/main.go

To start a client to connect with server

go run client/main.go <Optional: IP Address> <Optional: Port Number>

Here IP address and/or Port number are optional. If port number is not provided, program will assume 50051 as the port number. And if IP address is not provided, then the program will assume localhost

Test details:

On the server:
```shell
root:~/bits-wilp-s1/square# go run server/main.go 
2022/11/18 03:59:37 server listening at [::]:50051
2022/11/18 04:00:51 Received: 11
2022/11/18 04:01:00 Received: 123456789123456789
```
On the client:
```shell
root:~/bits-wilp-s1/square# go run client/main.go 172.44.0.19 50051
Enter a number to be squared
->11
2022/11/18 04:00:51 The square of 11 is: 121
root:~/bits-wilp-s1/square# go run client/main.go 172.44.0.19 50051
Enter a number to be squared
->123456789123456789
2022/11/18 04:01:00 The square of 123456789123456789 is: 15241578780673678515622620750190521
root:~/bits-wilp-s1/square# 
```
