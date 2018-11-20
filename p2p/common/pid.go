package common

import (
	"path"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/libp2p/go-libp2p-protocol"
	"github.com/pkg/errors"
)

type Pid struct {
	name    string
	version *version.Version
}

func NewPid(name, rawVersion string) (Pid, error) {
	v, err := version.NewVersion(rawVersion)
	if err != nil {
		return Pid{}, errors.Wrap(err, "failed to parse verison")
	}

	return Pid{
		name:    name,
		version: v,
	}, nil
}

func ParsePid(pid string) (Pid, error) {
	if pid[0] != '/' {
		return Pid{}, errors.New("pid has no prefix '/'")
	}
	elems := strings.Split(pid, "/")
	name, rawVersion := elems[1], elems[2]

	v, err := version.NewVersion(rawVersion)
	if err != nil {
		return Pid{}, errors.Wrap(err, "failed to parse version")
	}

	return Pid{
		name:    name,
		version: v,
	}, nil
}

func (pid Pid) Name() string {
	return pid.name
}

func (pid Pid) Version() *version.Version {
	return pid.version
}

func (pid Pid) ProtocolID() protocol.ID {
	return protocol.ID(path.Join("/", pid.name, pid.version.String()))
}

type Pids []Pid

func (pids Pids) ProtocolID() []protocol.ID {
	ids := make([]protocol.ID, len(pids))
	for i, pid := range pids {
		ids[i] = pid.ProtocolID()
	}
	return ids
}
