FROM node:20-alpine as build-stage

WORKDIR /app
COPY package*.json ./
COPY yarn.lock* ./
RUN npm install --frozen-lockfile
COPY . .
RUN npm run build



FROM nginx:stable-alpine as production-stage

WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=build-stage /app/dist .
ENTRYPOINT ["nginx", "-g", "daemon off;"]
