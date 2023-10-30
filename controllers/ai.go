package controllers

import (
	"SaveHouse/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"io/ioutil"
	"net/http"
	"os"
)

func GetRecommendation(c echo.Context) error {
	//ganti jadi AI buat admin JANGAN LUPA ***********
	OpenAI_Key := os.Getenv("API_OPENAI")

	var reqData models.GudangRequest

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	if err := json.Unmarshal(body, &reqData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	client := openai.NewClient(OpenAI_Key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Anda merupakan asisten yang dapat membantu untuk memberikan rekomendasi Tipe Gudang dimana terdapat 3 gudaang yang tersedia yaitu Gudang Lotus dengan kapasitas sebanyak 100 yang alamatnya berada di jl. mekar sari jakarta, Gudang Sakura dengan kapasitas sebanyak 150 yang alamatnya berada di jl. mekar jaya jakarta, dan Gudang Rahayu dengan kapasitas sebanyak 200 yang alamatnya berada di jl. taman mini jakarta.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Rekomendasi Tipe Gudang untuk barang %s dengan banyak barang  %s .", reqData.NamaBarang, reqData.Quantity),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}
	recommendation := resp.Choices[0].Message.Content

	response := models.AIResponse{
		Status: "success",
		Data:   recommendation,
	}

	return c.JSON(http.StatusOK, response)
}
