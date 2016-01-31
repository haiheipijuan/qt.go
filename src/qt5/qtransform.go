package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtGui/qtransform.h
// dst-file: /src/gui/qtransform.go
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
  // proto:  qreal QTransform::dx();
extern void C_ZNK10QTransform2dxEv(void* qthis); // 2
  // proto:  void QTransform::setMatrix(qreal m11, qreal m12, qreal m13, qreal m21, qreal m22, qreal m23, qreal m31, qreal m32, qreal m33);
extern void C_ZN10QTransform9setMatrixEddddddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4, double arg5, double arg6, double arg7, double arg8); // 4
  // proto:  QRect QTransform::mapRect(const QRect & );
extern void C_ZNK10QTransform7mapRectERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QRectF QTransform::mapRect(const QRectF & );
extern void C_ZNK10QTransform7mapRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  qreal QTransform::m21();
extern void C_ZNK10QTransform3m21Ev(void* qthis); // 2
  // proto:  qreal QTransform::m11();
extern void C_ZNK10QTransform3m11Ev(void* qthis); // 2
  // proto:  qreal QTransform::m13();
extern void C_ZNK10QTransform3m13Ev(void* qthis); // 2
  // proto:  bool QTransform::isInvertible();
extern void C_ZNK10QTransform12isInvertibleEv(void* qthis); // 2
  // proto:  qreal QTransform::m23();
extern void C_ZNK10QTransform3m23Ev(void* qthis); // 2
  // proto:  QTransform & QTransform::scale(qreal sx, qreal sy);
extern void C_ZN10QTransform5scaleEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QPolygon QTransform::mapToPolygon(const QRect & r);
extern void C_ZNK10QTransform12mapToPolygonERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QTransform QTransform::adjoint();
extern void C_ZNK10QTransform7adjointEv(void* qthis); // 4
  // proto:  qreal QTransform::m33();
extern void C_ZNK10QTransform3m33Ev(void* qthis); // 2
  // proto:  qreal QTransform::m32();
extern void C_ZNK10QTransform3m32Ev(void* qthis); // 2
  // proto:  qreal QTransform::m31();
extern void C_ZNK10QTransform3m31Ev(void* qthis); // 2
  // proto:  QTransform & QTransform::translate(qreal dx, qreal dy);
extern void C_ZN10QTransform9translateEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QTransform::TransformationType QTransform::type();
extern void C_ZNK10QTransform4typeEv(void* qthis); // 4
  // proto:  QTransform & QTransform::shear(qreal sh, qreal sv);
extern void C_ZN10QTransform5shearEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  QPoint QTransform::map(const QPoint & p);
extern void C_ZNK10QTransform3mapERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QPainterPath QTransform::map(const QPainterPath & p);
extern void C_ZNK10QTransform3mapERK12QPainterPath(void* qthis, void* arg0); // 4
  // proto:  QLine QTransform::map(const QLine & l);
extern void C_ZNK10QTransform3mapERK5QLine(void* qthis, void* arg0); // 4
  // proto:  QPolygonF QTransform::map(const QPolygonF & a);
extern void C_ZNK10QTransform3mapERK9QPolygonF(void* qthis, void* arg0); // 4
  // proto:  QPointF QTransform::map(const QPointF & p);
extern void C_ZNK10QTransform3mapERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QTransform::map(qreal x, qreal y, qreal * tx, qreal * ty);
extern void C_ZNK10QTransform3mapEddPdS0_(void* qthis, double arg0, double arg1, double* arg2, double* arg3); // 4
  // proto:  QPolygon QTransform::map(const QPolygon & a);
extern void C_ZNK10QTransform3mapERK8QPolygon(void* qthis, void* arg0); // 4
  // proto:  void QTransform::map(int x, int y, int * tx, int * ty);
extern void C_ZNK10QTransform3mapEiiPiS0_(void* qthis, int32_t arg0, int32_t arg1, int32_t* arg2, int32_t* arg3); // 4
  // proto:  QRegion QTransform::map(const QRegion & r);
extern void C_ZNK10QTransform3mapERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  QLineF QTransform::map(const QLineF & l);
extern void C_ZNK10QTransform3mapERK6QLineF(void* qthis, void* arg0); // 4
  // proto:  qreal QTransform::determinant();
extern void C_ZNK10QTransform11determinantEv(void* qthis); // 2
  // proto:  qreal QTransform::dy();
extern void C_ZNK10QTransform2dyEv(void* qthis); // 2
  // proto:  const QMatrix & QTransform::toAffine();
