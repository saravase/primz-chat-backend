# primz-chat-backend

### Mongodb container:

    docker run -it --rm --name mongodb_container -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin -v mongodata:/data/db -d -p 27017:27017 mongo:4.4

### Execute mongodb container:

    docker exec -it mongodb_container /bin/bash

#### Connect admin database:
    mongo -u admin -p admin --authenticationDatabase admin

#### Create database:
    use primz_chat_backend;

#### Create user and set role :
    db.createUser({user: 'primz', pwd: 'primz', roles:[{'role': 'readWrite', 'db': 'primz_chat_backend'}]})

