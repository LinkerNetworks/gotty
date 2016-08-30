package app

import "strings"

// isLinkerContainer check container ID and only connections to containers created via Linker DCOS Portal is allowed.
// return true if cId like `mesos-865ef1d5-d353-4420-8284-3b83f553d3b4-S0.b3b6ed49-efd5-4984-9752-a26a7221afd8`
func isLinkerDcosContainer(cId string) bool {
	if strings.HasPrefix(cId, "mesos-") && strings.Contains(cId, ".") && len(cId) == 82 {
		return true
	}
	return false
}
