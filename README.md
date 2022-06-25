# Secure Golang connection

## Context

### Generating Certficate

#### Print dry-run certificate-request

```bash
openssl req -new \
    -newkey rsa:2048 \
    -nodes \
    -x509 \
    -subj "/C=US/ST=GD/L=SZ/O=Startup, Inc./CN=domain.com/emailAddress=info@dmoain.com" \
    -addext "subjectAltName = DNS:domain.com" \
    -text \
    -noout
```

#### Generate

CSR first

```bash
openssl req -new \
    -newkey rsa:2048 \
    -nodes \
    -subj "/C=US/ST=GD/L=SZ/O=Startup, Inc./CN=domain.com/emailAddress=info@dmoain.com" \
    -addext "subjectAltName = DNS:domain.com, DNS:*.domain.com" \
    -keyout domain.key \
    -out domain.csr
```
.
```bash
openssl x509 -req \
    -in domain.csr \
    -days 365 \
    -signkey domain.key \
    -out domain.crt
```

### Generate certificate for `domain.com` signed with created CA

Since this certificate is not signed by a trusted CA, we need to install it on our system and tweak its trust parameters as described in the same article.

```bash
openssl x509 -req -CAcreateserial -days 365 \
    -CA domain.crt \
    -CAkey domain.key \
    -out domain.crt \
    -in domain.csr 
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

