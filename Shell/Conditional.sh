#!/bin/bash

# # Basic conditional.
day="Monday"

case $day in 
    "Monday")
        echo "Frontend"
        ;;
    "Friday")
        echo "Backend"
        ;;
    "Saturday" | "Sunday")
        echo "Weekend"
        ;;
    *)
        echo "Pull stack - stackoverflow"
        ;;
esac

# # Shell script to execute commands on pg-box
