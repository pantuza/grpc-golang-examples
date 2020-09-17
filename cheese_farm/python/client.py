from random import randint
from time import sleep

import grpc

import cheese_pb2 as pb
import cheese_pb2_grpc

def main ():

    # Load Certificate file
    trusted_cert = open("../ssl/cert.pem", "rb").read()
    credentials = grpc.ssl_channel_credentials(root_certificates=trusted_cert)

    try:

        with grpc.secure_channel('localhost:4000', credentials) as channel:

            client = cheese_pb2_grpc.CheeseServiceStub(channel)

            while True:
                try:
                    cheese_type = randint(0, 4)
                    order = pb.CheeseRequest(type=cheese_type)

                    cheese = client.Order(order)
                    print("[gRPC] Received={0}".format(pb.CheeseType.Name(cheese.type)))

                except grpc._channel._Rendezvous as err:
                    print(err)
                    sleep(1)

    except KeyboardInterrupt:
        print("[gRPC] Teardown")


if __name__ == "__main__":

    main()
