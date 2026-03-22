#!/usr/bin/env bash

BASE_URL="http://localhost:8080"
BASE_UPLOAD_URL="/home/jaspa/linux_repos/agora/testing_data/videos"
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

echo "=== Upload Videos"
curl -v -X POST "$BASE_URL/posts" \
    -H "Authorization: Bearer $TOKEN" \
    -F "title=one" \
    -F "description=desription one" \
    -F "file=@$BASE_UPLOAD_URL/test_video_1.mp4"
echo "=== 1 / 3 done ==="
curl -X POST "$BASE_URL/posts" \
    -H "Authorization: Bearer $TOKEN" \
    -F "title=two" \
    -F "description=description two" \
    -F "file=@$BASE_UPLOAD_URL/test_video_2.mp4"
echo "=== 2 / 3 done ==="
curl -X POST "$BASE_URL/posts" \
    -H "Authorization: Bearer $TOKEN" \
    -F "title=three" \
    -F "description=desription three" \
    -F "file=@$BASE_UPLOAD_URL/test_video_3.mp4"
echo "=== Uploading successful ==="
echo "=== Getting Posts ==="

curl -X GET "$BASE_URL/posts" \
    -H "$AUTH" -H "Content-Type: application/json" | jq