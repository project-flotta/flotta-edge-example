FROM golang:1.18-alpine AS builder

# this file for the development only
RUN mkdir /app
ADD . /app
WORKDIR /app

# RUN go get -u github.com/cosmtrek/air
RUN go mod tidy && go mod vendor
RUN go build -o app ./

FROM alpine:latest AS runner

WORKDIR /app
COPY --from=builder /app .
RUN chmod +x app

EXPOSE 8080
CMD ["./app"]
