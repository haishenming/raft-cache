package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	
	// 节点配置
	config := raft.DefaultConfig()
	config.LocalID = "1"
	config.Logger = hclog.New(&hclog.LoggerOptions{
		Name:   "myraft",
		Level:  hclog.LevelFromString("DEBUG"),
		Output: os.Stderr,
	})
	
	// logStore保存配置
	dir, _ := os.Getwd()
	logStore, err := raftboltdb.NewBoltStore(dir + "/n1/log_store.bolt")
	if err != nil {
		log.Fatal(err)
	}
	
	// 保存节点信息
	stableStore, err := raftboltdb.NewBoltStore(dir + "/n1/stable_store.bolt")
	if err != nil {
		log.Fatal(err)
	}
	
	// 不存储快照
	snapshotStore := raft.NewDiscardSnapshotStore()
	
	// 节点之间通讯方式
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:3000")
	transport, err := raft.NewTCPTransport(addr.String(), addr, 5, time.Second * 10, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	
	
	
	
	
	node, err := raft.NewRaft(config)
	
}
