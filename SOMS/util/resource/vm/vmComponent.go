package resource

type IVmComponent interface {
	initialize() bool
	getSerialID() string
}