extern void C_ZNK10QTransform8toAffineEv(void* qthis); // 4
  // proto:  bool QTransform::isRotating();
extern void C_ZNK10QTransform10isRotatingEv(void* qthis); // 2
  // proto:  QTransform QTransform::inverted(bool * invertible);
extern void C_ZNK10QTransform8invertedEPb(void* qthis, bool* arg0); // 4
  // proto: static bool QTransform::squareToQuad(const QPolygonF & square, QTransform & result);
extern void C_ZN10QTransform12squareToQuadERK9QPolygonFRS_(void* arg0, void* arg1); // 4
  // proto:  void QTransform::reset();
extern void C_ZN10QTransform5resetEv(void* qthis); // 4
  // proto:  QTransform QTransform::transposed();
extern void C_ZNK10QTransform10transposedEv(void* qthis); // 4
  // proto:  qreal QTransform::det();
extern void C_ZNK10QTransform3detEv(void* qthis); // 2
  // proto:  void QTransform::QTransform();
extern void* C_ZN10QTransformC2Ev(); // 3
  // proto:  void QTransform::QTransform(const QMatrix & mtx);
extern void* C_ZN10QTransformC2ERK7QMatrix(void* arg0); // 3
  // proto:  void QTransform::QTransform(qreal h11, qreal h12, qreal h13, qreal h21, qreal h22, qreal h23, qreal h31, qreal h32, qreal h33);
extern void* C_ZN10QTransformC2Eddddddddd(double arg0, double arg1, double arg2, double arg3, double arg4, double arg5, double arg6, double arg7, double arg8); // 3
  // proto:  void QTransform::QTransform(qreal h11, qreal h12, qreal h21, qreal h22, qreal dx, qreal dy);
extern void* C_ZN10QTransformC2Edddddd(double arg0, double arg1, double arg2, double arg3, double arg4, double arg5); // 3
  // proto:  qreal QTransform::m22();
extern void C_ZNK10QTransform3m22Ev(void* qthis); // 2
  // proto: static QTransform QTransform::fromScale(qreal dx, qreal dy);
extern void C_ZN10QTransform9fromScaleEdd(double arg0, double arg1); // 4
  // proto: static bool QTransform::quadToQuad(const QPolygonF & one, const QPolygonF & two, QTransform & result);
extern void C_ZN10QTransform10quadToQuadERK9QPolygonFS2_RS_(void* arg0, void* arg1, void* arg2); // 4
  // proto:  qreal QTransform::m12();
extern void C_ZNK10QTransform3m12Ev(void* qthis); // 2
  // proto:  bool QTransform::isIdentity();
extern void C_ZNK10QTransform10isIdentityEv(void* qthis); // 2
  // proto:  bool QTransform::isScaling();
extern void C_ZNK10QTransform9isScalingEv(void* qthis); // 2
  // proto: static bool QTransform::quadToSquare(const QPolygonF & quad, QTransform & result);
extern void C_ZN10QTransform12quadToSquareERK9QPolygonFRS_(void* arg0, void* arg1); // 4
  // proto:  bool QTransform::isAffine();
extern void C_ZNK10QTransform8isAffineEv(void* qthis); // 2
  // proto:  bool QTransform::isTranslating();
extern void C_ZNK10QTransform13isTranslatingEv(void* qthis); // 2
  // proto: static QTransform QTransform::fromTranslate(qreal dx, qreal dy);
extern void C_ZN10QTransform13fromTranslateEdd(double arg0, double arg1); // 4
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

// class sizeof(QTransform)=88
type QTransform struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// dx()
func (this *QTransform) dx(args ...interface{}) () {
  // dx()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform2dxEv
    // invoke: qreal dx()
    var ret = C.C_ZNK10QTransform2dxEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "dx", args)
  }

}

// setMatrix(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
func (this *QTransform) setMatrix(args ...interface{}) () {
  // setMatrix(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][5] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][6] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][7] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][8] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform9setMatrixEddddddddd
    // invoke: void setMatrix(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(args[4].(float64))
    if false {fmt.Println(arg4)}
    var arg5 = C.double(args[5].(float64))
    if false {fmt.Println(arg5)}
    var arg6 = C.double(args[6].(float64))
    if false {fmt.Println(arg6)}
    var arg7 = C.double(args[7].(float64))
    if false {fmt.Println(arg7)}
    var arg8 = C.double(args[8].(float64))
    if false {fmt.Println(arg8)}
    C.C_ZN10QTransform9setMatrixEddddddddd(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
  default:
    qtrt.ErrorResolve("QTransform", "setMatrix", args)
  }

}

