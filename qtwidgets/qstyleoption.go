package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QStyleOption struct {
	*qtrt.CObject
}
type QStyleOption_ITF interface {
	QStyleOption_PTR() *QStyleOption
}

func (ptr *QStyleOption) QStyleOption_PTR() *QStyleOption { return ptr }

func (this *QStyleOption) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStyleOption) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStyleOptionFromPointer(cthis unsafe.Pointer) *QStyleOption {
	return &QStyleOption{&qtrt.CObject{cthis}}
}
func (*QStyleOption) NewFromPointer(cthis unsafe.Pointer) *QStyleOption {
	return NewQStyleOptionFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOption(int, int)

/*
Constructs a QStyleOption with the specified version and type.

The version has no special meaning for QStyleOption; it can be used by subclasses to distinguish between different version of the same option type.

The state member variable is initialized to QStyle::State_None.

See also version and type.
*/
func NewQStyleOption(version int, type_ int) *QStyleOption {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStyleOptionC2Eii", qtrt.FFI_TYPE_POINTER, version, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOption)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOption(int, int)

/*
Constructs a QStyleOption with the specified version and type.

The version has no special meaning for QStyleOption; it can be used by subclasses to distinguish between different version of the same option type.

The state member variable is initialized to QStyle::State_None.

See also version and type.
*/
func NewQStyleOption__() *QStyleOption {
	// arg: 0, int=Int, =Invalid,
	version := 0 /*QStyleOption::Version*/
	// arg: 1, int=Int, =Invalid,
	type_ := QStyleOption__SO_Default
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStyleOptionC2Eii", qtrt.FFI_TYPE_POINTER, version, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOption)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOption(int, int)

/*
Constructs a QStyleOption with the specified version and type.

The version has no special meaning for QStyleOption; it can be used by subclasses to distinguish between different version of the same option type.

The state member variable is initialized to QStyle::State_None.

See also version and type.
*/
func NewQStyleOption__1(version int) *QStyleOption {
	// arg: 1, int=Int, =Invalid,
	type_ := QStyleOption__SO_Default
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStyleOptionC2Eii", qtrt.FFI_TYPE_POINTER, version, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOption)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QStyleOption()

/*

 */
func DeleteQStyleOption(this *QStyleOption) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStyleOptionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 64)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void init(const QWidget *)

/*

 */
func (this *QStyleOption) Init(w QWidget_ITF /*777 const QWidget **/) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStyleOption4initEPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:107
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void initFrom(const QWidget *)

/*
Initializes the state, direction, rect, palette, fontMetrics and styleObject member variables based on the specified widget.

This is a convenience function; the member variables can also be initialized manually.

This function was introduced in  Qt 4.1.

See also QWidget::layoutDirection(), QWidget::rect(), QWidget::palette(), and QWidget::fontMetrics().
*/
func (this *QStyleOption) InitFrom(w QWidget_ITF /*777 const QWidget **/) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStyleOption8initFromEPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:108
// index:0
// Public Visibility=Default Availability=Available
// [64] QStyleOption & operator=(const QStyleOption &)

/*

 */
func (this *QStyleOption) Operator_equal(other QStyleOption_ITF) *QStyleOption {
	var convArg0 unsafe.Pointer
	if other != nil && other.QStyleOption_PTR() != nil {
		convArg0 = other.QStyleOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QStyleOptionaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStyleOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStyleOption)
	return rv2
}

/*
This enum is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

QStyleOption::SO_ComboBox?QStyleOptionComboBox
QStyleOption::SO_GroupBox?QStyleOptionGroupBox
QStyleOption::SO_SizeGrip?QStyleOptionSizeGrip
QStyleOption::SO_Slider?QStyleOptionSlider
QStyleOption::SO_SpinBox?QStyleOptionSpinBox
QStyleOption::SO_TitleBar?QStyleOptionTitleBar
QStyleOption::SO_ToolButton?QStyleOptionToolButton


The following values are used for custom controls:



See also type.

*/
type QStyleOption__OptionType = int

// QStyleOption
const QStyleOption__SO_Default QStyleOption__OptionType = 0

// QStyleOptionFocusRect
const QStyleOption__SO_FocusRect QStyleOption__OptionType = 1

// QStyleOptionButton
const QStyleOption__SO_Button QStyleOption__OptionType = 2

// QStyleOptionTab
const QStyleOption__SO_Tab QStyleOption__OptionType = 3

// QStyleOptionMenuItem
const QStyleOption__SO_MenuItem QStyleOption__OptionType = 4

// QStyleOptionFrame
const QStyleOption__SO_Frame QStyleOption__OptionType = 5

// QStyleOptionProgressBar
const QStyleOption__SO_ProgressBar QStyleOption__OptionType = 6

// QStyleOptionToolBox
const QStyleOption__SO_ToolBox QStyleOption__OptionType = 7

// QStyleOptionHeader
const QStyleOption__SO_Header QStyleOption__OptionType = 8

// QStyleOptionDockWidget
const QStyleOption__SO_DockWidget QStyleOption__OptionType = 9

//
const QStyleOption__SO_ViewItem QStyleOption__OptionType = 10

//
const QStyleOption__SO_TabWidgetFrame QStyleOption__OptionType = 11

//
const QStyleOption__SO_TabBarBase QStyleOption__OptionType = 12

//
const QStyleOption__SO_RubberBand QStyleOption__OptionType = 13

//
const QStyleOption__SO_ToolBar QStyleOption__OptionType = 14

//
const QStyleOption__SO_GraphicsItem QStyleOption__OptionType = 15

//
const QStyleOption__SO_Complex QStyleOption__OptionType = 983040

//
const QStyleOption__SO_Slider QStyleOption__OptionType = 983041

//
const QStyleOption__SO_SpinBox QStyleOption__OptionType = 983042

//
const QStyleOption__SO_ToolButton QStyleOption__OptionType = 983043

//
const QStyleOption__SO_ComboBox QStyleOption__OptionType = 983044

//
const QStyleOption__SO_TitleBar QStyleOption__OptionType = 983045

//
const QStyleOption__SO_GroupBox QStyleOption__OptionType = 983046

//
const QStyleOption__SO_SizeGrip QStyleOption__OptionType = 983047

//
const QStyleOption__SO_CustomBase QStyleOption__OptionType = 3840

//
const QStyleOption__SO_ComplexCustomBase QStyleOption__OptionType = 251658240

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOption__StyleOptionType = int

//
const QStyleOption__Type QStyleOption__StyleOptionType = 0

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOption__StyleOptionVersion = int

// 1
const QStyleOption__Version QStyleOption__StyleOptionVersion = 1

//  body block end

//  keep block begin

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
