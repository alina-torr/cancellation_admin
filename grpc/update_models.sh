cd ../ml/
rm ml_pb2_grpc.py ml_pb2.py
python3 -m grpc_tools.protoc --proto_path=../grpc ../grpc/ml.proto --python_out=. --grpc_python_out=.

cd ../backend/
rm -r ml
protoc --proto_path=../grpc ../grpc/ml.proto --go_out=. --go-grpc_out=. --experimental_allow_proto3_optional