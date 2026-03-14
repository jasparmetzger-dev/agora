#!/usr/bin/env bash

BASE_URL="http://localhost:8080"

TOKEN=$(curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123"}' | jq -r '.token')

AUTH="Authorization: Bearer $TOKEN"
echo "Token: $TOKEN"

echo "=== GET /profile ==="
echo $(curl -s -X GET "$BASE_URL/profile" \
  -H "$AUTH" )

echo "=== PATCH /profile ==="
echo $(curl -s -X PATCH "$BASE_URL/profile" \
  -H "$AUTH" \
  -H "Content-Type: application/json" \
  -d '{"username": "updateduser", "email": "updated@test.com"}' )

echo "=== PATCH /profile/changepassword ==="
echo $(curl -s -X PATCH "$BASE_URL/profile/changepassword" \
  -H "$AUTH" \
  -H "Content-Type: application/json" \
  -d '{"old_password": "password123", "new_password": "newpassword123"}' )

echo "=== POST /posts ==="
POST_RESPONSE=$(curl -s -X POST "$BASE_URL/posts" \
  -H "$AUTH" \
  -H "Content-Type: application/json" \
  -d '{"title": "Test Post", "content": "This is a test post content"}')

echo "$POST_RESPONSE" | jq

# Extract post ID, assuming the response has an "id" field
POST_ID=$(echo "$POST_RESPONSE" | jq -r '.id')

echo "=== GET /posts ==="
echo $(curl -s -X GET "$BASE_URL/posts" \
  -H "$AUTH" | jq)

echo "=== PATCH /posts/$POST_ID ==="
echo $(curl -s -X PATCH "$BASE_URL/posts/$POST_ID" \
  -H "$AUTH" \
  -H "Content-Type: application/json" \
  -d '{"title": "Updated Test Post", "content": "Updated content"}' | jq)

echo "=== GET /feed ==="
echo $(curl -s -X GET "$BASE_URL/feed" \
  -H "$AUTH" | jq)

echo "=== DELETE /posts/$POST_ID ==="
echo $(curl -s -X DELETE "$BASE_URL/posts/$POST_ID" \
  -H "$AUTH" | jq)