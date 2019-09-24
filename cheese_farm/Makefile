#
# Makefile to rule project tasks
#


# Binary directory
BIN := bin

help:
	@echo "gRPC Golang examples"
	@echo
	@echo "Available target rule:"
	@echo
	@echo " build                     Builds all binary files"
	@echo " server                    Builds Server binary"
	@echo " client                    Builds Client binary"
	@echo " run_server                Runs server binary"
	@echo " run_client                Runs client binary"
	@echo " protoc_python             Runs the protoc compiler for Python"
	@echo " build_pyclient            Builds virtualenv and install deps"
	@echo " run_pyclient              Runs the Python client"
	@echo " clean                     Removes binary files"
	@echo


build: server client


.PHONY:server
server:
	go build -race -o $(BIN)/server cmd/server/main.go


.PHONY:client
client:
	go build -race -o $(BIN)/client cmd/client/main.go


run_server: $(BIN)/server
	@./$<


run_client: $(BIN)/client
	@./$<


protoc_python:
	@cd python && \
		pipenv run python -m grpc_tools.protoc \
						  -I ../cheese/ \
						  --python_out=. \
						  --grpc_python_out=. \
						  cheese.proto


build_pyclient:
	@cd python && pipenv install

run_pyclient:
	@cd python && pipenv run python client.py


clean:
	@rm -vf $(BIN)/{server,client}
