# Payment Gateway Documentation

## How to run

### Prerequisites

- Docker installed on your machine.
- Git installed on your machine.

### Step-by-Step Instructions

1. Clone the repository from https://github.com/processout-hiring/payment-gateway-Lionel-Wilson
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
    "currency": "USD",
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
    "error": "Validation failed: <details>"
  }
  ```

### 2. Retrieve Payment Details

- **Endpoint**: `/payments/{id}`
- **Method**: `GET`
- **Description**: Retrieves the details of a previously made payment using its identifier..
- **Request Parameters**: Retrieves the details of a previously made payment using its identifier..
  - **Path Parameter**: id (string) - The unique identifier of the payment.

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
    "currency": "USD",
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

## Assumptions

- The card number validation is performed using the creditcard tag, assuming a basic credit card format check.
- The expiry date validation assumes a MM/YY format.
- The currency code must be a three-letter alphabetic code (e.g. USD, EUR).
- The CVV is assumed to be a 3-digit numeric string.
- Payments are stored in memory for simplicity.
- I have to create the "CKO bank simulator" myself.
- Only possible responses from the bank simulator are a successful payment or a decline due to insufficient funds.(Obviously there can be more reasons and responses)
- You want to be able to see the payment details for both approved and declined payments when requesting payment details using the id.

## Areas for improvement

- Store payments in a persistent database instead of memory for production use.
- Use a more secure way to handle sensitive information (e.g., encryption).
- Implement better error handling and logging.
- Use HTTPS instead of HTTP.
- Create a frontend UI.
- Add a check to see if currency code submitted by user is valid/exists.
- Add user friendly errors for ProccessPaymentRequest field validation.
- Implement authentication and authorization for the API.

## Cloud technologies I’d use and why

### Azure App Service:

Purpose: To host the payment gateway application.

Why: Azure App Service provides a fully managed platform for building, deploying, and scaling web apps. It supports Docker containers and offers features like auto-scaling, load balancing, and CI/CD integration, making it easier to deploy and manage containerized applications.

### Azure SQL Database:

Purpose: To store payment and user data securely.

Why: Azure SQL Database is a fully managed relational database service that provides high availability, security, and scalability. It supports automatic backups, point-in-time restore, and advanced threat protection, ensuring the payment data is secure.

### Azure Key Vault:

Purpose: To manage secrets, keys, and certificates.

Why: Azure Key Vault helps you safeguard cryptographic keys and secrets used by your application. It is essential for storing sensitive information such as database connection strings, API keys, and encryption keys securely.

### Azure Active Directory (Azure AD):

Purpose: To manage user authentication and authorization.

Why: Azure AD provides identity and access management capabilities. It helps secure your payment gateway by enabling features like single sign-on (SSO), multi-factor authentication (MFA), and role-based access control (RBAC).
Azure Security Center:

## Extra mile areas

- Added unit tests
- Created a Dockerfile to containerize the application for easier deployment.
