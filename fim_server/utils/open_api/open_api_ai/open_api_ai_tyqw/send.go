package open_api_ai_tyqw

import (
	"bufio"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/https"
	"fim_server/utils/stores/logs"
)

type response struct {
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
		FinishReason interface{} `json:"finish_reason"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
	} `json:"choices"`
	Object            string      `json:"object"`
	Usage             interface{} `json:"usage"`
	Created           int         `json:"created"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
	Model             string      `json:"model"`
	Id                string      `json:"id"`
}
type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type request struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
	Stream   bool      `json:"stream"` // 流式输出
}

func Send(c Config) (messageChan chan string, err error) {

	messageChan = make(chan string)

	body := request{
		Model: "qwen-plus",
		Messages: []message{
			{Role: "system", Content: "你叫白音"},
			{Role: "user", Content: c.Content},
		},
		Stream: true,
	}

	r := https.Post("https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions", map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + c.ApiKey,
	}, conv.Json().Marshal(body))
	if r.Error != nil {
		return messageChan, logs.Error(r.Error)
	}

	scan := bufio.NewScanner(r.Response.Body) // 分片读
	scan.Split(bufio.ScanLines)               // 按行读

	go func() {
		for scan.Scan() {

			text := scan.Text()
			if text == "" {
				continue
			}

			if text == "data: [DONE]" {
				close(messageChan)
				return
			}

			var result response
			if !conv.Json().Unmarshal([]byte(text[6:]), &result) {
				return
			}

			for _, choice := range result.Choices {
				t := choice.Delta.Content
				if t == "" {
					continue
				}
				messageChan <- t
			}
		}
	}()
	return
}
