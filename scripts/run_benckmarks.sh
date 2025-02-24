#!/bin/bash

PROXY_URL="http://localhost:8080"

echo "Running benchmarks on $PROXY_URL..."

if command -v wrk &> /dev/null; then
    echo "Running wrk benchmark..."
    wrk -t4 -c100 -d30s $PROXY_URL
else
    echo "wrk not found. Skipping wrk benchmark."
fi

if command -v hey &> /dev/null; then
    echo "Running hey benchmark..."
    hey -n 10000 -c 100 $PROXY_URL
else
    echo "hey not found. Skipping hey benchmark."
fi
