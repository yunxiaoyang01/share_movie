package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"share_movie/server/config"
	"share_movie/server/models"
	"strconv"
)

// GetUsers 获取用户列表
func GetUsers(c *gin.Context) {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户列表失败"})
		return
	}

	// 过滤敏感信息
	var userList []gin.H
	for _, user := range users {
		userList = append(userList, gin.H{
			"id":        user.ID,
			"username":  user.Username,
			"email":     user.Email,
			"nickname":  user.Nickname,
			"role":      user.Role,
			"status":    user.Status,
			"avatar":    user.Avatar,
			"lastLogin": user.LastLogin,
		})
	}

	c.JSON(http.StatusOK, userList)
}

// UpdateUserStatus 更新用户状态
func UpdateUserStatus(c *gin.Context) {
	// 获取用户ID
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	// 获取请求体
	var req struct {
		Status int `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 检查状态值是否有效
	if req.Status != 0 && req.Status != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的状态值"})
		return
	}

	// 查找用户
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	// 更新用户状态
	user.Status = req.Status
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户状态失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "用户状态更新成功",
		"user": gin.H{
			"id":     user.ID,
			"status": user.Status,
		},
	})
}