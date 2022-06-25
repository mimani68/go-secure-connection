# Secure Golang connection

## Context

### Generating Certficate

#### Print dry-run certificate-request

```bash
openssl req -new -newkey rsa:2048 -nodes -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=domain.com" -addext "subjectAltName = DNS:domain.com" -text -noout
```

#### Generate

```bash
openssl req -new -newkey rsa:2048 -nodes -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=domain.com" -addext "subjectAltName = DNS:domain.com" -keyout domain.key -out domain.csr
openssl x509 -req -in domain.csr -days 365 -signkey domain.key -out domain.crt
```

### Generate certificate for `domain.com` signed with created CA

Since this certificate is not signed by a trusted CA, we need to install it on our system and tweak its trust parameters as described in the same article.

```bash
openssl x509 -req -in domain.csr -days 365 -CA domain.crt -CAkey domain.key -CAcreateserial -out domain.crt
```

## Development 

```bash
go run server.go
go run client.go
```

## cURL test

```bash
curl --cacert ./certs/domain.crt https://domain.com:9000
```

