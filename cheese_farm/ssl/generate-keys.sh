#!/bin/bash

openssl req -x509 -newkey rsa:4096 -keyout private.key -out cert.pem -days 36500 -nodes -subj "/CN=localhost"
