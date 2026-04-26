FROM node:20-alpine AS build

WORKDIR /src

COPY package*.json ./
RUN npm ci

COPY . .
RUN npm run build

FROM nginx:1.29-alpine

COPY nginx/default.conf.template /etc/nginx/templates/default.conf.template
COPY nginx/docker-entrypoint.d/40-aiwiki-config.sh /docker-entrypoint.d/40-aiwiki-config.sh
COPY --from=build /src/dist /usr/share/nginx/html

RUN chmod +x /docker-entrypoint.d/40-aiwiki-config.sh

ENV BACKEND_UPSTREAM=http://backend:8080
ENV AIWIKI_API_BASE_URL=
ENV PORT=80
EXPOSE 80

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget -q -O /dev/null "http://127.0.0.1:${PORT}/" || exit 1
