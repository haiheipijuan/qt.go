package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtGui/qpixelformat.h
// dst-file: /src/gui/qpixelformat.go
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
  // proto:  QPixelFormat::YUVLayout QPixelFormat::yuvLayout();
extern void C_ZNK12QPixelFormat9yuvLayoutEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::yellowSize();
extern void C_ZNK12QPixelFormat10yellowSizeEv(void* qthis); // 2
  // proto:  QPixelFormat::AlphaPosition QPixelFormat::alphaPosition();
extern void C_ZNK12QPixelFormat13alphaPositionEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::cyanSize();
extern void C_ZNK12QPixelFormat8cyanSizeEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::lightnessSize();
extern void C_ZNK12QPixelFormat13lightnessSizeEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::brightnessSize();
extern void C_ZNK12QPixelFormat14brightnessSizeEv(void* qthis); // 2
  // proto:  void QPixelFormat::QPixelFormat();
extern void* C_ZN12QPixelFormatC2Ev(); // 1
  // proto:  uchar QPixelFormat::bitsPerPixel();
extern void C_ZNK12QPixelFormat12bitsPerPixelEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::blackSize();
extern void C_ZNK12QPixelFormat9blackSizeEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::redSize();
extern void C_ZNK12QPixelFormat7redSizeEv(void* qthis); // 2
  // proto:  QPixelFormat::ByteOrder QPixelFormat::byteOrder();
extern void C_ZNK12QPixelFormat9byteOrderEv(void* qthis); // 2
  // proto:  QPixelFormat::TypeInterpretation QPixelFormat::typeInterpretation();
extern void C_ZNK12QPixelFormat18typeInterpretationEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::channelCount();
extern void C_ZNK12QPixelFormat12channelCountEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::alphaSize();
extern void C_ZNK12QPixelFormat9alphaSizeEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::hueSize();
extern void C_ZNK12QPixelFormat7hueSizeEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::saturationSize();
extern void C_ZNK12QPixelFormat14saturationSizeEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::magentaSize();
extern void C_ZNK12QPixelFormat11magentaSizeEv(void* qthis); // 2
  // proto:  QPixelFormat::AlphaPremultiplied QPixelFormat::premultiplied();
extern void C_ZNK12QPixelFormat13premultipliedEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::subEnum();
extern void C_ZNK12QPixelFormat7subEnumEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::greenSize();
extern void C_ZNK12QPixelFormat9greenSizeEv(void* qthis); // 2
  // proto:  QPixelFormat::ColorModel QPixelFormat::colorModel();
extern void C_ZNK12QPixelFormat10colorModelEv(void* qthis); // 2
  // proto:  QPixelFormat::AlphaUsage QPixelFormat::alphaUsage();
extern void C_ZNK12QPixelFormat10alphaUsageEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::blueSize();
extern void C_ZNK12QPixelFormat8blueSizeEv(void* qthis); // 2
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

// class sizeof(QPixelFormat)=8
type QPixelFormat struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// yuvLayout()
func (this *QPixelFormat) yuvLayout(args ...interface{}) () {
  // yuvLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat9yuvLayoutEv
    // invoke: QPixelFormat::YUVLayout yuvLayout()
    C.C_ZNK12QPixelFormat9yuvLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "yuvLayout", args)
  }

}

// yellowSize()
func (this *QPixelFormat) yellowSize(args ...interface{}) () {
  // yellowSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat10yellowSizeEv
    // invoke: uchar yellowSize()
    var ret = C.C_ZNK12QPixelFormat10yellowSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "yellowSize", args)
  }

}

// alphaPosition()
func (this *QPixelFormat) alphaPosition(args ...interface{}) () {
  // alphaPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat13alphaPositionEv
    // invoke: QPixelFormat::AlphaPosition alphaPosition()
    C.C_ZNK12QPixelFormat13alphaPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "alphaPosition", args)
  }

}

// cyanSize()
func (this *QPixelFormat) cyanSize(args ...interface{}) () {
  // cyanSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat8cyanSizeEv
    // invoke: uchar cyanSize()
    var ret = C.C_ZNK12QPixelFormat8cyanSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "cyanSize", args)
  }

}

// lightnessSize()
func (this *QPixelFormat) lightnessSize(args ...interface{}) () {
  // lightnessSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat13lightnessSizeEv
    // invoke: uchar lightnessSize()
    var ret = C.C_ZNK12QPixelFormat13lightnessSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "lightnessSize", args)
  }

}

// brightnessSize()
func (this *QPixelFormat) brightnessSize(args ...interface{}) () {
  // brightnessSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat14brightnessSizeEv
    // invoke: uchar brightnessSize()
    var ret = C.C_ZNK12QPixelFormat14brightnessSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "brightnessSize", args)
  }

}

