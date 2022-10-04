package handler

import (
	"assignment2/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ordService struct {
	ord []*model.Orders
	ite []*model.Items
}

type OrderIface interface {
	CreateOrder(c *gin.Context)
	GetOrders(c *gin.Context)
	// GetOrderID(c *gin.Context) int
	UpdateOrder(c *gin.Context)
	DeleteOrder(c *gin.Context)
}

func NewOrderService(ord []*model.Orders, ite []*model.Items) OrderIface {
	return &ordService{
		ord: ord,
		ite: ite,
	}
}

// Create One Order
// @Summary Create New Order
// @Description Create New Order
// @Tags Orders
// @Accept  */*
// @Produce  json
// @Param data body Request true "Order"
// @Success 200 {object} Response
// @Failure 500 {object} string "error"
// @Router /orders [post]
func (s *ordService) CreateOrder(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		var request model.Request
		var response model.Response
		var item []model.Items

		// Decode Request
		err := json.NewDecoder(c.Request.Body).Decode(&request)
		if err != nil {
			fmt.Println(err.Error())
		}

		orderID := s.GetOrderID(c)

		// Insert Order
		s.ord = append(s.ord, &model.Orders{
			OrderID:      orderID,
			CustomerName: request.CustomerName,
			OrderedAt:    time.Now(),
		})

		// Create Response
		for i := range request.Item {
			itemID := s.GetItemID(c)

			// Insert Item
			s.ite = append(s.ite, &model.Items{
				ItemID:      itemID,
				ItemCode:    request.Item[i].ItemCode,
				Description: request.Item[i].Description,
				Quantity:    request.Item[i].Quantity,
				OrderId:     orderID,
			})

			item = append(item, model.Items{
				ItemID:      itemID,
				ItemCode:    request.Item[i].ItemCode,
				Description: request.Item[i].Description,
				Quantity:    request.Item[i].Quantity,
				OrderId:     orderID,
			})
		}

		response.CustomerName = request.CustomerName
		response.OrderID = orderID
		response.OrderedAt = s.ord[0].OrderedAt
		response.Item = item

		c.JSON(http.StatusCreated, response)
	}
}

// Get Order
// @Summary Get All Order
// @Description Get All Order
// @Tags Orders
// @Accept  */*
// @Produce  json
// @Success 200 {object} Response
// @Failure 500 {object} string "error"
// @Router /orders [get]
func (s *ordService) GetOrders(c *gin.Context) {
	var response []model.Response

	for _, v := range s.ord {
		response = append(response, model.Response{
			CustomerName: v.CustomerName,
			OrderID:      v.OrderID,
			OrderedAt:    v.OrderedAt,
		})
	}

	for i := 0; i < len(response); i++ {
		for j := 0; j < len(s.ite); j++ {
			if s.ite[j].OrderId == response[i].OrderID {
				response[i].Item = append(response[i].Item, model.Items{
					ItemID:      s.ite[j].ItemID,
					ItemCode:    s.ite[j].ItemCode,
					Description: s.ite[j].Description,
					Quantity:    s.ite[j].Quantity,
					OrderId:     s.ite[j].OrderId,
				})
			}
		}
	}

	c.JSON(http.StatusOK, response)
}

// Update Order
// @Summary Update Order
// @Description Update Order
// @Tags Orders
// @Accept  */*
// @Produce  json
// @Success 200 {object} Response
// @Failure 500 {object} string "error"
// @Router /order/:id [put]
func (s *ordService) UpdateOrder(c *gin.Context) {
	var request model.Request
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	// Get id order
	id := c.Params.ByName("id")
	orderId, _ := strconv.Atoi(id)

	for _, v := range s.ord {
		if v.OrderID == orderId {
			v.CustomerName = request.CustomerName
		}
	}

	for i, v := range s.ite {
		if v.OrderId == orderId {
			if v.ItemID == orderId {
				v.ItemCode = request.Item[i].ItemCode
				v.Description = request.Item[i].Description
				v.Quantity = request.Item[i].Quantity
			}
		}
	}

	c.JSON(http.StatusOK, request)
}

// Delete Order
// @Summary Delete Order
// @Description Delete Order
// @Tags Orders
// @Accept  */*
// @Produce  json
// @Success 200 {object} Response
// @Failure 500 {object} string "error"
// @Router /order/:id [delete]
func (s *ordService) DeleteOrder(c *gin.Context) {
	// Get id order
	id := c.Params.ByName("id")
	orderId, _ := strconv.Atoi(id)

	for i := range s.ord {
		if orderId == s.ord[i].OrderID {
			s.ord[i] = &model.Orders{}
		}
	}

	s.ord = DelIndex(s.ord, 0)
	c.JSON(http.StatusOK, s.ord)
}

func (s *ordService) GetOrderID(c *gin.Context) int {
	max := 0
	for _, v := range s.ord {
		if v.OrderID > max {
			max = v.OrderID
		}
	}

	return max + 1
}

func (s *ordService) GetItemID(c *gin.Context) int {
	max := 0
	for _, v := range s.ite {
		if v.ItemID > max {
			max = v.ItemID
		}
	}
	return max + 1
}

func DelIndex(s []*model.Orders, index int) []*model.Orders {
	return append(s[:index], s[index+1:]...)
}
