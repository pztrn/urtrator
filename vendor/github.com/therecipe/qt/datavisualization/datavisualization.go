// +build !minimal

package datavisualization

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "datavisualization.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtDataVisualization_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type Q3DBars struct {
	ptr unsafe.Pointer
}

type Q3DBars_ITF interface {
	Q3DBars_PTR() *Q3DBars
}

func (ptr *Q3DBars) Q3DBars_PTR() *Q3DBars {
	return ptr
}

func (ptr *Q3DBars) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *Q3DBars) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQ3DBars(ptr Q3DBars_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DBars_PTR().Pointer()
	}
	return nil
}

func NewQ3DBarsFromPointer(ptr unsafe.Pointer) *Q3DBars {
	var n = new(Q3DBars)
	n.SetPointer(ptr)
	return n
}
func NewQ3DBars(format gui.QSurfaceFormat_ITF, parent gui.QWindow_ITF) *Q3DBars {
	return NewQ3DBarsFromPointer(C.Q3DBars_NewQ3DBars(gui.PointerFromQSurfaceFormat(format), gui.PointerFromQWindow(parent)))
}

func (ptr *Q3DBars) AddAxis(axis QAbstract3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_AddAxis(ptr.Pointer(), PointerFromQAbstract3DAxis(axis))
	}
}

func (ptr *Q3DBars) AddSeries(series QBar3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_AddSeries(ptr.Pointer(), PointerFromQBar3DSeries(series))
	}
}

//export callbackQ3DBars_BarSpacingChanged
func callbackQ3DBars_BarSpacingChanged(ptr unsafe.Pointer, spacing unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "barSpacingChanged"); signal != nil {
		signal.(func(*core.QSizeF))(core.NewQSizeFFromPointer(spacing))
	}

}

func (ptr *Q3DBars) ConnectBarSpacingChanged(f func(spacing *core.QSizeF)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "barSpacingChanged") {
			C.Q3DBars_ConnectBarSpacingChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "barSpacingChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "barSpacingChanged", func(spacing *core.QSizeF) {
				signal.(func(*core.QSizeF))(spacing)
				f(spacing)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "barSpacingChanged", f)
		}
	}
}

func (ptr *Q3DBars) DisconnectBarSpacingChanged() {
	if ptr.Pointer() != nil {
		C.Q3DBars_DisconnectBarSpacingChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "barSpacingChanged")
	}
}

func (ptr *Q3DBars) BarSpacingChanged(spacing core.QSizeF_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_BarSpacingChanged(ptr.Pointer(), core.PointerFromQSizeF(spacing))
	}
}

//export callbackQ3DBars_BarSpacingRelativeChanged
func callbackQ3DBars_BarSpacingRelativeChanged(ptr unsafe.Pointer, relative C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "barSpacingRelativeChanged"); signal != nil {
		signal.(func(bool))(int8(relative) != 0)
	}

}

func (ptr *Q3DBars) ConnectBarSpacingRelativeChanged(f func(relative bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "barSpacingRelativeChanged") {
			C.Q3DBars_ConnectBarSpacingRelativeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "barSpacingRelativeChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "barSpacingRelativeChanged", func(relative bool) {
				signal.(func(bool))(relative)
				f(relative)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "barSpacingRelativeChanged", f)
		}
	}
}

func (ptr *Q3DBars) DisconnectBarSpacingRelativeChanged() {
	if ptr.Pointer() != nil {
		C.Q3DBars_DisconnectBarSpacingRelativeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "barSpacingRelativeChanged")
	}
}

func (ptr *Q3DBars) BarSpacingRelativeChanged(relative bool) {
	if ptr.Pointer() != nil {
		C.Q3DBars_BarSpacingRelativeChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(relative))))
	}
}

//export callbackQ3DBars_BarThicknessChanged
func callbackQ3DBars_BarThicknessChanged(ptr unsafe.Pointer, thicknessRatio C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "barThicknessChanged"); signal != nil {
		signal.(func(float32))(float32(thicknessRatio))
	}

}

func (ptr *Q3DBars) ConnectBarThicknessChanged(f func(thicknessRatio float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "barThicknessChanged") {
			C.Q3DBars_ConnectBarThicknessChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "barThicknessChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "barThicknessChanged", func(thicknessRatio float32) {
				signal.(func(float32))(thicknessRatio)
				f(thicknessRatio)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "barThicknessChanged", f)
		}
	}
}

func (ptr *Q3DBars) DisconnectBarThicknessChanged() {
	if ptr.Pointer() != nil {
		C.Q3DBars_DisconnectBarThicknessChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "barThicknessChanged")
	}
}

func (ptr *Q3DBars) BarThicknessChanged(thicknessRatio float32) {
	if ptr.Pointer() != nil {
		C.Q3DBars_BarThicknessChanged(ptr.Pointer(), C.float(thicknessRatio))
	}
}

//export callbackQ3DBars_ColumnAxisChanged
func callbackQ3DBars_ColumnAxisChanged(ptr unsafe.Pointer, axis unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "columnAxisChanged"); signal != nil {
		signal.(func(*QCategory3DAxis))(NewQCategory3DAxisFromPointer(axis))
	}

}

func (ptr *Q3DBars) ConnectColumnAxisChanged(f func(axis *QCategory3DAxis)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "columnAxisChanged") {
			C.Q3DBars_ConnectColumnAxisChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "columnAxisChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnAxisChanged", func(axis *QCategory3DAxis) {
				signal.(func(*QCategory3DAxis))(axis)
				f(axis)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnAxisChanged", f)
		}
	}
}

func (ptr *Q3DBars) DisconnectColumnAxisChanged() {
	if ptr.Pointer() != nil {
		C.Q3DBars_DisconnectColumnAxisChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "columnAxisChanged")
	}
}

func (ptr *Q3DBars) ColumnAxisChanged(axis QCategory3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_ColumnAxisChanged(ptr.Pointer(), PointerFromQCategory3DAxis(axis))
	}
}

//export callbackQ3DBars_FloorLevelChanged
func callbackQ3DBars_FloorLevelChanged(ptr unsafe.Pointer, level C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "floorLevelChanged"); signal != nil {
		signal.(func(float32))(float32(level))
	}

}

func (ptr *Q3DBars) ConnectFloorLevelChanged(f func(level float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "floorLevelChanged") {
			C.Q3DBars_ConnectFloorLevelChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "floorLevelChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "floorLevelChanged", func(level float32) {
				signal.(func(float32))(level)
				f(level)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "floorLevelChanged", f)
		}
	}
}

func (ptr *Q3DBars) DisconnectFloorLevelChanged() {
	if ptr.Pointer() != nil {
		C.Q3DBars_DisconnectFloorLevelChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "floorLevelChanged")
	}
}

func (ptr *Q3DBars) FloorLevelChanged(level float32) {
	if ptr.Pointer() != nil {
		C.Q3DBars_FloorLevelChanged(ptr.Pointer(), C.float(level))
	}
}

func (ptr *Q3DBars) InsertSeries(index int, series QBar3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_InsertSeries(ptr.Pointer(), C.int(int32(index)), PointerFromQBar3DSeries(series))
	}
}

//export callbackQ3DBars_MultiSeriesUniformChanged
func callbackQ3DBars_MultiSeriesUniformChanged(ptr unsafe.Pointer, uniform C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "multiSeriesUniformChanged"); signal != nil {
		signal.(func(bool))(int8(uniform) != 0)
	}

}

func (ptr *Q3DBars) ConnectMultiSeriesUniformChanged(f func(uniform bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "multiSeriesUniformChanged") {
			C.Q3DBars_ConnectMultiSeriesUniformChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "multiSeriesUniformChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiSeriesUniformChanged", func(uniform bool) {
				signal.(func(bool))(uniform)
				f(uniform)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiSeriesUniformChanged", f)
		}
	}
}

func (ptr *Q3DBars) DisconnectMultiSeriesUniformChanged() {
	if ptr.Pointer() != nil {
		C.Q3DBars_DisconnectMultiSeriesUniformChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "multiSeriesUniformChanged")
	}
}

func (ptr *Q3DBars) MultiSeriesUniformChanged(uniform bool) {
	if ptr.Pointer() != nil {
		C.Q3DBars_MultiSeriesUniformChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(uniform))))
	}
}

//export callbackQ3DBars_PrimarySeriesChanged
func callbackQ3DBars_PrimarySeriesChanged(ptr unsafe.Pointer, series unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "primarySeriesChanged"); signal != nil {
		signal.(func(*QBar3DSeries))(NewQBar3DSeriesFromPointer(series))
	}

}

func (ptr *Q3DBars) ConnectPrimarySeriesChanged(f func(series *QBar3DSeries)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "primarySeriesChanged") {
			C.Q3DBars_ConnectPrimarySeriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "primarySeriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "primarySeriesChanged", func(series *QBar3DSeries) {
				signal.(func(*QBar3DSeries))(series)
				f(series)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "primarySeriesChanged", f)
		}
	}
}

func (ptr *Q3DBars) DisconnectPrimarySeriesChanged() {
	if ptr.Pointer() != nil {
		C.Q3DBars_DisconnectPrimarySeriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "primarySeriesChanged")
	}
}

func (ptr *Q3DBars) PrimarySeriesChanged(series QBar3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_PrimarySeriesChanged(ptr.Pointer(), PointerFromQBar3DSeries(series))
	}
}

func (ptr *Q3DBars) ReleaseAxis(axis QAbstract3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_ReleaseAxis(ptr.Pointer(), PointerFromQAbstract3DAxis(axis))
	}
}

func (ptr *Q3DBars) RemoveSeries(series QBar3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_RemoveSeries(ptr.Pointer(), PointerFromQBar3DSeries(series))
	}
}

//export callbackQ3DBars_RowAxisChanged
func callbackQ3DBars_RowAxisChanged(ptr unsafe.Pointer, axis unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowAxisChanged"); signal != nil {
		signal.(func(*QCategory3DAxis))(NewQCategory3DAxisFromPointer(axis))
	}

}

func (ptr *Q3DBars) ConnectRowAxisChanged(f func(axis *QCategory3DAxis)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowAxisChanged") {
			C.Q3DBars_ConnectRowAxisChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowAxisChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowAxisChanged", func(axis *QCategory3DAxis) {
				signal.(func(*QCategory3DAxis))(axis)
				f(axis)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowAxisChanged", f)
		}
	}
}

func (ptr *Q3DBars) DisconnectRowAxisChanged() {
	if ptr.Pointer() != nil {
		C.Q3DBars_DisconnectRowAxisChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowAxisChanged")
	}
}

func (ptr *Q3DBars) RowAxisChanged(axis QCategory3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_RowAxisChanged(ptr.Pointer(), PointerFromQCategory3DAxis(axis))
	}
}

//export callbackQ3DBars_SelectedSeriesChanged
func callbackQ3DBars_SelectedSeriesChanged(ptr unsafe.Pointer, series unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "selectedSeriesChanged"); signal != nil {
		signal.(func(*QBar3DSeries))(NewQBar3DSeriesFromPointer(series))
	}

}

func (ptr *Q3DBars) ConnectSelectedSeriesChanged(f func(series *QBar3DSeries)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged") {
			C.Q3DBars_ConnectSelectedSeriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged", func(series *QBar3DSeries) {
				signal.(func(*QBar3DSeries))(series)
				f(series)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged", f)
		}
	}
}

func (ptr *Q3DBars) DisconnectSelectedSeriesChanged() {
	if ptr.Pointer() != nil {
		C.Q3DBars_DisconnectSelectedSeriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged")
	}
}

func (ptr *Q3DBars) SelectedSeriesChanged(series QBar3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_SelectedSeriesChanged(ptr.Pointer(), PointerFromQBar3DSeries(series))
	}
}

func (ptr *Q3DBars) SetBarSpacing(spacing core.QSizeF_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_SetBarSpacing(ptr.Pointer(), core.PointerFromQSizeF(spacing))
	}
}

func (ptr *Q3DBars) SetBarSpacingRelative(relative bool) {
	if ptr.Pointer() != nil {
		C.Q3DBars_SetBarSpacingRelative(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(relative))))
	}
}

func (ptr *Q3DBars) SetBarThickness(thicknessRatio float32) {
	if ptr.Pointer() != nil {
		C.Q3DBars_SetBarThickness(ptr.Pointer(), C.float(thicknessRatio))
	}
}

func (ptr *Q3DBars) SetColumnAxis(axis QCategory3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_SetColumnAxis(ptr.Pointer(), PointerFromQCategory3DAxis(axis))
	}
}

func (ptr *Q3DBars) SetFloorLevel(level float32) {
	if ptr.Pointer() != nil {
		C.Q3DBars_SetFloorLevel(ptr.Pointer(), C.float(level))
	}
}

func (ptr *Q3DBars) SetMultiSeriesUniform(uniform bool) {
	if ptr.Pointer() != nil {
		C.Q3DBars_SetMultiSeriesUniform(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(uniform))))
	}
}

func (ptr *Q3DBars) SetPrimarySeries(series QBar3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_SetPrimarySeries(ptr.Pointer(), PointerFromQBar3DSeries(series))
	}
}

func (ptr *Q3DBars) SetRowAxis(axis QCategory3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_SetRowAxis(ptr.Pointer(), PointerFromQCategory3DAxis(axis))
	}
}

func (ptr *Q3DBars) SetValueAxis(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_SetValueAxis(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

//export callbackQ3DBars_ValueAxisChanged
func callbackQ3DBars_ValueAxisChanged(ptr unsafe.Pointer, axis unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "valueAxisChanged"); signal != nil {
		signal.(func(*QValue3DAxis))(NewQValue3DAxisFromPointer(axis))
	}

}

func (ptr *Q3DBars) ConnectValueAxisChanged(f func(axis *QValue3DAxis)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "valueAxisChanged") {
			C.Q3DBars_ConnectValueAxisChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "valueAxisChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "valueAxisChanged", func(axis *QValue3DAxis) {
				signal.(func(*QValue3DAxis))(axis)
				f(axis)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "valueAxisChanged", f)
		}
	}
}

func (ptr *Q3DBars) DisconnectValueAxisChanged() {
	if ptr.Pointer() != nil {
		C.Q3DBars_DisconnectValueAxisChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "valueAxisChanged")
	}
}

func (ptr *Q3DBars) ValueAxisChanged(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars_ValueAxisChanged(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

//export callbackQ3DBars_DestroyQ3DBars
func callbackQ3DBars_DestroyQ3DBars(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~Q3DBars"); signal != nil {
		signal.(func())()
	} else {
		NewQ3DBarsFromPointer(ptr).DestroyQ3DBarsDefault()
	}
}

func (ptr *Q3DBars) ConnectDestroyQ3DBars(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~Q3DBars"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DBars", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DBars", f)
		}
	}
}

func (ptr *Q3DBars) DisconnectDestroyQ3DBars() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DBars")
	}
}

func (ptr *Q3DBars) DestroyQ3DBars() {
	if ptr.Pointer() != nil {
		C.Q3DBars_DestroyQ3DBars(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DBars) DestroyQ3DBarsDefault() {
	if ptr.Pointer() != nil {
		C.Q3DBars_DestroyQ3DBarsDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DBars) PrimarySeries() *QBar3DSeries {
	if ptr.Pointer() != nil {
		return NewQBar3DSeriesFromPointer(C.Q3DBars_PrimarySeries(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DBars) SelectedSeries() *QBar3DSeries {
	if ptr.Pointer() != nil {
		return NewQBar3DSeriesFromPointer(C.Q3DBars_SelectedSeries(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DBars) ColumnAxis() *QCategory3DAxis {
	if ptr.Pointer() != nil {
		return NewQCategory3DAxisFromPointer(C.Q3DBars_ColumnAxis(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DBars) RowAxis() *QCategory3DAxis {
	if ptr.Pointer() != nil {
		return NewQCategory3DAxisFromPointer(C.Q3DBars_RowAxis(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DBars) Axes() []*QAbstract3DAxis {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []*QAbstract3DAxis {
			var out = make([]*QAbstract3DAxis, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQ3DBarsFromPointer(l.data).__axes_atList(i)
			}
			return out
		}(C.Q3DBars_Axes(ptr.Pointer()))
	}
	return make([]*QAbstract3DAxis, 0)
}

func (ptr *Q3DBars) SeriesList() []*QBar3DSeries {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []*QBar3DSeries {
			var out = make([]*QBar3DSeries, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQ3DBarsFromPointer(l.data).__seriesList_atList(i)
			}
			return out
		}(C.Q3DBars_SeriesList(ptr.Pointer()))
	}
	return make([]*QBar3DSeries, 0)
}

func (ptr *Q3DBars) BarSpacing() *core.QSizeF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFFromPointer(C.Q3DBars_BarSpacing(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DBars) ValueAxis() *QValue3DAxis {
	if ptr.Pointer() != nil {
		return NewQValue3DAxisFromPointer(C.Q3DBars_ValueAxis(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DBars) IsBarSpacingRelative() bool {
	if ptr.Pointer() != nil {
		return C.Q3DBars_IsBarSpacingRelative(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DBars) IsMultiSeriesUniform() bool {
	if ptr.Pointer() != nil {
		return C.Q3DBars_IsMultiSeriesUniform(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DBars) BarThickness() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Q3DBars_BarThickness(ptr.Pointer()))
	}
	return 0
}

func (ptr *Q3DBars) FloorLevel() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Q3DBars_FloorLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *Q3DBars) __axes_atList(i int) *QAbstract3DAxis {
	if ptr.Pointer() != nil {
		return NewQAbstract3DAxisFromPointer(C.Q3DBars___axes_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *Q3DBars) __axes_setList(i QAbstract3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars___axes_setList(ptr.Pointer(), PointerFromQAbstract3DAxis(i))
	}
}

func (ptr *Q3DBars) __axes_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Q3DBars___axes_newList(ptr.Pointer()))
}

func (ptr *Q3DBars) __seriesList_atList(i int) *QBar3DSeries {
	if ptr.Pointer() != nil {
		return NewQBar3DSeriesFromPointer(C.Q3DBars___seriesList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *Q3DBars) __seriesList_setList(i QBar3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DBars___seriesList_setList(ptr.Pointer(), PointerFromQBar3DSeries(i))
	}
}

func (ptr *Q3DBars) __seriesList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Q3DBars___seriesList_newList(ptr.Pointer()))
}

type Q3DCamera struct {
	ptr unsafe.Pointer
}

type Q3DCamera_ITF interface {
	Q3DCamera_PTR() *Q3DCamera
}

func (ptr *Q3DCamera) Q3DCamera_PTR() *Q3DCamera {
	return ptr
}

func (ptr *Q3DCamera) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *Q3DCamera) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQ3DCamera(ptr Q3DCamera_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DCamera_PTR().Pointer()
	}
	return nil
}

func NewQ3DCameraFromPointer(ptr unsafe.Pointer) *Q3DCamera {
	var n = new(Q3DCamera)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=Q3DCamera__CameraPreset
//Q3DCamera::CameraPreset
type Q3DCamera__CameraPreset int64

const (
	Q3DCamera__CameraPresetNone               Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(-1)
	Q3DCamera__CameraPresetFrontLow           Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(0)
	Q3DCamera__CameraPresetFront              Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(1)
	Q3DCamera__CameraPresetFrontHigh          Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(2)
	Q3DCamera__CameraPresetLeftLow            Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(3)
	Q3DCamera__CameraPresetLeft               Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(4)
	Q3DCamera__CameraPresetLeftHigh           Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(5)
	Q3DCamera__CameraPresetRightLow           Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(6)
	Q3DCamera__CameraPresetRight              Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(7)
	Q3DCamera__CameraPresetRightHigh          Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(8)
	Q3DCamera__CameraPresetBehindLow          Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(9)
	Q3DCamera__CameraPresetBehind             Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(10)
	Q3DCamera__CameraPresetBehindHigh         Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(11)
	Q3DCamera__CameraPresetIsometricLeft      Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(12)
	Q3DCamera__CameraPresetIsometricLeftHigh  Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(13)
	Q3DCamera__CameraPresetIsometricRight     Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(14)
	Q3DCamera__CameraPresetIsometricRightHigh Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(15)
	Q3DCamera__CameraPresetDirectlyAbove      Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(16)
	Q3DCamera__CameraPresetDirectlyAboveCW45  Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(17)
	Q3DCamera__CameraPresetDirectlyAboveCCW45 Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(18)
	Q3DCamera__CameraPresetFrontBelow         Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(19)
	Q3DCamera__CameraPresetLeftBelow          Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(20)
	Q3DCamera__CameraPresetRightBelow         Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(21)
	Q3DCamera__CameraPresetBehindBelow        Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(22)
	Q3DCamera__CameraPresetDirectlyBelow      Q3DCamera__CameraPreset = Q3DCamera__CameraPreset(23)
)

func (ptr *Q3DCamera) SetCameraPosition(horizontal float32, vertical float32, zoom float32) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_SetCameraPosition(ptr.Pointer(), C.float(horizontal), C.float(vertical), C.float(zoom))
	}
}

func (ptr *Q3DCamera) SetTarget(target gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_SetTarget(ptr.Pointer(), gui.PointerFromQVector3D(target))
	}
}

func NewQ3DCamera(parent core.QObject_ITF) *Q3DCamera {
	return NewQ3DCameraFromPointer(C.Q3DCamera_NewQ3DCamera(core.PointerFromQObject(parent)))
}

//export callbackQ3DCamera_CameraPresetChanged
func callbackQ3DCamera_CameraPresetChanged(ptr unsafe.Pointer, preset C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "cameraPresetChanged"); signal != nil {
		signal.(func(Q3DCamera__CameraPreset))(Q3DCamera__CameraPreset(preset))
	}

}

func (ptr *Q3DCamera) ConnectCameraPresetChanged(f func(preset Q3DCamera__CameraPreset)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "cameraPresetChanged") {
			C.Q3DCamera_ConnectCameraPresetChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "cameraPresetChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "cameraPresetChanged", func(preset Q3DCamera__CameraPreset) {
				signal.(func(Q3DCamera__CameraPreset))(preset)
				f(preset)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "cameraPresetChanged", f)
		}
	}
}

func (ptr *Q3DCamera) DisconnectCameraPresetChanged() {
	if ptr.Pointer() != nil {
		C.Q3DCamera_DisconnectCameraPresetChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "cameraPresetChanged")
	}
}

func (ptr *Q3DCamera) CameraPresetChanged(preset Q3DCamera__CameraPreset) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_CameraPresetChanged(ptr.Pointer(), C.longlong(preset))
	}
}

//export callbackQ3DCamera_CopyValuesFrom
func callbackQ3DCamera_CopyValuesFrom(ptr unsafe.Pointer, source unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "copyValuesFrom"); signal != nil {
		signal.(func(*Q3DObject))(NewQ3DObjectFromPointer(source))
	} else {
		NewQ3DCameraFromPointer(ptr).CopyValuesFromDefault(NewQ3DObjectFromPointer(source))
	}
}

func (ptr *Q3DCamera) ConnectCopyValuesFrom(f func(source *Q3DObject)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "copyValuesFrom"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "copyValuesFrom", func(source *Q3DObject) {
				signal.(func(*Q3DObject))(source)
				f(source)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "copyValuesFrom", f)
		}
	}
}

func (ptr *Q3DCamera) DisconnectCopyValuesFrom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "copyValuesFrom")
	}
}

func (ptr *Q3DCamera) CopyValuesFrom(source Q3DObject_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_CopyValuesFrom(ptr.Pointer(), PointerFromQ3DObject(source))
	}
}

func (ptr *Q3DCamera) CopyValuesFromDefault(source Q3DObject_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_CopyValuesFromDefault(ptr.Pointer(), PointerFromQ3DObject(source))
	}
}

//export callbackQ3DCamera_MaxZoomLevelChanged
func callbackQ3DCamera_MaxZoomLevelChanged(ptr unsafe.Pointer, zoomLevel C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "maxZoomLevelChanged"); signal != nil {
		signal.(func(float32))(float32(zoomLevel))
	}

}

func (ptr *Q3DCamera) ConnectMaxZoomLevelChanged(f func(zoomLevel float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "maxZoomLevelChanged") {
			C.Q3DCamera_ConnectMaxZoomLevelChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "maxZoomLevelChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "maxZoomLevelChanged", func(zoomLevel float32) {
				signal.(func(float32))(zoomLevel)
				f(zoomLevel)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "maxZoomLevelChanged", f)
		}
	}
}

func (ptr *Q3DCamera) DisconnectMaxZoomLevelChanged() {
	if ptr.Pointer() != nil {
		C.Q3DCamera_DisconnectMaxZoomLevelChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "maxZoomLevelChanged")
	}
}

func (ptr *Q3DCamera) MaxZoomLevelChanged(zoomLevel float32) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_MaxZoomLevelChanged(ptr.Pointer(), C.float(zoomLevel))
	}
}

//export callbackQ3DCamera_MinZoomLevelChanged
func callbackQ3DCamera_MinZoomLevelChanged(ptr unsafe.Pointer, zoomLevel C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "minZoomLevelChanged"); signal != nil {
		signal.(func(float32))(float32(zoomLevel))
	}

}

func (ptr *Q3DCamera) ConnectMinZoomLevelChanged(f func(zoomLevel float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "minZoomLevelChanged") {
			C.Q3DCamera_ConnectMinZoomLevelChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "minZoomLevelChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "minZoomLevelChanged", func(zoomLevel float32) {
				signal.(func(float32))(zoomLevel)
				f(zoomLevel)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "minZoomLevelChanged", f)
		}
	}
}

func (ptr *Q3DCamera) DisconnectMinZoomLevelChanged() {
	if ptr.Pointer() != nil {
		C.Q3DCamera_DisconnectMinZoomLevelChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "minZoomLevelChanged")
	}
}

func (ptr *Q3DCamera) MinZoomLevelChanged(zoomLevel float32) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_MinZoomLevelChanged(ptr.Pointer(), C.float(zoomLevel))
	}
}

func (ptr *Q3DCamera) SetCameraPreset(preset Q3DCamera__CameraPreset) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_SetCameraPreset(ptr.Pointer(), C.longlong(preset))
	}
}

func (ptr *Q3DCamera) SetMaxZoomLevel(zoomLevel float32) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_SetMaxZoomLevel(ptr.Pointer(), C.float(zoomLevel))
	}
}

func (ptr *Q3DCamera) SetMinZoomLevel(zoomLevel float32) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_SetMinZoomLevel(ptr.Pointer(), C.float(zoomLevel))
	}
}

func (ptr *Q3DCamera) SetWrapXRotation(isEnabled bool) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_SetWrapXRotation(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isEnabled))))
	}
}

func (ptr *Q3DCamera) SetWrapYRotation(isEnabled bool) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_SetWrapYRotation(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isEnabled))))
	}
}

func (ptr *Q3DCamera) SetXRotation(rotation float32) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_SetXRotation(ptr.Pointer(), C.float(rotation))
	}
}

func (ptr *Q3DCamera) SetYRotation(rotation float32) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_SetYRotation(ptr.Pointer(), C.float(rotation))
	}
}

func (ptr *Q3DCamera) SetZoomLevel(zoomLevel float32) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_SetZoomLevel(ptr.Pointer(), C.float(zoomLevel))
	}
}

//export callbackQ3DCamera_TargetChanged
func callbackQ3DCamera_TargetChanged(ptr unsafe.Pointer, target unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "targetChanged"); signal != nil {
		signal.(func(*gui.QVector3D))(gui.NewQVector3DFromPointer(target))
	}

}

func (ptr *Q3DCamera) ConnectTargetChanged(f func(target *gui.QVector3D)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "targetChanged") {
			C.Q3DCamera_ConnectTargetChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "targetChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "targetChanged", func(target *gui.QVector3D) {
				signal.(func(*gui.QVector3D))(target)
				f(target)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "targetChanged", f)
		}
	}
}

func (ptr *Q3DCamera) DisconnectTargetChanged() {
	if ptr.Pointer() != nil {
		C.Q3DCamera_DisconnectTargetChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "targetChanged")
	}
}

func (ptr *Q3DCamera) TargetChanged(target gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_TargetChanged(ptr.Pointer(), gui.PointerFromQVector3D(target))
	}
}

//export callbackQ3DCamera_WrapXRotationChanged
func callbackQ3DCamera_WrapXRotationChanged(ptr unsafe.Pointer, isEnabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "wrapXRotationChanged"); signal != nil {
		signal.(func(bool))(int8(isEnabled) != 0)
	}

}

func (ptr *Q3DCamera) ConnectWrapXRotationChanged(f func(isEnabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "wrapXRotationChanged") {
			C.Q3DCamera_ConnectWrapXRotationChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "wrapXRotationChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "wrapXRotationChanged", func(isEnabled bool) {
				signal.(func(bool))(isEnabled)
				f(isEnabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "wrapXRotationChanged", f)
		}
	}
}

func (ptr *Q3DCamera) DisconnectWrapXRotationChanged() {
	if ptr.Pointer() != nil {
		C.Q3DCamera_DisconnectWrapXRotationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "wrapXRotationChanged")
	}
}

func (ptr *Q3DCamera) WrapXRotationChanged(isEnabled bool) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_WrapXRotationChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isEnabled))))
	}
}

//export callbackQ3DCamera_WrapYRotationChanged
func callbackQ3DCamera_WrapYRotationChanged(ptr unsafe.Pointer, isEnabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "wrapYRotationChanged"); signal != nil {
		signal.(func(bool))(int8(isEnabled) != 0)
	}

}

func (ptr *Q3DCamera) ConnectWrapYRotationChanged(f func(isEnabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "wrapYRotationChanged") {
			C.Q3DCamera_ConnectWrapYRotationChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "wrapYRotationChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "wrapYRotationChanged", func(isEnabled bool) {
				signal.(func(bool))(isEnabled)
				f(isEnabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "wrapYRotationChanged", f)
		}
	}
}

func (ptr *Q3DCamera) DisconnectWrapYRotationChanged() {
	if ptr.Pointer() != nil {
		C.Q3DCamera_DisconnectWrapYRotationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "wrapYRotationChanged")
	}
}

func (ptr *Q3DCamera) WrapYRotationChanged(isEnabled bool) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_WrapYRotationChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isEnabled))))
	}
}

//export callbackQ3DCamera_XRotationChanged
func callbackQ3DCamera_XRotationChanged(ptr unsafe.Pointer, rotation C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "xRotationChanged"); signal != nil {
		signal.(func(float32))(float32(rotation))
	}

}

func (ptr *Q3DCamera) ConnectXRotationChanged(f func(rotation float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "xRotationChanged") {
			C.Q3DCamera_ConnectXRotationChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "xRotationChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xRotationChanged", func(rotation float32) {
				signal.(func(float32))(rotation)
				f(rotation)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xRotationChanged", f)
		}
	}
}

func (ptr *Q3DCamera) DisconnectXRotationChanged() {
	if ptr.Pointer() != nil {
		C.Q3DCamera_DisconnectXRotationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "xRotationChanged")
	}
}

func (ptr *Q3DCamera) XRotationChanged(rotation float32) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_XRotationChanged(ptr.Pointer(), C.float(rotation))
	}
}

//export callbackQ3DCamera_YRotationChanged
func callbackQ3DCamera_YRotationChanged(ptr unsafe.Pointer, rotation C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "yRotationChanged"); signal != nil {
		signal.(func(float32))(float32(rotation))
	}

}

func (ptr *Q3DCamera) ConnectYRotationChanged(f func(rotation float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "yRotationChanged") {
			C.Q3DCamera_ConnectYRotationChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "yRotationChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yRotationChanged", func(rotation float32) {
				signal.(func(float32))(rotation)
				f(rotation)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yRotationChanged", f)
		}
	}
}

func (ptr *Q3DCamera) DisconnectYRotationChanged() {
	if ptr.Pointer() != nil {
		C.Q3DCamera_DisconnectYRotationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "yRotationChanged")
	}
}

func (ptr *Q3DCamera) YRotationChanged(rotation float32) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_YRotationChanged(ptr.Pointer(), C.float(rotation))
	}
}

//export callbackQ3DCamera_ZoomLevelChanged
func callbackQ3DCamera_ZoomLevelChanged(ptr unsafe.Pointer, zoomLevel C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "zoomLevelChanged"); signal != nil {
		signal.(func(float32))(float32(zoomLevel))
	}

}

func (ptr *Q3DCamera) ConnectZoomLevelChanged(f func(zoomLevel float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "zoomLevelChanged") {
			C.Q3DCamera_ConnectZoomLevelChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "zoomLevelChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zoomLevelChanged", func(zoomLevel float32) {
				signal.(func(float32))(zoomLevel)
				f(zoomLevel)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zoomLevelChanged", f)
		}
	}
}

func (ptr *Q3DCamera) DisconnectZoomLevelChanged() {
	if ptr.Pointer() != nil {
		C.Q3DCamera_DisconnectZoomLevelChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "zoomLevelChanged")
	}
}

func (ptr *Q3DCamera) ZoomLevelChanged(zoomLevel float32) {
	if ptr.Pointer() != nil {
		C.Q3DCamera_ZoomLevelChanged(ptr.Pointer(), C.float(zoomLevel))
	}
}

//export callbackQ3DCamera_DestroyQ3DCamera
func callbackQ3DCamera_DestroyQ3DCamera(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~Q3DCamera"); signal != nil {
		signal.(func())()
	} else {
		NewQ3DCameraFromPointer(ptr).DestroyQ3DCameraDefault()
	}
}

func (ptr *Q3DCamera) ConnectDestroyQ3DCamera(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~Q3DCamera"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DCamera", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DCamera", f)
		}
	}
}

func (ptr *Q3DCamera) DisconnectDestroyQ3DCamera() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DCamera")
	}
}

func (ptr *Q3DCamera) DestroyQ3DCamera() {
	if ptr.Pointer() != nil {
		C.Q3DCamera_DestroyQ3DCamera(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DCamera) DestroyQ3DCameraDefault() {
	if ptr.Pointer() != nil {
		C.Q3DCamera_DestroyQ3DCameraDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DCamera) CameraPreset() Q3DCamera__CameraPreset {
	if ptr.Pointer() != nil {
		return Q3DCamera__CameraPreset(C.Q3DCamera_CameraPreset(ptr.Pointer()))
	}
	return 0
}

func (ptr *Q3DCamera) Target() *gui.QVector3D {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQVector3DFromPointer(C.Q3DCamera_Target(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QVector3D).DestroyQVector3D)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DCamera) WrapXRotation() bool {
	if ptr.Pointer() != nil {
		return C.Q3DCamera_WrapXRotation(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DCamera) WrapYRotation() bool {
	if ptr.Pointer() != nil {
		return C.Q3DCamera_WrapYRotation(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DCamera) MaxZoomLevel() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Q3DCamera_MaxZoomLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *Q3DCamera) MinZoomLevel() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Q3DCamera_MinZoomLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *Q3DCamera) XRotation() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Q3DCamera_XRotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *Q3DCamera) YRotation() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Q3DCamera_YRotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *Q3DCamera) ZoomLevel() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Q3DCamera_ZoomLevel(ptr.Pointer()))
	}
	return 0
}

type Q3DInputHandler struct {
	ptr unsafe.Pointer
}

type Q3DInputHandler_ITF interface {
	Q3DInputHandler_PTR() *Q3DInputHandler
}

func (ptr *Q3DInputHandler) Q3DInputHandler_PTR() *Q3DInputHandler {
	return ptr
}

func (ptr *Q3DInputHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *Q3DInputHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQ3DInputHandler(ptr Q3DInputHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DInputHandler_PTR().Pointer()
	}
	return nil
}

func NewQ3DInputHandlerFromPointer(ptr unsafe.Pointer) *Q3DInputHandler {
	var n = new(Q3DInputHandler)
	n.SetPointer(ptr)
	return n
}
func NewQ3DInputHandler(parent core.QObject_ITF) *Q3DInputHandler {
	return NewQ3DInputHandlerFromPointer(C.Q3DInputHandler_NewQ3DInputHandler(core.PointerFromQObject(parent)))
}

//export callbackQ3DInputHandler_MouseMoveEvent
func callbackQ3DInputHandler_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer, mousePos unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent, *core.QPoint))(gui.NewQMouseEventFromPointer(event), core.NewQPointFromPointer(mousePos))
	} else {
		NewQ3DInputHandlerFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event), core.NewQPointFromPointer(mousePos))
	}
}

func (ptr *Q3DInputHandler) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent, mousePos *core.QPoint)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "mouseMoveEvent"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mouseMoveEvent", func(event *gui.QMouseEvent, mousePos *core.QPoint) {
				signal.(func(*gui.QMouseEvent, *core.QPoint))(event, mousePos)
				f(event, mousePos)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mouseMoveEvent", f)
		}
	}
}

func (ptr *Q3DInputHandler) DisconnectMouseMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "mouseMoveEvent")
	}
}

func (ptr *Q3DInputHandler) MouseMoveEvent(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event), core.PointerFromQPoint(mousePos))
	}
}

func (ptr *Q3DInputHandler) MouseMoveEventDefault(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event), core.PointerFromQPoint(mousePos))
	}
}

//export callbackQ3DInputHandler_MousePressEvent
func callbackQ3DInputHandler_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer, mousePos unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent, *core.QPoint))(gui.NewQMouseEventFromPointer(event), core.NewQPointFromPointer(mousePos))
	} else {
		NewQ3DInputHandlerFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event), core.NewQPointFromPointer(mousePos))
	}
}

func (ptr *Q3DInputHandler) ConnectMousePressEvent(f func(event *gui.QMouseEvent, mousePos *core.QPoint)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "mousePressEvent"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mousePressEvent", func(event *gui.QMouseEvent, mousePos *core.QPoint) {
				signal.(func(*gui.QMouseEvent, *core.QPoint))(event, mousePos)
				f(event, mousePos)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mousePressEvent", f)
		}
	}
}

func (ptr *Q3DInputHandler) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "mousePressEvent")
	}
}

func (ptr *Q3DInputHandler) MousePressEvent(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event), core.PointerFromQPoint(mousePos))
	}
}

func (ptr *Q3DInputHandler) MousePressEventDefault(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event), core.PointerFromQPoint(mousePos))
	}
}

//export callbackQ3DInputHandler_MouseReleaseEvent
func callbackQ3DInputHandler_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer, mousePos unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent, *core.QPoint))(gui.NewQMouseEventFromPointer(event), core.NewQPointFromPointer(mousePos))
	} else {
		NewQ3DInputHandlerFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event), core.NewQPointFromPointer(mousePos))
	}
}

func (ptr *Q3DInputHandler) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent, mousePos *core.QPoint)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "mouseReleaseEvent"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mouseReleaseEvent", func(event *gui.QMouseEvent, mousePos *core.QPoint) {
				signal.(func(*gui.QMouseEvent, *core.QPoint))(event, mousePos)
				f(event, mousePos)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mouseReleaseEvent", f)
		}
	}
}

func (ptr *Q3DInputHandler) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "mouseReleaseEvent")
	}
}

func (ptr *Q3DInputHandler) MouseReleaseEvent(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event), core.PointerFromQPoint(mousePos))
	}
}

func (ptr *Q3DInputHandler) MouseReleaseEventDefault(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event), core.PointerFromQPoint(mousePos))
	}
}

//export callbackQ3DInputHandler_RotationEnabledChanged
func callbackQ3DInputHandler_RotationEnabledChanged(ptr unsafe.Pointer, enable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rotationEnabledChanged"); signal != nil {
		signal.(func(bool))(int8(enable) != 0)
	}

}

func (ptr *Q3DInputHandler) ConnectRotationEnabledChanged(f func(enable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rotationEnabledChanged") {
			C.Q3DInputHandler_ConnectRotationEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rotationEnabledChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationEnabledChanged", func(enable bool) {
				signal.(func(bool))(enable)
				f(enable)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationEnabledChanged", f)
		}
	}
}

func (ptr *Q3DInputHandler) DisconnectRotationEnabledChanged() {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_DisconnectRotationEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rotationEnabledChanged")
	}
}

func (ptr *Q3DInputHandler) RotationEnabledChanged(enable bool) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_RotationEnabledChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQ3DInputHandler_SelectionEnabledChanged
func callbackQ3DInputHandler_SelectionEnabledChanged(ptr unsafe.Pointer, enable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "selectionEnabledChanged"); signal != nil {
		signal.(func(bool))(int8(enable) != 0)
	}

}

func (ptr *Q3DInputHandler) ConnectSelectionEnabledChanged(f func(enable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "selectionEnabledChanged") {
			C.Q3DInputHandler_ConnectSelectionEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "selectionEnabledChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectionEnabledChanged", func(enable bool) {
				signal.(func(bool))(enable)
				f(enable)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectionEnabledChanged", f)
		}
	}
}

func (ptr *Q3DInputHandler) DisconnectSelectionEnabledChanged() {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_DisconnectSelectionEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "selectionEnabledChanged")
	}
}

func (ptr *Q3DInputHandler) SelectionEnabledChanged(enable bool) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_SelectionEnabledChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *Q3DInputHandler) SetRotationEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_SetRotationEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *Q3DInputHandler) SetSelectionEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_SetSelectionEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *Q3DInputHandler) SetZoomAtTargetEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_SetZoomAtTargetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *Q3DInputHandler) SetZoomEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_SetZoomEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQ3DInputHandler_WheelEvent
func callbackQ3DInputHandler_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQ3DInputHandlerFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *Q3DInputHandler) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "wheelEvent"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "wheelEvent", func(event *gui.QWheelEvent) {
				signal.(func(*gui.QWheelEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "wheelEvent", f)
		}
	}
}

func (ptr *Q3DInputHandler) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "wheelEvent")
	}
}

func (ptr *Q3DInputHandler) WheelEvent(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *Q3DInputHandler) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQ3DInputHandler_ZoomAtTargetEnabledChanged
func callbackQ3DInputHandler_ZoomAtTargetEnabledChanged(ptr unsafe.Pointer, enable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "zoomAtTargetEnabledChanged"); signal != nil {
		signal.(func(bool))(int8(enable) != 0)
	}

}

func (ptr *Q3DInputHandler) ConnectZoomAtTargetEnabledChanged(f func(enable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "zoomAtTargetEnabledChanged") {
			C.Q3DInputHandler_ConnectZoomAtTargetEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "zoomAtTargetEnabledChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zoomAtTargetEnabledChanged", func(enable bool) {
				signal.(func(bool))(enable)
				f(enable)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zoomAtTargetEnabledChanged", f)
		}
	}
}

func (ptr *Q3DInputHandler) DisconnectZoomAtTargetEnabledChanged() {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_DisconnectZoomAtTargetEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "zoomAtTargetEnabledChanged")
	}
}

func (ptr *Q3DInputHandler) ZoomAtTargetEnabledChanged(enable bool) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_ZoomAtTargetEnabledChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQ3DInputHandler_ZoomEnabledChanged
func callbackQ3DInputHandler_ZoomEnabledChanged(ptr unsafe.Pointer, enable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "zoomEnabledChanged"); signal != nil {
		signal.(func(bool))(int8(enable) != 0)
	}

}

func (ptr *Q3DInputHandler) ConnectZoomEnabledChanged(f func(enable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "zoomEnabledChanged") {
			C.Q3DInputHandler_ConnectZoomEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "zoomEnabledChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zoomEnabledChanged", func(enable bool) {
				signal.(func(bool))(enable)
				f(enable)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zoomEnabledChanged", f)
		}
	}
}

func (ptr *Q3DInputHandler) DisconnectZoomEnabledChanged() {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_DisconnectZoomEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "zoomEnabledChanged")
	}
}

func (ptr *Q3DInputHandler) ZoomEnabledChanged(enable bool) {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_ZoomEnabledChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQ3DInputHandler_DestroyQ3DInputHandler
func callbackQ3DInputHandler_DestroyQ3DInputHandler(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~Q3DInputHandler"); signal != nil {
		signal.(func())()
	} else {
		NewQ3DInputHandlerFromPointer(ptr).DestroyQ3DInputHandlerDefault()
	}
}

func (ptr *Q3DInputHandler) ConnectDestroyQ3DInputHandler(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~Q3DInputHandler"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DInputHandler", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DInputHandler", f)
		}
	}
}

func (ptr *Q3DInputHandler) DisconnectDestroyQ3DInputHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DInputHandler")
	}
}

func (ptr *Q3DInputHandler) DestroyQ3DInputHandler() {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_DestroyQ3DInputHandler(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DInputHandler) DestroyQ3DInputHandlerDefault() {
	if ptr.Pointer() != nil {
		C.Q3DInputHandler_DestroyQ3DInputHandlerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DInputHandler) IsRotationEnabled() bool {
	if ptr.Pointer() != nil {
		return C.Q3DInputHandler_IsRotationEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DInputHandler) IsSelectionEnabled() bool {
	if ptr.Pointer() != nil {
		return C.Q3DInputHandler_IsSelectionEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DInputHandler) IsZoomAtTargetEnabled() bool {
	if ptr.Pointer() != nil {
		return C.Q3DInputHandler_IsZoomAtTargetEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DInputHandler) IsZoomEnabled() bool {
	if ptr.Pointer() != nil {
		return C.Q3DInputHandler_IsZoomEnabled(ptr.Pointer()) != 0
	}
	return false
}

type Q3DLight struct {
	ptr unsafe.Pointer
}

type Q3DLight_ITF interface {
	Q3DLight_PTR() *Q3DLight
}

func (ptr *Q3DLight) Q3DLight_PTR() *Q3DLight {
	return ptr
}

func (ptr *Q3DLight) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *Q3DLight) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQ3DLight(ptr Q3DLight_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DLight_PTR().Pointer()
	}
	return nil
}

func NewQ3DLightFromPointer(ptr unsafe.Pointer) *Q3DLight {
	var n = new(Q3DLight)
	n.SetPointer(ptr)
	return n
}
func NewQ3DLight(parent core.QObject_ITF) *Q3DLight {
	return NewQ3DLightFromPointer(C.Q3DLight_NewQ3DLight(core.PointerFromQObject(parent)))
}

//export callbackQ3DLight_DestroyQ3DLight
func callbackQ3DLight_DestroyQ3DLight(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~Q3DLight"); signal != nil {
		signal.(func())()
	} else {
		NewQ3DLightFromPointer(ptr).DestroyQ3DLightDefault()
	}
}

func (ptr *Q3DLight) ConnectDestroyQ3DLight(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~Q3DLight"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DLight", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DLight", f)
		}
	}
}

func (ptr *Q3DLight) DisconnectDestroyQ3DLight() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DLight")
	}
}

func (ptr *Q3DLight) DestroyQ3DLight() {
	if ptr.Pointer() != nil {
		C.Q3DLight_DestroyQ3DLight(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DLight) DestroyQ3DLightDefault() {
	if ptr.Pointer() != nil {
		C.Q3DLight_DestroyQ3DLightDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type Q3DObject struct {
	ptr unsafe.Pointer
}

type Q3DObject_ITF interface {
	Q3DObject_PTR() *Q3DObject
}

func (ptr *Q3DObject) Q3DObject_PTR() *Q3DObject {
	return ptr
}

func (ptr *Q3DObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *Q3DObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQ3DObject(ptr Q3DObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DObject_PTR().Pointer()
	}
	return nil
}

func NewQ3DObjectFromPointer(ptr unsafe.Pointer) *Q3DObject {
	var n = new(Q3DObject)
	n.SetPointer(ptr)
	return n
}
func NewQ3DObject(parent core.QObject_ITF) *Q3DObject {
	return NewQ3DObjectFromPointer(C.Q3DObject_NewQ3DObject(core.PointerFromQObject(parent)))
}

func (ptr *Q3DObject) ParentScene() *Q3DScene {
	if ptr.Pointer() != nil {
		return NewQ3DSceneFromPointer(C.Q3DObject_ParentScene(ptr.Pointer()))
	}
	return nil
}

//export callbackQ3DObject_CopyValuesFrom
func callbackQ3DObject_CopyValuesFrom(ptr unsafe.Pointer, source unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "copyValuesFrom"); signal != nil {
		signal.(func(*Q3DObject))(NewQ3DObjectFromPointer(source))
	} else {
		NewQ3DObjectFromPointer(ptr).CopyValuesFromDefault(NewQ3DObjectFromPointer(source))
	}
}

func (ptr *Q3DObject) ConnectCopyValuesFrom(f func(source *Q3DObject)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "copyValuesFrom"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "copyValuesFrom", func(source *Q3DObject) {
				signal.(func(*Q3DObject))(source)
				f(source)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "copyValuesFrom", f)
		}
	}
}

func (ptr *Q3DObject) DisconnectCopyValuesFrom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "copyValuesFrom")
	}
}

func (ptr *Q3DObject) CopyValuesFrom(source Q3DObject_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DObject_CopyValuesFrom(ptr.Pointer(), PointerFromQ3DObject(source))
	}
}

func (ptr *Q3DObject) CopyValuesFromDefault(source Q3DObject_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DObject_CopyValuesFromDefault(ptr.Pointer(), PointerFromQ3DObject(source))
	}
}

//export callbackQ3DObject_PositionChanged
func callbackQ3DObject_PositionChanged(ptr unsafe.Pointer, position unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "positionChanged"); signal != nil {
		signal.(func(*gui.QVector3D))(gui.NewQVector3DFromPointer(position))
	}

}

func (ptr *Q3DObject) ConnectPositionChanged(f func(position *gui.QVector3D)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "positionChanged") {
			C.Q3DObject_ConnectPositionChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "positionChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "positionChanged", func(position *gui.QVector3D) {
				signal.(func(*gui.QVector3D))(position)
				f(position)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "positionChanged", f)
		}
	}
}

func (ptr *Q3DObject) DisconnectPositionChanged() {
	if ptr.Pointer() != nil {
		C.Q3DObject_DisconnectPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "positionChanged")
	}
}

func (ptr *Q3DObject) PositionChanged(position gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DObject_PositionChanged(ptr.Pointer(), gui.PointerFromQVector3D(position))
	}
}

func (ptr *Q3DObject) SetDirty(dirty bool) {
	if ptr.Pointer() != nil {
		C.Q3DObject_SetDirty(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(dirty))))
	}
}

func (ptr *Q3DObject) SetPosition(position gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DObject_SetPosition(ptr.Pointer(), gui.PointerFromQVector3D(position))
	}
}

//export callbackQ3DObject_DestroyQ3DObject
func callbackQ3DObject_DestroyQ3DObject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~Q3DObject"); signal != nil {
		signal.(func())()
	} else {
		NewQ3DObjectFromPointer(ptr).DestroyQ3DObjectDefault()
	}
}

func (ptr *Q3DObject) ConnectDestroyQ3DObject(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~Q3DObject"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DObject", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DObject", f)
		}
	}
}

func (ptr *Q3DObject) DisconnectDestroyQ3DObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DObject")
	}
}

func (ptr *Q3DObject) DestroyQ3DObject() {
	if ptr.Pointer() != nil {
		C.Q3DObject_DestroyQ3DObject(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DObject) DestroyQ3DObjectDefault() {
	if ptr.Pointer() != nil {
		C.Q3DObject_DestroyQ3DObjectDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DObject) Position() *gui.QVector3D {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQVector3DFromPointer(C.Q3DObject_Position(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QVector3D).DestroyQVector3D)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DObject) IsDirty() bool {
	if ptr.Pointer() != nil {
		return C.Q3DObject_IsDirty(ptr.Pointer()) != 0
	}
	return false
}

type Q3DScatter struct {
	ptr unsafe.Pointer
}

type Q3DScatter_ITF interface {
	Q3DScatter_PTR() *Q3DScatter
}

func (ptr *Q3DScatter) Q3DScatter_PTR() *Q3DScatter {
	return ptr
}

func (ptr *Q3DScatter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *Q3DScatter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQ3DScatter(ptr Q3DScatter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DScatter_PTR().Pointer()
	}
	return nil
}

func NewQ3DScatterFromPointer(ptr unsafe.Pointer) *Q3DScatter {
	var n = new(Q3DScatter)
	n.SetPointer(ptr)
	return n
}
func NewQ3DScatter(format gui.QSurfaceFormat_ITF, parent gui.QWindow_ITF) *Q3DScatter {
	return NewQ3DScatterFromPointer(C.Q3DScatter_NewQ3DScatter(gui.PointerFromQSurfaceFormat(format), gui.PointerFromQWindow(parent)))
}

func (ptr *Q3DScatter) AddAxis(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScatter_AddAxis(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

func (ptr *Q3DScatter) AddSeries(series QScatter3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScatter_AddSeries(ptr.Pointer(), PointerFromQScatter3DSeries(series))
	}
}

//export callbackQ3DScatter_AxisXChanged
func callbackQ3DScatter_AxisXChanged(ptr unsafe.Pointer, axis unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "axisXChanged"); signal != nil {
		signal.(func(*QValue3DAxis))(NewQValue3DAxisFromPointer(axis))
	}

}

func (ptr *Q3DScatter) ConnectAxisXChanged(f func(axis *QValue3DAxis)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "axisXChanged") {
			C.Q3DScatter_ConnectAxisXChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "axisXChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisXChanged", func(axis *QValue3DAxis) {
				signal.(func(*QValue3DAxis))(axis)
				f(axis)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisXChanged", f)
		}
	}
}

func (ptr *Q3DScatter) DisconnectAxisXChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScatter_DisconnectAxisXChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "axisXChanged")
	}
}

func (ptr *Q3DScatter) AxisXChanged(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScatter_AxisXChanged(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

//export callbackQ3DScatter_AxisYChanged
func callbackQ3DScatter_AxisYChanged(ptr unsafe.Pointer, axis unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "axisYChanged"); signal != nil {
		signal.(func(*QValue3DAxis))(NewQValue3DAxisFromPointer(axis))
	}

}

func (ptr *Q3DScatter) ConnectAxisYChanged(f func(axis *QValue3DAxis)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "axisYChanged") {
			C.Q3DScatter_ConnectAxisYChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "axisYChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisYChanged", func(axis *QValue3DAxis) {
				signal.(func(*QValue3DAxis))(axis)
				f(axis)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisYChanged", f)
		}
	}
}

func (ptr *Q3DScatter) DisconnectAxisYChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScatter_DisconnectAxisYChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "axisYChanged")
	}
}

func (ptr *Q3DScatter) AxisYChanged(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScatter_AxisYChanged(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

//export callbackQ3DScatter_AxisZChanged
func callbackQ3DScatter_AxisZChanged(ptr unsafe.Pointer, axis unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "axisZChanged"); signal != nil {
		signal.(func(*QValue3DAxis))(NewQValue3DAxisFromPointer(axis))
	}

}

func (ptr *Q3DScatter) ConnectAxisZChanged(f func(axis *QValue3DAxis)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "axisZChanged") {
			C.Q3DScatter_ConnectAxisZChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "axisZChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisZChanged", func(axis *QValue3DAxis) {
				signal.(func(*QValue3DAxis))(axis)
				f(axis)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisZChanged", f)
		}
	}
}

func (ptr *Q3DScatter) DisconnectAxisZChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScatter_DisconnectAxisZChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "axisZChanged")
	}
}

func (ptr *Q3DScatter) AxisZChanged(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScatter_AxisZChanged(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

func (ptr *Q3DScatter) ReleaseAxis(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScatter_ReleaseAxis(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

func (ptr *Q3DScatter) RemoveSeries(series QScatter3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScatter_RemoveSeries(ptr.Pointer(), PointerFromQScatter3DSeries(series))
	}
}

//export callbackQ3DScatter_SelectedSeriesChanged
func callbackQ3DScatter_SelectedSeriesChanged(ptr unsafe.Pointer, series unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "selectedSeriesChanged"); signal != nil {
		signal.(func(*QScatter3DSeries))(NewQScatter3DSeriesFromPointer(series))
	}

}

func (ptr *Q3DScatter) ConnectSelectedSeriesChanged(f func(series *QScatter3DSeries)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged") {
			C.Q3DScatter_ConnectSelectedSeriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged", func(series *QScatter3DSeries) {
				signal.(func(*QScatter3DSeries))(series)
				f(series)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged", f)
		}
	}
}

func (ptr *Q3DScatter) DisconnectSelectedSeriesChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScatter_DisconnectSelectedSeriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged")
	}
}

func (ptr *Q3DScatter) SelectedSeriesChanged(series QScatter3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScatter_SelectedSeriesChanged(ptr.Pointer(), PointerFromQScatter3DSeries(series))
	}
}

func (ptr *Q3DScatter) SetAxisX(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScatter_SetAxisX(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

func (ptr *Q3DScatter) SetAxisY(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScatter_SetAxisY(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

func (ptr *Q3DScatter) SetAxisZ(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScatter_SetAxisZ(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

//export callbackQ3DScatter_DestroyQ3DScatter
func callbackQ3DScatter_DestroyQ3DScatter(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~Q3DScatter"); signal != nil {
		signal.(func())()
	} else {
		NewQ3DScatterFromPointer(ptr).DestroyQ3DScatterDefault()
	}
}

func (ptr *Q3DScatter) ConnectDestroyQ3DScatter(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~Q3DScatter"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DScatter", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DScatter", f)
		}
	}
}

func (ptr *Q3DScatter) DisconnectDestroyQ3DScatter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DScatter")
	}
}

func (ptr *Q3DScatter) DestroyQ3DScatter() {
	if ptr.Pointer() != nil {
		C.Q3DScatter_DestroyQ3DScatter(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DScatter) DestroyQ3DScatterDefault() {
	if ptr.Pointer() != nil {
		C.Q3DScatter_DestroyQ3DScatterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DScatter) SeriesList() []*QScatter3DSeries {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []*QScatter3DSeries {
			var out = make([]*QScatter3DSeries, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQ3DScatterFromPointer(l.data).__seriesList_atList(i)
			}
			return out
		}(C.Q3DScatter_SeriesList(ptr.Pointer()))
	}
	return make([]*QScatter3DSeries, 0)
}

func (ptr *Q3DScatter) Axes() []*QValue3DAxis {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []*QValue3DAxis {
			var out = make([]*QValue3DAxis, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQ3DScatterFromPointer(l.data).__axes_atList(i)
			}
			return out
		}(C.Q3DScatter_Axes(ptr.Pointer()))
	}
	return make([]*QValue3DAxis, 0)
}

func (ptr *Q3DScatter) SelectedSeries() *QScatter3DSeries {
	if ptr.Pointer() != nil {
		return NewQScatter3DSeriesFromPointer(C.Q3DScatter_SelectedSeries(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DScatter) AxisX() *QValue3DAxis {
	if ptr.Pointer() != nil {
		return NewQValue3DAxisFromPointer(C.Q3DScatter_AxisX(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DScatter) AxisY() *QValue3DAxis {
	if ptr.Pointer() != nil {
		return NewQValue3DAxisFromPointer(C.Q3DScatter_AxisY(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DScatter) AxisZ() *QValue3DAxis {
	if ptr.Pointer() != nil {
		return NewQValue3DAxisFromPointer(C.Q3DScatter_AxisZ(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DScatter) __seriesList_atList(i int) *QScatter3DSeries {
	if ptr.Pointer() != nil {
		return NewQScatter3DSeriesFromPointer(C.Q3DScatter___seriesList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *Q3DScatter) __seriesList_setList(i QScatter3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScatter___seriesList_setList(ptr.Pointer(), PointerFromQScatter3DSeries(i))
	}
}

func (ptr *Q3DScatter) __seriesList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Q3DScatter___seriesList_newList(ptr.Pointer()))
}

func (ptr *Q3DScatter) __axes_atList(i int) *QValue3DAxis {
	if ptr.Pointer() != nil {
		return NewQValue3DAxisFromPointer(C.Q3DScatter___axes_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *Q3DScatter) __axes_setList(i QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScatter___axes_setList(ptr.Pointer(), PointerFromQValue3DAxis(i))
	}
}

func (ptr *Q3DScatter) __axes_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Q3DScatter___axes_newList(ptr.Pointer()))
}

type Q3DScene struct {
	ptr unsafe.Pointer
}

type Q3DScene_ITF interface {
	Q3DScene_PTR() *Q3DScene
}

func (ptr *Q3DScene) Q3DScene_PTR() *Q3DScene {
	return ptr
}

func (ptr *Q3DScene) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *Q3DScene) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQ3DScene(ptr Q3DScene_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DScene_PTR().Pointer()
	}
	return nil
}

func NewQ3DSceneFromPointer(ptr unsafe.Pointer) *Q3DScene {
	var n = new(Q3DScene)
	n.SetPointer(ptr)
	return n
}
func NewQ3DScene(parent core.QObject_ITF) *Q3DScene {
	return NewQ3DSceneFromPointer(C.Q3DScene_NewQ3DScene(core.PointerFromQObject(parent)))
}

func Q3DScene_InvalidSelectionPoint() *core.QPoint {
	var tmpValue = core.NewQPointFromPointer(C.Q3DScene_Q3DScene_InvalidSelectionPoint())
	runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
	return tmpValue
}

func (ptr *Q3DScene) InvalidSelectionPoint() *core.QPoint {
	var tmpValue = core.NewQPointFromPointer(C.Q3DScene_Q3DScene_InvalidSelectionPoint())
	runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
	return tmpValue
}

func (ptr *Q3DScene) IsPointInPrimarySubView(point core.QPoint_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Q3DScene_IsPointInPrimarySubView(ptr.Pointer(), core.PointerFromQPoint(point)) != 0
	}
	return false
}

func (ptr *Q3DScene) IsPointInSecondarySubView(point core.QPoint_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Q3DScene_IsPointInSecondarySubView(ptr.Pointer(), core.PointerFromQPoint(point)) != 0
	}
	return false
}

//export callbackQ3DScene_ActiveCameraChanged
func callbackQ3DScene_ActiveCameraChanged(ptr unsafe.Pointer, camera unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "activeCameraChanged"); signal != nil {
		signal.(func(*Q3DCamera))(NewQ3DCameraFromPointer(camera))
	}

}

func (ptr *Q3DScene) ConnectActiveCameraChanged(f func(camera *Q3DCamera)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "activeCameraChanged") {
			C.Q3DScene_ConnectActiveCameraChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "activeCameraChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "activeCameraChanged", func(camera *Q3DCamera) {
				signal.(func(*Q3DCamera))(camera)
				f(camera)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "activeCameraChanged", f)
		}
	}
}

func (ptr *Q3DScene) DisconnectActiveCameraChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScene_DisconnectActiveCameraChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "activeCameraChanged")
	}
}

func (ptr *Q3DScene) ActiveCameraChanged(camera Q3DCamera_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScene_ActiveCameraChanged(ptr.Pointer(), PointerFromQ3DCamera(camera))
	}
}

//export callbackQ3DScene_ActiveLightChanged
func callbackQ3DScene_ActiveLightChanged(ptr unsafe.Pointer, light unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "activeLightChanged"); signal != nil {
		signal.(func(*Q3DLight))(NewQ3DLightFromPointer(light))
	}

}

func (ptr *Q3DScene) ConnectActiveLightChanged(f func(light *Q3DLight)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "activeLightChanged") {
			C.Q3DScene_ConnectActiveLightChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "activeLightChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "activeLightChanged", func(light *Q3DLight) {
				signal.(func(*Q3DLight))(light)
				f(light)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "activeLightChanged", f)
		}
	}
}

func (ptr *Q3DScene) DisconnectActiveLightChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScene_DisconnectActiveLightChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "activeLightChanged")
	}
}

func (ptr *Q3DScene) ActiveLightChanged(light Q3DLight_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScene_ActiveLightChanged(ptr.Pointer(), PointerFromQ3DLight(light))
	}
}

//export callbackQ3DScene_DevicePixelRatioChanged
func callbackQ3DScene_DevicePixelRatioChanged(ptr unsafe.Pointer, pixelRatio C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "devicePixelRatioChanged"); signal != nil {
		signal.(func(float32))(float32(pixelRatio))
	}

}

func (ptr *Q3DScene) ConnectDevicePixelRatioChanged(f func(pixelRatio float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "devicePixelRatioChanged") {
			C.Q3DScene_ConnectDevicePixelRatioChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "devicePixelRatioChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "devicePixelRatioChanged", func(pixelRatio float32) {
				signal.(func(float32))(pixelRatio)
				f(pixelRatio)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "devicePixelRatioChanged", f)
		}
	}
}

func (ptr *Q3DScene) DisconnectDevicePixelRatioChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScene_DisconnectDevicePixelRatioChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "devicePixelRatioChanged")
	}
}

func (ptr *Q3DScene) DevicePixelRatioChanged(pixelRatio float32) {
	if ptr.Pointer() != nil {
		C.Q3DScene_DevicePixelRatioChanged(ptr.Pointer(), C.float(pixelRatio))
	}
}

//export callbackQ3DScene_GraphPositionQueryChanged
func callbackQ3DScene_GraphPositionQueryChanged(ptr unsafe.Pointer, position unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "graphPositionQueryChanged"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(position))
	}

}

func (ptr *Q3DScene) ConnectGraphPositionQueryChanged(f func(position *core.QPoint)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "graphPositionQueryChanged") {
			C.Q3DScene_ConnectGraphPositionQueryChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "graphPositionQueryChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "graphPositionQueryChanged", func(position *core.QPoint) {
				signal.(func(*core.QPoint))(position)
				f(position)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "graphPositionQueryChanged", f)
		}
	}
}

func (ptr *Q3DScene) DisconnectGraphPositionQueryChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScene_DisconnectGraphPositionQueryChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "graphPositionQueryChanged")
	}
}

func (ptr *Q3DScene) GraphPositionQueryChanged(position core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScene_GraphPositionQueryChanged(ptr.Pointer(), core.PointerFromQPoint(position))
	}
}

//export callbackQ3DScene_PrimarySubViewportChanged
func callbackQ3DScene_PrimarySubViewportChanged(ptr unsafe.Pointer, subViewport unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "primarySubViewportChanged"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(subViewport))
	}

}

func (ptr *Q3DScene) ConnectPrimarySubViewportChanged(f func(subViewport *core.QRect)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "primarySubViewportChanged") {
			C.Q3DScene_ConnectPrimarySubViewportChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "primarySubViewportChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "primarySubViewportChanged", func(subViewport *core.QRect) {
				signal.(func(*core.QRect))(subViewport)
				f(subViewport)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "primarySubViewportChanged", f)
		}
	}
}

func (ptr *Q3DScene) DisconnectPrimarySubViewportChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScene_DisconnectPrimarySubViewportChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "primarySubViewportChanged")
	}
}

func (ptr *Q3DScene) PrimarySubViewportChanged(subViewport core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScene_PrimarySubViewportChanged(ptr.Pointer(), core.PointerFromQRect(subViewport))
	}
}

//export callbackQ3DScene_SecondarySubViewportChanged
func callbackQ3DScene_SecondarySubViewportChanged(ptr unsafe.Pointer, subViewport unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "secondarySubViewportChanged"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(subViewport))
	}

}

func (ptr *Q3DScene) ConnectSecondarySubViewportChanged(f func(subViewport *core.QRect)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "secondarySubViewportChanged") {
			C.Q3DScene_ConnectSecondarySubViewportChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "secondarySubViewportChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "secondarySubViewportChanged", func(subViewport *core.QRect) {
				signal.(func(*core.QRect))(subViewport)
				f(subViewport)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "secondarySubViewportChanged", f)
		}
	}
}

func (ptr *Q3DScene) DisconnectSecondarySubViewportChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScene_DisconnectSecondarySubViewportChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "secondarySubViewportChanged")
	}
}

func (ptr *Q3DScene) SecondarySubViewportChanged(subViewport core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScene_SecondarySubViewportChanged(ptr.Pointer(), core.PointerFromQRect(subViewport))
	}
}

//export callbackQ3DScene_SecondarySubviewOnTopChanged
func callbackQ3DScene_SecondarySubviewOnTopChanged(ptr unsafe.Pointer, isSecondaryOnTop C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "secondarySubviewOnTopChanged"); signal != nil {
		signal.(func(bool))(int8(isSecondaryOnTop) != 0)
	}

}

func (ptr *Q3DScene) ConnectSecondarySubviewOnTopChanged(f func(isSecondaryOnTop bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "secondarySubviewOnTopChanged") {
			C.Q3DScene_ConnectSecondarySubviewOnTopChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "secondarySubviewOnTopChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "secondarySubviewOnTopChanged", func(isSecondaryOnTop bool) {
				signal.(func(bool))(isSecondaryOnTop)
				f(isSecondaryOnTop)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "secondarySubviewOnTopChanged", f)
		}
	}
}

func (ptr *Q3DScene) DisconnectSecondarySubviewOnTopChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScene_DisconnectSecondarySubviewOnTopChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "secondarySubviewOnTopChanged")
	}
}

func (ptr *Q3DScene) SecondarySubviewOnTopChanged(isSecondaryOnTop bool) {
	if ptr.Pointer() != nil {
		C.Q3DScene_SecondarySubviewOnTopChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isSecondaryOnTop))))
	}
}

//export callbackQ3DScene_SelectionQueryPositionChanged
func callbackQ3DScene_SelectionQueryPositionChanged(ptr unsafe.Pointer, position unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "selectionQueryPositionChanged"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(position))
	}

}

func (ptr *Q3DScene) ConnectSelectionQueryPositionChanged(f func(position *core.QPoint)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "selectionQueryPositionChanged") {
			C.Q3DScene_ConnectSelectionQueryPositionChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "selectionQueryPositionChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectionQueryPositionChanged", func(position *core.QPoint) {
				signal.(func(*core.QPoint))(position)
				f(position)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectionQueryPositionChanged", f)
		}
	}
}

func (ptr *Q3DScene) DisconnectSelectionQueryPositionChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScene_DisconnectSelectionQueryPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "selectionQueryPositionChanged")
	}
}

func (ptr *Q3DScene) SelectionQueryPositionChanged(position core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScene_SelectionQueryPositionChanged(ptr.Pointer(), core.PointerFromQPoint(position))
	}
}

func (ptr *Q3DScene) SetActiveCamera(camera Q3DCamera_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScene_SetActiveCamera(ptr.Pointer(), PointerFromQ3DCamera(camera))
	}
}

func (ptr *Q3DScene) SetActiveLight(light Q3DLight_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScene_SetActiveLight(ptr.Pointer(), PointerFromQ3DLight(light))
	}
}

func (ptr *Q3DScene) SetDevicePixelRatio(pixelRatio float32) {
	if ptr.Pointer() != nil {
		C.Q3DScene_SetDevicePixelRatio(ptr.Pointer(), C.float(pixelRatio))
	}
}

func (ptr *Q3DScene) SetGraphPositionQuery(point core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScene_SetGraphPositionQuery(ptr.Pointer(), core.PointerFromQPoint(point))
	}
}

func (ptr *Q3DScene) SetPrimarySubViewport(primarySubViewport core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScene_SetPrimarySubViewport(ptr.Pointer(), core.PointerFromQRect(primarySubViewport))
	}
}

func (ptr *Q3DScene) SetSecondarySubViewport(secondarySubViewport core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScene_SetSecondarySubViewport(ptr.Pointer(), core.PointerFromQRect(secondarySubViewport))
	}
}

func (ptr *Q3DScene) SetSecondarySubviewOnTop(isSecondaryOnTop bool) {
	if ptr.Pointer() != nil {
		C.Q3DScene_SetSecondarySubviewOnTop(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isSecondaryOnTop))))
	}
}

func (ptr *Q3DScene) SetSelectionQueryPosition(point core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScene_SetSelectionQueryPosition(ptr.Pointer(), core.PointerFromQPoint(point))
	}
}

func (ptr *Q3DScene) SetSlicingActive(isSlicing bool) {
	if ptr.Pointer() != nil {
		C.Q3DScene_SetSlicingActive(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isSlicing))))
	}
}

//export callbackQ3DScene_SlicingActiveChanged
func callbackQ3DScene_SlicingActiveChanged(ptr unsafe.Pointer, isSlicingActive C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "slicingActiveChanged"); signal != nil {
		signal.(func(bool))(int8(isSlicingActive) != 0)
	}

}

func (ptr *Q3DScene) ConnectSlicingActiveChanged(f func(isSlicingActive bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "slicingActiveChanged") {
			C.Q3DScene_ConnectSlicingActiveChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "slicingActiveChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "slicingActiveChanged", func(isSlicingActive bool) {
				signal.(func(bool))(isSlicingActive)
				f(isSlicingActive)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "slicingActiveChanged", f)
		}
	}
}

func (ptr *Q3DScene) DisconnectSlicingActiveChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScene_DisconnectSlicingActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "slicingActiveChanged")
	}
}

func (ptr *Q3DScene) SlicingActiveChanged(isSlicingActive bool) {
	if ptr.Pointer() != nil {
		C.Q3DScene_SlicingActiveChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isSlicingActive))))
	}
}

//export callbackQ3DScene_ViewportChanged
func callbackQ3DScene_ViewportChanged(ptr unsafe.Pointer, viewport unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "viewportChanged"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(viewport))
	}

}

func (ptr *Q3DScene) ConnectViewportChanged(f func(viewport *core.QRect)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "viewportChanged") {
			C.Q3DScene_ConnectViewportChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "viewportChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "viewportChanged", func(viewport *core.QRect) {
				signal.(func(*core.QRect))(viewport)
				f(viewport)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "viewportChanged", f)
		}
	}
}

func (ptr *Q3DScene) DisconnectViewportChanged() {
	if ptr.Pointer() != nil {
		C.Q3DScene_DisconnectViewportChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "viewportChanged")
	}
}

func (ptr *Q3DScene) ViewportChanged(viewport core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DScene_ViewportChanged(ptr.Pointer(), core.PointerFromQRect(viewport))
	}
}

//export callbackQ3DScene_DestroyQ3DScene
func callbackQ3DScene_DestroyQ3DScene(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~Q3DScene"); signal != nil {
		signal.(func())()
	} else {
		NewQ3DSceneFromPointer(ptr).DestroyQ3DSceneDefault()
	}
}

func (ptr *Q3DScene) ConnectDestroyQ3DScene(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~Q3DScene"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DScene", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DScene", f)
		}
	}
}

func (ptr *Q3DScene) DisconnectDestroyQ3DScene() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DScene")
	}
}

func (ptr *Q3DScene) DestroyQ3DScene() {
	if ptr.Pointer() != nil {
		C.Q3DScene_DestroyQ3DScene(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DScene) DestroyQ3DSceneDefault() {
	if ptr.Pointer() != nil {
		C.Q3DScene_DestroyQ3DSceneDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DScene) ActiveCamera() *Q3DCamera {
	if ptr.Pointer() != nil {
		return NewQ3DCameraFromPointer(C.Q3DScene_ActiveCamera(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DScene) ActiveLight() *Q3DLight {
	if ptr.Pointer() != nil {
		return NewQ3DLightFromPointer(C.Q3DScene_ActiveLight(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DScene) GraphPositionQuery() *core.QPoint {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFromPointer(C.Q3DScene_GraphPositionQuery(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DScene) SelectionQueryPosition() *core.QPoint {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFromPointer(C.Q3DScene_SelectionQueryPosition(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DScene) PrimarySubViewport() *core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.Q3DScene_PrimarySubViewport(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DScene) SecondarySubViewport() *core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.Q3DScene_SecondarySubViewport(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DScene) Viewport() *core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.Q3DScene_Viewport(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DScene) IsSecondarySubviewOnTop() bool {
	if ptr.Pointer() != nil {
		return C.Q3DScene_IsSecondarySubviewOnTop(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DScene) IsSlicingActive() bool {
	if ptr.Pointer() != nil {
		return C.Q3DScene_IsSlicingActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DScene) DevicePixelRatio() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Q3DScene_DevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

type Q3DSurface struct {
	ptr unsafe.Pointer
}

type Q3DSurface_ITF interface {
	Q3DSurface_PTR() *Q3DSurface
}

func (ptr *Q3DSurface) Q3DSurface_PTR() *Q3DSurface {
	return ptr
}

func (ptr *Q3DSurface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *Q3DSurface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQ3DSurface(ptr Q3DSurface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DSurface_PTR().Pointer()
	}
	return nil
}

func NewQ3DSurfaceFromPointer(ptr unsafe.Pointer) *Q3DSurface {
	var n = new(Q3DSurface)
	n.SetPointer(ptr)
	return n
}
func NewQ3DSurface(format gui.QSurfaceFormat_ITF, parent gui.QWindow_ITF) *Q3DSurface {
	return NewQ3DSurfaceFromPointer(C.Q3DSurface_NewQ3DSurface(gui.PointerFromQSurfaceFormat(format), gui.PointerFromQWindow(parent)))
}

func (ptr *Q3DSurface) AddAxis(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DSurface_AddAxis(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

func (ptr *Q3DSurface) AddSeries(series QSurface3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DSurface_AddSeries(ptr.Pointer(), PointerFromQSurface3DSeries(series))
	}
}

//export callbackQ3DSurface_AxisXChanged
func callbackQ3DSurface_AxisXChanged(ptr unsafe.Pointer, axis unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "axisXChanged"); signal != nil {
		signal.(func(*QValue3DAxis))(NewQValue3DAxisFromPointer(axis))
	}

}

func (ptr *Q3DSurface) ConnectAxisXChanged(f func(axis *QValue3DAxis)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "axisXChanged") {
			C.Q3DSurface_ConnectAxisXChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "axisXChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisXChanged", func(axis *QValue3DAxis) {
				signal.(func(*QValue3DAxis))(axis)
				f(axis)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisXChanged", f)
		}
	}
}

func (ptr *Q3DSurface) DisconnectAxisXChanged() {
	if ptr.Pointer() != nil {
		C.Q3DSurface_DisconnectAxisXChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "axisXChanged")
	}
}

func (ptr *Q3DSurface) AxisXChanged(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DSurface_AxisXChanged(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

//export callbackQ3DSurface_AxisYChanged
func callbackQ3DSurface_AxisYChanged(ptr unsafe.Pointer, axis unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "axisYChanged"); signal != nil {
		signal.(func(*QValue3DAxis))(NewQValue3DAxisFromPointer(axis))
	}

}

func (ptr *Q3DSurface) ConnectAxisYChanged(f func(axis *QValue3DAxis)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "axisYChanged") {
			C.Q3DSurface_ConnectAxisYChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "axisYChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisYChanged", func(axis *QValue3DAxis) {
				signal.(func(*QValue3DAxis))(axis)
				f(axis)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisYChanged", f)
		}
	}
}

func (ptr *Q3DSurface) DisconnectAxisYChanged() {
	if ptr.Pointer() != nil {
		C.Q3DSurface_DisconnectAxisYChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "axisYChanged")
	}
}

func (ptr *Q3DSurface) AxisYChanged(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DSurface_AxisYChanged(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

//export callbackQ3DSurface_AxisZChanged
func callbackQ3DSurface_AxisZChanged(ptr unsafe.Pointer, axis unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "axisZChanged"); signal != nil {
		signal.(func(*QValue3DAxis))(NewQValue3DAxisFromPointer(axis))
	}

}

func (ptr *Q3DSurface) ConnectAxisZChanged(f func(axis *QValue3DAxis)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "axisZChanged") {
			C.Q3DSurface_ConnectAxisZChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "axisZChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisZChanged", func(axis *QValue3DAxis) {
				signal.(func(*QValue3DAxis))(axis)
				f(axis)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisZChanged", f)
		}
	}
}

func (ptr *Q3DSurface) DisconnectAxisZChanged() {
	if ptr.Pointer() != nil {
		C.Q3DSurface_DisconnectAxisZChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "axisZChanged")
	}
}

func (ptr *Q3DSurface) AxisZChanged(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DSurface_AxisZChanged(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

//export callbackQ3DSurface_FlipHorizontalGridChanged
func callbackQ3DSurface_FlipHorizontalGridChanged(ptr unsafe.Pointer, flip C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "flipHorizontalGridChanged"); signal != nil {
		signal.(func(bool))(int8(flip) != 0)
	}

}

func (ptr *Q3DSurface) ConnectFlipHorizontalGridChanged(f func(flip bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "flipHorizontalGridChanged") {
			C.Q3DSurface_ConnectFlipHorizontalGridChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "flipHorizontalGridChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "flipHorizontalGridChanged", func(flip bool) {
				signal.(func(bool))(flip)
				f(flip)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "flipHorizontalGridChanged", f)
		}
	}
}

func (ptr *Q3DSurface) DisconnectFlipHorizontalGridChanged() {
	if ptr.Pointer() != nil {
		C.Q3DSurface_DisconnectFlipHorizontalGridChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "flipHorizontalGridChanged")
	}
}

func (ptr *Q3DSurface) FlipHorizontalGridChanged(flip bool) {
	if ptr.Pointer() != nil {
		C.Q3DSurface_FlipHorizontalGridChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(flip))))
	}
}

func (ptr *Q3DSurface) ReleaseAxis(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DSurface_ReleaseAxis(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

func (ptr *Q3DSurface) RemoveSeries(series QSurface3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DSurface_RemoveSeries(ptr.Pointer(), PointerFromQSurface3DSeries(series))
	}
}

//export callbackQ3DSurface_SelectedSeriesChanged
func callbackQ3DSurface_SelectedSeriesChanged(ptr unsafe.Pointer, series unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "selectedSeriesChanged"); signal != nil {
		signal.(func(*QSurface3DSeries))(NewQSurface3DSeriesFromPointer(series))
	}

}

func (ptr *Q3DSurface) ConnectSelectedSeriesChanged(f func(series *QSurface3DSeries)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged") {
			C.Q3DSurface_ConnectSelectedSeriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged", func(series *QSurface3DSeries) {
				signal.(func(*QSurface3DSeries))(series)
				f(series)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged", f)
		}
	}
}

func (ptr *Q3DSurface) DisconnectSelectedSeriesChanged() {
	if ptr.Pointer() != nil {
		C.Q3DSurface_DisconnectSelectedSeriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "selectedSeriesChanged")
	}
}

func (ptr *Q3DSurface) SelectedSeriesChanged(series QSurface3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DSurface_SelectedSeriesChanged(ptr.Pointer(), PointerFromQSurface3DSeries(series))
	}
}

func (ptr *Q3DSurface) SetAxisX(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DSurface_SetAxisX(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

func (ptr *Q3DSurface) SetAxisY(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DSurface_SetAxisY(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

func (ptr *Q3DSurface) SetAxisZ(axis QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DSurface_SetAxisZ(ptr.Pointer(), PointerFromQValue3DAxis(axis))
	}
}

func (ptr *Q3DSurface) SetFlipHorizontalGrid(flip bool) {
	if ptr.Pointer() != nil {
		C.Q3DSurface_SetFlipHorizontalGrid(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(flip))))
	}
}

//export callbackQ3DSurface_DestroyQ3DSurface
func callbackQ3DSurface_DestroyQ3DSurface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~Q3DSurface"); signal != nil {
		signal.(func())()
	} else {
		NewQ3DSurfaceFromPointer(ptr).DestroyQ3DSurfaceDefault()
	}
}

func (ptr *Q3DSurface) ConnectDestroyQ3DSurface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~Q3DSurface"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DSurface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DSurface", f)
		}
	}
}

func (ptr *Q3DSurface) DisconnectDestroyQ3DSurface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DSurface")
	}
}

func (ptr *Q3DSurface) DestroyQ3DSurface() {
	if ptr.Pointer() != nil {
		C.Q3DSurface_DestroyQ3DSurface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DSurface) DestroyQ3DSurfaceDefault() {
	if ptr.Pointer() != nil {
		C.Q3DSurface_DestroyQ3DSurfaceDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DSurface) SeriesList() []*QSurface3DSeries {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []*QSurface3DSeries {
			var out = make([]*QSurface3DSeries, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQ3DSurfaceFromPointer(l.data).__seriesList_atList(i)
			}
			return out
		}(C.Q3DSurface_SeriesList(ptr.Pointer()))
	}
	return make([]*QSurface3DSeries, 0)
}

func (ptr *Q3DSurface) Axes() []*QValue3DAxis {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []*QValue3DAxis {
			var out = make([]*QValue3DAxis, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQ3DSurfaceFromPointer(l.data).__axes_atList(i)
			}
			return out
		}(C.Q3DSurface_Axes(ptr.Pointer()))
	}
	return make([]*QValue3DAxis, 0)
}

func (ptr *Q3DSurface) SelectedSeries() *QSurface3DSeries {
	if ptr.Pointer() != nil {
		return NewQSurface3DSeriesFromPointer(C.Q3DSurface_SelectedSeries(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DSurface) AxisX() *QValue3DAxis {
	if ptr.Pointer() != nil {
		return NewQValue3DAxisFromPointer(C.Q3DSurface_AxisX(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DSurface) AxisY() *QValue3DAxis {
	if ptr.Pointer() != nil {
		return NewQValue3DAxisFromPointer(C.Q3DSurface_AxisY(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DSurface) AxisZ() *QValue3DAxis {
	if ptr.Pointer() != nil {
		return NewQValue3DAxisFromPointer(C.Q3DSurface_AxisZ(ptr.Pointer()))
	}
	return nil
}

func (ptr *Q3DSurface) FlipHorizontalGrid() bool {
	if ptr.Pointer() != nil {
		return C.Q3DSurface_FlipHorizontalGrid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DSurface) __seriesList_atList(i int) *QSurface3DSeries {
	if ptr.Pointer() != nil {
		return NewQSurface3DSeriesFromPointer(C.Q3DSurface___seriesList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *Q3DSurface) __seriesList_setList(i QSurface3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DSurface___seriesList_setList(ptr.Pointer(), PointerFromQSurface3DSeries(i))
	}
}

func (ptr *Q3DSurface) __seriesList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Q3DSurface___seriesList_newList(ptr.Pointer()))
}

func (ptr *Q3DSurface) __axes_atList(i int) *QValue3DAxis {
	if ptr.Pointer() != nil {
		return NewQValue3DAxisFromPointer(C.Q3DSurface___axes_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *Q3DSurface) __axes_setList(i QValue3DAxis_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DSurface___axes_setList(ptr.Pointer(), PointerFromQValue3DAxis(i))
	}
}

func (ptr *Q3DSurface) __axes_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Q3DSurface___axes_newList(ptr.Pointer()))
}

type Q3DTheme struct {
	ptr unsafe.Pointer
}

type Q3DTheme_ITF interface {
	Q3DTheme_PTR() *Q3DTheme
}

func (ptr *Q3DTheme) Q3DTheme_PTR() *Q3DTheme {
	return ptr
}

func (ptr *Q3DTheme) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *Q3DTheme) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQ3DTheme(ptr Q3DTheme_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Q3DTheme_PTR().Pointer()
	}
	return nil
}

func NewQ3DThemeFromPointer(ptr unsafe.Pointer) *Q3DTheme {
	var n = new(Q3DTheme)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=Q3DTheme__ColorStyle
//Q3DTheme::ColorStyle
type Q3DTheme__ColorStyle int64

var (
	Q3DTheme__ColorStyleUniform        Q3DTheme__ColorStyle = Q3DTheme__ColorStyle(0)
	Q3DTheme__ColorStyleObjectGradient Q3DTheme__ColorStyle = Q3DTheme__ColorStyle(1)
	Q3DTheme__ColorStyleRangeGradient  Q3DTheme__ColorStyle = Q3DTheme__ColorStyle(2)
)

//go:generate stringer -type=Q3DTheme__Theme
//Q3DTheme::Theme
type Q3DTheme__Theme int64

const (
	Q3DTheme__ThemeQt            Q3DTheme__Theme = Q3DTheme__Theme(0)
	Q3DTheme__ThemePrimaryColors Q3DTheme__Theme = Q3DTheme__Theme(1)
	Q3DTheme__ThemeDigia         Q3DTheme__Theme = Q3DTheme__Theme(2)
	Q3DTheme__ThemeStoneMoss     Q3DTheme__Theme = Q3DTheme__Theme(3)
	Q3DTheme__ThemeArmyBlue      Q3DTheme__Theme = Q3DTheme__Theme(4)
	Q3DTheme__ThemeRetro         Q3DTheme__Theme = Q3DTheme__Theme(5)
	Q3DTheme__ThemeEbony         Q3DTheme__Theme = Q3DTheme__Theme(6)
	Q3DTheme__ThemeIsabelle      Q3DTheme__Theme = Q3DTheme__Theme(7)
	Q3DTheme__ThemeUserDefined   Q3DTheme__Theme = Q3DTheme__Theme(8)
)

func NewQ3DTheme(parent core.QObject_ITF) *Q3DTheme {
	return NewQ3DThemeFromPointer(C.Q3DTheme_NewQ3DTheme(core.PointerFromQObject(parent)))
}

func NewQ3DTheme2(themeType Q3DTheme__Theme, parent core.QObject_ITF) *Q3DTheme {
	return NewQ3DThemeFromPointer(C.Q3DTheme_NewQ3DTheme2(C.longlong(themeType), core.PointerFromQObject(parent)))
}

//export callbackQ3DTheme_AmbientLightStrengthChanged
func callbackQ3DTheme_AmbientLightStrengthChanged(ptr unsafe.Pointer, strength C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "ambientLightStrengthChanged"); signal != nil {
		signal.(func(float32))(float32(strength))
	}

}

func (ptr *Q3DTheme) ConnectAmbientLightStrengthChanged(f func(strength float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "ambientLightStrengthChanged") {
			C.Q3DTheme_ConnectAmbientLightStrengthChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "ambientLightStrengthChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "ambientLightStrengthChanged", func(strength float32) {
				signal.(func(float32))(strength)
				f(strength)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "ambientLightStrengthChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectAmbientLightStrengthChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectAmbientLightStrengthChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "ambientLightStrengthChanged")
	}
}

func (ptr *Q3DTheme) AmbientLightStrengthChanged(strength float32) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_AmbientLightStrengthChanged(ptr.Pointer(), C.float(strength))
	}
}

//export callbackQ3DTheme_BackgroundColorChanged
func callbackQ3DTheme_BackgroundColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "backgroundColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *Q3DTheme) ConnectBackgroundColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "backgroundColorChanged") {
			C.Q3DTheme_ConnectBackgroundColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "backgroundColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "backgroundColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "backgroundColorChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectBackgroundColorChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectBackgroundColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "backgroundColorChanged")
	}
}

func (ptr *Q3DTheme) BackgroundColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_BackgroundColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQ3DTheme_BackgroundEnabledChanged
func callbackQ3DTheme_BackgroundEnabledChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "backgroundEnabledChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *Q3DTheme) ConnectBackgroundEnabledChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "backgroundEnabledChanged") {
			C.Q3DTheme_ConnectBackgroundEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "backgroundEnabledChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "backgroundEnabledChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "backgroundEnabledChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectBackgroundEnabledChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectBackgroundEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "backgroundEnabledChanged")
	}
}

func (ptr *Q3DTheme) BackgroundEnabledChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_BackgroundEnabledChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQ3DTheme_BaseColorsChanged
func callbackQ3DTheme_BaseColorsChanged(ptr unsafe.Pointer, colors C.struct_QtDataVisualization_PackedList) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "baseColorsChanged"); signal != nil {
		signal.(func([]*gui.QColor))(func(l C.struct_QtDataVisualization_PackedList) []*gui.QColor {
			var out = make([]*gui.QColor, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQ3DThemeFromPointer(l.data).__baseColorsChanged_colors_atList(i)
			}
			return out
		}(colors))
	}

}

func (ptr *Q3DTheme) ConnectBaseColorsChanged(f func(colors []*gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "baseColorsChanged") {
			C.Q3DTheme_ConnectBaseColorsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "baseColorsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "baseColorsChanged", func(colors []*gui.QColor) {
				signal.(func([]*gui.QColor))(colors)
				f(colors)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "baseColorsChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectBaseColorsChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectBaseColorsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "baseColorsChanged")
	}
}

func (ptr *Q3DTheme) BaseColorsChanged(colors []*gui.QColor) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_BaseColorsChanged(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQ3DThemeFromPointer(NewQ3DThemeFromPointer(nil).__baseColorsChanged_colors_newList())
			for _, v := range colors {
				tmpList.__baseColorsChanged_colors_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQ3DTheme_BaseGradientsChanged
func callbackQ3DTheme_BaseGradientsChanged(ptr unsafe.Pointer, gradients C.struct_QtDataVisualization_PackedList) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "baseGradientsChanged"); signal != nil {
		signal.(func([]*gui.QLinearGradient))(func(l C.struct_QtDataVisualization_PackedList) []*gui.QLinearGradient {
			var out = make([]*gui.QLinearGradient, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQ3DThemeFromPointer(l.data).__baseGradientsChanged_gradients_atList(i)
			}
			return out
		}(gradients))
	}

}

func (ptr *Q3DTheme) ConnectBaseGradientsChanged(f func(gradients []*gui.QLinearGradient)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "baseGradientsChanged") {
			C.Q3DTheme_ConnectBaseGradientsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "baseGradientsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "baseGradientsChanged", func(gradients []*gui.QLinearGradient) {
				signal.(func([]*gui.QLinearGradient))(gradients)
				f(gradients)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "baseGradientsChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectBaseGradientsChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectBaseGradientsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "baseGradientsChanged")
	}
}

func (ptr *Q3DTheme) BaseGradientsChanged(gradients []*gui.QLinearGradient) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_BaseGradientsChanged(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQ3DThemeFromPointer(NewQ3DThemeFromPointer(nil).__baseGradientsChanged_gradients_newList())
			for _, v := range gradients {
				tmpList.__baseGradientsChanged_gradients_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQ3DTheme_ColorStyleChanged
func callbackQ3DTheme_ColorStyleChanged(ptr unsafe.Pointer, style C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "colorStyleChanged"); signal != nil {
		signal.(func(Q3DTheme__ColorStyle))(Q3DTheme__ColorStyle(style))
	}

}

func (ptr *Q3DTheme) ConnectColorStyleChanged(f func(style Q3DTheme__ColorStyle)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "colorStyleChanged") {
			C.Q3DTheme_ConnectColorStyleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "colorStyleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "colorStyleChanged", func(style Q3DTheme__ColorStyle) {
				signal.(func(Q3DTheme__ColorStyle))(style)
				f(style)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "colorStyleChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectColorStyleChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectColorStyleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "colorStyleChanged")
	}
}

func (ptr *Q3DTheme) ColorStyleChanged(style Q3DTheme__ColorStyle) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_ColorStyleChanged(ptr.Pointer(), C.longlong(style))
	}
}

//export callbackQ3DTheme_FontChanged
func callbackQ3DTheme_FontChanged(ptr unsafe.Pointer, font unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "fontChanged"); signal != nil {
		signal.(func(*gui.QFont))(gui.NewQFontFromPointer(font))
	}

}

func (ptr *Q3DTheme) ConnectFontChanged(f func(font *gui.QFont)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "fontChanged") {
			C.Q3DTheme_ConnectFontChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "fontChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "fontChanged", func(font *gui.QFont) {
				signal.(func(*gui.QFont))(font)
				f(font)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "fontChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectFontChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectFontChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "fontChanged")
	}
}

func (ptr *Q3DTheme) FontChanged(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_FontChanged(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

//export callbackQ3DTheme_GridEnabledChanged
func callbackQ3DTheme_GridEnabledChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "gridEnabledChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *Q3DTheme) ConnectGridEnabledChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "gridEnabledChanged") {
			C.Q3DTheme_ConnectGridEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "gridEnabledChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "gridEnabledChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "gridEnabledChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectGridEnabledChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectGridEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "gridEnabledChanged")
	}
}

func (ptr *Q3DTheme) GridEnabledChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_GridEnabledChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQ3DTheme_GridLineColorChanged
func callbackQ3DTheme_GridLineColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "gridLineColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *Q3DTheme) ConnectGridLineColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "gridLineColorChanged") {
			C.Q3DTheme_ConnectGridLineColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "gridLineColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "gridLineColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "gridLineColorChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectGridLineColorChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectGridLineColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "gridLineColorChanged")
	}
}

func (ptr *Q3DTheme) GridLineColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_GridLineColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQ3DTheme_HighlightLightStrengthChanged
func callbackQ3DTheme_HighlightLightStrengthChanged(ptr unsafe.Pointer, strength C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "highlightLightStrengthChanged"); signal != nil {
		signal.(func(float32))(float32(strength))
	}

}

func (ptr *Q3DTheme) ConnectHighlightLightStrengthChanged(f func(strength float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "highlightLightStrengthChanged") {
			C.Q3DTheme_ConnectHighlightLightStrengthChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "highlightLightStrengthChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "highlightLightStrengthChanged", func(strength float32) {
				signal.(func(float32))(strength)
				f(strength)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "highlightLightStrengthChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectHighlightLightStrengthChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectHighlightLightStrengthChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "highlightLightStrengthChanged")
	}
}

func (ptr *Q3DTheme) HighlightLightStrengthChanged(strength float32) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_HighlightLightStrengthChanged(ptr.Pointer(), C.float(strength))
	}
}

//export callbackQ3DTheme_LabelBackgroundColorChanged
func callbackQ3DTheme_LabelBackgroundColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "labelBackgroundColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *Q3DTheme) ConnectLabelBackgroundColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "labelBackgroundColorChanged") {
			C.Q3DTheme_ConnectLabelBackgroundColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "labelBackgroundColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelBackgroundColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelBackgroundColorChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectLabelBackgroundColorChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectLabelBackgroundColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "labelBackgroundColorChanged")
	}
}

func (ptr *Q3DTheme) LabelBackgroundColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_LabelBackgroundColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQ3DTheme_LabelBackgroundEnabledChanged
func callbackQ3DTheme_LabelBackgroundEnabledChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "labelBackgroundEnabledChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *Q3DTheme) ConnectLabelBackgroundEnabledChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "labelBackgroundEnabledChanged") {
			C.Q3DTheme_ConnectLabelBackgroundEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "labelBackgroundEnabledChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelBackgroundEnabledChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelBackgroundEnabledChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectLabelBackgroundEnabledChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectLabelBackgroundEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "labelBackgroundEnabledChanged")
	}
}

func (ptr *Q3DTheme) LabelBackgroundEnabledChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_LabelBackgroundEnabledChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQ3DTheme_LabelBorderEnabledChanged
func callbackQ3DTheme_LabelBorderEnabledChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "labelBorderEnabledChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *Q3DTheme) ConnectLabelBorderEnabledChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "labelBorderEnabledChanged") {
			C.Q3DTheme_ConnectLabelBorderEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "labelBorderEnabledChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelBorderEnabledChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelBorderEnabledChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectLabelBorderEnabledChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectLabelBorderEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "labelBorderEnabledChanged")
	}
}

func (ptr *Q3DTheme) LabelBorderEnabledChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_LabelBorderEnabledChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQ3DTheme_LabelTextColorChanged
func callbackQ3DTheme_LabelTextColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "labelTextColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *Q3DTheme) ConnectLabelTextColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "labelTextColorChanged") {
			C.Q3DTheme_ConnectLabelTextColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "labelTextColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelTextColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelTextColorChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectLabelTextColorChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectLabelTextColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "labelTextColorChanged")
	}
}

func (ptr *Q3DTheme) LabelTextColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_LabelTextColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQ3DTheme_LightColorChanged
func callbackQ3DTheme_LightColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "lightColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *Q3DTheme) ConnectLightColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "lightColorChanged") {
			C.Q3DTheme_ConnectLightColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "lightColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "lightColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "lightColorChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectLightColorChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectLightColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "lightColorChanged")
	}
}

func (ptr *Q3DTheme) LightColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_LightColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQ3DTheme_LightStrengthChanged
func callbackQ3DTheme_LightStrengthChanged(ptr unsafe.Pointer, strength C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "lightStrengthChanged"); signal != nil {
		signal.(func(float32))(float32(strength))
	}

}

func (ptr *Q3DTheme) ConnectLightStrengthChanged(f func(strength float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "lightStrengthChanged") {
			C.Q3DTheme_ConnectLightStrengthChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "lightStrengthChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "lightStrengthChanged", func(strength float32) {
				signal.(func(float32))(strength)
				f(strength)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "lightStrengthChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectLightStrengthChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectLightStrengthChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "lightStrengthChanged")
	}
}

func (ptr *Q3DTheme) LightStrengthChanged(strength float32) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_LightStrengthChanged(ptr.Pointer(), C.float(strength))
	}
}

//export callbackQ3DTheme_MultiHighlightColorChanged
func callbackQ3DTheme_MultiHighlightColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "multiHighlightColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *Q3DTheme) ConnectMultiHighlightColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightColorChanged") {
			C.Q3DTheme_ConnectMultiHighlightColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightColorChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectMultiHighlightColorChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectMultiHighlightColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightColorChanged")
	}
}

func (ptr *Q3DTheme) MultiHighlightColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_MultiHighlightColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQ3DTheme_MultiHighlightGradientChanged
func callbackQ3DTheme_MultiHighlightGradientChanged(ptr unsafe.Pointer, gradient unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "multiHighlightGradientChanged"); signal != nil {
		signal.(func(*gui.QLinearGradient))(gui.NewQLinearGradientFromPointer(gradient))
	}

}

func (ptr *Q3DTheme) ConnectMultiHighlightGradientChanged(f func(gradient *gui.QLinearGradient)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightGradientChanged") {
			C.Q3DTheme_ConnectMultiHighlightGradientChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightGradientChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightGradientChanged", func(gradient *gui.QLinearGradient) {
				signal.(func(*gui.QLinearGradient))(gradient)
				f(gradient)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightGradientChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectMultiHighlightGradientChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectMultiHighlightGradientChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightGradientChanged")
	}
}

func (ptr *Q3DTheme) MultiHighlightGradientChanged(gradient gui.QLinearGradient_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_MultiHighlightGradientChanged(ptr.Pointer(), gui.PointerFromQLinearGradient(gradient))
	}
}

func (ptr *Q3DTheme) SetAmbientLightStrength(strength float32) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetAmbientLightStrength(ptr.Pointer(), C.float(strength))
	}
}

func (ptr *Q3DTheme) SetBackgroundColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetBackgroundColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *Q3DTheme) SetBackgroundEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetBackgroundEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *Q3DTheme) SetBaseColors(colors []*gui.QColor) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetBaseColors(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQ3DThemeFromPointer(NewQ3DThemeFromPointer(nil).__setBaseColors_colors_newList())
			for _, v := range colors {
				tmpList.__setBaseColors_colors_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *Q3DTheme) SetBaseGradients(gradients []*gui.QLinearGradient) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetBaseGradients(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQ3DThemeFromPointer(NewQ3DThemeFromPointer(nil).__setBaseGradients_gradients_newList())
			for _, v := range gradients {
				tmpList.__setBaseGradients_gradients_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *Q3DTheme) SetColorStyle(style Q3DTheme__ColorStyle) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetColorStyle(ptr.Pointer(), C.longlong(style))
	}
}

func (ptr *Q3DTheme) SetFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *Q3DTheme) SetGridEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetGridEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *Q3DTheme) SetGridLineColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetGridLineColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *Q3DTheme) SetHighlightLightStrength(strength float32) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetHighlightLightStrength(ptr.Pointer(), C.float(strength))
	}
}

func (ptr *Q3DTheme) SetLabelBackgroundColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetLabelBackgroundColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *Q3DTheme) SetLabelBackgroundEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetLabelBackgroundEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *Q3DTheme) SetLabelBorderEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetLabelBorderEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *Q3DTheme) SetLabelTextColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetLabelTextColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *Q3DTheme) SetLightColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetLightColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *Q3DTheme) SetLightStrength(strength float32) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetLightStrength(ptr.Pointer(), C.float(strength))
	}
}

func (ptr *Q3DTheme) SetMultiHighlightColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetMultiHighlightColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *Q3DTheme) SetMultiHighlightGradient(gradient gui.QLinearGradient_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetMultiHighlightGradient(ptr.Pointer(), gui.PointerFromQLinearGradient(gradient))
	}
}

func (ptr *Q3DTheme) SetSingleHighlightColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetSingleHighlightColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *Q3DTheme) SetSingleHighlightGradient(gradient gui.QLinearGradient_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetSingleHighlightGradient(ptr.Pointer(), gui.PointerFromQLinearGradient(gradient))
	}
}

func (ptr *Q3DTheme) SetType(themeType Q3DTheme__Theme) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetType(ptr.Pointer(), C.longlong(themeType))
	}
}

func (ptr *Q3DTheme) SetWindowColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SetWindowColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQ3DTheme_SingleHighlightColorChanged
func callbackQ3DTheme_SingleHighlightColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "singleHighlightColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *Q3DTheme) ConnectSingleHighlightColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightColorChanged") {
			C.Q3DTheme_ConnectSingleHighlightColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightColorChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectSingleHighlightColorChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectSingleHighlightColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightColorChanged")
	}
}

func (ptr *Q3DTheme) SingleHighlightColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SingleHighlightColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQ3DTheme_SingleHighlightGradientChanged
func callbackQ3DTheme_SingleHighlightGradientChanged(ptr unsafe.Pointer, gradient unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "singleHighlightGradientChanged"); signal != nil {
		signal.(func(*gui.QLinearGradient))(gui.NewQLinearGradientFromPointer(gradient))
	}

}

func (ptr *Q3DTheme) ConnectSingleHighlightGradientChanged(f func(gradient *gui.QLinearGradient)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightGradientChanged") {
			C.Q3DTheme_ConnectSingleHighlightGradientChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightGradientChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightGradientChanged", func(gradient *gui.QLinearGradient) {
				signal.(func(*gui.QLinearGradient))(gradient)
				f(gradient)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightGradientChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectSingleHighlightGradientChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectSingleHighlightGradientChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightGradientChanged")
	}
}

func (ptr *Q3DTheme) SingleHighlightGradientChanged(gradient gui.QLinearGradient_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_SingleHighlightGradientChanged(ptr.Pointer(), gui.PointerFromQLinearGradient(gradient))
	}
}

//export callbackQ3DTheme_TypeChanged
func callbackQ3DTheme_TypeChanged(ptr unsafe.Pointer, themeType C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "typeChanged"); signal != nil {
		signal.(func(Q3DTheme__Theme))(Q3DTheme__Theme(themeType))
	}

}

func (ptr *Q3DTheme) ConnectTypeChanged(f func(themeType Q3DTheme__Theme)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "typeChanged") {
			C.Q3DTheme_ConnectTypeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "typeChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "typeChanged", func(themeType Q3DTheme__Theme) {
				signal.(func(Q3DTheme__Theme))(themeType)
				f(themeType)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "typeChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectTypeChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "typeChanged")
	}
}

func (ptr *Q3DTheme) TypeChanged(themeType Q3DTheme__Theme) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_TypeChanged(ptr.Pointer(), C.longlong(themeType))
	}
}

//export callbackQ3DTheme_WindowColorChanged
func callbackQ3DTheme_WindowColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "windowColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *Q3DTheme) ConnectWindowColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "windowColorChanged") {
			C.Q3DTheme_ConnectWindowColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "windowColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "windowColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "windowColorChanged", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectWindowColorChanged() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DisconnectWindowColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "windowColorChanged")
	}
}

func (ptr *Q3DTheme) WindowColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme_WindowColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQ3DTheme_DestroyQ3DTheme
func callbackQ3DTheme_DestroyQ3DTheme(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~Q3DTheme"); signal != nil {
		signal.(func())()
	} else {
		NewQ3DThemeFromPointer(ptr).DestroyQ3DThemeDefault()
	}
}

func (ptr *Q3DTheme) ConnectDestroyQ3DTheme(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~Q3DTheme"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DTheme", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DTheme", f)
		}
	}
}

func (ptr *Q3DTheme) DisconnectDestroyQ3DTheme() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~Q3DTheme")
	}
}

func (ptr *Q3DTheme) DestroyQ3DTheme() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DestroyQ3DTheme(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DTheme) DestroyQ3DThemeDefault() {
	if ptr.Pointer() != nil {
		C.Q3DTheme_DestroyQ3DThemeDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *Q3DTheme) ColorStyle() Q3DTheme__ColorStyle {
	if ptr.Pointer() != nil {
		return Q3DTheme__ColorStyle(C.Q3DTheme_ColorStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *Q3DTheme) BackgroundColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.Q3DTheme_BackgroundColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) GridLineColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.Q3DTheme_GridLineColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) LabelBackgroundColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.Q3DTheme_LabelBackgroundColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) LabelTextColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.Q3DTheme_LabelTextColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) LightColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.Q3DTheme_LightColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) MultiHighlightColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.Q3DTheme_MultiHighlightColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) SingleHighlightColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.Q3DTheme_SingleHighlightColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) WindowColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.Q3DTheme_WindowColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) Font() *gui.QFont {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQFontFromPointer(C.Q3DTheme_Font(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) MultiHighlightGradient() *gui.QLinearGradient {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQLinearGradientFromPointer(C.Q3DTheme_MultiHighlightGradient(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QLinearGradient).DestroyQLinearGradient)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) SingleHighlightGradient() *gui.QLinearGradient {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQLinearGradientFromPointer(C.Q3DTheme_SingleHighlightGradient(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QLinearGradient).DestroyQLinearGradient)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) BaseColors() []*gui.QColor {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []*gui.QColor {
			var out = make([]*gui.QColor, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQ3DThemeFromPointer(l.data).__baseColors_atList(i)
			}
			return out
		}(C.Q3DTheme_BaseColors(ptr.Pointer()))
	}
	return make([]*gui.QColor, 0)
}

func (ptr *Q3DTheme) BaseGradients() []*gui.QLinearGradient {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []*gui.QLinearGradient {
			var out = make([]*gui.QLinearGradient, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQ3DThemeFromPointer(l.data).__baseGradients_atList(i)
			}
			return out
		}(C.Q3DTheme_BaseGradients(ptr.Pointer()))
	}
	return make([]*gui.QLinearGradient, 0)
}

func (ptr *Q3DTheme) Type() Q3DTheme__Theme {
	if ptr.Pointer() != nil {
		return Q3DTheme__Theme(C.Q3DTheme_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *Q3DTheme) IsBackgroundEnabled() bool {
	if ptr.Pointer() != nil {
		return C.Q3DTheme_IsBackgroundEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DTheme) IsGridEnabled() bool {
	if ptr.Pointer() != nil {
		return C.Q3DTheme_IsGridEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DTheme) IsLabelBackgroundEnabled() bool {
	if ptr.Pointer() != nil {
		return C.Q3DTheme_IsLabelBackgroundEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DTheme) IsLabelBorderEnabled() bool {
	if ptr.Pointer() != nil {
		return C.Q3DTheme_IsLabelBorderEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *Q3DTheme) AmbientLightStrength() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Q3DTheme_AmbientLightStrength(ptr.Pointer()))
	}
	return 0
}

func (ptr *Q3DTheme) HighlightLightStrength() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Q3DTheme_HighlightLightStrength(ptr.Pointer()))
	}
	return 0
}

func (ptr *Q3DTheme) LightStrength() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Q3DTheme_LightStrength(ptr.Pointer()))
	}
	return 0
}

func (ptr *Q3DTheme) __baseColorsChanged_colors_atList(i int) *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.Q3DTheme___baseColorsChanged_colors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) __baseColorsChanged_colors_setList(i gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme___baseColorsChanged_colors_setList(ptr.Pointer(), gui.PointerFromQColor(i))
	}
}

func (ptr *Q3DTheme) __baseColorsChanged_colors_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Q3DTheme___baseColorsChanged_colors_newList(ptr.Pointer()))
}

func (ptr *Q3DTheme) __baseGradientsChanged_gradients_atList(i int) *gui.QLinearGradient {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQLinearGradientFromPointer(C.Q3DTheme___baseGradientsChanged_gradients_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*gui.QLinearGradient).DestroyQLinearGradient)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) __baseGradientsChanged_gradients_setList(i gui.QLinearGradient_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme___baseGradientsChanged_gradients_setList(ptr.Pointer(), gui.PointerFromQLinearGradient(i))
	}
}

func (ptr *Q3DTheme) __baseGradientsChanged_gradients_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Q3DTheme___baseGradientsChanged_gradients_newList(ptr.Pointer()))
}

func (ptr *Q3DTheme) __setBaseColors_colors_atList(i int) *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.Q3DTheme___setBaseColors_colors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) __setBaseColors_colors_setList(i gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme___setBaseColors_colors_setList(ptr.Pointer(), gui.PointerFromQColor(i))
	}
}

func (ptr *Q3DTheme) __setBaseColors_colors_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Q3DTheme___setBaseColors_colors_newList(ptr.Pointer()))
}

func (ptr *Q3DTheme) __setBaseGradients_gradients_atList(i int) *gui.QLinearGradient {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQLinearGradientFromPointer(C.Q3DTheme___setBaseGradients_gradients_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*gui.QLinearGradient).DestroyQLinearGradient)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) __setBaseGradients_gradients_setList(i gui.QLinearGradient_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme___setBaseGradients_gradients_setList(ptr.Pointer(), gui.PointerFromQLinearGradient(i))
	}
}

func (ptr *Q3DTheme) __setBaseGradients_gradients_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Q3DTheme___setBaseGradients_gradients_newList(ptr.Pointer()))
}

func (ptr *Q3DTheme) __baseColors_atList(i int) *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.Q3DTheme___baseColors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) __baseColors_setList(i gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme___baseColors_setList(ptr.Pointer(), gui.PointerFromQColor(i))
	}
}

func (ptr *Q3DTheme) __baseColors_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Q3DTheme___baseColors_newList(ptr.Pointer()))
}

func (ptr *Q3DTheme) __baseGradients_atList(i int) *gui.QLinearGradient {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQLinearGradientFromPointer(C.Q3DTheme___baseGradients_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*gui.QLinearGradient).DestroyQLinearGradient)
		return tmpValue
	}
	return nil
}

func (ptr *Q3DTheme) __baseGradients_setList(i gui.QLinearGradient_ITF) {
	if ptr.Pointer() != nil {
		C.Q3DTheme___baseGradients_setList(ptr.Pointer(), gui.PointerFromQLinearGradient(i))
	}
}

func (ptr *Q3DTheme) __baseGradients_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Q3DTheme___baseGradients_newList(ptr.Pointer()))
}

type QAbstract3DAxis struct {
	ptr unsafe.Pointer
}

type QAbstract3DAxis_ITF interface {
	QAbstract3DAxis_PTR() *QAbstract3DAxis
}

func (ptr *QAbstract3DAxis) QAbstract3DAxis_PTR() *QAbstract3DAxis {
	return ptr
}

func (ptr *QAbstract3DAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstract3DAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAbstract3DAxis(ptr QAbstract3DAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstract3DAxis_PTR().Pointer()
	}
	return nil
}

func NewQAbstract3DAxisFromPointer(ptr unsafe.Pointer) *QAbstract3DAxis {
	var n = new(QAbstract3DAxis)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QAbstract3DAxis__AxisOrientation
//QAbstract3DAxis::AxisOrientation
type QAbstract3DAxis__AxisOrientation int64

const (
	QAbstract3DAxis__AxisOrientationNone QAbstract3DAxis__AxisOrientation = QAbstract3DAxis__AxisOrientation(0)
	QAbstract3DAxis__AxisOrientationX    QAbstract3DAxis__AxisOrientation = QAbstract3DAxis__AxisOrientation(1)
	QAbstract3DAxis__AxisOrientationY    QAbstract3DAxis__AxisOrientation = QAbstract3DAxis__AxisOrientation(2)
	QAbstract3DAxis__AxisOrientationZ    QAbstract3DAxis__AxisOrientation = QAbstract3DAxis__AxisOrientation(4)
)

//go:generate stringer -type=QAbstract3DAxis__AxisType
//QAbstract3DAxis::AxisType
type QAbstract3DAxis__AxisType int64

const (
	QAbstract3DAxis__AxisTypeNone     QAbstract3DAxis__AxisType = QAbstract3DAxis__AxisType(0)
	QAbstract3DAxis__AxisTypeCategory QAbstract3DAxis__AxisType = QAbstract3DAxis__AxisType(1)
	QAbstract3DAxis__AxisTypeValue    QAbstract3DAxis__AxisType = QAbstract3DAxis__AxisType(2)
)

//export callbackQAbstract3DAxis_RangeChanged
func callbackQAbstract3DAxis_RangeChanged(ptr unsafe.Pointer, min C.float, max C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rangeChanged"); signal != nil {
		signal.(func(float32, float32))(float32(min), float32(max))
	}

}

func (ptr *QAbstract3DAxis) ConnectRangeChanged(f func(min float32, max float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rangeChanged") {
			C.QAbstract3DAxis_ConnectRangeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rangeChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rangeChanged", func(min float32, max float32) {
				signal.(func(float32, float32))(min, max)
				f(min, max)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rangeChanged", f)
		}
	}
}

func (ptr *QAbstract3DAxis) DisconnectRangeChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_DisconnectRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rangeChanged")
	}
}

func (ptr *QAbstract3DAxis) RangeChanged(min float32, max float32) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_RangeChanged(ptr.Pointer(), C.float(min), C.float(max))
	}
}

func (ptr *QAbstract3DAxis) SetTitleFixed(fixed bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_SetTitleFixed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(fixed))))
	}
}

//export callbackQAbstract3DAxis_AutoAdjustRangeChanged
func callbackQAbstract3DAxis_AutoAdjustRangeChanged(ptr unsafe.Pointer, autoAdjust C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "autoAdjustRangeChanged"); signal != nil {
		signal.(func(bool))(int8(autoAdjust) != 0)
	}

}

func (ptr *QAbstract3DAxis) ConnectAutoAdjustRangeChanged(f func(autoAdjust bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "autoAdjustRangeChanged") {
			C.QAbstract3DAxis_ConnectAutoAdjustRangeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "autoAdjustRangeChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "autoAdjustRangeChanged", func(autoAdjust bool) {
				signal.(func(bool))(autoAdjust)
				f(autoAdjust)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "autoAdjustRangeChanged", f)
		}
	}
}

func (ptr *QAbstract3DAxis) DisconnectAutoAdjustRangeChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_DisconnectAutoAdjustRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "autoAdjustRangeChanged")
	}
}

func (ptr *QAbstract3DAxis) AutoAdjustRangeChanged(autoAdjust bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_AutoAdjustRangeChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(autoAdjust))))
	}
}

//export callbackQAbstract3DAxis_LabelAutoRotationChanged
func callbackQAbstract3DAxis_LabelAutoRotationChanged(ptr unsafe.Pointer, angle C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "labelAutoRotationChanged"); signal != nil {
		signal.(func(float32))(float32(angle))
	}

}

func (ptr *QAbstract3DAxis) ConnectLabelAutoRotationChanged(f func(angle float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "labelAutoRotationChanged") {
			C.QAbstract3DAxis_ConnectLabelAutoRotationChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "labelAutoRotationChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelAutoRotationChanged", func(angle float32) {
				signal.(func(float32))(angle)
				f(angle)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelAutoRotationChanged", f)
		}
	}
}

func (ptr *QAbstract3DAxis) DisconnectLabelAutoRotationChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_DisconnectLabelAutoRotationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "labelAutoRotationChanged")
	}
}

func (ptr *QAbstract3DAxis) LabelAutoRotationChanged(angle float32) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_LabelAutoRotationChanged(ptr.Pointer(), C.float(angle))
	}
}

//export callbackQAbstract3DAxis_LabelsChanged
func callbackQAbstract3DAxis_LabelsChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "labelsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstract3DAxis) ConnectLabelsChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "labelsChanged") {
			C.QAbstract3DAxis_ConnectLabelsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "labelsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelsChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelsChanged", f)
		}
	}
}

func (ptr *QAbstract3DAxis) DisconnectLabelsChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_DisconnectLabelsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "labelsChanged")
	}
}

func (ptr *QAbstract3DAxis) LabelsChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_LabelsChanged(ptr.Pointer())
	}
}

//export callbackQAbstract3DAxis_MaxChanged
func callbackQAbstract3DAxis_MaxChanged(ptr unsafe.Pointer, value C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "maxChanged"); signal != nil {
		signal.(func(float32))(float32(value))
	}

}

func (ptr *QAbstract3DAxis) ConnectMaxChanged(f func(value float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "maxChanged") {
			C.QAbstract3DAxis_ConnectMaxChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "maxChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "maxChanged", func(value float32) {
				signal.(func(float32))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "maxChanged", f)
		}
	}
}

func (ptr *QAbstract3DAxis) DisconnectMaxChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_DisconnectMaxChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "maxChanged")
	}
}

func (ptr *QAbstract3DAxis) MaxChanged(value float32) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_MaxChanged(ptr.Pointer(), C.float(value))
	}
}

//export callbackQAbstract3DAxis_MinChanged
func callbackQAbstract3DAxis_MinChanged(ptr unsafe.Pointer, value C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "minChanged"); signal != nil {
		signal.(func(float32))(float32(value))
	}

}

func (ptr *QAbstract3DAxis) ConnectMinChanged(f func(value float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "minChanged") {
			C.QAbstract3DAxis_ConnectMinChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "minChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "minChanged", func(value float32) {
				signal.(func(float32))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "minChanged", f)
		}
	}
}

func (ptr *QAbstract3DAxis) DisconnectMinChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_DisconnectMinChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "minChanged")
	}
}

func (ptr *QAbstract3DAxis) MinChanged(value float32) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_MinChanged(ptr.Pointer(), C.float(value))
	}
}

//export callbackQAbstract3DAxis_OrientationChanged
func callbackQAbstract3DAxis_OrientationChanged(ptr unsafe.Pointer, orientation C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "orientationChanged"); signal != nil {
		signal.(func(QAbstract3DAxis__AxisOrientation))(QAbstract3DAxis__AxisOrientation(orientation))
	}

}

func (ptr *QAbstract3DAxis) ConnectOrientationChanged(f func(orientation QAbstract3DAxis__AxisOrientation)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "orientationChanged") {
			C.QAbstract3DAxis_ConnectOrientationChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "orientationChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "orientationChanged", func(orientation QAbstract3DAxis__AxisOrientation) {
				signal.(func(QAbstract3DAxis__AxisOrientation))(orientation)
				f(orientation)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "orientationChanged", f)
		}
	}
}

func (ptr *QAbstract3DAxis) DisconnectOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_DisconnectOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "orientationChanged")
	}
}

func (ptr *QAbstract3DAxis) OrientationChanged(orientation QAbstract3DAxis__AxisOrientation) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_OrientationChanged(ptr.Pointer(), C.longlong(orientation))
	}
}

func (ptr *QAbstract3DAxis) SetAutoAdjustRange(autoAdjust bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_SetAutoAdjustRange(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(autoAdjust))))
	}
}

func (ptr *QAbstract3DAxis) SetLabelAutoRotation(angle float32) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_SetLabelAutoRotation(ptr.Pointer(), C.float(angle))
	}
}

func (ptr *QAbstract3DAxis) SetLabels(labels []string) {
	if ptr.Pointer() != nil {
		var labelsC = C.CString(strings.Join(labels, "|"))
		defer C.free(unsafe.Pointer(labelsC))
		C.QAbstract3DAxis_SetLabels(ptr.Pointer(), labelsC)
	}
}

func (ptr *QAbstract3DAxis) SetMax(max float32) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_SetMax(ptr.Pointer(), C.float(max))
	}
}

func (ptr *QAbstract3DAxis) SetMin(min float32) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_SetMin(ptr.Pointer(), C.float(min))
	}
}

func (ptr *QAbstract3DAxis) SetRange(min float32, max float32) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_SetRange(ptr.Pointer(), C.float(min), C.float(max))
	}
}

func (ptr *QAbstract3DAxis) SetTitle(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QAbstract3DAxis_SetTitle(ptr.Pointer(), titleC)
	}
}

func (ptr *QAbstract3DAxis) SetTitleVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_SetTitleVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstract3DAxis_TitleChanged
func callbackQAbstract3DAxis_TitleChanged(ptr unsafe.Pointer, newTitle C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "titleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(newTitle))
	}

}

func (ptr *QAbstract3DAxis) ConnectTitleChanged(f func(newTitle string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "titleChanged") {
			C.QAbstract3DAxis_ConnectTitleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "titleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "titleChanged", func(newTitle string) {
				signal.(func(string))(newTitle)
				f(newTitle)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "titleChanged", f)
		}
	}
}

func (ptr *QAbstract3DAxis) DisconnectTitleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "titleChanged")
	}
}

func (ptr *QAbstract3DAxis) TitleChanged(newTitle string) {
	if ptr.Pointer() != nil {
		var newTitleC *C.char
		if newTitle != "" {
			newTitleC = C.CString(newTitle)
			defer C.free(unsafe.Pointer(newTitleC))
		}
		C.QAbstract3DAxis_TitleChanged(ptr.Pointer(), newTitleC)
	}
}

//export callbackQAbstract3DAxis_TitleFixedChanged
func callbackQAbstract3DAxis_TitleFixedChanged(ptr unsafe.Pointer, fixed C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "titleFixedChanged"); signal != nil {
		signal.(func(bool))(int8(fixed) != 0)
	}

}

func (ptr *QAbstract3DAxis) ConnectTitleFixedChanged(f func(fixed bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "titleFixedChanged") {
			C.QAbstract3DAxis_ConnectTitleFixedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "titleFixedChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "titleFixedChanged", func(fixed bool) {
				signal.(func(bool))(fixed)
				f(fixed)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "titleFixedChanged", f)
		}
	}
}

func (ptr *QAbstract3DAxis) DisconnectTitleFixedChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_DisconnectTitleFixedChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "titleFixedChanged")
	}
}

func (ptr *QAbstract3DAxis) TitleFixedChanged(fixed bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_TitleFixedChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(fixed))))
	}
}

//export callbackQAbstract3DAxis_TitleVisibilityChanged
func callbackQAbstract3DAxis_TitleVisibilityChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "titleVisibilityChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QAbstract3DAxis) ConnectTitleVisibilityChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "titleVisibilityChanged") {
			C.QAbstract3DAxis_ConnectTitleVisibilityChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "titleVisibilityChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "titleVisibilityChanged", func(visible bool) {
				signal.(func(bool))(visible)
				f(visible)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "titleVisibilityChanged", f)
		}
	}
}

func (ptr *QAbstract3DAxis) DisconnectTitleVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_DisconnectTitleVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "titleVisibilityChanged")
	}
}

func (ptr *QAbstract3DAxis) TitleVisibilityChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_TitleVisibilityChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstract3DAxis_DestroyQAbstract3DAxis
func callbackQAbstract3DAxis_DestroyQAbstract3DAxis(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QAbstract3DAxis"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstract3DAxisFromPointer(ptr).DestroyQAbstract3DAxisDefault()
	}
}

func (ptr *QAbstract3DAxis) ConnectDestroyQAbstract3DAxis(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DAxis"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DAxis", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DAxis", f)
		}
	}
}

func (ptr *QAbstract3DAxis) DisconnectDestroyQAbstract3DAxis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DAxis")
	}
}

func (ptr *QAbstract3DAxis) DestroyQAbstract3DAxis() {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_DestroyQAbstract3DAxis(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstract3DAxis) DestroyQAbstract3DAxisDefault() {
	if ptr.Pointer() != nil {
		C.QAbstract3DAxis_DestroyQAbstract3DAxisDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstract3DAxis) Orientation() QAbstract3DAxis__AxisOrientation {
	if ptr.Pointer() != nil {
		return QAbstract3DAxis__AxisOrientation(C.QAbstract3DAxis_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DAxis) Type() QAbstract3DAxis__AxisType {
	if ptr.Pointer() != nil {
		return QAbstract3DAxis__AxisType(C.QAbstract3DAxis_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DAxis) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstract3DAxis_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstract3DAxis) Labels() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QAbstract3DAxis_Labels(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QAbstract3DAxis) IsAutoAdjustRange() bool {
	if ptr.Pointer() != nil {
		return C.QAbstract3DAxis_IsAutoAdjustRange(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstract3DAxis) IsTitleFixed() bool {
	if ptr.Pointer() != nil {
		return C.QAbstract3DAxis_IsTitleFixed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstract3DAxis) IsTitleVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAbstract3DAxis_IsTitleVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstract3DAxis) LabelAutoRotation() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QAbstract3DAxis_LabelAutoRotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DAxis) Max() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QAbstract3DAxis_Max(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DAxis) Min() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QAbstract3DAxis_Min(ptr.Pointer()))
	}
	return 0
}

type QAbstract3DGraph struct {
	ptr unsafe.Pointer
}

type QAbstract3DGraph_ITF interface {
	QAbstract3DGraph_PTR() *QAbstract3DGraph
}

func (ptr *QAbstract3DGraph) QAbstract3DGraph_PTR() *QAbstract3DGraph {
	return ptr
}

func (ptr *QAbstract3DGraph) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstract3DGraph) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAbstract3DGraph(ptr QAbstract3DGraph_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstract3DGraph_PTR().Pointer()
	}
	return nil
}

func NewQAbstract3DGraphFromPointer(ptr unsafe.Pointer) *QAbstract3DGraph {
	var n = new(QAbstract3DGraph)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QAbstract3DGraph__ElementType
//QAbstract3DGraph::ElementType
type QAbstract3DGraph__ElementType int64

const (
	QAbstract3DGraph__ElementNone       QAbstract3DGraph__ElementType = QAbstract3DGraph__ElementType(0)
	QAbstract3DGraph__ElementSeries     QAbstract3DGraph__ElementType = QAbstract3DGraph__ElementType(1)
	QAbstract3DGraph__ElementAxisXLabel QAbstract3DGraph__ElementType = QAbstract3DGraph__ElementType(2)
	QAbstract3DGraph__ElementAxisYLabel QAbstract3DGraph__ElementType = QAbstract3DGraph__ElementType(3)
	QAbstract3DGraph__ElementAxisZLabel QAbstract3DGraph__ElementType = QAbstract3DGraph__ElementType(4)
	QAbstract3DGraph__ElementCustomItem QAbstract3DGraph__ElementType = QAbstract3DGraph__ElementType(5)
)

//go:generate stringer -type=QAbstract3DGraph__OptimizationHint
//QAbstract3DGraph::OptimizationHint
type QAbstract3DGraph__OptimizationHint int64

const (
	QAbstract3DGraph__OptimizationDefault QAbstract3DGraph__OptimizationHint = QAbstract3DGraph__OptimizationHint(0)
	QAbstract3DGraph__OptimizationStatic  QAbstract3DGraph__OptimizationHint = QAbstract3DGraph__OptimizationHint(1)
)

//go:generate stringer -type=QAbstract3DGraph__SelectionFlag
//QAbstract3DGraph::SelectionFlag
type QAbstract3DGraph__SelectionFlag int64

const (
	QAbstract3DGraph__SelectionNone             QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(0)
	QAbstract3DGraph__SelectionItem             QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(1)
	QAbstract3DGraph__SelectionRow              QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(2)
	QAbstract3DGraph__SelectionItemAndRow       QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(QAbstract3DGraph__SelectionItem | QAbstract3DGraph__SelectionRow)
	QAbstract3DGraph__SelectionColumn           QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(4)
	QAbstract3DGraph__SelectionItemAndColumn    QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(QAbstract3DGraph__SelectionItem | QAbstract3DGraph__SelectionColumn)
	QAbstract3DGraph__SelectionRowAndColumn     QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(QAbstract3DGraph__SelectionRow | QAbstract3DGraph__SelectionColumn)
	QAbstract3DGraph__SelectionItemRowAndColumn QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(QAbstract3DGraph__SelectionItem | QAbstract3DGraph__SelectionRow | QAbstract3DGraph__SelectionColumn)
	QAbstract3DGraph__SelectionSlice            QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(8)
	QAbstract3DGraph__SelectionMultiSeries      QAbstract3DGraph__SelectionFlag = QAbstract3DGraph__SelectionFlag(16)
)

//go:generate stringer -type=QAbstract3DGraph__ShadowQuality
//QAbstract3DGraph::ShadowQuality
type QAbstract3DGraph__ShadowQuality int64

const (
	QAbstract3DGraph__ShadowQualityNone       QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(0)
	QAbstract3DGraph__ShadowQualityLow        QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(1)
	QAbstract3DGraph__ShadowQualityMedium     QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(2)
	QAbstract3DGraph__ShadowQualityHigh       QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(3)
	QAbstract3DGraph__ShadowQualitySoftLow    QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(4)
	QAbstract3DGraph__ShadowQualitySoftMedium QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(5)
	QAbstract3DGraph__ShadowQualitySoftHigh   QAbstract3DGraph__ShadowQuality = QAbstract3DGraph__ShadowQuality(6)
)

func (ptr *QAbstract3DGraph) RenderToImage(msaaSamples int, imageSize core.QSize_ITF) *gui.QImage {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQImageFromPointer(C.QAbstract3DGraph_RenderToImage(ptr.Pointer(), C.int(int32(msaaSamples)), core.PointerFromQSize(imageSize)))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstract3DGraph) AddCustomItem(item QCustom3DItem_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstract3DGraph_AddCustomItem(ptr.Pointer(), PointerFromQCustom3DItem(item))))
	}
	return 0
}

//export callbackQAbstract3DGraph_ActiveInputHandlerChanged
func callbackQAbstract3DGraph_ActiveInputHandlerChanged(ptr unsafe.Pointer, inputHandler unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "activeInputHandlerChanged"); signal != nil {
		signal.(func(*QAbstract3DInputHandler))(NewQAbstract3DInputHandlerFromPointer(inputHandler))
	}

}

func (ptr *QAbstract3DGraph) ConnectActiveInputHandlerChanged(f func(inputHandler *QAbstract3DInputHandler)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "activeInputHandlerChanged") {
			C.QAbstract3DGraph_ConnectActiveInputHandlerChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "activeInputHandlerChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "activeInputHandlerChanged", func(inputHandler *QAbstract3DInputHandler) {
				signal.(func(*QAbstract3DInputHandler))(inputHandler)
				f(inputHandler)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "activeInputHandlerChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectActiveInputHandlerChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectActiveInputHandlerChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "activeInputHandlerChanged")
	}
}

func (ptr *QAbstract3DGraph) ActiveInputHandlerChanged(inputHandler QAbstract3DInputHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_ActiveInputHandlerChanged(ptr.Pointer(), PointerFromQAbstract3DInputHandler(inputHandler))
	}
}

//export callbackQAbstract3DGraph_ActiveThemeChanged
func callbackQAbstract3DGraph_ActiveThemeChanged(ptr unsafe.Pointer, theme unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "activeThemeChanged"); signal != nil {
		signal.(func(*Q3DTheme))(NewQ3DThemeFromPointer(theme))
	}

}

func (ptr *QAbstract3DGraph) ConnectActiveThemeChanged(f func(theme *Q3DTheme)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "activeThemeChanged") {
			C.QAbstract3DGraph_ConnectActiveThemeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "activeThemeChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "activeThemeChanged", func(theme *Q3DTheme) {
				signal.(func(*Q3DTheme))(theme)
				f(theme)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "activeThemeChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectActiveThemeChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectActiveThemeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "activeThemeChanged")
	}
}

func (ptr *QAbstract3DGraph) ActiveThemeChanged(theme Q3DTheme_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_ActiveThemeChanged(ptr.Pointer(), PointerFromQ3DTheme(theme))
	}
}

func (ptr *QAbstract3DGraph) AddInputHandler(inputHandler QAbstract3DInputHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_AddInputHandler(ptr.Pointer(), PointerFromQAbstract3DInputHandler(inputHandler))
	}
}

func (ptr *QAbstract3DGraph) AddTheme(theme Q3DTheme_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_AddTheme(ptr.Pointer(), PointerFromQ3DTheme(theme))
	}
}

//export callbackQAbstract3DGraph_AspectRatioChanged
func callbackQAbstract3DGraph_AspectRatioChanged(ptr unsafe.Pointer, ratio C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "aspectRatioChanged"); signal != nil {
		signal.(func(float64))(float64(ratio))
	}

}

func (ptr *QAbstract3DGraph) ConnectAspectRatioChanged(f func(ratio float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "aspectRatioChanged") {
			C.QAbstract3DGraph_ConnectAspectRatioChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "aspectRatioChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "aspectRatioChanged", func(ratio float64) {
				signal.(func(float64))(ratio)
				f(ratio)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "aspectRatioChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectAspectRatioChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectAspectRatioChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "aspectRatioChanged")
	}
}

func (ptr *QAbstract3DGraph) AspectRatioChanged(ratio float64) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_AspectRatioChanged(ptr.Pointer(), C.double(ratio))
	}
}

func (ptr *QAbstract3DGraph) ClearSelection() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_ClearSelection(ptr.Pointer())
	}
}

//export callbackQAbstract3DGraph_CurrentFpsChanged
func callbackQAbstract3DGraph_CurrentFpsChanged(ptr unsafe.Pointer, fps C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "currentFpsChanged"); signal != nil {
		signal.(func(float64))(float64(fps))
	}

}

func (ptr *QAbstract3DGraph) ConnectCurrentFpsChanged(f func(fps float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "currentFpsChanged") {
			C.QAbstract3DGraph_ConnectCurrentFpsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "currentFpsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "currentFpsChanged", func(fps float64) {
				signal.(func(float64))(fps)
				f(fps)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "currentFpsChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectCurrentFpsChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectCurrentFpsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "currentFpsChanged")
	}
}

func (ptr *QAbstract3DGraph) CurrentFpsChanged(fps float64) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_CurrentFpsChanged(ptr.Pointer(), C.double(fps))
	}
}

//export callbackQAbstract3DGraph_HorizontalAspectRatioChanged
func callbackQAbstract3DGraph_HorizontalAspectRatioChanged(ptr unsafe.Pointer, ratio C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "horizontalAspectRatioChanged"); signal != nil {
		signal.(func(float64))(float64(ratio))
	}

}

func (ptr *QAbstract3DGraph) ConnectHorizontalAspectRatioChanged(f func(ratio float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "horizontalAspectRatioChanged") {
			C.QAbstract3DGraph_ConnectHorizontalAspectRatioChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "horizontalAspectRatioChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "horizontalAspectRatioChanged", func(ratio float64) {
				signal.(func(float64))(ratio)
				f(ratio)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "horizontalAspectRatioChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectHorizontalAspectRatioChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectHorizontalAspectRatioChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "horizontalAspectRatioChanged")
	}
}

func (ptr *QAbstract3DGraph) HorizontalAspectRatioChanged(ratio float64) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_HorizontalAspectRatioChanged(ptr.Pointer(), C.double(ratio))
	}
}

//export callbackQAbstract3DGraph_LocaleChanged
func callbackQAbstract3DGraph_LocaleChanged(ptr unsafe.Pointer, locale unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "localeChanged"); signal != nil {
		signal.(func(*core.QLocale))(core.NewQLocaleFromPointer(locale))
	}

}

func (ptr *QAbstract3DGraph) ConnectLocaleChanged(f func(locale *core.QLocale)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "localeChanged") {
			C.QAbstract3DGraph_ConnectLocaleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "localeChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "localeChanged", func(locale *core.QLocale) {
				signal.(func(*core.QLocale))(locale)
				f(locale)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "localeChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectLocaleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectLocaleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "localeChanged")
	}
}

func (ptr *QAbstract3DGraph) LocaleChanged(locale core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_LocaleChanged(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

//export callbackQAbstract3DGraph_MarginChanged
func callbackQAbstract3DGraph_MarginChanged(ptr unsafe.Pointer, margin C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "marginChanged"); signal != nil {
		signal.(func(float64))(float64(margin))
	}

}

func (ptr *QAbstract3DGraph) ConnectMarginChanged(f func(margin float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "marginChanged") {
			C.QAbstract3DGraph_ConnectMarginChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "marginChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "marginChanged", func(margin float64) {
				signal.(func(float64))(margin)
				f(margin)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "marginChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectMarginChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectMarginChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "marginChanged")
	}
}

func (ptr *QAbstract3DGraph) MarginChanged(margin float64) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_MarginChanged(ptr.Pointer(), C.double(margin))
	}
}

//export callbackQAbstract3DGraph_MeasureFpsChanged
func callbackQAbstract3DGraph_MeasureFpsChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "measureFpsChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QAbstract3DGraph) ConnectMeasureFpsChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "measureFpsChanged") {
			C.QAbstract3DGraph_ConnectMeasureFpsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "measureFpsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "measureFpsChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "measureFpsChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectMeasureFpsChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectMeasureFpsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "measureFpsChanged")
	}
}

func (ptr *QAbstract3DGraph) MeasureFpsChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_MeasureFpsChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQAbstract3DGraph_OptimizationHintsChanged
func callbackQAbstract3DGraph_OptimizationHintsChanged(ptr unsafe.Pointer, hints C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "optimizationHintsChanged"); signal != nil {
		signal.(func(QAbstract3DGraph__OptimizationHint))(QAbstract3DGraph__OptimizationHint(hints))
	}

}

func (ptr *QAbstract3DGraph) ConnectOptimizationHintsChanged(f func(hints QAbstract3DGraph__OptimizationHint)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "optimizationHintsChanged") {
			C.QAbstract3DGraph_ConnectOptimizationHintsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "optimizationHintsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "optimizationHintsChanged", func(hints QAbstract3DGraph__OptimizationHint) {
				signal.(func(QAbstract3DGraph__OptimizationHint))(hints)
				f(hints)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "optimizationHintsChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectOptimizationHintsChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectOptimizationHintsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "optimizationHintsChanged")
	}
}

func (ptr *QAbstract3DGraph) OptimizationHintsChanged(hints QAbstract3DGraph__OptimizationHint) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_OptimizationHintsChanged(ptr.Pointer(), C.longlong(hints))
	}
}

//export callbackQAbstract3DGraph_OrthoProjectionChanged
func callbackQAbstract3DGraph_OrthoProjectionChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "orthoProjectionChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QAbstract3DGraph) ConnectOrthoProjectionChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "orthoProjectionChanged") {
			C.QAbstract3DGraph_ConnectOrthoProjectionChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "orthoProjectionChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "orthoProjectionChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "orthoProjectionChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectOrthoProjectionChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectOrthoProjectionChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "orthoProjectionChanged")
	}
}

func (ptr *QAbstract3DGraph) OrthoProjectionChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_OrthoProjectionChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQAbstract3DGraph_PolarChanged
func callbackQAbstract3DGraph_PolarChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "polarChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QAbstract3DGraph) ConnectPolarChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "polarChanged") {
			C.QAbstract3DGraph_ConnectPolarChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "polarChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "polarChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "polarChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectPolarChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectPolarChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "polarChanged")
	}
}

func (ptr *QAbstract3DGraph) PolarChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_PolarChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQAbstract3DGraph_QueriedGraphPositionChanged
func callbackQAbstract3DGraph_QueriedGraphPositionChanged(ptr unsafe.Pointer, data unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "queriedGraphPositionChanged"); signal != nil {
		signal.(func(*gui.QVector3D))(gui.NewQVector3DFromPointer(data))
	}

}

func (ptr *QAbstract3DGraph) ConnectQueriedGraphPositionChanged(f func(data *gui.QVector3D)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "queriedGraphPositionChanged") {
			C.QAbstract3DGraph_ConnectQueriedGraphPositionChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "queriedGraphPositionChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "queriedGraphPositionChanged", func(data *gui.QVector3D) {
				signal.(func(*gui.QVector3D))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "queriedGraphPositionChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectQueriedGraphPositionChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectQueriedGraphPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "queriedGraphPositionChanged")
	}
}

func (ptr *QAbstract3DGraph) QueriedGraphPositionChanged(data gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_QueriedGraphPositionChanged(ptr.Pointer(), gui.PointerFromQVector3D(data))
	}
}

//export callbackQAbstract3DGraph_RadialLabelOffsetChanged
func callbackQAbstract3DGraph_RadialLabelOffsetChanged(ptr unsafe.Pointer, offset C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "radialLabelOffsetChanged"); signal != nil {
		signal.(func(float32))(float32(offset))
	}

}

func (ptr *QAbstract3DGraph) ConnectRadialLabelOffsetChanged(f func(offset float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "radialLabelOffsetChanged") {
			C.QAbstract3DGraph_ConnectRadialLabelOffsetChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "radialLabelOffsetChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "radialLabelOffsetChanged", func(offset float32) {
				signal.(func(float32))(offset)
				f(offset)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "radialLabelOffsetChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectRadialLabelOffsetChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectRadialLabelOffsetChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "radialLabelOffsetChanged")
	}
}

func (ptr *QAbstract3DGraph) RadialLabelOffsetChanged(offset float32) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_RadialLabelOffsetChanged(ptr.Pointer(), C.float(offset))
	}
}

//export callbackQAbstract3DGraph_ReflectionChanged
func callbackQAbstract3DGraph_ReflectionChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "reflectionChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QAbstract3DGraph) ConnectReflectionChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "reflectionChanged") {
			C.QAbstract3DGraph_ConnectReflectionChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "reflectionChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "reflectionChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "reflectionChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectReflectionChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectReflectionChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "reflectionChanged")
	}
}

func (ptr *QAbstract3DGraph) ReflectionChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_ReflectionChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQAbstract3DGraph_ReflectivityChanged
func callbackQAbstract3DGraph_ReflectivityChanged(ptr unsafe.Pointer, reflectivity C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "reflectivityChanged"); signal != nil {
		signal.(func(float64))(float64(reflectivity))
	}

}

func (ptr *QAbstract3DGraph) ConnectReflectivityChanged(f func(reflectivity float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "reflectivityChanged") {
			C.QAbstract3DGraph_ConnectReflectivityChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "reflectivityChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "reflectivityChanged", func(reflectivity float64) {
				signal.(func(float64))(reflectivity)
				f(reflectivity)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "reflectivityChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectReflectivityChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectReflectivityChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "reflectivityChanged")
	}
}

func (ptr *QAbstract3DGraph) ReflectivityChanged(reflectivity float64) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_ReflectivityChanged(ptr.Pointer(), C.double(reflectivity))
	}
}

func (ptr *QAbstract3DGraph) ReleaseCustomItem(item QCustom3DItem_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_ReleaseCustomItem(ptr.Pointer(), PointerFromQCustom3DItem(item))
	}
}

func (ptr *QAbstract3DGraph) ReleaseInputHandler(inputHandler QAbstract3DInputHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_ReleaseInputHandler(ptr.Pointer(), PointerFromQAbstract3DInputHandler(inputHandler))
	}
}

func (ptr *QAbstract3DGraph) ReleaseTheme(theme Q3DTheme_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_ReleaseTheme(ptr.Pointer(), PointerFromQ3DTheme(theme))
	}
}

func (ptr *QAbstract3DGraph) RemoveCustomItem(item QCustom3DItem_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_RemoveCustomItem(ptr.Pointer(), PointerFromQCustom3DItem(item))
	}
}

func (ptr *QAbstract3DGraph) RemoveCustomItemAt(position gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_RemoveCustomItemAt(ptr.Pointer(), gui.PointerFromQVector3D(position))
	}
}

func (ptr *QAbstract3DGraph) RemoveCustomItems() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_RemoveCustomItems(ptr.Pointer())
	}
}

//export callbackQAbstract3DGraph_SelectedElementChanged
func callbackQAbstract3DGraph_SelectedElementChanged(ptr unsafe.Pointer, ty C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "selectedElementChanged"); signal != nil {
		signal.(func(QAbstract3DGraph__ElementType))(QAbstract3DGraph__ElementType(ty))
	}

}

func (ptr *QAbstract3DGraph) ConnectSelectedElementChanged(f func(ty QAbstract3DGraph__ElementType)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "selectedElementChanged") {
			C.QAbstract3DGraph_ConnectSelectedElementChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "selectedElementChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedElementChanged", func(ty QAbstract3DGraph__ElementType) {
				signal.(func(QAbstract3DGraph__ElementType))(ty)
				f(ty)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedElementChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectSelectedElementChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectSelectedElementChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "selectedElementChanged")
	}
}

func (ptr *QAbstract3DGraph) SelectedElementChanged(ty QAbstract3DGraph__ElementType) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SelectedElementChanged(ptr.Pointer(), C.longlong(ty))
	}
}

//export callbackQAbstract3DGraph_SelectionModeChanged
func callbackQAbstract3DGraph_SelectionModeChanged(ptr unsafe.Pointer, mode C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "selectionModeChanged"); signal != nil {
		signal.(func(QAbstract3DGraph__SelectionFlag))(QAbstract3DGraph__SelectionFlag(mode))
	}

}

func (ptr *QAbstract3DGraph) ConnectSelectionModeChanged(f func(mode QAbstract3DGraph__SelectionFlag)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "selectionModeChanged") {
			C.QAbstract3DGraph_ConnectSelectionModeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "selectionModeChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectionModeChanged", func(mode QAbstract3DGraph__SelectionFlag) {
				signal.(func(QAbstract3DGraph__SelectionFlag))(mode)
				f(mode)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectionModeChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectSelectionModeChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectSelectionModeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "selectionModeChanged")
	}
}

func (ptr *QAbstract3DGraph) SelectionModeChanged(mode QAbstract3DGraph__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SelectionModeChanged(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QAbstract3DGraph) SetActiveInputHandler(inputHandler QAbstract3DInputHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetActiveInputHandler(ptr.Pointer(), PointerFromQAbstract3DInputHandler(inputHandler))
	}
}

func (ptr *QAbstract3DGraph) SetActiveTheme(theme Q3DTheme_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetActiveTheme(ptr.Pointer(), PointerFromQ3DTheme(theme))
	}
}

func (ptr *QAbstract3DGraph) SetAspectRatio(ratio float64) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetAspectRatio(ptr.Pointer(), C.double(ratio))
	}
}

func (ptr *QAbstract3DGraph) SetHorizontalAspectRatio(ratio float64) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetHorizontalAspectRatio(ptr.Pointer(), C.double(ratio))
	}
}

func (ptr *QAbstract3DGraph) SetLocale(locale core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

func (ptr *QAbstract3DGraph) SetMargin(margin float64) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QAbstract3DGraph) SetMeasureFps(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetMeasureFps(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QAbstract3DGraph) SetOptimizationHints(hints QAbstract3DGraph__OptimizationHint) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetOptimizationHints(ptr.Pointer(), C.longlong(hints))
	}
}

func (ptr *QAbstract3DGraph) SetOrthoProjection(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetOrthoProjection(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QAbstract3DGraph) SetPolar(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetPolar(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QAbstract3DGraph) SetRadialLabelOffset(offset float32) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetRadialLabelOffset(ptr.Pointer(), C.float(offset))
	}
}

func (ptr *QAbstract3DGraph) SetReflection(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetReflection(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QAbstract3DGraph) SetReflectivity(reflectivity float64) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetReflectivity(ptr.Pointer(), C.double(reflectivity))
	}
}

func (ptr *QAbstract3DGraph) SetSelectionMode(mode QAbstract3DGraph__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetSelectionMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QAbstract3DGraph) SetShadowQuality(quality QAbstract3DGraph__ShadowQuality) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_SetShadowQuality(ptr.Pointer(), C.longlong(quality))
	}
}

//export callbackQAbstract3DGraph_ShadowQualityChanged
func callbackQAbstract3DGraph_ShadowQualityChanged(ptr unsafe.Pointer, quality C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "shadowQualityChanged"); signal != nil {
		signal.(func(QAbstract3DGraph__ShadowQuality))(QAbstract3DGraph__ShadowQuality(quality))
	}

}

func (ptr *QAbstract3DGraph) ConnectShadowQualityChanged(f func(quality QAbstract3DGraph__ShadowQuality)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "shadowQualityChanged") {
			C.QAbstract3DGraph_ConnectShadowQualityChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "shadowQualityChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "shadowQualityChanged", func(quality QAbstract3DGraph__ShadowQuality) {
				signal.(func(QAbstract3DGraph__ShadowQuality))(quality)
				f(quality)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "shadowQualityChanged", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectShadowQualityChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DisconnectShadowQualityChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "shadowQualityChanged")
	}
}

func (ptr *QAbstract3DGraph) ShadowQualityChanged(quality QAbstract3DGraph__ShadowQuality) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_ShadowQualityChanged(ptr.Pointer(), C.longlong(quality))
	}
}

//export callbackQAbstract3DGraph_DestroyQAbstract3DGraph
func callbackQAbstract3DGraph_DestroyQAbstract3DGraph(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QAbstract3DGraph"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstract3DGraphFromPointer(ptr).DestroyQAbstract3DGraphDefault()
	}
}

func (ptr *QAbstract3DGraph) ConnectDestroyQAbstract3DGraph(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DGraph"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DGraph", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DGraph", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectDestroyQAbstract3DGraph() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DGraph")
	}
}

func (ptr *QAbstract3DGraph) DestroyQAbstract3DGraph() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DestroyQAbstract3DGraph(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstract3DGraph) DestroyQAbstract3DGraphDefault() {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph_DestroyQAbstract3DGraphDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstract3DGraph) SelectedElement() QAbstract3DGraph__ElementType {
	if ptr.Pointer() != nil {
		return QAbstract3DGraph__ElementType(C.QAbstract3DGraph_SelectedElement(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DGraph) OptimizationHints() QAbstract3DGraph__OptimizationHint {
	if ptr.Pointer() != nil {
		return QAbstract3DGraph__OptimizationHint(C.QAbstract3DGraph_OptimizationHints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DGraph) Scene() *Q3DScene {
	if ptr.Pointer() != nil {
		return NewQ3DSceneFromPointer(C.QAbstract3DGraph_Scene(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstract3DGraph) ActiveTheme() *Q3DTheme {
	if ptr.Pointer() != nil {
		return NewQ3DThemeFromPointer(C.QAbstract3DGraph_ActiveTheme(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstract3DGraph) SelectedAxis() *QAbstract3DAxis {
	if ptr.Pointer() != nil {
		return NewQAbstract3DAxisFromPointer(C.QAbstract3DGraph_SelectedAxis(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstract3DGraph) ActiveInputHandler() *QAbstract3DInputHandler {
	if ptr.Pointer() != nil {
		return NewQAbstract3DInputHandlerFromPointer(C.QAbstract3DGraph_ActiveInputHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstract3DGraph) SelectedCustomItem() *QCustom3DItem {
	if ptr.Pointer() != nil {
		return NewQCustom3DItemFromPointer(C.QAbstract3DGraph_SelectedCustomItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstract3DGraph) Themes() []*Q3DTheme {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []*Q3DTheme {
			var out = make([]*Q3DTheme, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQAbstract3DGraphFromPointer(l.data).__themes_atList(i)
			}
			return out
		}(C.QAbstract3DGraph_Themes(ptr.Pointer()))
	}
	return make([]*Q3DTheme, 0)
}

func (ptr *QAbstract3DGraph) InputHandlers() []*QAbstract3DInputHandler {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []*QAbstract3DInputHandler {
			var out = make([]*QAbstract3DInputHandler, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQAbstract3DGraphFromPointer(l.data).__inputHandlers_atList(i)
			}
			return out
		}(C.QAbstract3DGraph_InputHandlers(ptr.Pointer()))
	}
	return make([]*QAbstract3DInputHandler, 0)
}

func (ptr *QAbstract3DGraph) CustomItems() []*QCustom3DItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []*QCustom3DItem {
			var out = make([]*QCustom3DItem, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQAbstract3DGraphFromPointer(l.data).__customItems_atList(i)
			}
			return out
		}(C.QAbstract3DGraph_CustomItems(ptr.Pointer()))
	}
	return make([]*QCustom3DItem, 0)
}

func (ptr *QAbstract3DGraph) Locale() *core.QLocale {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQLocaleFromPointer(C.QAbstract3DGraph_Locale(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QLocale).DestroyQLocale)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstract3DGraph) QueriedGraphPosition() *gui.QVector3D {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQVector3DFromPointer(C.QAbstract3DGraph_QueriedGraphPosition(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QVector3D).DestroyQVector3D)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstract3DGraph) SelectionMode() QAbstract3DGraph__SelectionFlag {
	if ptr.Pointer() != nil {
		return QAbstract3DGraph__SelectionFlag(C.QAbstract3DGraph_SelectionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DGraph) ShadowQuality() QAbstract3DGraph__ShadowQuality {
	if ptr.Pointer() != nil {
		return QAbstract3DGraph__ShadowQuality(C.QAbstract3DGraph_ShadowQuality(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DGraph) HasContext() bool {
	if ptr.Pointer() != nil {
		return C.QAbstract3DGraph_HasContext(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstract3DGraph) IsOrthoProjection() bool {
	if ptr.Pointer() != nil {
		return C.QAbstract3DGraph_IsOrthoProjection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstract3DGraph) IsPolar() bool {
	if ptr.Pointer() != nil {
		return C.QAbstract3DGraph_IsPolar(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstract3DGraph) IsReflection() bool {
	if ptr.Pointer() != nil {
		return C.QAbstract3DGraph_IsReflection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstract3DGraph) MeasureFps() bool {
	if ptr.Pointer() != nil {
		return C.QAbstract3DGraph_MeasureFps(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQAbstract3DGraph_ShadowsSupported
func callbackQAbstract3DGraph_ShadowsSupported(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "shadowsSupported"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstract3DGraphFromPointer(ptr).ShadowsSupportedDefault())))
}

func (ptr *QAbstract3DGraph) ConnectShadowsSupported(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "shadowsSupported"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "shadowsSupported", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "shadowsSupported", f)
		}
	}
}

func (ptr *QAbstract3DGraph) DisconnectShadowsSupported() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "shadowsSupported")
	}
}

func (ptr *QAbstract3DGraph) ShadowsSupported() bool {
	if ptr.Pointer() != nil {
		return C.QAbstract3DGraph_ShadowsSupported(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstract3DGraph) ShadowsSupportedDefault() bool {
	if ptr.Pointer() != nil {
		return C.QAbstract3DGraph_ShadowsSupportedDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstract3DGraph) RadialLabelOffset() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QAbstract3DGraph_RadialLabelOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DGraph) SelectedCustomItemIndex() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstract3DGraph_SelectedCustomItemIndex(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstract3DGraph) SelectedLabelIndex() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstract3DGraph_SelectedLabelIndex(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstract3DGraph) AspectRatio() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAbstract3DGraph_AspectRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DGraph) CurrentFps() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAbstract3DGraph_CurrentFps(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DGraph) HorizontalAspectRatio() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAbstract3DGraph_HorizontalAspectRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DGraph) Margin() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAbstract3DGraph_Margin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DGraph) Reflectivity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAbstract3DGraph_Reflectivity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DGraph) __themes_atList(i int) *Q3DTheme {
	if ptr.Pointer() != nil {
		return NewQ3DThemeFromPointer(C.QAbstract3DGraph___themes_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QAbstract3DGraph) __themes_setList(i Q3DTheme_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph___themes_setList(ptr.Pointer(), PointerFromQ3DTheme(i))
	}
}

func (ptr *QAbstract3DGraph) __themes_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QAbstract3DGraph___themes_newList(ptr.Pointer()))
}

func (ptr *QAbstract3DGraph) __inputHandlers_atList(i int) *QAbstract3DInputHandler {
	if ptr.Pointer() != nil {
		return NewQAbstract3DInputHandlerFromPointer(C.QAbstract3DGraph___inputHandlers_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QAbstract3DGraph) __inputHandlers_setList(i QAbstract3DInputHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph___inputHandlers_setList(ptr.Pointer(), PointerFromQAbstract3DInputHandler(i))
	}
}

func (ptr *QAbstract3DGraph) __inputHandlers_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QAbstract3DGraph___inputHandlers_newList(ptr.Pointer()))
}

func (ptr *QAbstract3DGraph) __customItems_atList(i int) *QCustom3DItem {
	if ptr.Pointer() != nil {
		return NewQCustom3DItemFromPointer(C.QAbstract3DGraph___customItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QAbstract3DGraph) __customItems_setList(i QCustom3DItem_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DGraph___customItems_setList(ptr.Pointer(), PointerFromQCustom3DItem(i))
	}
}

func (ptr *QAbstract3DGraph) __customItems_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QAbstract3DGraph___customItems_newList(ptr.Pointer()))
}

type QAbstract3DInputHandler struct {
	ptr unsafe.Pointer
}

type QAbstract3DInputHandler_ITF interface {
	QAbstract3DInputHandler_PTR() *QAbstract3DInputHandler
}

func (ptr *QAbstract3DInputHandler) QAbstract3DInputHandler_PTR() *QAbstract3DInputHandler {
	return ptr
}

func (ptr *QAbstract3DInputHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstract3DInputHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAbstract3DInputHandler(ptr QAbstract3DInputHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstract3DInputHandler_PTR().Pointer()
	}
	return nil
}

func NewQAbstract3DInputHandlerFromPointer(ptr unsafe.Pointer) *QAbstract3DInputHandler {
	var n = new(QAbstract3DInputHandler)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QAbstract3DInputHandler__InputView
//QAbstract3DInputHandler::InputView
type QAbstract3DInputHandler__InputView int64

const (
	QAbstract3DInputHandler__InputViewNone        QAbstract3DInputHandler__InputView = QAbstract3DInputHandler__InputView(0)
	QAbstract3DInputHandler__InputViewOnPrimary   QAbstract3DInputHandler__InputView = QAbstract3DInputHandler__InputView(1)
	QAbstract3DInputHandler__InputViewOnSecondary QAbstract3DInputHandler__InputView = QAbstract3DInputHandler__InputView(2)
)

func (ptr *QAbstract3DInputHandler) SetInputPosition(position core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_SetInputPosition(ptr.Pointer(), core.PointerFromQPoint(position))
	}
}

func (ptr *QAbstract3DInputHandler) SetInputView(inputView QAbstract3DInputHandler__InputView) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_SetInputView(ptr.Pointer(), C.longlong(inputView))
	}
}

func (ptr *QAbstract3DInputHandler) InputView() QAbstract3DInputHandler__InputView {
	if ptr.Pointer() != nil {
		return QAbstract3DInputHandler__InputView(C.QAbstract3DInputHandler_InputView(ptr.Pointer()))
	}
	return 0
}

func NewQAbstract3DInputHandler(parent core.QObject_ITF) *QAbstract3DInputHandler {
	return NewQAbstract3DInputHandlerFromPointer(C.QAbstract3DInputHandler_NewQAbstract3DInputHandler(core.PointerFromQObject(parent)))
}

//export callbackQAbstract3DInputHandler_InputViewChanged
func callbackQAbstract3DInputHandler_InputViewChanged(ptr unsafe.Pointer, view C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "inputViewChanged"); signal != nil {
		signal.(func(QAbstract3DInputHandler__InputView))(QAbstract3DInputHandler__InputView(view))
	}

}

func (ptr *QAbstract3DInputHandler) ConnectInputViewChanged(f func(view QAbstract3DInputHandler__InputView)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "inputViewChanged") {
			C.QAbstract3DInputHandler_ConnectInputViewChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "inputViewChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "inputViewChanged", func(view QAbstract3DInputHandler__InputView) {
				signal.(func(QAbstract3DInputHandler__InputView))(view)
				f(view)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "inputViewChanged", f)
		}
	}
}

func (ptr *QAbstract3DInputHandler) DisconnectInputViewChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_DisconnectInputViewChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "inputViewChanged")
	}
}

func (ptr *QAbstract3DInputHandler) InputViewChanged(view QAbstract3DInputHandler__InputView) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_InputViewChanged(ptr.Pointer(), C.longlong(view))
	}
}

//export callbackQAbstract3DInputHandler_MouseDoubleClickEvent
func callbackQAbstract3DInputHandler_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstract3DInputHandlerFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstract3DInputHandler) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "mouseDoubleClickEvent"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mouseDoubleClickEvent", func(event *gui.QMouseEvent) {
				signal.(func(*gui.QMouseEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mouseDoubleClickEvent", f)
		}
	}
}

func (ptr *QAbstract3DInputHandler) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "mouseDoubleClickEvent")
	}
}

func (ptr *QAbstract3DInputHandler) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstract3DInputHandler) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQAbstract3DInputHandler_MouseMoveEvent
func callbackQAbstract3DInputHandler_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer, mousePos unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent, *core.QPoint))(gui.NewQMouseEventFromPointer(event), core.NewQPointFromPointer(mousePos))
	} else {
		NewQAbstract3DInputHandlerFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event), core.NewQPointFromPointer(mousePos))
	}
}

func (ptr *QAbstract3DInputHandler) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent, mousePos *core.QPoint)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "mouseMoveEvent"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mouseMoveEvent", func(event *gui.QMouseEvent, mousePos *core.QPoint) {
				signal.(func(*gui.QMouseEvent, *core.QPoint))(event, mousePos)
				f(event, mousePos)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mouseMoveEvent", f)
		}
	}
}

func (ptr *QAbstract3DInputHandler) DisconnectMouseMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "mouseMoveEvent")
	}
}

func (ptr *QAbstract3DInputHandler) MouseMoveEvent(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event), core.PointerFromQPoint(mousePos))
	}
}

func (ptr *QAbstract3DInputHandler) MouseMoveEventDefault(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event), core.PointerFromQPoint(mousePos))
	}
}

//export callbackQAbstract3DInputHandler_MousePressEvent
func callbackQAbstract3DInputHandler_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer, mousePos unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent, *core.QPoint))(gui.NewQMouseEventFromPointer(event), core.NewQPointFromPointer(mousePos))
	} else {
		NewQAbstract3DInputHandlerFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event), core.NewQPointFromPointer(mousePos))
	}
}

func (ptr *QAbstract3DInputHandler) ConnectMousePressEvent(f func(event *gui.QMouseEvent, mousePos *core.QPoint)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "mousePressEvent"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mousePressEvent", func(event *gui.QMouseEvent, mousePos *core.QPoint) {
				signal.(func(*gui.QMouseEvent, *core.QPoint))(event, mousePos)
				f(event, mousePos)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mousePressEvent", f)
		}
	}
}

func (ptr *QAbstract3DInputHandler) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "mousePressEvent")
	}
}

func (ptr *QAbstract3DInputHandler) MousePressEvent(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event), core.PointerFromQPoint(mousePos))
	}
}

func (ptr *QAbstract3DInputHandler) MousePressEventDefault(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event), core.PointerFromQPoint(mousePos))
	}
}

//export callbackQAbstract3DInputHandler_MouseReleaseEvent
func callbackQAbstract3DInputHandler_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer, mousePos unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent, *core.QPoint))(gui.NewQMouseEventFromPointer(event), core.NewQPointFromPointer(mousePos))
	} else {
		NewQAbstract3DInputHandlerFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event), core.NewQPointFromPointer(mousePos))
	}
}

func (ptr *QAbstract3DInputHandler) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent, mousePos *core.QPoint)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "mouseReleaseEvent"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mouseReleaseEvent", func(event *gui.QMouseEvent, mousePos *core.QPoint) {
				signal.(func(*gui.QMouseEvent, *core.QPoint))(event, mousePos)
				f(event, mousePos)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "mouseReleaseEvent", f)
		}
	}
}

func (ptr *QAbstract3DInputHandler) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "mouseReleaseEvent")
	}
}

func (ptr *QAbstract3DInputHandler) MouseReleaseEvent(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event), core.PointerFromQPoint(mousePos))
	}
}

func (ptr *QAbstract3DInputHandler) MouseReleaseEventDefault(event gui.QMouseEvent_ITF, mousePos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event), core.PointerFromQPoint(mousePos))
	}
}

//export callbackQAbstract3DInputHandler_PositionChanged
func callbackQAbstract3DInputHandler_PositionChanged(ptr unsafe.Pointer, position unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "positionChanged"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(position))
	}

}

func (ptr *QAbstract3DInputHandler) ConnectPositionChanged(f func(position *core.QPoint)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "positionChanged") {
			C.QAbstract3DInputHandler_ConnectPositionChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "positionChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "positionChanged", func(position *core.QPoint) {
				signal.(func(*core.QPoint))(position)
				f(position)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "positionChanged", f)
		}
	}
}

func (ptr *QAbstract3DInputHandler) DisconnectPositionChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_DisconnectPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "positionChanged")
	}
}

func (ptr *QAbstract3DInputHandler) PositionChanged(position core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_PositionChanged(ptr.Pointer(), core.PointerFromQPoint(position))
	}
}

//export callbackQAbstract3DInputHandler_SceneChanged
func callbackQAbstract3DInputHandler_SceneChanged(ptr unsafe.Pointer, scene unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "sceneChanged"); signal != nil {
		signal.(func(*Q3DScene))(NewQ3DSceneFromPointer(scene))
	}

}

func (ptr *QAbstract3DInputHandler) ConnectSceneChanged(f func(scene *Q3DScene)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "sceneChanged") {
			C.QAbstract3DInputHandler_ConnectSceneChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "sceneChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sceneChanged", func(scene *Q3DScene) {
				signal.(func(*Q3DScene))(scene)
				f(scene)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sceneChanged", f)
		}
	}
}

func (ptr *QAbstract3DInputHandler) DisconnectSceneChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_DisconnectSceneChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "sceneChanged")
	}
}

func (ptr *QAbstract3DInputHandler) SceneChanged(scene Q3DScene_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_SceneChanged(ptr.Pointer(), PointerFromQ3DScene(scene))
	}
}

func (ptr *QAbstract3DInputHandler) SetPrevDistance(distance int) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_SetPrevDistance(ptr.Pointer(), C.int(int32(distance)))
	}
}

func (ptr *QAbstract3DInputHandler) SetPreviousInputPos(position core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_SetPreviousInputPos(ptr.Pointer(), core.PointerFromQPoint(position))
	}
}

func (ptr *QAbstract3DInputHandler) SetScene(scene Q3DScene_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_SetScene(ptr.Pointer(), PointerFromQ3DScene(scene))
	}
}

//export callbackQAbstract3DInputHandler_TouchEvent
func callbackQAbstract3DInputHandler_TouchEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
	} else {
		NewQAbstract3DInputHandlerFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(event))
	}
}

func (ptr *QAbstract3DInputHandler) ConnectTouchEvent(f func(event *gui.QTouchEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "touchEvent"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "touchEvent", func(event *gui.QTouchEvent) {
				signal.(func(*gui.QTouchEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "touchEvent", f)
		}
	}
}

func (ptr *QAbstract3DInputHandler) DisconnectTouchEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "touchEvent")
	}
}

func (ptr *QAbstract3DInputHandler) TouchEvent(event gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_TouchEvent(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

func (ptr *QAbstract3DInputHandler) TouchEventDefault(event gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

//export callbackQAbstract3DInputHandler_WheelEvent
func callbackQAbstract3DInputHandler_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQAbstract3DInputHandlerFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QAbstract3DInputHandler) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "wheelEvent"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "wheelEvent", func(event *gui.QWheelEvent) {
				signal.(func(*gui.QWheelEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "wheelEvent", f)
		}
	}
}

func (ptr *QAbstract3DInputHandler) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "wheelEvent")
	}
}

func (ptr *QAbstract3DInputHandler) WheelEvent(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QAbstract3DInputHandler) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQAbstract3DInputHandler_DestroyQAbstract3DInputHandler
func callbackQAbstract3DInputHandler_DestroyQAbstract3DInputHandler(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QAbstract3DInputHandler"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstract3DInputHandlerFromPointer(ptr).DestroyQAbstract3DInputHandlerDefault()
	}
}

func (ptr *QAbstract3DInputHandler) ConnectDestroyQAbstract3DInputHandler(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DInputHandler"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DInputHandler", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DInputHandler", f)
		}
	}
}

func (ptr *QAbstract3DInputHandler) DisconnectDestroyQAbstract3DInputHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DInputHandler")
	}
}

func (ptr *QAbstract3DInputHandler) DestroyQAbstract3DInputHandler() {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_DestroyQAbstract3DInputHandler(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstract3DInputHandler) DestroyQAbstract3DInputHandlerDefault() {
	if ptr.Pointer() != nil {
		C.QAbstract3DInputHandler_DestroyQAbstract3DInputHandlerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstract3DInputHandler) Scene() *Q3DScene {
	if ptr.Pointer() != nil {
		return NewQ3DSceneFromPointer(C.QAbstract3DInputHandler_Scene(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstract3DInputHandler) InputPosition() *core.QPoint {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFromPointer(C.QAbstract3DInputHandler_InputPosition(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstract3DInputHandler) PreviousInputPos() *core.QPoint {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFromPointer(C.QAbstract3DInputHandler_PreviousInputPos(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstract3DInputHandler) PrevDistance() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstract3DInputHandler_PrevDistance(ptr.Pointer())))
	}
	return 0
}

type QAbstract3DSeries struct {
	ptr unsafe.Pointer
}

type QAbstract3DSeries_ITF interface {
	QAbstract3DSeries_PTR() *QAbstract3DSeries
}

func (ptr *QAbstract3DSeries) QAbstract3DSeries_PTR() *QAbstract3DSeries {
	return ptr
}

func (ptr *QAbstract3DSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstract3DSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAbstract3DSeries(ptr QAbstract3DSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstract3DSeries_PTR().Pointer()
	}
	return nil
}

func NewQAbstract3DSeriesFromPointer(ptr unsafe.Pointer) *QAbstract3DSeries {
	var n = new(QAbstract3DSeries)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QAbstract3DSeries__Mesh
//QAbstract3DSeries::Mesh
type QAbstract3DSeries__Mesh int64

const (
	QAbstract3DSeries__MeshUserDefined QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(0)
	QAbstract3DSeries__MeshBar         QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(1)
	QAbstract3DSeries__MeshCube        QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(2)
	QAbstract3DSeries__MeshPyramid     QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(3)
	QAbstract3DSeries__MeshCone        QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(4)
	QAbstract3DSeries__MeshCylinder    QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(5)
	QAbstract3DSeries__MeshBevelBar    QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(6)
	QAbstract3DSeries__MeshBevelCube   QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(7)
	QAbstract3DSeries__MeshSphere      QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(8)
	QAbstract3DSeries__MeshMinimal     QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(9)
	QAbstract3DSeries__MeshArrow       QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(10)
	QAbstract3DSeries__MeshPoint       QAbstract3DSeries__Mesh = QAbstract3DSeries__Mesh(11)
)

//go:generate stringer -type=QAbstract3DSeries__SeriesType
//QAbstract3DSeries::SeriesType
type QAbstract3DSeries__SeriesType int64

const (
	QAbstract3DSeries__SeriesTypeNone    QAbstract3DSeries__SeriesType = QAbstract3DSeries__SeriesType(0)
	QAbstract3DSeries__SeriesTypeBar     QAbstract3DSeries__SeriesType = QAbstract3DSeries__SeriesType(1)
	QAbstract3DSeries__SeriesTypeScatter QAbstract3DSeries__SeriesType = QAbstract3DSeries__SeriesType(2)
	QAbstract3DSeries__SeriesTypeSurface QAbstract3DSeries__SeriesType = QAbstract3DSeries__SeriesType(4)
)

//export callbackQAbstract3DSeries_MultiHighlightGradientChanged
func callbackQAbstract3DSeries_MultiHighlightGradientChanged(ptr unsafe.Pointer, gradient unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "multiHighlightGradientChanged"); signal != nil {
		signal.(func(*gui.QLinearGradient))(gui.NewQLinearGradientFromPointer(gradient))
	}

}

func (ptr *QAbstract3DSeries) ConnectMultiHighlightGradientChanged(f func(gradient *gui.QLinearGradient)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightGradientChanged") {
			C.QAbstract3DSeries_ConnectMultiHighlightGradientChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightGradientChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightGradientChanged", func(gradient *gui.QLinearGradient) {
				signal.(func(*gui.QLinearGradient))(gradient)
				f(gradient)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightGradientChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectMultiHighlightGradientChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectMultiHighlightGradientChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightGradientChanged")
	}
}

func (ptr *QAbstract3DSeries) MultiHighlightGradientChanged(gradient gui.QLinearGradient_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_MultiHighlightGradientChanged(ptr.Pointer(), gui.PointerFromQLinearGradient(gradient))
	}
}

func (ptr *QAbstract3DSeries) SetMeshRotation(rotation gui.QQuaternion_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SetMeshRotation(ptr.Pointer(), gui.PointerFromQQuaternion(rotation))
	}
}

//export callbackQAbstract3DSeries_BaseColorChanged
func callbackQAbstract3DSeries_BaseColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "baseColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAbstract3DSeries) ConnectBaseColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "baseColorChanged") {
			C.QAbstract3DSeries_ConnectBaseColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "baseColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "baseColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "baseColorChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectBaseColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectBaseColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "baseColorChanged")
	}
}

func (ptr *QAbstract3DSeries) BaseColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_BaseColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQAbstract3DSeries_BaseGradientChanged
func callbackQAbstract3DSeries_BaseGradientChanged(ptr unsafe.Pointer, gradient unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "baseGradientChanged"); signal != nil {
		signal.(func(*gui.QLinearGradient))(gui.NewQLinearGradientFromPointer(gradient))
	}

}

func (ptr *QAbstract3DSeries) ConnectBaseGradientChanged(f func(gradient *gui.QLinearGradient)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "baseGradientChanged") {
			C.QAbstract3DSeries_ConnectBaseGradientChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "baseGradientChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "baseGradientChanged", func(gradient *gui.QLinearGradient) {
				signal.(func(*gui.QLinearGradient))(gradient)
				f(gradient)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "baseGradientChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectBaseGradientChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectBaseGradientChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "baseGradientChanged")
	}
}

func (ptr *QAbstract3DSeries) BaseGradientChanged(gradient gui.QLinearGradient_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_BaseGradientChanged(ptr.Pointer(), gui.PointerFromQLinearGradient(gradient))
	}
}

//export callbackQAbstract3DSeries_ColorStyleChanged
func callbackQAbstract3DSeries_ColorStyleChanged(ptr unsafe.Pointer, style C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "colorStyleChanged"); signal != nil {
		signal.(func(Q3DTheme__ColorStyle))(Q3DTheme__ColorStyle(style))
	}

}

func (ptr *QAbstract3DSeries) ConnectColorStyleChanged(f func(style Q3DTheme__ColorStyle)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "colorStyleChanged") {
			C.QAbstract3DSeries_ConnectColorStyleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "colorStyleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "colorStyleChanged", func(style Q3DTheme__ColorStyle) {
				signal.(func(Q3DTheme__ColorStyle))(style)
				f(style)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "colorStyleChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectColorStyleChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectColorStyleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "colorStyleChanged")
	}
}

func (ptr *QAbstract3DSeries) ColorStyleChanged(style Q3DTheme__ColorStyle) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_ColorStyleChanged(ptr.Pointer(), C.longlong(style))
	}
}

//export callbackQAbstract3DSeries_ItemLabelChanged
func callbackQAbstract3DSeries_ItemLabelChanged(ptr unsafe.Pointer, label C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemLabelChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(label))
	}

}

func (ptr *QAbstract3DSeries) ConnectItemLabelChanged(f func(label string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemLabelChanged") {
			C.QAbstract3DSeries_ConnectItemLabelChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemLabelChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemLabelChanged", func(label string) {
				signal.(func(string))(label)
				f(label)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemLabelChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectItemLabelChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectItemLabelChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemLabelChanged")
	}
}

func (ptr *QAbstract3DSeries) ItemLabelChanged(label string) {
	if ptr.Pointer() != nil {
		var labelC *C.char
		if label != "" {
			labelC = C.CString(label)
			defer C.free(unsafe.Pointer(labelC))
		}
		C.QAbstract3DSeries_ItemLabelChanged(ptr.Pointer(), labelC)
	}
}

//export callbackQAbstract3DSeries_ItemLabelFormatChanged
func callbackQAbstract3DSeries_ItemLabelFormatChanged(ptr unsafe.Pointer, format C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemLabelFormatChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(format))
	}

}

func (ptr *QAbstract3DSeries) ConnectItemLabelFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemLabelFormatChanged") {
			C.QAbstract3DSeries_ConnectItemLabelFormatChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemLabelFormatChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemLabelFormatChanged", func(format string) {
				signal.(func(string))(format)
				f(format)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemLabelFormatChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectItemLabelFormatChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectItemLabelFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemLabelFormatChanged")
	}
}

func (ptr *QAbstract3DSeries) ItemLabelFormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QAbstract3DSeries_ItemLabelFormatChanged(ptr.Pointer(), formatC)
	}
}

//export callbackQAbstract3DSeries_ItemLabelVisibilityChanged
func callbackQAbstract3DSeries_ItemLabelVisibilityChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemLabelVisibilityChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QAbstract3DSeries) ConnectItemLabelVisibilityChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemLabelVisibilityChanged") {
			C.QAbstract3DSeries_ConnectItemLabelVisibilityChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemLabelVisibilityChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemLabelVisibilityChanged", func(visible bool) {
				signal.(func(bool))(visible)
				f(visible)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemLabelVisibilityChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectItemLabelVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectItemLabelVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemLabelVisibilityChanged")
	}
}

func (ptr *QAbstract3DSeries) ItemLabelVisibilityChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_ItemLabelVisibilityChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstract3DSeries_MeshChanged
func callbackQAbstract3DSeries_MeshChanged(ptr unsafe.Pointer, mesh C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "meshChanged"); signal != nil {
		signal.(func(QAbstract3DSeries__Mesh))(QAbstract3DSeries__Mesh(mesh))
	}

}

func (ptr *QAbstract3DSeries) ConnectMeshChanged(f func(mesh QAbstract3DSeries__Mesh)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "meshChanged") {
			C.QAbstract3DSeries_ConnectMeshChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "meshChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "meshChanged", func(mesh QAbstract3DSeries__Mesh) {
				signal.(func(QAbstract3DSeries__Mesh))(mesh)
				f(mesh)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "meshChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectMeshChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectMeshChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "meshChanged")
	}
}

func (ptr *QAbstract3DSeries) MeshChanged(mesh QAbstract3DSeries__Mesh) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_MeshChanged(ptr.Pointer(), C.longlong(mesh))
	}
}

//export callbackQAbstract3DSeries_MeshRotationChanged
func callbackQAbstract3DSeries_MeshRotationChanged(ptr unsafe.Pointer, rotation unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "meshRotationChanged"); signal != nil {
		signal.(func(*gui.QQuaternion))(gui.NewQQuaternionFromPointer(rotation))
	}

}

func (ptr *QAbstract3DSeries) ConnectMeshRotationChanged(f func(rotation *gui.QQuaternion)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "meshRotationChanged") {
			C.QAbstract3DSeries_ConnectMeshRotationChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "meshRotationChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "meshRotationChanged", func(rotation *gui.QQuaternion) {
				signal.(func(*gui.QQuaternion))(rotation)
				f(rotation)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "meshRotationChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectMeshRotationChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectMeshRotationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "meshRotationChanged")
	}
}

func (ptr *QAbstract3DSeries) MeshRotationChanged(rotation gui.QQuaternion_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_MeshRotationChanged(ptr.Pointer(), gui.PointerFromQQuaternion(rotation))
	}
}

//export callbackQAbstract3DSeries_MeshSmoothChanged
func callbackQAbstract3DSeries_MeshSmoothChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "meshSmoothChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QAbstract3DSeries) ConnectMeshSmoothChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "meshSmoothChanged") {
			C.QAbstract3DSeries_ConnectMeshSmoothChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "meshSmoothChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "meshSmoothChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "meshSmoothChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectMeshSmoothChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectMeshSmoothChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "meshSmoothChanged")
	}
}

func (ptr *QAbstract3DSeries) MeshSmoothChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_MeshSmoothChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQAbstract3DSeries_MultiHighlightColorChanged
func callbackQAbstract3DSeries_MultiHighlightColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "multiHighlightColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAbstract3DSeries) ConnectMultiHighlightColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightColorChanged") {
			C.QAbstract3DSeries_ConnectMultiHighlightColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightColorChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectMultiHighlightColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectMultiHighlightColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "multiHighlightColorChanged")
	}
}

func (ptr *QAbstract3DSeries) MultiHighlightColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_MultiHighlightColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQAbstract3DSeries_NameChanged
func callbackQAbstract3DSeries_NameChanged(ptr unsafe.Pointer, name C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "nameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(name))
	}

}

func (ptr *QAbstract3DSeries) ConnectNameChanged(f func(name string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "nameChanged") {
			C.QAbstract3DSeries_ConnectNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "nameChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "nameChanged", func(name string) {
				signal.(func(string))(name)
				f(name)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "nameChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectNameChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "nameChanged")
	}
}

func (ptr *QAbstract3DSeries) NameChanged(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QAbstract3DSeries_NameChanged(ptr.Pointer(), nameC)
	}
}

func (ptr *QAbstract3DSeries) SetBaseColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SetBaseColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstract3DSeries) SetBaseGradient(gradient gui.QLinearGradient_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SetBaseGradient(ptr.Pointer(), gui.PointerFromQLinearGradient(gradient))
	}
}

func (ptr *QAbstract3DSeries) SetColorStyle(style Q3DTheme__ColorStyle) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SetColorStyle(ptr.Pointer(), C.longlong(style))
	}
}

func (ptr *QAbstract3DSeries) SetItemLabelFormat(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QAbstract3DSeries_SetItemLabelFormat(ptr.Pointer(), formatC)
	}
}

func (ptr *QAbstract3DSeries) SetItemLabelVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SetItemLabelVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QAbstract3DSeries) SetMesh(mesh QAbstract3DSeries__Mesh) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SetMesh(ptr.Pointer(), C.longlong(mesh))
	}
}

func (ptr *QAbstract3DSeries) SetMeshAxisAndAngle(axis gui.QVector3D_ITF, angle float32) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SetMeshAxisAndAngle(ptr.Pointer(), gui.PointerFromQVector3D(axis), C.float(angle))
	}
}

func (ptr *QAbstract3DSeries) SetMeshSmooth(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SetMeshSmooth(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QAbstract3DSeries) SetMultiHighlightColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SetMultiHighlightColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstract3DSeries) SetMultiHighlightGradient(gradient gui.QLinearGradient_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SetMultiHighlightGradient(ptr.Pointer(), gui.PointerFromQLinearGradient(gradient))
	}
}

func (ptr *QAbstract3DSeries) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QAbstract3DSeries_SetName(ptr.Pointer(), nameC)
	}
}

func (ptr *QAbstract3DSeries) SetSingleHighlightColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SetSingleHighlightColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QAbstract3DSeries) SetSingleHighlightGradient(gradient gui.QLinearGradient_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SetSingleHighlightGradient(ptr.Pointer(), gui.PointerFromQLinearGradient(gradient))
	}
}

func (ptr *QAbstract3DSeries) SetUserDefinedMesh(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QAbstract3DSeries_SetUserDefinedMesh(ptr.Pointer(), fileNameC)
	}
}

func (ptr *QAbstract3DSeries) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstract3DSeries_SingleHighlightColorChanged
func callbackQAbstract3DSeries_SingleHighlightColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "singleHighlightColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QAbstract3DSeries) ConnectSingleHighlightColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightColorChanged") {
			C.QAbstract3DSeries_ConnectSingleHighlightColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightColorChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectSingleHighlightColorChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectSingleHighlightColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightColorChanged")
	}
}

func (ptr *QAbstract3DSeries) SingleHighlightColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SingleHighlightColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQAbstract3DSeries_SingleHighlightGradientChanged
func callbackQAbstract3DSeries_SingleHighlightGradientChanged(ptr unsafe.Pointer, gradient unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "singleHighlightGradientChanged"); signal != nil {
		signal.(func(*gui.QLinearGradient))(gui.NewQLinearGradientFromPointer(gradient))
	}

}

func (ptr *QAbstract3DSeries) ConnectSingleHighlightGradientChanged(f func(gradient *gui.QLinearGradient)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightGradientChanged") {
			C.QAbstract3DSeries_ConnectSingleHighlightGradientChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightGradientChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightGradientChanged", func(gradient *gui.QLinearGradient) {
				signal.(func(*gui.QLinearGradient))(gradient)
				f(gradient)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightGradientChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectSingleHighlightGradientChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectSingleHighlightGradientChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "singleHighlightGradientChanged")
	}
}

func (ptr *QAbstract3DSeries) SingleHighlightGradientChanged(gradient gui.QLinearGradient_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_SingleHighlightGradientChanged(ptr.Pointer(), gui.PointerFromQLinearGradient(gradient))
	}
}

//export callbackQAbstract3DSeries_UserDefinedMeshChanged
func callbackQAbstract3DSeries_UserDefinedMeshChanged(ptr unsafe.Pointer, fileName C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "userDefinedMeshChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(fileName))
	}

}

func (ptr *QAbstract3DSeries) ConnectUserDefinedMeshChanged(f func(fileName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "userDefinedMeshChanged") {
			C.QAbstract3DSeries_ConnectUserDefinedMeshChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "userDefinedMeshChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "userDefinedMeshChanged", func(fileName string) {
				signal.(func(string))(fileName)
				f(fileName)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "userDefinedMeshChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectUserDefinedMeshChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectUserDefinedMeshChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "userDefinedMeshChanged")
	}
}

func (ptr *QAbstract3DSeries) UserDefinedMeshChanged(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QAbstract3DSeries_UserDefinedMeshChanged(ptr.Pointer(), fileNameC)
	}
}

//export callbackQAbstract3DSeries_VisibilityChanged
func callbackQAbstract3DSeries_VisibilityChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "visibilityChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QAbstract3DSeries) ConnectVisibilityChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "visibilityChanged") {
			C.QAbstract3DSeries_ConnectVisibilityChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "visibilityChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "visibilityChanged", func(visible bool) {
				signal.(func(bool))(visible)
				f(visible)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "visibilityChanged", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DisconnectVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "visibilityChanged")
	}
}

func (ptr *QAbstract3DSeries) VisibilityChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_VisibilityChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQAbstract3DSeries_DestroyQAbstract3DSeries
func callbackQAbstract3DSeries_DestroyQAbstract3DSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QAbstract3DSeries"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstract3DSeriesFromPointer(ptr).DestroyQAbstract3DSeriesDefault()
	}
}

func (ptr *QAbstract3DSeries) ConnectDestroyQAbstract3DSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DSeries"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DSeries", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DSeries", f)
		}
	}
}

func (ptr *QAbstract3DSeries) DisconnectDestroyQAbstract3DSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstract3DSeries")
	}
}

func (ptr *QAbstract3DSeries) DestroyQAbstract3DSeries() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DestroyQAbstract3DSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstract3DSeries) DestroyQAbstract3DSeriesDefault() {
	if ptr.Pointer() != nil {
		C.QAbstract3DSeries_DestroyQAbstract3DSeriesDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstract3DSeries) Mesh() QAbstract3DSeries__Mesh {
	if ptr.Pointer() != nil {
		return QAbstract3DSeries__Mesh(C.QAbstract3DSeries_Mesh(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DSeries) ColorStyle() Q3DTheme__ColorStyle {
	if ptr.Pointer() != nil {
		return Q3DTheme__ColorStyle(C.QAbstract3DSeries_ColorStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DSeries) BaseColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QAbstract3DSeries_BaseColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstract3DSeries) MultiHighlightColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QAbstract3DSeries_MultiHighlightColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstract3DSeries) SingleHighlightColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QAbstract3DSeries_SingleHighlightColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstract3DSeries) BaseGradient() *gui.QLinearGradient {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQLinearGradientFromPointer(C.QAbstract3DSeries_BaseGradient(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QLinearGradient).DestroyQLinearGradient)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstract3DSeries) MultiHighlightGradient() *gui.QLinearGradient {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQLinearGradientFromPointer(C.QAbstract3DSeries_MultiHighlightGradient(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QLinearGradient).DestroyQLinearGradient)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstract3DSeries) SingleHighlightGradient() *gui.QLinearGradient {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQLinearGradientFromPointer(C.QAbstract3DSeries_SingleHighlightGradient(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QLinearGradient).DestroyQLinearGradient)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstract3DSeries) MeshRotation() *gui.QQuaternion {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQQuaternionFromPointer(C.QAbstract3DSeries_MeshRotation(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QQuaternion).DestroyQQuaternion)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstract3DSeries) ItemLabel() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstract3DSeries_ItemLabel(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstract3DSeries) ItemLabelFormat() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstract3DSeries_ItemLabelFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstract3DSeries) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstract3DSeries_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstract3DSeries) UserDefinedMesh() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstract3DSeries_UserDefinedMesh(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstract3DSeries) Type() QAbstract3DSeries__SeriesType {
	if ptr.Pointer() != nil {
		return QAbstract3DSeries__SeriesType(C.QAbstract3DSeries_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstract3DSeries) IsItemLabelVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAbstract3DSeries_IsItemLabelVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstract3DSeries) IsMeshSmooth() bool {
	if ptr.Pointer() != nil {
		return C.QAbstract3DSeries_IsMeshSmooth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstract3DSeries) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAbstract3DSeries_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

type QAbstractDataProxy struct {
	ptr unsafe.Pointer
}

type QAbstractDataProxy_ITF interface {
	QAbstractDataProxy_PTR() *QAbstractDataProxy
}

func (ptr *QAbstractDataProxy) QAbstractDataProxy_PTR() *QAbstractDataProxy {
	return ptr
}

func (ptr *QAbstractDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstractDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAbstractDataProxy(ptr QAbstractDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractDataProxy_PTR().Pointer()
	}
	return nil
}

func NewQAbstractDataProxyFromPointer(ptr unsafe.Pointer) *QAbstractDataProxy {
	var n = new(QAbstractDataProxy)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QAbstractDataProxy__DataType
//QAbstractDataProxy::DataType
type QAbstractDataProxy__DataType int64

const (
	QAbstractDataProxy__DataTypeNone    QAbstractDataProxy__DataType = QAbstractDataProxy__DataType(0)
	QAbstractDataProxy__DataTypeBar     QAbstractDataProxy__DataType = QAbstractDataProxy__DataType(1)
	QAbstractDataProxy__DataTypeScatter QAbstractDataProxy__DataType = QAbstractDataProxy__DataType(2)
	QAbstractDataProxy__DataTypeSurface QAbstractDataProxy__DataType = QAbstractDataProxy__DataType(4)
)

func (ptr *QAbstractDataProxy) Type() QAbstractDataProxy__DataType {
	if ptr.Pointer() != nil {
		return QAbstractDataProxy__DataType(C.QAbstractDataProxy_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractDataProxy_DestroyQAbstractDataProxy
func callbackQAbstractDataProxy_DestroyQAbstractDataProxy(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QAbstractDataProxy"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractDataProxyFromPointer(ptr).DestroyQAbstractDataProxyDefault()
	}
}

func (ptr *QAbstractDataProxy) ConnectDestroyQAbstractDataProxy(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QAbstractDataProxy"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstractDataProxy", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstractDataProxy", f)
		}
	}
}

func (ptr *QAbstractDataProxy) DisconnectDestroyQAbstractDataProxy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QAbstractDataProxy")
	}
}

func (ptr *QAbstractDataProxy) DestroyQAbstractDataProxy() {
	if ptr.Pointer() != nil {
		C.QAbstractDataProxy_DestroyQAbstractDataProxy(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractDataProxy) DestroyQAbstractDataProxyDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractDataProxy_DestroyQAbstractDataProxyDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QBar3DSeries struct {
	ptr unsafe.Pointer
}

type QBar3DSeries_ITF interface {
	QBar3DSeries_PTR() *QBar3DSeries
}

func (ptr *QBar3DSeries) QBar3DSeries_PTR() *QBar3DSeries {
	return ptr
}

func (ptr *QBar3DSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBar3DSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQBar3DSeries(ptr QBar3DSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBar3DSeries_PTR().Pointer()
	}
	return nil
}

func NewQBar3DSeriesFromPointer(ptr unsafe.Pointer) *QBar3DSeries {
	var n = new(QBar3DSeries)
	n.SetPointer(ptr)
	return n
}
func NewQBar3DSeries2(dataProxy QBarDataProxy_ITF, parent core.QObject_ITF) *QBar3DSeries {
	return NewQBar3DSeriesFromPointer(C.QBar3DSeries_NewQBar3DSeries2(PointerFromQBarDataProxy(dataProxy), core.PointerFromQObject(parent)))
}

func NewQBar3DSeries(parent core.QObject_ITF) *QBar3DSeries {
	return NewQBar3DSeriesFromPointer(C.QBar3DSeries_NewQBar3DSeries(core.PointerFromQObject(parent)))
}

func QBar3DSeries_InvalidSelectionPosition() *core.QPoint {
	var tmpValue = core.NewQPointFromPointer(C.QBar3DSeries_QBar3DSeries_InvalidSelectionPosition())
	runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
	return tmpValue
}

func (ptr *QBar3DSeries) InvalidSelectionPosition() *core.QPoint {
	var tmpValue = core.NewQPointFromPointer(C.QBar3DSeries_QBar3DSeries_InvalidSelectionPosition())
	runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
	return tmpValue
}

//export callbackQBar3DSeries_DataProxyChanged
func callbackQBar3DSeries_DataProxyChanged(ptr unsafe.Pointer, proxy unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dataProxyChanged"); signal != nil {
		signal.(func(*QBarDataProxy))(NewQBarDataProxyFromPointer(proxy))
	}

}

func (ptr *QBar3DSeries) ConnectDataProxyChanged(f func(proxy *QBarDataProxy)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged") {
			C.QBar3DSeries_ConnectDataProxyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged", func(proxy *QBarDataProxy) {
				signal.(func(*QBarDataProxy))(proxy)
				f(proxy)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged", f)
		}
	}
}

func (ptr *QBar3DSeries) DisconnectDataProxyChanged() {
	if ptr.Pointer() != nil {
		C.QBar3DSeries_DisconnectDataProxyChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged")
	}
}

func (ptr *QBar3DSeries) DataProxyChanged(proxy QBarDataProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QBar3DSeries_DataProxyChanged(ptr.Pointer(), PointerFromQBarDataProxy(proxy))
	}
}

//export callbackQBar3DSeries_MeshAngleChanged
func callbackQBar3DSeries_MeshAngleChanged(ptr unsafe.Pointer, angle C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "meshAngleChanged"); signal != nil {
		signal.(func(float32))(float32(angle))
	}

}

func (ptr *QBar3DSeries) ConnectMeshAngleChanged(f func(angle float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "meshAngleChanged") {
			C.QBar3DSeries_ConnectMeshAngleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "meshAngleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "meshAngleChanged", func(angle float32) {
				signal.(func(float32))(angle)
				f(angle)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "meshAngleChanged", f)
		}
	}
}

func (ptr *QBar3DSeries) DisconnectMeshAngleChanged() {
	if ptr.Pointer() != nil {
		C.QBar3DSeries_DisconnectMeshAngleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "meshAngleChanged")
	}
}

func (ptr *QBar3DSeries) MeshAngleChanged(angle float32) {
	if ptr.Pointer() != nil {
		C.QBar3DSeries_MeshAngleChanged(ptr.Pointer(), C.float(angle))
	}
}

//export callbackQBar3DSeries_SelectedBarChanged
func callbackQBar3DSeries_SelectedBarChanged(ptr unsafe.Pointer, position unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "selectedBarChanged"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(position))
	}

}

func (ptr *QBar3DSeries) ConnectSelectedBarChanged(f func(position *core.QPoint)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "selectedBarChanged") {
			C.QBar3DSeries_ConnectSelectedBarChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "selectedBarChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedBarChanged", func(position *core.QPoint) {
				signal.(func(*core.QPoint))(position)
				f(position)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedBarChanged", f)
		}
	}
}

func (ptr *QBar3DSeries) DisconnectSelectedBarChanged() {
	if ptr.Pointer() != nil {
		C.QBar3DSeries_DisconnectSelectedBarChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "selectedBarChanged")
	}
}

func (ptr *QBar3DSeries) SelectedBarChanged(position core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QBar3DSeries_SelectedBarChanged(ptr.Pointer(), core.PointerFromQPoint(position))
	}
}

func (ptr *QBar3DSeries) SetDataProxy(proxy QBarDataProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QBar3DSeries_SetDataProxy(ptr.Pointer(), PointerFromQBarDataProxy(proxy))
	}
}

func (ptr *QBar3DSeries) SetMeshAngle(angle float32) {
	if ptr.Pointer() != nil {
		C.QBar3DSeries_SetMeshAngle(ptr.Pointer(), C.float(angle))
	}
}

func (ptr *QBar3DSeries) SetSelectedBar(position core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QBar3DSeries_SetSelectedBar(ptr.Pointer(), core.PointerFromQPoint(position))
	}
}

//export callbackQBar3DSeries_DestroyQBar3DSeries
func callbackQBar3DSeries_DestroyQBar3DSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QBar3DSeries"); signal != nil {
		signal.(func())()
	} else {
		NewQBar3DSeriesFromPointer(ptr).DestroyQBar3DSeriesDefault()
	}
}

func (ptr *QBar3DSeries) ConnectDestroyQBar3DSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QBar3DSeries"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QBar3DSeries", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QBar3DSeries", f)
		}
	}
}

func (ptr *QBar3DSeries) DisconnectDestroyQBar3DSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QBar3DSeries")
	}
}

func (ptr *QBar3DSeries) DestroyQBar3DSeries() {
	if ptr.Pointer() != nil {
		C.QBar3DSeries_DestroyQBar3DSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBar3DSeries) DestroyQBar3DSeriesDefault() {
	if ptr.Pointer() != nil {
		C.QBar3DSeries_DestroyQBar3DSeriesDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBar3DSeries) DataProxy() *QBarDataProxy {
	if ptr.Pointer() != nil {
		return NewQBarDataProxyFromPointer(C.QBar3DSeries_DataProxy(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBar3DSeries) SelectedBar() *core.QPoint {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFromPointer(C.QBar3DSeries_SelectedBar(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QBar3DSeries) MeshAngle() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QBar3DSeries_MeshAngle(ptr.Pointer()))
	}
	return 0
}

type QBarDataItem struct {
	ptr unsafe.Pointer
}

type QBarDataItem_ITF interface {
	QBarDataItem_PTR() *QBarDataItem
}

func (ptr *QBarDataItem) QBarDataItem_PTR() *QBarDataItem {
	return ptr
}

func (ptr *QBarDataItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBarDataItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQBarDataItem(ptr QBarDataItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBarDataItem_PTR().Pointer()
	}
	return nil
}

func NewQBarDataItemFromPointer(ptr unsafe.Pointer) *QBarDataItem {
	var n = new(QBarDataItem)
	n.SetPointer(ptr)
	return n
}
func NewQBarDataItem() *QBarDataItem {
	var tmpValue = NewQBarDataItemFromPointer(C.QBarDataItem_NewQBarDataItem())
	runtime.SetFinalizer(tmpValue, (*QBarDataItem).DestroyQBarDataItem)
	return tmpValue
}

func NewQBarDataItem4(other QBarDataItem_ITF) *QBarDataItem {
	var tmpValue = NewQBarDataItemFromPointer(C.QBarDataItem_NewQBarDataItem4(PointerFromQBarDataItem(other)))
	runtime.SetFinalizer(tmpValue, (*QBarDataItem).DestroyQBarDataItem)
	return tmpValue
}

func NewQBarDataItem2(value float32) *QBarDataItem {
	var tmpValue = NewQBarDataItemFromPointer(C.QBarDataItem_NewQBarDataItem2(C.float(value)))
	runtime.SetFinalizer(tmpValue, (*QBarDataItem).DestroyQBarDataItem)
	return tmpValue
}

func NewQBarDataItem3(value float32, angle float32) *QBarDataItem {
	var tmpValue = NewQBarDataItemFromPointer(C.QBarDataItem_NewQBarDataItem3(C.float(value), C.float(angle)))
	runtime.SetFinalizer(tmpValue, (*QBarDataItem).DestroyQBarDataItem)
	return tmpValue
}

func (ptr *QBarDataItem) SetRotation(angle float32) {
	if ptr.Pointer() != nil {
		C.QBarDataItem_SetRotation(ptr.Pointer(), C.float(angle))
	}
}

func (ptr *QBarDataItem) SetValue(val float32) {
	if ptr.Pointer() != nil {
		C.QBarDataItem_SetValue(ptr.Pointer(), C.float(val))
	}
}

func (ptr *QBarDataItem) DestroyQBarDataItem() {
	if ptr.Pointer() != nil {
		C.QBarDataItem_DestroyQBarDataItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBarDataItem) Rotation() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QBarDataItem_Rotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBarDataItem) Value() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QBarDataItem_Value(ptr.Pointer()))
	}
	return 0
}

type QBarDataProxy struct {
	ptr unsafe.Pointer
}

type QBarDataProxy_ITF interface {
	QBarDataProxy_PTR() *QBarDataProxy
}

func (ptr *QBarDataProxy) QBarDataProxy_PTR() *QBarDataProxy {
	return ptr
}

func (ptr *QBarDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBarDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQBarDataProxy(ptr QBarDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBarDataProxy_PTR().Pointer()
	}
	return nil
}

func NewQBarDataProxyFromPointer(ptr unsafe.Pointer) *QBarDataProxy {
	var n = new(QBarDataProxy)
	n.SetPointer(ptr)
	return n
}
func NewQBarDataProxy(parent core.QObject_ITF) *QBarDataProxy {
	return NewQBarDataProxyFromPointer(C.QBarDataProxy_NewQBarDataProxy(core.PointerFromQObject(parent)))
}

//export callbackQBarDataProxy_ArrayReset
func callbackQBarDataProxy_ArrayReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "arrayReset"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBarDataProxy) ConnectArrayReset(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "arrayReset") {
			C.QBarDataProxy_ConnectArrayReset(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "arrayReset"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "arrayReset", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "arrayReset", f)
		}
	}
}

func (ptr *QBarDataProxy) DisconnectArrayReset() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_DisconnectArrayReset(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "arrayReset")
	}
}

func (ptr *QBarDataProxy) ArrayReset() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_ArrayReset(ptr.Pointer())
	}
}

//export callbackQBarDataProxy_ColumnLabelsChanged
func callbackQBarDataProxy_ColumnLabelsChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "columnLabelsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBarDataProxy) ConnectColumnLabelsChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "columnLabelsChanged") {
			C.QBarDataProxy_ConnectColumnLabelsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "columnLabelsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnLabelsChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnLabelsChanged", f)
		}
	}
}

func (ptr *QBarDataProxy) DisconnectColumnLabelsChanged() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_DisconnectColumnLabelsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "columnLabelsChanged")
	}
}

func (ptr *QBarDataProxy) ColumnLabelsChanged() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_ColumnLabelsChanged(ptr.Pointer())
	}
}

//export callbackQBarDataProxy_ItemChanged
func callbackQBarDataProxy_ItemChanged(ptr unsafe.Pointer, rowIndex C.int, columnIndex C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemChanged"); signal != nil {
		signal.(func(int, int))(int(int32(rowIndex)), int(int32(columnIndex)))
	}

}

func (ptr *QBarDataProxy) ConnectItemChanged(f func(rowIndex int, columnIndex int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemChanged") {
			C.QBarDataProxy_ConnectItemChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemChanged", func(rowIndex int, columnIndex int) {
				signal.(func(int, int))(rowIndex, columnIndex)
				f(rowIndex, columnIndex)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemChanged", f)
		}
	}
}

func (ptr *QBarDataProxy) DisconnectItemChanged() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_DisconnectItemChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemChanged")
	}
}

func (ptr *QBarDataProxy) ItemChanged(rowIndex int, columnIndex int) {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_ItemChanged(ptr.Pointer(), C.int(int32(rowIndex)), C.int(int32(columnIndex)))
	}
}

func (ptr *QBarDataProxy) RemoveRows(rowIndex int, removeCount int, removeLabels bool) {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_RemoveRows(ptr.Pointer(), C.int(int32(rowIndex)), C.int(int32(removeCount)), C.char(int8(qt.GoBoolToInt(removeLabels))))
	}
}

func (ptr *QBarDataProxy) ResetArray() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_ResetArray(ptr.Pointer())
	}
}

//export callbackQBarDataProxy_RowCountChanged
func callbackQBarDataProxy_RowCountChanged(ptr unsafe.Pointer, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowCountChanged"); signal != nil {
		signal.(func(int))(int(int32(count)))
	}

}

func (ptr *QBarDataProxy) ConnectRowCountChanged(f func(count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowCountChanged") {
			C.QBarDataProxy_ConnectRowCountChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowCountChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowCountChanged", func(count int) {
				signal.(func(int))(count)
				f(count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowCountChanged", f)
		}
	}
}

func (ptr *QBarDataProxy) DisconnectRowCountChanged() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_DisconnectRowCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowCountChanged")
	}
}

func (ptr *QBarDataProxy) RowCountChanged(count int) {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_RowCountChanged(ptr.Pointer(), C.int(int32(count)))
	}
}

//export callbackQBarDataProxy_RowLabelsChanged
func callbackQBarDataProxy_RowLabelsChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowLabelsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBarDataProxy) ConnectRowLabelsChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowLabelsChanged") {
			C.QBarDataProxy_ConnectRowLabelsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowLabelsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowLabelsChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowLabelsChanged", f)
		}
	}
}

func (ptr *QBarDataProxy) DisconnectRowLabelsChanged() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_DisconnectRowLabelsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowLabelsChanged")
	}
}

func (ptr *QBarDataProxy) RowLabelsChanged() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_RowLabelsChanged(ptr.Pointer())
	}
}

//export callbackQBarDataProxy_RowsAdded
func callbackQBarDataProxy_RowsAdded(ptr unsafe.Pointer, startIndex C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowsAdded"); signal != nil {
		signal.(func(int, int))(int(int32(startIndex)), int(int32(count)))
	}

}

func (ptr *QBarDataProxy) ConnectRowsAdded(f func(startIndex int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowsAdded") {
			C.QBarDataProxy_ConnectRowsAdded(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowsAdded"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsAdded", func(startIndex int, count int) {
				signal.(func(int, int))(startIndex, count)
				f(startIndex, count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsAdded", f)
		}
	}
}

func (ptr *QBarDataProxy) DisconnectRowsAdded() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_DisconnectRowsAdded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowsAdded")
	}
}

func (ptr *QBarDataProxy) RowsAdded(startIndex int, count int) {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_RowsAdded(ptr.Pointer(), C.int(int32(startIndex)), C.int(int32(count)))
	}
}

//export callbackQBarDataProxy_RowsChanged
func callbackQBarDataProxy_RowsChanged(ptr unsafe.Pointer, startIndex C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowsChanged"); signal != nil {
		signal.(func(int, int))(int(int32(startIndex)), int(int32(count)))
	}

}

func (ptr *QBarDataProxy) ConnectRowsChanged(f func(startIndex int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowsChanged") {
			C.QBarDataProxy_ConnectRowsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsChanged", func(startIndex int, count int) {
				signal.(func(int, int))(startIndex, count)
				f(startIndex, count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsChanged", f)
		}
	}
}

func (ptr *QBarDataProxy) DisconnectRowsChanged() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_DisconnectRowsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowsChanged")
	}
}

func (ptr *QBarDataProxy) RowsChanged(startIndex int, count int) {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_RowsChanged(ptr.Pointer(), C.int(int32(startIndex)), C.int(int32(count)))
	}
}

//export callbackQBarDataProxy_RowsInserted
func callbackQBarDataProxy_RowsInserted(ptr unsafe.Pointer, startIndex C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowsInserted"); signal != nil {
		signal.(func(int, int))(int(int32(startIndex)), int(int32(count)))
	}

}

func (ptr *QBarDataProxy) ConnectRowsInserted(f func(startIndex int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowsInserted") {
			C.QBarDataProxy_ConnectRowsInserted(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowsInserted"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsInserted", func(startIndex int, count int) {
				signal.(func(int, int))(startIndex, count)
				f(startIndex, count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsInserted", f)
		}
	}
}

func (ptr *QBarDataProxy) DisconnectRowsInserted() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_DisconnectRowsInserted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowsInserted")
	}
}

func (ptr *QBarDataProxy) RowsInserted(startIndex int, count int) {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_RowsInserted(ptr.Pointer(), C.int(int32(startIndex)), C.int(int32(count)))
	}
}

//export callbackQBarDataProxy_RowsRemoved
func callbackQBarDataProxy_RowsRemoved(ptr unsafe.Pointer, startIndex C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowsRemoved"); signal != nil {
		signal.(func(int, int))(int(int32(startIndex)), int(int32(count)))
	}

}

func (ptr *QBarDataProxy) ConnectRowsRemoved(f func(startIndex int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowsRemoved") {
			C.QBarDataProxy_ConnectRowsRemoved(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowsRemoved"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsRemoved", func(startIndex int, count int) {
				signal.(func(int, int))(startIndex, count)
				f(startIndex, count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsRemoved", f)
		}
	}
}

func (ptr *QBarDataProxy) DisconnectRowsRemoved() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_DisconnectRowsRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowsRemoved")
	}
}

func (ptr *QBarDataProxy) RowsRemoved(startIndex int, count int) {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_RowsRemoved(ptr.Pointer(), C.int(int32(startIndex)), C.int(int32(count)))
	}
}

//export callbackQBarDataProxy_SeriesChanged
func callbackQBarDataProxy_SeriesChanged(ptr unsafe.Pointer, series unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "seriesChanged"); signal != nil {
		signal.(func(*QBar3DSeries))(NewQBar3DSeriesFromPointer(series))
	}

}

func (ptr *QBarDataProxy) ConnectSeriesChanged(f func(series *QBar3DSeries)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged") {
			C.QBarDataProxy_ConnectSeriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged", func(series *QBar3DSeries) {
				signal.(func(*QBar3DSeries))(series)
				f(series)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged", f)
		}
	}
}

func (ptr *QBarDataProxy) DisconnectSeriesChanged() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_DisconnectSeriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged")
	}
}

func (ptr *QBarDataProxy) SeriesChanged(series QBar3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_SeriesChanged(ptr.Pointer(), PointerFromQBar3DSeries(series))
	}
}

func (ptr *QBarDataProxy) SetColumnLabels(labels []string) {
	if ptr.Pointer() != nil {
		var labelsC = C.CString(strings.Join(labels, "|"))
		defer C.free(unsafe.Pointer(labelsC))
		C.QBarDataProxy_SetColumnLabels(ptr.Pointer(), labelsC)
	}
}

func (ptr *QBarDataProxy) SetItem2(position core.QPoint_ITF, item QBarDataItem_ITF) {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_SetItem2(ptr.Pointer(), core.PointerFromQPoint(position), PointerFromQBarDataItem(item))
	}
}

func (ptr *QBarDataProxy) SetItem(rowIndex int, columnIndex int, item QBarDataItem_ITF) {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_SetItem(ptr.Pointer(), C.int(int32(rowIndex)), C.int(int32(columnIndex)), PointerFromQBarDataItem(item))
	}
}

func (ptr *QBarDataProxy) SetRowLabels(labels []string) {
	if ptr.Pointer() != nil {
		var labelsC = C.CString(strings.Join(labels, "|"))
		defer C.free(unsafe.Pointer(labelsC))
		C.QBarDataProxy_SetRowLabels(ptr.Pointer(), labelsC)
	}
}

//export callbackQBarDataProxy_DestroyQBarDataProxy
func callbackQBarDataProxy_DestroyQBarDataProxy(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QBarDataProxy"); signal != nil {
		signal.(func())()
	} else {
		NewQBarDataProxyFromPointer(ptr).DestroyQBarDataProxyDefault()
	}
}

func (ptr *QBarDataProxy) ConnectDestroyQBarDataProxy(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QBarDataProxy"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QBarDataProxy", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QBarDataProxy", f)
		}
	}
}

func (ptr *QBarDataProxy) DisconnectDestroyQBarDataProxy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QBarDataProxy")
	}
}

func (ptr *QBarDataProxy) DestroyQBarDataProxy() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_DestroyQBarDataProxy(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBarDataProxy) DestroyQBarDataProxyDefault() {
	if ptr.Pointer() != nil {
		C.QBarDataProxy_DestroyQBarDataProxyDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBarDataProxy) Series() *QBar3DSeries {
	if ptr.Pointer() != nil {
		return NewQBar3DSeriesFromPointer(C.QBarDataProxy_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBarDataProxy) ColumnLabels() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QBarDataProxy_ColumnLabels(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QBarDataProxy) RowLabels() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QBarDataProxy_RowLabels(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QBarDataProxy) ItemAt2(position core.QPoint_ITF) *QBarDataItem {
	if ptr.Pointer() != nil {
		return NewQBarDataItemFromPointer(C.QBarDataProxy_ItemAt2(ptr.Pointer(), core.PointerFromQPoint(position)))
	}
	return nil
}

func (ptr *QBarDataProxy) ItemAt(rowIndex int, columnIndex int) *QBarDataItem {
	if ptr.Pointer() != nil {
		return NewQBarDataItemFromPointer(C.QBarDataProxy_ItemAt(ptr.Pointer(), C.int(int32(rowIndex)), C.int(int32(columnIndex))))
	}
	return nil
}

func (ptr *QBarDataProxy) RowCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBarDataProxy_RowCount(ptr.Pointer())))
	}
	return 0
}

type QCategory3DAxis struct {
	ptr unsafe.Pointer
}

type QCategory3DAxis_ITF interface {
	QCategory3DAxis_PTR() *QCategory3DAxis
}

func (ptr *QCategory3DAxis) QCategory3DAxis_PTR() *QCategory3DAxis {
	return ptr
}

func (ptr *QCategory3DAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCategory3DAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCategory3DAxis(ptr QCategory3DAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCategory3DAxis_PTR().Pointer()
	}
	return nil
}

func NewQCategory3DAxisFromPointer(ptr unsafe.Pointer) *QCategory3DAxis {
	var n = new(QCategory3DAxis)
	n.SetPointer(ptr)
	return n
}
func NewQCategory3DAxis(parent core.QObject_ITF) *QCategory3DAxis {
	return NewQCategory3DAxisFromPointer(C.QCategory3DAxis_NewQCategory3DAxis(core.PointerFromQObject(parent)))
}

//export callbackQCategory3DAxis_LabelsChanged
func callbackQCategory3DAxis_LabelsChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "labelsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCategory3DAxis) ConnectLabelsChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "labelsChanged") {
			C.QCategory3DAxis_ConnectLabelsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "labelsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelsChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelsChanged", f)
		}
	}
}

func (ptr *QCategory3DAxis) DisconnectLabelsChanged() {
	if ptr.Pointer() != nil {
		C.QCategory3DAxis_DisconnectLabelsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "labelsChanged")
	}
}

func (ptr *QCategory3DAxis) LabelsChanged() {
	if ptr.Pointer() != nil {
		C.QCategory3DAxis_LabelsChanged(ptr.Pointer())
	}
}

func (ptr *QCategory3DAxis) SetLabels(labels []string) {
	if ptr.Pointer() != nil {
		var labelsC = C.CString(strings.Join(labels, "|"))
		defer C.free(unsafe.Pointer(labelsC))
		C.QCategory3DAxis_SetLabels(ptr.Pointer(), labelsC)
	}
}

//export callbackQCategory3DAxis_DestroyQCategory3DAxis
func callbackQCategory3DAxis_DestroyQCategory3DAxis(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QCategory3DAxis"); signal != nil {
		signal.(func())()
	} else {
		NewQCategory3DAxisFromPointer(ptr).DestroyQCategory3DAxisDefault()
	}
}

func (ptr *QCategory3DAxis) ConnectDestroyQCategory3DAxis(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QCategory3DAxis"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QCategory3DAxis", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QCategory3DAxis", f)
		}
	}
}

func (ptr *QCategory3DAxis) DisconnectDestroyQCategory3DAxis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QCategory3DAxis")
	}
}

func (ptr *QCategory3DAxis) DestroyQCategory3DAxis() {
	if ptr.Pointer() != nil {
		C.QCategory3DAxis_DestroyQCategory3DAxis(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCategory3DAxis) DestroyQCategory3DAxisDefault() {
	if ptr.Pointer() != nil {
		C.QCategory3DAxis_DestroyQCategory3DAxisDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCategory3DAxis) Labels() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QCategory3DAxis_Labels(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

type QCustom3DItem struct {
	ptr unsafe.Pointer
}

type QCustom3DItem_ITF interface {
	QCustom3DItem_PTR() *QCustom3DItem
}

func (ptr *QCustom3DItem) QCustom3DItem_PTR() *QCustom3DItem {
	return ptr
}

func (ptr *QCustom3DItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCustom3DItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCustom3DItem(ptr QCustom3DItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCustom3DItem_PTR().Pointer()
	}
	return nil
}

func NewQCustom3DItemFromPointer(ptr unsafe.Pointer) *QCustom3DItem {
	var n = new(QCustom3DItem)
	n.SetPointer(ptr)
	return n
}
func NewQCustom3DItem(parent core.QObject_ITF) *QCustom3DItem {
	return NewQCustom3DItemFromPointer(C.QCustom3DItem_NewQCustom3DItem(core.PointerFromQObject(parent)))
}

func NewQCustom3DItem2(meshFile string, position gui.QVector3D_ITF, scaling gui.QVector3D_ITF, rotation gui.QQuaternion_ITF, texture gui.QImage_ITF, parent core.QObject_ITF) *QCustom3DItem {
	var meshFileC *C.char
	if meshFile != "" {
		meshFileC = C.CString(meshFile)
		defer C.free(unsafe.Pointer(meshFileC))
	}
	return NewQCustom3DItemFromPointer(C.QCustom3DItem_NewQCustom3DItem2(meshFileC, gui.PointerFromQVector3D(position), gui.PointerFromQVector3D(scaling), gui.PointerFromQQuaternion(rotation), gui.PointerFromQImage(texture), core.PointerFromQObject(parent)))
}

func (ptr *QCustom3DItem) Rotation() *gui.QQuaternion {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQQuaternionFromPointer(C.QCustom3DItem_Rotation(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QQuaternion).DestroyQQuaternion)
		return tmpValue
	}
	return nil
}

//export callbackQCustom3DItem_MeshFileChanged
func callbackQCustom3DItem_MeshFileChanged(ptr unsafe.Pointer, meshFile C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "meshFileChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(meshFile))
	}

}

func (ptr *QCustom3DItem) ConnectMeshFileChanged(f func(meshFile string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "meshFileChanged") {
			C.QCustom3DItem_ConnectMeshFileChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "meshFileChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "meshFileChanged", func(meshFile string) {
				signal.(func(string))(meshFile)
				f(meshFile)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "meshFileChanged", f)
		}
	}
}

func (ptr *QCustom3DItem) DisconnectMeshFileChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_DisconnectMeshFileChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "meshFileChanged")
	}
}

func (ptr *QCustom3DItem) MeshFileChanged(meshFile string) {
	if ptr.Pointer() != nil {
		var meshFileC *C.char
		if meshFile != "" {
			meshFileC = C.CString(meshFile)
			defer C.free(unsafe.Pointer(meshFileC))
		}
		C.QCustom3DItem_MeshFileChanged(ptr.Pointer(), meshFileC)
	}
}

//export callbackQCustom3DItem_PositionAbsoluteChanged
func callbackQCustom3DItem_PositionAbsoluteChanged(ptr unsafe.Pointer, positionAbsolute C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "positionAbsoluteChanged"); signal != nil {
		signal.(func(bool))(int8(positionAbsolute) != 0)
	}

}

func (ptr *QCustom3DItem) ConnectPositionAbsoluteChanged(f func(positionAbsolute bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "positionAbsoluteChanged") {
			C.QCustom3DItem_ConnectPositionAbsoluteChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "positionAbsoluteChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "positionAbsoluteChanged", func(positionAbsolute bool) {
				signal.(func(bool))(positionAbsolute)
				f(positionAbsolute)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "positionAbsoluteChanged", f)
		}
	}
}

func (ptr *QCustom3DItem) DisconnectPositionAbsoluteChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_DisconnectPositionAbsoluteChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "positionAbsoluteChanged")
	}
}

func (ptr *QCustom3DItem) PositionAbsoluteChanged(positionAbsolute bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_PositionAbsoluteChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(positionAbsolute))))
	}
}

//export callbackQCustom3DItem_PositionChanged
func callbackQCustom3DItem_PositionChanged(ptr unsafe.Pointer, position unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "positionChanged"); signal != nil {
		signal.(func(*gui.QVector3D))(gui.NewQVector3DFromPointer(position))
	}

}

func (ptr *QCustom3DItem) ConnectPositionChanged(f func(position *gui.QVector3D)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "positionChanged") {
			C.QCustom3DItem_ConnectPositionChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "positionChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "positionChanged", func(position *gui.QVector3D) {
				signal.(func(*gui.QVector3D))(position)
				f(position)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "positionChanged", f)
		}
	}
}

func (ptr *QCustom3DItem) DisconnectPositionChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_DisconnectPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "positionChanged")
	}
}

func (ptr *QCustom3DItem) PositionChanged(position gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_PositionChanged(ptr.Pointer(), gui.PointerFromQVector3D(position))
	}
}

//export callbackQCustom3DItem_RotationChanged
func callbackQCustom3DItem_RotationChanged(ptr unsafe.Pointer, rotation unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rotationChanged"); signal != nil {
		signal.(func(*gui.QQuaternion))(gui.NewQQuaternionFromPointer(rotation))
	}

}

func (ptr *QCustom3DItem) ConnectRotationChanged(f func(rotation *gui.QQuaternion)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rotationChanged") {
			C.QCustom3DItem_ConnectRotationChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rotationChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationChanged", func(rotation *gui.QQuaternion) {
				signal.(func(*gui.QQuaternion))(rotation)
				f(rotation)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationChanged", f)
		}
	}
}

func (ptr *QCustom3DItem) DisconnectRotationChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_DisconnectRotationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rotationChanged")
	}
}

func (ptr *QCustom3DItem) RotationChanged(rotation gui.QQuaternion_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_RotationChanged(ptr.Pointer(), gui.PointerFromQQuaternion(rotation))
	}
}

//export callbackQCustom3DItem_ScalingAbsoluteChanged
func callbackQCustom3DItem_ScalingAbsoluteChanged(ptr unsafe.Pointer, scalingAbsolute C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "scalingAbsoluteChanged"); signal != nil {
		signal.(func(bool))(int8(scalingAbsolute) != 0)
	}

}

func (ptr *QCustom3DItem) ConnectScalingAbsoluteChanged(f func(scalingAbsolute bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "scalingAbsoluteChanged") {
			C.QCustom3DItem_ConnectScalingAbsoluteChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "scalingAbsoluteChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "scalingAbsoluteChanged", func(scalingAbsolute bool) {
				signal.(func(bool))(scalingAbsolute)
				f(scalingAbsolute)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "scalingAbsoluteChanged", f)
		}
	}
}

func (ptr *QCustom3DItem) DisconnectScalingAbsoluteChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_DisconnectScalingAbsoluteChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "scalingAbsoluteChanged")
	}
}

func (ptr *QCustom3DItem) ScalingAbsoluteChanged(scalingAbsolute bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_ScalingAbsoluteChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(scalingAbsolute))))
	}
}

//export callbackQCustom3DItem_ScalingChanged
func callbackQCustom3DItem_ScalingChanged(ptr unsafe.Pointer, scaling unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "scalingChanged"); signal != nil {
		signal.(func(*gui.QVector3D))(gui.NewQVector3DFromPointer(scaling))
	}

}

func (ptr *QCustom3DItem) ConnectScalingChanged(f func(scaling *gui.QVector3D)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "scalingChanged") {
			C.QCustom3DItem_ConnectScalingChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "scalingChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "scalingChanged", func(scaling *gui.QVector3D) {
				signal.(func(*gui.QVector3D))(scaling)
				f(scaling)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "scalingChanged", f)
		}
	}
}

func (ptr *QCustom3DItem) DisconnectScalingChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_DisconnectScalingChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "scalingChanged")
	}
}

func (ptr *QCustom3DItem) ScalingChanged(scaling gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_ScalingChanged(ptr.Pointer(), gui.PointerFromQVector3D(scaling))
	}
}

func (ptr *QCustom3DItem) SetMeshFile(meshFile string) {
	if ptr.Pointer() != nil {
		var meshFileC *C.char
		if meshFile != "" {
			meshFileC = C.CString(meshFile)
			defer C.free(unsafe.Pointer(meshFileC))
		}
		C.QCustom3DItem_SetMeshFile(ptr.Pointer(), meshFileC)
	}
}

func (ptr *QCustom3DItem) SetPosition(position gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_SetPosition(ptr.Pointer(), gui.PointerFromQVector3D(position))
	}
}

func (ptr *QCustom3DItem) SetPositionAbsolute(positionAbsolute bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_SetPositionAbsolute(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(positionAbsolute))))
	}
}

func (ptr *QCustom3DItem) SetRotation(rotation gui.QQuaternion_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_SetRotation(ptr.Pointer(), gui.PointerFromQQuaternion(rotation))
	}
}

func (ptr *QCustom3DItem) SetRotationAxisAndAngle(axis gui.QVector3D_ITF, angle float32) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_SetRotationAxisAndAngle(ptr.Pointer(), gui.PointerFromQVector3D(axis), C.float(angle))
	}
}

func (ptr *QCustom3DItem) SetScaling(scaling gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_SetScaling(ptr.Pointer(), gui.PointerFromQVector3D(scaling))
	}
}

func (ptr *QCustom3DItem) SetScalingAbsolute(scalingAbsolute bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_SetScalingAbsolute(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(scalingAbsolute))))
	}
}

func (ptr *QCustom3DItem) SetShadowCasting(enabled bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_SetShadowCasting(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QCustom3DItem) SetTextureFile(textureFile string) {
	if ptr.Pointer() != nil {
		var textureFileC *C.char
		if textureFile != "" {
			textureFileC = C.CString(textureFile)
			defer C.free(unsafe.Pointer(textureFileC))
		}
		C.QCustom3DItem_SetTextureFile(ptr.Pointer(), textureFileC)
	}
}

func (ptr *QCustom3DItem) SetTextureImage(textureImage gui.QImage_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_SetTextureImage(ptr.Pointer(), gui.PointerFromQImage(textureImage))
	}
}

func (ptr *QCustom3DItem) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQCustom3DItem_ShadowCastingChanged
func callbackQCustom3DItem_ShadowCastingChanged(ptr unsafe.Pointer, shadowCasting C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "shadowCastingChanged"); signal != nil {
		signal.(func(bool))(int8(shadowCasting) != 0)
	}

}

func (ptr *QCustom3DItem) ConnectShadowCastingChanged(f func(shadowCasting bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "shadowCastingChanged") {
			C.QCustom3DItem_ConnectShadowCastingChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "shadowCastingChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "shadowCastingChanged", func(shadowCasting bool) {
				signal.(func(bool))(shadowCasting)
				f(shadowCasting)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "shadowCastingChanged", f)
		}
	}
}

func (ptr *QCustom3DItem) DisconnectShadowCastingChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_DisconnectShadowCastingChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "shadowCastingChanged")
	}
}

func (ptr *QCustom3DItem) ShadowCastingChanged(shadowCasting bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_ShadowCastingChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(shadowCasting))))
	}
}

//export callbackQCustom3DItem_TextureFileChanged
func callbackQCustom3DItem_TextureFileChanged(ptr unsafe.Pointer, textureFile C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "textureFileChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(textureFile))
	}

}

func (ptr *QCustom3DItem) ConnectTextureFileChanged(f func(textureFile string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "textureFileChanged") {
			C.QCustom3DItem_ConnectTextureFileChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "textureFileChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureFileChanged", func(textureFile string) {
				signal.(func(string))(textureFile)
				f(textureFile)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureFileChanged", f)
		}
	}
}

func (ptr *QCustom3DItem) DisconnectTextureFileChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_DisconnectTextureFileChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "textureFileChanged")
	}
}

func (ptr *QCustom3DItem) TextureFileChanged(textureFile string) {
	if ptr.Pointer() != nil {
		var textureFileC *C.char
		if textureFile != "" {
			textureFileC = C.CString(textureFile)
			defer C.free(unsafe.Pointer(textureFileC))
		}
		C.QCustom3DItem_TextureFileChanged(ptr.Pointer(), textureFileC)
	}
}

//export callbackQCustom3DItem_VisibleChanged
func callbackQCustom3DItem_VisibleChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "visibleChanged"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QCustom3DItem) ConnectVisibleChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "visibleChanged") {
			C.QCustom3DItem_ConnectVisibleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "visibleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "visibleChanged", func(visible bool) {
				signal.(func(bool))(visible)
				f(visible)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "visibleChanged", f)
		}
	}
}

func (ptr *QCustom3DItem) DisconnectVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "visibleChanged")
	}
}

func (ptr *QCustom3DItem) VisibleChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_VisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQCustom3DItem_DestroyQCustom3DItem
func callbackQCustom3DItem_DestroyQCustom3DItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QCustom3DItem"); signal != nil {
		signal.(func())()
	} else {
		NewQCustom3DItemFromPointer(ptr).DestroyQCustom3DItemDefault()
	}
}

func (ptr *QCustom3DItem) ConnectDestroyQCustom3DItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QCustom3DItem"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QCustom3DItem", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QCustom3DItem", f)
		}
	}
}

func (ptr *QCustom3DItem) DisconnectDestroyQCustom3DItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QCustom3DItem")
	}
}

func (ptr *QCustom3DItem) DestroyQCustom3DItem() {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_DestroyQCustom3DItem(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCustom3DItem) DestroyQCustom3DItemDefault() {
	if ptr.Pointer() != nil {
		C.QCustom3DItem_DestroyQCustom3DItemDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCustom3DItem) MeshFile() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QCustom3DItem_MeshFile(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCustom3DItem) TextureFile() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QCustom3DItem_TextureFile(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCustom3DItem) Position() *gui.QVector3D {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQVector3DFromPointer(C.QCustom3DItem_Position(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QVector3D).DestroyQVector3D)
		return tmpValue
	}
	return nil
}

func (ptr *QCustom3DItem) Scaling() *gui.QVector3D {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQVector3DFromPointer(C.QCustom3DItem_Scaling(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QVector3D).DestroyQVector3D)
		return tmpValue
	}
	return nil
}

func (ptr *QCustom3DItem) IsPositionAbsolute() bool {
	if ptr.Pointer() != nil {
		return C.QCustom3DItem_IsPositionAbsolute(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCustom3DItem) IsScalingAbsolute() bool {
	if ptr.Pointer() != nil {
		return C.QCustom3DItem_IsScalingAbsolute(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCustom3DItem) IsShadowCasting() bool {
	if ptr.Pointer() != nil {
		return C.QCustom3DItem_IsShadowCasting(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCustom3DItem) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QCustom3DItem_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

type QCustom3DLabel struct {
	ptr unsafe.Pointer
}

type QCustom3DLabel_ITF interface {
	QCustom3DLabel_PTR() *QCustom3DLabel
}

func (ptr *QCustom3DLabel) QCustom3DLabel_PTR() *QCustom3DLabel {
	return ptr
}

func (ptr *QCustom3DLabel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCustom3DLabel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCustom3DLabel(ptr QCustom3DLabel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCustom3DLabel_PTR().Pointer()
	}
	return nil
}

func NewQCustom3DLabelFromPointer(ptr unsafe.Pointer) *QCustom3DLabel {
	var n = new(QCustom3DLabel)
	n.SetPointer(ptr)
	return n
}
func NewQCustom3DLabel(parent core.QObject_ITF) *QCustom3DLabel {
	return NewQCustom3DLabelFromPointer(C.QCustom3DLabel_NewQCustom3DLabel(core.PointerFromQObject(parent)))
}

func NewQCustom3DLabel2(text string, font gui.QFont_ITF, position gui.QVector3D_ITF, scaling gui.QVector3D_ITF, rotation gui.QQuaternion_ITF, parent core.QObject_ITF) *QCustom3DLabel {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	return NewQCustom3DLabelFromPointer(C.QCustom3DLabel_NewQCustom3DLabel2(textC, gui.PointerFromQFont(font), gui.PointerFromQVector3D(position), gui.PointerFromQVector3D(scaling), gui.PointerFromQQuaternion(rotation), core.PointerFromQObject(parent)))
}

//export callbackQCustom3DLabel_BackgroundColorChanged
func callbackQCustom3DLabel_BackgroundColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "backgroundColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QCustom3DLabel) ConnectBackgroundColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "backgroundColorChanged") {
			C.QCustom3DLabel_ConnectBackgroundColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "backgroundColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "backgroundColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "backgroundColorChanged", f)
		}
	}
}

func (ptr *QCustom3DLabel) DisconnectBackgroundColorChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_DisconnectBackgroundColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "backgroundColorChanged")
	}
}

func (ptr *QCustom3DLabel) BackgroundColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_BackgroundColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQCustom3DLabel_BackgroundEnabledChanged
func callbackQCustom3DLabel_BackgroundEnabledChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "backgroundEnabledChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QCustom3DLabel) ConnectBackgroundEnabledChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "backgroundEnabledChanged") {
			C.QCustom3DLabel_ConnectBackgroundEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "backgroundEnabledChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "backgroundEnabledChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "backgroundEnabledChanged", f)
		}
	}
}

func (ptr *QCustom3DLabel) DisconnectBackgroundEnabledChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_DisconnectBackgroundEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "backgroundEnabledChanged")
	}
}

func (ptr *QCustom3DLabel) BackgroundEnabledChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_BackgroundEnabledChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQCustom3DLabel_BorderEnabledChanged
func callbackQCustom3DLabel_BorderEnabledChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "borderEnabledChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QCustom3DLabel) ConnectBorderEnabledChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "borderEnabledChanged") {
			C.QCustom3DLabel_ConnectBorderEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "borderEnabledChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "borderEnabledChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "borderEnabledChanged", f)
		}
	}
}

func (ptr *QCustom3DLabel) DisconnectBorderEnabledChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_DisconnectBorderEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "borderEnabledChanged")
	}
}

func (ptr *QCustom3DLabel) BorderEnabledChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_BorderEnabledChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQCustom3DLabel_FacingCameraChanged
func callbackQCustom3DLabel_FacingCameraChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "facingCameraChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QCustom3DLabel) ConnectFacingCameraChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "facingCameraChanged") {
			C.QCustom3DLabel_ConnectFacingCameraChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "facingCameraChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "facingCameraChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "facingCameraChanged", f)
		}
	}
}

func (ptr *QCustom3DLabel) DisconnectFacingCameraChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_DisconnectFacingCameraChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "facingCameraChanged")
	}
}

func (ptr *QCustom3DLabel) FacingCameraChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_FacingCameraChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQCustom3DLabel_FontChanged
func callbackQCustom3DLabel_FontChanged(ptr unsafe.Pointer, font unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "fontChanged"); signal != nil {
		signal.(func(*gui.QFont))(gui.NewQFontFromPointer(font))
	}

}

func (ptr *QCustom3DLabel) ConnectFontChanged(f func(font *gui.QFont)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "fontChanged") {
			C.QCustom3DLabel_ConnectFontChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "fontChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "fontChanged", func(font *gui.QFont) {
				signal.(func(*gui.QFont))(font)
				f(font)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "fontChanged", f)
		}
	}
}

func (ptr *QCustom3DLabel) DisconnectFontChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_DisconnectFontChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "fontChanged")
	}
}

func (ptr *QCustom3DLabel) FontChanged(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_FontChanged(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QCustom3DLabel) SetBackgroundColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_SetBackgroundColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QCustom3DLabel) SetBackgroundEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_SetBackgroundEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QCustom3DLabel) SetBorderEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_SetBorderEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QCustom3DLabel) SetFacingCamera(enabled bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_SetFacingCamera(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QCustom3DLabel) SetFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QCustom3DLabel) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QCustom3DLabel_SetText(ptr.Pointer(), textC)
	}
}

func (ptr *QCustom3DLabel) SetTextColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_SetTextColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQCustom3DLabel_TextChanged
func callbackQCustom3DLabel_TextChanged(ptr unsafe.Pointer, text C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "textChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(text))
	}

}

func (ptr *QCustom3DLabel) ConnectTextChanged(f func(text string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "textChanged") {
			C.QCustom3DLabel_ConnectTextChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "textChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textChanged", func(text string) {
				signal.(func(string))(text)
				f(text)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textChanged", f)
		}
	}
}

func (ptr *QCustom3DLabel) DisconnectTextChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_DisconnectTextChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "textChanged")
	}
}

func (ptr *QCustom3DLabel) TextChanged(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QCustom3DLabel_TextChanged(ptr.Pointer(), textC)
	}
}

//export callbackQCustom3DLabel_TextColorChanged
func callbackQCustom3DLabel_TextColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "textColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QCustom3DLabel) ConnectTextColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "textColorChanged") {
			C.QCustom3DLabel_ConnectTextColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "textColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textColorChanged", f)
		}
	}
}

func (ptr *QCustom3DLabel) DisconnectTextColorChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_DisconnectTextColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "textColorChanged")
	}
}

func (ptr *QCustom3DLabel) TextColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_TextColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQCustom3DLabel_DestroyQCustom3DLabel
func callbackQCustom3DLabel_DestroyQCustom3DLabel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QCustom3DLabel"); signal != nil {
		signal.(func())()
	} else {
		NewQCustom3DLabelFromPointer(ptr).DestroyQCustom3DLabelDefault()
	}
}

func (ptr *QCustom3DLabel) ConnectDestroyQCustom3DLabel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QCustom3DLabel"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QCustom3DLabel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QCustom3DLabel", f)
		}
	}
}

func (ptr *QCustom3DLabel) DisconnectDestroyQCustom3DLabel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QCustom3DLabel")
	}
}

func (ptr *QCustom3DLabel) DestroyQCustom3DLabel() {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_DestroyQCustom3DLabel(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCustom3DLabel) DestroyQCustom3DLabelDefault() {
	if ptr.Pointer() != nil {
		C.QCustom3DLabel_DestroyQCustom3DLabelDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCustom3DLabel) BackgroundColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QCustom3DLabel_BackgroundColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QCustom3DLabel) TextColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QCustom3DLabel_TextColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QCustom3DLabel) Font() *gui.QFont {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQFontFromPointer(C.QCustom3DLabel_Font(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QCustom3DLabel) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QCustom3DLabel_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCustom3DLabel) IsBackgroundEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QCustom3DLabel_IsBackgroundEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCustom3DLabel) IsBorderEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QCustom3DLabel_IsBorderEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCustom3DLabel) IsFacingCamera() bool {
	if ptr.Pointer() != nil {
		return C.QCustom3DLabel_IsFacingCamera(ptr.Pointer()) != 0
	}
	return false
}

type QCustom3DVolume struct {
	ptr unsafe.Pointer
}

type QCustom3DVolume_ITF interface {
	QCustom3DVolume_PTR() *QCustom3DVolume
}

func (ptr *QCustom3DVolume) QCustom3DVolume_PTR() *QCustom3DVolume {
	return ptr
}

func (ptr *QCustom3DVolume) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCustom3DVolume) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCustom3DVolume(ptr QCustom3DVolume_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCustom3DVolume_PTR().Pointer()
	}
	return nil
}

func NewQCustom3DVolumeFromPointer(ptr unsafe.Pointer) *QCustom3DVolume {
	var n = new(QCustom3DVolume)
	n.SetPointer(ptr)
	return n
}
func NewQCustom3DVolume(parent core.QObject_ITF) *QCustom3DVolume {
	return NewQCustom3DVolumeFromPointer(C.QCustom3DVolume_NewQCustom3DVolume(core.PointerFromQObject(parent)))
}

func (ptr *QCustom3DVolume) RenderSlice(axis core.Qt__Axis, index int) *gui.QImage {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQImageFromPointer(C.QCustom3DVolume_RenderSlice(ptr.Pointer(), C.longlong(axis), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

//export callbackQCustom3DVolume_AlphaMultiplierChanged
func callbackQCustom3DVolume_AlphaMultiplierChanged(ptr unsafe.Pointer, mult C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "alphaMultiplierChanged"); signal != nil {
		signal.(func(float32))(float32(mult))
	}

}

func (ptr *QCustom3DVolume) ConnectAlphaMultiplierChanged(f func(mult float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "alphaMultiplierChanged") {
			C.QCustom3DVolume_ConnectAlphaMultiplierChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "alphaMultiplierChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "alphaMultiplierChanged", func(mult float32) {
				signal.(func(float32))(mult)
				f(mult)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "alphaMultiplierChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectAlphaMultiplierChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectAlphaMultiplierChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "alphaMultiplierChanged")
	}
}

func (ptr *QCustom3DVolume) AlphaMultiplierChanged(mult float32) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_AlphaMultiplierChanged(ptr.Pointer(), C.float(mult))
	}
}

//export callbackQCustom3DVolume_ColorTableChanged
func callbackQCustom3DVolume_ColorTableChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "colorTableChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCustom3DVolume) ConnectColorTableChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "colorTableChanged") {
			C.QCustom3DVolume_ConnectColorTableChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "colorTableChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "colorTableChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "colorTableChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectColorTableChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectColorTableChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "colorTableChanged")
	}
}

func (ptr *QCustom3DVolume) ColorTableChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_ColorTableChanged(ptr.Pointer())
	}
}

//export callbackQCustom3DVolume_DrawSliceFramesChanged
func callbackQCustom3DVolume_DrawSliceFramesChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "drawSliceFramesChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QCustom3DVolume) ConnectDrawSliceFramesChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "drawSliceFramesChanged") {
			C.QCustom3DVolume_ConnectDrawSliceFramesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "drawSliceFramesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "drawSliceFramesChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "drawSliceFramesChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectDrawSliceFramesChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectDrawSliceFramesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "drawSliceFramesChanged")
	}
}

func (ptr *QCustom3DVolume) DrawSliceFramesChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DrawSliceFramesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQCustom3DVolume_DrawSlicesChanged
func callbackQCustom3DVolume_DrawSlicesChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "drawSlicesChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QCustom3DVolume) ConnectDrawSlicesChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "drawSlicesChanged") {
			C.QCustom3DVolume_ConnectDrawSlicesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "drawSlicesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "drawSlicesChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "drawSlicesChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectDrawSlicesChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectDrawSlicesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "drawSlicesChanged")
	}
}

func (ptr *QCustom3DVolume) DrawSlicesChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DrawSlicesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQCustom3DVolume_PreserveOpacityChanged
func callbackQCustom3DVolume_PreserveOpacityChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "preserveOpacityChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QCustom3DVolume) ConnectPreserveOpacityChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "preserveOpacityChanged") {
			C.QCustom3DVolume_ConnectPreserveOpacityChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "preserveOpacityChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "preserveOpacityChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "preserveOpacityChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectPreserveOpacityChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectPreserveOpacityChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "preserveOpacityChanged")
	}
}

func (ptr *QCustom3DVolume) PreserveOpacityChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_PreserveOpacityChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QCustom3DVolume) SetAlphaMultiplier(mult float32) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetAlphaMultiplier(ptr.Pointer(), C.float(mult))
	}
}

func (ptr *QCustom3DVolume) SetDrawSliceFrames(enable bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetDrawSliceFrames(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QCustom3DVolume) SetDrawSlices(enable bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetDrawSlices(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QCustom3DVolume) SetPreserveOpacity(enable bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetPreserveOpacity(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QCustom3DVolume) SetSliceFrameColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetSliceFrameColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QCustom3DVolume) SetSliceFrameGaps(values gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetSliceFrameGaps(ptr.Pointer(), gui.PointerFromQVector3D(values))
	}
}

func (ptr *QCustom3DVolume) SetSliceFrameThicknesses(values gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetSliceFrameThicknesses(ptr.Pointer(), gui.PointerFromQVector3D(values))
	}
}

func (ptr *QCustom3DVolume) SetSliceFrameWidths(values gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetSliceFrameWidths(ptr.Pointer(), gui.PointerFromQVector3D(values))
	}
}

func (ptr *QCustom3DVolume) SetSliceIndexX(value int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetSliceIndexX(ptr.Pointer(), C.int(int32(value)))
	}
}

func (ptr *QCustom3DVolume) SetSliceIndexY(value int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetSliceIndexY(ptr.Pointer(), C.int(int32(value)))
	}
}

func (ptr *QCustom3DVolume) SetSliceIndexZ(value int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetSliceIndexZ(ptr.Pointer(), C.int(int32(value)))
	}
}

func (ptr *QCustom3DVolume) SetSliceIndices(x int, y int, z int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetSliceIndices(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)), C.int(int32(z)))
	}
}

func (ptr *QCustom3DVolume) SetSubTextureData2(axis core.Qt__Axis, index int, image gui.QImage_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetSubTextureData2(ptr.Pointer(), C.longlong(axis), C.int(int32(index)), gui.PointerFromQImage(image))
	}
}

func (ptr *QCustom3DVolume) SetSubTextureData(axis core.Qt__Axis, index int, data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.QCustom3DVolume_SetSubTextureData(ptr.Pointer(), C.longlong(axis), C.int(int32(index)), dataC)
	}
}

func (ptr *QCustom3DVolume) SetTextureDepth(value int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetTextureDepth(ptr.Pointer(), C.int(int32(value)))
	}
}

func (ptr *QCustom3DVolume) SetTextureDimensions(width int, height int, depth int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetTextureDimensions(ptr.Pointer(), C.int(int32(width)), C.int(int32(height)), C.int(int32(depth)))
	}
}

func (ptr *QCustom3DVolume) SetTextureFormat(format gui.QImage__Format) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetTextureFormat(ptr.Pointer(), C.longlong(format))
	}
}

func (ptr *QCustom3DVolume) SetTextureHeight(value int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetTextureHeight(ptr.Pointer(), C.int(int32(value)))
	}
}

func (ptr *QCustom3DVolume) SetTextureWidth(value int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetTextureWidth(ptr.Pointer(), C.int(int32(value)))
	}
}

func (ptr *QCustom3DVolume) SetUseHighDefShader(enable bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SetUseHighDefShader(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQCustom3DVolume_SliceFrameColorChanged
func callbackQCustom3DVolume_SliceFrameColorChanged(ptr unsafe.Pointer, color unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "sliceFrameColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QCustom3DVolume) ConnectSliceFrameColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameColorChanged") {
			C.QCustom3DVolume_ConnectSliceFrameColorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameColorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameColorChanged", func(color *gui.QColor) {
				signal.(func(*gui.QColor))(color)
				f(color)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameColorChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectSliceFrameColorChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectSliceFrameColorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameColorChanged")
	}
}

func (ptr *QCustom3DVolume) SliceFrameColorChanged(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SliceFrameColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

//export callbackQCustom3DVolume_SliceFrameGapsChanged
func callbackQCustom3DVolume_SliceFrameGapsChanged(ptr unsafe.Pointer, values unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "sliceFrameGapsChanged"); signal != nil {
		signal.(func(*gui.QVector3D))(gui.NewQVector3DFromPointer(values))
	}

}

func (ptr *QCustom3DVolume) ConnectSliceFrameGapsChanged(f func(values *gui.QVector3D)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameGapsChanged") {
			C.QCustom3DVolume_ConnectSliceFrameGapsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameGapsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameGapsChanged", func(values *gui.QVector3D) {
				signal.(func(*gui.QVector3D))(values)
				f(values)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameGapsChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectSliceFrameGapsChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectSliceFrameGapsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameGapsChanged")
	}
}

func (ptr *QCustom3DVolume) SliceFrameGapsChanged(values gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SliceFrameGapsChanged(ptr.Pointer(), gui.PointerFromQVector3D(values))
	}
}

//export callbackQCustom3DVolume_SliceFrameThicknessesChanged
func callbackQCustom3DVolume_SliceFrameThicknessesChanged(ptr unsafe.Pointer, values unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "sliceFrameThicknessesChanged"); signal != nil {
		signal.(func(*gui.QVector3D))(gui.NewQVector3DFromPointer(values))
	}

}

func (ptr *QCustom3DVolume) ConnectSliceFrameThicknessesChanged(f func(values *gui.QVector3D)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameThicknessesChanged") {
			C.QCustom3DVolume_ConnectSliceFrameThicknessesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameThicknessesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameThicknessesChanged", func(values *gui.QVector3D) {
				signal.(func(*gui.QVector3D))(values)
				f(values)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameThicknessesChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectSliceFrameThicknessesChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectSliceFrameThicknessesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameThicknessesChanged")
	}
}

func (ptr *QCustom3DVolume) SliceFrameThicknessesChanged(values gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SliceFrameThicknessesChanged(ptr.Pointer(), gui.PointerFromQVector3D(values))
	}
}

//export callbackQCustom3DVolume_SliceFrameWidthsChanged
func callbackQCustom3DVolume_SliceFrameWidthsChanged(ptr unsafe.Pointer, values unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "sliceFrameWidthsChanged"); signal != nil {
		signal.(func(*gui.QVector3D))(gui.NewQVector3DFromPointer(values))
	}

}

func (ptr *QCustom3DVolume) ConnectSliceFrameWidthsChanged(f func(values *gui.QVector3D)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameWidthsChanged") {
			C.QCustom3DVolume_ConnectSliceFrameWidthsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameWidthsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameWidthsChanged", func(values *gui.QVector3D) {
				signal.(func(*gui.QVector3D))(values)
				f(values)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameWidthsChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectSliceFrameWidthsChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectSliceFrameWidthsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "sliceFrameWidthsChanged")
	}
}

func (ptr *QCustom3DVolume) SliceFrameWidthsChanged(values gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SliceFrameWidthsChanged(ptr.Pointer(), gui.PointerFromQVector3D(values))
	}
}

//export callbackQCustom3DVolume_SliceIndexXChanged
func callbackQCustom3DVolume_SliceIndexXChanged(ptr unsafe.Pointer, value C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "sliceIndexXChanged"); signal != nil {
		signal.(func(int))(int(int32(value)))
	}

}

func (ptr *QCustom3DVolume) ConnectSliceIndexXChanged(f func(value int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexXChanged") {
			C.QCustom3DVolume_ConnectSliceIndexXChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexXChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexXChanged", func(value int) {
				signal.(func(int))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexXChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectSliceIndexXChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectSliceIndexXChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexXChanged")
	}
}

func (ptr *QCustom3DVolume) SliceIndexXChanged(value int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SliceIndexXChanged(ptr.Pointer(), C.int(int32(value)))
	}
}

//export callbackQCustom3DVolume_SliceIndexYChanged
func callbackQCustom3DVolume_SliceIndexYChanged(ptr unsafe.Pointer, value C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "sliceIndexYChanged"); signal != nil {
		signal.(func(int))(int(int32(value)))
	}

}

func (ptr *QCustom3DVolume) ConnectSliceIndexYChanged(f func(value int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexYChanged") {
			C.QCustom3DVolume_ConnectSliceIndexYChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexYChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexYChanged", func(value int) {
				signal.(func(int))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexYChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectSliceIndexYChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectSliceIndexYChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexYChanged")
	}
}

func (ptr *QCustom3DVolume) SliceIndexYChanged(value int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SliceIndexYChanged(ptr.Pointer(), C.int(int32(value)))
	}
}

//export callbackQCustom3DVolume_SliceIndexZChanged
func callbackQCustom3DVolume_SliceIndexZChanged(ptr unsafe.Pointer, value C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "sliceIndexZChanged"); signal != nil {
		signal.(func(int))(int(int32(value)))
	}

}

func (ptr *QCustom3DVolume) ConnectSliceIndexZChanged(f func(value int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexZChanged") {
			C.QCustom3DVolume_ConnectSliceIndexZChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexZChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexZChanged", func(value int) {
				signal.(func(int))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexZChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectSliceIndexZChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectSliceIndexZChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "sliceIndexZChanged")
	}
}

func (ptr *QCustom3DVolume) SliceIndexZChanged(value int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_SliceIndexZChanged(ptr.Pointer(), C.int(int32(value)))
	}
}

//export callbackQCustom3DVolume_TextureDepthChanged
func callbackQCustom3DVolume_TextureDepthChanged(ptr unsafe.Pointer, value C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "textureDepthChanged"); signal != nil {
		signal.(func(int))(int(int32(value)))
	}

}

func (ptr *QCustom3DVolume) ConnectTextureDepthChanged(f func(value int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "textureDepthChanged") {
			C.QCustom3DVolume_ConnectTextureDepthChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "textureDepthChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureDepthChanged", func(value int) {
				signal.(func(int))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureDepthChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectTextureDepthChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectTextureDepthChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "textureDepthChanged")
	}
}

func (ptr *QCustom3DVolume) TextureDepthChanged(value int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_TextureDepthChanged(ptr.Pointer(), C.int(int32(value)))
	}
}

//export callbackQCustom3DVolume_TextureFormatChanged
func callbackQCustom3DVolume_TextureFormatChanged(ptr unsafe.Pointer, format C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "textureFormatChanged"); signal != nil {
		signal.(func(gui.QImage__Format))(gui.QImage__Format(format))
	}

}

func (ptr *QCustom3DVolume) ConnectTextureFormatChanged(f func(format gui.QImage__Format)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "textureFormatChanged") {
			C.QCustom3DVolume_ConnectTextureFormatChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "textureFormatChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureFormatChanged", func(format gui.QImage__Format) {
				signal.(func(gui.QImage__Format))(format)
				f(format)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureFormatChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectTextureFormatChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectTextureFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "textureFormatChanged")
	}
}

func (ptr *QCustom3DVolume) TextureFormatChanged(format gui.QImage__Format) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_TextureFormatChanged(ptr.Pointer(), C.longlong(format))
	}
}

//export callbackQCustom3DVolume_TextureHeightChanged
func callbackQCustom3DVolume_TextureHeightChanged(ptr unsafe.Pointer, value C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "textureHeightChanged"); signal != nil {
		signal.(func(int))(int(int32(value)))
	}

}

func (ptr *QCustom3DVolume) ConnectTextureHeightChanged(f func(value int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "textureHeightChanged") {
			C.QCustom3DVolume_ConnectTextureHeightChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "textureHeightChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureHeightChanged", func(value int) {
				signal.(func(int))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureHeightChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectTextureHeightChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectTextureHeightChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "textureHeightChanged")
	}
}

func (ptr *QCustom3DVolume) TextureHeightChanged(value int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_TextureHeightChanged(ptr.Pointer(), C.int(int32(value)))
	}
}

//export callbackQCustom3DVolume_TextureWidthChanged
func callbackQCustom3DVolume_TextureWidthChanged(ptr unsafe.Pointer, value C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "textureWidthChanged"); signal != nil {
		signal.(func(int))(int(int32(value)))
	}

}

func (ptr *QCustom3DVolume) ConnectTextureWidthChanged(f func(value int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "textureWidthChanged") {
			C.QCustom3DVolume_ConnectTextureWidthChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "textureWidthChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureWidthChanged", func(value int) {
				signal.(func(int))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureWidthChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectTextureWidthChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectTextureWidthChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "textureWidthChanged")
	}
}

func (ptr *QCustom3DVolume) TextureWidthChanged(value int) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_TextureWidthChanged(ptr.Pointer(), C.int(int32(value)))
	}
}

//export callbackQCustom3DVolume_UseHighDefShaderChanged
func callbackQCustom3DVolume_UseHighDefShaderChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "useHighDefShaderChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QCustom3DVolume) ConnectUseHighDefShaderChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "useHighDefShaderChanged") {
			C.QCustom3DVolume_ConnectUseHighDefShaderChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "useHighDefShaderChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "useHighDefShaderChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "useHighDefShaderChanged", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectUseHighDefShaderChanged() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DisconnectUseHighDefShaderChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "useHighDefShaderChanged")
	}
}

func (ptr *QCustom3DVolume) UseHighDefShaderChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_UseHighDefShaderChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQCustom3DVolume_DestroyQCustom3DVolume
func callbackQCustom3DVolume_DestroyQCustom3DVolume(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QCustom3DVolume"); signal != nil {
		signal.(func())()
	} else {
		NewQCustom3DVolumeFromPointer(ptr).DestroyQCustom3DVolumeDefault()
	}
}

func (ptr *QCustom3DVolume) ConnectDestroyQCustom3DVolume(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QCustom3DVolume"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QCustom3DVolume", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QCustom3DVolume", f)
		}
	}
}

func (ptr *QCustom3DVolume) DisconnectDestroyQCustom3DVolume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QCustom3DVolume")
	}
}

func (ptr *QCustom3DVolume) DestroyQCustom3DVolume() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DestroyQCustom3DVolume(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCustom3DVolume) DestroyQCustom3DVolumeDefault() {
	if ptr.Pointer() != nil {
		C.QCustom3DVolume_DestroyQCustom3DVolumeDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCustom3DVolume) SliceFrameColor() *gui.QColor {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQColorFromPointer(C.QCustom3DVolume_SliceFrameColor(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QCustom3DVolume) TextureFormat() gui.QImage__Format {
	if ptr.Pointer() != nil {
		return gui.QImage__Format(C.QCustom3DVolume_TextureFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCustom3DVolume) SliceFrameGaps() *gui.QVector3D {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQVector3DFromPointer(C.QCustom3DVolume_SliceFrameGaps(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QVector3D).DestroyQVector3D)
		return tmpValue
	}
	return nil
}

func (ptr *QCustom3DVolume) SliceFrameThicknesses() *gui.QVector3D {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQVector3DFromPointer(C.QCustom3DVolume_SliceFrameThicknesses(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QVector3D).DestroyQVector3D)
		return tmpValue
	}
	return nil
}

func (ptr *QCustom3DVolume) SliceFrameWidths() *gui.QVector3D {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQVector3DFromPointer(C.QCustom3DVolume_SliceFrameWidths(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QVector3D).DestroyQVector3D)
		return tmpValue
	}
	return nil
}

func (ptr *QCustom3DVolume) DrawSliceFrames() bool {
	if ptr.Pointer() != nil {
		return C.QCustom3DVolume_DrawSliceFrames(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCustom3DVolume) DrawSlices() bool {
	if ptr.Pointer() != nil {
		return C.QCustom3DVolume_DrawSlices(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCustom3DVolume) PreserveOpacity() bool {
	if ptr.Pointer() != nil {
		return C.QCustom3DVolume_PreserveOpacity(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCustom3DVolume) UseHighDefShader() bool {
	if ptr.Pointer() != nil {
		return C.QCustom3DVolume_UseHighDefShader(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCustom3DVolume) AlphaMultiplier() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QCustom3DVolume_AlphaMultiplier(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCustom3DVolume) SliceIndexX() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCustom3DVolume_SliceIndexX(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCustom3DVolume) SliceIndexY() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCustom3DVolume_SliceIndexY(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCustom3DVolume) SliceIndexZ() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCustom3DVolume_SliceIndexZ(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCustom3DVolume) TextureDepth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCustom3DVolume_TextureDepth(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCustom3DVolume) TextureHeight() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCustom3DVolume_TextureHeight(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCustom3DVolume) TextureWidth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCustom3DVolume_TextureWidth(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCustom3DVolume) __QCustom3DVolume_colorTable_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QCustom3DVolume___QCustom3DVolume_colorTable_newList2(ptr.Pointer()))
}

func (ptr *QCustom3DVolume) __setColorTable_colors_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QCustom3DVolume___setColorTable_colors_newList(ptr.Pointer()))
}

func (ptr *QCustom3DVolume) __colorTable_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QCustom3DVolume___colorTable_newList(ptr.Pointer()))
}

type QHeightMapSurfaceDataProxy struct {
	ptr unsafe.Pointer
}

type QHeightMapSurfaceDataProxy_ITF interface {
	QHeightMapSurfaceDataProxy_PTR() *QHeightMapSurfaceDataProxy
}

func (ptr *QHeightMapSurfaceDataProxy) QHeightMapSurfaceDataProxy_PTR() *QHeightMapSurfaceDataProxy {
	return ptr
}

func (ptr *QHeightMapSurfaceDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QHeightMapSurfaceDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQHeightMapSurfaceDataProxy(ptr QHeightMapSurfaceDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHeightMapSurfaceDataProxy_PTR().Pointer()
	}
	return nil
}

func NewQHeightMapSurfaceDataProxyFromPointer(ptr unsafe.Pointer) *QHeightMapSurfaceDataProxy {
	var n = new(QHeightMapSurfaceDataProxy)
	n.SetPointer(ptr)
	return n
}
func NewQHeightMapSurfaceDataProxy(parent core.QObject_ITF) *QHeightMapSurfaceDataProxy {
	return NewQHeightMapSurfaceDataProxyFromPointer(C.QHeightMapSurfaceDataProxy_NewQHeightMapSurfaceDataProxy(core.PointerFromQObject(parent)))
}

func NewQHeightMapSurfaceDataProxy2(image gui.QImage_ITF, parent core.QObject_ITF) *QHeightMapSurfaceDataProxy {
	return NewQHeightMapSurfaceDataProxyFromPointer(C.QHeightMapSurfaceDataProxy_NewQHeightMapSurfaceDataProxy2(gui.PointerFromQImage(image), core.PointerFromQObject(parent)))
}

func NewQHeightMapSurfaceDataProxy3(filename string, parent core.QObject_ITF) *QHeightMapSurfaceDataProxy {
	var filenameC *C.char
	if filename != "" {
		filenameC = C.CString(filename)
		defer C.free(unsafe.Pointer(filenameC))
	}
	return NewQHeightMapSurfaceDataProxyFromPointer(C.QHeightMapSurfaceDataProxy_NewQHeightMapSurfaceDataProxy3(filenameC, core.PointerFromQObject(parent)))
}

//export callbackQHeightMapSurfaceDataProxy_HeightMapChanged
func callbackQHeightMapSurfaceDataProxy_HeightMapChanged(ptr unsafe.Pointer, image unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "heightMapChanged"); signal != nil {
		signal.(func(*gui.QImage))(gui.NewQImageFromPointer(image))
	}

}

func (ptr *QHeightMapSurfaceDataProxy) ConnectHeightMapChanged(f func(image *gui.QImage)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "heightMapChanged") {
			C.QHeightMapSurfaceDataProxy_ConnectHeightMapChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "heightMapChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "heightMapChanged", func(image *gui.QImage) {
				signal.(func(*gui.QImage))(image)
				f(image)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "heightMapChanged", f)
		}
	}
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectHeightMapChanged() {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_DisconnectHeightMapChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "heightMapChanged")
	}
}

func (ptr *QHeightMapSurfaceDataProxy) HeightMapChanged(image gui.QImage_ITF) {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_HeightMapChanged(ptr.Pointer(), gui.PointerFromQImage(image))
	}
}

//export callbackQHeightMapSurfaceDataProxy_HeightMapFileChanged
func callbackQHeightMapSurfaceDataProxy_HeightMapFileChanged(ptr unsafe.Pointer, filename C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "heightMapFileChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(filename))
	}

}

func (ptr *QHeightMapSurfaceDataProxy) ConnectHeightMapFileChanged(f func(filename string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "heightMapFileChanged") {
			C.QHeightMapSurfaceDataProxy_ConnectHeightMapFileChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "heightMapFileChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "heightMapFileChanged", func(filename string) {
				signal.(func(string))(filename)
				f(filename)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "heightMapFileChanged", f)
		}
	}
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectHeightMapFileChanged() {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_DisconnectHeightMapFileChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "heightMapFileChanged")
	}
}

func (ptr *QHeightMapSurfaceDataProxy) HeightMapFileChanged(filename string) {
	if ptr.Pointer() != nil {
		var filenameC *C.char
		if filename != "" {
			filenameC = C.CString(filename)
			defer C.free(unsafe.Pointer(filenameC))
		}
		C.QHeightMapSurfaceDataProxy_HeightMapFileChanged(ptr.Pointer(), filenameC)
	}
}

//export callbackQHeightMapSurfaceDataProxy_MaxXValueChanged
func callbackQHeightMapSurfaceDataProxy_MaxXValueChanged(ptr unsafe.Pointer, value C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "maxXValueChanged"); signal != nil {
		signal.(func(float32))(float32(value))
	}

}

func (ptr *QHeightMapSurfaceDataProxy) ConnectMaxXValueChanged(f func(value float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "maxXValueChanged") {
			C.QHeightMapSurfaceDataProxy_ConnectMaxXValueChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "maxXValueChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "maxXValueChanged", func(value float32) {
				signal.(func(float32))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "maxXValueChanged", f)
		}
	}
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectMaxXValueChanged() {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_DisconnectMaxXValueChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "maxXValueChanged")
	}
}

func (ptr *QHeightMapSurfaceDataProxy) MaxXValueChanged(value float32) {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_MaxXValueChanged(ptr.Pointer(), C.float(value))
	}
}

//export callbackQHeightMapSurfaceDataProxy_MaxZValueChanged
func callbackQHeightMapSurfaceDataProxy_MaxZValueChanged(ptr unsafe.Pointer, value C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "maxZValueChanged"); signal != nil {
		signal.(func(float32))(float32(value))
	}

}

func (ptr *QHeightMapSurfaceDataProxy) ConnectMaxZValueChanged(f func(value float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "maxZValueChanged") {
			C.QHeightMapSurfaceDataProxy_ConnectMaxZValueChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "maxZValueChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "maxZValueChanged", func(value float32) {
				signal.(func(float32))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "maxZValueChanged", f)
		}
	}
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectMaxZValueChanged() {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_DisconnectMaxZValueChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "maxZValueChanged")
	}
}

func (ptr *QHeightMapSurfaceDataProxy) MaxZValueChanged(value float32) {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_MaxZValueChanged(ptr.Pointer(), C.float(value))
	}
}

//export callbackQHeightMapSurfaceDataProxy_MinXValueChanged
func callbackQHeightMapSurfaceDataProxy_MinXValueChanged(ptr unsafe.Pointer, value C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "minXValueChanged"); signal != nil {
		signal.(func(float32))(float32(value))
	}

}

func (ptr *QHeightMapSurfaceDataProxy) ConnectMinXValueChanged(f func(value float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "minXValueChanged") {
			C.QHeightMapSurfaceDataProxy_ConnectMinXValueChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "minXValueChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "minXValueChanged", func(value float32) {
				signal.(func(float32))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "minXValueChanged", f)
		}
	}
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectMinXValueChanged() {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_DisconnectMinXValueChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "minXValueChanged")
	}
}

func (ptr *QHeightMapSurfaceDataProxy) MinXValueChanged(value float32) {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_MinXValueChanged(ptr.Pointer(), C.float(value))
	}
}

//export callbackQHeightMapSurfaceDataProxy_MinZValueChanged
func callbackQHeightMapSurfaceDataProxy_MinZValueChanged(ptr unsafe.Pointer, value C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "minZValueChanged"); signal != nil {
		signal.(func(float32))(float32(value))
	}

}

func (ptr *QHeightMapSurfaceDataProxy) ConnectMinZValueChanged(f func(value float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "minZValueChanged") {
			C.QHeightMapSurfaceDataProxy_ConnectMinZValueChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "minZValueChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "minZValueChanged", func(value float32) {
				signal.(func(float32))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "minZValueChanged", f)
		}
	}
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectMinZValueChanged() {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_DisconnectMinZValueChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "minZValueChanged")
	}
}

func (ptr *QHeightMapSurfaceDataProxy) MinZValueChanged(value float32) {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_MinZValueChanged(ptr.Pointer(), C.float(value))
	}
}

func (ptr *QHeightMapSurfaceDataProxy) SetHeightMap(image gui.QImage_ITF) {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_SetHeightMap(ptr.Pointer(), gui.PointerFromQImage(image))
	}
}

func (ptr *QHeightMapSurfaceDataProxy) SetHeightMapFile(filename string) {
	if ptr.Pointer() != nil {
		var filenameC *C.char
		if filename != "" {
			filenameC = C.CString(filename)
			defer C.free(unsafe.Pointer(filenameC))
		}
		C.QHeightMapSurfaceDataProxy_SetHeightMapFile(ptr.Pointer(), filenameC)
	}
}

func (ptr *QHeightMapSurfaceDataProxy) SetMaxXValue(max float32) {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_SetMaxXValue(ptr.Pointer(), C.float(max))
	}
}

func (ptr *QHeightMapSurfaceDataProxy) SetMaxZValue(max float32) {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_SetMaxZValue(ptr.Pointer(), C.float(max))
	}
}

func (ptr *QHeightMapSurfaceDataProxy) SetMinXValue(min float32) {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_SetMinXValue(ptr.Pointer(), C.float(min))
	}
}

func (ptr *QHeightMapSurfaceDataProxy) SetMinZValue(min float32) {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_SetMinZValue(ptr.Pointer(), C.float(min))
	}
}

func (ptr *QHeightMapSurfaceDataProxy) SetValueRanges(minX float32, maxX float32, minZ float32, maxZ float32) {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_SetValueRanges(ptr.Pointer(), C.float(minX), C.float(maxX), C.float(minZ), C.float(maxZ))
	}
}

//export callbackQHeightMapSurfaceDataProxy_DestroyQHeightMapSurfaceDataProxy
func callbackQHeightMapSurfaceDataProxy_DestroyQHeightMapSurfaceDataProxy(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QHeightMapSurfaceDataProxy"); signal != nil {
		signal.(func())()
	} else {
		NewQHeightMapSurfaceDataProxyFromPointer(ptr).DestroyQHeightMapSurfaceDataProxyDefault()
	}
}

func (ptr *QHeightMapSurfaceDataProxy) ConnectDestroyQHeightMapSurfaceDataProxy(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QHeightMapSurfaceDataProxy"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QHeightMapSurfaceDataProxy", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QHeightMapSurfaceDataProxy", f)
		}
	}
}

func (ptr *QHeightMapSurfaceDataProxy) DisconnectDestroyQHeightMapSurfaceDataProxy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QHeightMapSurfaceDataProxy")
	}
}

func (ptr *QHeightMapSurfaceDataProxy) DestroyQHeightMapSurfaceDataProxy() {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_DestroyQHeightMapSurfaceDataProxy(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHeightMapSurfaceDataProxy) DestroyQHeightMapSurfaceDataProxyDefault() {
	if ptr.Pointer() != nil {
		C.QHeightMapSurfaceDataProxy_DestroyQHeightMapSurfaceDataProxyDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHeightMapSurfaceDataProxy) HeightMap() *gui.QImage {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQImageFromPointer(C.QHeightMapSurfaceDataProxy_HeightMap(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QHeightMapSurfaceDataProxy) HeightMapFile() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHeightMapSurfaceDataProxy_HeightMapFile(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHeightMapSurfaceDataProxy) MaxXValue() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QHeightMapSurfaceDataProxy_MaxXValue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeightMapSurfaceDataProxy) MaxZValue() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QHeightMapSurfaceDataProxy_MaxZValue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeightMapSurfaceDataProxy) MinXValue() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QHeightMapSurfaceDataProxy_MinXValue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeightMapSurfaceDataProxy) MinZValue() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QHeightMapSurfaceDataProxy_MinZValue(ptr.Pointer()))
	}
	return 0
}

type QItemModelBarDataProxy struct {
	ptr unsafe.Pointer
}

type QItemModelBarDataProxy_ITF interface {
	QItemModelBarDataProxy_PTR() *QItemModelBarDataProxy
}

func (ptr *QItemModelBarDataProxy) QItemModelBarDataProxy_PTR() *QItemModelBarDataProxy {
	return ptr
}

func (ptr *QItemModelBarDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QItemModelBarDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQItemModelBarDataProxy(ptr QItemModelBarDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemModelBarDataProxy_PTR().Pointer()
	}
	return nil
}

func NewQItemModelBarDataProxyFromPointer(ptr unsafe.Pointer) *QItemModelBarDataProxy {
	var n = new(QItemModelBarDataProxy)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QItemModelBarDataProxy__MultiMatchBehavior
//QItemModelBarDataProxy::MultiMatchBehavior
type QItemModelBarDataProxy__MultiMatchBehavior int64

const (
	QItemModelBarDataProxy__MMBFirst      QItemModelBarDataProxy__MultiMatchBehavior = QItemModelBarDataProxy__MultiMatchBehavior(0)
	QItemModelBarDataProxy__MMBLast       QItemModelBarDataProxy__MultiMatchBehavior = QItemModelBarDataProxy__MultiMatchBehavior(1)
	QItemModelBarDataProxy__MMBAverage    QItemModelBarDataProxy__MultiMatchBehavior = QItemModelBarDataProxy__MultiMatchBehavior(2)
	QItemModelBarDataProxy__MMBCumulative QItemModelBarDataProxy__MultiMatchBehavior = QItemModelBarDataProxy__MultiMatchBehavior(3)
)

func NewQItemModelBarDataProxy2(itemModel core.QAbstractItemModel_ITF, parent core.QObject_ITF) *QItemModelBarDataProxy {
	return NewQItemModelBarDataProxyFromPointer(C.QItemModelBarDataProxy_NewQItemModelBarDataProxy2(core.PointerFromQAbstractItemModel(itemModel), core.PointerFromQObject(parent)))
}

func NewQItemModelBarDataProxy4(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, valueRole string, parent core.QObject_ITF) *QItemModelBarDataProxy {
	var rowRoleC *C.char
	if rowRole != "" {
		rowRoleC = C.CString(rowRole)
		defer C.free(unsafe.Pointer(rowRoleC))
	}
	var columnRoleC *C.char
	if columnRole != "" {
		columnRoleC = C.CString(columnRole)
		defer C.free(unsafe.Pointer(columnRoleC))
	}
	var valueRoleC *C.char
	if valueRole != "" {
		valueRoleC = C.CString(valueRole)
		defer C.free(unsafe.Pointer(valueRoleC))
	}
	return NewQItemModelBarDataProxyFromPointer(C.QItemModelBarDataProxy_NewQItemModelBarDataProxy4(core.PointerFromQAbstractItemModel(itemModel), rowRoleC, columnRoleC, valueRoleC, core.PointerFromQObject(parent)))
}

func NewQItemModelBarDataProxy5(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, valueRole string, rotationRole string, parent core.QObject_ITF) *QItemModelBarDataProxy {
	var rowRoleC *C.char
	if rowRole != "" {
		rowRoleC = C.CString(rowRole)
		defer C.free(unsafe.Pointer(rowRoleC))
	}
	var columnRoleC *C.char
	if columnRole != "" {
		columnRoleC = C.CString(columnRole)
		defer C.free(unsafe.Pointer(columnRoleC))
	}
	var valueRoleC *C.char
	if valueRole != "" {
		valueRoleC = C.CString(valueRole)
		defer C.free(unsafe.Pointer(valueRoleC))
	}
	var rotationRoleC *C.char
	if rotationRole != "" {
		rotationRoleC = C.CString(rotationRole)
		defer C.free(unsafe.Pointer(rotationRoleC))
	}
	return NewQItemModelBarDataProxyFromPointer(C.QItemModelBarDataProxy_NewQItemModelBarDataProxy5(core.PointerFromQAbstractItemModel(itemModel), rowRoleC, columnRoleC, valueRoleC, rotationRoleC, core.PointerFromQObject(parent)))
}

func NewQItemModelBarDataProxy7(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, valueRole string, rotationRole string, rowCategories []string, columnCategories []string, parent core.QObject_ITF) *QItemModelBarDataProxy {
	var rowRoleC *C.char
	if rowRole != "" {
		rowRoleC = C.CString(rowRole)
		defer C.free(unsafe.Pointer(rowRoleC))
	}
	var columnRoleC *C.char
	if columnRole != "" {
		columnRoleC = C.CString(columnRole)
		defer C.free(unsafe.Pointer(columnRoleC))
	}
	var valueRoleC *C.char
	if valueRole != "" {
		valueRoleC = C.CString(valueRole)
		defer C.free(unsafe.Pointer(valueRoleC))
	}
	var rotationRoleC *C.char
	if rotationRole != "" {
		rotationRoleC = C.CString(rotationRole)
		defer C.free(unsafe.Pointer(rotationRoleC))
	}
	var rowCategoriesC = C.CString(strings.Join(rowCategories, "|"))
	defer C.free(unsafe.Pointer(rowCategoriesC))
	var columnCategoriesC = C.CString(strings.Join(columnCategories, "|"))
	defer C.free(unsafe.Pointer(columnCategoriesC))
	return NewQItemModelBarDataProxyFromPointer(C.QItemModelBarDataProxy_NewQItemModelBarDataProxy7(core.PointerFromQAbstractItemModel(itemModel), rowRoleC, columnRoleC, valueRoleC, rotationRoleC, rowCategoriesC, columnCategoriesC, core.PointerFromQObject(parent)))
}

func NewQItemModelBarDataProxy6(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, valueRole string, rowCategories []string, columnCategories []string, parent core.QObject_ITF) *QItemModelBarDataProxy {
	var rowRoleC *C.char
	if rowRole != "" {
		rowRoleC = C.CString(rowRole)
		defer C.free(unsafe.Pointer(rowRoleC))
	}
	var columnRoleC *C.char
	if columnRole != "" {
		columnRoleC = C.CString(columnRole)
		defer C.free(unsafe.Pointer(columnRoleC))
	}
	var valueRoleC *C.char
	if valueRole != "" {
		valueRoleC = C.CString(valueRole)
		defer C.free(unsafe.Pointer(valueRoleC))
	}
	var rowCategoriesC = C.CString(strings.Join(rowCategories, "|"))
	defer C.free(unsafe.Pointer(rowCategoriesC))
	var columnCategoriesC = C.CString(strings.Join(columnCategories, "|"))
	defer C.free(unsafe.Pointer(columnCategoriesC))
	return NewQItemModelBarDataProxyFromPointer(C.QItemModelBarDataProxy_NewQItemModelBarDataProxy6(core.PointerFromQAbstractItemModel(itemModel), rowRoleC, columnRoleC, valueRoleC, rowCategoriesC, columnCategoriesC, core.PointerFromQObject(parent)))
}

func NewQItemModelBarDataProxy3(itemModel core.QAbstractItemModel_ITF, valueRole string, parent core.QObject_ITF) *QItemModelBarDataProxy {
	var valueRoleC *C.char
	if valueRole != "" {
		valueRoleC = C.CString(valueRole)
		defer C.free(unsafe.Pointer(valueRoleC))
	}
	return NewQItemModelBarDataProxyFromPointer(C.QItemModelBarDataProxy_NewQItemModelBarDataProxy3(core.PointerFromQAbstractItemModel(itemModel), valueRoleC, core.PointerFromQObject(parent)))
}

func NewQItemModelBarDataProxy(parent core.QObject_ITF) *QItemModelBarDataProxy {
	return NewQItemModelBarDataProxyFromPointer(C.QItemModelBarDataProxy_NewQItemModelBarDataProxy(core.PointerFromQObject(parent)))
}

func (ptr *QItemModelBarDataProxy) ColumnCategoryIndex(category string) int {
	if ptr.Pointer() != nil {
		var categoryC *C.char
		if category != "" {
			categoryC = C.CString(category)
			defer C.free(unsafe.Pointer(categoryC))
		}
		return int(int32(C.QItemModelBarDataProxy_ColumnCategoryIndex(ptr.Pointer(), categoryC)))
	}
	return 0
}

func (ptr *QItemModelBarDataProxy) RowCategoryIndex(category string) int {
	if ptr.Pointer() != nil {
		var categoryC *C.char
		if category != "" {
			categoryC = C.CString(category)
			defer C.free(unsafe.Pointer(categoryC))
		}
		return int(int32(C.QItemModelBarDataProxy_RowCategoryIndex(ptr.Pointer(), categoryC)))
	}
	return 0
}

//export callbackQItemModelBarDataProxy_AutoColumnCategoriesChanged
func callbackQItemModelBarDataProxy_AutoColumnCategoriesChanged(ptr unsafe.Pointer, enable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "autoColumnCategoriesChanged"); signal != nil {
		signal.(func(bool))(int8(enable) != 0)
	}

}

func (ptr *QItemModelBarDataProxy) ConnectAutoColumnCategoriesChanged(f func(enable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "autoColumnCategoriesChanged") {
			C.QItemModelBarDataProxy_ConnectAutoColumnCategoriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "autoColumnCategoriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "autoColumnCategoriesChanged", func(enable bool) {
				signal.(func(bool))(enable)
				f(enable)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "autoColumnCategoriesChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectAutoColumnCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectAutoColumnCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "autoColumnCategoriesChanged")
	}
}

func (ptr *QItemModelBarDataProxy) AutoColumnCategoriesChanged(enable bool) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_AutoColumnCategoriesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQItemModelBarDataProxy_AutoRowCategoriesChanged
func callbackQItemModelBarDataProxy_AutoRowCategoriesChanged(ptr unsafe.Pointer, enable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "autoRowCategoriesChanged"); signal != nil {
		signal.(func(bool))(int8(enable) != 0)
	}

}

func (ptr *QItemModelBarDataProxy) ConnectAutoRowCategoriesChanged(f func(enable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "autoRowCategoriesChanged") {
			C.QItemModelBarDataProxy_ConnectAutoRowCategoriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "autoRowCategoriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "autoRowCategoriesChanged", func(enable bool) {
				signal.(func(bool))(enable)
				f(enable)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "autoRowCategoriesChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectAutoRowCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectAutoRowCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "autoRowCategoriesChanged")
	}
}

func (ptr *QItemModelBarDataProxy) AutoRowCategoriesChanged(enable bool) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_AutoRowCategoriesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQItemModelBarDataProxy_ColumnCategoriesChanged
func callbackQItemModelBarDataProxy_ColumnCategoriesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "columnCategoriesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QItemModelBarDataProxy) ConnectColumnCategoriesChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "columnCategoriesChanged") {
			C.QItemModelBarDataProxy_ConnectColumnCategoriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "columnCategoriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnCategoriesChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnCategoriesChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectColumnCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectColumnCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "columnCategoriesChanged")
	}
}

func (ptr *QItemModelBarDataProxy) ColumnCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_ColumnCategoriesChanged(ptr.Pointer())
	}
}

//export callbackQItemModelBarDataProxy_ColumnRoleChanged
func callbackQItemModelBarDataProxy_ColumnRoleChanged(ptr unsafe.Pointer, role C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "columnRoleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(role))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectColumnRoleChanged(f func(role string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "columnRoleChanged") {
			C.QItemModelBarDataProxy_ConnectColumnRoleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "columnRoleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnRoleChanged", func(role string) {
				signal.(func(string))(role)
				f(role)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnRoleChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectColumnRoleChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectColumnRoleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "columnRoleChanged")
	}
}

func (ptr *QItemModelBarDataProxy) ColumnRoleChanged(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelBarDataProxy_ColumnRoleChanged(ptr.Pointer(), roleC)
	}
}

//export callbackQItemModelBarDataProxy_ColumnRolePatternChanged
func callbackQItemModelBarDataProxy_ColumnRolePatternChanged(ptr unsafe.Pointer, pattern unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "columnRolePatternChanged"); signal != nil {
		signal.(func(*core.QRegExp))(core.NewQRegExpFromPointer(pattern))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectColumnRolePatternChanged(f func(pattern *core.QRegExp)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "columnRolePatternChanged") {
			C.QItemModelBarDataProxy_ConnectColumnRolePatternChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "columnRolePatternChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnRolePatternChanged", func(pattern *core.QRegExp) {
				signal.(func(*core.QRegExp))(pattern)
				f(pattern)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnRolePatternChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectColumnRolePatternChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectColumnRolePatternChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "columnRolePatternChanged")
	}
}

func (ptr *QItemModelBarDataProxy) ColumnRolePatternChanged(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_ColumnRolePatternChanged(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

//export callbackQItemModelBarDataProxy_ColumnRoleReplaceChanged
func callbackQItemModelBarDataProxy_ColumnRoleReplaceChanged(ptr unsafe.Pointer, replace C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "columnRoleReplaceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(replace))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectColumnRoleReplaceChanged(f func(replace string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "columnRoleReplaceChanged") {
			C.QItemModelBarDataProxy_ConnectColumnRoleReplaceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "columnRoleReplaceChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnRoleReplaceChanged", func(replace string) {
				signal.(func(string))(replace)
				f(replace)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnRoleReplaceChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectColumnRoleReplaceChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectColumnRoleReplaceChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "columnRoleReplaceChanged")
	}
}

func (ptr *QItemModelBarDataProxy) ColumnRoleReplaceChanged(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelBarDataProxy_ColumnRoleReplaceChanged(ptr.Pointer(), replaceC)
	}
}

//export callbackQItemModelBarDataProxy_ItemModelChanged
func callbackQItemModelBarDataProxy_ItemModelChanged(ptr unsafe.Pointer, itemModel unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemModelChanged"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(itemModel))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectItemModelChanged(f func(itemModel *core.QAbstractItemModel)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged") {
			C.QItemModelBarDataProxy_ConnectItemModelChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged", func(itemModel *core.QAbstractItemModel) {
				signal.(func(*core.QAbstractItemModel))(itemModel)
				f(itemModel)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectItemModelChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectItemModelChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged")
	}
}

func (ptr *QItemModelBarDataProxy) ItemModelChanged(itemModel core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_ItemModelChanged(ptr.Pointer(), core.PointerFromQAbstractItemModel(itemModel))
	}
}

//export callbackQItemModelBarDataProxy_MultiMatchBehaviorChanged
func callbackQItemModelBarDataProxy_MultiMatchBehaviorChanged(ptr unsafe.Pointer, behavior C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "multiMatchBehaviorChanged"); signal != nil {
		signal.(func(QItemModelBarDataProxy__MultiMatchBehavior))(QItemModelBarDataProxy__MultiMatchBehavior(behavior))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectMultiMatchBehaviorChanged(f func(behavior QItemModelBarDataProxy__MultiMatchBehavior)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "multiMatchBehaviorChanged") {
			C.QItemModelBarDataProxy_ConnectMultiMatchBehaviorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "multiMatchBehaviorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiMatchBehaviorChanged", func(behavior QItemModelBarDataProxy__MultiMatchBehavior) {
				signal.(func(QItemModelBarDataProxy__MultiMatchBehavior))(behavior)
				f(behavior)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiMatchBehaviorChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectMultiMatchBehaviorChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectMultiMatchBehaviorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "multiMatchBehaviorChanged")
	}
}

func (ptr *QItemModelBarDataProxy) MultiMatchBehaviorChanged(behavior QItemModelBarDataProxy__MultiMatchBehavior) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_MultiMatchBehaviorChanged(ptr.Pointer(), C.longlong(behavior))
	}
}

func (ptr *QItemModelBarDataProxy) Remap(rowRole string, columnRole string, valueRole string, rotationRole string, rowCategories []string, columnCategories []string) {
	if ptr.Pointer() != nil {
		var rowRoleC *C.char
		if rowRole != "" {
			rowRoleC = C.CString(rowRole)
			defer C.free(unsafe.Pointer(rowRoleC))
		}
		var columnRoleC *C.char
		if columnRole != "" {
			columnRoleC = C.CString(columnRole)
			defer C.free(unsafe.Pointer(columnRoleC))
		}
		var valueRoleC *C.char
		if valueRole != "" {
			valueRoleC = C.CString(valueRole)
			defer C.free(unsafe.Pointer(valueRoleC))
		}
		var rotationRoleC *C.char
		if rotationRole != "" {
			rotationRoleC = C.CString(rotationRole)
			defer C.free(unsafe.Pointer(rotationRoleC))
		}
		var rowCategoriesC = C.CString(strings.Join(rowCategories, "|"))
		defer C.free(unsafe.Pointer(rowCategoriesC))
		var columnCategoriesC = C.CString(strings.Join(columnCategories, "|"))
		defer C.free(unsafe.Pointer(columnCategoriesC))
		C.QItemModelBarDataProxy_Remap(ptr.Pointer(), rowRoleC, columnRoleC, valueRoleC, rotationRoleC, rowCategoriesC, columnCategoriesC)
	}
}

//export callbackQItemModelBarDataProxy_RotationRoleChanged
func callbackQItemModelBarDataProxy_RotationRoleChanged(ptr unsafe.Pointer, role C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rotationRoleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(role))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectRotationRoleChanged(f func(role string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleChanged") {
			C.QItemModelBarDataProxy_ConnectRotationRoleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleChanged", func(role string) {
				signal.(func(string))(role)
				f(role)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectRotationRoleChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectRotationRoleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleChanged")
	}
}

func (ptr *QItemModelBarDataProxy) RotationRoleChanged(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelBarDataProxy_RotationRoleChanged(ptr.Pointer(), roleC)
	}
}

//export callbackQItemModelBarDataProxy_RotationRolePatternChanged
func callbackQItemModelBarDataProxy_RotationRolePatternChanged(ptr unsafe.Pointer, pattern unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rotationRolePatternChanged"); signal != nil {
		signal.(func(*core.QRegExp))(core.NewQRegExpFromPointer(pattern))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectRotationRolePatternChanged(f func(pattern *core.QRegExp)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rotationRolePatternChanged") {
			C.QItemModelBarDataProxy_ConnectRotationRolePatternChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rotationRolePatternChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRolePatternChanged", func(pattern *core.QRegExp) {
				signal.(func(*core.QRegExp))(pattern)
				f(pattern)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRolePatternChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectRotationRolePatternChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectRotationRolePatternChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRolePatternChanged")
	}
}

func (ptr *QItemModelBarDataProxy) RotationRolePatternChanged(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_RotationRolePatternChanged(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

//export callbackQItemModelBarDataProxy_RotationRoleReplaceChanged
func callbackQItemModelBarDataProxy_RotationRoleReplaceChanged(ptr unsafe.Pointer, replace C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rotationRoleReplaceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(replace))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectRotationRoleReplaceChanged(f func(replace string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleReplaceChanged") {
			C.QItemModelBarDataProxy_ConnectRotationRoleReplaceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleReplaceChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleReplaceChanged", func(replace string) {
				signal.(func(string))(replace)
				f(replace)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleReplaceChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectRotationRoleReplaceChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectRotationRoleReplaceChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleReplaceChanged")
	}
}

func (ptr *QItemModelBarDataProxy) RotationRoleReplaceChanged(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelBarDataProxy_RotationRoleReplaceChanged(ptr.Pointer(), replaceC)
	}
}

//export callbackQItemModelBarDataProxy_RowCategoriesChanged
func callbackQItemModelBarDataProxy_RowCategoriesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowCategoriesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QItemModelBarDataProxy) ConnectRowCategoriesChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowCategoriesChanged") {
			C.QItemModelBarDataProxy_ConnectRowCategoriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowCategoriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowCategoriesChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowCategoriesChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectRowCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectRowCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowCategoriesChanged")
	}
}

func (ptr *QItemModelBarDataProxy) RowCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_RowCategoriesChanged(ptr.Pointer())
	}
}

//export callbackQItemModelBarDataProxy_RowRoleChanged
func callbackQItemModelBarDataProxy_RowRoleChanged(ptr unsafe.Pointer, role C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowRoleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(role))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectRowRoleChanged(f func(role string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowRoleChanged") {
			C.QItemModelBarDataProxy_ConnectRowRoleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowRoleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowRoleChanged", func(role string) {
				signal.(func(string))(role)
				f(role)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowRoleChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectRowRoleChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectRowRoleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowRoleChanged")
	}
}

func (ptr *QItemModelBarDataProxy) RowRoleChanged(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelBarDataProxy_RowRoleChanged(ptr.Pointer(), roleC)
	}
}

//export callbackQItemModelBarDataProxy_RowRolePatternChanged
func callbackQItemModelBarDataProxy_RowRolePatternChanged(ptr unsafe.Pointer, pattern unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowRolePatternChanged"); signal != nil {
		signal.(func(*core.QRegExp))(core.NewQRegExpFromPointer(pattern))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectRowRolePatternChanged(f func(pattern *core.QRegExp)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowRolePatternChanged") {
			C.QItemModelBarDataProxy_ConnectRowRolePatternChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowRolePatternChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowRolePatternChanged", func(pattern *core.QRegExp) {
				signal.(func(*core.QRegExp))(pattern)
				f(pattern)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowRolePatternChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectRowRolePatternChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectRowRolePatternChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowRolePatternChanged")
	}
}

func (ptr *QItemModelBarDataProxy) RowRolePatternChanged(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_RowRolePatternChanged(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

//export callbackQItemModelBarDataProxy_RowRoleReplaceChanged
func callbackQItemModelBarDataProxy_RowRoleReplaceChanged(ptr unsafe.Pointer, replace C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowRoleReplaceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(replace))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectRowRoleReplaceChanged(f func(replace string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowRoleReplaceChanged") {
			C.QItemModelBarDataProxy_ConnectRowRoleReplaceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowRoleReplaceChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowRoleReplaceChanged", func(replace string) {
				signal.(func(string))(replace)
				f(replace)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowRoleReplaceChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectRowRoleReplaceChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectRowRoleReplaceChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowRoleReplaceChanged")
	}
}

func (ptr *QItemModelBarDataProxy) RowRoleReplaceChanged(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelBarDataProxy_RowRoleReplaceChanged(ptr.Pointer(), replaceC)
	}
}

func (ptr *QItemModelBarDataProxy) SetAutoColumnCategories(enable bool) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_SetAutoColumnCategories(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QItemModelBarDataProxy) SetAutoRowCategories(enable bool) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_SetAutoRowCategories(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QItemModelBarDataProxy) SetColumnCategories(categories []string) {
	if ptr.Pointer() != nil {
		var categoriesC = C.CString(strings.Join(categories, "|"))
		defer C.free(unsafe.Pointer(categoriesC))
		C.QItemModelBarDataProxy_SetColumnCategories(ptr.Pointer(), categoriesC)
	}
}

func (ptr *QItemModelBarDataProxy) SetColumnRole(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelBarDataProxy_SetColumnRole(ptr.Pointer(), roleC)
	}
}

func (ptr *QItemModelBarDataProxy) SetColumnRolePattern(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_SetColumnRolePattern(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

func (ptr *QItemModelBarDataProxy) SetColumnRoleReplace(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelBarDataProxy_SetColumnRoleReplace(ptr.Pointer(), replaceC)
	}
}

func (ptr *QItemModelBarDataProxy) SetItemModel(itemModel core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_SetItemModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(itemModel))
	}
}

func (ptr *QItemModelBarDataProxy) SetMultiMatchBehavior(behavior QItemModelBarDataProxy__MultiMatchBehavior) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_SetMultiMatchBehavior(ptr.Pointer(), C.longlong(behavior))
	}
}

func (ptr *QItemModelBarDataProxy) SetRotationRole(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelBarDataProxy_SetRotationRole(ptr.Pointer(), roleC)
	}
}

func (ptr *QItemModelBarDataProxy) SetRotationRolePattern(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_SetRotationRolePattern(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

func (ptr *QItemModelBarDataProxy) SetRotationRoleReplace(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelBarDataProxy_SetRotationRoleReplace(ptr.Pointer(), replaceC)
	}
}

func (ptr *QItemModelBarDataProxy) SetRowCategories(categories []string) {
	if ptr.Pointer() != nil {
		var categoriesC = C.CString(strings.Join(categories, "|"))
		defer C.free(unsafe.Pointer(categoriesC))
		C.QItemModelBarDataProxy_SetRowCategories(ptr.Pointer(), categoriesC)
	}
}

func (ptr *QItemModelBarDataProxy) SetRowRole(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelBarDataProxy_SetRowRole(ptr.Pointer(), roleC)
	}
}

func (ptr *QItemModelBarDataProxy) SetRowRolePattern(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_SetRowRolePattern(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

func (ptr *QItemModelBarDataProxy) SetRowRoleReplace(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelBarDataProxy_SetRowRoleReplace(ptr.Pointer(), replaceC)
	}
}

func (ptr *QItemModelBarDataProxy) SetUseModelCategories(enable bool) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_SetUseModelCategories(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QItemModelBarDataProxy) SetValueRole(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelBarDataProxy_SetValueRole(ptr.Pointer(), roleC)
	}
}

func (ptr *QItemModelBarDataProxy) SetValueRolePattern(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_SetValueRolePattern(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

func (ptr *QItemModelBarDataProxy) SetValueRoleReplace(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelBarDataProxy_SetValueRoleReplace(ptr.Pointer(), replaceC)
	}
}

//export callbackQItemModelBarDataProxy_UseModelCategoriesChanged
func callbackQItemModelBarDataProxy_UseModelCategoriesChanged(ptr unsafe.Pointer, enable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "useModelCategoriesChanged"); signal != nil {
		signal.(func(bool))(int8(enable) != 0)
	}

}

func (ptr *QItemModelBarDataProxy) ConnectUseModelCategoriesChanged(f func(enable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "useModelCategoriesChanged") {
			C.QItemModelBarDataProxy_ConnectUseModelCategoriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "useModelCategoriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "useModelCategoriesChanged", func(enable bool) {
				signal.(func(bool))(enable)
				f(enable)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "useModelCategoriesChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectUseModelCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectUseModelCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "useModelCategoriesChanged")
	}
}

func (ptr *QItemModelBarDataProxy) UseModelCategoriesChanged(enable bool) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_UseModelCategoriesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQItemModelBarDataProxy_ValueRoleChanged
func callbackQItemModelBarDataProxy_ValueRoleChanged(ptr unsafe.Pointer, role C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "valueRoleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(role))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectValueRoleChanged(f func(role string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "valueRoleChanged") {
			C.QItemModelBarDataProxy_ConnectValueRoleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "valueRoleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "valueRoleChanged", func(role string) {
				signal.(func(string))(role)
				f(role)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "valueRoleChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectValueRoleChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectValueRoleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "valueRoleChanged")
	}
}

func (ptr *QItemModelBarDataProxy) ValueRoleChanged(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelBarDataProxy_ValueRoleChanged(ptr.Pointer(), roleC)
	}
}

//export callbackQItemModelBarDataProxy_ValueRolePatternChanged
func callbackQItemModelBarDataProxy_ValueRolePatternChanged(ptr unsafe.Pointer, pattern unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "valueRolePatternChanged"); signal != nil {
		signal.(func(*core.QRegExp))(core.NewQRegExpFromPointer(pattern))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectValueRolePatternChanged(f func(pattern *core.QRegExp)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "valueRolePatternChanged") {
			C.QItemModelBarDataProxy_ConnectValueRolePatternChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "valueRolePatternChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "valueRolePatternChanged", func(pattern *core.QRegExp) {
				signal.(func(*core.QRegExp))(pattern)
				f(pattern)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "valueRolePatternChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectValueRolePatternChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectValueRolePatternChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "valueRolePatternChanged")
	}
}

func (ptr *QItemModelBarDataProxy) ValueRolePatternChanged(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_ValueRolePatternChanged(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

//export callbackQItemModelBarDataProxy_ValueRoleReplaceChanged
func callbackQItemModelBarDataProxy_ValueRoleReplaceChanged(ptr unsafe.Pointer, replace C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "valueRoleReplaceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(replace))
	}

}

func (ptr *QItemModelBarDataProxy) ConnectValueRoleReplaceChanged(f func(replace string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "valueRoleReplaceChanged") {
			C.QItemModelBarDataProxy_ConnectValueRoleReplaceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "valueRoleReplaceChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "valueRoleReplaceChanged", func(replace string) {
				signal.(func(string))(replace)
				f(replace)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "valueRoleReplaceChanged", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectValueRoleReplaceChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DisconnectValueRoleReplaceChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "valueRoleReplaceChanged")
	}
}

func (ptr *QItemModelBarDataProxy) ValueRoleReplaceChanged(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelBarDataProxy_ValueRoleReplaceChanged(ptr.Pointer(), replaceC)
	}
}

//export callbackQItemModelBarDataProxy_DestroyQItemModelBarDataProxy
func callbackQItemModelBarDataProxy_DestroyQItemModelBarDataProxy(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QItemModelBarDataProxy"); signal != nil {
		signal.(func())()
	} else {
		NewQItemModelBarDataProxyFromPointer(ptr).DestroyQItemModelBarDataProxyDefault()
	}
}

func (ptr *QItemModelBarDataProxy) ConnectDestroyQItemModelBarDataProxy(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QItemModelBarDataProxy"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QItemModelBarDataProxy", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QItemModelBarDataProxy", f)
		}
	}
}

func (ptr *QItemModelBarDataProxy) DisconnectDestroyQItemModelBarDataProxy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QItemModelBarDataProxy")
	}
}

func (ptr *QItemModelBarDataProxy) DestroyQItemModelBarDataProxy() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DestroyQItemModelBarDataProxy(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QItemModelBarDataProxy) DestroyQItemModelBarDataProxyDefault() {
	if ptr.Pointer() != nil {
		C.QItemModelBarDataProxy_DestroyQItemModelBarDataProxyDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QItemModelBarDataProxy) MultiMatchBehavior() QItemModelBarDataProxy__MultiMatchBehavior {
	if ptr.Pointer() != nil {
		return QItemModelBarDataProxy__MultiMatchBehavior(C.QItemModelBarDataProxy_MultiMatchBehavior(ptr.Pointer()))
	}
	return 0
}

func (ptr *QItemModelBarDataProxy) ItemModel() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQAbstractItemModelFromPointer(C.QItemModelBarDataProxy_ItemModel(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelBarDataProxy) ColumnRolePattern() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QItemModelBarDataProxy_ColumnRolePattern(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelBarDataProxy) RotationRolePattern() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QItemModelBarDataProxy_RotationRolePattern(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelBarDataProxy) RowRolePattern() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QItemModelBarDataProxy_RowRolePattern(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelBarDataProxy) ValueRolePattern() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QItemModelBarDataProxy_ValueRolePattern(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelBarDataProxy) ColumnRole() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelBarDataProxy_ColumnRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelBarDataProxy) ColumnRoleReplace() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelBarDataProxy_ColumnRoleReplace(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelBarDataProxy) RotationRole() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelBarDataProxy_RotationRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelBarDataProxy) RotationRoleReplace() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelBarDataProxy_RotationRoleReplace(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelBarDataProxy) RowRole() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelBarDataProxy_RowRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelBarDataProxy) RowRoleReplace() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelBarDataProxy_RowRoleReplace(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelBarDataProxy) ValueRole() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelBarDataProxy_ValueRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelBarDataProxy) ValueRoleReplace() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelBarDataProxy_ValueRoleReplace(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelBarDataProxy) ColumnCategories() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QItemModelBarDataProxy_ColumnCategories(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QItemModelBarDataProxy) RowCategories() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QItemModelBarDataProxy_RowCategories(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QItemModelBarDataProxy) AutoColumnCategories() bool {
	if ptr.Pointer() != nil {
		return C.QItemModelBarDataProxy_AutoColumnCategories(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QItemModelBarDataProxy) AutoRowCategories() bool {
	if ptr.Pointer() != nil {
		return C.QItemModelBarDataProxy_AutoRowCategories(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QItemModelBarDataProxy) UseModelCategories() bool {
	if ptr.Pointer() != nil {
		return C.QItemModelBarDataProxy_UseModelCategories(ptr.Pointer()) != 0
	}
	return false
}

type QItemModelScatterDataProxy struct {
	ptr unsafe.Pointer
}

type QItemModelScatterDataProxy_ITF interface {
	QItemModelScatterDataProxy_PTR() *QItemModelScatterDataProxy
}

func (ptr *QItemModelScatterDataProxy) QItemModelScatterDataProxy_PTR() *QItemModelScatterDataProxy {
	return ptr
}

func (ptr *QItemModelScatterDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QItemModelScatterDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQItemModelScatterDataProxy(ptr QItemModelScatterDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemModelScatterDataProxy_PTR().Pointer()
	}
	return nil
}

func NewQItemModelScatterDataProxyFromPointer(ptr unsafe.Pointer) *QItemModelScatterDataProxy {
	var n = new(QItemModelScatterDataProxy)
	n.SetPointer(ptr)
	return n
}
func NewQItemModelScatterDataProxy2(itemModel core.QAbstractItemModel_ITF, parent core.QObject_ITF) *QItemModelScatterDataProxy {
	return NewQItemModelScatterDataProxyFromPointer(C.QItemModelScatterDataProxy_NewQItemModelScatterDataProxy2(core.PointerFromQAbstractItemModel(itemModel), core.PointerFromQObject(parent)))
}

func NewQItemModelScatterDataProxy3(itemModel core.QAbstractItemModel_ITF, xPosRole string, yPosRole string, zPosRole string, parent core.QObject_ITF) *QItemModelScatterDataProxy {
	var xPosRoleC *C.char
	if xPosRole != "" {
		xPosRoleC = C.CString(xPosRole)
		defer C.free(unsafe.Pointer(xPosRoleC))
	}
	var yPosRoleC *C.char
	if yPosRole != "" {
		yPosRoleC = C.CString(yPosRole)
		defer C.free(unsafe.Pointer(yPosRoleC))
	}
	var zPosRoleC *C.char
	if zPosRole != "" {
		zPosRoleC = C.CString(zPosRole)
		defer C.free(unsafe.Pointer(zPosRoleC))
	}
	return NewQItemModelScatterDataProxyFromPointer(C.QItemModelScatterDataProxy_NewQItemModelScatterDataProxy3(core.PointerFromQAbstractItemModel(itemModel), xPosRoleC, yPosRoleC, zPosRoleC, core.PointerFromQObject(parent)))
}

func NewQItemModelScatterDataProxy4(itemModel core.QAbstractItemModel_ITF, xPosRole string, yPosRole string, zPosRole string, rotationRole string, parent core.QObject_ITF) *QItemModelScatterDataProxy {
	var xPosRoleC *C.char
	if xPosRole != "" {
		xPosRoleC = C.CString(xPosRole)
		defer C.free(unsafe.Pointer(xPosRoleC))
	}
	var yPosRoleC *C.char
	if yPosRole != "" {
		yPosRoleC = C.CString(yPosRole)
		defer C.free(unsafe.Pointer(yPosRoleC))
	}
	var zPosRoleC *C.char
	if zPosRole != "" {
		zPosRoleC = C.CString(zPosRole)
		defer C.free(unsafe.Pointer(zPosRoleC))
	}
	var rotationRoleC *C.char
	if rotationRole != "" {
		rotationRoleC = C.CString(rotationRole)
		defer C.free(unsafe.Pointer(rotationRoleC))
	}
	return NewQItemModelScatterDataProxyFromPointer(C.QItemModelScatterDataProxy_NewQItemModelScatterDataProxy4(core.PointerFromQAbstractItemModel(itemModel), xPosRoleC, yPosRoleC, zPosRoleC, rotationRoleC, core.PointerFromQObject(parent)))
}

func NewQItemModelScatterDataProxy(parent core.QObject_ITF) *QItemModelScatterDataProxy {
	return NewQItemModelScatterDataProxyFromPointer(C.QItemModelScatterDataProxy_NewQItemModelScatterDataProxy(core.PointerFromQObject(parent)))
}

//export callbackQItemModelScatterDataProxy_ItemModelChanged
func callbackQItemModelScatterDataProxy_ItemModelChanged(ptr unsafe.Pointer, itemModel unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemModelChanged"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(itemModel))
	}

}

func (ptr *QItemModelScatterDataProxy) ConnectItemModelChanged(f func(itemModel *core.QAbstractItemModel)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged") {
			C.QItemModelScatterDataProxy_ConnectItemModelChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged", func(itemModel *core.QAbstractItemModel) {
				signal.(func(*core.QAbstractItemModel))(itemModel)
				f(itemModel)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectItemModelChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DisconnectItemModelChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged")
	}
}

func (ptr *QItemModelScatterDataProxy) ItemModelChanged(itemModel core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_ItemModelChanged(ptr.Pointer(), core.PointerFromQAbstractItemModel(itemModel))
	}
}

func (ptr *QItemModelScatterDataProxy) Remap(xPosRole string, yPosRole string, zPosRole string, rotationRole string) {
	if ptr.Pointer() != nil {
		var xPosRoleC *C.char
		if xPosRole != "" {
			xPosRoleC = C.CString(xPosRole)
			defer C.free(unsafe.Pointer(xPosRoleC))
		}
		var yPosRoleC *C.char
		if yPosRole != "" {
			yPosRoleC = C.CString(yPosRole)
			defer C.free(unsafe.Pointer(yPosRoleC))
		}
		var zPosRoleC *C.char
		if zPosRole != "" {
			zPosRoleC = C.CString(zPosRole)
			defer C.free(unsafe.Pointer(zPosRoleC))
		}
		var rotationRoleC *C.char
		if rotationRole != "" {
			rotationRoleC = C.CString(rotationRole)
			defer C.free(unsafe.Pointer(rotationRoleC))
		}
		C.QItemModelScatterDataProxy_Remap(ptr.Pointer(), xPosRoleC, yPosRoleC, zPosRoleC, rotationRoleC)
	}
}

//export callbackQItemModelScatterDataProxy_RotationRoleChanged
func callbackQItemModelScatterDataProxy_RotationRoleChanged(ptr unsafe.Pointer, role C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rotationRoleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(role))
	}

}

func (ptr *QItemModelScatterDataProxy) ConnectRotationRoleChanged(f func(role string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleChanged") {
			C.QItemModelScatterDataProxy_ConnectRotationRoleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleChanged", func(role string) {
				signal.(func(string))(role)
				f(role)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleChanged", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectRotationRoleChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DisconnectRotationRoleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleChanged")
	}
}

func (ptr *QItemModelScatterDataProxy) RotationRoleChanged(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelScatterDataProxy_RotationRoleChanged(ptr.Pointer(), roleC)
	}
}

//export callbackQItemModelScatterDataProxy_RotationRolePatternChanged
func callbackQItemModelScatterDataProxy_RotationRolePatternChanged(ptr unsafe.Pointer, pattern unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rotationRolePatternChanged"); signal != nil {
		signal.(func(*core.QRegExp))(core.NewQRegExpFromPointer(pattern))
	}

}

func (ptr *QItemModelScatterDataProxy) ConnectRotationRolePatternChanged(f func(pattern *core.QRegExp)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rotationRolePatternChanged") {
			C.QItemModelScatterDataProxy_ConnectRotationRolePatternChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rotationRolePatternChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRolePatternChanged", func(pattern *core.QRegExp) {
				signal.(func(*core.QRegExp))(pattern)
				f(pattern)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRolePatternChanged", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectRotationRolePatternChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DisconnectRotationRolePatternChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRolePatternChanged")
	}
}

func (ptr *QItemModelScatterDataProxy) RotationRolePatternChanged(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_RotationRolePatternChanged(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

//export callbackQItemModelScatterDataProxy_RotationRoleReplaceChanged
func callbackQItemModelScatterDataProxy_RotationRoleReplaceChanged(ptr unsafe.Pointer, replace C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rotationRoleReplaceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(replace))
	}

}

func (ptr *QItemModelScatterDataProxy) ConnectRotationRoleReplaceChanged(f func(replace string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleReplaceChanged") {
			C.QItemModelScatterDataProxy_ConnectRotationRoleReplaceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleReplaceChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleReplaceChanged", func(replace string) {
				signal.(func(string))(replace)
				f(replace)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleReplaceChanged", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectRotationRoleReplaceChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DisconnectRotationRoleReplaceChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rotationRoleReplaceChanged")
	}
}

func (ptr *QItemModelScatterDataProxy) RotationRoleReplaceChanged(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelScatterDataProxy_RotationRoleReplaceChanged(ptr.Pointer(), replaceC)
	}
}

func (ptr *QItemModelScatterDataProxy) SetItemModel(itemModel core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_SetItemModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(itemModel))
	}
}

func (ptr *QItemModelScatterDataProxy) SetRotationRole(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelScatterDataProxy_SetRotationRole(ptr.Pointer(), roleC)
	}
}

func (ptr *QItemModelScatterDataProxy) SetRotationRolePattern(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_SetRotationRolePattern(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

func (ptr *QItemModelScatterDataProxy) SetRotationRoleReplace(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelScatterDataProxy_SetRotationRoleReplace(ptr.Pointer(), replaceC)
	}
}

func (ptr *QItemModelScatterDataProxy) SetXPosRole(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelScatterDataProxy_SetXPosRole(ptr.Pointer(), roleC)
	}
}

func (ptr *QItemModelScatterDataProxy) SetXPosRolePattern(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_SetXPosRolePattern(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

func (ptr *QItemModelScatterDataProxy) SetXPosRoleReplace(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelScatterDataProxy_SetXPosRoleReplace(ptr.Pointer(), replaceC)
	}
}

func (ptr *QItemModelScatterDataProxy) SetYPosRole(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelScatterDataProxy_SetYPosRole(ptr.Pointer(), roleC)
	}
}

func (ptr *QItemModelScatterDataProxy) SetYPosRolePattern(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_SetYPosRolePattern(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

func (ptr *QItemModelScatterDataProxy) SetYPosRoleReplace(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelScatterDataProxy_SetYPosRoleReplace(ptr.Pointer(), replaceC)
	}
}

func (ptr *QItemModelScatterDataProxy) SetZPosRole(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelScatterDataProxy_SetZPosRole(ptr.Pointer(), roleC)
	}
}

func (ptr *QItemModelScatterDataProxy) SetZPosRolePattern(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_SetZPosRolePattern(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

func (ptr *QItemModelScatterDataProxy) SetZPosRoleReplace(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelScatterDataProxy_SetZPosRoleReplace(ptr.Pointer(), replaceC)
	}
}

//export callbackQItemModelScatterDataProxy_XPosRoleChanged
func callbackQItemModelScatterDataProxy_XPosRoleChanged(ptr unsafe.Pointer, role C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "xPosRoleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(role))
	}

}

func (ptr *QItemModelScatterDataProxy) ConnectXPosRoleChanged(f func(role string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleChanged") {
			C.QItemModelScatterDataProxy_ConnectXPosRoleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleChanged", func(role string) {
				signal.(func(string))(role)
				f(role)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleChanged", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectXPosRoleChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DisconnectXPosRoleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleChanged")
	}
}

func (ptr *QItemModelScatterDataProxy) XPosRoleChanged(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelScatterDataProxy_XPosRoleChanged(ptr.Pointer(), roleC)
	}
}

//export callbackQItemModelScatterDataProxy_XPosRolePatternChanged
func callbackQItemModelScatterDataProxy_XPosRolePatternChanged(ptr unsafe.Pointer, pattern unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "xPosRolePatternChanged"); signal != nil {
		signal.(func(*core.QRegExp))(core.NewQRegExpFromPointer(pattern))
	}

}

func (ptr *QItemModelScatterDataProxy) ConnectXPosRolePatternChanged(f func(pattern *core.QRegExp)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "xPosRolePatternChanged") {
			C.QItemModelScatterDataProxy_ConnectXPosRolePatternChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "xPosRolePatternChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRolePatternChanged", func(pattern *core.QRegExp) {
				signal.(func(*core.QRegExp))(pattern)
				f(pattern)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRolePatternChanged", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectXPosRolePatternChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DisconnectXPosRolePatternChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRolePatternChanged")
	}
}

func (ptr *QItemModelScatterDataProxy) XPosRolePatternChanged(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_XPosRolePatternChanged(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

//export callbackQItemModelScatterDataProxy_XPosRoleReplaceChanged
func callbackQItemModelScatterDataProxy_XPosRoleReplaceChanged(ptr unsafe.Pointer, replace C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "xPosRoleReplaceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(replace))
	}

}

func (ptr *QItemModelScatterDataProxy) ConnectXPosRoleReplaceChanged(f func(replace string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleReplaceChanged") {
			C.QItemModelScatterDataProxy_ConnectXPosRoleReplaceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleReplaceChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleReplaceChanged", func(replace string) {
				signal.(func(string))(replace)
				f(replace)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleReplaceChanged", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectXPosRoleReplaceChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DisconnectXPosRoleReplaceChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleReplaceChanged")
	}
}

func (ptr *QItemModelScatterDataProxy) XPosRoleReplaceChanged(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelScatterDataProxy_XPosRoleReplaceChanged(ptr.Pointer(), replaceC)
	}
}

//export callbackQItemModelScatterDataProxy_YPosRoleChanged
func callbackQItemModelScatterDataProxy_YPosRoleChanged(ptr unsafe.Pointer, role C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "yPosRoleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(role))
	}

}

func (ptr *QItemModelScatterDataProxy) ConnectYPosRoleChanged(f func(role string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleChanged") {
			C.QItemModelScatterDataProxy_ConnectYPosRoleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleChanged", func(role string) {
				signal.(func(string))(role)
				f(role)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleChanged", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectYPosRoleChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DisconnectYPosRoleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleChanged")
	}
}

func (ptr *QItemModelScatterDataProxy) YPosRoleChanged(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelScatterDataProxy_YPosRoleChanged(ptr.Pointer(), roleC)
	}
}

//export callbackQItemModelScatterDataProxy_YPosRolePatternChanged
func callbackQItemModelScatterDataProxy_YPosRolePatternChanged(ptr unsafe.Pointer, pattern unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "yPosRolePatternChanged"); signal != nil {
		signal.(func(*core.QRegExp))(core.NewQRegExpFromPointer(pattern))
	}

}

func (ptr *QItemModelScatterDataProxy) ConnectYPosRolePatternChanged(f func(pattern *core.QRegExp)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "yPosRolePatternChanged") {
			C.QItemModelScatterDataProxy_ConnectYPosRolePatternChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "yPosRolePatternChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRolePatternChanged", func(pattern *core.QRegExp) {
				signal.(func(*core.QRegExp))(pattern)
				f(pattern)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRolePatternChanged", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectYPosRolePatternChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DisconnectYPosRolePatternChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRolePatternChanged")
	}
}

func (ptr *QItemModelScatterDataProxy) YPosRolePatternChanged(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_YPosRolePatternChanged(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

//export callbackQItemModelScatterDataProxy_YPosRoleReplaceChanged
func callbackQItemModelScatterDataProxy_YPosRoleReplaceChanged(ptr unsafe.Pointer, replace C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "yPosRoleReplaceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(replace))
	}

}

func (ptr *QItemModelScatterDataProxy) ConnectYPosRoleReplaceChanged(f func(replace string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleReplaceChanged") {
			C.QItemModelScatterDataProxy_ConnectYPosRoleReplaceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleReplaceChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleReplaceChanged", func(replace string) {
				signal.(func(string))(replace)
				f(replace)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleReplaceChanged", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectYPosRoleReplaceChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DisconnectYPosRoleReplaceChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleReplaceChanged")
	}
}

func (ptr *QItemModelScatterDataProxy) YPosRoleReplaceChanged(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelScatterDataProxy_YPosRoleReplaceChanged(ptr.Pointer(), replaceC)
	}
}

//export callbackQItemModelScatterDataProxy_ZPosRoleChanged
func callbackQItemModelScatterDataProxy_ZPosRoleChanged(ptr unsafe.Pointer, role C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "zPosRoleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(role))
	}

}

func (ptr *QItemModelScatterDataProxy) ConnectZPosRoleChanged(f func(role string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleChanged") {
			C.QItemModelScatterDataProxy_ConnectZPosRoleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleChanged", func(role string) {
				signal.(func(string))(role)
				f(role)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleChanged", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectZPosRoleChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DisconnectZPosRoleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleChanged")
	}
}

func (ptr *QItemModelScatterDataProxy) ZPosRoleChanged(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelScatterDataProxy_ZPosRoleChanged(ptr.Pointer(), roleC)
	}
}

//export callbackQItemModelScatterDataProxy_ZPosRolePatternChanged
func callbackQItemModelScatterDataProxy_ZPosRolePatternChanged(ptr unsafe.Pointer, pattern unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "zPosRolePatternChanged"); signal != nil {
		signal.(func(*core.QRegExp))(core.NewQRegExpFromPointer(pattern))
	}

}

func (ptr *QItemModelScatterDataProxy) ConnectZPosRolePatternChanged(f func(pattern *core.QRegExp)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "zPosRolePatternChanged") {
			C.QItemModelScatterDataProxy_ConnectZPosRolePatternChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "zPosRolePatternChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRolePatternChanged", func(pattern *core.QRegExp) {
				signal.(func(*core.QRegExp))(pattern)
				f(pattern)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRolePatternChanged", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectZPosRolePatternChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DisconnectZPosRolePatternChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRolePatternChanged")
	}
}

func (ptr *QItemModelScatterDataProxy) ZPosRolePatternChanged(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_ZPosRolePatternChanged(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

//export callbackQItemModelScatterDataProxy_ZPosRoleReplaceChanged
func callbackQItemModelScatterDataProxy_ZPosRoleReplaceChanged(ptr unsafe.Pointer, replace C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "zPosRoleReplaceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(replace))
	}

}

func (ptr *QItemModelScatterDataProxy) ConnectZPosRoleReplaceChanged(f func(replace string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleReplaceChanged") {
			C.QItemModelScatterDataProxy_ConnectZPosRoleReplaceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleReplaceChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleReplaceChanged", func(replace string) {
				signal.(func(string))(replace)
				f(replace)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleReplaceChanged", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectZPosRoleReplaceChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DisconnectZPosRoleReplaceChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleReplaceChanged")
	}
}

func (ptr *QItemModelScatterDataProxy) ZPosRoleReplaceChanged(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelScatterDataProxy_ZPosRoleReplaceChanged(ptr.Pointer(), replaceC)
	}
}

//export callbackQItemModelScatterDataProxy_DestroyQItemModelScatterDataProxy
func callbackQItemModelScatterDataProxy_DestroyQItemModelScatterDataProxy(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QItemModelScatterDataProxy"); signal != nil {
		signal.(func())()
	} else {
		NewQItemModelScatterDataProxyFromPointer(ptr).DestroyQItemModelScatterDataProxyDefault()
	}
}

func (ptr *QItemModelScatterDataProxy) ConnectDestroyQItemModelScatterDataProxy(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QItemModelScatterDataProxy"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QItemModelScatterDataProxy", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QItemModelScatterDataProxy", f)
		}
	}
}

func (ptr *QItemModelScatterDataProxy) DisconnectDestroyQItemModelScatterDataProxy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QItemModelScatterDataProxy")
	}
}

func (ptr *QItemModelScatterDataProxy) DestroyQItemModelScatterDataProxy() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DestroyQItemModelScatterDataProxy(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QItemModelScatterDataProxy) DestroyQItemModelScatterDataProxyDefault() {
	if ptr.Pointer() != nil {
		C.QItemModelScatterDataProxy_DestroyQItemModelScatterDataProxyDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QItemModelScatterDataProxy) ItemModel() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQAbstractItemModelFromPointer(C.QItemModelScatterDataProxy_ItemModel(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelScatterDataProxy) RotationRolePattern() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QItemModelScatterDataProxy_RotationRolePattern(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelScatterDataProxy) XPosRolePattern() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QItemModelScatterDataProxy_XPosRolePattern(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelScatterDataProxy) YPosRolePattern() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QItemModelScatterDataProxy_YPosRolePattern(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelScatterDataProxy) ZPosRolePattern() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QItemModelScatterDataProxy_ZPosRolePattern(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelScatterDataProxy) RotationRole() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelScatterDataProxy_RotationRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelScatterDataProxy) RotationRoleReplace() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelScatterDataProxy_RotationRoleReplace(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelScatterDataProxy) XPosRole() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelScatterDataProxy_XPosRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelScatterDataProxy) XPosRoleReplace() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelScatterDataProxy_XPosRoleReplace(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelScatterDataProxy) YPosRole() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelScatterDataProxy_YPosRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelScatterDataProxy) YPosRoleReplace() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelScatterDataProxy_YPosRoleReplace(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelScatterDataProxy) ZPosRole() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelScatterDataProxy_ZPosRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelScatterDataProxy) ZPosRoleReplace() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelScatterDataProxy_ZPosRoleReplace(ptr.Pointer()))
	}
	return ""
}

type QItemModelSurfaceDataProxy struct {
	ptr unsafe.Pointer
}

type QItemModelSurfaceDataProxy_ITF interface {
	QItemModelSurfaceDataProxy_PTR() *QItemModelSurfaceDataProxy
}

func (ptr *QItemModelSurfaceDataProxy) QItemModelSurfaceDataProxy_PTR() *QItemModelSurfaceDataProxy {
	return ptr
}

func (ptr *QItemModelSurfaceDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QItemModelSurfaceDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQItemModelSurfaceDataProxy(ptr QItemModelSurfaceDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemModelSurfaceDataProxy_PTR().Pointer()
	}
	return nil
}

func NewQItemModelSurfaceDataProxyFromPointer(ptr unsafe.Pointer) *QItemModelSurfaceDataProxy {
	var n = new(QItemModelSurfaceDataProxy)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QItemModelSurfaceDataProxy__MultiMatchBehavior
//QItemModelSurfaceDataProxy::MultiMatchBehavior
type QItemModelSurfaceDataProxy__MultiMatchBehavior int64

const (
	QItemModelSurfaceDataProxy__MMBFirst       QItemModelSurfaceDataProxy__MultiMatchBehavior = QItemModelSurfaceDataProxy__MultiMatchBehavior(0)
	QItemModelSurfaceDataProxy__MMBLast        QItemModelSurfaceDataProxy__MultiMatchBehavior = QItemModelSurfaceDataProxy__MultiMatchBehavior(1)
	QItemModelSurfaceDataProxy__MMBAverage     QItemModelSurfaceDataProxy__MultiMatchBehavior = QItemModelSurfaceDataProxy__MultiMatchBehavior(2)
	QItemModelSurfaceDataProxy__MMBCumulativeY QItemModelSurfaceDataProxy__MultiMatchBehavior = QItemModelSurfaceDataProxy__MultiMatchBehavior(3)
)

func NewQItemModelSurfaceDataProxy2(itemModel core.QAbstractItemModel_ITF, parent core.QObject_ITF) *QItemModelSurfaceDataProxy {
	return NewQItemModelSurfaceDataProxyFromPointer(C.QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy2(core.PointerFromQAbstractItemModel(itemModel), core.PointerFromQObject(parent)))
}

func NewQItemModelSurfaceDataProxy5(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, xPosRole string, yPosRole string, zPosRole string, parent core.QObject_ITF) *QItemModelSurfaceDataProxy {
	var rowRoleC *C.char
	if rowRole != "" {
		rowRoleC = C.CString(rowRole)
		defer C.free(unsafe.Pointer(rowRoleC))
	}
	var columnRoleC *C.char
	if columnRole != "" {
		columnRoleC = C.CString(columnRole)
		defer C.free(unsafe.Pointer(columnRoleC))
	}
	var xPosRoleC *C.char
	if xPosRole != "" {
		xPosRoleC = C.CString(xPosRole)
		defer C.free(unsafe.Pointer(xPosRoleC))
	}
	var yPosRoleC *C.char
	if yPosRole != "" {
		yPosRoleC = C.CString(yPosRole)
		defer C.free(unsafe.Pointer(yPosRoleC))
	}
	var zPosRoleC *C.char
	if zPosRole != "" {
		zPosRoleC = C.CString(zPosRole)
		defer C.free(unsafe.Pointer(zPosRoleC))
	}
	return NewQItemModelSurfaceDataProxyFromPointer(C.QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy5(core.PointerFromQAbstractItemModel(itemModel), rowRoleC, columnRoleC, xPosRoleC, yPosRoleC, zPosRoleC, core.PointerFromQObject(parent)))
}

func NewQItemModelSurfaceDataProxy7(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, xPosRole string, yPosRole string, zPosRole string, rowCategories []string, columnCategories []string, parent core.QObject_ITF) *QItemModelSurfaceDataProxy {
	var rowRoleC *C.char
	if rowRole != "" {
		rowRoleC = C.CString(rowRole)
		defer C.free(unsafe.Pointer(rowRoleC))
	}
	var columnRoleC *C.char
	if columnRole != "" {
		columnRoleC = C.CString(columnRole)
		defer C.free(unsafe.Pointer(columnRoleC))
	}
	var xPosRoleC *C.char
	if xPosRole != "" {
		xPosRoleC = C.CString(xPosRole)
		defer C.free(unsafe.Pointer(xPosRoleC))
	}
	var yPosRoleC *C.char
	if yPosRole != "" {
		yPosRoleC = C.CString(yPosRole)
		defer C.free(unsafe.Pointer(yPosRoleC))
	}
	var zPosRoleC *C.char
	if zPosRole != "" {
		zPosRoleC = C.CString(zPosRole)
		defer C.free(unsafe.Pointer(zPosRoleC))
	}
	var rowCategoriesC = C.CString(strings.Join(rowCategories, "|"))
	defer C.free(unsafe.Pointer(rowCategoriesC))
	var columnCategoriesC = C.CString(strings.Join(columnCategories, "|"))
	defer C.free(unsafe.Pointer(columnCategoriesC))
	return NewQItemModelSurfaceDataProxyFromPointer(C.QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy7(core.PointerFromQAbstractItemModel(itemModel), rowRoleC, columnRoleC, xPosRoleC, yPosRoleC, zPosRoleC, rowCategoriesC, columnCategoriesC, core.PointerFromQObject(parent)))
}

func NewQItemModelSurfaceDataProxy4(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, yPosRole string, parent core.QObject_ITF) *QItemModelSurfaceDataProxy {
	var rowRoleC *C.char
	if rowRole != "" {
		rowRoleC = C.CString(rowRole)
		defer C.free(unsafe.Pointer(rowRoleC))
	}
	var columnRoleC *C.char
	if columnRole != "" {
		columnRoleC = C.CString(columnRole)
		defer C.free(unsafe.Pointer(columnRoleC))
	}
	var yPosRoleC *C.char
	if yPosRole != "" {
		yPosRoleC = C.CString(yPosRole)
		defer C.free(unsafe.Pointer(yPosRoleC))
	}
	return NewQItemModelSurfaceDataProxyFromPointer(C.QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy4(core.PointerFromQAbstractItemModel(itemModel), rowRoleC, columnRoleC, yPosRoleC, core.PointerFromQObject(parent)))
}

func NewQItemModelSurfaceDataProxy6(itemModel core.QAbstractItemModel_ITF, rowRole string, columnRole string, yPosRole string, rowCategories []string, columnCategories []string, parent core.QObject_ITF) *QItemModelSurfaceDataProxy {
	var rowRoleC *C.char
	if rowRole != "" {
		rowRoleC = C.CString(rowRole)
		defer C.free(unsafe.Pointer(rowRoleC))
	}
	var columnRoleC *C.char
	if columnRole != "" {
		columnRoleC = C.CString(columnRole)
		defer C.free(unsafe.Pointer(columnRoleC))
	}
	var yPosRoleC *C.char
	if yPosRole != "" {
		yPosRoleC = C.CString(yPosRole)
		defer C.free(unsafe.Pointer(yPosRoleC))
	}
	var rowCategoriesC = C.CString(strings.Join(rowCategories, "|"))
	defer C.free(unsafe.Pointer(rowCategoriesC))
	var columnCategoriesC = C.CString(strings.Join(columnCategories, "|"))
	defer C.free(unsafe.Pointer(columnCategoriesC))
	return NewQItemModelSurfaceDataProxyFromPointer(C.QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy6(core.PointerFromQAbstractItemModel(itemModel), rowRoleC, columnRoleC, yPosRoleC, rowCategoriesC, columnCategoriesC, core.PointerFromQObject(parent)))
}

func NewQItemModelSurfaceDataProxy3(itemModel core.QAbstractItemModel_ITF, yPosRole string, parent core.QObject_ITF) *QItemModelSurfaceDataProxy {
	var yPosRoleC *C.char
	if yPosRole != "" {
		yPosRoleC = C.CString(yPosRole)
		defer C.free(unsafe.Pointer(yPosRoleC))
	}
	return NewQItemModelSurfaceDataProxyFromPointer(C.QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy3(core.PointerFromQAbstractItemModel(itemModel), yPosRoleC, core.PointerFromQObject(parent)))
}

func NewQItemModelSurfaceDataProxy(parent core.QObject_ITF) *QItemModelSurfaceDataProxy {
	return NewQItemModelSurfaceDataProxyFromPointer(C.QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy(core.PointerFromQObject(parent)))
}

func (ptr *QItemModelSurfaceDataProxy) ColumnCategoryIndex(category string) int {
	if ptr.Pointer() != nil {
		var categoryC *C.char
		if category != "" {
			categoryC = C.CString(category)
			defer C.free(unsafe.Pointer(categoryC))
		}
		return int(int32(C.QItemModelSurfaceDataProxy_ColumnCategoryIndex(ptr.Pointer(), categoryC)))
	}
	return 0
}

func (ptr *QItemModelSurfaceDataProxy) RowCategoryIndex(category string) int {
	if ptr.Pointer() != nil {
		var categoryC *C.char
		if category != "" {
			categoryC = C.CString(category)
			defer C.free(unsafe.Pointer(categoryC))
		}
		return int(int32(C.QItemModelSurfaceDataProxy_RowCategoryIndex(ptr.Pointer(), categoryC)))
	}
	return 0
}

//export callbackQItemModelSurfaceDataProxy_AutoColumnCategoriesChanged
func callbackQItemModelSurfaceDataProxy_AutoColumnCategoriesChanged(ptr unsafe.Pointer, enable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "autoColumnCategoriesChanged"); signal != nil {
		signal.(func(bool))(int8(enable) != 0)
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectAutoColumnCategoriesChanged(f func(enable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "autoColumnCategoriesChanged") {
			C.QItemModelSurfaceDataProxy_ConnectAutoColumnCategoriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "autoColumnCategoriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "autoColumnCategoriesChanged", func(enable bool) {
				signal.(func(bool))(enable)
				f(enable)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "autoColumnCategoriesChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectAutoColumnCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectAutoColumnCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "autoColumnCategoriesChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) AutoColumnCategoriesChanged(enable bool) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_AutoColumnCategoriesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQItemModelSurfaceDataProxy_AutoRowCategoriesChanged
func callbackQItemModelSurfaceDataProxy_AutoRowCategoriesChanged(ptr unsafe.Pointer, enable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "autoRowCategoriesChanged"); signal != nil {
		signal.(func(bool))(int8(enable) != 0)
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectAutoRowCategoriesChanged(f func(enable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "autoRowCategoriesChanged") {
			C.QItemModelSurfaceDataProxy_ConnectAutoRowCategoriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "autoRowCategoriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "autoRowCategoriesChanged", func(enable bool) {
				signal.(func(bool))(enable)
				f(enable)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "autoRowCategoriesChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectAutoRowCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectAutoRowCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "autoRowCategoriesChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) AutoRowCategoriesChanged(enable bool) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_AutoRowCategoriesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQItemModelSurfaceDataProxy_ColumnCategoriesChanged
func callbackQItemModelSurfaceDataProxy_ColumnCategoriesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "columnCategoriesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectColumnCategoriesChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "columnCategoriesChanged") {
			C.QItemModelSurfaceDataProxy_ConnectColumnCategoriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "columnCategoriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnCategoriesChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnCategoriesChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectColumnCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectColumnCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "columnCategoriesChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) ColumnCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_ColumnCategoriesChanged(ptr.Pointer())
	}
}

//export callbackQItemModelSurfaceDataProxy_ColumnRoleChanged
func callbackQItemModelSurfaceDataProxy_ColumnRoleChanged(ptr unsafe.Pointer, role C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "columnRoleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(role))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectColumnRoleChanged(f func(role string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "columnRoleChanged") {
			C.QItemModelSurfaceDataProxy_ConnectColumnRoleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "columnRoleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnRoleChanged", func(role string) {
				signal.(func(string))(role)
				f(role)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnRoleChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectColumnRoleChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectColumnRoleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "columnRoleChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) ColumnRoleChanged(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelSurfaceDataProxy_ColumnRoleChanged(ptr.Pointer(), roleC)
	}
}

//export callbackQItemModelSurfaceDataProxy_ColumnRolePatternChanged
func callbackQItemModelSurfaceDataProxy_ColumnRolePatternChanged(ptr unsafe.Pointer, pattern unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "columnRolePatternChanged"); signal != nil {
		signal.(func(*core.QRegExp))(core.NewQRegExpFromPointer(pattern))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectColumnRolePatternChanged(f func(pattern *core.QRegExp)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "columnRolePatternChanged") {
			C.QItemModelSurfaceDataProxy_ConnectColumnRolePatternChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "columnRolePatternChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnRolePatternChanged", func(pattern *core.QRegExp) {
				signal.(func(*core.QRegExp))(pattern)
				f(pattern)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnRolePatternChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectColumnRolePatternChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectColumnRolePatternChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "columnRolePatternChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) ColumnRolePatternChanged(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_ColumnRolePatternChanged(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

//export callbackQItemModelSurfaceDataProxy_ColumnRoleReplaceChanged
func callbackQItemModelSurfaceDataProxy_ColumnRoleReplaceChanged(ptr unsafe.Pointer, replace C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "columnRoleReplaceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(replace))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectColumnRoleReplaceChanged(f func(replace string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "columnRoleReplaceChanged") {
			C.QItemModelSurfaceDataProxy_ConnectColumnRoleReplaceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "columnRoleReplaceChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnRoleReplaceChanged", func(replace string) {
				signal.(func(string))(replace)
				f(replace)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnRoleReplaceChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectColumnRoleReplaceChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectColumnRoleReplaceChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "columnRoleReplaceChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) ColumnRoleReplaceChanged(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelSurfaceDataProxy_ColumnRoleReplaceChanged(ptr.Pointer(), replaceC)
	}
}

//export callbackQItemModelSurfaceDataProxy_ItemModelChanged
func callbackQItemModelSurfaceDataProxy_ItemModelChanged(ptr unsafe.Pointer, itemModel unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemModelChanged"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(itemModel))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectItemModelChanged(f func(itemModel *core.QAbstractItemModel)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged") {
			C.QItemModelSurfaceDataProxy_ConnectItemModelChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged", func(itemModel *core.QAbstractItemModel) {
				signal.(func(*core.QAbstractItemModel))(itemModel)
				f(itemModel)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectItemModelChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectItemModelChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemModelChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) ItemModelChanged(itemModel core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_ItemModelChanged(ptr.Pointer(), core.PointerFromQAbstractItemModel(itemModel))
	}
}

//export callbackQItemModelSurfaceDataProxy_MultiMatchBehaviorChanged
func callbackQItemModelSurfaceDataProxy_MultiMatchBehaviorChanged(ptr unsafe.Pointer, behavior C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "multiMatchBehaviorChanged"); signal != nil {
		signal.(func(QItemModelSurfaceDataProxy__MultiMatchBehavior))(QItemModelSurfaceDataProxy__MultiMatchBehavior(behavior))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectMultiMatchBehaviorChanged(f func(behavior QItemModelSurfaceDataProxy__MultiMatchBehavior)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "multiMatchBehaviorChanged") {
			C.QItemModelSurfaceDataProxy_ConnectMultiMatchBehaviorChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "multiMatchBehaviorChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiMatchBehaviorChanged", func(behavior QItemModelSurfaceDataProxy__MultiMatchBehavior) {
				signal.(func(QItemModelSurfaceDataProxy__MultiMatchBehavior))(behavior)
				f(behavior)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "multiMatchBehaviorChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectMultiMatchBehaviorChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectMultiMatchBehaviorChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "multiMatchBehaviorChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) MultiMatchBehaviorChanged(behavior QItemModelSurfaceDataProxy__MultiMatchBehavior) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_MultiMatchBehaviorChanged(ptr.Pointer(), C.longlong(behavior))
	}
}

func (ptr *QItemModelSurfaceDataProxy) Remap(rowRole string, columnRole string, xPosRole string, yPosRole string, zPosRole string, rowCategories []string, columnCategories []string) {
	if ptr.Pointer() != nil {
		var rowRoleC *C.char
		if rowRole != "" {
			rowRoleC = C.CString(rowRole)
			defer C.free(unsafe.Pointer(rowRoleC))
		}
		var columnRoleC *C.char
		if columnRole != "" {
			columnRoleC = C.CString(columnRole)
			defer C.free(unsafe.Pointer(columnRoleC))
		}
		var xPosRoleC *C.char
		if xPosRole != "" {
			xPosRoleC = C.CString(xPosRole)
			defer C.free(unsafe.Pointer(xPosRoleC))
		}
		var yPosRoleC *C.char
		if yPosRole != "" {
			yPosRoleC = C.CString(yPosRole)
			defer C.free(unsafe.Pointer(yPosRoleC))
		}
		var zPosRoleC *C.char
		if zPosRole != "" {
			zPosRoleC = C.CString(zPosRole)
			defer C.free(unsafe.Pointer(zPosRoleC))
		}
		var rowCategoriesC = C.CString(strings.Join(rowCategories, "|"))
		defer C.free(unsafe.Pointer(rowCategoriesC))
		var columnCategoriesC = C.CString(strings.Join(columnCategories, "|"))
		defer C.free(unsafe.Pointer(columnCategoriesC))
		C.QItemModelSurfaceDataProxy_Remap(ptr.Pointer(), rowRoleC, columnRoleC, xPosRoleC, yPosRoleC, zPosRoleC, rowCategoriesC, columnCategoriesC)
	}
}

//export callbackQItemModelSurfaceDataProxy_RowCategoriesChanged
func callbackQItemModelSurfaceDataProxy_RowCategoriesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowCategoriesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectRowCategoriesChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowCategoriesChanged") {
			C.QItemModelSurfaceDataProxy_ConnectRowCategoriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowCategoriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowCategoriesChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowCategoriesChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectRowCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectRowCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowCategoriesChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) RowCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_RowCategoriesChanged(ptr.Pointer())
	}
}

//export callbackQItemModelSurfaceDataProxy_RowRoleChanged
func callbackQItemModelSurfaceDataProxy_RowRoleChanged(ptr unsafe.Pointer, role C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowRoleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(role))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectRowRoleChanged(f func(role string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowRoleChanged") {
			C.QItemModelSurfaceDataProxy_ConnectRowRoleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowRoleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowRoleChanged", func(role string) {
				signal.(func(string))(role)
				f(role)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowRoleChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectRowRoleChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectRowRoleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowRoleChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) RowRoleChanged(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelSurfaceDataProxy_RowRoleChanged(ptr.Pointer(), roleC)
	}
}

//export callbackQItemModelSurfaceDataProxy_RowRolePatternChanged
func callbackQItemModelSurfaceDataProxy_RowRolePatternChanged(ptr unsafe.Pointer, pattern unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowRolePatternChanged"); signal != nil {
		signal.(func(*core.QRegExp))(core.NewQRegExpFromPointer(pattern))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectRowRolePatternChanged(f func(pattern *core.QRegExp)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowRolePatternChanged") {
			C.QItemModelSurfaceDataProxy_ConnectRowRolePatternChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowRolePatternChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowRolePatternChanged", func(pattern *core.QRegExp) {
				signal.(func(*core.QRegExp))(pattern)
				f(pattern)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowRolePatternChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectRowRolePatternChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectRowRolePatternChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowRolePatternChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) RowRolePatternChanged(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_RowRolePatternChanged(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

//export callbackQItemModelSurfaceDataProxy_RowRoleReplaceChanged
func callbackQItemModelSurfaceDataProxy_RowRoleReplaceChanged(ptr unsafe.Pointer, replace C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowRoleReplaceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(replace))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectRowRoleReplaceChanged(f func(replace string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowRoleReplaceChanged") {
			C.QItemModelSurfaceDataProxy_ConnectRowRoleReplaceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowRoleReplaceChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowRoleReplaceChanged", func(replace string) {
				signal.(func(string))(replace)
				f(replace)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowRoleReplaceChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectRowRoleReplaceChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectRowRoleReplaceChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowRoleReplaceChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) RowRoleReplaceChanged(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelSurfaceDataProxy_RowRoleReplaceChanged(ptr.Pointer(), replaceC)
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetAutoColumnCategories(enable bool) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_SetAutoColumnCategories(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetAutoRowCategories(enable bool) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_SetAutoRowCategories(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetColumnCategories(categories []string) {
	if ptr.Pointer() != nil {
		var categoriesC = C.CString(strings.Join(categories, "|"))
		defer C.free(unsafe.Pointer(categoriesC))
		C.QItemModelSurfaceDataProxy_SetColumnCategories(ptr.Pointer(), categoriesC)
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetColumnRole(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelSurfaceDataProxy_SetColumnRole(ptr.Pointer(), roleC)
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetColumnRolePattern(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_SetColumnRolePattern(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetColumnRoleReplace(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelSurfaceDataProxy_SetColumnRoleReplace(ptr.Pointer(), replaceC)
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetItemModel(itemModel core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_SetItemModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(itemModel))
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetMultiMatchBehavior(behavior QItemModelSurfaceDataProxy__MultiMatchBehavior) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_SetMultiMatchBehavior(ptr.Pointer(), C.longlong(behavior))
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetRowCategories(categories []string) {
	if ptr.Pointer() != nil {
		var categoriesC = C.CString(strings.Join(categories, "|"))
		defer C.free(unsafe.Pointer(categoriesC))
		C.QItemModelSurfaceDataProxy_SetRowCategories(ptr.Pointer(), categoriesC)
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetRowRole(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelSurfaceDataProxy_SetRowRole(ptr.Pointer(), roleC)
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetRowRolePattern(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_SetRowRolePattern(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetRowRoleReplace(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelSurfaceDataProxy_SetRowRoleReplace(ptr.Pointer(), replaceC)
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetUseModelCategories(enable bool) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_SetUseModelCategories(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetXPosRole(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelSurfaceDataProxy_SetXPosRole(ptr.Pointer(), roleC)
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetXPosRolePattern(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_SetXPosRolePattern(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetXPosRoleReplace(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelSurfaceDataProxy_SetXPosRoleReplace(ptr.Pointer(), replaceC)
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetYPosRole(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelSurfaceDataProxy_SetYPosRole(ptr.Pointer(), roleC)
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetYPosRolePattern(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_SetYPosRolePattern(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetYPosRoleReplace(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelSurfaceDataProxy_SetYPosRoleReplace(ptr.Pointer(), replaceC)
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetZPosRole(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelSurfaceDataProxy_SetZPosRole(ptr.Pointer(), roleC)
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetZPosRolePattern(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_SetZPosRolePattern(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

func (ptr *QItemModelSurfaceDataProxy) SetZPosRoleReplace(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelSurfaceDataProxy_SetZPosRoleReplace(ptr.Pointer(), replaceC)
	}
}

//export callbackQItemModelSurfaceDataProxy_UseModelCategoriesChanged
func callbackQItemModelSurfaceDataProxy_UseModelCategoriesChanged(ptr unsafe.Pointer, enable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "useModelCategoriesChanged"); signal != nil {
		signal.(func(bool))(int8(enable) != 0)
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectUseModelCategoriesChanged(f func(enable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "useModelCategoriesChanged") {
			C.QItemModelSurfaceDataProxy_ConnectUseModelCategoriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "useModelCategoriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "useModelCategoriesChanged", func(enable bool) {
				signal.(func(bool))(enable)
				f(enable)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "useModelCategoriesChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectUseModelCategoriesChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectUseModelCategoriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "useModelCategoriesChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) UseModelCategoriesChanged(enable bool) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_UseModelCategoriesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQItemModelSurfaceDataProxy_XPosRoleChanged
func callbackQItemModelSurfaceDataProxy_XPosRoleChanged(ptr unsafe.Pointer, role C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "xPosRoleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(role))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectXPosRoleChanged(f func(role string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleChanged") {
			C.QItemModelSurfaceDataProxy_ConnectXPosRoleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleChanged", func(role string) {
				signal.(func(string))(role)
				f(role)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectXPosRoleChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectXPosRoleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) XPosRoleChanged(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelSurfaceDataProxy_XPosRoleChanged(ptr.Pointer(), roleC)
	}
}

//export callbackQItemModelSurfaceDataProxy_XPosRolePatternChanged
func callbackQItemModelSurfaceDataProxy_XPosRolePatternChanged(ptr unsafe.Pointer, pattern unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "xPosRolePatternChanged"); signal != nil {
		signal.(func(*core.QRegExp))(core.NewQRegExpFromPointer(pattern))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectXPosRolePatternChanged(f func(pattern *core.QRegExp)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "xPosRolePatternChanged") {
			C.QItemModelSurfaceDataProxy_ConnectXPosRolePatternChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "xPosRolePatternChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRolePatternChanged", func(pattern *core.QRegExp) {
				signal.(func(*core.QRegExp))(pattern)
				f(pattern)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRolePatternChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectXPosRolePatternChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectXPosRolePatternChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRolePatternChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) XPosRolePatternChanged(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_XPosRolePatternChanged(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

//export callbackQItemModelSurfaceDataProxy_XPosRoleReplaceChanged
func callbackQItemModelSurfaceDataProxy_XPosRoleReplaceChanged(ptr unsafe.Pointer, replace C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "xPosRoleReplaceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(replace))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectXPosRoleReplaceChanged(f func(replace string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleReplaceChanged") {
			C.QItemModelSurfaceDataProxy_ConnectXPosRoleReplaceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleReplaceChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleReplaceChanged", func(replace string) {
				signal.(func(string))(replace)
				f(replace)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleReplaceChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectXPosRoleReplaceChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectXPosRoleReplaceChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "xPosRoleReplaceChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) XPosRoleReplaceChanged(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelSurfaceDataProxy_XPosRoleReplaceChanged(ptr.Pointer(), replaceC)
	}
}

//export callbackQItemModelSurfaceDataProxy_YPosRoleChanged
func callbackQItemModelSurfaceDataProxy_YPosRoleChanged(ptr unsafe.Pointer, role C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "yPosRoleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(role))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectYPosRoleChanged(f func(role string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleChanged") {
			C.QItemModelSurfaceDataProxy_ConnectYPosRoleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleChanged", func(role string) {
				signal.(func(string))(role)
				f(role)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectYPosRoleChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectYPosRoleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) YPosRoleChanged(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelSurfaceDataProxy_YPosRoleChanged(ptr.Pointer(), roleC)
	}
}

//export callbackQItemModelSurfaceDataProxy_YPosRolePatternChanged
func callbackQItemModelSurfaceDataProxy_YPosRolePatternChanged(ptr unsafe.Pointer, pattern unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "yPosRolePatternChanged"); signal != nil {
		signal.(func(*core.QRegExp))(core.NewQRegExpFromPointer(pattern))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectYPosRolePatternChanged(f func(pattern *core.QRegExp)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "yPosRolePatternChanged") {
			C.QItemModelSurfaceDataProxy_ConnectYPosRolePatternChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "yPosRolePatternChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRolePatternChanged", func(pattern *core.QRegExp) {
				signal.(func(*core.QRegExp))(pattern)
				f(pattern)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRolePatternChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectYPosRolePatternChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectYPosRolePatternChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRolePatternChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) YPosRolePatternChanged(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_YPosRolePatternChanged(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

//export callbackQItemModelSurfaceDataProxy_YPosRoleReplaceChanged
func callbackQItemModelSurfaceDataProxy_YPosRoleReplaceChanged(ptr unsafe.Pointer, replace C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "yPosRoleReplaceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(replace))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectYPosRoleReplaceChanged(f func(replace string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleReplaceChanged") {
			C.QItemModelSurfaceDataProxy_ConnectYPosRoleReplaceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleReplaceChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleReplaceChanged", func(replace string) {
				signal.(func(string))(replace)
				f(replace)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleReplaceChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectYPosRoleReplaceChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectYPosRoleReplaceChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "yPosRoleReplaceChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) YPosRoleReplaceChanged(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelSurfaceDataProxy_YPosRoleReplaceChanged(ptr.Pointer(), replaceC)
	}
}

//export callbackQItemModelSurfaceDataProxy_ZPosRoleChanged
func callbackQItemModelSurfaceDataProxy_ZPosRoleChanged(ptr unsafe.Pointer, role C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "zPosRoleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(role))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectZPosRoleChanged(f func(role string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleChanged") {
			C.QItemModelSurfaceDataProxy_ConnectZPosRoleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleChanged", func(role string) {
				signal.(func(string))(role)
				f(role)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectZPosRoleChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectZPosRoleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) ZPosRoleChanged(role string) {
	if ptr.Pointer() != nil {
		var roleC *C.char
		if role != "" {
			roleC = C.CString(role)
			defer C.free(unsafe.Pointer(roleC))
		}
		C.QItemModelSurfaceDataProxy_ZPosRoleChanged(ptr.Pointer(), roleC)
	}
}

//export callbackQItemModelSurfaceDataProxy_ZPosRolePatternChanged
func callbackQItemModelSurfaceDataProxy_ZPosRolePatternChanged(ptr unsafe.Pointer, pattern unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "zPosRolePatternChanged"); signal != nil {
		signal.(func(*core.QRegExp))(core.NewQRegExpFromPointer(pattern))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectZPosRolePatternChanged(f func(pattern *core.QRegExp)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "zPosRolePatternChanged") {
			C.QItemModelSurfaceDataProxy_ConnectZPosRolePatternChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "zPosRolePatternChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRolePatternChanged", func(pattern *core.QRegExp) {
				signal.(func(*core.QRegExp))(pattern)
				f(pattern)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRolePatternChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectZPosRolePatternChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectZPosRolePatternChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRolePatternChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) ZPosRolePatternChanged(pattern core.QRegExp_ITF) {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_ZPosRolePatternChanged(ptr.Pointer(), core.PointerFromQRegExp(pattern))
	}
}

//export callbackQItemModelSurfaceDataProxy_ZPosRoleReplaceChanged
func callbackQItemModelSurfaceDataProxy_ZPosRoleReplaceChanged(ptr unsafe.Pointer, replace C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "zPosRoleReplaceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(replace))
	}

}

func (ptr *QItemModelSurfaceDataProxy) ConnectZPosRoleReplaceChanged(f func(replace string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleReplaceChanged") {
			C.QItemModelSurfaceDataProxy_ConnectZPosRoleReplaceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleReplaceChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleReplaceChanged", func(replace string) {
				signal.(func(string))(replace)
				f(replace)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleReplaceChanged", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectZPosRoleReplaceChanged() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DisconnectZPosRoleReplaceChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "zPosRoleReplaceChanged")
	}
}

func (ptr *QItemModelSurfaceDataProxy) ZPosRoleReplaceChanged(replace string) {
	if ptr.Pointer() != nil {
		var replaceC *C.char
		if replace != "" {
			replaceC = C.CString(replace)
			defer C.free(unsafe.Pointer(replaceC))
		}
		C.QItemModelSurfaceDataProxy_ZPosRoleReplaceChanged(ptr.Pointer(), replaceC)
	}
}

//export callbackQItemModelSurfaceDataProxy_DestroyQItemModelSurfaceDataProxy
func callbackQItemModelSurfaceDataProxy_DestroyQItemModelSurfaceDataProxy(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QItemModelSurfaceDataProxy"); signal != nil {
		signal.(func())()
	} else {
		NewQItemModelSurfaceDataProxyFromPointer(ptr).DestroyQItemModelSurfaceDataProxyDefault()
	}
}

func (ptr *QItemModelSurfaceDataProxy) ConnectDestroyQItemModelSurfaceDataProxy(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QItemModelSurfaceDataProxy"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QItemModelSurfaceDataProxy", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QItemModelSurfaceDataProxy", f)
		}
	}
}

func (ptr *QItemModelSurfaceDataProxy) DisconnectDestroyQItemModelSurfaceDataProxy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QItemModelSurfaceDataProxy")
	}
}

func (ptr *QItemModelSurfaceDataProxy) DestroyQItemModelSurfaceDataProxy() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DestroyQItemModelSurfaceDataProxy(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QItemModelSurfaceDataProxy) DestroyQItemModelSurfaceDataProxyDefault() {
	if ptr.Pointer() != nil {
		C.QItemModelSurfaceDataProxy_DestroyQItemModelSurfaceDataProxyDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QItemModelSurfaceDataProxy) MultiMatchBehavior() QItemModelSurfaceDataProxy__MultiMatchBehavior {
	if ptr.Pointer() != nil {
		return QItemModelSurfaceDataProxy__MultiMatchBehavior(C.QItemModelSurfaceDataProxy_MultiMatchBehavior(ptr.Pointer()))
	}
	return 0
}

func (ptr *QItemModelSurfaceDataProxy) ItemModel() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQAbstractItemModelFromPointer(C.QItemModelSurfaceDataProxy_ItemModel(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelSurfaceDataProxy) ColumnRolePattern() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QItemModelSurfaceDataProxy_ColumnRolePattern(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelSurfaceDataProxy) RowRolePattern() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QItemModelSurfaceDataProxy_RowRolePattern(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelSurfaceDataProxy) XPosRolePattern() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QItemModelSurfaceDataProxy_XPosRolePattern(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelSurfaceDataProxy) YPosRolePattern() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QItemModelSurfaceDataProxy_YPosRolePattern(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelSurfaceDataProxy) ZPosRolePattern() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QItemModelSurfaceDataProxy_ZPosRolePattern(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QItemModelSurfaceDataProxy) ColumnRole() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelSurfaceDataProxy_ColumnRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelSurfaceDataProxy) ColumnRoleReplace() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelSurfaceDataProxy_ColumnRoleReplace(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelSurfaceDataProxy) RowRole() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelSurfaceDataProxy_RowRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelSurfaceDataProxy) RowRoleReplace() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelSurfaceDataProxy_RowRoleReplace(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelSurfaceDataProxy) XPosRole() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelSurfaceDataProxy_XPosRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelSurfaceDataProxy) XPosRoleReplace() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelSurfaceDataProxy_XPosRoleReplace(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelSurfaceDataProxy) YPosRole() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelSurfaceDataProxy_YPosRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelSurfaceDataProxy) YPosRoleReplace() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelSurfaceDataProxy_YPosRoleReplace(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelSurfaceDataProxy) ZPosRole() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelSurfaceDataProxy_ZPosRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelSurfaceDataProxy) ZPosRoleReplace() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QItemModelSurfaceDataProxy_ZPosRoleReplace(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemModelSurfaceDataProxy) ColumnCategories() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QItemModelSurfaceDataProxy_ColumnCategories(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QItemModelSurfaceDataProxy) RowCategories() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QItemModelSurfaceDataProxy_RowCategories(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QItemModelSurfaceDataProxy) AutoColumnCategories() bool {
	if ptr.Pointer() != nil {
		return C.QItemModelSurfaceDataProxy_AutoColumnCategories(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QItemModelSurfaceDataProxy) AutoRowCategories() bool {
	if ptr.Pointer() != nil {
		return C.QItemModelSurfaceDataProxy_AutoRowCategories(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QItemModelSurfaceDataProxy) UseModelCategories() bool {
	if ptr.Pointer() != nil {
		return C.QItemModelSurfaceDataProxy_UseModelCategories(ptr.Pointer()) != 0
	}
	return false
}

type QLogValue3DAxisFormatter struct {
	ptr unsafe.Pointer
}

type QLogValue3DAxisFormatter_ITF interface {
	QLogValue3DAxisFormatter_PTR() *QLogValue3DAxisFormatter
}

func (ptr *QLogValue3DAxisFormatter) QLogValue3DAxisFormatter_PTR() *QLogValue3DAxisFormatter {
	return ptr
}

func (ptr *QLogValue3DAxisFormatter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLogValue3DAxisFormatter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQLogValue3DAxisFormatter(ptr QLogValue3DAxisFormatter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLogValue3DAxisFormatter_PTR().Pointer()
	}
	return nil
}

func NewQLogValue3DAxisFormatterFromPointer(ptr unsafe.Pointer) *QLogValue3DAxisFormatter {
	var n = new(QLogValue3DAxisFormatter)
	n.SetPointer(ptr)
	return n
}
func NewQLogValue3DAxisFormatter(parent core.QObject_ITF) *QLogValue3DAxisFormatter {
	return NewQLogValue3DAxisFormatterFromPointer(C.QLogValue3DAxisFormatter_NewQLogValue3DAxisFormatter(core.PointerFromQObject(parent)))
}

//export callbackQLogValue3DAxisFormatter_AutoSubGridChanged
func callbackQLogValue3DAxisFormatter_AutoSubGridChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "autoSubGridChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QLogValue3DAxisFormatter) ConnectAutoSubGridChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "autoSubGridChanged") {
			C.QLogValue3DAxisFormatter_ConnectAutoSubGridChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "autoSubGridChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "autoSubGridChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "autoSubGridChanged", f)
		}
	}
}

func (ptr *QLogValue3DAxisFormatter) DisconnectAutoSubGridChanged() {
	if ptr.Pointer() != nil {
		C.QLogValue3DAxisFormatter_DisconnectAutoSubGridChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "autoSubGridChanged")
	}
}

func (ptr *QLogValue3DAxisFormatter) AutoSubGridChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QLogValue3DAxisFormatter_AutoSubGridChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQLogValue3DAxisFormatter_BaseChanged
func callbackQLogValue3DAxisFormatter_BaseChanged(ptr unsafe.Pointer, base C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "baseChanged"); signal != nil {
		signal.(func(float64))(float64(base))
	}

}

func (ptr *QLogValue3DAxisFormatter) ConnectBaseChanged(f func(base float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "baseChanged") {
			C.QLogValue3DAxisFormatter_ConnectBaseChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "baseChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "baseChanged", func(base float64) {
				signal.(func(float64))(base)
				f(base)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "baseChanged", f)
		}
	}
}

func (ptr *QLogValue3DAxisFormatter) DisconnectBaseChanged() {
	if ptr.Pointer() != nil {
		C.QLogValue3DAxisFormatter_DisconnectBaseChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "baseChanged")
	}
}

func (ptr *QLogValue3DAxisFormatter) BaseChanged(base float64) {
	if ptr.Pointer() != nil {
		C.QLogValue3DAxisFormatter_BaseChanged(ptr.Pointer(), C.double(base))
	}
}

func (ptr *QLogValue3DAxisFormatter) SetAutoSubGrid(enabled bool) {
	if ptr.Pointer() != nil {
		C.QLogValue3DAxisFormatter_SetAutoSubGrid(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QLogValue3DAxisFormatter) SetBase(base float64) {
	if ptr.Pointer() != nil {
		C.QLogValue3DAxisFormatter_SetBase(ptr.Pointer(), C.double(base))
	}
}

func (ptr *QLogValue3DAxisFormatter) SetShowEdgeLabels(enabled bool) {
	if ptr.Pointer() != nil {
		C.QLogValue3DAxisFormatter_SetShowEdgeLabels(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQLogValue3DAxisFormatter_ShowEdgeLabelsChanged
func callbackQLogValue3DAxisFormatter_ShowEdgeLabelsChanged(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "showEdgeLabelsChanged"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	}

}

func (ptr *QLogValue3DAxisFormatter) ConnectShowEdgeLabelsChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "showEdgeLabelsChanged") {
			C.QLogValue3DAxisFormatter_ConnectShowEdgeLabelsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "showEdgeLabelsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "showEdgeLabelsChanged", func(enabled bool) {
				signal.(func(bool))(enabled)
				f(enabled)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "showEdgeLabelsChanged", f)
		}
	}
}

func (ptr *QLogValue3DAxisFormatter) DisconnectShowEdgeLabelsChanged() {
	if ptr.Pointer() != nil {
		C.QLogValue3DAxisFormatter_DisconnectShowEdgeLabelsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "showEdgeLabelsChanged")
	}
}

func (ptr *QLogValue3DAxisFormatter) ShowEdgeLabelsChanged(enabled bool) {
	if ptr.Pointer() != nil {
		C.QLogValue3DAxisFormatter_ShowEdgeLabelsChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQLogValue3DAxisFormatter_DestroyQLogValue3DAxisFormatter
func callbackQLogValue3DAxisFormatter_DestroyQLogValue3DAxisFormatter(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QLogValue3DAxisFormatter"); signal != nil {
		signal.(func())()
	} else {
		NewQLogValue3DAxisFormatterFromPointer(ptr).DestroyQLogValue3DAxisFormatterDefault()
	}
}

func (ptr *QLogValue3DAxisFormatter) ConnectDestroyQLogValue3DAxisFormatter(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QLogValue3DAxisFormatter"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QLogValue3DAxisFormatter", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QLogValue3DAxisFormatter", f)
		}
	}
}

func (ptr *QLogValue3DAxisFormatter) DisconnectDestroyQLogValue3DAxisFormatter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QLogValue3DAxisFormatter")
	}
}

func (ptr *QLogValue3DAxisFormatter) DestroyQLogValue3DAxisFormatter() {
	if ptr.Pointer() != nil {
		C.QLogValue3DAxisFormatter_DestroyQLogValue3DAxisFormatter(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLogValue3DAxisFormatter) DestroyQLogValue3DAxisFormatterDefault() {
	if ptr.Pointer() != nil {
		C.QLogValue3DAxisFormatter_DestroyQLogValue3DAxisFormatterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLogValue3DAxisFormatter) AutoSubGrid() bool {
	if ptr.Pointer() != nil {
		return C.QLogValue3DAxisFormatter_AutoSubGrid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLogValue3DAxisFormatter) ShowEdgeLabels() bool {
	if ptr.Pointer() != nil {
		return C.QLogValue3DAxisFormatter_ShowEdgeLabels(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLogValue3DAxisFormatter) Base() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLogValue3DAxisFormatter_Base(ptr.Pointer()))
	}
	return 0
}

type QScatter3DSeries struct {
	ptr unsafe.Pointer
}

type QScatter3DSeries_ITF interface {
	QScatter3DSeries_PTR() *QScatter3DSeries
}

func (ptr *QScatter3DSeries) QScatter3DSeries_PTR() *QScatter3DSeries {
	return ptr
}

func (ptr *QScatter3DSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScatter3DSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScatter3DSeries(ptr QScatter3DSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScatter3DSeries_PTR().Pointer()
	}
	return nil
}

func NewQScatter3DSeriesFromPointer(ptr unsafe.Pointer) *QScatter3DSeries {
	var n = new(QScatter3DSeries)
	n.SetPointer(ptr)
	return n
}
func NewQScatter3DSeries(parent core.QObject_ITF) *QScatter3DSeries {
	return NewQScatter3DSeriesFromPointer(C.QScatter3DSeries_NewQScatter3DSeries(core.PointerFromQObject(parent)))
}

func NewQScatter3DSeries2(dataProxy QScatterDataProxy_ITF, parent core.QObject_ITF) *QScatter3DSeries {
	return NewQScatter3DSeriesFromPointer(C.QScatter3DSeries_NewQScatter3DSeries2(PointerFromQScatterDataProxy(dataProxy), core.PointerFromQObject(parent)))
}

func QScatter3DSeries_InvalidSelectionIndex() int {
	return int(int32(C.QScatter3DSeries_QScatter3DSeries_InvalidSelectionIndex()))
}

func (ptr *QScatter3DSeries) InvalidSelectionIndex() int {
	return int(int32(C.QScatter3DSeries_QScatter3DSeries_InvalidSelectionIndex()))
}

//export callbackQScatter3DSeries_DataProxyChanged
func callbackQScatter3DSeries_DataProxyChanged(ptr unsafe.Pointer, proxy unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dataProxyChanged"); signal != nil {
		signal.(func(*QScatterDataProxy))(NewQScatterDataProxyFromPointer(proxy))
	}

}

func (ptr *QScatter3DSeries) ConnectDataProxyChanged(f func(proxy *QScatterDataProxy)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged") {
			C.QScatter3DSeries_ConnectDataProxyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged", func(proxy *QScatterDataProxy) {
				signal.(func(*QScatterDataProxy))(proxy)
				f(proxy)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged", f)
		}
	}
}

func (ptr *QScatter3DSeries) DisconnectDataProxyChanged() {
	if ptr.Pointer() != nil {
		C.QScatter3DSeries_DisconnectDataProxyChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged")
	}
}

func (ptr *QScatter3DSeries) DataProxyChanged(proxy QScatterDataProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QScatter3DSeries_DataProxyChanged(ptr.Pointer(), PointerFromQScatterDataProxy(proxy))
	}
}

//export callbackQScatter3DSeries_ItemSizeChanged
func callbackQScatter3DSeries_ItemSizeChanged(ptr unsafe.Pointer, size C.float) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemSizeChanged"); signal != nil {
		signal.(func(float32))(float32(size))
	}

}

func (ptr *QScatter3DSeries) ConnectItemSizeChanged(f func(size float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemSizeChanged") {
			C.QScatter3DSeries_ConnectItemSizeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemSizeChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemSizeChanged", func(size float32) {
				signal.(func(float32))(size)
				f(size)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemSizeChanged", f)
		}
	}
}

func (ptr *QScatter3DSeries) DisconnectItemSizeChanged() {
	if ptr.Pointer() != nil {
		C.QScatter3DSeries_DisconnectItemSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemSizeChanged")
	}
}

func (ptr *QScatter3DSeries) ItemSizeChanged(size float32) {
	if ptr.Pointer() != nil {
		C.QScatter3DSeries_ItemSizeChanged(ptr.Pointer(), C.float(size))
	}
}

//export callbackQScatter3DSeries_SelectedItemChanged
func callbackQScatter3DSeries_SelectedItemChanged(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "selectedItemChanged"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QScatter3DSeries) ConnectSelectedItemChanged(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "selectedItemChanged") {
			C.QScatter3DSeries_ConnectSelectedItemChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "selectedItemChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedItemChanged", func(index int) {
				signal.(func(int))(index)
				f(index)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedItemChanged", f)
		}
	}
}

func (ptr *QScatter3DSeries) DisconnectSelectedItemChanged() {
	if ptr.Pointer() != nil {
		C.QScatter3DSeries_DisconnectSelectedItemChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "selectedItemChanged")
	}
}

func (ptr *QScatter3DSeries) SelectedItemChanged(index int) {
	if ptr.Pointer() != nil {
		C.QScatter3DSeries_SelectedItemChanged(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QScatter3DSeries) SetDataProxy(proxy QScatterDataProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QScatter3DSeries_SetDataProxy(ptr.Pointer(), PointerFromQScatterDataProxy(proxy))
	}
}

func (ptr *QScatter3DSeries) SetItemSize(size float32) {
	if ptr.Pointer() != nil {
		C.QScatter3DSeries_SetItemSize(ptr.Pointer(), C.float(size))
	}
}

func (ptr *QScatter3DSeries) SetSelectedItem(index int) {
	if ptr.Pointer() != nil {
		C.QScatter3DSeries_SetSelectedItem(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQScatter3DSeries_DestroyQScatter3DSeries
func callbackQScatter3DSeries_DestroyQScatter3DSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QScatter3DSeries"); signal != nil {
		signal.(func())()
	} else {
		NewQScatter3DSeriesFromPointer(ptr).DestroyQScatter3DSeriesDefault()
	}
}

func (ptr *QScatter3DSeries) ConnectDestroyQScatter3DSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QScatter3DSeries"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QScatter3DSeries", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QScatter3DSeries", f)
		}
	}
}

func (ptr *QScatter3DSeries) DisconnectDestroyQScatter3DSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QScatter3DSeries")
	}
}

func (ptr *QScatter3DSeries) DestroyQScatter3DSeries() {
	if ptr.Pointer() != nil {
		C.QScatter3DSeries_DestroyQScatter3DSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScatter3DSeries) DestroyQScatter3DSeriesDefault() {
	if ptr.Pointer() != nil {
		C.QScatter3DSeries_DestroyQScatter3DSeriesDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScatter3DSeries) DataProxy() *QScatterDataProxy {
	if ptr.Pointer() != nil {
		return NewQScatterDataProxyFromPointer(C.QScatter3DSeries_DataProxy(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScatter3DSeries) ItemSize() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QScatter3DSeries_ItemSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScatter3DSeries) SelectedItem() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScatter3DSeries_SelectedItem(ptr.Pointer())))
	}
	return 0
}

type QScatterDataItem struct {
	ptr unsafe.Pointer
}

type QScatterDataItem_ITF interface {
	QScatterDataItem_PTR() *QScatterDataItem
}

func (ptr *QScatterDataItem) QScatterDataItem_PTR() *QScatterDataItem {
	return ptr
}

func (ptr *QScatterDataItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScatterDataItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScatterDataItem(ptr QScatterDataItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScatterDataItem_PTR().Pointer()
	}
	return nil
}

func NewQScatterDataItemFromPointer(ptr unsafe.Pointer) *QScatterDataItem {
	var n = new(QScatterDataItem)
	n.SetPointer(ptr)
	return n
}
func NewQScatterDataItem() *QScatterDataItem {
	var tmpValue = NewQScatterDataItemFromPointer(C.QScatterDataItem_NewQScatterDataItem())
	runtime.SetFinalizer(tmpValue, (*QScatterDataItem).DestroyQScatterDataItem)
	return tmpValue
}

func NewQScatterDataItem4(other QScatterDataItem_ITF) *QScatterDataItem {
	var tmpValue = NewQScatterDataItemFromPointer(C.QScatterDataItem_NewQScatterDataItem4(PointerFromQScatterDataItem(other)))
	runtime.SetFinalizer(tmpValue, (*QScatterDataItem).DestroyQScatterDataItem)
	return tmpValue
}

func NewQScatterDataItem2(position gui.QVector3D_ITF) *QScatterDataItem {
	var tmpValue = NewQScatterDataItemFromPointer(C.QScatterDataItem_NewQScatterDataItem2(gui.PointerFromQVector3D(position)))
	runtime.SetFinalizer(tmpValue, (*QScatterDataItem).DestroyQScatterDataItem)
	return tmpValue
}

func NewQScatterDataItem3(position gui.QVector3D_ITF, rotation gui.QQuaternion_ITF) *QScatterDataItem {
	var tmpValue = NewQScatterDataItemFromPointer(C.QScatterDataItem_NewQScatterDataItem3(gui.PointerFromQVector3D(position), gui.PointerFromQQuaternion(rotation)))
	runtime.SetFinalizer(tmpValue, (*QScatterDataItem).DestroyQScatterDataItem)
	return tmpValue
}

func (ptr *QScatterDataItem) SetPosition(pos gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterDataItem_SetPosition(ptr.Pointer(), gui.PointerFromQVector3D(pos))
	}
}

func (ptr *QScatterDataItem) SetRotation(rot gui.QQuaternion_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterDataItem_SetRotation(ptr.Pointer(), gui.PointerFromQQuaternion(rot))
	}
}

func (ptr *QScatterDataItem) SetX(value float32) {
	if ptr.Pointer() != nil {
		C.QScatterDataItem_SetX(ptr.Pointer(), C.float(value))
	}
}

func (ptr *QScatterDataItem) SetY(value float32) {
	if ptr.Pointer() != nil {
		C.QScatterDataItem_SetY(ptr.Pointer(), C.float(value))
	}
}

func (ptr *QScatterDataItem) SetZ(value float32) {
	if ptr.Pointer() != nil {
		C.QScatterDataItem_SetZ(ptr.Pointer(), C.float(value))
	}
}

func (ptr *QScatterDataItem) DestroyQScatterDataItem() {
	if ptr.Pointer() != nil {
		C.QScatterDataItem_DestroyQScatterDataItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScatterDataItem) Rotation() *gui.QQuaternion {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQQuaternionFromPointer(C.QScatterDataItem_Rotation(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QQuaternion).DestroyQQuaternion)
		return tmpValue
	}
	return nil
}

func (ptr *QScatterDataItem) Position() *gui.QVector3D {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQVector3DFromPointer(C.QScatterDataItem_Position(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QVector3D).DestroyQVector3D)
		return tmpValue
	}
	return nil
}

func (ptr *QScatterDataItem) X() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QScatterDataItem_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScatterDataItem) Y() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QScatterDataItem_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScatterDataItem) Z() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QScatterDataItem_Z(ptr.Pointer()))
	}
	return 0
}

type QScatterDataProxy struct {
	ptr unsafe.Pointer
}

type QScatterDataProxy_ITF interface {
	QScatterDataProxy_PTR() *QScatterDataProxy
}

func (ptr *QScatterDataProxy) QScatterDataProxy_PTR() *QScatterDataProxy {
	return ptr
}

func (ptr *QScatterDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScatterDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScatterDataProxy(ptr QScatterDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScatterDataProxy_PTR().Pointer()
	}
	return nil
}

func NewQScatterDataProxyFromPointer(ptr unsafe.Pointer) *QScatterDataProxy {
	var n = new(QScatterDataProxy)
	n.SetPointer(ptr)
	return n
}
func NewQScatterDataProxy(parent core.QObject_ITF) *QScatterDataProxy {
	return NewQScatterDataProxyFromPointer(C.QScatterDataProxy_NewQScatterDataProxy(core.PointerFromQObject(parent)))
}

func (ptr *QScatterDataProxy) AddItem(item QScatterDataItem_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScatterDataProxy_AddItem(ptr.Pointer(), PointerFromQScatterDataItem(item))))
	}
	return 0
}

//export callbackQScatterDataProxy_ArrayReset
func callbackQScatterDataProxy_ArrayReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "arrayReset"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScatterDataProxy) ConnectArrayReset(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "arrayReset") {
			C.QScatterDataProxy_ConnectArrayReset(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "arrayReset"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "arrayReset", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "arrayReset", f)
		}
	}
}

func (ptr *QScatterDataProxy) DisconnectArrayReset() {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_DisconnectArrayReset(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "arrayReset")
	}
}

func (ptr *QScatterDataProxy) ArrayReset() {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_ArrayReset(ptr.Pointer())
	}
}

func (ptr *QScatterDataProxy) InsertItem(index int, item QScatterDataItem_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_InsertItem(ptr.Pointer(), C.int(int32(index)), PointerFromQScatterDataItem(item))
	}
}

//export callbackQScatterDataProxy_ItemCountChanged
func callbackQScatterDataProxy_ItemCountChanged(ptr unsafe.Pointer, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemCountChanged"); signal != nil {
		signal.(func(int))(int(int32(count)))
	}

}

func (ptr *QScatterDataProxy) ConnectItemCountChanged(f func(count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemCountChanged") {
			C.QScatterDataProxy_ConnectItemCountChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemCountChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemCountChanged", func(count int) {
				signal.(func(int))(count)
				f(count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemCountChanged", f)
		}
	}
}

func (ptr *QScatterDataProxy) DisconnectItemCountChanged() {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_DisconnectItemCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemCountChanged")
	}
}

func (ptr *QScatterDataProxy) ItemCountChanged(count int) {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_ItemCountChanged(ptr.Pointer(), C.int(int32(count)))
	}
}

//export callbackQScatterDataProxy_ItemsAdded
func callbackQScatterDataProxy_ItemsAdded(ptr unsafe.Pointer, startIndex C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemsAdded"); signal != nil {
		signal.(func(int, int))(int(int32(startIndex)), int(int32(count)))
	}

}

func (ptr *QScatterDataProxy) ConnectItemsAdded(f func(startIndex int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemsAdded") {
			C.QScatterDataProxy_ConnectItemsAdded(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemsAdded"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemsAdded", func(startIndex int, count int) {
				signal.(func(int, int))(startIndex, count)
				f(startIndex, count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemsAdded", f)
		}
	}
}

func (ptr *QScatterDataProxy) DisconnectItemsAdded() {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_DisconnectItemsAdded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemsAdded")
	}
}

func (ptr *QScatterDataProxy) ItemsAdded(startIndex int, count int) {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_ItemsAdded(ptr.Pointer(), C.int(int32(startIndex)), C.int(int32(count)))
	}
}

//export callbackQScatterDataProxy_ItemsChanged
func callbackQScatterDataProxy_ItemsChanged(ptr unsafe.Pointer, startIndex C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemsChanged"); signal != nil {
		signal.(func(int, int))(int(int32(startIndex)), int(int32(count)))
	}

}

func (ptr *QScatterDataProxy) ConnectItemsChanged(f func(startIndex int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemsChanged") {
			C.QScatterDataProxy_ConnectItemsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemsChanged", func(startIndex int, count int) {
				signal.(func(int, int))(startIndex, count)
				f(startIndex, count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemsChanged", f)
		}
	}
}

func (ptr *QScatterDataProxy) DisconnectItemsChanged() {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_DisconnectItemsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemsChanged")
	}
}

func (ptr *QScatterDataProxy) ItemsChanged(startIndex int, count int) {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_ItemsChanged(ptr.Pointer(), C.int(int32(startIndex)), C.int(int32(count)))
	}
}

//export callbackQScatterDataProxy_ItemsInserted
func callbackQScatterDataProxy_ItemsInserted(ptr unsafe.Pointer, startIndex C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemsInserted"); signal != nil {
		signal.(func(int, int))(int(int32(startIndex)), int(int32(count)))
	}

}

func (ptr *QScatterDataProxy) ConnectItemsInserted(f func(startIndex int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemsInserted") {
			C.QScatterDataProxy_ConnectItemsInserted(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemsInserted"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemsInserted", func(startIndex int, count int) {
				signal.(func(int, int))(startIndex, count)
				f(startIndex, count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemsInserted", f)
		}
	}
}

func (ptr *QScatterDataProxy) DisconnectItemsInserted() {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_DisconnectItemsInserted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemsInserted")
	}
}

func (ptr *QScatterDataProxy) ItemsInserted(startIndex int, count int) {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_ItemsInserted(ptr.Pointer(), C.int(int32(startIndex)), C.int(int32(count)))
	}
}

//export callbackQScatterDataProxy_ItemsRemoved
func callbackQScatterDataProxy_ItemsRemoved(ptr unsafe.Pointer, startIndex C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemsRemoved"); signal != nil {
		signal.(func(int, int))(int(int32(startIndex)), int(int32(count)))
	}

}

func (ptr *QScatterDataProxy) ConnectItemsRemoved(f func(startIndex int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemsRemoved") {
			C.QScatterDataProxy_ConnectItemsRemoved(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemsRemoved"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemsRemoved", func(startIndex int, count int) {
				signal.(func(int, int))(startIndex, count)
				f(startIndex, count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemsRemoved", f)
		}
	}
}

func (ptr *QScatterDataProxy) DisconnectItemsRemoved() {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_DisconnectItemsRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemsRemoved")
	}
}

func (ptr *QScatterDataProxy) ItemsRemoved(startIndex int, count int) {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_ItemsRemoved(ptr.Pointer(), C.int(int32(startIndex)), C.int(int32(count)))
	}
}

func (ptr *QScatterDataProxy) RemoveItems(index int, removeCount int) {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_RemoveItems(ptr.Pointer(), C.int(int32(index)), C.int(int32(removeCount)))
	}
}

//export callbackQScatterDataProxy_SeriesChanged
func callbackQScatterDataProxy_SeriesChanged(ptr unsafe.Pointer, series unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "seriesChanged"); signal != nil {
		signal.(func(*QScatter3DSeries))(NewQScatter3DSeriesFromPointer(series))
	}

}

func (ptr *QScatterDataProxy) ConnectSeriesChanged(f func(series *QScatter3DSeries)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged") {
			C.QScatterDataProxy_ConnectSeriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged", func(series *QScatter3DSeries) {
				signal.(func(*QScatter3DSeries))(series)
				f(series)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged", f)
		}
	}
}

func (ptr *QScatterDataProxy) DisconnectSeriesChanged() {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_DisconnectSeriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged")
	}
}

func (ptr *QScatterDataProxy) SeriesChanged(series QScatter3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_SeriesChanged(ptr.Pointer(), PointerFromQScatter3DSeries(series))
	}
}

func (ptr *QScatterDataProxy) SetItem(index int, item QScatterDataItem_ITF) {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_SetItem(ptr.Pointer(), C.int(int32(index)), PointerFromQScatterDataItem(item))
	}
}

//export callbackQScatterDataProxy_DestroyQScatterDataProxy
func callbackQScatterDataProxy_DestroyQScatterDataProxy(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QScatterDataProxy"); signal != nil {
		signal.(func())()
	} else {
		NewQScatterDataProxyFromPointer(ptr).DestroyQScatterDataProxyDefault()
	}
}

func (ptr *QScatterDataProxy) ConnectDestroyQScatterDataProxy(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QScatterDataProxy"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QScatterDataProxy", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QScatterDataProxy", f)
		}
	}
}

func (ptr *QScatterDataProxy) DisconnectDestroyQScatterDataProxy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QScatterDataProxy")
	}
}

func (ptr *QScatterDataProxy) DestroyQScatterDataProxy() {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_DestroyQScatterDataProxy(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScatterDataProxy) DestroyQScatterDataProxyDefault() {
	if ptr.Pointer() != nil {
		C.QScatterDataProxy_DestroyQScatterDataProxyDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScatterDataProxy) Series() *QScatter3DSeries {
	if ptr.Pointer() != nil {
		return NewQScatter3DSeriesFromPointer(C.QScatterDataProxy_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScatterDataProxy) ItemAt(index int) *QScatterDataItem {
	if ptr.Pointer() != nil {
		return NewQScatterDataItemFromPointer(C.QScatterDataProxy_ItemAt(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

func (ptr *QScatterDataProxy) ItemCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScatterDataProxy_ItemCount(ptr.Pointer())))
	}
	return 0
}

type QSurface3DSeries struct {
	ptr unsafe.Pointer
}

type QSurface3DSeries_ITF interface {
	QSurface3DSeries_PTR() *QSurface3DSeries
}

func (ptr *QSurface3DSeries) QSurface3DSeries_PTR() *QSurface3DSeries {
	return ptr
}

func (ptr *QSurface3DSeries) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSurface3DSeries) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSurface3DSeries(ptr QSurface3DSeries_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSurface3DSeries_PTR().Pointer()
	}
	return nil
}

func NewQSurface3DSeriesFromPointer(ptr unsafe.Pointer) *QSurface3DSeries {
	var n = new(QSurface3DSeries)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QSurface3DSeries__DrawFlag
//QSurface3DSeries::DrawFlag
type QSurface3DSeries__DrawFlag int64

const (
	QSurface3DSeries__DrawWireframe           QSurface3DSeries__DrawFlag = QSurface3DSeries__DrawFlag(1)
	QSurface3DSeries__DrawSurface             QSurface3DSeries__DrawFlag = QSurface3DSeries__DrawFlag(2)
	QSurface3DSeries__DrawSurfaceAndWireframe QSurface3DSeries__DrawFlag = QSurface3DSeries__DrawFlag(QSurface3DSeries__DrawWireframe | QSurface3DSeries__DrawSurface)
)

//export callbackQSurface3DSeries_SelectedPointChanged
func callbackQSurface3DSeries_SelectedPointChanged(ptr unsafe.Pointer, position unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "selectedPointChanged"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(position))
	}

}

func (ptr *QSurface3DSeries) ConnectSelectedPointChanged(f func(position *core.QPoint)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "selectedPointChanged") {
			C.QSurface3DSeries_ConnectSelectedPointChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "selectedPointChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedPointChanged", func(position *core.QPoint) {
				signal.(func(*core.QPoint))(position)
				f(position)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "selectedPointChanged", f)
		}
	}
}

func (ptr *QSurface3DSeries) DisconnectSelectedPointChanged() {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_DisconnectSelectedPointChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "selectedPointChanged")
	}
}

func (ptr *QSurface3DSeries) SelectedPointChanged(position core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_SelectedPointChanged(ptr.Pointer(), core.PointerFromQPoint(position))
	}
}

func (ptr *QSurface3DSeries) SetDataProxy(proxy QSurfaceDataProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_SetDataProxy(ptr.Pointer(), PointerFromQSurfaceDataProxy(proxy))
	}
}

func QSurface3DSeries_InvalidSelectionPosition() *core.QPoint {
	var tmpValue = core.NewQPointFromPointer(C.QSurface3DSeries_QSurface3DSeries_InvalidSelectionPosition())
	runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
	return tmpValue
}

func (ptr *QSurface3DSeries) InvalidSelectionPosition() *core.QPoint {
	var tmpValue = core.NewQPointFromPointer(C.QSurface3DSeries_QSurface3DSeries_InvalidSelectionPosition())
	runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
	return tmpValue
}

func NewQSurface3DSeries(parent core.QObject_ITF) *QSurface3DSeries {
	return NewQSurface3DSeriesFromPointer(C.QSurface3DSeries_NewQSurface3DSeries(core.PointerFromQObject(parent)))
}

func NewQSurface3DSeries2(dataProxy QSurfaceDataProxy_ITF, parent core.QObject_ITF) *QSurface3DSeries {
	return NewQSurface3DSeriesFromPointer(C.QSurface3DSeries_NewQSurface3DSeries2(PointerFromQSurfaceDataProxy(dataProxy), core.PointerFromQObject(parent)))
}

//export callbackQSurface3DSeries_DataProxyChanged
func callbackQSurface3DSeries_DataProxyChanged(ptr unsafe.Pointer, proxy unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "dataProxyChanged"); signal != nil {
		signal.(func(*QSurfaceDataProxy))(NewQSurfaceDataProxyFromPointer(proxy))
	}

}

func (ptr *QSurface3DSeries) ConnectDataProxyChanged(f func(proxy *QSurfaceDataProxy)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged") {
			C.QSurface3DSeries_ConnectDataProxyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged", func(proxy *QSurfaceDataProxy) {
				signal.(func(*QSurfaceDataProxy))(proxy)
				f(proxy)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged", f)
		}
	}
}

func (ptr *QSurface3DSeries) DisconnectDataProxyChanged() {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_DisconnectDataProxyChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "dataProxyChanged")
	}
}

func (ptr *QSurface3DSeries) DataProxyChanged(proxy QSurfaceDataProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_DataProxyChanged(ptr.Pointer(), PointerFromQSurfaceDataProxy(proxy))
	}
}

//export callbackQSurface3DSeries_DrawModeChanged
func callbackQSurface3DSeries_DrawModeChanged(ptr unsafe.Pointer, mode C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "drawModeChanged"); signal != nil {
		signal.(func(QSurface3DSeries__DrawFlag))(QSurface3DSeries__DrawFlag(mode))
	}

}

func (ptr *QSurface3DSeries) ConnectDrawModeChanged(f func(mode QSurface3DSeries__DrawFlag)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "drawModeChanged") {
			C.QSurface3DSeries_ConnectDrawModeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "drawModeChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "drawModeChanged", func(mode QSurface3DSeries__DrawFlag) {
				signal.(func(QSurface3DSeries__DrawFlag))(mode)
				f(mode)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "drawModeChanged", f)
		}
	}
}

func (ptr *QSurface3DSeries) DisconnectDrawModeChanged() {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_DisconnectDrawModeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "drawModeChanged")
	}
}

func (ptr *QSurface3DSeries) DrawModeChanged(mode QSurface3DSeries__DrawFlag) {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_DrawModeChanged(ptr.Pointer(), C.longlong(mode))
	}
}

//export callbackQSurface3DSeries_FlatShadingEnabledChanged
func callbackQSurface3DSeries_FlatShadingEnabledChanged(ptr unsafe.Pointer, enable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "flatShadingEnabledChanged"); signal != nil {
		signal.(func(bool))(int8(enable) != 0)
	}

}

func (ptr *QSurface3DSeries) ConnectFlatShadingEnabledChanged(f func(enable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "flatShadingEnabledChanged") {
			C.QSurface3DSeries_ConnectFlatShadingEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "flatShadingEnabledChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "flatShadingEnabledChanged", func(enable bool) {
				signal.(func(bool))(enable)
				f(enable)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "flatShadingEnabledChanged", f)
		}
	}
}

func (ptr *QSurface3DSeries) DisconnectFlatShadingEnabledChanged() {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_DisconnectFlatShadingEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "flatShadingEnabledChanged")
	}
}

func (ptr *QSurface3DSeries) FlatShadingEnabledChanged(enable bool) {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_FlatShadingEnabledChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQSurface3DSeries_FlatShadingSupportedChanged
func callbackQSurface3DSeries_FlatShadingSupportedChanged(ptr unsafe.Pointer, enable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "flatShadingSupportedChanged"); signal != nil {
		signal.(func(bool))(int8(enable) != 0)
	}

}

func (ptr *QSurface3DSeries) ConnectFlatShadingSupportedChanged(f func(enable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "flatShadingSupportedChanged") {
			C.QSurface3DSeries_ConnectFlatShadingSupportedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "flatShadingSupportedChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "flatShadingSupportedChanged", func(enable bool) {
				signal.(func(bool))(enable)
				f(enable)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "flatShadingSupportedChanged", f)
		}
	}
}

func (ptr *QSurface3DSeries) DisconnectFlatShadingSupportedChanged() {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_DisconnectFlatShadingSupportedChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "flatShadingSupportedChanged")
	}
}

func (ptr *QSurface3DSeries) FlatShadingSupportedChanged(enable bool) {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_FlatShadingSupportedChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QSurface3DSeries) SetDrawMode(mode QSurface3DSeries__DrawFlag) {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_SetDrawMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QSurface3DSeries) SetFlatShadingEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_SetFlatShadingEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QSurface3DSeries) SetSelectedPoint(position core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_SetSelectedPoint(ptr.Pointer(), core.PointerFromQPoint(position))
	}
}

func (ptr *QSurface3DSeries) SetTexture(texture gui.QImage_ITF) {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_SetTexture(ptr.Pointer(), gui.PointerFromQImage(texture))
	}
}

func (ptr *QSurface3DSeries) SetTextureFile(filename string) {
	if ptr.Pointer() != nil {
		var filenameC *C.char
		if filename != "" {
			filenameC = C.CString(filename)
			defer C.free(unsafe.Pointer(filenameC))
		}
		C.QSurface3DSeries_SetTextureFile(ptr.Pointer(), filenameC)
	}
}

//export callbackQSurface3DSeries_TextureChanged
func callbackQSurface3DSeries_TextureChanged(ptr unsafe.Pointer, image unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "textureChanged"); signal != nil {
		signal.(func(*gui.QImage))(gui.NewQImageFromPointer(image))
	}

}

func (ptr *QSurface3DSeries) ConnectTextureChanged(f func(image *gui.QImage)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "textureChanged") {
			C.QSurface3DSeries_ConnectTextureChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "textureChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureChanged", func(image *gui.QImage) {
				signal.(func(*gui.QImage))(image)
				f(image)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureChanged", f)
		}
	}
}

func (ptr *QSurface3DSeries) DisconnectTextureChanged() {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_DisconnectTextureChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "textureChanged")
	}
}

func (ptr *QSurface3DSeries) TextureChanged(image gui.QImage_ITF) {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_TextureChanged(ptr.Pointer(), gui.PointerFromQImage(image))
	}
}

//export callbackQSurface3DSeries_TextureFileChanged
func callbackQSurface3DSeries_TextureFileChanged(ptr unsafe.Pointer, filename C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "textureFileChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(filename))
	}

}

func (ptr *QSurface3DSeries) ConnectTextureFileChanged(f func(filename string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "textureFileChanged") {
			C.QSurface3DSeries_ConnectTextureFileChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "textureFileChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureFileChanged", func(filename string) {
				signal.(func(string))(filename)
				f(filename)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "textureFileChanged", f)
		}
	}
}

func (ptr *QSurface3DSeries) DisconnectTextureFileChanged() {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_DisconnectTextureFileChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "textureFileChanged")
	}
}

func (ptr *QSurface3DSeries) TextureFileChanged(filename string) {
	if ptr.Pointer() != nil {
		var filenameC *C.char
		if filename != "" {
			filenameC = C.CString(filename)
			defer C.free(unsafe.Pointer(filenameC))
		}
		C.QSurface3DSeries_TextureFileChanged(ptr.Pointer(), filenameC)
	}
}

//export callbackQSurface3DSeries_DestroyQSurface3DSeries
func callbackQSurface3DSeries_DestroyQSurface3DSeries(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QSurface3DSeries"); signal != nil {
		signal.(func())()
	} else {
		NewQSurface3DSeriesFromPointer(ptr).DestroyQSurface3DSeriesDefault()
	}
}

func (ptr *QSurface3DSeries) ConnectDestroyQSurface3DSeries(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QSurface3DSeries"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QSurface3DSeries", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QSurface3DSeries", f)
		}
	}
}

func (ptr *QSurface3DSeries) DisconnectDestroyQSurface3DSeries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QSurface3DSeries")
	}
}

func (ptr *QSurface3DSeries) DestroyQSurface3DSeries() {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_DestroyQSurface3DSeries(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSurface3DSeries) DestroyQSurface3DSeriesDefault() {
	if ptr.Pointer() != nil {
		C.QSurface3DSeries_DestroyQSurface3DSeriesDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSurface3DSeries) Texture() *gui.QImage {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQImageFromPointer(C.QSurface3DSeries_Texture(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QSurface3DSeries) SelectedPoint() *core.QPoint {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFromPointer(C.QSurface3DSeries_SelectedPoint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QSurface3DSeries) TextureFile() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSurface3DSeries_TextureFile(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSurface3DSeries) DrawMode() QSurface3DSeries__DrawFlag {
	if ptr.Pointer() != nil {
		return QSurface3DSeries__DrawFlag(C.QSurface3DSeries_DrawMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurface3DSeries) DataProxy() *QSurfaceDataProxy {
	if ptr.Pointer() != nil {
		return NewQSurfaceDataProxyFromPointer(C.QSurface3DSeries_DataProxy(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSurface3DSeries) IsFlatShadingEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QSurface3DSeries_IsFlatShadingEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSurface3DSeries) IsFlatShadingSupported() bool {
	if ptr.Pointer() != nil {
		return C.QSurface3DSeries_IsFlatShadingSupported(ptr.Pointer()) != 0
	}
	return false
}

type QSurfaceDataItem struct {
	ptr unsafe.Pointer
}

type QSurfaceDataItem_ITF interface {
	QSurfaceDataItem_PTR() *QSurfaceDataItem
}

func (ptr *QSurfaceDataItem) QSurfaceDataItem_PTR() *QSurfaceDataItem {
	return ptr
}

func (ptr *QSurfaceDataItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSurfaceDataItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSurfaceDataItem(ptr QSurfaceDataItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSurfaceDataItem_PTR().Pointer()
	}
	return nil
}

func NewQSurfaceDataItemFromPointer(ptr unsafe.Pointer) *QSurfaceDataItem {
	var n = new(QSurfaceDataItem)
	n.SetPointer(ptr)
	return n
}
func NewQSurfaceDataItem() *QSurfaceDataItem {
	var tmpValue = NewQSurfaceDataItemFromPointer(C.QSurfaceDataItem_NewQSurfaceDataItem())
	runtime.SetFinalizer(tmpValue, (*QSurfaceDataItem).DestroyQSurfaceDataItem)
	return tmpValue
}

func NewQSurfaceDataItem3(other QSurfaceDataItem_ITF) *QSurfaceDataItem {
	var tmpValue = NewQSurfaceDataItemFromPointer(C.QSurfaceDataItem_NewQSurfaceDataItem3(PointerFromQSurfaceDataItem(other)))
	runtime.SetFinalizer(tmpValue, (*QSurfaceDataItem).DestroyQSurfaceDataItem)
	return tmpValue
}

func NewQSurfaceDataItem2(position gui.QVector3D_ITF) *QSurfaceDataItem {
	var tmpValue = NewQSurfaceDataItemFromPointer(C.QSurfaceDataItem_NewQSurfaceDataItem2(gui.PointerFromQVector3D(position)))
	runtime.SetFinalizer(tmpValue, (*QSurfaceDataItem).DestroyQSurfaceDataItem)
	return tmpValue
}

func (ptr *QSurfaceDataItem) SetPosition(pos gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataItem_SetPosition(ptr.Pointer(), gui.PointerFromQVector3D(pos))
	}
}

func (ptr *QSurfaceDataItem) SetX(value float32) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataItem_SetX(ptr.Pointer(), C.float(value))
	}
}

func (ptr *QSurfaceDataItem) SetY(value float32) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataItem_SetY(ptr.Pointer(), C.float(value))
	}
}

func (ptr *QSurfaceDataItem) SetZ(value float32) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataItem_SetZ(ptr.Pointer(), C.float(value))
	}
}

func (ptr *QSurfaceDataItem) DestroyQSurfaceDataItem() {
	if ptr.Pointer() != nil {
		C.QSurfaceDataItem_DestroyQSurfaceDataItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSurfaceDataItem) Position() *gui.QVector3D {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQVector3DFromPointer(C.QSurfaceDataItem_Position(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QVector3D).DestroyQVector3D)
		return tmpValue
	}
	return nil
}

func (ptr *QSurfaceDataItem) X() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QSurfaceDataItem_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceDataItem) Y() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QSurfaceDataItem_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceDataItem) Z() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QSurfaceDataItem_Z(ptr.Pointer()))
	}
	return 0
}

type QSurfaceDataProxy struct {
	ptr unsafe.Pointer
}

type QSurfaceDataProxy_ITF interface {
	QSurfaceDataProxy_PTR() *QSurfaceDataProxy
}

func (ptr *QSurfaceDataProxy) QSurfaceDataProxy_PTR() *QSurfaceDataProxy {
	return ptr
}

func (ptr *QSurfaceDataProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSurfaceDataProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSurfaceDataProxy(ptr QSurfaceDataProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSurfaceDataProxy_PTR().Pointer()
	}
	return nil
}

func NewQSurfaceDataProxyFromPointer(ptr unsafe.Pointer) *QSurfaceDataProxy {
	var n = new(QSurfaceDataProxy)
	n.SetPointer(ptr)
	return n
}
func NewQSurfaceDataProxy(parent core.QObject_ITF) *QSurfaceDataProxy {
	return NewQSurfaceDataProxyFromPointer(C.QSurfaceDataProxy_NewQSurfaceDataProxy(core.PointerFromQObject(parent)))
}

//export callbackQSurfaceDataProxy_ArrayReset
func callbackQSurfaceDataProxy_ArrayReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "arrayReset"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSurfaceDataProxy) ConnectArrayReset(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "arrayReset") {
			C.QSurfaceDataProxy_ConnectArrayReset(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "arrayReset"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "arrayReset", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "arrayReset", f)
		}
	}
}

func (ptr *QSurfaceDataProxy) DisconnectArrayReset() {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_DisconnectArrayReset(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "arrayReset")
	}
}

func (ptr *QSurfaceDataProxy) ArrayReset() {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_ArrayReset(ptr.Pointer())
	}
}

//export callbackQSurfaceDataProxy_ColumnCountChanged
func callbackQSurfaceDataProxy_ColumnCountChanged(ptr unsafe.Pointer, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "columnCountChanged"); signal != nil {
		signal.(func(int))(int(int32(count)))
	}

}

func (ptr *QSurfaceDataProxy) ConnectColumnCountChanged(f func(count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "columnCountChanged") {
			C.QSurfaceDataProxy_ConnectColumnCountChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "columnCountChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnCountChanged", func(count int) {
				signal.(func(int))(count)
				f(count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "columnCountChanged", f)
		}
	}
}

func (ptr *QSurfaceDataProxy) DisconnectColumnCountChanged() {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_DisconnectColumnCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "columnCountChanged")
	}
}

func (ptr *QSurfaceDataProxy) ColumnCountChanged(count int) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_ColumnCountChanged(ptr.Pointer(), C.int(int32(count)))
	}
}

//export callbackQSurfaceDataProxy_ItemChanged
func callbackQSurfaceDataProxy_ItemChanged(ptr unsafe.Pointer, rowIndex C.int, columnIndex C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "itemChanged"); signal != nil {
		signal.(func(int, int))(int(int32(rowIndex)), int(int32(columnIndex)))
	}

}

func (ptr *QSurfaceDataProxy) ConnectItemChanged(f func(rowIndex int, columnIndex int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "itemChanged") {
			C.QSurfaceDataProxy_ConnectItemChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "itemChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemChanged", func(rowIndex int, columnIndex int) {
				signal.(func(int, int))(rowIndex, columnIndex)
				f(rowIndex, columnIndex)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "itemChanged", f)
		}
	}
}

func (ptr *QSurfaceDataProxy) DisconnectItemChanged() {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_DisconnectItemChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "itemChanged")
	}
}

func (ptr *QSurfaceDataProxy) ItemChanged(rowIndex int, columnIndex int) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_ItemChanged(ptr.Pointer(), C.int(int32(rowIndex)), C.int(int32(columnIndex)))
	}
}

func (ptr *QSurfaceDataProxy) RemoveRows(rowIndex int, removeCount int) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_RemoveRows(ptr.Pointer(), C.int(int32(rowIndex)), C.int(int32(removeCount)))
	}
}

//export callbackQSurfaceDataProxy_RowCountChanged
func callbackQSurfaceDataProxy_RowCountChanged(ptr unsafe.Pointer, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowCountChanged"); signal != nil {
		signal.(func(int))(int(int32(count)))
	}

}

func (ptr *QSurfaceDataProxy) ConnectRowCountChanged(f func(count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowCountChanged") {
			C.QSurfaceDataProxy_ConnectRowCountChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowCountChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowCountChanged", func(count int) {
				signal.(func(int))(count)
				f(count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowCountChanged", f)
		}
	}
}

func (ptr *QSurfaceDataProxy) DisconnectRowCountChanged() {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_DisconnectRowCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowCountChanged")
	}
}

func (ptr *QSurfaceDataProxy) RowCountChanged(count int) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_RowCountChanged(ptr.Pointer(), C.int(int32(count)))
	}
}

//export callbackQSurfaceDataProxy_RowsAdded
func callbackQSurfaceDataProxy_RowsAdded(ptr unsafe.Pointer, startIndex C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowsAdded"); signal != nil {
		signal.(func(int, int))(int(int32(startIndex)), int(int32(count)))
	}

}

func (ptr *QSurfaceDataProxy) ConnectRowsAdded(f func(startIndex int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowsAdded") {
			C.QSurfaceDataProxy_ConnectRowsAdded(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowsAdded"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsAdded", func(startIndex int, count int) {
				signal.(func(int, int))(startIndex, count)
				f(startIndex, count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsAdded", f)
		}
	}
}

func (ptr *QSurfaceDataProxy) DisconnectRowsAdded() {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_DisconnectRowsAdded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowsAdded")
	}
}

func (ptr *QSurfaceDataProxy) RowsAdded(startIndex int, count int) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_RowsAdded(ptr.Pointer(), C.int(int32(startIndex)), C.int(int32(count)))
	}
}

//export callbackQSurfaceDataProxy_RowsChanged
func callbackQSurfaceDataProxy_RowsChanged(ptr unsafe.Pointer, startIndex C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowsChanged"); signal != nil {
		signal.(func(int, int))(int(int32(startIndex)), int(int32(count)))
	}

}

func (ptr *QSurfaceDataProxy) ConnectRowsChanged(f func(startIndex int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowsChanged") {
			C.QSurfaceDataProxy_ConnectRowsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowsChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsChanged", func(startIndex int, count int) {
				signal.(func(int, int))(startIndex, count)
				f(startIndex, count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsChanged", f)
		}
	}
}

func (ptr *QSurfaceDataProxy) DisconnectRowsChanged() {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_DisconnectRowsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowsChanged")
	}
}

func (ptr *QSurfaceDataProxy) RowsChanged(startIndex int, count int) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_RowsChanged(ptr.Pointer(), C.int(int32(startIndex)), C.int(int32(count)))
	}
}

//export callbackQSurfaceDataProxy_RowsInserted
func callbackQSurfaceDataProxy_RowsInserted(ptr unsafe.Pointer, startIndex C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowsInserted"); signal != nil {
		signal.(func(int, int))(int(int32(startIndex)), int(int32(count)))
	}

}

func (ptr *QSurfaceDataProxy) ConnectRowsInserted(f func(startIndex int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowsInserted") {
			C.QSurfaceDataProxy_ConnectRowsInserted(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowsInserted"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsInserted", func(startIndex int, count int) {
				signal.(func(int, int))(startIndex, count)
				f(startIndex, count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsInserted", f)
		}
	}
}

func (ptr *QSurfaceDataProxy) DisconnectRowsInserted() {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_DisconnectRowsInserted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowsInserted")
	}
}

func (ptr *QSurfaceDataProxy) RowsInserted(startIndex int, count int) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_RowsInserted(ptr.Pointer(), C.int(int32(startIndex)), C.int(int32(count)))
	}
}

//export callbackQSurfaceDataProxy_RowsRemoved
func callbackQSurfaceDataProxy_RowsRemoved(ptr unsafe.Pointer, startIndex C.int, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rowsRemoved"); signal != nil {
		signal.(func(int, int))(int(int32(startIndex)), int(int32(count)))
	}

}

func (ptr *QSurfaceDataProxy) ConnectRowsRemoved(f func(startIndex int, count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "rowsRemoved") {
			C.QSurfaceDataProxy_ConnectRowsRemoved(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "rowsRemoved"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsRemoved", func(startIndex int, count int) {
				signal.(func(int, int))(startIndex, count)
				f(startIndex, count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rowsRemoved", f)
		}
	}
}

func (ptr *QSurfaceDataProxy) DisconnectRowsRemoved() {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_DisconnectRowsRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rowsRemoved")
	}
}

func (ptr *QSurfaceDataProxy) RowsRemoved(startIndex int, count int) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_RowsRemoved(ptr.Pointer(), C.int(int32(startIndex)), C.int(int32(count)))
	}
}

//export callbackQSurfaceDataProxy_SeriesChanged
func callbackQSurfaceDataProxy_SeriesChanged(ptr unsafe.Pointer, series unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "seriesChanged"); signal != nil {
		signal.(func(*QSurface3DSeries))(NewQSurface3DSeriesFromPointer(series))
	}

}

func (ptr *QSurfaceDataProxy) ConnectSeriesChanged(f func(series *QSurface3DSeries)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged") {
			C.QSurfaceDataProxy_ConnectSeriesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged", func(series *QSurface3DSeries) {
				signal.(func(*QSurface3DSeries))(series)
				f(series)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged", f)
		}
	}
}

func (ptr *QSurfaceDataProxy) DisconnectSeriesChanged() {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_DisconnectSeriesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "seriesChanged")
	}
}

func (ptr *QSurfaceDataProxy) SeriesChanged(series QSurface3DSeries_ITF) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_SeriesChanged(ptr.Pointer(), PointerFromQSurface3DSeries(series))
	}
}

func (ptr *QSurfaceDataProxy) SetItem2(position core.QPoint_ITF, item QSurfaceDataItem_ITF) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_SetItem2(ptr.Pointer(), core.PointerFromQPoint(position), PointerFromQSurfaceDataItem(item))
	}
}

func (ptr *QSurfaceDataProxy) SetItem(rowIndex int, columnIndex int, item QSurfaceDataItem_ITF) {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_SetItem(ptr.Pointer(), C.int(int32(rowIndex)), C.int(int32(columnIndex)), PointerFromQSurfaceDataItem(item))
	}
}

//export callbackQSurfaceDataProxy_DestroyQSurfaceDataProxy
func callbackQSurfaceDataProxy_DestroyQSurfaceDataProxy(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QSurfaceDataProxy"); signal != nil {
		signal.(func())()
	} else {
		NewQSurfaceDataProxyFromPointer(ptr).DestroyQSurfaceDataProxyDefault()
	}
}

func (ptr *QSurfaceDataProxy) ConnectDestroyQSurfaceDataProxy(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QSurfaceDataProxy"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QSurfaceDataProxy", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QSurfaceDataProxy", f)
		}
	}
}

func (ptr *QSurfaceDataProxy) DisconnectDestroyQSurfaceDataProxy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QSurfaceDataProxy")
	}
}

func (ptr *QSurfaceDataProxy) DestroyQSurfaceDataProxy() {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_DestroyQSurfaceDataProxy(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSurfaceDataProxy) DestroyQSurfaceDataProxyDefault() {
	if ptr.Pointer() != nil {
		C.QSurfaceDataProxy_DestroyQSurfaceDataProxyDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSurfaceDataProxy) Series() *QSurface3DSeries {
	if ptr.Pointer() != nil {
		return NewQSurface3DSeriesFromPointer(C.QSurfaceDataProxy_Series(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSurfaceDataProxy) ItemAt2(position core.QPoint_ITF) *QSurfaceDataItem {
	if ptr.Pointer() != nil {
		return NewQSurfaceDataItemFromPointer(C.QSurfaceDataProxy_ItemAt2(ptr.Pointer(), core.PointerFromQPoint(position)))
	}
	return nil
}

func (ptr *QSurfaceDataProxy) ItemAt(rowIndex int, columnIndex int) *QSurfaceDataItem {
	if ptr.Pointer() != nil {
		return NewQSurfaceDataItemFromPointer(C.QSurfaceDataProxy_ItemAt(ptr.Pointer(), C.int(int32(rowIndex)), C.int(int32(columnIndex))))
	}
	return nil
}

func (ptr *QSurfaceDataProxy) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSurfaceDataProxy_ColumnCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSurfaceDataProxy) RowCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSurfaceDataProxy_RowCount(ptr.Pointer())))
	}
	return 0
}

type QTouch3DInputHandler struct {
	ptr unsafe.Pointer
}

type QTouch3DInputHandler_ITF interface {
	QTouch3DInputHandler_PTR() *QTouch3DInputHandler
}

func (ptr *QTouch3DInputHandler) QTouch3DInputHandler_PTR() *QTouch3DInputHandler {
	return ptr
}

func (ptr *QTouch3DInputHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTouch3DInputHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTouch3DInputHandler(ptr QTouch3DInputHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTouch3DInputHandler_PTR().Pointer()
	}
	return nil
}

func NewQTouch3DInputHandlerFromPointer(ptr unsafe.Pointer) *QTouch3DInputHandler {
	var n = new(QTouch3DInputHandler)
	n.SetPointer(ptr)
	return n
}
func NewQTouch3DInputHandler(parent core.QObject_ITF) *QTouch3DInputHandler {
	return NewQTouch3DInputHandlerFromPointer(C.QTouch3DInputHandler_NewQTouch3DInputHandler(core.PointerFromQObject(parent)))
}

//export callbackQTouch3DInputHandler_TouchEvent
func callbackQTouch3DInputHandler_TouchEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
	} else {
		NewQTouch3DInputHandlerFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(event))
	}
}

func (ptr *QTouch3DInputHandler) ConnectTouchEvent(f func(event *gui.QTouchEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "touchEvent"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "touchEvent", func(event *gui.QTouchEvent) {
				signal.(func(*gui.QTouchEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "touchEvent", f)
		}
	}
}

func (ptr *QTouch3DInputHandler) DisconnectTouchEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "touchEvent")
	}
}

func (ptr *QTouch3DInputHandler) TouchEvent(event gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTouch3DInputHandler_TouchEvent(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

func (ptr *QTouch3DInputHandler) TouchEventDefault(event gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTouch3DInputHandler_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

//export callbackQTouch3DInputHandler_DestroyQTouch3DInputHandler
func callbackQTouch3DInputHandler_DestroyQTouch3DInputHandler(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QTouch3DInputHandler"); signal != nil {
		signal.(func())()
	} else {
		NewQTouch3DInputHandlerFromPointer(ptr).DestroyQTouch3DInputHandlerDefault()
	}
}

func (ptr *QTouch3DInputHandler) ConnectDestroyQTouch3DInputHandler(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QTouch3DInputHandler"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QTouch3DInputHandler", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QTouch3DInputHandler", f)
		}
	}
}

func (ptr *QTouch3DInputHandler) DisconnectDestroyQTouch3DInputHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QTouch3DInputHandler")
	}
}

func (ptr *QTouch3DInputHandler) DestroyQTouch3DInputHandler() {
	if ptr.Pointer() != nil {
		C.QTouch3DInputHandler_DestroyQTouch3DInputHandler(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTouch3DInputHandler) DestroyQTouch3DInputHandlerDefault() {
	if ptr.Pointer() != nil {
		C.QTouch3DInputHandler_DestroyQTouch3DInputHandlerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QValue3DAxis struct {
	ptr unsafe.Pointer
}

type QValue3DAxis_ITF interface {
	QValue3DAxis_PTR() *QValue3DAxis
}

func (ptr *QValue3DAxis) QValue3DAxis_PTR() *QValue3DAxis {
	return ptr
}

func (ptr *QValue3DAxis) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QValue3DAxis) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQValue3DAxis(ptr QValue3DAxis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QValue3DAxis_PTR().Pointer()
	}
	return nil
}

func NewQValue3DAxisFromPointer(ptr unsafe.Pointer) *QValue3DAxis {
	var n = new(QValue3DAxis)
	n.SetPointer(ptr)
	return n
}
func NewQValue3DAxis(parent core.QObject_ITF) *QValue3DAxis {
	return NewQValue3DAxisFromPointer(C.QValue3DAxis_NewQValue3DAxis(core.PointerFromQObject(parent)))
}

//export callbackQValue3DAxis_FormatterChanged
func callbackQValue3DAxis_FormatterChanged(ptr unsafe.Pointer, formatter unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "formatterChanged"); signal != nil {
		signal.(func(*QValue3DAxisFormatter))(NewQValue3DAxisFormatterFromPointer(formatter))
	}

}

func (ptr *QValue3DAxis) ConnectFormatterChanged(f func(formatter *QValue3DAxisFormatter)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "formatterChanged") {
			C.QValue3DAxis_ConnectFormatterChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "formatterChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "formatterChanged", func(formatter *QValue3DAxisFormatter) {
				signal.(func(*QValue3DAxisFormatter))(formatter)
				f(formatter)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "formatterChanged", f)
		}
	}
}

func (ptr *QValue3DAxis) DisconnectFormatterChanged() {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_DisconnectFormatterChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "formatterChanged")
	}
}

func (ptr *QValue3DAxis) FormatterChanged(formatter QValue3DAxisFormatter_ITF) {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_FormatterChanged(ptr.Pointer(), PointerFromQValue3DAxisFormatter(formatter))
	}
}

//export callbackQValue3DAxis_LabelFormatChanged
func callbackQValue3DAxis_LabelFormatChanged(ptr unsafe.Pointer, format C.struct_QtDataVisualization_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "labelFormatChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(format))
	}

}

func (ptr *QValue3DAxis) ConnectLabelFormatChanged(f func(format string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "labelFormatChanged") {
			C.QValue3DAxis_ConnectLabelFormatChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "labelFormatChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelFormatChanged", func(format string) {
				signal.(func(string))(format)
				f(format)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "labelFormatChanged", f)
		}
	}
}

func (ptr *QValue3DAxis) DisconnectLabelFormatChanged() {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_DisconnectLabelFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "labelFormatChanged")
	}
}

func (ptr *QValue3DAxis) LabelFormatChanged(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QValue3DAxis_LabelFormatChanged(ptr.Pointer(), formatC)
	}
}

//export callbackQValue3DAxis_ReversedChanged
func callbackQValue3DAxis_ReversedChanged(ptr unsafe.Pointer, enable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "reversedChanged"); signal != nil {
		signal.(func(bool))(int8(enable) != 0)
	}

}

func (ptr *QValue3DAxis) ConnectReversedChanged(f func(enable bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "reversedChanged") {
			C.QValue3DAxis_ConnectReversedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "reversedChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "reversedChanged", func(enable bool) {
				signal.(func(bool))(enable)
				f(enable)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "reversedChanged", f)
		}
	}
}

func (ptr *QValue3DAxis) DisconnectReversedChanged() {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_DisconnectReversedChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "reversedChanged")
	}
}

func (ptr *QValue3DAxis) ReversedChanged(enable bool) {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_ReversedChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQValue3DAxis_SegmentCountChanged
func callbackQValue3DAxis_SegmentCountChanged(ptr unsafe.Pointer, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "segmentCountChanged"); signal != nil {
		signal.(func(int))(int(int32(count)))
	}

}

func (ptr *QValue3DAxis) ConnectSegmentCountChanged(f func(count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "segmentCountChanged") {
			C.QValue3DAxis_ConnectSegmentCountChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "segmentCountChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "segmentCountChanged", func(count int) {
				signal.(func(int))(count)
				f(count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "segmentCountChanged", f)
		}
	}
}

func (ptr *QValue3DAxis) DisconnectSegmentCountChanged() {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_DisconnectSegmentCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "segmentCountChanged")
	}
}

func (ptr *QValue3DAxis) SegmentCountChanged(count int) {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_SegmentCountChanged(ptr.Pointer(), C.int(int32(count)))
	}
}

func (ptr *QValue3DAxis) SetFormatter(formatter QValue3DAxisFormatter_ITF) {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_SetFormatter(ptr.Pointer(), PointerFromQValue3DAxisFormatter(formatter))
	}
}

func (ptr *QValue3DAxis) SetLabelFormat(format string) {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		C.QValue3DAxis_SetLabelFormat(ptr.Pointer(), formatC)
	}
}

func (ptr *QValue3DAxis) SetReversed(enable bool) {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_SetReversed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QValue3DAxis) SetSegmentCount(count int) {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_SetSegmentCount(ptr.Pointer(), C.int(int32(count)))
	}
}

func (ptr *QValue3DAxis) SetSubSegmentCount(count int) {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_SetSubSegmentCount(ptr.Pointer(), C.int(int32(count)))
	}
}

//export callbackQValue3DAxis_SubSegmentCountChanged
func callbackQValue3DAxis_SubSegmentCountChanged(ptr unsafe.Pointer, count C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "subSegmentCountChanged"); signal != nil {
		signal.(func(int))(int(int32(count)))
	}

}

func (ptr *QValue3DAxis) ConnectSubSegmentCountChanged(f func(count int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "subSegmentCountChanged") {
			C.QValue3DAxis_ConnectSubSegmentCountChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "subSegmentCountChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "subSegmentCountChanged", func(count int) {
				signal.(func(int))(count)
				f(count)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "subSegmentCountChanged", f)
		}
	}
}

func (ptr *QValue3DAxis) DisconnectSubSegmentCountChanged() {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_DisconnectSubSegmentCountChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "subSegmentCountChanged")
	}
}

func (ptr *QValue3DAxis) SubSegmentCountChanged(count int) {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_SubSegmentCountChanged(ptr.Pointer(), C.int(int32(count)))
	}
}

//export callbackQValue3DAxis_DestroyQValue3DAxis
func callbackQValue3DAxis_DestroyQValue3DAxis(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QValue3DAxis"); signal != nil {
		signal.(func())()
	} else {
		NewQValue3DAxisFromPointer(ptr).DestroyQValue3DAxisDefault()
	}
}

func (ptr *QValue3DAxis) ConnectDestroyQValue3DAxis(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QValue3DAxis"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QValue3DAxis", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QValue3DAxis", f)
		}
	}
}

func (ptr *QValue3DAxis) DisconnectDestroyQValue3DAxis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QValue3DAxis")
	}
}

func (ptr *QValue3DAxis) DestroyQValue3DAxis() {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_DestroyQValue3DAxis(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QValue3DAxis) DestroyQValue3DAxisDefault() {
	if ptr.Pointer() != nil {
		C.QValue3DAxis_DestroyQValue3DAxisDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QValue3DAxis) LabelFormat() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QValue3DAxis_LabelFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QValue3DAxis) Formatter() *QValue3DAxisFormatter {
	if ptr.Pointer() != nil {
		return NewQValue3DAxisFormatterFromPointer(C.QValue3DAxis_Formatter(ptr.Pointer()))
	}
	return nil
}

func (ptr *QValue3DAxis) Reversed() bool {
	if ptr.Pointer() != nil {
		return C.QValue3DAxis_Reversed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QValue3DAxis) SegmentCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QValue3DAxis_SegmentCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QValue3DAxis) SubSegmentCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QValue3DAxis_SubSegmentCount(ptr.Pointer())))
	}
	return 0
}

type QValue3DAxisFormatter struct {
	ptr unsafe.Pointer
}

type QValue3DAxisFormatter_ITF interface {
	QValue3DAxisFormatter_PTR() *QValue3DAxisFormatter
}

func (ptr *QValue3DAxisFormatter) QValue3DAxisFormatter_PTR() *QValue3DAxisFormatter {
	return ptr
}

func (ptr *QValue3DAxisFormatter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QValue3DAxisFormatter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQValue3DAxisFormatter(ptr QValue3DAxisFormatter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QValue3DAxisFormatter_PTR().Pointer()
	}
	return nil
}

func NewQValue3DAxisFormatterFromPointer(ptr unsafe.Pointer) *QValue3DAxisFormatter {
	var n = new(QValue3DAxisFormatter)
	n.SetPointer(ptr)
	return n
}
func NewQValue3DAxisFormatter(parent core.QObject_ITF) *QValue3DAxisFormatter {
	return NewQValue3DAxisFormatterFromPointer(C.QValue3DAxisFormatter_NewQValue3DAxisFormatter(core.PointerFromQObject(parent)))
}

func (ptr *QValue3DAxisFormatter) MarkDirty(labelsChange bool) {
	if ptr.Pointer() != nil {
		C.QValue3DAxisFormatter_MarkDirty(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(labelsChange))))
	}
}

//export callbackQValue3DAxisFormatter_Recalculate
func callbackQValue3DAxisFormatter_Recalculate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "recalculate"); signal != nil {
		signal.(func())()
	} else {
		NewQValue3DAxisFormatterFromPointer(ptr).RecalculateDefault()
	}
}

func (ptr *QValue3DAxisFormatter) ConnectRecalculate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "recalculate"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "recalculate", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "recalculate", f)
		}
	}
}

func (ptr *QValue3DAxisFormatter) DisconnectRecalculate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "recalculate")
	}
}

func (ptr *QValue3DAxisFormatter) Recalculate() {
	if ptr.Pointer() != nil {
		C.QValue3DAxisFormatter_Recalculate(ptr.Pointer())
	}
}

func (ptr *QValue3DAxisFormatter) RecalculateDefault() {
	if ptr.Pointer() != nil {
		C.QValue3DAxisFormatter_RecalculateDefault(ptr.Pointer())
	}
}

func (ptr *QValue3DAxisFormatter) SetAllowNegatives(allow bool) {
	if ptr.Pointer() != nil {
		C.QValue3DAxisFormatter_SetAllowNegatives(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(allow))))
	}
}

func (ptr *QValue3DAxisFormatter) SetAllowZero(allow bool) {
	if ptr.Pointer() != nil {
		C.QValue3DAxisFormatter_SetAllowZero(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(allow))))
	}
}

func (ptr *QValue3DAxisFormatter) SetLocale(locale core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QValue3DAxisFormatter_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

//export callbackQValue3DAxisFormatter_DestroyQValue3DAxisFormatter
func callbackQValue3DAxisFormatter_DestroyQValue3DAxisFormatter(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QValue3DAxisFormatter"); signal != nil {
		signal.(func())()
	} else {
		NewQValue3DAxisFormatterFromPointer(ptr).DestroyQValue3DAxisFormatterDefault()
	}
}

func (ptr *QValue3DAxisFormatter) ConnectDestroyQValue3DAxisFormatter(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "~QValue3DAxisFormatter"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QValue3DAxisFormatter", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QValue3DAxisFormatter", f)
		}
	}
}

func (ptr *QValue3DAxisFormatter) DisconnectDestroyQValue3DAxisFormatter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QValue3DAxisFormatter")
	}
}

func (ptr *QValue3DAxisFormatter) DestroyQValue3DAxisFormatter() {
	if ptr.Pointer() != nil {
		C.QValue3DAxisFormatter_DestroyQValue3DAxisFormatter(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QValue3DAxisFormatter) DestroyQValue3DAxisFormatterDefault() {
	if ptr.Pointer() != nil {
		C.QValue3DAxisFormatter_DestroyQValue3DAxisFormatterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QValue3DAxisFormatter) Locale() *core.QLocale {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQLocaleFromPointer(C.QValue3DAxisFormatter_Locale(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QLocale).DestroyQLocale)
		return tmpValue
	}
	return nil
}

//export callbackQValue3DAxisFormatter_StringForValue
func callbackQValue3DAxisFormatter_StringForValue(ptr unsafe.Pointer, value C.double, format C.struct_QtDataVisualization_PackedString) *C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "stringForValue"); signal != nil {
		return C.CString(signal.(func(float64, string) string)(float64(value), cGoUnpackString(format)))
	}

	return C.CString(NewQValue3DAxisFormatterFromPointer(ptr).StringForValueDefault(float64(value), cGoUnpackString(format)))
}

func (ptr *QValue3DAxisFormatter) ConnectStringForValue(f func(value float64, format string) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "stringForValue"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "stringForValue", func(value float64, format string) string {
				signal.(func(float64, string) string)(value, format)
				return f(value, format)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "stringForValue", f)
		}
	}
}

func (ptr *QValue3DAxisFormatter) DisconnectStringForValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "stringForValue")
	}
}

func (ptr *QValue3DAxisFormatter) StringForValue(value float64, format string) string {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return cGoUnpackString(C.QValue3DAxisFormatter_StringForValue(ptr.Pointer(), C.double(value), formatC))
	}
	return ""
}

func (ptr *QValue3DAxisFormatter) StringForValueDefault(value float64, format string) string {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return cGoUnpackString(C.QValue3DAxisFormatter_StringForValueDefault(ptr.Pointer(), C.double(value), formatC))
	}
	return ""
}

func (ptr *QValue3DAxisFormatter) LabelStrings() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QValue3DAxisFormatter_LabelStrings(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QValue3DAxisFormatter) Axis() *QValue3DAxis {
	if ptr.Pointer() != nil {
		return NewQValue3DAxisFromPointer(C.QValue3DAxisFormatter_Axis(ptr.Pointer()))
	}
	return nil
}

//export callbackQValue3DAxisFormatter_CreateNewInstance
func callbackQValue3DAxisFormatter_CreateNewInstance(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "createNewInstance"); signal != nil {
		return PointerFromQValue3DAxisFormatter(signal.(func() *QValue3DAxisFormatter)())
	}

	return PointerFromQValue3DAxisFormatter(NewQValue3DAxisFormatterFromPointer(ptr).CreateNewInstanceDefault())
}

func (ptr *QValue3DAxisFormatter) ConnectCreateNewInstance(f func() *QValue3DAxisFormatter) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "createNewInstance"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "createNewInstance", func() *QValue3DAxisFormatter {
				signal.(func() *QValue3DAxisFormatter)()
				return f()
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "createNewInstance", f)
		}
	}
}

func (ptr *QValue3DAxisFormatter) DisconnectCreateNewInstance() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "createNewInstance")
	}
}

func (ptr *QValue3DAxisFormatter) CreateNewInstance() *QValue3DAxisFormatter {
	if ptr.Pointer() != nil {
		return NewQValue3DAxisFormatterFromPointer(C.QValue3DAxisFormatter_CreateNewInstance(ptr.Pointer()))
	}
	return nil
}

func (ptr *QValue3DAxisFormatter) CreateNewInstanceDefault() *QValue3DAxisFormatter {
	if ptr.Pointer() != nil {
		return NewQValue3DAxisFormatterFromPointer(C.QValue3DAxisFormatter_CreateNewInstanceDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QValue3DAxisFormatter) GridPositions() []float32 {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []float32 {
			var out = make([]float32, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQValue3DAxisFormatterFromPointer(l.data).__gridPositions_atList(i)
			}
			return out
		}(C.QValue3DAxisFormatter_GridPositions(ptr.Pointer()))
	}
	return make([]float32, 0)
}

func (ptr *QValue3DAxisFormatter) LabelPositions() []float32 {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []float32 {
			var out = make([]float32, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQValue3DAxisFormatterFromPointer(l.data).__labelPositions_atList(i)
			}
			return out
		}(C.QValue3DAxisFormatter_LabelPositions(ptr.Pointer()))
	}
	return make([]float32, 0)
}

func (ptr *QValue3DAxisFormatter) SubGridPositions() []float32 {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDataVisualization_PackedList) []float32 {
			var out = make([]float32, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQValue3DAxisFormatterFromPointer(l.data).__subGridPositions_atList(i)
			}
			return out
		}(C.QValue3DAxisFormatter_SubGridPositions(ptr.Pointer()))
	}
	return make([]float32, 0)
}

func (ptr *QValue3DAxisFormatter) AllowNegatives() bool {
	if ptr.Pointer() != nil {
		return C.QValue3DAxisFormatter_AllowNegatives(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QValue3DAxisFormatter) AllowZero() bool {
	if ptr.Pointer() != nil {
		return C.QValue3DAxisFormatter_AllowZero(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQValue3DAxisFormatter_PositionAt
func callbackQValue3DAxisFormatter_PositionAt(ptr unsafe.Pointer, value C.float) C.float {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "positionAt"); signal != nil {
		return C.float(signal.(func(float32) float32)(float32(value)))
	}

	return C.float(NewQValue3DAxisFormatterFromPointer(ptr).PositionAtDefault(float32(value)))
}

func (ptr *QValue3DAxisFormatter) ConnectPositionAt(f func(value float32) float32) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "positionAt"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "positionAt", func(value float32) float32 {
				signal.(func(float32) float32)(value)
				return f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "positionAt", f)
		}
	}
}

func (ptr *QValue3DAxisFormatter) DisconnectPositionAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "positionAt")
	}
}

func (ptr *QValue3DAxisFormatter) PositionAt(value float32) float32 {
	if ptr.Pointer() != nil {
		return float32(C.QValue3DAxisFormatter_PositionAt(ptr.Pointer(), C.float(value)))
	}
	return 0
}

func (ptr *QValue3DAxisFormatter) PositionAtDefault(value float32) float32 {
	if ptr.Pointer() != nil {
		return float32(C.QValue3DAxisFormatter_PositionAtDefault(ptr.Pointer(), C.float(value)))
	}
	return 0
}

//export callbackQValue3DAxisFormatter_ValueAt
func callbackQValue3DAxisFormatter_ValueAt(ptr unsafe.Pointer, position C.float) C.float {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "valueAt"); signal != nil {
		return C.float(signal.(func(float32) float32)(float32(position)))
	}

	return C.float(NewQValue3DAxisFormatterFromPointer(ptr).ValueAtDefault(float32(position)))
}

func (ptr *QValue3DAxisFormatter) ConnectValueAt(f func(position float32) float32) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "valueAt"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "valueAt", func(position float32) float32 {
				signal.(func(float32) float32)(position)
				return f(position)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "valueAt", f)
		}
	}
}

func (ptr *QValue3DAxisFormatter) DisconnectValueAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "valueAt")
	}
}

func (ptr *QValue3DAxisFormatter) ValueAt(position float32) float32 {
	if ptr.Pointer() != nil {
		return float32(C.QValue3DAxisFormatter_ValueAt(ptr.Pointer(), C.float(position)))
	}
	return 0
}

func (ptr *QValue3DAxisFormatter) ValueAtDefault(position float32) float32 {
	if ptr.Pointer() != nil {
		return float32(C.QValue3DAxisFormatter_ValueAtDefault(ptr.Pointer(), C.float(position)))
	}
	return 0
}

//export callbackQValue3DAxisFormatter_PopulateCopy
func callbackQValue3DAxisFormatter_PopulateCopy(ptr unsafe.Pointer, copy unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "populateCopy"); signal != nil {
		signal.(func(*QValue3DAxisFormatter))(NewQValue3DAxisFormatterFromPointer(copy))
	} else {
		NewQValue3DAxisFormatterFromPointer(ptr).PopulateCopyDefault(NewQValue3DAxisFormatterFromPointer(copy))
	}
}

func (ptr *QValue3DAxisFormatter) ConnectPopulateCopy(f func(copy *QValue3DAxisFormatter)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "populateCopy"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "populateCopy", func(copy *QValue3DAxisFormatter) {
				signal.(func(*QValue3DAxisFormatter))(copy)
				f(copy)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "populateCopy", f)
		}
	}
}

func (ptr *QValue3DAxisFormatter) DisconnectPopulateCopy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "populateCopy")
	}
}

func (ptr *QValue3DAxisFormatter) PopulateCopy(copy QValue3DAxisFormatter_ITF) {
	if ptr.Pointer() != nil {
		C.QValue3DAxisFormatter_PopulateCopy(ptr.Pointer(), PointerFromQValue3DAxisFormatter(copy))
	}
}

func (ptr *QValue3DAxisFormatter) PopulateCopyDefault(copy QValue3DAxisFormatter_ITF) {
	if ptr.Pointer() != nil {
		C.QValue3DAxisFormatter_PopulateCopyDefault(ptr.Pointer(), PointerFromQValue3DAxisFormatter(copy))
	}
}

func (ptr *QValue3DAxisFormatter) __gridPositions_atList(i int) float32 {
	if ptr.Pointer() != nil {
		return float32(C.QValue3DAxisFormatter___gridPositions_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QValue3DAxisFormatter) __gridPositions_setList(i float32) {
	if ptr.Pointer() != nil {
		C.QValue3DAxisFormatter___gridPositions_setList(ptr.Pointer(), C.float(i))
	}
}

func (ptr *QValue3DAxisFormatter) __gridPositions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QValue3DAxisFormatter___gridPositions_newList(ptr.Pointer()))
}

func (ptr *QValue3DAxisFormatter) __labelPositions_atList(i int) float32 {
	if ptr.Pointer() != nil {
		return float32(C.QValue3DAxisFormatter___labelPositions_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QValue3DAxisFormatter) __labelPositions_setList(i float32) {
	if ptr.Pointer() != nil {
		C.QValue3DAxisFormatter___labelPositions_setList(ptr.Pointer(), C.float(i))
	}
}

func (ptr *QValue3DAxisFormatter) __labelPositions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QValue3DAxisFormatter___labelPositions_newList(ptr.Pointer()))
}

func (ptr *QValue3DAxisFormatter) __subGridPositions_atList(i int) float32 {
	if ptr.Pointer() != nil {
		return float32(C.QValue3DAxisFormatter___subGridPositions_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QValue3DAxisFormatter) __subGridPositions_setList(i float32) {
	if ptr.Pointer() != nil {
		C.QValue3DAxisFormatter___subGridPositions_setList(ptr.Pointer(), C.float(i))
	}
}

func (ptr *QValue3DAxisFormatter) __subGridPositions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QValue3DAxisFormatter___subGridPositions_newList(ptr.Pointer()))
}
