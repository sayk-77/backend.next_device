package controllers

import (
	"fmt"
	"next_device/backend/models"
	"next_device/backend/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (pc *ProductController) CreateProduct(c *fiber.Ctx) error {
	var product *models.Products

	if err := c.BodyParser(&product); err != nil {
		fmt.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product data",
		})
	}

	productId, err := pc.productService.CreateProduct(product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"productId": productId,
	})
}

func (pc *ProductController) GetAllProducts(c *fiber.Ctx) error {
	products, err := pc.productService.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve products",
		})
	}

	return c.Status(fiber.StatusOK).JSON(products)
}

func (pc *ProductController) GetProductByIdOrName(c *fiber.Ctx) error {
	param := c.Params("param")

	id, err := strconv.Atoi(param)
	if err == nil {
		product, err := pc.productService.GetProductByID(uint(id))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Product not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(product)
	}

	product, err := pc.productService.GetProductByName(param)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func (pc *ProductController) UpdateProduct(c *fiber.Ctx) error {
	var product models.Products

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product data",
		})
	}

	if err := pc.productService.UpdateProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}

func (pc *ProductController) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	if err := pc.productService.DeleteProduct(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}

func (pc *ProductController) GetProductsByCategory(c *fiber.Ctx) error {
	category := c.Params("category")
	limit, err := strconv.Atoi(c.Query("limit", "30"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit parameter"})
	}
	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid offset parameter"})
	}

	products, err := pc.productService.GetProductsByCategory(category, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

func (pc *ProductController) GetDiscountedProducts(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.Query("limit", "30"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit parameter"})
	}
	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid offset parameter"})
	}

	brand := c.Query("brand")

	products, err := pc.productService.GetDiscountedProducts(limit, offset, brand)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

func (pc *ProductController) GetNewProducts(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.Query("limit", "30"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit parameter"})
	}
	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid offset parameter"})
	}

	products, err := pc.productService.GetNewProducts(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

func (pc *ProductController) GetProductsByBrandAndCategory(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.Query("limit", "30"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit parameter"})
	}
	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid offset parameter"})
	}
	brandId, err := strconv.Atoi(c.Query("brand_id"))
	fmt.Print(brandId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid brand ID",
		})
	}
	categoryName := c.Query("category")
	fmt.Print(categoryName)

	products, err := pc.productService.GetProductsByBrandAndCategory(brandId, categoryName, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

func (pc *ProductController) SearchProduct(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.Query("limit", "30"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit parameter"})
	}
	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid offset parameter"})
	}

	query := c.Query("query")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameter"})
	}

	result, err := pc.productService.SearchAll(query, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to search product"})
	}
	return c.JSON(result)
}

func (pc *ProductController) GetFilteredProducts(c *fiber.Ctx) error {
	category := c.Params("category")

	var filters struct {
		PriceFrom       *int     `json:"priceFrom"`
		PriceTo         *int     `json:"priceTo"`
		Brands          []string `json:"brands"`
		ScreenFrom      *float64 `json:"screenFrom"`
		ScreenTo        *float64 `json:"screenTo"`
		Memories        []string `json:"memories"`
		RAM             []string `json:"ram"`
		Ratings         []string `json:"ratings"`
		CameraQualities []string `json:"cameraQualities"`
		OS              []string `json:"os"`
	}

	if err := c.BodyParser(&filters); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	products, err := pc.productService.GetFilteredProducts(
		category,
		filters.PriceFrom,
		filters.PriceTo,
		filters.Brands,
		filters.ScreenFrom,
		filters.ScreenTo,
		filters.Memories,
		filters.RAM,
		filters.Ratings,
		filters.CameraQualities,
		filters.OS,
		10,
		0,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(products)
}

func (pc *ProductController) GetFilteredLaptops(c *fiber.Ctx) error {
	var filters struct {
		PriceFrom  *int     `json:"priceFrom"`
		PriceTo    *int     `json:"priceTo"`
		Brands     []string `json:"brands"`
		ScreenFrom *float64 `json:"screenFrom"`
		ScreenTo   *float64 `json:"screenTo"`
		Memories   []string `json:"memories"`
		RAM        []string `json:"ram"`
		CpuType    []string `json:"cpuType"`
		GpuType    []string `json:"gpuType"`
	}

	if err := c.BodyParser(&filters); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	products, err := pc.productService.GetFilteredLaptops(
		filters.PriceFrom,
		filters.PriceTo,
		filters.Brands,
		filters.ScreenFrom,
		filters.ScreenTo,
		filters.Memories,
		filters.RAM,
		filters.CpuType,
		filters.GpuType,
		10,
		0,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(products)
}

func (pc *ProductController) CreateImageProduct(c *fiber.Ctx) error {
	productId := c.FormValue("productId")
	if productId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "productId is required",
		})
	}

	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid productId format",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse multipart form",
		})
	}

	files := form.File["images"]

	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No images found in request",
		})
	}

	var imageRecords []models.ProductImage

	for i, file := range files {
		uniqueFileName := fmt.Sprintf("image_%d_%s", productIdInt, file.Filename)

		savePath := fmt.Sprintf("./images_product/%s", uniqueFileName)

		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save image",
			})
		}

		isMain := false
		if i == 0 {
			isMain = true
		}

		imageRecord := models.ProductImage{
			ProductID: uint(productIdInt),
			ImageURL:  uniqueFileName,
			IsMain:    isMain,
		}
		imageRecords = append(imageRecords, imageRecord)
	}

	if err := pc.productService.SaveProductImages(imageRecords); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save images to database",
		})
	}

	return c.Status(200).JSON("success")
}
