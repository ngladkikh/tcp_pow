# POW

## Task

Design and implement “Word of Wisdom” tcp server. 

* TCP server should be protected from DDOS attacks with the Proof of Work, the challenge-response protocol should be used. 
* The choice of the POW algorithm should be explained. 
* After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes. 
* Docker file should be provided both for the server and for the client that solves the POW challenge

## PoW Algorithm

The chosen PoW algorithm is based on the concept of finding a nonce such that the hash of the nonce concatenated with a given challenge string results in a hash with a specific number of leading zeros. 
This algorithm was chosen for its simplicity and effectiveness in demonstrating the PoW concept. 
It is computationally difficult enough to deter denial-of-service attacks but still feasible for a legitimate client to solve within a reasonable amount of time.

## HOWTO

### Using Make Targets

This project uses a Makefile for easy building, testing, and running of the server and client. The available targets are:

- `make test`: Runs Go unit tests for the entire project.
- `make server`: Runs the server locally.
- `make client`: Runs the client locally.
- `make build`: Builds Docker images for both the server and the client.
- `make integration`: Runs both the server and client using `docker-compose` for integration testing.

### Structure

The repository is structured as follows:

- cmd/: Contains the main applications for the client and server.
    - client/: Client application.
    - server/: Server application.
- internal/: Internal application code.
    - pow/: Implementation of the Proof of Work algorithm.
    - quotes/: Management of the quotes used in the server response.
- pkg/: Library code that could be used by external applications.
    - tcp/: TCP communication handling.
    - config/: Shared between sever and client configuration helpers.
- Dockerfile.server and Dockerfile.client: Dockerfiles for building the server and client applications.