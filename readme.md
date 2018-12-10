## Generate Certificate for HTTPS
```sh
openssl req -x509 -nodes -days 365 -sha256 -newkey rsa:2048 -keyout keys/key.pem -out keys/cert.pem
```