# prip
A simpe go tool that prints the IP addresses in the range from 1 to 255

Usage:
```
prip 127.0.0.1
```
or can pipe using

```
echo '127.0.0.1' | prip  
```

Installation:

Simply use go build to build from binnary and set it up in your $GOPATH 

for example:

```
go build .\prip.go
```
Copy/Move to your defalt go path

```
mv prip.go \go\bin
```
