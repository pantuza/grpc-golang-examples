#
# Makefile to rule project tasks
#


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
	@echo " clean                     Removes binary files"
	@echo


build: server client


.PHONY:server
server:
	go build -race -o cheese_server server/main.go


.PHONY:client
client:
	go build -race -o cheese_client client/main.go


run_server: cheese_server
	@./cheese_server


run_client: cheese_client
	@./cheese_client


clean:
	@rm -vf cheese_server cheese_client
