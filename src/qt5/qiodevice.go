package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qiodevice.h
// dst-file: /src/core/qiodevice.go
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
  // proto:  qint64 QIODevice::bytesAvailable();
extern void C_ZNK9QIODevice14bytesAvailableEv(void* qthis); // 4
  // proto:  qint64 QIODevice::pos();
extern void C_ZNK9QIODevice3posEv(void* qthis); // 4
  // proto:  qint64 QIODevice::bytesToWrite();
extern void C_ZNK9QIODevice12bytesToWriteEv(void* qthis); // 4
  // proto:  void QIODevice::close();
extern void C_ZN9QIODevice5closeEv(void* qthis); // 4
  // proto:  bool QIODevice::seek(qint64 pos);
extern void C_ZN9QIODevice4seekEx(void* qthis, int64_t arg0); // 4
  // proto:  bool QIODevice::isSequential();
extern void C_ZNK9QIODevice12isSequentialEv(void* qthis); // 4
  // proto:  void QIODevice::ungetChar(char c);
extern void C_ZN9QIODevice9ungetCharEc(void* qthis, unsigned char arg0); // 4
  // proto:  qint64 QIODevice::size();
extern void C_ZNK9QIODevice4sizeEv(void* qthis); // 4
  // proto:  qint64 QIODevice::write(const char * data, qint64 len);
extern void C_ZN9QIODevice5writeEPKcx(void* qthis, unsigned char* arg0, int64_t arg1); // 4
  // proto:  qint64 QIODevice::write(const char * data);
extern void C_ZN9QIODevice5writeEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  qint64 QIODevice::write(const QByteArray & data);
extern void C_ZN9QIODevice5writeERK10QByteArray(void* qthis, void* arg0); // 2
  // proto:  bool QIODevice::canReadLine();
extern void C_ZNK9QIODevice11canReadLineEv(void* qthis); // 4
  // proto:  bool QIODevice::waitForBytesWritten(int msecs);
extern void C_ZN9QIODevice19waitForBytesWrittenEi(void* qthis, int32_t arg0); // 4
  // proto:  qint64 QIODevice::read(char * data, qint64 maxlen);
extern void C_ZN9QIODevice4readEPcx(void* qthis, unsigned char* arg0, int64_t arg1); // 4
  // proto:  QByteArray QIODevice::read(qint64 maxlen);
extern void C_ZN9QIODevice4readEx(void* qthis, int64_t arg0); // 4
  // proto:  qint64 QIODevice::peek(char * data, qint64 maxlen);
extern void C_ZN9QIODevice4peekEPcx(void* qthis, unsigned char* arg0, int64_t arg1); // 4
  // proto:  QByteArray QIODevice::peek(qint64 maxlen);
extern void C_ZN9QIODevice4peekEx(void* qthis, int64_t arg0); // 4
  // proto:  bool QIODevice::putChar(char c);
extern void C_ZN9QIODevice7putCharEc(void* qthis, unsigned char arg0); // 4
  // proto:  QString QIODevice::errorString();
extern void C_ZNK9QIODevice11errorStringEv(void* qthis); // 4
  // proto:  void QIODevice::~QIODevice();
extern void C_ZN9QIODeviceD2Ev(void* qthis); // 4
  // proto:  QByteArray QIODevice::readAll();
extern void C_ZN9QIODevice7readAllEv(void* qthis); // 4
  // proto:  void QIODevice::setTextModeEnabled(bool enabled);
extern void C_ZN9QIODevice18setTextModeEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QIODevice::QIODevice(QObject * parent);
extern void* C_ZN9QIODeviceC2EP7QObject(void* arg0); // 3
  // proto:  void QIODevice::QIODevice();
extern void* C_ZN9QIODeviceC2Ev(); // 3
  // proto:  bool QIODevice::reset();
extern void C_ZN9QIODevice5resetEv(void* qthis); // 4
  // proto:  OpenMode QIODevice::openMode();
extern void C_ZNK9QIODevice8openModeEv(void* qthis); // 4
  // proto:  bool QIODevice::isTextModeEnabled();
extern void C_ZNK9QIODevice17isTextModeEnabledEv(void* qthis); // 4
  // proto:  const QMetaObject * QIODevice::metaObject();
extern void C_ZNK9QIODevice10metaObjectEv(void* qthis); // 4
  // proto:  bool QIODevice::isOpen();
extern void C_ZNK9QIODevice6isOpenEv(void* qthis); // 4
  // proto:  bool QIODevice::atEnd();
extern void C_ZNK9QIODevice5atEndEv(void* qthis); // 4
  // proto:  bool QIODevice::waitForReadyRead(int msecs);
extern void C_ZN9QIODevice16waitForReadyReadEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QIODevice::isReadable();
extern void C_ZNK9QIODevice10isReadableEv(void* qthis); // 4
  // proto:  QByteArray QIODevice::readLine(qint64 maxlen);
