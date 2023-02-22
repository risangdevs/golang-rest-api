package controllers

import (
	"golang-rest-api/db"
	"golang-rest-api/repositories"
	"golang-rest-api/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTransaction(c *gin.Context) {
	var (
		result gin.H
	)

	transactions, err := repositories.GetAllTransaction(db.DbConnection)
	if err != nil {
		result = gin.H{"result": err}
	} else {
		result = gin.H{"result": transactions}
	}
	c.JSON(http.StatusOK, result)

}

func InsertTransaction(c *gin.Context) {
	var transaction structs.Transaction

	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		panic(err)
	}
	err2 := repositories.InsertTransaction(db.DbConnection, transaction)
	if err2 != nil {
		panic(err2)
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Insert Transaction"})
}

func UpdateTransaction(c *gin.Context) {
	var transaction structs.Transaction
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		panic(err)
	}
	transaction.ID = int64(id)
	err = repositories.UpdateTransaction(db.DbConnection, transaction)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success Update Transaction"})

}

func DeleteTransaction(c *gin.Context) {
	var transaction structs.Transaction
	id, err := strconv.Atoi(c.Param("id"))

	transaction.ID = int64(id)

	err = repositories.DeleteTransaction(db.DbConnection, transaction)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success Delete Transaction"})
}
