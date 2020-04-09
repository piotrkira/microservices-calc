FROM python:3

ADD main.py /
ADD service_pb2_grpc.py /
ADD service_pb2.py /

RUN pip install grpcio
RUN pip install grpcio-tools

CMD [ "python", "./main.py"]