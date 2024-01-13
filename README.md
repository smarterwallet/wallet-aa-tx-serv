# Wallet AA Transaction Server

In order for the front-end to display the details and status of transactions initiated by the wallet, we need a back-end program to record account abstraction transactions (UserOperation) and monitor the transaction status until the transaction succeeds or fails or is rolled back.

This service needs to provide an interface to query and save transactions externally.

## Features

- [ ] Record user's sending transactions, including the UserOperation details of the transactions
- [ ] Periodic updates the status of the user's sending transactions
- [x] Query by user address or tx hash

## Getting Started

To get started with this template, follow these simple steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/smarterwallet/wallet-aa-tx-serv.git
   ```

2. Navigate to the project directory:
   ```bash
   cd wallet-aa-tx-serv
   ```

3. Install dependencies (make sure you have Go installed on your system):
   ```bash
   go mod tidy
   ```

4. Copy the `.env.example` file to `.env` and update the variables to match your local setup for MySQL or PostgreSQL.

5. Run the application:
   ```bash
   go run main.go
   ```