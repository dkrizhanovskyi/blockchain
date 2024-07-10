### Blockchain Project on Go

---

## Table of Contents
1. [Project Description](#project-description)
2. [Features](#features)
3. [Technologies Used](#technologies-used)
4. [Installation](#installation)
5. [Usage](#usage)
6. [API Endpoints](#api-endpoints)
7. [Optimization](#optimization)
8. [Directory Structure](#directory-structure)
9. [Contributing](#contributing)
10. [License](#license)

---

## Project Description
This project is a blockchain implementation written in Go using the Proof of Stake (PoS) consensus algorithm. It features essential blockchain functionalities including block creation, transaction handling, and automatic block addition.

## Features
- **Block Creation**: Automatically creates blocks every 30 seconds.
- **Transactions**: Adds transactions to blocks with signature verification.
- **Caching**: Uses caching to improve API performance.
- **User Registration and Login**: Secure registration and login for users.
- **Parallel Transaction Processing**: Uses goroutines to validate transactions.

## Technologies Used
- **Programming Language**: Go
- **Cryptography**: ECDSA
- **Data Storage**: In-memory
- **API**: net/http

## Installation
1. **Clone the repository**:
   ```bash
   git clone https://github.com/dkrizhanovskyi/blockchain.git
   ```
2. **Navigate to the project directory**:
   ```bash
   cd blockchain
   ```
3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

## Usage
1. **Run the server**:
   ```bash
   go run main.go
   ```
2. **Access the API**:
   The API is accessible at `https://localhost:8082`.

## API Endpoints
- **Get Blockchain**:
  ```bash
  curl -X GET https://localhost:8082/blockchain
  ```
- **Send Transaction**:
  ```bash
  curl -X POST https://localhost:8082/send -d '{"sender":"<sender>", "recipient":"<recipient>", "amount":<amount>, "signature":"<signature>"}'
  ```

### Endpoint Details
1. **GET /blockchain**
   - Returns the current state of the blockchain.
2. **POST /write**
   - Adds a block with the provided transactions.
3. **POST /register**
   - Registers a new user.
4. **POST /login**
   - Logs in a user.
5. **POST /send**
   - Sends a transaction.
6. **POST /mine**
   - Mines new coins by creating a block with a transaction from the system to the recipient.

## Optimization
- **Profiling and Performance Analysis**: Uses `pprof` for profiling.
- **Parallel Processing**: Utilizes goroutines for transaction validation.
- **Minimizing Locking**: Reduces lock contention for improved performance.

## Directory Structure
```plaintext
blockchain/
│
├── api/
│   ├── api.go
│   ├── api_test.go
│   └── ...
│
├── blockchain/
│   ├── blockchain.go
│   ├── blockchain_test.go
│   └── ...
│
├── crypto/
│   ├── crypto.go
│   ├── crypto_test.go
│   └── ...
│
├── network/
│   ├── network.go
│   ├── network_test.go
│   └── ...
│
├── static/
│   ├── index.html
│   ├── style.css
│   └── script.js
│
├── main.go
├── go.mod
└── go.sum
```

## Contributing
1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

