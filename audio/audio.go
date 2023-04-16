package audio

import (
	"github.com/vault-thirteen/SDLW/sdl"
	m "github.com/vault-thirteen/SDLW/sdl/model"
)

// GetInfo collects information about audio subsystem.
// SDL audio must be already initialized before using this function.
func GetInfo() (info *Info, err error) {
	info = &Info{}

	// Drivers.
	info.DriversCount = sdl.GetNumAudioDrivers()
	info.Drivers = make([]*Driver, 0, info.DriversCount)
	for i := 0; i < info.DriversCount; i++ {
		driverName := sdl.GetAudioDriver(i)
		info.Drivers = append(info.Drivers,
			&Driver{Index: i, Name: driverName})
	}

	// Playback Devices.
	info.PlaybackDevices, err = getDevices(0)
	if err != nil {
		return nil, err
	}
	info.PlaybackDevicesCount = len(info.PlaybackDevices)

	// Recording Devices.
	info.RecorderDevices, err = getDevices(1)
	if err != nil {
		return nil, err
	}
	info.RecorderDevicesCount = len(info.RecorderDevices)

	return info, nil
}

// getDevices lists devices for the specified mode.
func getDevices(devMode int) (devices []*Device, err error) {
	n := sdl.GetNumAudioDevices(devMode)
	devices = make([]*Device, 0, n)

	var device *Device
	for i := 0; i < n; i++ {
		device, err = getDeviceInfo(i, devMode)
		if err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}

	return devices, nil
}

// getDeviceInfo gets device information.
// 'devMode' is an 'isCapture' argument for 'GetAudioDeviceName' function.
func getDeviceInfo(index int, devMode int) (dev *Device, err error) {
	var deviceName = sdl.GetAudioDeviceName(index, devMode)

	var deviceSpec m.AudioSpec
	ret := sdl.GetAudioDeviceSpec(index, devMode, &deviceSpec)
	if ret != 0 {
		return nil, sdl.GetError()
	}

	return &Device{
		Index: index,
		Name:  deviceName,
		Spec:  &deviceSpec,
	}, nil
}