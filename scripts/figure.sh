#!/bin/bash

curl -X POST http://localhost:17000 -d "white"
curl -X POST http://localhost:17000 -d "bgrect 200 200 600 600"
curl -X POST http://localhost:17000 -d "figure 400 400"
curl -X POST http://localhost:17000 -d "green"
curl -X POST http://localhost:17000 -d "figure 480 480"
curl -X POST http://localhost:17000 -d "update"