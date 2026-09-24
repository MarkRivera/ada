#!/usr/bin/env bash

curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"title":"Hello","description":"Test123","fileType":"video/mp4"}' \
  http://localhost:8080/upload