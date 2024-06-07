#!/bin/bash

# Function to check for jq command and print install instructions if not present
check_jq() {
  if ! command -v jq &> /dev/null; then
    echo "jq could not be found."
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
      echo "Please install jq using the following command:"
      echo "sudo apt-get update && sudo apt-get install -y jq"
    elif [[ "$OSTYPE" == "darwin"* ]]; then
      echo "Please install jq using Homebrew with the following command:"
      echo "brew install jq"
    else
      echo "Unsupported OS. Please install jq manually."
      exit 1
    fi
    exit 1
  fi
}

# Check and prompt to install jq if necessary
check_jq

# Create user
USER_RESPONSE=$(curl -s -X POST http://127.0.0.1/core/v1/user/register \
-H "Content-Type: application/json" \
-d '{
    "username": "user1",
    "email": "user1@user.com",
    "password": "pass"
}')

echo "User response: $USER_RESPONSE"

# Log in user
LOGIN_RESPONSE=$(curl -s -X POST http://127.0.0.1/core/v1/user/login \
-H "Content-Type: application/json" \
-d '{
    "username": "user1",
    "password": "pass"
}')

echo "Login response: $LOGIN_RESPONSE"  

# Extract token from login response
TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.token')

echo "Token is: $TOKEN"

# Create products and update stock
for i in {1..10}
do
  PRODUCT_RESPONSE=$(curl -s -X POST http://127.0.0.1/core/v1/product \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d "{
      \"name\": \"Producto $i\",
      \"pricing\": $((RANDOM % 100 + 1)),
      \"description\": \"test-$i\"
  }")

  STOCK_RESPONSE=$(curl -s -X POST http://127.0.0.1/core/v1/stock/$i \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d "{
      \"quantity\": $((RANDOM % 101))
  }")

  echo "Product $i response: $PRODUCT_RESPONSE"
  echo "Stock $i response: $STOCK_RESPONSE"
done
