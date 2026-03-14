#!/usr/bin/env bash

BASE_URL="http://localhost:8080"


echo "=== Login testuser ==="
TOKEN=$(curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123"}' | jq -r '.token')

echo "Token: $TOKEN"

echo "=== Get Profile ==="
curl -s -X GET "$BASE_URL/profile" \
  -H "Authorization: Bearer $TOKEN" | jq