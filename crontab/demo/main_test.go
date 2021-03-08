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

// 调度多个任务

type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func Test_Crons(t *testing.T) {
	// 需要有1个调度协程, 定时检测所有的cron任务, 谁过期了就执行谁
	now := time.Now()
	expr := cronexpr.MustParse("*/5 * * * * * *")
	cronJob := &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	scheduleTable := map[string]*CronJob{}
	// 任务注册到调度表
	scheduleTable["job1"] = cronJob

	now2 := time.Now()
	expr2 := cronexpr.MustParse("*/5 * * * * * *")
	cronJob2 := &CronJob{
		expr:     expr2,
		nextTime: expr.Next(now2),
	}
	// 任务注册到调度表
	scheduleTable["job2"] = cronJob2
	// 启动调度协程
	go func() {
		for {
			now := time.Now()
			for jobName, cronJob := range scheduleTable {
				// 判断是否过期
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					// 启动一个携程, 去之执行这个任务
					go func(name string) {
						fmt.Println("执行:", jobName)
					}(jobName)
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println("下次执行时间:", cronJob.nextTime)
				}
			}
			select {
			case <- time.NewTimer(100 * time.Microsecond).C:// 将在100ms后可读, 返回
			}
		}
	}()
	time.Sleep(100 * time.Second)
}
