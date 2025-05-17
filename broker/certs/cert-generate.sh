#!/bin/bash

echo "Generating server cert..."

openssl genrsa -out ca.key 2048

openssl req -x509 -new -nodes -key ca.key -sha256 -days 3650 -out ca.crt -subj "/CN=client1"

openssl genrsa -out server.key 2048

openssl req -new -key server.key -out server.csr -config openssl-san.cnf

openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
-out server.crt -days 365 -sha256 -extfile openssl-san.cnf -extensions v3_req



echo "Generating client cert..."

openssl genrsa -out client.key 2048

openssl req -new -key client.key -out client.csr -config openssl-san.cnf

openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
-out client.crt -days 365 -sha256 -extfile openssl-san.cnf -extensions v3_req


