# 3-7 POST data

Like a `curl -d test=vaue http://localhost:18888`


```bash
# server log
POST / HTTP/1.1
Host: localhost:18888
Accept-Encoding: gzip
Content-Length: 10
Content-Type: application/x-www-form-urlencoded
User-Agent: Go-http-client/1.1

test=value
Content-Type [application/x-www-form-urlencoded]
Accept-Encoding [gzip]
User-Agent [Go-http-client/1.1]
Content-Length [10]
HEAD elements= map[Content-Type:[application/x-www-form-urlencoded] Accept-Encoding:[gzip] User-Agent:[Go-http-client/1.1] Content-Length:[10]]
BODY = test=value
```
