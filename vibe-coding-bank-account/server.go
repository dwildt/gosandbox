package main

import (
	"net/http"
	"vibe-coding-bank-account/bank"

	"github.com/gin-gonic/gin"
)

var accounts = make(map[string]*bank.Account)

func main() {
	r := gin.Default()

	// Route to create a new account
	r.POST("/accounts", func(c *gin.Context) {
		var req struct {
			Owner          string  `json:"owner"`
			InitialBalance float64 `json:"initial_balance"`
			Limit          float64 `json:"limit"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var account *bank.Account
		if req.Limit != 0 {
			account = bank.NewSpecialAccount(req.Owner, req.InitialBalance, req.Limit)
		} else {
			account = bank.NewAccount(req.Owner, req.InitialBalance)
		}
		accounts[req.Owner] = account
		c.JSON(http.StatusCreated, account)
	})

	// Route to get account details
	r.GET("/accounts/:owner", func(c *gin.Context) {
		owner := c.Param("owner")
		account, exists := accounts[owner]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
			return
		}
		c.JSON(http.StatusOK, account)
	})

	// Route to deposit money
	r.POST("/accounts/:owner/deposit", func(c *gin.Context) {
		owner := c.Param("owner")
		account, exists := accounts[owner]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
			return
		}

		var req struct {
			Amount float64 `json:"amount"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		account.Deposit(req.Amount)
		c.JSON(http.StatusOK, account)
	})

	// Route to withdraw money
	r.POST("/accounts/:owner/withdraw", func(c *gin.Context) {
		owner := c.Param("owner")
		account, exists := accounts[owner]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
			return
		}

		var req struct {
			Amount float64 `json:"amount"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		success := account.Withdraw(req.Amount)
		if !success {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient funds or limit exceeded"})
			return
		}
		c.JSON(http.StatusOK, account)
	})

	// Start the server
	r.Run(":8080") // Runs on http://localhost:8080
}
