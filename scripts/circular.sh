#!/bin/bash

start_x=400
start_y=100
top_to_left=300
left_to_bottom=600
bottom_to_right=900
right_to_top=1200
step=10

curl -X POST http://localhost:17000 -d "figure $start_x $start_y"

while true; do
    for ((i = 0; i < top_to_left; i += step)); do
        curl -X POST http://localhost:17000 -d "move $((-step)) $((step))"
    done

    for ((i = top_to_left; i < left_to_bottom; i += step)); do
        curl -X POST http://localhost:17000 -d "move $step $step"
    done

    for ((i = left_to_bottom; i < bottom_to_right; i += step)); do
        curl -X POST http://localhost:17000 -d "move $step $((-step))"
    done

    for ((i = bottom_to_right; i < right_to_top; i += step)); do
        curl -X POST http://localhost:17000 -d "move $((-step)) $((-step))"
    done
done