export PATH=$PATH:../.bin

grpcurl -plaintext -d '{"id": "1"}' \
localhost:50051 userapi.UserService/GetUser
