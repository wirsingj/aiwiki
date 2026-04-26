FROM golang:1.26-alpine AS build

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/aiwiki-server ./cmd/server

FROM alpine:3.22

RUN apk add --no-cache ca-certificates wget
RUN addgroup -S aiwiki && adduser -S aiwiki -G aiwiki
USER aiwiki

WORKDIR /app
COPY --from=build /out/aiwiki-server /app/aiwiki-server

ENV PORT=8080
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget -q -O /dev/null "http://127.0.0.1:${PORT}/api/health" || exit 1

CMD ["/app/aiwiki-server"]
