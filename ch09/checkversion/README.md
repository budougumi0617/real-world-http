# 9.1 HTTP/2 Reuslt

```bash
$ GODEBUG=http2client=0 go run checkversion/main.go
Protocol Version: HTTP/1.1
$ GODEBUG=http2client=1 go run checkversion/main.go
Protocol Version: HTTP/2.0
$ GODEBUG=http2debug=1 go run checkversion/main.go
2018/07/17 13:52:58 http2: Transport failed to get client conn for google.com:443: http2: no cached connection was available
2018/07/17 13:52:59 http2: Transport creating client conn 0xc420158000 to 172.217.161.46:443
2018/07/17 13:52:59 http2: Transport encoding header ":authority" = "google.com"
2018/07/17 13:52:59 http2: Transport encoding header ":method" = "GET"
2018/07/17 13:52:59 http2: Transport encoding header ":path" = "/"
2018/07/17 13:52:59 http2: Transport encoding header ":scheme" = "https"
2018/07/17 13:52:59 http2: Transport encoding header "accept-encoding" = "gzip"
2018/07/17 13:52:59 http2: Transport received SETTINGS len=18, settings: MAX_CONCURRENT_STREAMS=100, INITIAL_WINDOW_SIZE=1048576, MAX_HEADER_LIST_SIZE=16384
2018/07/17 13:52:59 http2: Transport encoding header "user-agent" = "Go-http-client/2.0"
2018/07/17 13:52:59 http2: Transport received WINDOW_UPDATE len=4 (conn) incr=983041
2018/07/17 13:52:59 http2: Transport received SETTINGS flags=ACK len=0
2018/07/17 13:52:59 http2: Transport received HEADERS flags=END_HEADERS stream=1 len=218
2018/07/17 13:52:59 http2: Transport received DATA stream=1 len=220 data="<HTML><HEAD><meta http-equiv=\"content-type\" content=\"text/html;charset=utf-8\">\n<TITLE>301 Moved</TITLE></HEAD><BODY>\n<H1>301 Moved</H1>\nThe document has moved\n<A HREF=\"https://www.google.com/\">here</A>.\r\n</BODY></HTML>\r\n"
2018/07/17 13:52:59 http2: Transport received DATA flags=END_STREAM stream=1 len=0 data=""
2018/07/17 13:52:59 http2: Transport received PING len=8 ping="\x00\x00\x00\x00\x00\x00\x00\x00"
2018/07/17 13:52:59 http2: Transport failed to get client conn for www.google.com:443: http2: no cached connection was available
2018/07/17 13:52:59 http2: Transport creating client conn 0xc4203021c0 to 172.217.161.36:443
2018/07/17 13:52:59 http2: Transport encoding header ":authority" = "www.google.com"
2018/07/17 13:52:59 http2: Transport encoding header ":method" = "GET"
2018/07/17 13:52:59 http2: Transport encoding header ":path" = "/"
2018/07/17 13:52:59 http2: Transport encoding header ":scheme" = "https"
2018/07/17 13:52:59 http2: Transport encoding header "referer" = "https://google.com"
2018/07/17 13:52:59 http2: Transport encoding header "accept-encoding" = "gzip"
2018/07/17 13:52:59 http2: Transport encoding header "user-agent" = "Go-http-client/2.0"
2018/07/17 13:52:59 http2: Transport received SETTINGS len=18, settings: MAX_CONCURRENT_STREAMS=100, INITIAL_WINDOW_SIZE=1048576, MAX_HEADER_LIST_SIZE=16384
2018/07/17 13:52:59 http2: Transport received WINDOW_UPDATE len=4 (conn) incr=983041
2018/07/17 13:52:59 http2: Transport received SETTINGS flags=ACK len=0
2018/07/17 13:52:59 http2: Transport received HEADERS flags=END_HEADERS stream=1 len=480
Protocol Version: HTTP/2.0
2018/07/17 13:52:59 http2: Transport received DATA flags=PADDED stream=1 len=5509 data="\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xcd[ys\xdbF\xb2\xff?\x9f\x02\x82\xca\x14\xf9\x04\x918I\x904\xe4\xb5e\xc5Q6\xc9:\x9bd\xb3\x89\xe3U\r\x80!\t\t\x97\x00PG(~\xf7\xf7\xeb\xc1A\x90\xa2\xecT\xde\xd6\xdb-J\xc4\x1c\xdd=}MO\xcf\f\xf8\xf2\xc0O\xbc\xe2!\xe5Ң\x88\xc2ӗ\xf4-\x05\x05\x8fr/I\xb9#ˢB\x00\x8e\xbc(\x8at2\x18\xe4ނG\xac\x9fd\xf3\xc1\xcf\xdc}\xcf\xe6\\\x96B\x16\xcf\x1d\xf9\x8aɠ\xc0\x99\u007f\xfa2\xe2\x05\x93\xbc$.x\\8r\xe7P\x1b\x8f\xc7\xea\xb4sh\xa8\xaan㩫\xaaf\xe0\xa9\xe9\xc6\xd8.\x9f\x96)\x9e\xa6nUO\xbdz\x8e\b\xde\x1c\x8d\b^\u05edѨl7\b^\x1f\x8eU\xc27t\xd5(\xf1\x8d\xba_/\x9f\x86]\x8ecj\xa3\xad\xf1L\xbb\x84\xb7L\xb5|\xeaF\x9b\xaeeY\xba\xe0\xd3\x18\x8d+\xbaV\x85\xafV\xfc\x0e+\xba\xdb\xe3\xeau\xff" (5026 bytes omitted)
```

