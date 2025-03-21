# Vibe Coding Bank Account

This project implements a simple bank account system in Go, with functionality to create accounts, deposit money, and withdraw money. It also includes unit tests to validate the functionality and a web interface for managing accounts.

## Prerequisites

- Go 1.20 or later installed on your system.
- A terminal or command-line interface.

## How to Run the Code

1. Clone or navigate to the project directory (gosandbox/vibe-coding-bank-account):

2. Initialize the Go module (if not already done):
   ```bash
   go mod init vibe-coding-bank-account
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

   This will execute the `main.go` file and display the output in the terminal.

## How to Run the Web Application

1. Navigate to the project directory (gosandbox/vibe-coding-bank-account).

2. Run the web server:
   ```bash
   go run server.go
   ```

   The server will start on `http://localhost:8080`.

3. Use `curl` or a tool like Postman to interact with the API.

## Testing the Web Application with `curl`

Here are some example `curl` commands to test the web application:

1. **Create a new account**:
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"owner":"John Doe","initial_balance":1000,"limit":0}' http://localhost:8080/accounts
   ```

2. **View account details**:
   ```bash
   curl http://localhost:8080/accounts/John%20Doe
   ```

3. **Deposit money into an account**:
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"amount":500}' http://localhost:8080/accounts/John%20Doe/deposit
   ```

4. **Withdraw money from an account**:
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"amount":1200}' http://localhost:8080/accounts/John%20Doe/withdraw
   ```

5. **Create a special account with a negative limit**:
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"owner":"Jane Doe","initial_balance":1000,"limit":-500}' http://localhost:8080/accounts
   ```

6. **Withdraw money from a special account**:
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"amount":1400}' http://localhost:8080/accounts/Jane%20Doe/withdraw
   ```

## Project Structure

```
vibe-coding-bank-account/
├── bank/
│   ├── account.go        // Contains the Bank Account logic
│   └── account_test.go   // Contains unit tests for the Bank Account
├── go.mod                // Go module file
├── main.go               // Entry point for the application
└── server.go             // Web server for managing accounts
```

## Example Usage

When you run the application (main.go), it will create a bank account for "John Doe" with an initial balance of 1000.0 and display the account details.

For more business rules, such as deposit and withdrawal behavior, refer to the test files located in `bank/account_test.go`. These tests demonstrate how the account logic handles various scenarios, including deposits, withdrawals, and insufficient funds.

## License

This project is for educational purposes and does not include a specific license. The code was generated using the GPT-4o model and "Vibe Coding".