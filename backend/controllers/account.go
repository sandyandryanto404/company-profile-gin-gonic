/**
 * This file is part of the Sandy Andryanto Company Profile Website.
 *
 * @author     Sandy Andryanto <sandy.andryanto404@gmail.com>
 * @copyright  2024
 *
 * For the full copyright and license information,
 * please view the LICENSE.md file that was distributed
 * with this source code.
 */
 
package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

func ProfileDetail(c *gin.Context) {

	auth := c.MustGet("claims").(jwt.MapClaims)
	db := c.MustGet("db").(*gorm.DB)
	
	var user models.User
	if err := db.Where("id = ?", auth["id"]).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": user})
}

func ProfileUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}

func PasswordUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}

func ProfileUpload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}