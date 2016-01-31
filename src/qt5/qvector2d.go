package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtGui/qvector2d.h
// dst-file: /src/gui/qvector2d.go
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
  // proto:  QVector4D QVector2D::toVector4D();
extern void C_ZNK9QVector2D10toVector4DEv(void* qthis); // 4
  // proto:  void QVector2D::normalize();
extern void C_ZN9QVector2D9normalizeEv(void* qthis); // 4
  // proto: static float QVector2D::dotProduct(const QVector2D & v1, const QVector2D & v2);
extern void C_ZN9QVector2D10dotProductERKS_S1_(void* arg0, void* arg1); // 4
  // proto:  void QVector2D::QVector2D(const QPoint & point);
extern void* C_ZN9QVector2DC2ERK6QPoint(void* arg0); // 1
  // proto:  void QVector2D::QVector2D(float xpos, float ypos);
extern void* C_ZN9QVector2DC2Eff(float arg0, float arg1); // 1
  // proto:  void QVector2D::QVector2D(const QVector4D & vector);
extern void* C_ZN9QVector2DC2ERK9QVector4D(void* arg0); // 3
  // proto:  void QVector2D::QVector2D();
extern void* C_ZN9QVector2DC2Ev(); // 1
  // proto:  void QVector2D::QVector2D(const QVector3D & vector);
extern void* C_ZN9QVector2DC2ERK9QVector3D(void* arg0); // 3
  // proto:  void QVector2D::QVector2D(const QPointF & point);
extern void* C_ZN9QVector2DC2ERK7QPointF(void* arg0); // 1
  // proto:  QPointF QVector2D::toPointF();
extern void C_ZNK9QVector2D8toPointFEv(void* qthis); // 2
  // proto:  QVector3D QVector2D::toVector3D();
extern void C_ZNK9QVector2D10toVector3DEv(void* qthis); // 4
  // proto:  float QVector2D::lengthSquared();
extern void C_ZNK9QVector2D13lengthSquaredEv(void* qthis); // 4
  // proto:  QVector2D QVector2D::normalized();
extern void C_ZNK9QVector2D10normalizedEv(void* qthis); // 4
  // proto:  void QVector2D::setX(float x);
extern void C_ZN9QVector2D4setXEf(void* qthis, float arg0); // 2
  // proto:  void QVector2D::setY(float y);
extern void C_ZN9QVector2D4setYEf(void* qthis, float arg0); // 2
  // proto:  QPoint QVector2D::toPoint();
extern void C_ZNK9QVector2D7toPointEv(void* qthis); // 2
  // proto:  float QVector2D::distanceToLine(const QVector2D & point, const QVector2D & direction);
extern void C_ZNK9QVector2D14distanceToLineERKS_S1_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QVector2D::isNull();
extern void C_ZNK9QVector2D6isNullEv(void* qthis); // 2
  // proto:  float QVector2D::length();
extern void C_ZNK9QVector2D6lengthEv(void* qthis); // 4
  // proto:  float QVector2D::distanceToPoint(const QVector2D & point);
extern void C_ZNK9QVector2D15distanceToPointERKS_(void* qthis, void* arg0); // 4
  // proto:  float QVector2D::y();
extern void C_ZNK9QVector2D1yEv(void* qthis); // 2
  // proto:  float QVector2D::x();
extern void C_ZNK9QVector2D1xEv(void* qthis); // 2
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

// class sizeof(QVector2D)=8
type QVector2D struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// toVector4D()
func (this *QVector2D) toVector4D(args ...interface{}) () {
  // toVector4D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D10toVector4DEv
    // invoke: QVector4D toVector4D()
    var ret = C.C_ZNK9QVector2D10toVector4DEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QVector2D", "toVector4D", args)
  }

}

// normalize()
func (this *QVector2D) normalize(args ...interface{}) () {
  // normalize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D9normalizeEv
    // invoke: void normalize()
    C.C_ZN9QVector2D9normalizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector2D", "normalize", args)
  }

}

// dotProduct(const class QVector2D &, const class QVector2D &)
func (this *QVector2D) dotProduct_s(args ...interface{}) () {
  // dotProduct(const class QVector2D &, const class QVector2D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[0][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D10dotProductERKS_S1_
    // invoke: float dotProduct(const class QVector2D &, const class QVector2D &)
    var arg0 = args[0].(QVector2D).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector2D).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN9QVector2D10dotProductERKS_S1_(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QVector2D", "dotProduct", args)
  }

}

