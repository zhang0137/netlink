package netlink

import "syscall"

// TCPSocketList lists tcp sockets in the system.
// Equivalent to: ss -t
func TCPSocketList() ([]*InetDiagMsg, error) {
	ret := make([]*InetDiagMsg, 0)
	req := NewInetDiagRequest()
	data := NewInetDiagReqV2(syscall.AF_INET, syscall.IPPROTO_TCP, TCP_ALL)
	req.AddData(data)
	msgs, err := req.Execute(syscall.NETLINK_INET_DIAG, 0)
	if err != nil {
		return ret, err
	}
	for _, msg := range msgs {
		ret = append(ret, ParseInetDiagMsg(msg))
	}
	return ret, nil
}
