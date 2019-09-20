from random import randint

import grpc

import cheese_pb2 as pb


def main ():

    with grpc.insecure_channel('localhost:4000') as channel:
        client = pb.CheeseStub(channel)

        while True:

            cheese_type = randint(100) % 5
            order = pb.CheeseRequest(type=cheese_type)

            cheese = client.Order(order)
            print("[gRPC] Received={0}", cheese.type)
