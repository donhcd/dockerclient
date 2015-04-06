package dockerclient

import (
	"io"
)

type Callback func(*Event, chan error, ...interface{})

type StatCallback func(string, *Stats, chan error, ...interface{})

type Client interface {
	Info() (*Info, error)
	ListContainers(all, size bool, filters string) ([]Container, error)
	InspectContainer(id string) (*ContainerInfo, error)
	CreateContainer(config *ContainerConfig, name string) (string, error)
	ContainerLogs(id string, options *LogOptions) (io.ReadCloser, error)
	ContainerChanges(id string) ([]*ContainerChanges, error)
	// ContainerStats returns a stats channel, an error channel, and a close
	// channel. Users should select on the stats and error channels. If
	// anything is sent on the error channels, then no more stats will be
	// sent. Users must close the close channel when they are done reading stats,
	// even if an error was sent.
	ContainerStats(id string) (<-chan Stats, <-chan error, chan<- struct{}, error)
	Exec(config *ExecConfig) (string, error)
	StartContainer(id string, config *HostConfig) error
	StopContainer(id string, timeout int) error
	RestartContainer(id string, timeout int) error
	KillContainer(id, signal string) error
	StartMonitorEvents(cb Callback, ec chan error, args ...interface{})
	StopAllMonitorEvents()
	StartMonitorStats(id string, cb StatCallback, ec chan error, args ...interface{})
	StopAllMonitorStats()
	TagImage(nameOrID string, repo string, tag string, force bool) error
	Version() (*Version, error)
	PullImage(name string, auth *AuthConfig) error
	RemoveContainer(id string, force, volumes bool) error
	ListImages() ([]*Image, error)
	RemoveImage(name string) ([]*ImageDelete, error)
	PauseContainer(name string) error
	UnpauseContainer(name string) error
}
