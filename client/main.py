# Import grpc
import grpc
import os
import glob
import json

# import the generated classes
import redisimport_pb2
import redisimport_pb2_grpc

# # connect grpc server
grpc_server = os.environ.get('GRPC_SERVER', 'localhost:50051')

with grpc.insecure_channel(grpc_server) as channel:
    stub = redisimport_pb2_grpc.RedisImportStub(channel)
    # List user files
    for filename in glob.glob('users/*.json'):
        # Append users from file
        print("Reading file: " + filename)
        with open(os.path.join(os.getcwd(), filename), 'r') as file:
            data = json.load(file)
            for user in data:
                response = stub.Import(redisimport_pb2.User(
                    id=user["id"],
                    first_name=user["first_name"],
                    last_name=user["last_name"],
                    email=user["email"],
                    gender=user["gender"],
                    ip_address=user["ip_address"],
                    user_name=user["user_name"],
                    agent=user["agent"],
                    country=user["country"]
                ))

                if response:
                    print(response.message)
        print("File success: " + filename)
