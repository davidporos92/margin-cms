FROM node:24.21.0-alpine

WORKDIR /app

COPY package.json package.json
COPY package-lock.json package-lock.json

RUN npm install -D

COPY openapi openapi

CMD ["npm", "run", "spec:docs"]