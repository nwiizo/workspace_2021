package main

import (
	"fmt"
	"reflect"

	"github.com/coreos/go-iptables/iptables"
)

func contains(list []string, value string) bool {
	for _, val := range list {
		if val == value {
			return true
		}
	}
	return false
}

func main() {
	ipt, err := iptables.New()
	chain := "sample1"
	// Saving the list of chains before executing tests
	originaListChain, err := ipt.ListChains("filter")
	if err != nil {
		fmt.Printf("ListChains of Initial failed: %v", err)
	}

	// chain shouldn't exist, this will create new
	err = ipt.ClearChain("filter", chain)
	if err != nil {
		fmt.Printf("ClearChain (of missing) failed: %v", err)
	}

	// chain should be in listChain
	listChain, err := ipt.ListChains("filter")
	if err != nil {
		fmt.Printf("ListChains failed: %v", err)
	}
	if !contains(listChain, chain) {
		fmt.Printf("ListChains doesn't contain the new chain %v", chain)
	}

	// chain now exists
	err = ipt.ClearChain("filter", chain)
	if err != nil {
		fmt.Printf("ClearChain (of empty) failed: %v", err)
	}

	// put a simple rule in
	err = ipt.Append("filter", chain, "-s", "0/0", "-j", "ACCEPT")
	if err != nil {
		fmt.Printf("Append failed: %v", err)
	}

	err = ipt.ClearChain("filter", chain)
	if err != nil {
		fmt.Printf("ClearChain (of non-empty) failed: %v\n", err)
	}

	// can't delete non-empty chain
	err = ipt.DeleteChain("filter", chain)
	if err == nil {
		fmt.Printf("DeleteChain of non-empty chain did not fail\n")
	}

	err = ipt.ClearChain("filter", chain)
	if err != nil {
		fmt.Printf("ClearChain (of non-empty) failed: %v", err)
	}

	// rename the chain
	newChain := "sample2"
	err = ipt.RenameChain("filter", chain, newChain)
	if err != nil {
		fmt.Printf("RenameChain failed: %v\n", err)
	}

	err = ipt.ClearChain("filter", newChain)
	if err != nil {
		fmt.Printf("ClearChain (of non-empty) failed: %v\n", err)
	}
	// chain empty, should be ok
	err = ipt.DeleteChain("filter", newChain)
	if err != nil {
		fmt.Printf("DeleteChain of empty chain failed: %v\n", err)
	}

	// check that chain is fully gone and that state similar to initial one
	listChain, err = ipt.ListChains("filter")
	if err != nil {
		fmt.Printf("ListChains failed: %v", err)
	}
	if !reflect.DeepEqual(originaListChain, listChain) {
		fmt.Printf("ListChains mismatch: \ngot  %#v \nneed %#v\n", originaListChain, listChain)
	}
	createMacFilter(ipt)
	createMacBasedProtFilter(ipt, 80)
	createInputToPortFilter(ipt)
}

const (
	tableName           = "filter"
	ProtFilterChainName = "port_jump"
	MacFilterChainName  = "mac_filter"
)

func createInputToPortFilter(ipt *iptables.IPTables) {
	chain := "INPUT"
	list, err := ipt.ListChains(tableName)
	fmt.Printf("chain list:%v", list)
	if err != nil {
		fmt.Printf("ListChains of Initial failed: %v", err)
	}
	isExists, err := ipt.Exists(tableName, chain, "-j", "port_jump")
	if !isExists {
		err = ipt.Append(tableName, chain, "-j", "port_jump")
		if err != nil {
			fmt.Printf("Append Input To Port Jump: %v\n", err)
		}
	}

}
func createMacBasedProtFilter(ipt *iptables.IPTables, port uint32) {
	chain := ProtFilterChainName
	err := ipt.ClearChain(tableName, chain)
	if err != nil {
		fmt.Printf("ClearChain (of non-empty) failed: %v\n", err)
	}
	err = ipt.Insert(tableName, chain, 1, "-p", "tcp", "--dport", fmt.Sprintf("%d", port), "-j", MacFilterChainName)
	err = ipt.Insert(tableName, chain, 1, "-p", "udp", "--dport", fmt.Sprintf("%d", port), "-j", MacFilterChainName)
}

func createMacFilter(ipt *iptables.IPTables) {
	chain := MacFilterChainName
	err := ipt.ClearChain(tableName, chain)
	if err != nil {
		fmt.Printf("ClearChain (of non-empty) failed: %v", err)
	}
	// put a simple rule in
	err = ipt.Insert(tableName, chain, 1, "-m", "mac", "--mac-source", "08:00:27:05:86:c6", "-j", "ACCEPT")
	if err != nil {
		fmt.Printf("Append failed: %v", err)
	}
	err = ipt.Append(tableName, chain, "-j", "DROP")
	if err != nil {
		fmt.Printf("Append failed: %v", err)
	}

}
