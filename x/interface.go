package x

//
// Rule is for desplay, XResource is actully worker
// XResource{
//  inEndId
//  enabled
//  status.
//}
//
type XResource interface {
	Test(inEndId string) bool      //0
	Register(inEndId string) error //1
	Start(e *RuleEngine) error     //2
	Enabled() bool
	Reload()
	Pause()
	Status(e *RuleEngine) State
	Stop()
}

//
// Stream from resource and to target
//
type XTarget interface {
	Test(outEndId string) bool      //0
	Register(outEndId string) error //1
	Start(e *RuleEngine) error      //2
	Enabled() bool
	Reload()
	Pause()
	Status(e *RuleEngine) State
	To(data interface{}) error
	Stop()
}

//
//
//
type XPlugin interface {
	Load(*RuleEngine) *XPluginEnv
	Init(*XPluginEnv) error
	Install(*XPluginEnv) (*XPluginMetaInfo, error)
	Start(*RuleEngine, *XPluginEnv) error
	Uninstall(*XPluginEnv) error
	Clean()
}

//
//
//
type XHook interface {
	Work(data string) error
	Name() string
}

//
//
//
type XStatus struct {
	inEndId string
	enabled bool
}
