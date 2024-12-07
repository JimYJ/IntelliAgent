package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	OpenAIBaseURL = "https://api.openai.com/v1"
)

// OpenAIClient OpenAI客户端结构体
type OpenAIClient struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

// ChatCompletionRequest 聊天请求结构体
type ChatCompletionRequest struct {
	Model        string                  `json:"model"`
	Messages     []ChatCompletionMessage `json:"messages"`
	Functions    []FunctionDefinition    `json:"functions,omitempty"`
	FunctionCall any                     `json:"function_call,omitempty"`
	Temperature  float32                 `json:"temperature,omitempty"`
	MaxTokens    int                     `json:"max_tokens,omitempty"`
}

// ChatCompletionMessage 聊天消息结构体
type ChatCompletionMessage struct {
	Role         string        `json:"role"`
	Content      string        `json:"content"`
	FunctionCall *FunctionCall `json:"function_call,omitempty"`
	Name         string        `json:"name,omitempty"`
}

// FunctionDefinition 函数定义结构体
type FunctionDefinition struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Parameters  Parameters `json:"parameters"`
}

// Parameters 函数参数结构体
type Parameters struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required,omitempty"`
}

// Property 参数属性结构体
type Property struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Enum        []string `json:"enum,omitempty"`
}

// FunctionCall 函数调用结构体
type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

// ChatCompletionChoice 聊天响应选项结构体
type ChatCompletionChoice struct {
	Message      ChatCompletionMessage `json:"message"`
	FinishReason string                `json:"finish_reason"`
}

// ChatCompletionUsage 聊天响应使用情况结构体
type ChatCompletionUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatCompletionResponse 聊天响应结构体
type ChatCompletionResponse struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Choices []ChatCompletionChoice `json:"choices"`
	Usage   ChatCompletionUsage    `json:"usage"`
}

// NewOpenAIClient 创建新的OpenAI客户端
func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		APIKey:     apiKey,
		BaseURL:    OpenAIBaseURL,
		HTTPClient: &http.Client{},
	}
}

// request 执行HTTP请求
func (c *OpenAIClient) request(method, path string, reqBody interface{}, result interface{}) error {
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("marshal request failed: %v", err)
	}
	request, err := http.NewRequest(method, c.BaseURL+path, bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("create request failed: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.APIKey)
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return fmt.Errorf("http request failed: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("read response body failed: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status %d: %s", response.StatusCode, string(body))
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("unmarshal response failed: %v", err)
	}

	return nil
}

// CreateChatCompletion 创建聊天完成请求
func (c *OpenAIClient) CreateChatCompletion(req ChatCompletionRequest) (*ChatCompletionResponse, error) {
	var result ChatCompletionResponse
	if err := c.request("POST", "/chat/completions", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
