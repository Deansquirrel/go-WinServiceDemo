package worker

import (
	"fmt"
	"github.com/Deansquirrel/go-tool"
	"time"
)

type Worker struct {

}

func (p *Worker)Do(){
	for {
		strOut := go_tool.GetDateTimeStr(time.Now()) + " run"
		go_tool.Log(strOut)
		fmt.Println(strOut)
		time.Sleep(time.Second * 1)
	}
}