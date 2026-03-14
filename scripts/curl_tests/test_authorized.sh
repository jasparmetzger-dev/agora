#!/usr/bin/env bash

BASE_URL="http://localhost:8080"
USERNAME="testuser_$RANDOM"
EMAIL="$USERNAME@test.com"
PASSWORD="password123"

echo "=== Register $USERNAME ==="
curl -s -X POST "$BASE_URL/register" \
  -H "Content-Type: application/json" \
  -d "{\"username\": \"$USERNAME\", \"email\": \"$EMAIL\", \"password\": \"$PASSWORD\"}" > /dev/null

echo "=== Login ==="
TOKEN=$(curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"username\": \"$USERNAME\", \"password\": \"$PASSWORD\"}" | jq -r '.token')
echo "Token: $TOKEN"

AUTH="Authorization: Bearer $TOKEN"

echo "=== GET /profile ==="
curl -s -X GET "$BASE_URL/profile" -H "$AUTH" | jq

UPDATED_USERNAME="updated_$USERNAME"
UPDATED_EMAIL="$USERNAME@test.com"
echo "=== PATCH /profile ==="
curl -s -X PATCH "$BASE_URL/profile" \
  -H "$AUTH" -H "Content-Type: application/json" \
  -d "{\"username\": \"$UPDATED_USERNAME\", \"email\": \"$UPDATED_EMAIL\"}" | jq

echo "=== PATCH /profile/changepassword ==="
curl -s -X PATCH "$BASE_URL/profile/changepassword" \
  -H "$AUTH" -H "Content-Type: application/json" \
  -d "{\"old_password\": \"$PASSWORD\", \"new_password\": \"newpassword123\"}" | jq

echo "=== POST /posts ==="
POST_RESPONSE=$(curl -s -X POST "$BASE_URL/posts" \
  -H "$AUTH" -H "Content-Type: application/json" \
  -d '{"title": "Test Post", "content": "This is test post content"}')
echo "$POST_RESPONSE" | jq
POST_ID=$(echo "$POST_RESPONSE" | jq -r '.post.ID')

echo "=== GET /posts ==="
curl -s -X GET "$BASE_URL/posts" -H "$AUTH" | jq

echo "=== PATCH /posts/$POST_ID ==="
curl -s -X PATCH "$BASE_URL/posts/$POST_ID" \
  -H "$AUTH" -H "Content-Type: application/json" \
  -d '{"title": "Updated Post", "content": "Updated content"}' | jq

echo "=== GET /feed ==="
curl -s -X GET "$BASE_URL/feed" -H "$AUTH" | jq

echo "=== DELETE /posts/$POST_ID ==="
curl -s -X DELETE "$BASE_URL/posts/$POST_ID" -H "$AUTH" | jq 