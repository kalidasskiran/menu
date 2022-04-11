docker run -d --name mongodb -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password -p 27017:27017 mongo:4.4.3
 openssl rand -base64 12 | docker secret create mongodb_password-
 docker swarm init
 docker swarm join-token manager
  -e MONGO_INITDB_ROOT_PASSWORD_FILE=/run/secrets/mongodb_password
   MONGO_URI="mongodb://admin:password@localhost:27017/test?authSource=admin" go run main.go

 MONGO_URI="mongodb://admin:password@localhost:27017/test?authSource=admin" MONGO_DATABASE=demo go run main.go

 