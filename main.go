package osversion

type Version struct {
	Major            uint32
	Minor            uint32
	Build            uint32
	PlatformId       uint32
	ServicePackMajor uint16
	ServicePackMinor uint16
}

func Query() *Version {
	return query()
}
