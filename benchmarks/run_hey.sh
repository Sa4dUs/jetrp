#!/bin/bash
echo "Ejecutando pruebas con hey..."
hey -z 30s -c 100 http://localhost:8080