// QVector2D(const class QPoint &)
func NewQVector2D(args ...interface{}) *QVector2D {
  // QVector2D(const class QPoint &)
  // QVector2D(float, float)
  // QVector2D(const class QVector4D &)
  // QVector2D()
  // QVector2D(const class QVector3D &)
  // QVector2D(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.FloatTy(false) // "float"
  vtys[1][1] = qtrt.FloatTy(false) // "float"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2DC1ERK6QPoint
    // invoke: void QVector2D(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector2DC2ERK6QPoint(arg0)
    return &QVector2D{qclsinst:qthis}
  case 1:
    // invoke: _ZN9QVector2DC1Eff
    // invoke: void QVector2D(float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector2DC2Eff(arg0, arg1)
    return &QVector2D{qclsinst:qthis}
  case 2:
    // invoke: _ZN9QVector2DC1ERK9QVector4D
    // invoke: void QVector2D(const class QVector4D &)
    var arg0 = args[0].(QVector4D).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector2DC2ERK9QVector4D(arg0)
    return &QVector2D{qclsinst:qthis}
  case 3:
    // invoke: _ZN9QVector2DC1Ev
    // invoke: void QVector2D()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector2DC2Ev()
    return &QVector2D{qclsinst:qthis}
  case 4:
    // invoke: _ZN9QVector2DC1ERK9QVector3D
    // invoke: void QVector2D(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector2DC2ERK9QVector3D(arg0)
    return &QVector2D{qclsinst:qthis}
  case 5:
    // invoke: _ZN9QVector2DC1ERK7QPointF
    // invoke: void QVector2D(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QVector2DC2ERK7QPointF(arg0)
    return &QVector2D{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QVector2D", "QVector2D", args)
  }

  return nil // QVector2D{qclsinst:qthis}
}

// toPointF()
func (this *QVector2D) toPointF(args ...interface{}) () {
  // toPointF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D8toPointFEv
    // invoke: QPointF toPointF()
    var ret = C.C_ZNK9QVector2D8toPointFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QVector2D", "toPointF", args)
  }

}

// toVector3D()
func (this *QVector2D) toVector3D(args ...interface{}) () {
  // toVector3D()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D10toVector3DEv
    // invoke: QVector3D toVector3D()
    var ret = C.C_ZNK9QVector2D10toVector3DEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QVector2D", "toVector3D", args)
  }

}

// lengthSquared()
func (this *QVector2D) lengthSquared(args ...interface{}) () {
  // lengthSquared()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D13lengthSquaredEv
    // invoke: float lengthSquared()
    var ret = C.C_ZNK9QVector2D13lengthSquaredEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QVector2D", "lengthSquared", args)
  }

}

// normalized()
func (this *QVector2D) normalized(args ...interface{}) () {
  // normalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D10normalizedEv
    // invoke: QVector2D normalized()
    var ret = C.C_ZNK9QVector2D10normalizedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QVector2D", "normalized", args)
  }

}

// setX(float)
func (this *QVector2D) setX(args ...interface{}) () {
  // setX(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D4setXEf
    // invoke: void setX(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QVector2D4setXEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector2D", "setX", args)
  }

}

// setY(float)
func (this *QVector2D) setY(args ...interface{}) () {
  // setY(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QVector2D4setYEf
    // invoke: void setY(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QVector2D4setYEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVector2D", "setY", args)
  }

}

// toPoint()
func (this *QVector2D) toPoint(args ...interface{}) () {
  // toPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D7toPointEv
    // invoke: QPoint toPoint()
    var ret = C.C_ZNK9QVector2D7toPointEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QVector2D", "toPoint", args)
  }

}

// distanceToLine(const class QVector2D &, const class QVector2D &)
func (this *QVector2D) distanceToLine(args ...interface{}) () {
  // distanceToLine(const class QVector2D &, const class QVector2D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[0][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D14distanceToLineERKS_S1_
    // invoke: float distanceToLine(const class QVector2D &, const class QVector2D &)
    var arg0 = args[0].(QVector2D).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector2D).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK9QVector2D14distanceToLineERKS_S1_(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QVector2D", "distanceToLine", args)
  }

}

// isNull()
func (this *QVector2D) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D6isNullEv
    // invoke: bool isNull()
    var ret = C.C_ZNK9QVector2D6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QVector2D", "isNull", args)
  }

}

// length()
func (this *QVector2D) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D6lengthEv
    // invoke: float length()
    var ret = C.C_ZNK9QVector2D6lengthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QVector2D", "length", args)
  }

}

// distanceToPoint(const class QVector2D &)
func (this *QVector2D) distanceToPoint(args ...interface{}) () {
  // distanceToPoint(const class QVector2D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D15distanceToPointERKS_
    // invoke: float distanceToPoint(const class QVector2D &)
    var arg0 = args[0].(QVector2D).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK9QVector2D15distanceToPointERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QVector2D", "distanceToPoint", args)
  }

}

// y()
func (this *QVector2D) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D1yEv
    // invoke: float y()
    C.C_ZNK9QVector2D1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector2D", "y", args)
  }

}

// x()
func (this *QVector2D) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QVector2D1xEv
    // invoke: float x()
    C.C_ZNK9QVector2D1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVector2D", "x", args)
  }

}

// <= body block end

