#!/bin/bash

# Script to create a virtual environment, install dependencies, and run the Python script

# Define variables
ENV_DIR="env"
REQUIREMENTS_FILE="requirements.txt"
PYTHON_SCRIPT="main.py"
CHROMEDRIVER_PATH=$(which chromedriver) # Get the chromedriver path

if [ -z "$CHROMEDRIVER_PATH" ]; then
    echo "chromedriver not found. Please install it using Homebrew or specify the correct path."
    exit 1
fi

# Check if the virtual environment directory exists
if [ -d "$ENV_DIR" ]; then
    echo "Virtual environment already exists. Skipping creation."
else
    # Create a virtual environment
    python3 -m venv $ENV_DIR
    echo "Virtual environment created."
fi

# Activate the virtual environment
source $ENV_DIR/bin/activate

# Upgrade pip
pip install --upgrade pip

# Install dependencies
pip install -r $REQUIREMENTS_FILE

# Check for debug flag
DEBUG_FLAG=""
if [ "$1" == "--debug" ]; then
    DEBUG_FLAG="--debug"
fi

# Run the Python script with the CHROMEDRIVER_PATH environment variable
CHROMEDRIVER_PATH=$CHROMEDRIVER_PATH python $PYTHON_SCRIPT $DEBUG_FLAG

# Deactivate the virtual environment
deactivate
