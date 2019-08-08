package plugins

import (
	"errors"
	"github.com/lwlwilliam/server/response"
	"os/exec"
)

func public(command string, args []string) (string, error) {
	// TODO: 暂时先用这种简陋的方法处理（在接收到动态文件请求时，调用相应解释器，接收解释器运行结果，再返回给浏览器），先理解大概运行机制，以后慢慢完善吧
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return response.StatusText(response.InternalServerError), errors.New("python execute error")
	}

	return string(output), nil
}
