# Golang grpc generator
protoc -I . --go_out=plugins=grpc:server redisimport/redisimport.proto

#Python grpc generator
# protoc -I . --python_out=client redisimport/redisimport.proto
python3 -m grpc_tools.protoc -I . --python_out=client --grpc_python_out=client redisimport/redisimport.proto