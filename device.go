package gst

/*
#include <gst/gst.h>
*/
import "C"
import "github.com/s-urbaniak/glib"

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

func (d *Device) GetDeviceClass() string {
	return C.GoString((*C.char)(C.gst_device_get_device_class(d.g())))
}

func (d *Device) GetProperties() (string, glib.Params) {
	s := C.gst_device_get_properties(d.g())
	defer C.gst_structure_free(s)

	return parseGstStructure(s)
}
