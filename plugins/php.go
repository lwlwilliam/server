package plugins

import (
	"errors"
	"github.com/lwlwilliam/server/response"
	"log"
	"os/exec"
)

func PHP(file string) (string, error) {
	log.Println("PHP file:", file)

	// TODO: 暂时先用这种简陋的方法处理 php（在接收到 php 文件请求时，调用 php 解释器，接收解释器运行结果，再返回给浏览器），先理解大概运行机制，以后慢慢完善吧
	cmd := exec.Command("php", "-f", file)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return response.StatusText(response.InternalServerError), errors.New("PHP execute error")
	}

	log.Printf("PHP output: %s\n", output)

	return string(output), nil
}
