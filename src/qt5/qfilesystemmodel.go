package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtWidgets/qfilesystemmodel.h
// dst-file: /src/widgets/qfilesystemmodel.go
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
  // proto:  void QFileSystemModel::setNameFilters(const QStringList & filters);
extern void C_ZN16QFileSystemModel14setNameFiltersERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  int QFileSystemModel::columnCount(const QModelIndex & parent);
extern void C_ZNK16QFileSystemModel11columnCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QVariant QFileSystemModel::myComputer(int role);
extern void C_ZNK16QFileSystemModel10myComputerEi(void* qthis, int32_t arg0); // 4
  // proto:  QModelIndex QFileSystemModel::mkdir(const QModelIndex & parent, const QString & name);
extern void C_ZN16QFileSystemModel5mkdirERK11QModelIndexRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QModelIndex QFileSystemModel::setRootPath(const QString & path);
extern void C_ZN16QFileSystemModel11setRootPathERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFileSystemModel::~QFileSystemModel();
extern void C_ZN16QFileSystemModelD2Ev(void* qthis); // 4
  // proto:  void QFileSystemModel::QFileSystemModel(QObject * parent);
extern void* C_ZN16QFileSystemModelC2EP7QObject(void* arg0); // 3
  // proto:  qint64 QFileSystemModel::size(const QModelIndex & index);
extern void C_ZNK16QFileSystemModel4sizeERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QFileSystemModel::index(const QString & path, int column);
extern void C_ZNK16QFileSystemModel5indexERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QModelIndex QFileSystemModel::index(int row, int column, const QModelIndex & parent);
extern void C_ZNK16QFileSystemModel5indexEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QFileSystemModel::resolveSymlinks();
extern void C_ZNK16QFileSystemModel15resolveSymlinksEv(void* qthis); // 4
  // proto:  bool QFileSystemModel::hasChildren(const QModelIndex & parent);
extern void C_ZNK16QFileSystemModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QFileSystemModel::isReadOnly();
extern void C_ZNK16QFileSystemModel10isReadOnlyEv(void* qthis); // 4
  // proto:  bool QFileSystemModel::isDir(const QModelIndex & index);
extern void C_ZNK16QFileSystemModel5isDirERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QStringList QFileSystemModel::nameFilters();
extern void C_ZNK16QFileSystemModel11nameFiltersEv(void* qthis); // 4
  // proto:  bool QFileSystemModel::rmdir(const QModelIndex & index);
extern void C_ZN16QFileSystemModel5rmdirERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QFileInfo QFileSystemModel::fileInfo(const QModelIndex & index);
extern void C_ZNK16QFileSystemModel8fileInfoERK11QModelIndex(void* qthis, void* arg0); // 2
  // proto:  QFile::Permissions QFileSystemModel::permissions(const QModelIndex & index);
extern void C_ZNK16QFileSystemModel11permissionsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QString QFileSystemModel::type(const QModelIndex & index);
extern void C_ZNK16QFileSystemModel4typeERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QFileSystemModel::setNameFilterDisables(bool enable);
extern void C_ZN16QFileSystemModel21setNameFilterDisablesEb(void* qthis, bool arg0); // 4
  // proto:  void QFileSystemModel::setIconProvider(QFileIconProvider * provider);
extern void C_ZN16QFileSystemModel15setIconProviderEP17QFileIconProvider(void* qthis, void* arg0); // 4
  // proto:  bool QFileSystemModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern void C_ZN16QFileSystemModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  bool QFileSystemModel::canFetchMore(const QModelIndex & parent);
extern void C_ZNK16QFileSystemModel12canFetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QFileSystemModel::parent(const QModelIndex & child);
extern void C_ZNK16QFileSystemModel6parentERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QString QFileSystemModel::filePath(const QModelIndex & index);
extern void C_ZNK16QFileSystemModel8filePathERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QString QFileSystemModel::fileName(const QModelIndex & index);
extern void C_ZNK16QFileSystemModel8fileNameERK11QModelIndex(void* qthis, void* arg0); // 2
  // proto:  int QFileSystemModel::rowCount(const QModelIndex & parent);
