## prip
A simpe go tool that prints the IP addresses in the range from 1 to 255

## Usage:

```
prip 127.0.0.1
```
![1](https://github.com/computerauditor/prip/assets/117805200/c16e339a-001b-4763-8843-7f14f2498cfd)

## Or can pipe using

```
echo '127.0.0.1' | prip  
```
![2](https://github.com/computerauditor/prip/assets/117805200/ca6ee422-404f-4a17-b5f2-8af3cd6e298d)

## Installation:

Simply use go build to build from binnary and set it up in your $GOPATH 

for example:

```
go build .\prip.go
```
Copy/Move to your defalt go path

```
mv prip.exe \go\bin
```
