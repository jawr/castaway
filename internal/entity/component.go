package entity

import "github.com/yourbasic/bit"

type ComponentFlag int

type ComponentFlags struct {
	set *bit.Set
}

func newComponentFlags() ComponentFlags {
	return ComponentFlags{
		set: bit.New(),
	}
}

func (f ComponentFlags) Add(flag ComponentFlag) {
	f.set = f.set.Add(int(flag))
}

func (f ComponentFlags) Remove(flag ComponentFlag) {
	f.set = f.set.Delete(int(flag))
}

func (f ComponentFlags) Contains(flags ...ComponentFlag) bool {
	for idx := range flags {
		if !f.set.Contains(int(flags[idx])) {
			return false
		}
	}
	return true
}

type Component interface {
	GetEntity() Entity
	GetFlag() ComponentFlag
}

const (
	ComponentAnimator ComponentFlag = iota
	ComponentPosition
	ComponentSprite
	ComponentSpeed
)
