#!/usr/bin/env bash
curl 'http://localhost:8080/print' \
       -H 'Content-Type: application/x-www-form-urlencoded' \
       --data-urlencode "url=https://www.example.com" \
       --output 'example.pdf' \
       -vv