extern void C_ZNK16QFileSystemModel8rowCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QFileIconProvider * QFileSystemModel::iconProvider();
extern void C_ZNK16QFileSystemModel12iconProviderEv(void* qthis); // 4
  // proto:  bool QFileSystemModel::nameFilterDisables();
extern void C_ZNK16QFileSystemModel18nameFilterDisablesEv(void* qthis); // 4
  // proto:  QVariant QFileSystemModel::data(const QModelIndex & index, int role);
extern void C_ZNK16QFileSystemModel4dataERK11QModelIndexi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QStringList QFileSystemModel::mimeTypes();
extern void C_ZNK16QFileSystemModel9mimeTypesEv(void* qthis); // 4
  // proto:  const QMetaObject * QFileSystemModel::metaObject();
extern void C_ZNK16QFileSystemModel10metaObjectEv(void* qthis); // 4
  // proto:  void QFileSystemModel::setResolveSymlinks(bool enable);
extern void C_ZN16QFileSystemModel18setResolveSymlinksEb(void* qthis, bool arg0); // 4
  // proto:  QString QFileSystemModel::rootPath();
extern void C_ZNK16QFileSystemModel8rootPathEv(void* qthis); // 4
  // proto:  bool QFileSystemModel::remove(const QModelIndex & index);
extern void C_ZN16QFileSystemModel6removeERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QDir::Filters QFileSystemModel::filter();
extern void C_ZNK16QFileSystemModel6filterEv(void* qthis); // 4
  // proto:  void QFileSystemModel::setReadOnly(bool enable);
extern void C_ZN16QFileSystemModel11setReadOnlyEb(void* qthis, bool arg0); // 4
  // proto:  Qt::ItemFlags QFileSystemModel::flags(const QModelIndex & index);
extern void C_ZNK16QFileSystemModel5flagsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QDir QFileSystemModel::rootDirectory();
extern void C_ZNK16QFileSystemModel13rootDirectoryEv(void* qthis); // 4
  // proto:  QDateTime QFileSystemModel::lastModified(const QModelIndex & index);
extern void C_ZNK16QFileSystemModel12lastModifiedERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QIcon QFileSystemModel::fileIcon(const QModelIndex & index);
extern void C_ZNK16QFileSystemModel8fileIconERK11QModelIndex(void* qthis, void* arg0); // 2
  // proto:  void QFileSystemModel::fetchMore(const QModelIndex & parent);
extern void C_ZN16QFileSystemModel9fetchMoreERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  Qt::DropActions QFileSystemModel::supportedDropActions();
extern void C_ZNK16QFileSystemModel20supportedDropActionsEv(void* qthis); // 4
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

// class sizeof(QFileSystemModel)=1
type QFileSystemModel struct {
  /*qbase*/ QAbstractItemModel;
  qclsinst unsafe.Pointer /* *C.void */;
//  _rootPathChanged QFileSystemModel_rootPathChanged_signal;
//  _directoryLoaded QFileSystemModel_directoryLoaded_signal;
//  _fileRenamed QFileSystemModel_fileRenamed_signal;
}

// setNameFilters(const class QStringList &)
func (this *QFileSystemModel) setNameFilters(args ...interface{}) () {
  // setNameFilters(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel14setNameFiltersERK11QStringList
    // invoke: void setNameFilters(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QFileSystemModel14setNameFiltersERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setNameFilters", args)
  }

}

// columnCount(const class QModelIndex &)
func (this *QFileSystemModel) columnCount(args ...interface{}) () {
  // columnCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel11columnCountERK11QModelIndex
    // invoke: int columnCount(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel11columnCountERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "columnCount", args)
  }

}

// myComputer(int)
func (this *QFileSystemModel) myComputer(args ...interface{}) () {
  // myComputer(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel10myComputerEi
    // invoke: QVariant myComputer(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel10myComputerEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "myComputer", args)
  }

}

