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

func PagePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}

func PageHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}

func PageAbout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}

func PageService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}

func PageFaq(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}

func PageContact(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}

func PageMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}

func PageSubscribe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}

func PageGetFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "status": true, "data": nil})
}