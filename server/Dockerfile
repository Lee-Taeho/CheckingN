FROM golang:1.17 as go-builder
WORKDIR /app
COPY ./ .
RUN go get .
RUN go build -o run

EXPOSE 8080

ENV HOST_IP_BINDING=0.0.0.0:8080
ENV FRONT_END_PATH=./client
ENV DB_URI="mongodb+srv://checkingn.sqysj.mongodb.net/myFirstDatabase?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile="
ENV DB_CERT_KEY="/app/database/certs/mongo_cert_key.pem"
ENV API_KEY="DYDKw4UzSsq53uN7hpgezA"
ENV API_SECRET="i6W7n221TAPgM7vzt0h2MYGt185STCAgSp6d"

CMD ["/app/run"]