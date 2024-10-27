package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wanchanok6698/web-blogs/api/v1/models"
	"github.com/wanchanok6698/web-blogs/api/v1/services"
	"github.com/wanchanok6698/web-blogs/util"
)

type BlogsController struct {
	service services.BlogService
}

func NewBlogsController(service services.BlogService) *BlogsController {
	return &BlogsController{service: service}
}

func (bc *BlogsController) GetAllBlogs(c *fiber.Ctx) error {
	authorId := c.Query("authorId")
	search := c.Query("search")
	sort := c.Query("sort")

	filterOptions := models.FilterBlogsOptions{
		AuthorID: authorId,
		Search:   search,
		Sort:     sort,
	}
	resultChan := make(chan []models.BlogPost)
	errorChan := make(chan error)

	go func() {
		result, err := bc.service.GetAllBlogs(c.Context(), filterOptions)
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- result
	}()

	select {
	case blogs := <-resultChan:
		return util.HandleSuccess(c, "ดึงข้อมูลสำเร็จ", blogs)
	case err := <-errorChan:
		return util.HandleError(c, "ดึงข้อมูล blogs ไม่สำเร็จ", err.Error(), fiber.StatusInternalServerError)
	}

}

func (bc *BlogsController) GetBlogByID(c *fiber.Ctx) error {
	id := c.Params("id")
	resultChan := make(chan *models.BlogPost)
	errorChan := make(chan error)

	go func() {
		result, err := bc.service.GetBlogByID(c.Context(), id)
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- result
	}()

	select {
	case blog := <-resultChan:
		return util.HandleSuccess(c, "ดึงข้อมูลสำเร็จ", blog)
	case err := <-errorChan:
		return util.HandleError(c, "ดึงข้อมูล blog ไม่สำเร็จ", err.Error(), fiber.StatusInternalServerError)
	}
}

func (bc *BlogsController) CreateBlog(c *fiber.Ctx) error {
	var blog models.BlogPost
	if err := c.BodyParser(&blog); err != nil {
		return util.HandleError(c, "ไม่สามารถพาร์สข้อมูลบล็อกได้", err.Error(), fiber.StatusBadRequest)
	}
	resultChan := make(chan *models.BlogPost)
	errorChan := make(chan error)

	go func() {
		result, err := bc.service.CreateBlog(c.Context(), &blog)
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- result
	}()

	select {
	case newBlog := <-resultChan:
		return util.HandleSuccess(c, "สร้างบล็อกสำเร็จ", newBlog)
	case err := <-errorChan:
		return util.HandleError(c, "สร้างบล็อกไม่สำเร็จ", err.Error(), fiber.StatusInternalServerError)
	}

}

func (bc *BlogsController) UpdateBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedBlog models.BlogUpdate
	if err := c.BodyParser(&updatedBlog); err != nil {
		return util.HandleError(c, "ไม่สามารถพาร์สข้อมูลบล็อกได้", err.Error(), fiber.StatusBadRequest)
	}

	resultChan := make(chan *models.BlogUpdate)
	errorChan := make(chan error)

	go func() {
		result, err := bc.service.UpdateBlog(c.Context(), id, &updatedBlog)
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- result
	}()

	select {
	case blog := <-resultChan:
		return util.HandleSuccess(c, "อัปเดตบล็อกสำเร็จ", blog)
	case err := <-errorChan:
		return util.HandleError(c, "อัปเดตบล็อกไม่สำเร็จ", err.Error(), fiber.StatusInternalServerError)
	}
}

func (bc *BlogsController) DeleteBlog(c *fiber.Ctx) error {
	id := c.Params("id")

	err := bc.service.DeleteBlog(c.Context(), id)
	if err != nil {
		return util.HandleError(c, "ลบบล็อกไม่สำเร็จ", err.Error(), fiber.StatusInternalServerError)
	}

	return util.HandleSuccess(c, "ลบบล็อกสำเร็จ", map[string]interface{}{})

}
