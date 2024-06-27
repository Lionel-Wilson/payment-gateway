# Payment Gateway Documentation

## How to run

### Prerequisites

- Docker installed on your machine.
- Git installed on your machine.

### Step-by-Step Instructions

1. Clone the repository from https://github.com/Lionel-Wilson/payment-gateway
2. Open a terminal and cd into the 'Backend' folder
3. Run the following command to create a Docker image named payment-gateway:

```
docker build -t payment-gateway .
```

4. Then run the following command to start the container, map port 8080 on your host to port 8080 in the container, and name the container payment-gateway-container:

```
docker run --publish 8080:8080 payment-gateway
```

5. Test the payment gateway using Postman or your preferred API testing tool.

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
    "errors": ["CurrencyCode is required", "CVV must be exactly 3 characters"]
  }
  ```

### 2. Retrieve Payment Details

- **Endpoint**: `/payments`
- **Method**: `GET`
- **Description**: Retrieves the details of a previously made payment using its identifier..
- **Query Parameters**: Retrieves the details of a previously made payment using its identifier..
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

## Project Status

Project is: _Complete_

## Areas for improvement

- Store payments in a persistent database instead of memory for production use.
- Use a more secure way to handle sensitive information (e.g., encryption).
- Implement better error handling and logging.
- Use HTTPS instead of HTTP.
- Create a frontend UI.
- Add a check to see if currency code submitted by user is valid/exists.
- Add user friendly errors for ProccessPaymentRequest field validation.
- Implement authentication and authorization for the API.
