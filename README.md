# Money Transfer API

This repository contains a Money Transfer API built with Go, utilizing PostgreSQL for data storage. This README outlines the steps to set up the project, run the database migrations, and execute sample requests using `curl`.

## Prerequisites

- Go (version 1.20 or later)
- PostgreSQL
- MockAPI account (for account validation. For testing, use account inside .env.example)

## Setup Instructions

### 1. Initialize Go Modules

Navigate to the project directory and run:

```bash
go mod init github.com/okyanawang/money-transfer-go
```

### 2. Tidy Up Dependencies

Ensure that all dependencies are correctly specified and download the required packages:

```bash
go mod tidy
```

### 3. Configure Environment Variables

Create a `.env` file in the project root and fill it with your PostgreSQL connection information. The file should look like this:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=<your_pg_user>
DB_PASSWORD=<your_pg_password>
DB_NAME=<your_db_name>
```

### 4. Run Database Seeding

Make sure to run the seed script to populate the database with sample accounts (Dena and Cynthia). This usually requires a command similar to:

```bash
go run cmd/seed/main.go
```

Adjust the command according to your project structure.

### 5. Validate Accounts using CURL

Run the following command to validate the account for Dena:

```bash
curl -X POST http://localhost:8088/api/v1/accounts/validate -H "Content-Type: application/json" -d "{\"account_number\": \"721539120\"}"
```

You should receive a response indicating whether the account is valid.

### 6. Transfer Funds using CURL

Execute the following command to initiate a transfer from Dena to Cynthia:

```bash
curl -X POST http://localhost:8088/api/v1/transfers -H "Content-Type: application/json" -d "{\"sender_account_id\": \"SENDER_ACCOUNT_ID\", \"receiver_account_id\": \"RECEIVER_ACCOUNT_ID\", \"amount\": 100000, \"transaction_date\": \"2024-10-07T14:56:36.899819+07:00\"}"
```

Make sure to replace "SENDER_ACCOUNT_ID" & "RECEIVER_ACCOUNT_ID" with the actual transaction ID that is generated when the transfer occurs.

### 7. Send Callback using CURL

Finally, send a callback indicating the status of the transaction:

```bash
curl -X POST http://localhost:8088/api/v1/transfers-callback -H "Content-Type: application/json" -d "{\"transaction_id\": \"YOUR_TRANSACTION_ID\", \"status\": \"success\", \"amount\": 100000, \"processed_at\": \"2024-10-07T14:57:00Z\", \"received_at\": \"2024-10-07T14:58:00Z\"}"
```

Make sure to replace "YOUR_TRANSACTION_ID" with the actual transaction ID that is generated when the transfer occurs.

## Future Development

Future development will implement database transactions and concurrency for improved reliability and performance.
