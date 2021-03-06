package etcd

import (
	"errors"
	"log"
	"strconv"
	"time"
)

// ServiceNode 服务注册对象
type ServiceNode struct {
	etcd *Etcd
	node Node
}

// NewServiceNode 新建一个服务注册对象
func NewServiceNode(endpoints []string, dc, nid, name string) *ServiceNode {
	etcd, err := NewEtcd(endpoints)
	if err != nil {
		log.Printf("NewServiceNode err = %v", err)
		return nil
	}

	var serverNode ServiceNode
	serverNode.etcd = etcd
	serverNode.node = Node{
		Ndc:  dc,
		Nid:  nid,
		Name: name,
		Npay: "0",
	}
	return &serverNode
}

// Close 关闭资源
func (serverNode *ServiceNode) Close() {
	if serverNode.etcd != nil {
		serverNode.etcd.Close()
	}
}

// NodeInfo 返回服务节点信息
func (serverNode *ServiceNode) NodeInfo() Node {
	return serverNode.node
}

// GetRPCChannel 获取RPC对象string
func (serverNode *ServiceNode) GetRPCChannel() string {
	return "rpc-" + serverNode.node.Nid
}

// GetEventChannel 获取广播对象string
func (serverNode *ServiceNode) GetEventChannel() string {
	return "event-" + serverNode.node.Nid
}

// RegisterNode 注册服务节点
func (serverNode *ServiceNode) RegisterNode() error {
	if serverNode.node.Ndc == "" || serverNode.node.Nid == "" || serverNode.node.Name == "" {
		return errors.New("Node dc id or name must be non empty")
	}
	go serverNode.keepRegistered(serverNode.node)
	return nil
}

// UpdateNodePayload 更新节点负载
func (serverNode *ServiceNode) UpdateNodePayload(payload int) error {
	if serverNode.node.Npay != strconv.Itoa(payload) {
		serverNode.node.Npay = strconv.Itoa(payload)
		go serverNode.updateRegistered(serverNode.node)
	}
	return nil
}

// keepRegistered 注册一个服务节点到etcd服务管理上
func (serverNode *ServiceNode) keepRegistered(node Node) {
	for {
		err := serverNode.etcd.Keep(node.Nid, node.GetNodeValue())
		if err != nil {
			log.Printf("keepRegistered err = %s", err)
			time.Sleep(5 * time.Second)
		} else {
			log.Printf("Node [%s] keepRegistered success!", node.Nid)
			return
		}
	}
}

// keepRegistered 更新一个服务节点到etcd服务管理上
func (serverNode *ServiceNode) updateRegistered(node Node) {
	for {
		err := serverNode.etcd.Update(node.Nid, node.GetNodeValue())
		if err != nil {
			log.Printf("updateRegistered err = %s", err)
			time.Sleep(5 * time.Second)
		} else {
			log.Printf("Node [%s] updateRegistered success!", node.Nid)
			return
		}
	}
}
