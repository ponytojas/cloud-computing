FROM node:20.15-alpine3.19
RUN apk add --no-cache dumb-init
ENV NODE_ENV production
WORKDIR /usr/src/app
COPY --chown=node:node --chmod=755 package*.json  ./
RUN npm ci --ignore-scripts --only=production
COPY --chown=node:node --chmod=755 index.js .
COPY --chown=node:node --chmod=755 src/ ./src
USER node
CMD ["dumb-init", "node", "index.js"]
