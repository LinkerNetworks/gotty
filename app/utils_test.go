package app

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestIsLinkerDcosContainer(t *testing.T) {
	var cases = []struct {
		cId    string
		expect bool
	}{
		{"", false},
		{"9db6", false},
		{"93ac35942009", false},
		{"ef56e0991b27224c0770984309aec90879b161c7093f7732a78f53dc66c43071", false},
		{"mesos-1234.hacker-container", false},
		{"test830one_cadvisormonitor_1", false},
		{"mesos-865ef1d5-d353-4420-8284-3b83f553d3b4-S0.b3b6ed49-efd5-4984-9752-a26a7221afd8", true},
		{"mesos-865ef1d5-d353-4420-8284-3b83f553d3b4-S1.d4cae85c-894d-4acc-a3b1-6ecd6a4abf63", true},
		{"mesos-865ef1d5-d353-4420-8284-3b83f553d3b4-S1.716d9d07-d57b-48ec-b242-ebcff4b83587", true},
	}

	for _, c := range cases {
		//call
		got := isLinkerDcosContainer(c.cId)
		//assert
		assert.Equal(t, c.expect, got)
	}
}
