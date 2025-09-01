# Go Learning Sandbox (gosandbox)

**ALWAYS follow these instructions first and only fallback to additional search and context gathering if the information here is incomplete or found to be in error.**

This is a Go learning repository containing multiple chapters demonstrating different Go concepts, from basic programs to web applications with APIs and testing.

## Prerequisites and Setup

### Required Software
- Go 1.20 or later is required (repository currently has Go 1.24.6 installed)
- curl for testing web APIs
- Standard Unix tools (bash, find, etc.)

### Repository Structure
```
gosandbox/
├── ch01-hello/              # Basic Go programs (hello world variations)
├── ch03-tipos-de-dados/     # Data types with unit tests
├── ch04-arr-slice-map/      # Arrays, slices, maps with unit tests  
├── ch05-cast/               # Type conversion examples
├── ch14-fizzbuzz/           # FizzBuzz implementation with unit tests
├── ch15-webfizzbuzz/        # Web FizzBuzz using gin-gonic framework
└── vibe-coding-bank-account/ # Complete banking app with web API (has go.mod)
```

## Working Effectively

### Building and Testing Individual Chapters

**CRITICAL TIMING NOTES:**
- Initial Go module downloads: 5-15 seconds per module. NEVER CANCEL.
- Testing: Typically under 1 second per test suite once dependencies are installed
- Web application startup: Under 1 second
- Gin framework download: 10-15 seconds first time. NEVER CANCEL. Set timeout to 30+ minutes.

#### Chapters Without Tests (ch01-hello, ch05-cast)
```bash
cd ch01-hello
go run hello.go                    # Takes ~7 seconds first run, <1 second after
go run hello_string.go             # Takes <1 second  
go run hello_bhaskara_chatgpt.go   # Individual program execution
```

```bash  
cd ch05-cast
go run converte.go                 # Takes <1 second
```

#### Chapters With Unit Tests (ch03, ch04, ch14)
**REQUIRED SETUP:** Each testing chapter needs Go modules initialized and dependencies installed:

```bash
cd ch03-tipos-de-dados
go mod init ch03-tipos-de-dados
go get gotest.tools/assert         # Downloads dependencies - takes 5-10 seconds. NEVER CANCEL.
go test                           # Takes <1 second after setup
```

```bash
cd ch04-arr-slice-map  
go mod init ch04-arr-slice-map
go get gotest.tools/assert         # Takes 3-5 seconds after first install
go test                           # Takes <1 second
```

```bash
cd ch14-fizzbuzz
go mod init ch14-fizzbuzz  
go get gotest.tools/assert         # Takes 3-5 seconds after first install
go test                           # Takes <1 second
```

#### Web Applications (ch15-webfizzbuzz)
**SETUP:** Requires gin-gonic framework:

```bash
cd ch15-webfizzbuzz
go mod init ch15-webfizzbuzz
go get github.com/gin-gonic/gin    # Takes 10-15 seconds first time. NEVER CANCEL.
go get gotest.tools/assert         # For tests
go test                           # Takes <1 second
go run fizzbuzzweb.go gobuzz.go    # IMPORTANT: Must include BOTH files
```

**Testing the Web API:**
```bash
curl http://localhost:8080/              # Returns usage instructions
curl http://localhost:8080/fb/5          # Returns {"result":"buzz","value":5}
curl http://localhost:8080/fb/15         # Returns {"result":"fizzbuzz","value":15}
```

#### Complete Banking Application (vibe-coding-bank-account)
**IMPORTANT:** This directory already has go.mod configured with all dependencies.

```bash
cd vibe-coding-bank-account
go test ./bank                    # Test only bank package - takes <1 second
go run main.go                    # Runs CLI demonstration - takes <1 second
go run server.go                  # Starts web API on :8080
```

**Banking Web API Testing:**
```bash
# Create account
curl -X POST -H "Content-Type: application/json" \
  -d '{"owner":"John Doe","initial_balance":1000,"limit":0}' \
  http://localhost:8080/accounts

# View account  
curl http://localhost:8080/accounts/John%20Doe

# Deposit money
curl -X POST -H "Content-Type: application/json" \
  -d '{"amount":500}' \
  http://localhost:8080/accounts/John%20Doe/deposit

# Withdraw money
curl -X POST -H "Content-Type: application/json" \
  -d '{"amount":1200}' \
  http://localhost:8080/accounts/John%20Doe/withdraw
```

### Code Quality and Validation

**Before making any changes, ALWAYS run:**

```bash
gofmt -l .                        # Check formatting (returns empty if properly formatted)
go vet ./...                      # Static analysis (may fail on vibe-coding-bank-account due to multiple main functions)
~/go/bin/golint .                 # Style checking (install first with: go install golang.org/x/lint/golint@latest)
```

**Format code automatically:**
```bash
gofmt -w filename.go              # Format specific file
```

### Validation Scenarios

**ALWAYS test these scenarios after making changes:**

1. **Basic Program Execution:**
   - Run `go run hello.go` in ch01-hello
   - Verify output is "hello, world"

2. **Unit Test Validation:**
   - Run tests in ch03, ch04, ch14 directories
   - All tests should pass

3. **Web Application Validation:**
   - Start ch15-webfizzbuzz server
   - Test fizzbuzz API endpoints (3→fizz, 5→buzz, 15→fizzbuzz)
   - Start bank server  
   - Test complete banking workflow (create account, deposit, withdraw)

4. **Code Quality:**
   - Run gofmt to ensure code is properly formatted
   - Run go vet (excluding vibe-coding-bank-account main conflict)

## Common Issues and Solutions

### Multiple Main Functions Error
In vibe-coding-bank-account:
- `go test ./...` fails due to main.go and server.go both having main functions
- **Solution:** Test bank package specifically: `go test ./bank`
- **Solution:** Run programs individually: `go run main.go` OR `go run server.go`

### Missing Dependencies
If you see import errors:
- **ch03, ch04, ch14:** Run `go get gotest.tools/assert`
- **ch15, vibe-coding-bank-account:** Run `go get github.com/gin-gonic/gin`

### Module Initialization Required
For chapters without go.mod:
- Run `go mod init [chapter-name]` first
- Then install dependencies as shown above

### Gin Framework First Download
- Takes 10-15 seconds to download all gin dependencies
- **NEVER CANCEL** - use 30+ minute timeout for first-time gin installation
- Subsequent runs are fast (<1 second)

## Testing and Development Workflow

1. **For simple Go programs:** Just use `go run filename.go`
2. **For chapters with tests:** Initialize module, install dependencies, then test
3. **For web applications:** Initialize module, install gin, run with all required files
4. **Always validate:** Format with gofmt, check with go vet, test APIs with curl

## Key Dependencies

- **gotest.tools/assert** - For unit testing (ch03, ch04, ch14)
- **github.com/gin-gonic/gin** - For web applications (ch15, vibe-coding-bank-account)

## Time Expectations

- **First-time dependency download:** 5-15 seconds per module. NEVER CANCEL.
- **Gin framework download:** 10-15 seconds first time. NEVER CANCEL. Set timeout to 30+ minutes.
- **Test execution:** <1 second once dependencies installed
- **Program execution:** <1 second for most programs, 5-7 seconds for first Go run
- **Web server startup:** <1 second

**CRITICAL:** Always wait for dependency downloads to complete. Build cancellation will require re-downloading dependencies.