FROM golang:alpine AS isbpbuildmaster
WORKDIR /go/src/github.com/nikitamirzani323/wl_super_backend_api
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest as totosvelterelease
WORKDIR /app
RUN apk add tzdata
RUN mkdir -p ./frontend/public
COPY --from=isbpbuildmaster /go/src/github.com/nikitamirzani323/wl_super_backend_api/app .
COPY --from=isbpbuildmaster /go/src/github.com/nikitamirzani323/wl_super_backend_api/env-sample /app/.env

ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

EXPOSE 2011
CMD ["./app"]