extern void C_ZN9QIODevice8readLineEx(void* qthis, int64_t arg0); // 4
  // proto:  qint64 QIODevice::readLine(char * data, qint64 maxlen);
extern void C_ZN9QIODevice8readLineEPcx(void* qthis, unsigned char* arg0, int64_t arg1); // 4
  // proto:  bool QIODevice::isWritable();
extern void C_ZNK9QIODevice10isWritableEv(void* qthis); // 4
  // proto:  bool QIODevice::getChar(char * c);
extern void C_ZN9QIODevice7getCharEPc(void* qthis, unsigned char* arg0); // 4
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

// class sizeof(QIODevice)=1
type QIODevice struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _readyRead QIODevice_readyRead_signal;
//  _readChannelFinished QIODevice_readChannelFinished_signal;
//  _aboutToClose QIODevice_aboutToClose_signal;
//  _bytesWritten QIODevice_bytesWritten_signal;
}

// bytesAvailable()
func (this *QIODevice) bytesAvailable(args ...interface{}) () {
  // bytesAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice14bytesAvailableEv
    // invoke: qint64 bytesAvailable()
    var ret = C.C_ZNK9QIODevice14bytesAvailableEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "bytesAvailable", args)
  }

}

// pos()
func (this *QIODevice) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice3posEv
    // invoke: qint64 pos()
    var ret = C.C_ZNK9QIODevice3posEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "pos", args)
  }

}

// bytesToWrite()
func (this *QIODevice) bytesToWrite(args ...interface{}) () {
  // bytesToWrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice12bytesToWriteEv
    // invoke: qint64 bytesToWrite()
    var ret = C.C_ZNK9QIODevice12bytesToWriteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "bytesToWrite", args)
  }

}

// close()
func (this *QIODevice) close(args ...interface{}) () {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice5closeEv
    // invoke: void close()
    C.C_ZN9QIODevice5closeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIODevice", "close", args)
  }

}

// seek(qint64)
func (this *QIODevice) seek(args ...interface{}) () {
  // seek(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice4seekEx
    // invoke: bool seek(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QIODevice4seekEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "seek", args)
  }

}

// isSequential()
func (this *QIODevice) isSequential(args ...interface{}) () {
  // isSequential()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice12isSequentialEv
    // invoke: bool isSequential()
    var ret = C.C_ZNK9QIODevice12isSequentialEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "isSequential", args)
  }

}

// ungetChar(char)
func (this *QIODevice) ungetChar(args ...interface{}) () {
  // ungetChar(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice9ungetCharEc
    // invoke: void ungetChar(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.C_ZN9QIODevice9ungetCharEc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIODevice", "ungetChar", args)
  }

}

// size()
func (this *QIODevice) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice4sizeEv
    // invoke: qint64 size()
    var ret = C.C_ZNK9QIODevice4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "size", args)
  }

}

// write(const char *, qint64)
func (this *QIODevice) write(args ...interface{}) () {
  // write(const char *, qint64)
  // write(const char *)
  // write(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int64Ty(false) // "qint64"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice5writeEPKcx
    // invoke: qint64 write(const char *, qint64)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int64_t(args[1].(int64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN9QIODevice5writeEPKcx(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN9QIODevice5writeEPKc
    // invoke: qint64 write(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QIODevice5writeEPKc(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZN9QIODevice5writeERK10QByteArray
    // invoke: qint64 write(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QIODevice5writeERK10QByteArray(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "write", args)
  }

}

// canReadLine()
func (this *QIODevice) canReadLine(args ...interface{}) () {
  // canReadLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice11canReadLineEv
    // invoke: bool canReadLine()
    var ret = C.C_ZNK9QIODevice11canReadLineEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "canReadLine", args)
  }

}

// waitForBytesWritten(int)
func (this *QIODevice) waitForBytesWritten(args ...interface{}) () {
  // waitForBytesWritten(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice19waitForBytesWrittenEi
    // invoke: bool waitForBytesWritten(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QIODevice19waitForBytesWrittenEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "waitForBytesWritten", args)
  }

}

// read(char *, qint64)
func (this *QIODevice) read(args ...interface{}) () {
  // read(char *, qint64)
  // read(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "char *"
  vtys[0][1] = qtrt.Int64Ty(false) // "qint64"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice4readEPcx
    // invoke: qint64 read(char *, qint64)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int64_t(args[1].(int64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN9QIODevice4readEPcx(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN9QIODevice4readEx
    // invoke: QByteArray read(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QIODevice4readEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "read", args)
  }

}

// peek(char *, qint64)
func (this *QIODevice) peek(args ...interface{}) () {
  // peek(char *, qint64)
  // peek(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "char *"
  vtys[0][1] = qtrt.Int64Ty(false) // "qint64"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice4peekEPcx
    // invoke: qint64 peek(char *, qint64)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int64_t(args[1].(int64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN9QIODevice4peekEPcx(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN9QIODevice4peekEx
    // invoke: QByteArray peek(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QIODevice4peekEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "peek", args)
  }

}

