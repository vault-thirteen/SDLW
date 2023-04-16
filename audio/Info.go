package audio

type Info struct {
	DriversCount int
	Drivers      []*Driver

	PlaybackDevicesCount int
	PlaybackDevices      []*Device

	RecorderDevicesCount int
	RecorderDevices      []*Device
}
