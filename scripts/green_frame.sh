#!/bin/bash

width=10

./reset_window.sh

curl -X POST http://localhost:17000 -d "green"
curl -X POST http://localhost:17000 -d "bgrect $width $width $((800-width)) $((800-width))"
curl -X POST http://localhost:17000 -d "update"
