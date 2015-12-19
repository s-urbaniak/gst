package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/ziutek/glib"
)

type StreamStatusType C.GstStreamStatusType

const (
	STREAM_CREATE  = StreamStatusType(C.GST_STREAM_STATUS_TYPE_CREATE)
	STREAM_ENTER   = StreamStatusType(C.GST_STREAM_STATUS_TYPE_ENTER)
	STREAM_LEAVE   = StreamStatusType(C.GST_STREAM_STATUS_TYPE_LEAVE)
	STREAM_DESTROY = StreamStatusType(C.GST_STREAM_STATUS_TYPE_DESTROY)
	STREAM_START   = StreamStatusType(C.GST_STREAM_STATUS_TYPE_START)
	STREAM_PAUSE   = StreamStatusType(C.GST_STREAM_STATUS_TYPE_PAUSE)
	STREAM_STOP    = StreamStatusType(C.GST_STREAM_STATUS_TYPE_STOP)
)

func (t StreamStatusType) String() string {
	switch t {
	case STREAM_CREATE:
		return "STREAM_CREATE"
	case STREAM_ENTER:
		return "STREAM_ENTER"
	case STREAM_LEAVE:
		return "STREAM_LEAVE"
	case STREAM_DESTROY:
		return "STREAM_DESTROY"
	case STREAM_START:
		return "STREAM_START"
	case STREAM_PAUSE:
		return "STREAM_PAUSE"
	case STREAM_STOP:
		return "STREAM_STOP"
	}
	return fmt.Sprintf("unknown stream status %v", t)
}

type MessageType C.GstMessageType

const (
	MESSAGE_UNKNOWN          = MessageType(C.GST_MESSAGE_UNKNOWN)
	MESSAGE_EOS              = MessageType(C.GST_MESSAGE_EOS)
	MESSAGE_ERROR            = MessageType(C.GST_MESSAGE_ERROR)
	MESSAGE_WARNING          = MessageType(C.GST_MESSAGE_WARNING)
	MESSAGE_INFO             = MessageType(C.GST_MESSAGE_INFO)
	MESSAGE_TAG              = MessageType(C.GST_MESSAGE_TAG)
	MESSAGE_BUFFERING        = MessageType(C.GST_MESSAGE_BUFFERING)
	MESSAGE_STATE_CHANGED    = MessageType(C.GST_MESSAGE_STATE_CHANGED)
	MESSAGE_STATE_DIRTY      = MessageType(C.GST_MESSAGE_STATE_DIRTY)
	MESSAGE_STEP_DONE        = MessageType(C.GST_MESSAGE_STEP_DONE)
	MESSAGE_CLOCK_PROVIDE    = MessageType(C.GST_MESSAGE_CLOCK_PROVIDE)
	MESSAGE_CLOCK_LOST       = MessageType(C.GST_MESSAGE_CLOCK_LOST)
	MESSAGE_NEW_CLOCK        = MessageType(C.GST_MESSAGE_NEW_CLOCK)
	MESSAGE_STRUCTURE_CHANGE = MessageType(C.GST_MESSAGE_STRUCTURE_CHANGE)
	MESSAGE_STREAM_STATUS    = MessageType(C.GST_MESSAGE_STREAM_STATUS)
	MESSAGE_APPLICATION      = MessageType(C.GST_MESSAGE_APPLICATION)
	MESSAGE_ELEMENT          = MessageType(C.GST_MESSAGE_ELEMENT)
	MESSAGE_SEGMENT_START    = MessageType(C.GST_MESSAGE_SEGMENT_START)
	MESSAGE_SEGMENT_DONE     = MessageType(C.GST_MESSAGE_SEGMENT_DONE)
	MESSAGE_DURATION_CHANGED = MessageType(C.GST_MESSAGE_DURATION_CHANGED)
	MESSAGE_LATENCY          = MessageType(C.GST_MESSAGE_LATENCY)
	MESSAGE_ASYNC_START      = MessageType(C.GST_MESSAGE_ASYNC_START)
	MESSAGE_ASYNC_DONE       = MessageType(C.GST_MESSAGE_ASYNC_DONE)
	MESSAGE_REQUEST_STATE    = MessageType(C.GST_MESSAGE_REQUEST_STATE)
	MESSAGE_STEP_START       = MessageType(C.GST_MESSAGE_STEP_START)
	MESSAGE_QOS              = MessageType(C.GST_MESSAGE_QOS)
	MESSAGE_PROGRESS         = MessageType(C.GST_MESSAGE_PROGRESS)
	MESSAGE_TOC              = MessageType(C.GST_MESSAGE_TOC)
	MESSAGE_RESET_TIME       = MessageType(C.GST_MESSAGE_RESET_TIME)
	MESSAGE_STREAM_START     = MessageType(C.GST_MESSAGE_STREAM_START)
	MESSAGE_NEED_CONTEXT     = MessageType(C.GST_MESSAGE_NEED_CONTEXT)
	MESSAGE_HAVE_CONTEXT     = MessageType(C.GST_MESSAGE_HAVE_CONTEXT)
	MESSAGE_EXTENDED         = MessageType(C.GST_MESSAGE_EXTENDED)
	MESSAGE_DEVICE_ADDED     = MessageType(C.GST_MESSAGE_DEVICE_ADDED)
	MESSAGE_DEVICE_REMOVED   = MessageType(C.GST_MESSAGE_DEVICE_REMOVED)
	MESSAGE_ANY              = MessageType(C.GST_MESSAGE_ANY)
)

