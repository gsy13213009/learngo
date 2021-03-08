package main

import (
	"context"
	"fmt"
	"github.com/gorhill/cronexpr"
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
	err    error
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
			err:    err,
			output: output,
		}
	}()
	// 继续往下走
	time.Sleep(1 * time.Second)
	// 取消任务
	cancelFunc()
	res := <-resultChan
	fmt.Println(res.err, string(res.output)) // 输出 `signal: killed`
}

// 调度一个任务, 每5s执行一次
func Test_Cron(t *testing.T) {
	// Seconds(0-59), Minutes(0-59), Hours(0-23), Day of month(1-31), Month(1-12), Day of week(0-6) Year(1970-2099)
	var expr *cronexpr.Expression
	var err error
	// 每5秒执行一次
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	now := time.Now()
	nextTime := expr.Next(now)
	fmt.Println(now, nextTime) // 2021-03-08 21:21:42.584238 +0800 CST m=+0.001131463 2021-03-08 21:21:45 +0800 CST
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了")
	})
	time.Sleep(10 * time.Second)
}
