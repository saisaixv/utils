package cron

import(
	"sync"

	"github.com/pkg/errors"
	cron_v3 "github.com/robfig/cron/v3"

)

type Crontab struct{
	inner *cron_v3.Cron
	ids map[string]cron_v3.EntryID
	mutex sync.Mutex
}

func NewCrontab()*Crontab  {
	return &Crontab{
		inner:cron_v3.New(),
		ids:make(map[string]cron_v3.EntryID),
	}
}

func (c *Crontab)IDs()[]string  {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	validIDs:=make([]string, 0,len(c.ids))
	invalidIDs:=make([]string,0)
	for sid,eid:=range c.ids{
		if e:=c.inner.Entry(eid);e.ID!=eid{
			invalidIDs=append(invalidIDs,sid)
			continue
		}
		validIDs=append(validIDs,sid)
	}
	for _,id:=range invalidIDs{
		delete(c.ids,id)
	}
	return validIDs
}

func (c *Crontab)Start()  {
	c.inner.Start()
}

func (c *Crontab)Stop()  {
	c.inner.Stop()
}

func (c *Crontab)DelByID(id string)  {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	eid,ok:=c.ids[id]
	if !ok{
		return
	}
	c.inner.Remove(eid)
	delete(c.ids,id)
}

func (c *Crontab)AddByJob(id string,spec string,cmd cron_v3.Job)error  {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _,ok:=c.ids[id];ok{
		return errors.Errorf("crontab id exists")
	}

	eid,err:=c.inner.AddJob(spec,cmd)
	if err!=nil{
		return err
	}
	c.ids[id]=eid
	return nil
}

func (c *Crontab)AddByFunc(id string,spec string,f func())error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _,ok:=c.ids[id];ok{
		return errors.Errorf("crontab id exists")
	}
	eid,err:=c.inner.AddFunc(spec,f)
	if err!=nil{
		return err
	}
	c.ids[id]=eid
	return nil
}

func (c *Crontab)IsExists(jid string)bool  {
	_,exist:=c.ids[jid]
	return exist
}