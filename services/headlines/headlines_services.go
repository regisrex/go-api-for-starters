package headlines_services

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitub.com/regisrex/golang-apis/helpers"
	"gitub.com/regisrex/golang-apis/models"
)

func CreateHeadline(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	var newHeadline struct {
		Title       string
		Quote       string
		Description string
		Body        string
	}
	c.Bind(&newHeadline)
	// validationError := helpers.Validate.Struct(newHeadline)
	// if validationError != nil {
	// 	c.JSON(406, gin.H{
	// 		"message": "Invalid information given",
	// 		"success": false,
	// 		"status":  406,
	// 	})
	// 	return
	// }
	helpers.Database.Create(&models.NewsHeadline{
		ID:          uuid.New().String(),
		Title:       newHeadline.Title,
		Quote:       newHeadline.Quote,
		Description: newHeadline.Description,
		UserRefer:   user_id.(string),
		Body:        newHeadline.Body,
	})

	c.JSON(200, gin.H{
		"message": "Headline created successfully",
		"success": false,
		"status":  200,
	})
}

func UpdateHeadline(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	headlineId := c.Param("id")

	var headline models.NewsHeadline
	result := helpers.Database.Where("id = ?", headlineId).First(&headline)
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"message": "Headline not found",
			"success": false,
			"status":  404,
		})
		c.Abort()
	}
	if headline.UserRefer != user_id {
		c.JSON(405, gin.H{
			"message": "You can't update a headline you don't own",
			"success": false,
			"status":  405,
		})
		c.Abort()
	}
	var body struct {
		Title       string
		Quote       string
		Description string
		Body        string
	}
	c.Bind(&body)
	// validationError := helpers.Validate.Struct(body)
	// if validationError != nil {
	// 	c.JSON(406, gin.H{
	// 		"message": "Invalid information given",
	// 		"success": false,
	// 		"status":  406,
	// 	})
	// 	return
	// }
	if body.Body != "" {
		headline.Body = body.Body
	}
	if body.Description != "" {
		headline.Description = body.Description
	}
	if body.Quote != "" {
		headline.Quote = body.Quote
	}
	if body.Title != "" {
		headline.Title = body.Title
	}

	helpers.Database.Save(&headline)
	c.JSON(200, gin.H{
		"message": "Headline updated successfully",
		"success": true,
		"status":  200,
	})
	return
}

func GetAllHeadlines(c *gin.Context) {
	var headlines []models.NewsHeadline
	helpers.Database.Find(&headlines)

	c.JSON(200, gin.H{
		"message": "Retrieved all headlines",
		"status":  200,
		"success": true,
		"data": map[string]interface{}{
			"headlines": headlines,
		},
	})

}

func GetSingleHeadline(c *gin.Context) {
	var headlineId string = c.Param("id")
	var headline models.NewsHeadline
	helpers.Database.Where("id = ?", headlineId).First(&headline)

	c.JSON(200, gin.H{
		"message": "Retrieved all headlines",
		"status":  200,
		"success": true,
		"data": map[string]interface{}{
			"headline": headline,
		},
	})

}

func DeleteHeadline(c *gin.Context) {
	id := c.Param("id")
	user_id, _ := c.Get("user_id")

	var headline models.NewsHeadline
	result := helpers.Database.Where("id = ?", id).First(&headline)
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"message": "Headline not found",
			"success": false,
			"status":  404,
		})
		c.Abort()
		return
	}
	if headline.UserRefer != user_id {
		c.JSON(405, gin.H{
			"message": "You can't delete a headline you don't own",
			"success": false,
			"status":  405,
		})
		c.Abort()
		return
	}
	helpers.Database.Where(&models.NewsHeadline{ID: headline.ID}).Delete(&models.NewsHeadline{})
	c.JSON(200, gin.H{
		"message": "Headline deleted",
		"success": true,
		"status":  200,
	})
	c.Abort()
}
