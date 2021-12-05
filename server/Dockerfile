FROM golang:1.17 as go-builder
WORKDIR /app
COPY ./ .
RUN go get .
RUN go build

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=go-builder /app/server /app/server

EXPOSE 8080

### Set Env Vars
ENV HOST_IP_BINDING=0.0.0.0:8080
ENV FRONT_END_PATH=./client
ENV DB_URI="mongodb+srv://ayush:hello123@checkingn.sqysj.mongodb.net/users?retryWrites=true&w=majority"
ENV API_KEY="DYDKw4UzSsq53uN7hpgezA"
ENV API_SECRET="i6W7n221TAPgM7vzt0h2MYGt185STCAgSp6d"

# Run the binaries
CMD ["/app/server"]