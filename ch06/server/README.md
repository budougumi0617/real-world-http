
```
$ curl https://localhost:18443
curl: (60) SSL certificate problem: unable to get local issuer certificate
More details here: https://curl.haxx.se/docs/sslcerts.html

$ curl --cacert ca.crt https://localhost:18443 -v
* Rebuilt URL to: https://localhost:18443/
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 18443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* Cipher selection: ALL:!EXPORT:!EXPORT40:!EXPORT56:!aNULL:!LOW:!RC4:@STRENGTH
* successfully set certificate verify locations:
*   CAfile: ca.crt
  CApath: none
* TLSv1.2 (OUT), TLS handshake, Client hello (1):
* TLSv1.2 (IN), TLS handshake, Server hello (2):
* TLSv1.2 (IN), TLS handshake, Certificate (11):
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
* TLSv1.2 (IN), TLS handshake, Server finished (14):
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
* TLSv1.2 (OUT), TLS change cipher, Client hello (1):
* TLSv1.2 (OUT), TLS handshake, Finished (20):
* TLSv1.2 (IN), TLS change cipher, Client hello (1):
* TLSv1.2 (IN), TLS handshake, Finished (20):
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256
* ALPN, server accepted to use h2
* Server certificate:
*  subject: C=JP; ST=Tokyo; L=Kawasaki; O=example.com; CN=localhost; emailAddress=webmaster@example.com
*  start date: Jun  3 04:58:11 2018 GMT
*  expire date: Jun  3 04:58:11 2019 GMT
*  common name: localhost (matched)
*  issuer: C=JP; ST=Tokyo; L=Kawasaki; O=example.com; emailAddress=webmaster@example.com
*  SSL certificate verify ok.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x7fcdb500aa00)
> GET / HTTP/2
> Host: localhost:18443
> User-Agent: curl/7.54.0
> Accept: */*
>
* Connection state changed (MAX_CONCURRENT_STREAMS updated)!
< HTTP/2 200
< content-type: text/html; charset=utf-8
< content-length: 31
< date: Sun, 03 Jun 2018 05:30:31 GMT
<
* Connection #0 to host localhost left intact
<html><body>hello</body></html>%
```
