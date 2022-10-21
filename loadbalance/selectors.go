package loadbalance

import (
	"aggregator/config"
	"aggregator/types"
	"sync"
)

var _selectors = map[string]*WrSelector{}
var _mutex sync.Mutex

func SetNodes(chain string, nodes []*types.Node) {
	_mutex.Lock()
	defer _mutex.Unlock()

	selector := &WrSelector{}
	selector.SetNodes(nodes)
	_selectors[chain] = selector
}

func NextNode(chain string) *types.Node {
	_mutex.Lock()
	defer _mutex.Unlock()

	selector := _selectors[chain]
	if selector != nil {
		return selector.NextNode()
	}

	return nil
}

func LoadFromConfig() {
	for chain, nodes := range config.Config.Nodes {
		SetNodes(chain, nodes)
	}
}
