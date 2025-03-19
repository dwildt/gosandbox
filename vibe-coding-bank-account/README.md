# Vibe Coding Bank Account

This project implements a simple bank account system in Go, with functionality to create accounts, deposit money, and withdraw money. It also includes unit tests to validate the functionality.

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

## How to Run the Tests

1. Navigate to the project directory (/gosandbox/vibe-coding-bank-account):

2. Run all tests:
   ```bash
   go test ./...
   ```

3. For detailed test output, use the `-v` flag:
   ```bash
   go test -v ./...
   ```

   This will execute the tests in `bank/account_test.go` and display the results.

## Project Structure

```
vibe-coding-bank-account/
├── bank/
│   ├── account.go        // Contains the Bank Account logic
│   └── account_test.go   // Contains unit tests for the Bank Account
├── go.mod                // Go module file
└── main.go               // Entry point for the application
```

## Example Usage

When you run the application, it will create a bank account for "John Doe" with an initial balance of 1000.0 and display the account details.

For more business rules, such as deposit and withdrawal behavior, refer to the test files located in `bank/account_test.go`. These tests demonstrate how the account logic handles various scenarios, including deposits, withdrawals, and insufficient funds.  

## License

This project is for educational purposes and does not include a specific license. The code was generated using the GPT-4o model and "Vibe Coding".