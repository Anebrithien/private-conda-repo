package types

import (
	"io"
)

type Conda interface {
	// Creates a new Conda channel. Essentially, a conda channel is a folder containing all the necessary platform folders
	CreateChannel(channel string) (Channel, error)

	// Retrieves the specified channel
	GetChannel(channel string) (Channel, error)

	// Removes the specified Conda channel
	RemoveChannel(channel string) error

	// renames the conda channel
	ChangeChannelName(oldChannel, newChannel string) (Channel, error)
}

type Channel interface {
	// Returns the file location of the channel's volume
	Dir() string

	// Indexes the channel. This should be done whenever a package is uploaded or removed from the channel.
	// During this process, the metadata in the channeldata.json file will be updated. This json file is used
	// by conda to know how to install the package requested by the client
	Index() error

	// Gets the channel's meta information
	GetMetaInfo() (*ChannelMetaInfo, error)

	// Adds a package into the channel
	AddPackage(file io.Reader, platform, packageName string) error

	// Removes a package from the channel
	RemovePackage(platform, packageName string) error
}
