package handler

import (
	"github.com/abdghn/alpha-indo-soft-be-test/models"
	"github.com/abdghn/alpha-indo-soft-be-test/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func GetArticles(c *gin.Context) {
//	var articles []models.Article
//	db := database.GetDB()
//	db.Find(&articles)
//
//	helper.RespondJSON(c, 200, "", articles, nil)
//}
//
//func CreateArticle(c *gin.Context) {
//	var article models.Article
//
//	if err := c.ShouldBind(&article); err != nil {
//		helper.RespondJSON(c, 422, "Could not create article", nil, err)
//		return
//	}
//
//	db := database.GetDB()
//
//	db.Create(&article)
//
//	if article.ID != 0 {
//		helper.RespondJSON(c, 201, "Article has been created successfully", article, nil)
//	} else {
//		helper.RespondJSON(c, 409, "Can not create article", nil, nil)
//	}
//}

type Handler interface {
	CreateArticle(c *gin.Context)
	ViewArticles(c *gin.Context)
}

type handler struct {
	r repository.Repository
}

func NewHandler(r repository.Repository) Handler {
	return &handler{r: r}
}

func (h *handler) CreateArticle(c *gin.Context) {
	var request models.CreateArticle

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := &models.Article{Body: request.Body, Title: request.Title, Author: request.Author}

	result, err := h.r.Create(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(200, result)

}

func (h *handler) ViewArticles(c *gin.Context) {
	var req models.GetArticle
	var err error
	criteria := make(map[string]interface{})

	err = c.ShouldBindUri(&req)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Oopss server someting wrong"})
		return
	}

	if req.Author != "" {
		criteria["author"] = req.Author
	}

	articles, err := h.r.GetArticles(criteria, req.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Oopss server someting wrong"})
		return
	}

	c.JSON(http.StatusOK, articles)

}
