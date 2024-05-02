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
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticleList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}

func ArticleDetail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}

func ArticleCommentList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}

func ArticleCommentCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}