func (t MessageType) String() string {
	switch t {
	case MESSAGE_UNKNOWN:
		return "MESSAGE_UNKNOWN"
	case MESSAGE_EOS:
		return "MESSAGE_EOS"
	case MESSAGE_ERROR:
		return "MESSAGE_ERROR"
	case MESSAGE_WARNING:
		return "MESSAGE_WARNING"
	case MESSAGE_INFO:
		return "MESSAGE_INFO"
	case MESSAGE_TAG:
		return "MESSAGE_TAG"
	case MESSAGE_BUFFERING:
		return "MESSAGE_BUFFERING"
	case MESSAGE_STATE_CHANGED:
		return "MESSAGE_STATE_CHANGED"
	case MESSAGE_STATE_DIRTY:
		return "MESSAGE_STATE_DIRTY"
	case MESSAGE_STEP_DONE:
		return "MESSAGE_STEP_DONE"
	case MESSAGE_CLOCK_PROVIDE:
		return "MESSAGE_CLOCK_PROVIDE"
	case MESSAGE_CLOCK_LOST:
		return "MESSAGE_CLOCK_LOST"
	case MESSAGE_NEW_CLOCK:
		return "MESSAGE_NEW_CLOCK"
	case MESSAGE_STRUCTURE_CHANGE:
		return "MESSAGE_STRUCTURE_CHANGE"
	case MESSAGE_STREAM_STATUS:
		return "MESSAGE_STREAM_STATUS"
	case MESSAGE_APPLICATION:
		return "MESSAGE_APPLICATION"
	case MESSAGE_ELEMENT:
		return "MESSAGE_ELEMENT"
	case MESSAGE_SEGMENT_START:
		return "MESSAGE_SEGMENT_START"
	case MESSAGE_SEGMENT_DONE:
		return "MESSAGE_SEGMENT_DONE"
	case MESSAGE_DURATION_CHANGED:
		return "MESSAGE_DURATION_CHANGED"
	case MESSAGE_LATENCY:
		return "MESSAGE_LATENCY"
	case MESSAGE_ASYNC_START:
		return "MESSAGE_ASYNC_START"
	case MESSAGE_ASYNC_DONE:
		return "MESSAGE_ASYNC_DONE"
	case MESSAGE_REQUEST_STATE:
		return "MESSAGE_REQUEST_STATE"
	case MESSAGE_STEP_START:
		return "MESSAGE_STEP_START"
	case MESSAGE_QOS:
		return "MESSAGE_QOS"
	case MESSAGE_PROGRESS:
		return "MESSAGE_PROGRESS"
	case MESSAGE_ANY:
		return "MESSAGE_ANY"
	case MESSAGE_TOC:
		return "MESSAGE_TOC"
	case MESSAGE_RESET_TIME:
		return "MESSAGE_RESET_TIME"
	case MESSAGE_STREAM_START:
		return "MESSAGE_STREAM_START"
	case MESSAGE_NEED_CONTEXT:
		return "MESSAGE_NEED_CONTEXT"
	case MESSAGE_HAVE_CONTEXT:
		return "MESSAGE_HAVE_CONTEXT"
	case MESSAGE_EXTENDED:
		return "MESSAGE_EXTENDED"
	case MESSAGE_DEVICE_ADDED:
		return "MESSAGE_DEVICE_ADDED"
	case MESSAGE_DEVICE_REMOVED:
		return "MESSAGE_DEVICE_REMOVED"
	}
	return fmt.Sprintf("unknown message type %b", t)
}

type Message C.GstMessage

func (m *Message) g() *C.GstMessage {
	return (*C.GstMessage)(m)
}

func (m *Message) Type() glib.Type {
	return glib.TypeFromName("GstMessage")
}

func (m *Message) Ref() *Message {
	return (*Message)(C.gst_message_ref(m.g()))
}

func (m *Message) Unref() {
	C.gst_message_unref(m.g())
}

func (m *Message) GetType() MessageType {
	return MessageType(m._type)
}

func (m *Message) GetStructure() (string, glib.Params) {
	s := C.gst_message_get_structure(m.g())
	if s == nil {
		return "", nil
	}
	return parseGstStructure(s)
}

func (m *Message) GetSrc() *GstObj {
	src := new(GstObj)
	src.SetPtr(glib.Pointer(m.src))
	return src
}

func (m *Message) ParseError() (err *glib.Error, debug string) {
	var d *C.gchar
	var e, ret_e *C.GError

	C.gst_message_parse_error(m.g(), &e, &d)
	defer C.free(unsafe.Pointer(e))
	defer C.free(unsafe.Pointer(d))

	debug = C.GoString((*C.char)(d))
	ret_e = new(C.GError)
	*ret_e = *e
	err = (*glib.Error)(unsafe.Pointer(ret_e))
	return
}

func (m *Message) ParseStateChanged() (new, old, pending State) {
	C.gst_message_parse_state_changed(m.g(), (&old).g(), (&new).g(), (&pending).g())
	return
}

func (m *Message) ParseStreamStatus() StreamStatusType {
	var t StreamStatusType

	C.gst_message_parse_stream_status(m.g(), (*C.GstStreamStatusType)(&t), nil)

	return t
}
