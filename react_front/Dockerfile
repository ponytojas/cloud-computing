# Use an official Node.js image as the base image
FROM node:20.15.1-alpine3.20

# Set environment variables
ENV NPM_CONFIG_PREFIX=/home/node/.npm-global
ENV PATH=$PATH:/home/node/.npm-global/bin

# Install pnpm globally
RUN npm i -g pnpm

# Install curl
RUN apk add --no-cache curl

# Create and set the working directory
WORKDIR /home/node/app

# Copy package.json and package-lock.json (if present) first to leverage Docker cache
COPY package*.json ./

# Install dependencies
RUN pnpm install

# Copy the rest of the application code
COPY . .

# Remove .env and rename .env.prod to .env
RUN rm -f .env && mv .env.prod .env

# Build the application
RUN pnpm build

# Change ownership of the application files
RUN chown -R node:node /home/node/app

# Switch to non-root user
USER node

# Specify the command to run the application
CMD ["pnpm", "preview", "--host", "0.0.0.0", "--port", "5555"]

# Expose the application port
EXPOSE 5555
