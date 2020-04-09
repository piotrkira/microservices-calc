import grpc
from concurrent import futures
import time
import os

import service_pb2
import service_pb2_grpc

class MulDivService(service_pb2_grpc.MulDivServicer):
    def Mul(self, request, context):
        print("multiply requested")
        response = service_pb2.Result()
        result = request.a * request.b
        response.content = "The result is {0} calculated by {1}".format(result, os.uname()[1])
        return response
    def Div(self, request, context):
        print("dividing requested")
        response = service_pb2.Result()
        result = request.a // request.b
        response.content = "The result is {0} calculated by {1}".format(result, os.uname()[1])
        return response


server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

service_pb2_grpc.add_MulDivServicer_to_server(
    MulDivService(), server)

print('muldiv service started')
server.add_insecure_port('[::]:7777')
server.start()

try:
    while True:
        time.sleep(100000)
except KeyboardInterrupt:
    server.stop(0)