#!/bin/bash

#available colors: white, green
color="white"

curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "$color"
curl -X POST http://localhost:17000 -d "update"