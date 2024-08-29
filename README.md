# Blockchain Project

This project implements a basic blockchain with a peer-to-peer (P2P) network and an API interface. The project is built using Go for the blockchain and network layers, and Rust for optimized cryptographic operations.

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [P2P Network](#p2p-network)
- [Tests](#tests)
- [Contribution](#contribution)
- [License](#license)

## Features

- **Blockchain Implementation**: A simple blockchain with basic functionality like adding blocks and validating the chain.
- **P2P Network**: Nodes can connect to each other to share blocks and validate the blockchain.
- **API Interface**: A RESTful API for interacting with the blockchain, adding blocks, and validating the chain.
- **Rust Integration**: Cryptographic functions (e.g., hashing) are implemented in Rust for better performance.
- **Modular Codebase**: The project is organized into packages for easy maintenance and extension.

## Project Structure

```plaintext
.github/
│   └── workflows/
│       └── go.yml          # GitHub Actions workflow for CI/CD
blockchain_app/             # Binary output of the compiled Go application
cmd/
│   └── blockchain/
│       └── main.go         # Main entry point of the application
docs/
│   └── api_docs/           # API documentation
internal/
│   ├── api/                # API handlers and routes
│   ├── blockchain/         # Blockchain implementation
│   ├── crypto/             # Cryptographic functions, including Rust integration
│   ├── network/            # P2P network implementation
│   ├── p2p/                # Node implementation for P2P network
│   └── utils/              # Utility functions (e.g., logging)
pkg/
│   ├── config/             # Configuration management
│   ├── errors/             # Custom error types and handling
│   └── middleware/         # HTTP middleware functions
rust/
│   └── rust_crypto/        # Rust project for cryptographic functions
static/
│   ├── index.html          # Frontend interface for interacting with the API
│   ├── script.js           # JavaScript for frontend logic
│   └── style.css           # CSS for frontend styling
README.md                   # This file
```

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/dkrizhanovskyi/blockchain.git
   cd blockchain
   ```

2. **Set up Go:**

   Ensure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).

3. **Set up Rust:**

   Ensure you have Rust installed. You can install it using [rustup](https://rustup.rs/).

4. **Build the Rust component:**

   ```bash
   cd rust/rust_crypto
   cargo build --release
   ```

5. **Build the Go application:**

   ```bash
   go build -o blockchain_app ./cmd/blockchain
   ```

## Running the Application

1. **Start a node:**

   ```bash
   ./blockchain_app
   ```

   By default, the node will listen on `localhost:3001` and the API server will run on `localhost:8080`.

2. **Start additional nodes:**

   You can start additional nodes by specifying a different `NODE_ADDRESS`:

   ```bash
   NODE_ADDRESS="localhost:3002" ./blockchain_app
   ```

3. **Connect nodes:**

   To connect nodes, you can use the `INITIAL_PEER` environment variable:

   ```bash
   INITIAL_PEER="localhost:3001" NODE_ADDRESS="localhost:3002" ./blockchain_app
   ```

## API Endpoints

- **`GET /getblockchain`**: Retrieves the entire blockchain.
- **`POST /addblock`**: Adds a new block to the blockchain.
- **`GET /block?index=INDEX`**: Retrieves a specific block by its index.
- **`GET /lastblock`**: Retrieves the last block in the blockchain.
- **`GET /validate`**: Validates the integrity of the blockchain.

## P2P Network

The P2P network allows nodes to connect to each other and share blocks. When a node receives a new block, it broadcasts the block to all its connected peers.

## Tests

To run the tests for both Go and Rust components:

```bash
# Run Go tests
go test ./...

# Run Rust tests
cd rust/rust_crypto
cargo test
```

## Contribution

Contributions are welcome! Please fork this repository and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
