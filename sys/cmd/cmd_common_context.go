package cmd

import (
	"context"
	trlogger "github.com/woaijssss/tros/logx"
	"os"
	"os/exec"
	"syscall"
	"time"
)

// Wait() on Cmd, and swallow ExitError accordingly (for INT/KILL)
func waitProcessCommonContext(ctx context.Context, runCmd *exec.Cmd) error {
	trlogger.Debugf(ctx, "waitProcessCommonContext is called")
	err := runCmd.Wait()
	if err != nil {
		ee, ok := err.(*exec.ExitError)
		if !ok {
			trlogger.Errorf(ctx, "waitProcessCommonContext Wait is general error: [%+v]", err)
			return err
		}
		trlogger.Infof(ctx, "waitProcessCommonContext Wait is ExitError: [%+v]", ee)
		// TODO further examine exit reason -- but go error is string and hard to cover all cases
		return nil
	}
	return nil
}

// 退出子进程
func StopProcessCommonContext(ctx context.Context, runCmd *exec.Cmd, timeout int64) (err error) {
	killCancel, _ := StopProcessCommonContextNonBlocking(ctx, runCmd, timeout) // err logged in NonBlocking()
	defer killCancel.Stop()
	return waitProcessCommonContext(ctx, runCmd) // Wait() is here
}

// 发送信号，但不对 cmd 执行Wait()。
// 注：每个进程必须Wait，否则会出现僵尸进程(zombie process)
// 返回的timer可以用于取消ForceKillProcessCommonContext()
func StopProcessCommonContextNonBlocking(ctx context.Context, runCmd *exec.Cmd, timeout int64) (killCancel *time.Timer, err error) {
	trlogger.Debugf(ctx, "StopProcessCommonContext pid info %+v ", runCmd.Process)

	killCancel = time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		err1 := ForceKillProcessCommonContext(ctx, runCmd, false)
		if err1 != nil {
			trlogger.Errorf(ctx, "ForceKillProcessCommonContext error: %v", err1)
		}
	})
	// 退出事件
	err = GraceKillProcessCommonContext(ctx, runCmd, false)
	if err != nil {
		trlogger.Errorf(ctx, "GraceKillProcessCommonContext error: %v", err)
	}
	return
}

func signalProcessCommonContext(ctx context.Context, runCmd *exec.Cmd, sig os.Signal, isWait bool) (err error) {
	trlogger.Debugf(ctx, "signal %v start %d ", sig, time.Now().Unix())
	err = runCmd.Process.Signal(sig)
	if err != nil {
		trlogger.Errorf(ctx, "cmd.Process.Signal err %v", err)
		return
	}
	if isWait {
		err = waitProcessCommonContext(ctx, runCmd)
		if err != nil {
			trlogger.Errorf(ctx, "cmd waitProcessCommonContext err %v", err)
			return
		}
	}
	trlogger.Debugf(ctx, "signal %v end at %d ", sig, time.Now().Unix())
	return
}

func ForceKillProcessCommonContext(ctx context.Context, runCmd *exec.Cmd, isWait bool) (err error) {
	return signalProcessCommonContext(ctx, runCmd, syscall.SIGKILL, isWait)
}

func GraceKillProcessCommonContext(ctx context.Context, runCmd *exec.Cmd, isWait bool) (err error) {
	return signalProcessCommonContext(ctx, runCmd, syscall.SIGINT, isWait)
}

type CmdContextCommonContext struct {
	Stdout         *os.File `json:"stdout"`           //进程执行的输出参数
	MaxExecuteTime int64    `json:"max_execute_time"` //进程最多执行时间
	WaitTimeOut    int64    `json:"wait_time_out"`    //进程执行超时后，暴力杀死进程前的等待时间
	BinPath        string   `json:"bin_path"`         //可执行文件路径
	Args           []string `json:"args"`             //可执行文件参数
	BinEnv         []string `json:"bin_env"`          //可执行文件需要的环境变量
}

func RunCmdWithTimeoutCommonContext(ctx context.Context, CmdContextCommonContext CmdContextCommonContext) (err error) {
	trlogger.Debugf(ctx, "RunCmdWithTimeout context %+v ", CmdContextCommonContext)

	notifier := make(ExitChanCommonContext, 1)
	cancel := time.AfterFunc(time.Duration(CmdContextCommonContext.MaxExecuteTime)*time.Second, func() {
		trlogger.Infof(ctx, "RunCmdWithTimeout execute timeout")
		notifier <- true
	})
	defer cancel.Stop()

	return RunCmdWithExitSignalCommonContext(ctx, CmdContextCommonContext, notifier)
}

type ExitChanCommonContext chan bool

func RunCmdWithExitSignalCommonContext(ctx context.Context, CmdContextCommonContext CmdContextCommonContext, forceExitChanCommonContext ExitChanCommonContext) (err error) {
	trlogger.Debugf(ctx, "RunCmdWithExitSignalCommonContext context %+v ", CmdContextCommonContext)
	trlogger.Debugf(ctx, "RunCmdWithExitSignalCommonContext start %d ", time.Now().Unix())
	runCmd := exec.Command(CmdContextCommonContext.BinPath, CmdContextCommonContext.Args...)

	//cmd env
	runCmd.Env = os.Environ()
	runCmd.Env = append(runCmd.Env, CmdContextCommonContext.BinEnv...)

	runCmd.Stdout = CmdContextCommonContext.Stdout // 重定向标准输出到文件
	runCmd.Stderr = CmdContextCommonContext.Stdout // 重定向标准错误到文件

	trlogger.Debugf(ctx, "RunCmdWithExitSignalCommonContext info %+v ", runCmd)

	err = runCmd.Start() // 开始运行进程
	if err != nil {
		trlogger.Errorf(ctx, "RunCmdWithExitSignalCommonContext process Start() is err: %v", err)
		return err
	}

	// 处理各种退出，保证在开始进程之后执行以免出现奇怪的错误
	doneChan := make(chan struct{}, 1)
	killCancelChan := make(chan *time.Timer, 1)
	go func() {
		select {
		case <-forceExitChanCommonContext: // 收到退出信号，强制退出进程
			trlogger.Infof(ctx, "RunCmdWithExitSignalCommonContext process exit signal received")
			killCancel, err1 := StopProcessCommonContextNonBlocking(ctx, runCmd, CmdContextCommonContext.WaitTimeOut)
			if err1 != nil {
				trlogger.Errorf(ctx, "RunCmdWithExitSignalCommonContext StopProcessCommonContext is err: %v", err1)
			}
			killCancelChan <- killCancel
			return
		case <-doneChan: // 进程正常完成，取消这个线程
			killCancelChan <- nil
			return
		}
	}()

	err = runCmd.Wait()    // 等待进程结束，阻塞
	doneChan <- struct{}{} // 进程正常完成，可以取消等待退出信号了

	// 取消强制退出任务（如有），因为Wait()已经返回
	killCancel := <-killCancelChan
	if killCancel != nil {
		killCancel.Stop()
	}

	trlogger.Debugf(ctx, "RunCmdWithExitSignalCommonContext process end, err is %v", err)

	trlogger.Debugf(ctx, "RunCmdWithExitSignalCommonContext end %d ", time.Now().Unix())
	return err
}

// sudo strace -o output.log -p 66887  todo 会存在大量资源消耗
// ps -C gst-launch-1.0 -mww -o pid,ppid,pgid,cpu,lwp,stime,time,stat,wchan,cmd
func TraceOneProcessCommonContext(ctx context.Context, pid int, path string) (err error) {

	return
}
