# Payment Gateway Documentation

## How to run

### Prerequisites

- Docker installed on your machine.
- Git installed on your machine.

### How to run the whole application

1. Clone the repository from https://github.com/Lionel-Wilson/payment-gateway
2. Open a terminal and run the following command. Make sure you're in the root of the repository:

```
docker-compose up --build
```

3. After it's built, visit the merchant site at http://localhost:4200 . Alternatively you can test the API using swagger at http://localhost:8080/swagger/index.html

## Endpoints

### 1. Process a Payment

- **Endpoint**: `/payments`
- **Method**: `POST`
- **Description**: Processes a payment through the payment gateway.
- **Request Body**:
  ```json
  {
    "firstName": "John",
    "lastName": "Doe",
    "cardNumber": "4111111111111111",
    "expiryDate": "12/24",
    "amount": 100.5,
    "currencyCode": "USD",
    "cvv": "123"
  }
  ```

#### Responses

- **Success (201 Created)**:

  ```json
  {
    "id": "PAY-1625843728243722000",
    "status": "payment_paid",
    "responseSummary": "Approved"
  }
  ```

- **Failure (402 Payment Required)**:

  ```json
  {
    "id": "PAY-1625843728243722000",
    "status": "payment_declined",
    "responseSummary": "Insufficient funds"
  }
  ```

- **Validation Error (422 Unprocessable Entity)**:

  ```json
  {
    "statusCode": 422,
    "message": "Validation failed",
    "errors": ["FirstName is required", "CardNumber is required"]
  }
  ```

- **Failure (500 Internal Server Error)**:

  ```json
  {
    "statusCode": 500,
    "message": "Something went wrong. Please try again later."
  }
  ```

### 2. Retrieve Payment Details

- **Endpoint**: `/payments/{id}`
- **Method**: `GET`
- **Description**: Retrieves the details of a previously made payment using its identifier..
- **Path Parameters**: Retrieves the details of a previously made payment using its identifier..
  - **id (string)**: The unique identifier of the payment.

#### Responses

- **Success (200 OK)**:

  ```json
  {
    "id": "PAY-1625843728243722000",
    "firstName": "John",
    "lastName": "Doe",
    "cardNumber": "************1111",
    "expiryDate": "12/24",
    "amount": 100.5,
    "currencyCode": "USD",
    "status": "payment_paid",
    "statusCode": 10000
  }
  ```

- **Not Found (404 Not Found)**:

  ```json
  {
    "error": "Payment not found"
  }
  ```

### 3. Retrieve All Payments

- **Endpoint**: `/payments`
- **Method**: `GET`
- **Description**: Retrieves the details of all previously made payments.

#### Responses

- **Success (200 OK)**:

  ```json
  [
    {
      "id": "PAY-1719555378154588956",
      "firstName": "Gee",
      "lastName": "Wilson",
      "cardNumber": "************1032",
      "expiryDate": "11/27",
      "amount": 254.5,
      "currencyCode": "GBP",
      "status": "payment_paid",
      "statusCode": 10000
    },
    {
      "id": "PAY-1719555406263509469",
      "firstName": "Lionel",
      "lastName": "Wilson",
      "cardNumber": "************1032",
      "expiryDate": "11/27",
      "amount": 154.5,
      "currencyCode": "GBP",
      "status": "payment_paid",
      "statusCode": 10000
    }
  ]
  ```

- **Not Found (404 Not Found)**:

  ```json
  {
    "error": "No payments available"
  }
  ```

## Project Status

Project is: _Complete_

## Areas for improvement

- Store payments in a persistent database instead of memory for production use.
- Use a more secure way to handle sensitive information (e.g., encryption).
- Use HTTPS instead of HTTP.
- Add a check to see if currency code submitted by user is valid/exists.
- Implement authentication and authorization for the API.
- Paginate view all payments endpoint.
