package x

//
//
//
type outEnd struct {
	Id          string                  `json:"id"`
	Type        string                  `json:"type"`
	State       State                   `json:"state"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Config      *map[string]interface{} `json:"config"`
	Target      XTarget                 `json:"-"`
}

func (o *outEnd) GetState() State {
	lock.Lock()
	defer lock.Unlock()
	return o.State
}

//
func (o *outEnd) SetState(s State) {
	lock.Lock()
	defer lock.Unlock()
	o.State = s
}

//
//
//
func NewOutEnd(t string,
	n string,
	d string,
	c *map[string]interface{}) *outEnd {
	return &outEnd{
		Id:          MakeUUID("OUTEND"),
		Type:        t,
		State:       DOWN,
		Name:        n,
		Description: d,
		Config:      c,
	}
}

//
//
//
func (out *outEnd) GetConfig(k string) interface{} {
	return (*out.Config)[k]
}
