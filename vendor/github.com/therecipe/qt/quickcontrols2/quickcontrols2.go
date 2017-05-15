// +build !minimal

package quickcontrols2

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "quickcontrols2.h"
import "C"
import (
	"unsafe"
)

func cGoUnpackString(s C.struct_QtQuickControls2_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QQuickStyle struct {
	ptr unsafe.Pointer
}

type QQuickStyle_ITF interface {
	QQuickStyle_PTR() *QQuickStyle
}

func (ptr *QQuickStyle) QQuickStyle_PTR() *QQuickStyle {
	return ptr
}

func (ptr *QQuickStyle) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQuickStyle) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQuickStyle(ptr QQuickStyle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickStyle_PTR().Pointer()
	}
	return nil
}

func NewQQuickStyleFromPointer(ptr unsafe.Pointer) *QQuickStyle {
	var n = new(QQuickStyle)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQuickStyle) DestroyQQuickStyle() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func QQuickStyle_Name() string {
	return cGoUnpackString(C.QQuickStyle_QQuickStyle_Name())
}

func (ptr *QQuickStyle) Name() string {
	return cGoUnpackString(C.QQuickStyle_QQuickStyle_Name())
}

func QQuickStyle_Path() string {
	return cGoUnpackString(C.QQuickStyle_QQuickStyle_Path())
}

func (ptr *QQuickStyle) Path() string {
	return cGoUnpackString(C.QQuickStyle_QQuickStyle_Path())
}

func QQuickStyle_SetFallbackStyle(style string) {
	var styleC *C.char
	if style != "" {
		styleC = C.CString(style)
		defer C.free(unsafe.Pointer(styleC))
	}
	C.QQuickStyle_QQuickStyle_SetFallbackStyle(styleC)
}

func (ptr *QQuickStyle) SetFallbackStyle(style string) {
	var styleC *C.char
	if style != "" {
		styleC = C.CString(style)
		defer C.free(unsafe.Pointer(styleC))
	}
	C.QQuickStyle_QQuickStyle_SetFallbackStyle(styleC)
}

func QQuickStyle_SetStyle(style string) {
	var styleC *C.char
	if style != "" {
		styleC = C.CString(style)
		defer C.free(unsafe.Pointer(styleC))
	}
	C.QQuickStyle_QQuickStyle_SetStyle(styleC)
}

func (ptr *QQuickStyle) SetStyle(style string) {
	var styleC *C.char
	if style != "" {
		styleC = C.CString(style)
		defer C.free(unsafe.Pointer(styleC))
	}
	C.QQuickStyle_QQuickStyle_SetStyle(styleC)
}
