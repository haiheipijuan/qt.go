package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtGui/qgenericplugin.h
// dst-file: /src/gui/qgenericplugin.go
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
  // proto:  const QMetaObject * QGenericPlugin::metaObject();
extern void C_ZNK14QGenericPlugin10metaObjectEv(void* qthis); // 4
  // proto:  void QGenericPlugin::QGenericPlugin(QObject * parent);
extern void* C_ZN14QGenericPluginC2EP7QObject(void* arg0); // 3
  // proto:  void QGenericPlugin::~QGenericPlugin();
extern void C_ZN14QGenericPluginD2Ev(void* qthis); // 4
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

// class sizeof(QGenericPlugin)=1
type QGenericPlugin struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QGenericPlugin) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGenericPlugin10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK14QGenericPlugin10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGenericPlugin", "metaObject", args)
  }

}

// QGenericPlugin(class QObject *)
func NewQGenericPlugin(args ...interface{}) *QGenericPlugin {
  // QGenericPlugin(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGenericPluginC1EP7QObject
    // invoke: void QGenericPlugin(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QGenericPluginC2EP7QObject(arg0)
    return &QGenericPlugin{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGenericPlugin", "QGenericPlugin", args)
  }

  return nil // QGenericPlugin{qclsinst:qthis}
}

// ~QGenericPlugin()
func (this *QGenericPlugin) FreeQGenericPlugin(args ...interface{}) () {
  // ~QGenericPlugin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGenericPluginD0Ev
    // invoke: void ~QGenericPlugin()
    C.C_ZN14QGenericPluginD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGenericPlugin", "~QGenericPlugin", args)
  }

}

// <= body block end

