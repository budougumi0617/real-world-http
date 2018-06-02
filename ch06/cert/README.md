# How to generate each SSL files

# Generate private CA
```
$ openssl genrsa -out ca.key 2048
$ openssl req -new -sha256 -key ca.key -out cs.csr -config openssl.cnf
$ openssl x509 -in ca.csr -days 365 -req -signkey ca.key -sha256 -out ca.crt -extfile ./openssl.cnf -extensions CA
```

# Confirm files
```bash
# Confirm secret key
$ openssl rsa -in ca.key -text
# Confirm CSR
$ openssl req -in ca.csr -text
# Confirm certificate
$ openssl x509 -in ca.crt -text
```


# Generate server certificate
```
$ openssl genrsa -out server.key 2048
$ openssl req -new -nodes -sha256 -key server.key -out server.csr -config ./openssl.cnf
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [JP]:
State or Province Name (full name) [Tokyo]:
Locality Name (eg, city) [Kawasaki]:
Organization Name (eg, company) [example.com]:
Organizational Unit Name (eg, section) []:
Common Name (eg, fully qualified host name) []:localhost
Email Address [webmaster@example.com]:

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:
$ openssl x509 -req -days 365 -in server.csr -sha256 -out -server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions Server
```
