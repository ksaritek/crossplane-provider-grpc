export PATH=$PATH:../.bin

grpcurl -plaintext -d '{"id": "1", "name": "John Updated", "email": "johnupdated@example.com"}' \
localhost:50051 userapi.UserService/UpdateUser
