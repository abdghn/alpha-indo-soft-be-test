package handler

import (
	"bytes"
	"encoding/json"
	"github.com/abdghn/alpha-indo-soft-be-test/models"
	"github.com/abdghn/alpha-indo-soft-be-test/repository/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func NewServer(repo *mocks.MockRepository) *gin.Engine {
	h := NewHandler(repo)
	server := gin.Default()
	server.GET("/articles", h.ViewArticles)
	server.POST("/articles", h.CreateArticle)
	return server
}

func requireBodyMatchArticle(t *testing.T, body *bytes.Buffer, article models.Article) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotArticle models.Article
	err = json.Unmarshal(data, &gotArticle)
	require.NoError(t, err)
	require.Equal(t, article, gotArticle)
}

func requireBodyMatchArticles(t *testing.T, body *bytes.Buffer, articles []*models.Article) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotArticles []*models.Article
	err = json.Unmarshal(data, &gotArticles)
	require.NoError(t, err)
	require.Equal(t, articles, gotArticles)
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
