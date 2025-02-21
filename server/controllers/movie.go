package controllers

import (
	"github.com/gin-gonic/gin"
	"share_movie/server/config"
	"share_movie/server/models"
	"net/http"
	"strconv"
)

// ListMovies 获取影视资源列表
func ListMovies(c *gin.Context) {
	var movies []models.Movie
	result := config.DB.Preload("Links").Preload("Tags").Find(&movies)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取影视列表失败"})
		return
	}
	c.JSON(http.StatusOK, movies)
}

// GetMovie 获取单个影视资源详情
func GetMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	var movie models.Movie
	result := config.DB.Preload("Links").Preload("Tags").First(&movie, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "影视资源不存在"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// CreateMovie 创建新的影视资源
func CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	result := config.DB.Create(&movie)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建影视资源失败"})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

// UpdateMovie 更新影视资源信息
func UpdateMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	var movie models.Movie
	if err := config.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "影视资源不存在"})
		return
	}

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	result := config.DB.Save(&movie)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新影视资源失败"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// DeleteMovie 删除影视资源
func DeleteMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	result := config.DB.Delete(&models.Movie{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除影视资源失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "影视资源已删除"})
}