// QPixelFormat()
func NewQPixelFormat(args ...interface{}) *QPixelFormat {
  // QPixelFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPixelFormatC1Ev
    // invoke: void QPixelFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QPixelFormatC2Ev()
    return &QPixelFormat{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPixelFormat", "QPixelFormat", args)
  }

  return nil // QPixelFormat{qclsinst:qthis}
}

// bitsPerPixel()
func (this *QPixelFormat) bitsPerPixel(args ...interface{}) () {
  // bitsPerPixel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat12bitsPerPixelEv
    // invoke: uchar bitsPerPixel()
    var ret = C.C_ZNK12QPixelFormat12bitsPerPixelEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "bitsPerPixel", args)
  }

}

// blackSize()
func (this *QPixelFormat) blackSize(args ...interface{}) () {
  // blackSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat9blackSizeEv
    // invoke: uchar blackSize()
    var ret = C.C_ZNK12QPixelFormat9blackSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "blackSize", args)
  }

}

// redSize()
func (this *QPixelFormat) redSize(args ...interface{}) () {
  // redSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat7redSizeEv
    // invoke: uchar redSize()
    var ret = C.C_ZNK12QPixelFormat7redSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "redSize", args)
  }

}

// byteOrder()
func (this *QPixelFormat) byteOrder(args ...interface{}) () {
  // byteOrder()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat9byteOrderEv
    // invoke: QPixelFormat::ByteOrder byteOrder()
    C.C_ZNK12QPixelFormat9byteOrderEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "byteOrder", args)
  }

}

// typeInterpretation()
func (this *QPixelFormat) typeInterpretation(args ...interface{}) () {
  // typeInterpretation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat18typeInterpretationEv
    // invoke: QPixelFormat::TypeInterpretation typeInterpretation()
    C.C_ZNK12QPixelFormat18typeInterpretationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "typeInterpretation", args)
  }

}

// channelCount()
func (this *QPixelFormat) channelCount(args ...interface{}) () {
  // channelCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat12channelCountEv
    // invoke: uchar channelCount()
    var ret = C.C_ZNK12QPixelFormat12channelCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "channelCount", args)
  }

}

// alphaSize()
func (this *QPixelFormat) alphaSize(args ...interface{}) () {
  // alphaSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat9alphaSizeEv
    // invoke: uchar alphaSize()
    var ret = C.C_ZNK12QPixelFormat9alphaSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "alphaSize", args)
  }

}

// hueSize()
func (this *QPixelFormat) hueSize(args ...interface{}) () {
  // hueSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat7hueSizeEv
    // invoke: uchar hueSize()
    var ret = C.C_ZNK12QPixelFormat7hueSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "hueSize", args)
  }

}

// saturationSize()
func (this *QPixelFormat) saturationSize(args ...interface{}) () {
  // saturationSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat14saturationSizeEv
    // invoke: uchar saturationSize()
    var ret = C.C_ZNK12QPixelFormat14saturationSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "saturationSize", args)
  }

}

// magentaSize()
func (this *QPixelFormat) magentaSize(args ...interface{}) () {
  // magentaSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat11magentaSizeEv
    // invoke: uchar magentaSize()
    var ret = C.C_ZNK12QPixelFormat11magentaSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "magentaSize", args)
  }

}

// premultiplied()
func (this *QPixelFormat) premultiplied(args ...interface{}) () {
  // premultiplied()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat13premultipliedEv
    // invoke: QPixelFormat::AlphaPremultiplied premultiplied()
    C.C_ZNK12QPixelFormat13premultipliedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "premultiplied", args)
  }

}

// subEnum()
func (this *QPixelFormat) subEnum(args ...interface{}) () {
  // subEnum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat7subEnumEv
    // invoke: uchar subEnum()
    var ret = C.C_ZNK12QPixelFormat7subEnumEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "subEnum", args)
  }

}

// greenSize()
func (this *QPixelFormat) greenSize(args ...interface{}) () {
  // greenSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat9greenSizeEv
    // invoke: uchar greenSize()
    var ret = C.C_ZNK12QPixelFormat9greenSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "greenSize", args)
  }

}

// colorModel()
func (this *QPixelFormat) colorModel(args ...interface{}) () {
  // colorModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat10colorModelEv
    // invoke: QPixelFormat::ColorModel colorModel()
    C.C_ZNK12QPixelFormat10colorModelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "colorModel", args)
  }

}

// alphaUsage()
func (this *QPixelFormat) alphaUsage(args ...interface{}) () {
  // alphaUsage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat10alphaUsageEv
    // invoke: QPixelFormat::AlphaUsage alphaUsage()
    C.C_ZNK12QPixelFormat10alphaUsageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "alphaUsage", args)
  }

}

// blueSize()
func (this *QPixelFormat) blueSize(args ...interface{}) () {
  // blueSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat8blueSizeEv
    // invoke: uchar blueSize()
    var ret = C.C_ZNK12QPixelFormat8blueSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPixelFormat", "blueSize", args)
  }

}

// <= body block end

