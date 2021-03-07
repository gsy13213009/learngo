package main

import (
	"context"
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)
	cmd = exec.Command("/bin/bash", "-c", "ls -a;echo 'hhh'")
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

type result struct {
	err error
	output []byte
}
func Test_main(t *testing.T) {
	resultChan := make(chan *result)
	// 执行1个cmd, 让他在一个携程里执行, 执行2s, sleep 2; echo hello;
	// 1s时, 杀死cmd
	ctx, cancelFunc := context.WithCancel(context.TODO())
	go func() {
		cmd := exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2;echo hello;")
		// 执行任务, 捕获输出
		output, err := cmd.CombinedOutput()
		resultChan <- &result{
			err: err,
			output: output,
		}
	}()
	// 继续往下走
	time.Sleep(1 * time.Second)
	// 取消任务
	cancelFunc()
	res := <- resultChan
	fmt.Println(res.err, string(res.output)) // 输出 `signal: killed`
}