package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtWidgets/qstyle.h
// dst-file: /src/widgets/qstyle.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QStyle::drawItemPixmap(QPainter * painter, const QRect & rect, int alignment, const QPixmap & pixmap);
extern void C_ZNK6QStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap(void* qthis, void* arg0, void* arg1, int32_t arg2, void* arg3); // 4
  // proto:  void QStyle::~QStyle();
extern void C_ZN6QStyleD2Ev(void* qthis); // 4
  // proto:  void QStyle::polish(QWidget * );
extern void C_ZN6QStyle6polishEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QStyle::polish(QPalette & );
extern void C_ZN6QStyle6polishER8QPalette(void* qthis, void* arg0); // 4
  // proto:  void QStyle::polish(QApplication * );
extern void C_ZN6QStyle6polishEP12QApplication(void* qthis, void* arg0); // 4
  // proto:  QPalette QStyle::standardPalette();
extern void C_ZNK6QStyle15standardPaletteEv(void* qthis); // 4
  // proto:  void QStyle::QStyle();
extern void* C_ZN6QStyleC2Ev(); // 3
  // proto: static int QStyle::sliderValueFromPosition(int min, int max, int pos, int space, bool upsideDown);
extern void C_ZN6QStyle23sliderValueFromPositionEiiiib(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, bool arg4); // 4
  // proto:  const QStyle * QStyle::proxy();
extern void C_ZNK6QStyle5proxyEv(void* qthis); // 4
  // proto:  const QMetaObject * QStyle::metaObject();
extern void C_ZNK6QStyle10metaObjectEv(void* qthis); // 4
  // proto:  QRect QStyle::itemTextRect(const QFontMetrics & fm, const QRect & r, int flags, bool enabled, const QString & text);
extern void C_ZNK6QStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString(void* qthis, void* arg0, void* arg1, int32_t arg2, bool arg3, void* arg4); // 4
  // proto:  void QStyle::unpolish(QWidget * );
extern void C_ZN6QStyle8unpolishEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QStyle::unpolish(QApplication * );
extern void C_ZN6QStyle8unpolishEP12QApplication(void* qthis, void* arg0); // 4
  // proto: static int QStyle::sliderPositionFromValue(int min, int max, int val, int space, bool upsideDown);
extern void C_ZN6QStyle23sliderPositionFromValueEiiiib(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, bool arg4); // 4
  // proto:  QRect QStyle::itemPixmapRect(const QRect & r, int flags, const QPixmap & pixmap);
extern void C_ZNK6QStyle14itemPixmapRectERK5QRectiRK7QPixmap(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QStyle)=1
type QStyle struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
func (this *QStyle) drawItemPixmap(args ...interface{}) () {
  // drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap
    // invoke: void drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QRect).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QPixmap).qclsinst
    if false {fmt.Println(arg3)}
    C.C_ZNK6QStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QStyle", "drawItemPixmap", args)
  }

}

// ~QStyle()
func (this *QStyle) FreeQStyle(args ...interface{}) () {
  // ~QStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStyleD0Ev
    // invoke: void ~QStyle()
    C.C_ZN6QStyleD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyle", "~QStyle", args)
  }

}

// polish(class QWidget *)
func (this *QStyle) polish(args ...interface{}) () {
  // polish(class QWidget *)
  // polish(class QPalette &)
  // polish(class QApplication *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPalette{}) // "QPalette &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QApplication{}) // "QApplication *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStyle6polishEP7QWidget
    // invoke: void polish(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QStyle6polishEP7QWidget(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN6QStyle6polishER8QPalette
    // invoke: void polish(class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QStyle6polishER8QPalette(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN6QStyle6polishEP12QApplication
    // invoke: void polish(class QApplication *)
    var arg0 = args[0].(QApplication).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QStyle6polishEP12QApplication(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyle", "polish", args)
  }

}

// standardPalette()
func (this *QStyle) standardPalette(args ...interface{}) () {
  // standardPalette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle15standardPaletteEv
    // invoke: QPalette standardPalette()
    var ret = C.C_ZNK6QStyle15standardPaletteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStyle", "standardPalette", args)
  }

}

// QStyle()
func NewQStyle(args ...interface{}) *QStyle {
  // QStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStyleC1Ev
    // invoke: void QStyle()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QStyleC2Ev()
    return &QStyle{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyle", "QStyle", args)
  }

  return nil // QStyle{qclsinst:qthis}
}

// sliderValueFromPosition(int, int, int, int, _Bool)
func (this *QStyle) sliderValueFromPosition_s(args ...interface{}) () {
  // sliderValueFromPosition(int, int, int, int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStyle23sliderValueFromPositionEiiiib
    // invoke: int sliderValueFromPosition(int, int, int, int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.bool(args[4].(bool))
    if false {fmt.Println(arg4)}
    var ret = C.C_ZN6QStyle23sliderValueFromPositionEiiiib(arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStyle", "sliderValueFromPosition", args)
  }

}

// proxy()
func (this *QStyle) proxy(args ...interface{}) () {
  // proxy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle5proxyEv
    // invoke: const QStyle * proxy()
    var ret = C.C_ZNK6QStyle5proxyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStyle", "proxy", args)
  }

}

// metaObject()
func (this *QStyle) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK6QStyle10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyle", "metaObject", args)
  }

}

// itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
func (this *QStyle) itemTextRect(args ...interface{}) () {
  // itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFontMetrics{}) // "const QFontMetrics &"
  vtys[0][1] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.BoolTy(false) // "bool"
  vtys[0][4] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString
    // invoke: QRect itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
    var arg0 = args[0].(QFontMetrics).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QRect).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.bool(args[3].(bool))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QString).qclsinst
    if false {fmt.Println(arg4)}
    var ret = C.C_ZNK6QStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStyle", "itemTextRect", args)
  }

}

// unpolish(class QWidget *)
func (this *QStyle) unpolish(args ...interface{}) () {
  // unpolish(class QWidget *)
  // unpolish(class QApplication *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QApplication{}) // "QApplication *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStyle8unpolishEP7QWidget
    // invoke: void unpolish(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QStyle8unpolishEP7QWidget(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN6QStyle8unpolishEP12QApplication
    // invoke: void unpolish(class QApplication *)
    var arg0 = args[0].(QApplication).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QStyle8unpolishEP12QApplication(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyle", "unpolish", args)
  }

}

// sliderPositionFromValue(int, int, int, int, _Bool)
func (this *QStyle) sliderPositionFromValue_s(args ...interface{}) () {
  // sliderPositionFromValue(int, int, int, int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStyle23sliderPositionFromValueEiiiib
    // invoke: int sliderPositionFromValue(int, int, int, int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.bool(args[4].(bool))
    if false {fmt.Println(arg4)}
    var ret = C.C_ZN6QStyle23sliderPositionFromValueEiiiib(arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStyle", "sliderPositionFromValue", args)
  }

}

// itemPixmapRect(const class QRect &, int, const class QPixmap &)
func (this *QStyle) itemPixmapRect(args ...interface{}) () {
  // itemPixmapRect(const class QRect &, int, const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle14itemPixmapRectERK5QRectiRK7QPixmap
    // invoke: QRect itemPixmapRect(const class QRect &, int, const class QPixmap &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPixmap).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZNK6QStyle14itemPixmapRectERK5QRectiRK7QPixmap(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStyle", "itemPixmapRect", args)
  }

}

// <= body block end

