FROM node:14-alpine as ui-builder
WORKDIR /build
COPY ui/package.json ui/yarn.lock ./
RUN yarn install --frozen-lockfile
COPY ui .
RUN yarn build --mode $MODE

FROM golang:1.16-alpine as go-builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=ui-builder /build/dist ui
RUN go build -o bin/server main.go

FROM alpine
WORKDIR /app
COPY --from=go-builder /build/bin/server .
CMD ./server
