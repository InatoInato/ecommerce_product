package product

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service *ProductService
}

func (h *ProductHandler) GetAllProducts(c *gin.Context){
	products, err := h.Service.GetAllProducts()
	if err != nil{
		c.JSON(500, gin.H{
			"message": "Error fetching",
		})
		return
	}

	c.JSON(200, products)
}

func (h *ProductHandler) FilterProduct (c *gin.Context){
	var filter struct{
		Name string `json:"name"`
		Type string `json:"type"`
		MinPrice float64 `json:"min_price"`
		MaxPrice float64 `json:"max_price"`
		MinRating float64 `json:"min_rating"`
	}

	if err := c.ShouldBindJSON(&filter); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	products, err := h.Service.FilterProducts(filter.Name, filter.Type, filter.MinPrice, filter.MaxPrice, filter.MinRating)

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Filter error!",
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func AdminMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenString := c.GetHeader("Authorization")
		if tokenString == ""{
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token isn't available",
			})
			c.Abort()
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		
		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid{
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		role, ok := (*claims)["role"].(string)
		if !ok || role != "admin"{
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied",
			})
			c.Abort()
			return
		}

		c.Set("role", role)

		if role != "admin"{
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context){
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	if err := h.Service.CreateProduct(&product); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't save product",
		})
		return
	}
	
	c.JSON(http.StatusCreated, product)
}