// mapRect(const class QRect &)
func (this *QTransform) mapRect(args ...interface{}) () {
  // mapRect(const class QRect &)
  // mapRect(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform7mapRectERK5QRect
    // invoke: QRect mapRect(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTransform7mapRectERK5QRect(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK10QTransform7mapRectERK6QRectF
    // invoke: QRectF mapRect(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTransform7mapRectERK6QRectF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "mapRect", args)
  }

}

// m21()
func (this *QTransform) m21(args ...interface{}) () {
  // m21()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m21Ev
    // invoke: qreal m21()
    var ret = C.C_ZNK10QTransform3m21Ev(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "m21", args)
  }

}

// m11()
func (this *QTransform) m11(args ...interface{}) () {
  // m11()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m11Ev
    // invoke: qreal m11()
    var ret = C.C_ZNK10QTransform3m11Ev(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "m11", args)
  }

}

// m13()
func (this *QTransform) m13(args ...interface{}) () {
  // m13()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m13Ev
    // invoke: qreal m13()
    var ret = C.C_ZNK10QTransform3m13Ev(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "m13", args)
  }

}

// isInvertible()
func (this *QTransform) isInvertible(args ...interface{}) () {
  // isInvertible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform12isInvertibleEv
    // invoke: bool isInvertible()
    var ret = C.C_ZNK10QTransform12isInvertibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "isInvertible", args)
  }

}

// m23()
func (this *QTransform) m23(args ...interface{}) () {
  // m23()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m23Ev
    // invoke: qreal m23()
    var ret = C.C_ZNK10QTransform3m23Ev(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "m23", args)
  }

}

// scale(qreal, qreal)
func (this *QTransform) scale(args ...interface{}) () {
  // scale(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform5scaleEdd
    // invoke: QTransform & scale(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN10QTransform5scaleEdd(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "scale", args)
  }

}

// mapToPolygon(const class QRect &)
func (this *QTransform) mapToPolygon(args ...interface{}) () {
  // mapToPolygon(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform12mapToPolygonERK5QRect
    // invoke: QPolygon mapToPolygon(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTransform12mapToPolygonERK5QRect(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "mapToPolygon", args)
  }

}

// adjoint()
func (this *QTransform) adjoint(args ...interface{}) () {
  // adjoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform7adjointEv
    // invoke: QTransform adjoint()
    var ret = C.C_ZNK10QTransform7adjointEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "adjoint", args)
  }

}

// m33()
func (this *QTransform) m33(args ...interface{}) () {
  // m33()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m33Ev
    // invoke: qreal m33()
    var ret = C.C_ZNK10QTransform3m33Ev(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "m33", args)
  }

}

// m32()
func (this *QTransform) m32(args ...interface{}) () {
  // m32()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m32Ev
    // invoke: qreal m32()
    var ret = C.C_ZNK10QTransform3m32Ev(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "m32", args)
  }

}

// m31()
func (this *QTransform) m31(args ...interface{}) () {
  // m31()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m31Ev
    // invoke: qreal m31()
    var ret = C.C_ZNK10QTransform3m31Ev(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "m31", args)
  }

}

// translate(qreal, qreal)
func (this *QTransform) translate(args ...interface{}) () {
  // translate(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform9translateEdd
    // invoke: QTransform & translate(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN10QTransform9translateEdd(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "translate", args)
  }

}

// type()
func (this *QTransform) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform4typeEv
    // invoke: QTransform::TransformationType type()
    C.C_ZNK10QTransform4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTransform", "type", args)
  }

}

// shear(qreal, qreal)
func (this *QTransform) shear(args ...interface{}) () {
  // shear(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform5shearEdd
    // invoke: QTransform & shear(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN10QTransform5shearEdd(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "shear", args)
  }

}

