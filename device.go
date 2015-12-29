package gst

/*
#include <gst/gst.h>
*/
import "C"

type Device struct {
	GstObj
}

func (d *Device) g() *C.GstDevice {
	return (*C.GstDevice)(d.GetPtr())
}

func (d *Device) GetCaps() *Caps {
	return (*Caps)(C.gst_device_get_caps(d.g()))
}

func (d *Device) GetDisplayName() string {
	return C.GoString((*C.char)(C.gst_device_get_display_name(d.g())))
}
