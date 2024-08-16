// automatically generated by stateify.

//go:build arm || mips || mipsle || 386
// +build arm mips mipsle 386

package atomicbitops

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (i *Int64) StateTypeName() string {
	return "pkg/atomicbitops.Int64"
}

func (i *Int64) StateFields() []string {
	return []string{
		"value",
		"value32",
	}
}

func (i *Int64) beforeSave() {}

// +checklocksignore
func (i *Int64) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.value)
	stateSinkObject.Save(1, &i.value32)
}

func (i *Int64) afterLoad(context.Context) {}

// +checklocksignore
func (i *Int64) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.value)
	stateSourceObject.Load(1, &i.value32)
}

func (u *Uint64) StateTypeName() string {
	return "pkg/atomicbitops.Uint64"
}

func (u *Uint64) StateFields() []string {
	return []string{
		"value",
		"value32",
	}
}

func (u *Uint64) beforeSave() {}

// +checklocksignore
func (u *Uint64) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	stateSinkObject.Save(0, &u.value)
	stateSinkObject.Save(1, &u.value32)
}

func (u *Uint64) afterLoad(context.Context) {}

// +checklocksignore
func (u *Uint64) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.value)
	stateSourceObject.Load(1, &u.value32)
}

func init() {
	state.Register((*Int64)(nil))
	state.Register((*Uint64)(nil))
}
