# Build stage
FROM node:20-alpine AS build-stage
WORKDIR /app
COPY package*.json package-lock.json ./
RUN npm install --frozen-lockfile --ignore-scripts

# Copy only the necessary source files
COPY vite.config.js .
COPY index.html .
COPY src src
COPY public public

RUN npm run build

# Production stage
FROM nginx:stable-alpine AS production-stage

# Create a non-root user and switch to it
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /usr/share/nginx/html
# Clear any existing files in the directory
RUN rm -rf ./*
# Copy only the necessary build files from the previous stage
COPY --from=build-stage /app/dist .

# Change ownership of the directory and its contents to the non-root user
RUN chown -R appuser:appgroup /usr/share/nginx/html

# Switch to non-root user
USER appuser

# Start Nginx and make sure it doesn't run as a daemon
ENTRYPOINT ["nginx", "-g", "daemon off;"]
