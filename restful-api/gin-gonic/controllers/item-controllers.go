package controllers

import (
	"strconv"

	"sinhnx.dev/restfulapi/gonic/models"

	"sinhnx.dev/restfulapi/gonic/dal"

	"github.com/gin-gonic/gin"
)

// SetupItemsRouter for item router
func SetupItemsRouter(r *gin.Engine) {
	var Items = []models.Item{
		{ItemId: 1, ItemName: "Item 1", UnitPrice: 12.5, Amount: 5, ItemStatus: 1, ItemDescription: "Item 1 Description"},
		{ItemId: 2, ItemName: "Item 2", UnitPrice: 15.5, Amount: 3, ItemStatus: 1, ItemDescription: "Item 2 Description"},
		{ItemId: 3, ItemName: "Item 3", UnitPrice: 18.5, Amount: 2, ItemStatus: 1, ItemDescription: "Item 3 Description"},
	}
	//Create
	r.POST("/items", func(c *gin.Context) {
		var item models.Item
		if err := c.ShouldBindJSON(&item); err == nil {
			Items = append(Items, item)
			c.JSON(200, gin.H{
				"messages": "Insert Item complete",
				"itemId":   item.ItemId,
			})
		} else {
			c.JSON(500, gin.H{
				"messages": "invalid item",
			})
		}
	})
	//Read
	r.GET("/items/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Invalid id",
			})
		} else {
			for _, item := range Items {
				if item.ItemId == id {
					c.JSON(200, item)
				}
			}
		}
	})
	//Update
	r.PUT("/items", func(c *gin.Context) {
		var item models.Item
		if err := c.ShouldBindJSON(&item); err == nil {
			for index, i := range Items {
				if item.ItemId == i.ItemId {
					Items[index] = item
					c.JSON(200, gin.H{
						"messages": "Update Item complete",
						"itemId":   item.ItemId,
					})
				}
			}
		} else {
			c.JSON(500, gin.H{
				"messages": "Invalid item",
			})
		}
	})
	//Delete
	r.DELETE("/items/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Invalid id",
			})
		} else {
			for index, item := range Items {
				// if our id path parameter matches one of our items
				if item.ItemId == id {
					// updates our Items array to remove the item
					Items = append(Items[:index], Items[index+1:]...)
					c.JSON(200, gin.H{
						"messages": "Update Item complete",
						"itemId":   item.ItemId,
					})
				}
			}
		}
	})
}

// SetupItemRouter for item router
func SetupItemRouter(r *gin.Engine) {
	//Create
	r.POST("/item", func(c *gin.Context) {
		var item models.Item
		if err := c.ShouldBindJSON(&item); err == nil {
			rowsAffected, lastInsertedId, err := dal.InsertItem(item)
			if err != nil {
				c.JSON(500, gin.H{
					"messages": "Insert Item error",
				})
			} else {
				if rowsAffected > 0 {
					c.JSON(200, gin.H{
						"messages": "Insert Item complete",
						"itemId":   lastInsertedId,
					})
				}
			}
		}
	})

	//Read
	r.GET("/item/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		item, err := dal.GetItem(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Item not found",
			})
		} else {
			c.JSON(200, item)
		}
	})

	//Update
	r.PUT("/item", func(c *gin.Context) {
		var item models.Item
		if err := c.ShouldBindJSON(&item); err == nil {
			rowsAffected, err := dal.UpdateItem(item)
			if err != nil {
				c.JSON(500, gin.H{
					"messages": "update Item error",
				})
			} else {
				if rowsAffected > 0 {
					c.JSON(200, gin.H{
						"messages": "update Item complete",
					})
				}
			}
		}
	})

	//Delete
	r.DELETE("/item/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		rowsDeletedAffected, err := dal.DeleteItem(id)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "delete error.",
			})
		} else {
			if rowsDeletedAffected > 0 {
				c.JSON(200, gin.H{
					"messages": "delete completed.",
				})
			}
		}
	})
}
