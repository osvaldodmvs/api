package controllers

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/osvaldodmvs/api/initializers"
	"github.com/osvaldodmvs/api/models"
	"github.com/osvaldodmvs/api/utils"
	"gorm.io/gorm"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.Bind(&product); err != nil {
		//if the json is not valid, return a bad request
		log.Println("Error binding JSON: ", err)
		c.JSON(400, gin.H{"message": "Error binding JSON"})
		return
	}

	valid := utils.IsValidCategory(product.Category)

	if !valid {
		c.JSON(400, gin.H{"message": "Invalid category"})
		return
	}

	//static testing
	//post := models.Product{Name: "Laptop", Description: "A laptop.", Price: 1000, Stock: 10, Rating: 5}
	post := models.Product{Name: product.Name, Description: product.Description, Price: product.Price, Stock: product.Stock, Rating: product.Rating, Category: product.Category}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		//if there is an error, print it and return a bad request
		log.Println("Error creating product: ", result.Error)
		c.JSON(400, gin.H{"message": "Error creating product"})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetProducts(c *gin.Context) {

	var products []models.Product

	result := initializers.DB.Find(&products)

	if result.Error != nil {
		//if there is an error, print it and return a bad request
		log.Println("Error finding products: ", result.Error)
		c.JSON(400, gin.H{"message": "Error finding products"})
		return
	}

	c.JSON(200, gin.H{
		"products": products,
	})
}

func GetProductById(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	result := initializers.DB.First(&product, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			//id not found
			c.JSON(404, gin.H{"message": "Resource not found"})
		} else {
			//400 for other errors
			c.JSON(400, gin.H{"message": "Bad request"})
		}
		return
	}

	c.JSON(200, gin.H{
		"product {" + id + "}": product,
	})
}

func UpdateProductById(c *gin.Context) {
	//get the product to be updated with the id
	//first thing to do because if it doesn't exist, we don't need to continue
	var product models.Product

	id := c.Param("id")

	result := initializers.DB.First(&product, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			//id not found
			c.JSON(404, gin.H{"message": "Resource not found"})
		} else {
			//400 for other errors
			c.JSON(400, gin.H{"message": "Bad request"})
		}
		return
	}
	//requested product with the filled in data to update the actual existing product
	var reqProduct models.Product

	//bind the json to the product struct
	if err := c.Bind(&reqProduct); err != nil {
		//if the json is not valid, return a bad request
		log.Println("Error binding JSON: ", err)
		c.JSON(400, gin.H{"message": "Error binding JSON"})
		return
	}

	initializers.DB.Model(&product).Updates(models.Product{
		Name: reqProduct.Name, Description: reqProduct.Description, Price: reqProduct.Price, Stock: reqProduct.Stock, Rating: reqProduct.Rating})

	c.JSON(200, gin.H{
		"product {" + id + "} updated": product,
	})
}

func DeleteProductById(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&models.Product{}, id)

	if result.RowsAffected == 0 {
		//id doesn't exist, can't delete anything
		c.JSON(404, gin.H{"message": "Resource not found"})
		return
	}

	c.JSON(200, gin.H{
		"product {" + id + "} deleted": "true",
	})
}
