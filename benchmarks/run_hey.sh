#!/bin/bash
echo "Running wrk benchmark..."
hey -z 30s -c 100 http://localhost:8080
