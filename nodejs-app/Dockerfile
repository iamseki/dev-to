FROM node:alpine

WORKDIR /app

COPY . .

RUN npm ci

ENTRYPOINT ["node", "app.js"]