// mkdir(const class QModelIndex &, const class QString &)
func (this *QFileSystemModel) mkdir(args ...interface{}) () {
  // mkdir(const class QModelIndex &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel5mkdirERK11QModelIndexRK7QString
    // invoke: QModelIndex mkdir(const class QModelIndex &, const class QString &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN16QFileSystemModel5mkdirERK11QModelIndexRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "mkdir", args)
  }

}

// setRootPath(const class QString &)
func (this *QFileSystemModel) setRootPath(args ...interface{}) () {
  // setRootPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel11setRootPathERK7QString
    // invoke: QModelIndex setRootPath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN16QFileSystemModel11setRootPathERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setRootPath", args)
  }

}

// ~QFileSystemModel()
func (this *QFileSystemModel) FreeQFileSystemModel(args ...interface{}) () {
  // ~QFileSystemModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModelD0Ev
    // invoke: void ~QFileSystemModel()
    C.C_ZN16QFileSystemModelD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "~QFileSystemModel", args)
  }

}

// QFileSystemModel(class QObject *)
func NewQFileSystemModel(args ...interface{}) *QFileSystemModel {
  // QFileSystemModel(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModelC1EP7QObject
    // invoke: void QFileSystemModel(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QFileSystemModelC2EP7QObject(arg0)
    return &QFileSystemModel{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "QFileSystemModel", args)
  }

  return nil // QFileSystemModel{qclsinst:qthis}
}

// size(const class QModelIndex &)
func (this *QFileSystemModel) size(args ...interface{}) () {
  // size(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel4sizeERK11QModelIndex
    // invoke: qint64 size(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel4sizeERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "size", args)
  }

}

// index(const class QString &, int)
func (this *QFileSystemModel) index(args ...interface{}) () {
  // index(const class QString &, int)
  // index(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel5indexERK7QStringi
    // invoke: QModelIndex index(const class QString &, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK16QFileSystemModel5indexERK7QStringi(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK16QFileSystemModel5indexEiiRK11QModelIndex
    // invoke: QModelIndex index(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZNK16QFileSystemModel5indexEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "index", args)
  }

}

// resolveSymlinks()
func (this *QFileSystemModel) resolveSymlinks(args ...interface{}) () {
  // resolveSymlinks()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel15resolveSymlinksEv
    // invoke: bool resolveSymlinks()
    var ret = C.C_ZNK16QFileSystemModel15resolveSymlinksEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "resolveSymlinks", args)
  }

}

// hasChildren(const class QModelIndex &)
func (this *QFileSystemModel) hasChildren(args ...interface{}) () {
  // hasChildren(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel11hasChildrenERK11QModelIndex
    // invoke: bool hasChildren(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel11hasChildrenERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "hasChildren", args)
  }

}

// isReadOnly()
func (this *QFileSystemModel) isReadOnly(args ...interface{}) () {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel10isReadOnlyEv
    // invoke: bool isReadOnly()
    var ret = C.C_ZNK16QFileSystemModel10isReadOnlyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "isReadOnly", args)
  }

}

// isDir(const class QModelIndex &)
func (this *QFileSystemModel) isDir(args ...interface{}) () {
  // isDir(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel5isDirERK11QModelIndex
    // invoke: bool isDir(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel5isDirERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "isDir", args)
  }

}

// nameFilters()
func (this *QFileSystemModel) nameFilters(args ...interface{}) () {
  // nameFilters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel11nameFiltersEv
    // invoke: QStringList nameFilters()
    C.C_ZNK16QFileSystemModel11nameFiltersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "nameFilters", args)
  }

}

// rmdir(const class QModelIndex &)
func (this *QFileSystemModel) rmdir(args ...interface{}) () {
  // rmdir(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel5rmdirERK11QModelIndex
    // invoke: bool rmdir(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN16QFileSystemModel5rmdirERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "rmdir", args)
  }

}

