package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Customer struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Cpf      string  `json:"cpf"`
	Income   float32 `json:"income"`
	Location string  `json:"location"`
}

type CustomerResponse struct {
	Name  string `json:"name"`
	Loans []Loan `json:"loans"`
}

type Loan struct {
	Type         string `json:"type"`
	InterestRate int    `json:"interest_rate"`
}

var (
	personal    Loan = Loan{Type: "PERSONAL", InterestRate: 4}
	guaranteed  Loan = Loan{Type: "GUARANTEED", InterestRate: 3}
	consignment Loan = Loan{Type: "CONSIGNMENT", InterestRate: 2}
)

func main() {
	router := gin.Default()

	router.POST("/customer-loans", postLoans)

	router.Run("localhost:8080")
}

func postLoans(c *gin.Context) {
	var newCustumer Customer

	// verifica se o body está correto
	if err := c.BindJSON(&newCustumer); err != nil {
		return
	}

	loans := getAvailableLoans(newCustumer)

	c.IndentedJSON(http.StatusOK, CustomerResponse{Name: newCustumer.Name, Loans: loans})
}

func getAvailableLoans(c Customer) []Loan {
	var loans []Loan

	// Personal loan
	if c.Income <= 3000 {
		loans = append(loans, personal)
	} else if c.Income > 3000 && c.Income <= 5000 && c.Age < 30 && c.Location == "SP" {
		loans = append(loans, personal)
	}

	// Consignment loan
	if c.Income >= 5000 {
		loans = append(loans, consignment)
	}

	// Guaranteed loan
	if c.Income <= 3000 {
		loans = append(loans, guaranteed)
	} else if c.Income > 3000 && c.Income <= 5000 && c.Age < 30 && c.Location == "SP" {
		loans = append(loans, guaranteed)
	}

	return loans
}
