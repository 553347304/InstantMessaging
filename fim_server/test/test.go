package main

import (
	"fim_server/utils/open_api/open_api_qq"
)

func main() {
	open_api_qq.ReceiveEmail(open_api_qq.EmailConfig{
		Code:        "bcawbpbjmmxhbede",
		ReceiveUser: "553347304@qq.com",
	})
}