// fileInfo(const class QModelIndex &)
func (this *QFileSystemModel) fileInfo(args ...interface{}) () {
  // fileInfo(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel8fileInfoERK11QModelIndex
    // invoke: QFileInfo fileInfo(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel8fileInfoERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "fileInfo", args)
  }

}

// permissions(const class QModelIndex &)
func (this *QFileSystemModel) permissions(args ...interface{}) () {
  // permissions(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel11permissionsERK11QModelIndex
    // invoke: QFile::Permissions permissions(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK16QFileSystemModel11permissionsERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "permissions", args)
  }

}

// type(const class QModelIndex &)
func (this *QFileSystemModel) type_(args ...interface{}) () {
  // type(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel4typeERK11QModelIndex
    // invoke: QString type(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel4typeERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "type", args)
  }

}

// setNameFilterDisables(_Bool)
func (this *QFileSystemModel) setNameFilterDisables(args ...interface{}) () {
  // setNameFilterDisables(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel21setNameFilterDisablesEb
    // invoke: void setNameFilterDisables(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QFileSystemModel21setNameFilterDisablesEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setNameFilterDisables", args)
  }

}

// setIconProvider(class QFileIconProvider *)
func (this *QFileSystemModel) setIconProvider(args ...interface{}) () {
  // setIconProvider(class QFileIconProvider *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFileIconProvider{}) // "QFileIconProvider *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel15setIconProviderEP17QFileIconProvider
    // invoke: void setIconProvider(class QFileIconProvider *)
    var arg0 = args[0].(QFileIconProvider).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QFileSystemModel15setIconProviderEP17QFileIconProvider(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setIconProvider", args)
  }

}

// setData(const class QModelIndex &, const class QVariant &, int)
func (this *QFileSystemModel) setData(args ...interface{}) () {
  // setData(const class QModelIndex &, const class QVariant &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel7setDataERK11QModelIndexRK8QVarianti
    // invoke: bool setData(const class QModelIndex &, const class QVariant &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN16QFileSystemModel7setDataERK11QModelIndexRK8QVarianti(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setData", args)
  }

}

// canFetchMore(const class QModelIndex &)
func (this *QFileSystemModel) canFetchMore(args ...interface{}) () {
  // canFetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel12canFetchMoreERK11QModelIndex
    // invoke: bool canFetchMore(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel12canFetchMoreERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "canFetchMore", args)
  }

}

// parent(const class QModelIndex &)
func (this *QFileSystemModel) parent(args ...interface{}) () {
  // parent(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel6parentERK11QModelIndex
    // invoke: QModelIndex parent(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel6parentERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "parent", args)
  }

}

// filePath(const class QModelIndex &)
func (this *QFileSystemModel) filePath(args ...interface{}) () {
  // filePath(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel8filePathERK11QModelIndex
    // invoke: QString filePath(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel8filePathERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "filePath", args)
  }

}

// fileName(const class QModelIndex &)
func (this *QFileSystemModel) fileName(args ...interface{}) () {
  // fileName(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel8fileNameERK11QModelIndex
    // invoke: QString fileName(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel8fileNameERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "fileName", args)
  }

}

// rowCount(const class QModelIndex &)
func (this *QFileSystemModel) rowCount(args ...interface{}) () {
  // rowCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel8rowCountERK11QModelIndex
    // invoke: int rowCount(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel8rowCountERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "rowCount", args)
  }

}

// iconProvider()
func (this *QFileSystemModel) iconProvider(args ...interface{}) () {
  // iconProvider()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel12iconProviderEv
    // invoke: QFileIconProvider * iconProvider()
    var ret = C.C_ZNK16QFileSystemModel12iconProviderEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "iconProvider", args)
  }

}

// nameFilterDisables()
func (this *QFileSystemModel) nameFilterDisables(args ...interface{}) () {
  // nameFilterDisables()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel18nameFilterDisablesEv
    // invoke: bool nameFilterDisables()
    var ret = C.C_ZNK16QFileSystemModel18nameFilterDisablesEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "nameFilterDisables", args)
  }

}

