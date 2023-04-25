#!/bin/bash

screen_size=800
figure_radius=100
start_pos=$figure_radius
finish_pos=$((screen_size - figure_radius))
max_dist=$((finish_pos - start_pos))
step=10
interval=0.01

curl -X POST http://localhost:17000 -d "figure $start_pos $start_pos"
pos=$start_pos
dist=0
sleep $interval

while true; do
    while (( pos < finish_pos )); do
    if ((dist + d > max_dist)); then
        curl -X POST http://localhost:17000 -d "move $((max_dist-dist)) $((max_dist-dist))"
        pos=$finish_pos
        dist=$max_dist
    else
        curl -X POST http://localhost:17000 -d "move $step $step"
        pos=$((pos + step))
        dist=$((dist + step))
    fi
    sleep $interval
    done

    while (( pos > start_pos )); do
        if ((dist - step < 0)); then
            curl -X POST http://localhost:17000 -d "move $((-dist)) $((-dist))"
            pos=$start_pos
            dist=0
        else
        curl -X POST http://localhost:17000 -d "move $((-step)) $((-step))"
        pos=$((pos - step))
        dist=$((dist - step))
        fi
        sleep $interval
    done
done
