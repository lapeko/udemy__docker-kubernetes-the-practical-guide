FROM node:alpine

WORKDIR /app

COPY package.json package-lock.json /

RUN npm ci

COPY . .

ARG API_URL=http://localhost:3001
RUN sed -i 's|$API_URL|'${API_URL}'|g' /app/src/environments/env.ts

EXPOSE 8000

CMD ["npm", "run", "dev"]