# gRPC Service with Clean Architecture
This repository is designed for learning purposes and showcases a simple gRPC service using **Clean Architecture**, a **Dependency Injection (DI) container**, and best practices. It demonstrates how to set up and structure a gRPC microservice for maintainability and scalability.

### `infrastructure` Branch
Please switch to the `infrastructure` branch, which contains the **docker-compose.yml** configuration file required to run the service locally. This configuration ensures that all necessary services are orchestrated seamlessly using Docker.

### Running and Testing
A **Makefile** is provided to simplify the process of running and testing the gRPC service. It contains the essential commands needed to operate the service locally and run tests.

#### Makefile commands
- **Run the service**: Spins up the gRPC service locally.
- **Testing the service**: This repository relies on grpcurl for calling and testing the gRPC service endpoints directly. grpcurl serves as a command-line tool for interacting with gRPC services, allowing you to invoke remote procedures similarly to curl for HTTP APIs.

## Intsalling `grpcurl`
macOS (using Homebrew)
```sh
brew install grpcurl
```

fedora (using dnf)
```sh
dnf install grpcurl
```

## Future Improvements
This repository will be continuously improved and expanded to include additional features, better error handling, enhanced logging, and further examples of best practices.

Contributions and suggestions are welcome to help make this service more robust and adaptable for real-world applications.
