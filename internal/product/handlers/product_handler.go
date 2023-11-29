package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hosnibounechada/go-api/internal/product/models"
	"github.com/hosnibounechada/go-api/internal/product/services"
	pkgUtil "github.com/hosnibounechada/go-api/pkg/util"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		productService: *services.NewProductService(),
	}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.productService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) Get(c *gin.Context) {
	productIdStr := c.Param("id")

	productId, err := strconv.ParseInt(productIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	product, err := h.productService.Get(productId)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var product models.CreateProductDTO

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": pkgUtil.FormatValidationErrors(err)})
		return
	}

	createdProduct, err := h.productService.Create(product)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductHandler) Update(c *gin.Context) {
	productIdStr := c.Param("id")

	productId, err := strconv.ParseInt(productIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	var product models.UpdateProductDTO

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": pkgUtil.FormatValidationErrors(err)})
		return
	}

	updatedProduct, err := h.productService.Update(productId, product)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	productIdStr := c.Param("id")

	productId, err := strconv.ParseInt(productIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	dbErr := h.productService.Delete(productId)

	if dbErr != nil {
		c.JSON(http.StatusGone, gin.H{
			"error": dbErr.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
