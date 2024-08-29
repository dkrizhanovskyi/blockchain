# API Documentation

This directory contains the OpenAPI/Swagger documentation for the Blockchain API.

## Overview

The API documentation provides a detailed description of all available endpoints, including their request parameters, responses, and data schemas. This documentation is essential for developers who want to interact with the blockchain through HTTP requests.

## Files

- **swagger.json**: The API specification in JSON format.
- **swagger.yaml**: The API specification in YAML format.

These files describe the RESTful API for interacting with the blockchain, including endpoints for adding blocks, retrieving blockchain data, and validating the blockchain.

## Viewing the Documentation

You can view and interact with the API documentation using Swagger UI or any other OpenAPI-compatible tool.

### Using Swagger UI

1. Open [Swagger Editor](https://editor.swagger.io/).
2. Load the `swagger.yaml` or `swagger.json` file by either:
   - Copy-pasting the file contents into the editor.
   - Using the "File" menu to import the file from your local machine.
3. Once loaded, you can explore the API endpoints, view detailed request/response formats, and try out the requests directly from the browser.

### Using a Local Swagger UI Instance

If you prefer to run Swagger UI locally:

1. Clone the Swagger UI repository:

   ```bash
   git clone https://github.com/swagger-api/swagger-ui.git
   cd swagger-ui
   ```

2. Serve the Swagger UI files using a simple HTTP server:

   ```bash
   python3 -m http.server
   ```

3. Open your web browser and navigate to `http://localhost:8000`.

4. Load the `swagger.yaml` or `swagger.json` file by pasting its contents or by specifying the URL to the file.

### Generating Client Code

You can generate client code for various programming languages using Swagger Codegen or other tools that support OpenAPI specifications.

For example, to generate a Go client:

```bash
swagger-codegen generate -i swagger.yaml -l go -o ./generated-client
```

Replace `go` with the desired programming language, and adjust the output directory as needed.

## Endpoints Overview

The API includes the following endpoints:

- **`GET /getblockchain`**: Retrieves the entire blockchain.
- **`POST /addblock`**: Adds a new block to the blockchain.
- **`GET /block?index=INDEX`**: Retrieves a specific block by its index.
- **`GET /lastblock`**: Retrieves the last block in the blockchain.
- **`GET /validate`**: Validates the integrity of the blockchain.

Each endpoint is fully documented in the `swagger.yaml` and `swagger.json` files, including the request parameters and expected responses.

## Updating the Documentation

To update the API documentation:

1. Modify the `swagger.yaml` or `swagger.json` files directly using an OpenAPI editor like Swagger Editor.
2. Commit the updated files to your repository.

Ensure that the documentation stays up-to-date with any changes made to the API in the codebase.

## Contribution

If you would like to contribute to the API documentation, please fork the repository and submit a pull request with your changes. Ensure that all new or modified endpoints are fully documented.

## License

The API documentation is licensed under the MIT License. See the [LICENSE](../../LICENSE) file for more information.
