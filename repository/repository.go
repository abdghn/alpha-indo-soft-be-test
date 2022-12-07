package repository

import (
	"context"
	"fmt"
	"github.com/abdghn/alpha-indo-soft-be-test/models"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Create(article *models.Article) (*models.Article, error)
	GetArticles(criteria map[string]interface{}, query string) ([]*models.Article, error)
}

type repository struct {
	DB    *gorm.DB
	redis *redis.Client
}

func NewRepository(DB *gorm.DB, redis *redis.Client) Repository {
	return &repository{DB: DB, redis: redis}
}

func (e *repository) Create(article *models.Article) (*models.Article, error) {
	err := e.DB.Save(&article).Error
	if err != nil {
		return nil, fmt.Errorf("failed insert data")
	}

	IDKEY := fmt.Sprintf("users_%d", article.ID)

	errRedis := e.setValueCache(IDKEY, &article)
	if errRedis != nil {
		return nil, fmt.Errorf("failed insert redis: %v", errRedis)
	}
	return article, nil
}

func (e *repository) GetArticles(criteria map[string]interface{}, query string) ([]*models.Article, error) {
	var articles []*models.Article
	model := e.DB.Where(criteria)
	if query != "" {
		model = model.Where("name LIKE ?", "%"+query+"%")
	}
	err := model.Find(&articles).Error
	if err != nil {
		return nil, fmt.Errorf("failed view all data")
	}
	return articles, nil
}

func (e *repository) setValueCache(key string, value interface{}) error {

	err := e.redis.Set(context.Background(), key, value, 0).Err()

	if err != nil {
		return err
	}

	return nil
}
