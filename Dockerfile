### Server
FROM golang:1.17 as go-builder
WORKDIR /app
COPY server .
RUN go get .
RUN go build

### Client
FROM node:17-alpine3.12 as node-builder
WORKDIR /app
ENV NODE_OPTIONS=--openssl-legacy-provider
COPY client/package.json ./
RUN npm install
COPY client/ .
RUN npm run build

### Define Listen Ports 
WORKDIR /app

EXPOSE 3000
EXPOSE 8080

### Set Env Vars
ENV HOST_IP_BINDING=localhost:8080
ENV FRONT_END_PATH=../client
ENV DB_URI="mongodb+srv://ayush:hello123@checkingn.sqysj.mongodb.net/users?retryWrites=true&w=majority"
ENV API_KEY="DYDKw4UzSsq53uN7hpgezA"
ENV API_SECRET="i6W7n221TAPgM7vzt0h2MYGt185STCAgSp6d"

# Run the binaries
CMD ["/app/server", "-FAP=./build"]