package controller

import "github.com/ahamlinman/typeparams-no-location/value"

type State int

const StateNone State = iota

type Status struct {
	State State
}

type Controller struct {
	status *value.Value[Status]
}

func NewController() *Controller {
	return &Controller{
		status: value.NewValue(Status{}),
	}
}

func (c *Controller) Status() Status {
	return c.status.Get()
}