// map(const class QPoint &)
func (this *QTransform) map_(args ...interface{}) () {
  // map(const class QPoint &)
  // map(const class QPainterPath &)
  // map(const class QLine &)
  // map(const class QPolygonF &)
  // map(const class QPointF &)
  // map(qreal, qreal, qreal *, qreal *)
  // map(const class QPolygon &)
  // map(int, int, int *, int *)
  // map(const class QRegion &)
  // map(const class QLineF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLine{}) // "const QLine &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[5][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[5][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[5][3] = qtrt.DoubleTy(true) // "qreal *"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = qtrt.Int32Ty(true) // "int *"
  vtys[7][3] = qtrt.Int32Ty(true) // "int *"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3mapERK6QPoint
    // invoke: QPoint map(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTransform3mapERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK10QTransform3mapERK12QPainterPath
    // invoke: QPainterPath map(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTransform3mapERK12QPainterPath(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZNK10QTransform3mapERK5QLine
    // invoke: QLine map(const class QLine &)
    var arg0 = args[0].(QLine).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTransform3mapERK5QLine(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 3:
    // invoke: _ZNK10QTransform3mapERK9QPolygonF
    // invoke: QPolygonF map(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTransform3mapERK9QPolygonF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 4:
    // invoke: _ZNK10QTransform3mapERK7QPointF
    // invoke: QPointF map(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTransform3mapERK7QPointF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 5:
    // invoke: _ZNK10QTransform3mapEddPdS0_
    // invoke: void map(qreal, qreal, qreal *, qreal *)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK10QTransform3mapEddPdS0_(this.qclsinst, arg0, arg1, arg2, arg3)
  case 6:
    // invoke: _ZNK10QTransform3mapERK8QPolygon
    // invoke: QPolygon map(const class QPolygon &)
    var arg0 = args[0].(QPolygon).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTransform3mapERK8QPolygon(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 7:
    // invoke: _ZNK10QTransform3mapEiiPiS0_
    // invoke: void map(int, int, int *, int *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK10QTransform3mapEiiPiS0_(this.qclsinst, arg0, arg1, arg2, arg3)
  case 8:
    // invoke: _ZNK10QTransform3mapERK7QRegion
    // invoke: QRegion map(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTransform3mapERK7QRegion(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 9:
    // invoke: _ZNK10QTransform3mapERK6QLineF
    // invoke: QLineF map(const class QLineF &)
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTransform3mapERK6QLineF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "map", args)
  }

}

// determinant()
func (this *QTransform) determinant(args ...interface{}) () {
  // determinant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform11determinantEv
    // invoke: qreal determinant()
    var ret = C.C_ZNK10QTransform11determinantEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "determinant", args)
  }

}

// dy()
func (this *QTransform) dy(args ...interface{}) () {
  // dy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform2dyEv
    // invoke: qreal dy()
    var ret = C.C_ZNK10QTransform2dyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "dy", args)
  }

}

// toAffine()
func (this *QTransform) toAffine(args ...interface{}) () {
  // toAffine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform8toAffineEv
    // invoke: const QMatrix & toAffine()
    var ret = C.C_ZNK10QTransform8toAffineEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "toAffine", args)
  }

}

// isRotating()
func (this *QTransform) isRotating(args ...interface{}) () {
  // isRotating()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform10isRotatingEv
    // invoke: bool isRotating()
    var ret = C.C_ZNK10QTransform10isRotatingEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "isRotating", args)
  }

}

// inverted(_Bool *)
func (this *QTransform) inverted(args ...interface{}) () {
  // inverted(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform8invertedEPb
    // invoke: QTransform inverted(_Bool *)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTransform8invertedEPb(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "inverted", args)
  }

}

// squareToQuad(const class QPolygonF &, class QTransform &)
func (this *QTransform) squareToQuad_s(args ...interface{}) () {
  // squareToQuad(const class QPolygonF &, class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[0][1] = reflect.TypeOf(QTransform{}) // "QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform12squareToQuadERK9QPolygonFRS_
    // invoke: bool squareToQuad(const class QPolygonF &, class QTransform &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTransform).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN10QTransform12squareToQuadERK9QPolygonFRS_(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "squareToQuad", args)
  }

}

// reset()
func (this *QTransform) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform5resetEv
    // invoke: void reset()
    C.C_ZN10QTransform5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTransform", "reset", args)
  }

}

// transposed()
func (this *QTransform) transposed(args ...interface{}) () {
  // transposed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform10transposedEv
    // invoke: QTransform transposed()
    var ret = C.C_ZNK10QTransform10transposedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "transposed", args)
  }

}

// det()
func (this *QTransform) det(args ...interface{}) () {
  // det()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3detEv
    // invoke: qreal det()
    var ret = C.C_ZNK10QTransform3detEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "det", args)
  }

}

