#!/bin/bash
echo "Running wrk benchmark..."
wrk -t12 -c400 -d30s http://localhost:8080