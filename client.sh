#!/usr/bin/env bash

# Target url is the host of the html you want to render
# passed as the first argument
target_url=$1

# If the target is empty we default to example.com
[ -z "${1}" ] && target_url="http://www.example.com"

# Request that the service retrieve the page, render it and send the result in the response
curl 'http://localhost:8080/print' \
       -H 'Content-Type: application/x-www-form-urlencoded' \
       --data-urlencode "url=${target_url}" \
       --output 'example.pdf' \
       -vv

