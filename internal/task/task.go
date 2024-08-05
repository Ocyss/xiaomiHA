package task

import (
	"github.com/alphadose/haxmap"
	"github.com/hibiken/asynq"
)

var T = &Tasks{
	cfgs:    haxmap.New[string, *asynq.PeriodicTaskConfig](),
	disable: haxmap.New[string, *asynq.PeriodicTaskConfig](),
	Mux:     asynq.NewServeMux(),
}

type Tasks struct {
	cfgs    *haxmap.Map[string, *asynq.PeriodicTaskConfig]
	disable *haxmap.Map[string, *asynq.PeriodicTaskConfig]
	Mux     *asynq.ServeMux
}

func (p *Tasks) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
	cfgs := make([]*asynq.PeriodicTaskConfig, 0, p.cfgs.Len()+10)
	p.cfgs.ForEach(func(key string, value *asynq.PeriodicTaskConfig) bool {
		cfgs = append(cfgs, value)
		return true
	})
	return cfgs, nil
}

func (p *Tasks) AddHandler(t *asynq.PeriodicTaskConfig, handler asynq.Handler) {
	name := t.Task.Type()
	p.Mux.Handle(name, handler)
	p.cfgs.Set(name, t)
}

func (p *Tasks) Disable(typename string) bool {
	if v, ok := p.cfgs.Get(typename); ok {
		p.cfgs.Del(typename)
		p.disable.Set(typename, v)
		return true
	} else {
		return false
	}
}
func (p *Tasks) Enable(typename string) bool {
	if v, ok := p.disable.Get(typename); ok {
		p.cfgs.Set(typename, v)
		p.disable.Del(typename)
		return true
	} else {
		return false
	}
}
