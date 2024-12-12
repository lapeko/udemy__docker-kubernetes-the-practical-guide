FROM node:22-alpine AS build
WORKDIR /app
COPY package*.json .
RUN npm install
COPY . .
RUN npm run build

FROM nginx:1.27.3-alpine
WORKDIR /usr/share/nginx/html
COPY --from=build /app/dist .
COPY nginx.conf.template /etc/nginx/nginx.conf.template

RUN apk add --no-cache gettext

EXPOSE 80

CMD ["sh", "-c", "envsubst '${TASKS_SERVICE_URL}' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf && nginx -g 'daemon off;'"]
