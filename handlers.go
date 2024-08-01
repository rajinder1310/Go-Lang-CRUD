package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    "json:id"
	Name  string "json: name"
	Email string "json: email"
}

func createUser(c *gin.Context) {
	fmt.Println("createUser Start")

	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := getDB()
	_, err := db.Exec("INSERT INTO users (name, email) values (?, ?)", user.Name, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
	fmt.Println("Handler.go End")
}

func getUser(c *gin.Context) {
	fmt.Println("getUser Start")
	idStr := c.Query("id")
	fmt.Println("ID before conversion is:", idStr)

	// Trim any leading/trailing whitespace
	idStr = strings.Trim(idStr, `"`)
	// Convert to integer
	id, idErr := strconv.Atoi(idStr)
	if idErr != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	db := getDB()

	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

func updateUser(c *gin.Context) {
	fmt.Println("Update User Start")
	id := c.Query("id")
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := getDB()
	_, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Updated Successfully", "data": user})
}

func deleteUser(c *gin.Context) {
	idStr := c.Query("id")
	id, idErr := strconv.Atoi(idStr)
	if idErr != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	db := getDB()
	var user User
	fmt.Println("IIDDDDDDD", id)
	err := db.QueryRow("SELECT id, name, email from users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	fmt.Println("DONE")
	
	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Delete Successfully"})
}
