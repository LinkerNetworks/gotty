package boomer

import (
	"errors"
	"time"
)

// Boomer is a count down timer and can trigger function `f` when timeout.
type Boomer struct {
	SECONDS_INIT uint64
	seconds      uint64
	f            interface{}
	armed        bool
	unarmed      bool
	boomed       bool
}

// NewBoomer creates a boomer, time is in seconds. You need to define function f.
func NewBoomer(seconds uint64, f interface{}) (*Boomer, error) {
	if seconds <= 0 {
		return nil, errors.New("invalid seconds")
	}
	p := &Boomer{}
	p.SECONDS_INIT = seconds
	p.seconds = seconds
	p.f = f
	return p, nil
}

// Arm let timer count down now.
func (p *Boomer) Arm() {
	go func() {
		for ; p.seconds+1 > 0; p.seconds-- {
			// fmt.Printf("\r%ds ", p.seconds)
			if p.seconds == 0 {
				p.boom(p.f)
			}
			time.Sleep(time.Second)
		}
	}()
	p.armed = true
	return
}

// Unarm cancels timer, the boomer will no longer boom.
func (p *Boomer) Unarm() {
	p.SECONDS_INIT = 0
	p.seconds = 0
	p.unarmed = true
	return
}

// Rewind count down timer to initial value(defined when you call Arm()).
func (p *Boomer) Rewind() error {
	if !p.armed {
		return errors.New("cannot rewind, boomer not armed")
	}
	if p.unarmed {
		return errors.New("cannot rewind, boomer unarmed")
	}
	if p.unarmed {
		return errors.New("cannot rewind, boomer has boomed")
	}
	p.seconds = p.SECONDS_INIT
	return nil
}

// Trigger function
func (p *Boomer) boom(f interface{}) {
	f.(func())()
	p.boomed = true
	return
}
