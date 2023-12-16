## prip
A simpe go tool that prints the IP addresses in the range from 1 to 255

## Usage:

```
prip 127.0.0.1
```
![image 1](https://github.com/computerauditor/prip/assets/117805200/28b34952-05bc-41fb-b2b2-b0646107f274)

## Or can pipe using

```
echo '127.0.0.1' | prip  
```
![image 2](https://github.com/computerauditor/prip/assets/117805200/7cf54123-d65b-4592-aebd-263359515276)

## Installation:

Simply use go build to build from binnary and set it up in your $GOPATH 

for example:

```
go build .\prip.go
```
Copy/Move to your defalt go path

```
mv prip.go \go\bin
```
