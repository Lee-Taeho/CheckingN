export HOST_IP_BINDING=localhost:8080
export FRONT_END_PATH=../client
export DB_URI="mongodb+srv://checkingn.sqysj.mongodb.net/myFirstDatabase?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile="
export DB_CERT_KEY="/Users/ayushmaheshwari/CS_classes/CS160/2021-fall-cs160-chocolate-cake/backend/database/certs/mongo_cert_key.pem"
export API_KEY="DYDKw4UzSsq53uN7hpgezA"
export API_SECRET="i6W7n221TAPgM7vzt0h2MYGt185STCAgSp6d"
go build
./server
