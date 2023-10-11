export PATH=$PATH:../.bin
grpcurl -plaintext -d '{"id": "1", "name": "John Doe", "email": "john@example.com"}' \
localhost:50051 userapi.UserService/CreateUser
