# Clinet output logs
```bash
$ gor main.go
2018/06/06 23:40:58 Status: 101 Switching Protocols
2018/06/06 23:40:58 Headers: map[Connection:[Upgrade] Upgrade:[MyProtocol]]
<- 1
-> 10
<- 2
-> 9
<- 3
-> 8
<- 4
-> 7
<- 5
-> 6
<- 6
-> 5
<- 7
-> 4
<- 8
-> 3
<- 9
-> 2
<- 10
-> 1
```

# Server output logs
```bash
$ gor main.go
2018/06/06 23:40:46 start http listening :18888
Upgrade to MyProtocol
-> 1
<- 10
-> 2
<- 9
-> 3
<- 8
-> 4
<- 7
-> 5
<- 6
-> 6
<- 5
-> 7
<- 4
-> 8
<- 3
-> 9
<- 2
-> 10
<- 1
```
