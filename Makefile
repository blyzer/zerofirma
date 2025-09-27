.PHONY: all ca validator storage zkproof test

all: ca validator storage

ca:
    docker-compose build ca

validator:
    docker-compose build validator

storage:
    docker-compose build storage

zkproof:
    cd zkproof && npm install && circom circuit/proof.circom --r1cs

test:
    go test ./ca-service/internal/issuance
    go test ./validator/pkg
    go test ./storage/pkg
    cd zkproof && go test ./pkg