// QTransform()
func NewQTransform(args ...interface{}) *QTransform {
  // QTransform()
  // QTransform(const class QMatrix &)
  // QTransform(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
  // QTransform(qreal, qreal, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][5] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][6] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][7] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][8] = qtrt.DoubleTy(false) // "qreal"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][4] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][5] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransformC1Ev
    // invoke: void QTransform()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTransformC2Ev()
    return &QTransform{qclsinst:qthis}
  case 1:
    // invoke: _ZN10QTransformC1ERK7QMatrix
    // invoke: void QTransform(const class QMatrix &)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTransformC2ERK7QMatrix(arg0)
    return &QTransform{qclsinst:qthis}
  case 2:
    // invoke: _ZN10QTransformC1Eddddddddd
    // invoke: void QTransform(qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(args[4].(float64))
    if false {fmt.Println(arg4)}
    var arg5 = C.double(args[5].(float64))
    if false {fmt.Println(arg5)}
    var arg6 = C.double(args[6].(float64))
    if false {fmt.Println(arg6)}
    var arg7 = C.double(args[7].(float64))
    if false {fmt.Println(arg7)}
    var arg8 = C.double(args[8].(float64))
    if false {fmt.Println(arg8)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTransformC2Eddddddddd(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return &QTransform{qclsinst:qthis}
  case 3:
    // invoke: _ZN10QTransformC1Edddddd
    // invoke: void QTransform(qreal, qreal, qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(args[4].(float64))
    if false {fmt.Println(arg4)}
    var arg5 = C.double(args[5].(float64))
    if false {fmt.Println(arg5)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTransformC2Edddddd(arg0, arg1, arg2, arg3, arg4, arg5)
    return &QTransform{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTransform", "QTransform", args)
  }

  return nil // QTransform{qclsinst:qthis}
}

// m22()
func (this *QTransform) m22(args ...interface{}) () {
  // m22()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m22Ev
    // invoke: qreal m22()
    var ret = C.C_ZNK10QTransform3m22Ev(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "m22", args)
  }

}

// fromScale(qreal, qreal)
func (this *QTransform) fromScale_s(args ...interface{}) () {
  // fromScale(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform9fromScaleEdd
    // invoke: QTransform fromScale(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN10QTransform9fromScaleEdd(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "fromScale", args)
  }

}

// quadToQuad(const class QPolygonF &, const class QPolygonF &, class QTransform &)
func (this *QTransform) quadToQuad_s(args ...interface{}) () {
  // quadToQuad(const class QPolygonF &, const class QPolygonF &, class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[0][1] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[0][2] = reflect.TypeOf(QTransform{}) // "QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform10quadToQuadERK9QPolygonFS2_RS_
    // invoke: bool quadToQuad(const class QPolygonF &, const class QPolygonF &, class QTransform &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPolygonF).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QTransform).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN10QTransform10quadToQuadERK9QPolygonFS2_RS_(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "quadToQuad", args)
  }

}

// m12()
func (this *QTransform) m12(args ...interface{}) () {
  // m12()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform3m12Ev
    // invoke: qreal m12()
    var ret = C.C_ZNK10QTransform3m12Ev(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "m12", args)
  }

}

// isIdentity()
func (this *QTransform) isIdentity(args ...interface{}) () {
  // isIdentity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform10isIdentityEv
    // invoke: bool isIdentity()
    var ret = C.C_ZNK10QTransform10isIdentityEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "isIdentity", args)
  }

}

// isScaling()
func (this *QTransform) isScaling(args ...interface{}) () {
  // isScaling()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform9isScalingEv
    // invoke: bool isScaling()
    var ret = C.C_ZNK10QTransform9isScalingEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "isScaling", args)
  }

}

// quadToSquare(const class QPolygonF &, class QTransform &)
func (this *QTransform) quadToSquare_s(args ...interface{}) () {
  // quadToSquare(const class QPolygonF &, class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[0][1] = reflect.TypeOf(QTransform{}) // "QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform12quadToSquareERK9QPolygonFRS_
    // invoke: bool quadToSquare(const class QPolygonF &, class QTransform &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTransform).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN10QTransform12quadToSquareERK9QPolygonFRS_(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "quadToSquare", args)
  }

}

// isAffine()
func (this *QTransform) isAffine(args ...interface{}) () {
  // isAffine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform8isAffineEv
    // invoke: bool isAffine()
    var ret = C.C_ZNK10QTransform8isAffineEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "isAffine", args)
  }

}

// isTranslating()
func (this *QTransform) isTranslating(args ...interface{}) () {
  // isTranslating()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTransform13isTranslatingEv
    // invoke: bool isTranslating()
    var ret = C.C_ZNK10QTransform13isTranslatingEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "isTranslating", args)
  }

}

// fromTranslate(qreal, qreal)
func (this *QTransform) fromTranslate_s(args ...interface{}) () {
  // fromTranslate(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTransform13fromTranslateEdd
    // invoke: QTransform fromTranslate(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN10QTransform13fromTranslateEdd(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTransform", "fromTranslate", args)
  }

}

// <= body block end

