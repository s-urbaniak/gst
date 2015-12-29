package gst

/*
#include <gst/gst.h>
*/
import "C"
import "github.com/s-urbaniak/glib"

const (
	GST_MESSAGE_DEVICE_ADDED   = C.GST_MESSAGE_DEVICE_ADDED
	GST_MESSAGE_DEVICE_REMOVED = C.GST_MESSAGE_DEVICE_REMOVED
)

type DeviceMonitor C.GstDeviceMonitor

func (dm *DeviceMonitor) g() *C.GstDeviceMonitor {
	return (*C.GstDeviceMonitor)(dm)
}

func NewDeviceMonitor() *DeviceMonitor {
	return (*DeviceMonitor)(C.gst_device_monitor_new())
}

func (dm *DeviceMonitor) Stop() {
	C.gst_device_monitor_stop(dm.g())
}

func (dm *DeviceMonitor) Start() {
	C.gst_device_monitor_start(dm.g())
}

func (dm *DeviceMonitor) AddFilter(classes string, caps *Caps) uint {
	ret := C.gst_device_monitor_add_filter(
		dm.g(),
		(*C.gchar)(C.CString(classes)),
		caps.g(),
	)

	return uint(ret)
}

func (dm *DeviceMonitor) GetBus() *Bus {
	gstBus := C.gst_device_monitor_get_bus(dm.g())

	b := new(Bus)
	b.SetPtr(glib.Pointer(gstBus))
	return b
}

func (dm *DeviceMonitor) GetDevices() []Device {
	l := glib.List{}
	l.SetPtr((glib.Pointer)(C.gst_device_monitor_get_devices(dm.g())))
	defer l.Free()

	ds := make([]Device, l.Length(), l.Length())
	var i uint
	for ; i < l.Length(); i++ {
		d := Device{}
		d.SetPtr(l.Nth(i))
		ds[i] = d
	}

	return ds
}
