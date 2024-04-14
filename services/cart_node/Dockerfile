FROM node:20.12.0-alpine3.19
RUN apk add --no-cache dumb-init
ENV NODE_ENV production
WORKDIR /usr/src/app
COPY --chown=node:node package*.json ./
RUN npm ci --only=production
COPY --chown=node:node . .
USER node
CMD ["dumb-init", "node", "index.js"]