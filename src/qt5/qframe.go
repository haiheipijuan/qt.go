package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtWidgets/qframe.h
// dst-file: /src/widgets/qframe.go
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
  // proto:  int QFrame::frameStyle();
extern void C_ZNK6QFrame10frameStyleEv(void* qthis); // 4
  // proto:  void QFrame::~QFrame();
extern void C_ZN6QFrameD2Ev(void* qthis); // 4
  // proto:  int QFrame::frameWidth();
extern void C_ZNK6QFrame10frameWidthEv(void* qthis); // 4
  // proto:  void QFrame::setLineWidth(int );
extern void C_ZN6QFrame12setLineWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QFrame::metaObject();
extern void C_ZNK6QFrame10metaObjectEv(void* qthis); // 4
  // proto:  void QFrame::setFrameStyle(int );
extern void C_ZN6QFrame13setFrameStyleEi(void* qthis, int32_t arg0); // 4
  // proto:  QRect QFrame::frameRect();
extern void C_ZNK6QFrame9frameRectEv(void* qthis); // 4
  // proto:  QFrame::Shadow QFrame::frameShadow();
extern void C_ZNK6QFrame11frameShadowEv(void* qthis); // 4
  // proto:  void QFrame::setFrameRect(const QRect & );
extern void C_ZN6QFrame12setFrameRectERK5QRect(void* qthis, void* arg0); // 4
  // proto:  int QFrame::midLineWidth();
extern void C_ZNK6QFrame12midLineWidthEv(void* qthis); // 4
  // proto:  QSize QFrame::sizeHint();
extern void C_ZNK6QFrame8sizeHintEv(void* qthis); // 4
  // proto:  QFrame::Shape QFrame::frameShape();
extern void C_ZNK6QFrame10frameShapeEv(void* qthis); // 4
  // proto:  void QFrame::setMidLineWidth(int );
extern void C_ZN6QFrame15setMidLineWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  int QFrame::lineWidth();
extern void C_ZNK6QFrame9lineWidthEv(void* qthis); // 4
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

// class sizeof(QFrame)=1
type QFrame struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
}

// frameStyle()
func (this *QFrame) frameStyle(args ...interface{}) () {
  // frameStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame10frameStyleEv
    // invoke: int frameStyle()
    var ret = C.C_ZNK6QFrame10frameStyleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFrame", "frameStyle", args)
  }

}

// ~QFrame()
func (this *QFrame) FreeQFrame(args ...interface{}) () {
  // ~QFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QFrameD0Ev
    // invoke: void ~QFrame()
    C.C_ZN6QFrameD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFrame", "~QFrame", args)
  }

}

// frameWidth()
func (this *QFrame) frameWidth(args ...interface{}) () {
  // frameWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame10frameWidthEv
    // invoke: int frameWidth()
    var ret = C.C_ZNK6QFrame10frameWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFrame", "frameWidth", args)
  }

}

// setLineWidth(int)
func (this *QFrame) setLineWidth(args ...interface{}) () {
  // setLineWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QFrame12setLineWidthEi
    // invoke: void setLineWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QFrame12setLineWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFrame", "setLineWidth", args)
  }

}

// metaObject()
func (this *QFrame) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK6QFrame10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFrame", "metaObject", args)
  }

}

// setFrameStyle(int)
func (this *QFrame) setFrameStyle(args ...interface{}) () {
  // setFrameStyle(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QFrame13setFrameStyleEi
    // invoke: void setFrameStyle(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QFrame13setFrameStyleEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFrame", "setFrameStyle", args)
  }

}

// frameRect()
func (this *QFrame) frameRect(args ...interface{}) () {
  // frameRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame9frameRectEv
    // invoke: QRect frameRect()
    var ret = C.C_ZNK6QFrame9frameRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFrame", "frameRect", args)
  }

}

// frameShadow()
func (this *QFrame) frameShadow(args ...interface{}) () {
  // frameShadow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame11frameShadowEv
    // invoke: QFrame::Shadow frameShadow()
    C.C_ZNK6QFrame11frameShadowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFrame", "frameShadow", args)
  }

}

// setFrameRect(const class QRect &)
func (this *QFrame) setFrameRect(args ...interface{}) () {
  // setFrameRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QFrame12setFrameRectERK5QRect
    // invoke: void setFrameRect(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QFrame12setFrameRectERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFrame", "setFrameRect", args)
  }

}

// midLineWidth()
func (this *QFrame) midLineWidth(args ...interface{}) () {
  // midLineWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame12midLineWidthEv
    // invoke: int midLineWidth()
    var ret = C.C_ZNK6QFrame12midLineWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFrame", "midLineWidth", args)
  }

}

// sizeHint()
func (this *QFrame) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame8sizeHintEv
    // invoke: QSize sizeHint()
    var ret = C.C_ZNK6QFrame8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFrame", "sizeHint", args)
  }

}

// frameShape()
func (this *QFrame) frameShape(args ...interface{}) () {
  // frameShape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame10frameShapeEv
    // invoke: QFrame::Shape frameShape()
    C.C_ZNK6QFrame10frameShapeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFrame", "frameShape", args)
  }

}

// setMidLineWidth(int)
func (this *QFrame) setMidLineWidth(args ...interface{}) () {
  // setMidLineWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QFrame15setMidLineWidthEi
    // invoke: void setMidLineWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QFrame15setMidLineWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFrame", "setMidLineWidth", args)
  }

}

// lineWidth()
func (this *QFrame) lineWidth(args ...interface{}) () {
  // lineWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame9lineWidthEv
    // invoke: int lineWidth()
    var ret = C.C_ZNK6QFrame9lineWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFrame", "lineWidth", args)
  }

}

// <= body block end