// putChar(char)
func (this *QIODevice) putChar(args ...interface{}) () {
  // putChar(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice7putCharEc
    // invoke: bool putChar(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QIODevice7putCharEc(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "putChar", args)
  }

}

// errorString()
func (this *QIODevice) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice11errorStringEv
    // invoke: QString errorString()
    var ret = C.C_ZNK9QIODevice11errorStringEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "errorString", args)
  }

}

// ~QIODevice()
func (this *QIODevice) FreeQIODevice(args ...interface{}) () {
  // ~QIODevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODeviceD0Ev
    // invoke: void ~QIODevice()
    C.C_ZN9QIODeviceD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIODevice", "~QIODevice", args)
  }

}

// readAll()
func (this *QIODevice) readAll(args ...interface{}) () {
  // readAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice7readAllEv
    // invoke: QByteArray readAll()
    var ret = C.C_ZN9QIODevice7readAllEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "readAll", args)
  }

}

// setTextModeEnabled(_Bool)
func (this *QIODevice) setTextModeEnabled(args ...interface{}) () {
  // setTextModeEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice18setTextModeEnabledEb
    // invoke: void setTextModeEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QIODevice18setTextModeEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIODevice", "setTextModeEnabled", args)
  }

}

// QIODevice(class QObject *)
func NewQIODevice(args ...interface{}) *QIODevice {
  // QIODevice(class QObject *)
  // QIODevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODeviceC1EP7QObject
    // invoke: void QIODevice(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QIODeviceC2EP7QObject(arg0)
    return &QIODevice{qclsinst:qthis}
  case 1:
    // invoke: _ZN9QIODeviceC1Ev
    // invoke: void QIODevice()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QIODeviceC2Ev()
    return &QIODevice{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QIODevice", "QIODevice", args)
  }

  return nil // QIODevice{qclsinst:qthis}
}

// reset()
func (this *QIODevice) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice5resetEv
    // invoke: bool reset()
    var ret = C.C_ZN9QIODevice5resetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "reset", args)
  }

}

// openMode()
func (this *QIODevice) openMode(args ...interface{}) () {
  // openMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice8openModeEv
    // invoke: OpenMode openMode()
    C.C_ZNK9QIODevice8openModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIODevice", "openMode", args)
  }

}

// isTextModeEnabled()
func (this *QIODevice) isTextModeEnabled(args ...interface{}) () {
  // isTextModeEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice17isTextModeEnabledEv
    // invoke: bool isTextModeEnabled()
    var ret = C.C_ZNK9QIODevice17isTextModeEnabledEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "isTextModeEnabled", args)
  }

}

// metaObject()
func (this *QIODevice) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QIODevice10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIODevice", "metaObject", args)
  }

}

// isOpen()
func (this *QIODevice) isOpen(args ...interface{}) () {
  // isOpen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice6isOpenEv
    // invoke: bool isOpen()
    var ret = C.C_ZNK9QIODevice6isOpenEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "isOpen", args)
  }

}

// atEnd()
func (this *QIODevice) atEnd(args ...interface{}) () {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice5atEndEv
    // invoke: bool atEnd()
    var ret = C.C_ZNK9QIODevice5atEndEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "atEnd", args)
  }

}

// waitForReadyRead(int)
func (this *QIODevice) waitForReadyRead(args ...interface{}) () {
  // waitForReadyRead(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice16waitForReadyReadEi
    // invoke: bool waitForReadyRead(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QIODevice16waitForReadyReadEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "waitForReadyRead", args)
  }

}

// isReadable()
func (this *QIODevice) isReadable(args ...interface{}) () {
  // isReadable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice10isReadableEv
    // invoke: bool isReadable()
    var ret = C.C_ZNK9QIODevice10isReadableEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "isReadable", args)
  }

}

// readLine(qint64)
func (this *QIODevice) readLine(args ...interface{}) () {
  // readLine(qint64)
  // readLine(char *, qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "char *"
  vtys[1][1] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice8readLineEx
    // invoke: QByteArray readLine(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QIODevice8readLineEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN9QIODevice8readLineEPcx
    // invoke: qint64 readLine(char *, qint64)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int64_t(args[1].(int64))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN9QIODevice8readLineEPcx(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "readLine", args)
  }

}

// isWritable()
func (this *QIODevice) isWritable(args ...interface{}) () {
  // isWritable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice10isWritableEv
    // invoke: bool isWritable()
    var ret = C.C_ZNK9QIODevice10isWritableEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "isWritable", args)
  }

}

// getChar(char *)
func (this *QIODevice) getChar(args ...interface{}) () {
  // getChar(char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice7getCharEPc
    // invoke: bool getChar(char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QIODevice7getCharEPc(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QIODevice", "getChar", args)
  }

}

// <= body block end

