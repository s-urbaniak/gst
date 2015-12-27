package gst

/*
#include <gst/gst.h>
*/
import "C"

type Event C.GstEvent

func (e *Event) g() *C.GstEvent {
	return (*C.GstEvent)(e)
}

func NewEventEOS() *Event {
	return (*Event)(C.gst_event_new_eos())
}