// data(const class QModelIndex &, int)
func (this *QFileSystemModel) data(args ...interface{}) () {
  // data(const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel4dataERK11QModelIndexi
    // invoke: QVariant data(const class QModelIndex &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK16QFileSystemModel4dataERK11QModelIndexi(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "data", args)
  }

}

// mimeTypes()
func (this *QFileSystemModel) mimeTypes(args ...interface{}) () {
  // mimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel9mimeTypesEv
    // invoke: QStringList mimeTypes()
    C.C_ZNK16QFileSystemModel9mimeTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "mimeTypes", args)
  }

}

// metaObject()
func (this *QFileSystemModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK16QFileSystemModel10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "metaObject", args)
  }

}

// setResolveSymlinks(_Bool)
func (this *QFileSystemModel) setResolveSymlinks(args ...interface{}) () {
  // setResolveSymlinks(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel18setResolveSymlinksEb
    // invoke: void setResolveSymlinks(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QFileSystemModel18setResolveSymlinksEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setResolveSymlinks", args)
  }

}

// rootPath()
func (this *QFileSystemModel) rootPath(args ...interface{}) () {
  // rootPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel8rootPathEv
    // invoke: QString rootPath()
    var ret = C.C_ZNK16QFileSystemModel8rootPathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "rootPath", args)
  }

}

// remove(const class QModelIndex &)
func (this *QFileSystemModel) remove(args ...interface{}) () {
  // remove(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel6removeERK11QModelIndex
    // invoke: bool remove(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN16QFileSystemModel6removeERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "remove", args)
  }

}

// filter()
func (this *QFileSystemModel) filter(args ...interface{}) () {
  // filter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel6filterEv
    // invoke: QDir::Filters filter()
    C.C_ZNK16QFileSystemModel6filterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "filter", args)
  }

}

// setReadOnly(_Bool)
func (this *QFileSystemModel) setReadOnly(args ...interface{}) () {
  // setReadOnly(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel11setReadOnlyEb
    // invoke: void setReadOnly(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QFileSystemModel11setReadOnlyEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setReadOnly", args)
  }

}

// flags(const class QModelIndex &)
func (this *QFileSystemModel) flags(args ...interface{}) () {
  // flags(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel5flagsERK11QModelIndex
    // invoke: Qt::ItemFlags flags(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK16QFileSystemModel5flagsERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "flags", args)
  }

}

// rootDirectory()
func (this *QFileSystemModel) rootDirectory(args ...interface{}) () {
  // rootDirectory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel13rootDirectoryEv
    // invoke: QDir rootDirectory()
    var ret = C.C_ZNK16QFileSystemModel13rootDirectoryEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "rootDirectory", args)
  }

}

// lastModified(const class QModelIndex &)
func (this *QFileSystemModel) lastModified(args ...interface{}) () {
  // lastModified(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel12lastModifiedERK11QModelIndex
    // invoke: QDateTime lastModified(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel12lastModifiedERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "lastModified", args)
  }

}

// fileIcon(const class QModelIndex &)
func (this *QFileSystemModel) fileIcon(args ...interface{}) () {
  // fileIcon(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel8fileIconERK11QModelIndex
    // invoke: QIcon fileIcon(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QFileSystemModel8fileIconERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "fileIcon", args)
  }

}

// fetchMore(const class QModelIndex &)
func (this *QFileSystemModel) fetchMore(args ...interface{}) () {
  // fetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel9fetchMoreERK11QModelIndex
    // invoke: void fetchMore(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QFileSystemModel9fetchMoreERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "fetchMore", args)
  }

}

// supportedDropActions()
func (this *QFileSystemModel) supportedDropActions(args ...interface{}) () {
  // supportedDropActions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel20supportedDropActionsEv
    // invoke: Qt::DropActions supportedDropActions()
    C.C_ZNK16QFileSystemModel20supportedDropActionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileSystemModel", "supportedDropActions", args)
  }

}

// <= body block end

