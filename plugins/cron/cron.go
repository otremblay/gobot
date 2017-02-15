package cron

import (
        "fmt"
        "os/exec"
        "strings"

        "github.com/gabeguz/gobot"
        "github.com/robfig/cron"
)

type Cron struct{
c *cron.Cron
}

func NewCron(cronname, cronline string, bot gobot.Bot) *Cron {
        halves := strings.Split(cronline, "|")
        c := cron.New()
        donechan := make(chan string)
        errchan := make(chan string)
        c.AddFunc(halves[0], func(){
        		cmd := strings.Split(halves[1]," ")
                output, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput()
                if err == nil {
                        donechan<-string(output)
                } else {
                        errchan<-string(output)
                }
        })
        c.Start()
        go func(){
                for{
                        select{
                                case output := <-donechan:
                                        bot.Send(fmt.Sprintf("Successfully executed cron job '%s'. Job output:\n```%s```",cronname, output))
                                case output := <-errchan:
                                        bot.Send(fmt.Sprintf("Error executing cron job '%s'. Job output:\n```%s```",cronname, output))
                        }
                }
        }()
        return nil
}

func (p Cron) Name() string {
        return "Cron v1.0"
}

func (p Cron) Execute(msg gobot.Message, bot gobot.Bot) error {
        return nil
}