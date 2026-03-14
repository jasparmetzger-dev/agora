#!/usr/bin/env bash

BASE_URL="http://localhost:8080"

echo "=== Register ==="
echo "=== Register Normally ==="
curl -s -X POST "$BASE_URL/register" \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "email": "test@test.com", "password": "password123"}' | jq

echo "=== Register duplicate ==="
curl -s -X POST "$BASE_URL/register" \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "email": "test@test.com", "password": "password123"}' | jq

echo "=== Register invalid email ==="
curl -s -X POST "$BASE_URL/register" \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser2", "email": "notanemail", "password": "password123"}' | jq

echo "=== Register short password ==="
curl -s -X POST "$BASE_URL/register" \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser2", "email": "test2@test.com", "password": "123"}' | jq

echo "=== Login ==="
TOKEN=$(curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123"}' | tee /dev/stderr | jq -r '.token')

echo "=== Login wrong password ==="
echo "=== Login Normally ==="
curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "wrongpassword"}' | jq

echo "=== Login unknown user ==="
curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d '{"username": "nobody", "password": "password123"}' | jq

echo "Token: $TOKEN"