package main

import (
	"github.com/iamcmnut/go-simple-blockchain/blockchain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	chain *blockchain.Blockchain
)

func main() {
	chain = &blockchain.Blockchain{}

	chain.Init()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", getChain)
	e.POST("/", createBlock)

	e.Logger.Fatal(e.Start(":1234"))
}

func getChain(c echo.Context) error {
	code := 200
	return c.JSON(code, &ResponseChains{
		Status:  "error",
		Code:    code,
		Message: "All blocks",
		Chains:  chain.Blocks,
	})
}

func createBlock(c echo.Context) error {
	req := &blockchain.Block{}
	if err := c.Bind(req); err != nil {
		var code = 400
		return c.JSON(code, &ResponseError{
			Status:  "error",
			Code:    code,
			Message: "Cannot bind data",
		})
	}

	if req.Data == "" {
		var code = 400
		return c.JSON(code, &ResponseError{
			Status:  "error",
			Code:    code,
			Message: "data is required",
		})
	}

	block, err := chain.CreateBlock(req.Data)
	if err != nil {
		var code = 500
		return c.JSON(code, &ResponseError{
			Status:  "error",
			Code:    code,
			Message: err.Error(),
		})
	}

	code := 200
	return c.JSON(code, &ResponseBlock{
		Status:  "success",
		Code:    code,
		Message: "New block has been created",
		Block:   block,
	})
}

type ResponseError struct {
	Status  string `json:"status" example:"error"`
	Code    int    `json:"code" example:"520"`
	Message string `json:"message" example:"error message"`
}

type ResponseChains struct {
	Status  string              `json:"status" example:"error"`
	Code    int                 `json:"code" example:"520"`
	Message string              `json:"message" example:"error message"`
	Chains  []*blockchain.Block `json:"chains"`
}

type ResponseBlock struct {
	Status  string            `json:"status" example:"error"`
	Code    int               `json:"code" example:"520"`
	Message string            `json:"message" example:"error message"`
	Block   *blockchain.Block `json:"block"`
}
