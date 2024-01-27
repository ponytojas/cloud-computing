# Etapa de construcción
FROM node:20.11.0-bookworm-slim as builder

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

ENV NODE_ENV=production
RUN npm run build

FROM node:20.11.0-bookworm-slim as runner

COPY --from=builder /app/.output /app/.output

WORKDIR /app
EXPOSE 3000
ENTRYPOINT ["node", ".output/server/index.mjs"]