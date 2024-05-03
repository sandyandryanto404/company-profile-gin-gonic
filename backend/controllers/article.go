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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ArticleResultDetail struct {
	Id          uint64    `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Image       string    `json:"image"`
	Gender      string    `json:"gender"`
	AboutMe     string    `json:"about_me"`
	Categories  string    `json:"categories"`
}

func ArticleList(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var articles []ArticleResult
	db.Raw(`
		SELECT 
			a.id,
			a.title,
			a.slug,
			a.description,
			a.created_at,
			u.first_name,
			u.last_name,
			u.image,
			u.gender,
			u.about_me,
			(SELECT 
				GROUP_CONCAT(r.name SEPARATOR ',') AS r 
				FROM references_contents r
				WHERE r.id IN (
					SELECT reference_id
					FROM articles_references
					WHERE article_id = a.id
				)
			) as categories
		FROM articles a
		INNER JOIN users u ON u.id = a.user_id
		WHERE a.status = 1 ORDER BY RAND() LIMIT 3 
	`).Scan(&articles)

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
