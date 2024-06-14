#!/bin/bash

# Check if the user has jq and parallel installed, if not, prompt the user to install them
if ! [ -x "$(command -v jq)" ]; then
  echo 'Error: jq is not installed. Please install jq before running this script.' >&2
  echo 'Ubuntu: sudo apt-get install jq'
  echo 'Mac: brew install jq'
  exit 1
fi

if ! [ -x "$(command -v parallel)" ]; then
  echo 'Error: parallel is not installed. Please install parallel before running this script.' >&2
  echo 'Ubuntu: sudo apt-get install parallel'
  echo 'Mac: brew install parallel'
  exit 1
fi

# URL for user registration
REGISTER_URL='http://127.0.0.1/core/v1/user/register'
# URL for user login
LOGIN_URL='http://127.0.0.1/core/v1/user/login'
# URL for product creation
PRODUCT_URL='http://127.0.0.1/core/v1/product'
# URL for stock update
STOCK_URL_BASE='http://127.0.0.1/core/v1/stock/'

# URL for adding to cart
ADD_TO_CART_URL='http://localhost/cart/v1/add-to-cart'
# URL for getting total payment
PAYMENT_TOTAL_URL='http://localhost/payment/v1/total/1'

TOKEN_FILE="token.txt"
> $TOKEN_FILE # Create or clear the token file

# Array of image URLs
IMAGE_URLS=("https://cdn.shopify.com/s/files/1/0746/7391/4132/files/LightIII-Float-ProductSite_1600x.png?v=1718070454" "https://www.nobullproject.com/cdn/shop/products/dd436e6d31839fdd4e9d827ecf83acc84abc2d79_1cd8725e-8adb-400c-89d6-6c2bedfaf546.jpg?v=1688016515" "https://bureau.ams3.digitaloceanspaces.com/goods/prod/uploads/store/59d7645003a61d648a4bd6b8d76ceff4.png")

# Check if IMAGE_URLS array is not empty
if [ ${#IMAGE_URLS[@]} -eq 0 ]; then
  echo 'Error: IMAGE_URLS array is empty. Please add URLs to the IMAGE_URLS array.' >&2
  exit 1
fi

register_and_process_user() {
  i=$1
  REGISTER_URL=$2
  LOGIN_URL=$3
  PRODUCT_URL=$4
  STOCK_URL_BASE=$5
  TOKEN_FILE=$6

  # Register the user
  register_response=$(curl --silent --write-out "HTTPSTATUS:%{http_code}" --location $REGISTER_URL \
  --header 'Content-Type: application/json' \
  --data-raw "{
    \"username\": \"user$i\",
    \"email\": \"user$i@user.com\",
    \"password\": \"pass\"
  }")

  # Extract the body and status
  register_body=$(echo $register_response | sed -e 's/HTTPSTATUS\:.*//g')
  register_status=$(echo $register_response | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')

  if [ $register_status -eq 200 ] || [ $register_status -eq 201 ]; then
    echo "Registration $i: OK"

    # Log in with the registered user
    login_response=$(curl --silent --write-out "HTTPSTATUS:%{http_code}" --location $LOGIN_URL \
    --header 'Content-Type: application/json' \
    --data-raw "{
      \"username\": \"user$i\",
      \"password\": \"pass\"
    }")

    # Extract the body and status
    login_body=$(echo $login_response | sed -e 's/HTTPSTATUS\:.*//g')
    login_status=$(echo $login_response | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')

    if [ $login_status -eq 200 ]; then
      echo "Login $i: OK"

      # Extract the token
      token=$(echo $login_body | jq -r '.token')

      # Save the token to the file if this is the first user
      if [ $i -eq 1 ]; then
        echo $token > $TOKEN_FILE
      fi

      price=$(( ( RANDOM % 200 )  + 1 ))
      rating=$(( ( RANDOM % 5 )  + 1 ))

      # Select a random image URL
      image_url=${IMAGE_URLS[$RANDOM % ${#IMAGE_URLS[@]}]}

      # Create a product
      product_response=$(curl --silent --write-out "HTTPSTATUS:%{http_code}" --location $PRODUCT_URL \
      --header 'Content-Type: application/json' \
      --header "Authorization: $token" \
      --data-raw "{
        \"name\": \"Producto $i\",
        \"pricing\": $price,
        \"description\": \"test\",
        \"rating\": $rating,
        \"imageUrl\": \"$image_url\"
      }")

      # Extract the product creation status
      product_status=$(echo $product_response | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')

      if [ $product_status -eq 200 ] || [ $product_status -eq 201 ]; then
        echo "Product $i: Created"

        # Generate a random quantity between 1 and 200
        quantity=$(( ( RANDOM % 200 )  + 1 ))

        # Update stock
        stock_response=$(curl --silent --write-out "HTTPSTATUS:%{http_code}" --location "${STOCK_URL_BASE}${i}" \
        --header 'Content-Type: application/json' \
        --header "Authorization: $token" \
        --data-raw "{
          \"quantity\": $quantity
        }")

        # Extract the stock update status
        stock_status=$(echo $stock_response | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')

        if [ $stock_status -eq 200 ] || [ $stock_status -eq 201 ]; then
          echo "Stock $i: Updated with quantity $quantity"
        else
          echo "Stock $i: Failed to update with status $stock_status"
        fi

      else
        echo "Product $i: Failed to create with status $product_status"
      fi

    else
      echo "Login $i: Failed with status $login_status"
    fi

  else
    echo "Registration $i: Failed with status $register_status"
  fi
}

export -f register_and_process_user

# Run the user registration and product creation in parallel
seq 1 200 | parallel -j 10 register_and_process_user {} $REGISTER_URL $LOGIN_URL $PRODUCT_URL $STOCK_URL_BASE $TOKEN_FILE

add_to_cart_and_check_total() {
  i=$1
  ADD_TO_CART_URL=$2
  PAYMENT_TOTAL_URL=$3
  TOKEN_FILE=$4

  # Read the stored token
  stored_token=$(cat $TOKEN_FILE)

  # Add to cart
  add_to_cart_response=$(curl --silent --write-out "HTTPSTATUS:%{http_code}" --location $ADD_TO_CART_URL \
  --header 'Content-Type: application/json' \
  --header "Authorization: $stored_token" \
  --data-raw "{
    \"userId\": 1,
    \"productId\": $i,
    \"quantity\": 1
  }")

  # Extract the add to cart status
  add_to_cart_status=$(echo $add_to_cart_response | tr -d '\n' | sed -e 's/.*HTTPSTATUS://') 
  if [ $add_to_cart_status -eq 200 ] || [ $add_to_cart_status -eq 201 ]; then
    echo "Add to cart $i: OK"
  else
    echo "Add to cart $i: Failed with status $add_to_cart_status"
  fi

  # Check payment total
  payment_total_response=$(curl --silent --write-out "HTTPSTATUS:%{http_code}" --location $PAYMENT_TOTAL_URL \
  --header 'Content-Type: application/json' \
  --header "Authorization: $stored_token")

  # Extract the payment total status
  payment_total_status=$(echo $payment_total_response | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')

  if [ $payment_total_status -eq 200 ]; then
    echo "Payment total $i: OK"
  else
    echo "Payment total $i: Failed with status $payment_total_status"
  fi
}

export -f add_to_cart_and_check_total

# Call the function in parallel
seq 1 150 | parallel -j 10 add_to_cart_and_check_total {} $ADD_TO_CART_URL $PAYMENT_TOTAL_URL $TOKEN_FILE
