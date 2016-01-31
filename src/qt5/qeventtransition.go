package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qeventtransition.h
// dst-file: /src/core/qeventtransition.go
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
  // proto:  const QMetaObject * QEventTransition::metaObject();
extern void C_ZNK16QEventTransition10metaObjectEv(void* qthis); // 4
  // proto:  void QEventTransition::setEventSource(QObject * object);
extern void C_ZN16QEventTransition14setEventSourceEP7QObject(void* qthis, void* arg0); // 4
  // proto:  void QEventTransition::QEventTransition(QState * sourceState);
extern void* C_ZN16QEventTransitionC2EP6QState(void* arg0); // 3
  // proto:  QObject * QEventTransition::eventSource();
extern void C_ZNK16QEventTransition11eventSourceEv(void* qthis); // 4
  // proto:  QEvent::Type QEventTransition::eventType();
extern void C_ZNK16QEventTransition9eventTypeEv(void* qthis); // 4
  // proto:  void QEventTransition::~QEventTransition();
extern void C_ZN16QEventTransitionD2Ev(void* qthis); // 4
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

// class sizeof(QEventTransition)=1
type QEventTransition struct {
  /*qbase*/ QAbstractTransition;
  qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QEventTransition) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QEventTransition10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK16QEventTransition10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventTransition", "metaObject", args)
  }

}

// setEventSource(class QObject *)
func (this *QEventTransition) setEventSource(args ...interface{}) () {
  // setEventSource(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QEventTransition14setEventSourceEP7QObject
    // invoke: void setEventSource(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QEventTransition14setEventSourceEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEventTransition", "setEventSource", args)
  }

}

// QEventTransition(class QState *)
func NewQEventTransition(args ...interface{}) *QEventTransition {
  // QEventTransition(class QState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QState{}) // "QState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QEventTransitionC1EP6QState
    // invoke: void QEventTransition(class QState *)
    var arg0 = args[0].(QState).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QEventTransitionC2EP6QState(arg0)
    return &QEventTransition{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QEventTransition", "QEventTransition", args)
  }

  return nil // QEventTransition{qclsinst:qthis}
}

// eventSource()
func (this *QEventTransition) eventSource(args ...interface{}) () {
  // eventSource()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QEventTransition11eventSourceEv
    // invoke: QObject * eventSource()
    var ret = C.C_ZNK16QEventTransition11eventSourceEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QEventTransition", "eventSource", args)
  }

}

// eventType()
func (this *QEventTransition) eventType(args ...interface{}) () {
  // eventType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QEventTransition9eventTypeEv
    // invoke: QEvent::Type eventType()
    C.C_ZNK16QEventTransition9eventTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventTransition", "eventType", args)
  }

}

// ~QEventTransition()
func (this *QEventTransition) FreeQEventTransition(args ...interface{}) () {
  // ~QEventTransition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QEventTransitionD0Ev
    // invoke: void ~QEventTransition()
    C.C_ZN16QEventTransitionD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEventTransition", "~QEventTransition", args)
  }

}

// <= body block end

