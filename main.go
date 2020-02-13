package main

import(
	"os"
	"fmt"
	"utils/cron"
)

type testTask struct{}

func (t *testTask)Run()  {
	fmt.Println("hello world")
}

func main()  {
	crontab:=cron.NewCrontab()
	task:=&testTask{}
	if err:=crontab.AddByJob("1","* * * * *",task);err!=nil{
		fmt.Printf("error to add crontab task:%s",err)
		os.Exit(-1)
	}

	taskFunc:=func(){
		fmt.Println("hello world")
	}

	if err:=crontab.AddByFunc("2","* * * * *",taskFunc);err!=nil{
		fmt.Printf("error to add crontab task:%s",err)
		os.Exit(-1)
	}
	crontab.Start()
	select{}
}