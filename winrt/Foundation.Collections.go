package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"log"
	"syscall"
	"unsafe"
)

// enums

// enum
type CollectionChange int32

const (
	CollectionChange_Reset        CollectionChange = 0
	CollectionChange_ItemInserted CollectionChange = 1
	CollectionChange_ItemRemoved  CollectionChange = 2
	CollectionChange_ItemChanged  CollectionChange = 3
)

// func types

// 179517F3-94EE-41F8-BDDC-768A895544F3
type MapChangedEventHandler[K any, V any] func(sender *IObservableMap[K, V], event *IMapChangedEventArgs[K]) com.Error

// 0C051752-9FBF-4C70-AA0C-0E4C82D9A761
type VectorChangedEventHandler[T any] func(sender *IObservableVector[T], event *IVectorChangedEventArgs) com.Error

// interfaces

// FAA585EA-6214-4217-AFDA-7F46DE5869B3
var IID_IIterable = syscall.GUID{0xFAA585EA, 0x6214, 0x4217,
	[8]byte{0xAF, 0xDA, 0x7F, 0x46, 0xDE, 0x58, 0x69, 0xB3}}

type IIterableInterface[T any] interface {
	win32.IInspectableInterface
	First() *IIterator[T]
}

type IIterableVtbl struct {
	win32.IInspectableVtbl
	First uintptr
}

type IIterable[T any] struct {
	win32.IInspectable
}

func (this *IIterable[T]) Vtbl() *IIterableVtbl {
	return (*IIterableVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IIterable[T]) First() *IIterator[T] {
	var _result *IIterator[T]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().First, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6A79E863-4300-459A-9966-CBB660963EE1
var IID_IIterator = syscall.GUID{0x6A79E863, 0x4300, 0x459A,
	[8]byte{0x99, 0x66, 0xCB, 0xB6, 0x60, 0x96, 0x3E, 0xE1}}

type IIteratorInterface[T any] interface {
	win32.IInspectableInterface
	Get_Current() T
	Get_HasCurrent() bool
	MoveNext() bool
	GetMany(itemsLength uint32, items *T) uint32
}

type IIteratorVtbl struct {
	win32.IInspectableVtbl
	Get_Current    uintptr
	Get_HasCurrent uintptr
	MoveNext       uintptr
	GetMany        uintptr
}

type IIterator[T any] struct {
	win32.IInspectable
}

func (this *IIterator[T]) Vtbl() *IIteratorVtbl {
	return (*IIteratorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IIterator[T]) Get_Current() T {
	var _result T
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Current, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return PostProcessGenericResult(_result)
}

func (this *IIterator[T]) Get_HasCurrent() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasCurrent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIterator[T]) MoveNext() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MoveNext, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IIterator[T]) GetMany(itemsLength uint32, items *T) uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMany, uintptr(unsafe.Pointer(this)), uintptr(itemsLength), uintptr(unsafe.Pointer(items)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 02B51929-C1C4-4A7E-8940-0312B5C18500
var IID_IKeyValuePair = syscall.GUID{0x02B51929, 0xC1C4, 0x4A7E,
	[8]byte{0x89, 0x40, 0x03, 0x12, 0xB5, 0xC1, 0x85, 0x00}}

type IKeyValuePairInterface[K any, V any] interface {
	win32.IInspectableInterface
	Get_Key() K
	Get_Value() V
}

type IKeyValuePairVtbl struct {
	win32.IInspectableVtbl
	Get_Key   uintptr
	Get_Value uintptr
}

type IKeyValuePair[K any, V any] struct {
	win32.IInspectable
}

func (this *IKeyValuePair[K, V]) Vtbl() *IKeyValuePairVtbl {
	return (*IKeyValuePairVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IKeyValuePair[K, V]) Get_Key() K {
	var _result K
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Key, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return PostProcessGenericResult(_result)
}

func (this *IKeyValuePair[K, V]) Get_Value() V {
	var _result V
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return PostProcessGenericResult(_result)
}

// 9939F4DF-050A-4C0F-AA60-77075F9C4777
var IID_IMapChangedEventArgs = syscall.GUID{0x9939F4DF, 0x050A, 0x4C0F,
	[8]byte{0xAA, 0x60, 0x77, 0x07, 0x5F, 0x9C, 0x47, 0x77}}

type IMapChangedEventArgsInterface[K any] interface {
	win32.IInspectableInterface
	Get_CollectionChange() CollectionChange
	Get_Key() K
}

type IMapChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_CollectionChange uintptr
	Get_Key              uintptr
}

type IMapChangedEventArgs[K any] struct {
	win32.IInspectable
}

func (this *IMapChangedEventArgs[K]) Vtbl() *IMapChangedEventArgsVtbl {
	return (*IMapChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMapChangedEventArgs[K]) Get_CollectionChange() CollectionChange {
	var _result CollectionChange
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CollectionChange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMapChangedEventArgs[K]) Get_Key() K {
	var _result K
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Key, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return PostProcessGenericResult(_result)
}

// E480CE40-A338-4ADA-ADCF-272272E48CB9
var IID_IMapView = syscall.GUID{0xE480CE40, 0xA338, 0x4ADA,
	[8]byte{0xAD, 0xCF, 0x27, 0x22, 0x72, 0xE4, 0x8C, 0xB9}}

type IMapViewInterface[K any, V any] interface {
	win32.IInspectableInterface
	Lookup(key K) V
	Get_Size() uint32
	HasKey(key K) bool
	Split(first *IMapView[K, V], second *IMapView[K, V])
}

type IMapViewVtbl struct {
	win32.IInspectableVtbl
	Lookup   uintptr
	Get_Size uintptr
	HasKey   uintptr
	Split    uintptr
}

type IMapView[K any, V any] struct {
	win32.IInspectable
}

func (this *IMapView[K, V]) Vtbl() *IMapViewVtbl {
	return (*IMapViewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMapView[K, V]) Lookup(key K) V {
	var _result V
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Lookup, uintptr(unsafe.Pointer(this)), uintptr(CastArgToPointer(key)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return PostProcessGenericResult(_result)
}

func (this *IMapView[K, V]) Get_Size() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMapView[K, V]) HasKey(key K) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().HasKey, uintptr(unsafe.Pointer(this)), uintptr(CastArgToPointer(key)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMapView[K, V]) Split(first *IMapView[K, V], second *IMapView[K, V]) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Split, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(first)), uintptr(unsafe.Pointer(second)))
	_ = _hr
}

// 3C2925FE-8519-45C1-AA79-197B6718C1C1
var IID_IMap = syscall.GUID{0x3C2925FE, 0x8519, 0x45C1,
	[8]byte{0xAA, 0x79, 0x19, 0x7B, 0x67, 0x18, 0xC1, 0xC1}}

type IMapInterface[K any, V any] interface {
	win32.IInspectableInterface
	Lookup(key K) V
	Get_Size() uint32
	HasKey(key K) bool
	GetView() *IMapView[K, V]
	Insert(key K, value V) bool
	Remove(key K)
	Clear()
}

type IMapVtbl struct {
	win32.IInspectableVtbl
	Lookup   uintptr
	Get_Size uintptr
	HasKey   uintptr
	GetView  uintptr
	Insert   uintptr
	Remove   uintptr
	Clear    uintptr
}

type IMap[K any, V any] struct {
	win32.IInspectable
}

func (this *IMap[K, V]) Vtbl() *IMapVtbl {
	return (*IMapVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMap[K, V]) Lookup(key K) V {
	var _result V
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Lookup, uintptr(unsafe.Pointer(this)), uintptr(CastArgToPointer(key)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return PostProcessGenericResult(_result)
}

func (this *IMap[K, V]) Get_Size() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMap[K, V]) HasKey(key K) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().HasKey, uintptr(unsafe.Pointer(this)), uintptr(CastArgToPointer(key)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMap[K, V]) GetView() *IMapView[K, V] {
	var _result *IMapView[K, V]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMap[K, V]) Insert(key K, value V) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Insert, uintptr(unsafe.Pointer(this)), uintptr(CastArgToPointer(key)), uintptr(CastArgToPointer(value)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMap[K, V]) Remove(key K) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove, uintptr(unsafe.Pointer(this)), uintptr(CastArgToPointer(key)))
	_ = _hr
}

func (this *IMap[K, V]) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

// 65DF2BF5-BF39-41B5-AEBC-5A9D865E472B
var IID_IObservableMap = syscall.GUID{0x65DF2BF5, 0xBF39, 0x41B5,
	[8]byte{0xAE, 0xBC, 0x5A, 0x9D, 0x86, 0x5E, 0x47, 0x2B}}

type IObservableMapInterface[K any, V any] interface {
	win32.IInspectableInterface
	Add_MapChanged(vhnd MapChangedEventHandler[K, V]) EventRegistrationToken
	Remove_MapChanged(token EventRegistrationToken)
}

type IObservableMapVtbl struct {
	win32.IInspectableVtbl
	Add_MapChanged    uintptr
	Remove_MapChanged uintptr
}

type IObservableMap[K any, V any] struct {
	win32.IInspectable
}

func (this *IObservableMap[K, V]) Vtbl() *IObservableMapVtbl {
	return (*IObservableMapVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IObservableMap[K, V]) Add_MapChanged(vhnd MapChangedEventHandler[K, V]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_MapChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(vhnd))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IObservableMap[K, V]) Remove_MapChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_MapChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 5917EB53-50B4-4A0D-B309-65862B3F1DBC
var IID_IObservableVector = syscall.GUID{0x5917EB53, 0x50B4, 0x4A0D,
	[8]byte{0xB3, 0x09, 0x65, 0x86, 0x2B, 0x3F, 0x1D, 0xBC}}

type IObservableVectorInterface[T any] interface {
	win32.IInspectableInterface
	Add_VectorChanged(vhnd VectorChangedEventHandler[T]) EventRegistrationToken
	Remove_VectorChanged(token EventRegistrationToken)
}

type IObservableVectorVtbl struct {
	win32.IInspectableVtbl
	Add_VectorChanged    uintptr
	Remove_VectorChanged uintptr
}

type IObservableVector[T any] struct {
	win32.IInspectable
}

func (this *IObservableVector[T]) Vtbl() *IObservableVectorVtbl {
	return (*IObservableVectorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IObservableVector[T]) Add_VectorChanged(vhnd VectorChangedEventHandler[T]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_VectorChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(vhnd))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IObservableVector[T]) Remove_VectorChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_VectorChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

// 8A43ED9F-F4E6-4421-ACF9-1DAB2986820C
var IID_IPropertySet = syscall.GUID{0x8A43ED9F, 0xF4E6, 0x4421,
	[8]byte{0xAC, 0xF9, 0x1D, 0xAB, 0x29, 0x86, 0x82, 0x0C}}

type IPropertySetInterface interface {
	win32.IInspectableInterface
}

type IPropertySetVtbl struct {
	win32.IInspectableVtbl
}

type IPropertySet struct {
	win32.IInspectable
}

func (this *IPropertySet) Vtbl() *IPropertySetVtbl {
	return (*IPropertySetVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 575933DF-34FE-4480-AF15-07691F3D5D9B
var IID_IVectorChangedEventArgs = syscall.GUID{0x575933DF, 0x34FE, 0x4480,
	[8]byte{0xAF, 0x15, 0x07, 0x69, 0x1F, 0x3D, 0x5D, 0x9B}}

type IVectorChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_CollectionChange() CollectionChange
	Get_Index() uint32
}

type IVectorChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_CollectionChange uintptr
	Get_Index            uintptr
}

type IVectorChangedEventArgs struct {
	win32.IInspectable
}

func (this *IVectorChangedEventArgs) Vtbl() *IVectorChangedEventArgsVtbl {
	return (*IVectorChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVectorChangedEventArgs) Get_CollectionChange() CollectionChange {
	var _result CollectionChange
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CollectionChange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVectorChangedEventArgs) Get_Index() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Index, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// BBE1FA4C-B0E3-4583-BAEF-1F1B2E483E56
var IID_IVectorView = syscall.GUID{0xBBE1FA4C, 0xB0E3, 0x4583,
	[8]byte{0xBA, 0xEF, 0x1F, 0x1B, 0x2E, 0x48, 0x3E, 0x56}}

type IVectorViewInterface[T any] interface {
	win32.IInspectableInterface
	GetAt(index uint32) T
	Get_Size() uint32
	IndexOf(value T, index uint32) bool
	GetMany(startIndex uint32, itemsLength uint32, items *T) uint32
}

type IVectorViewVtbl struct {
	win32.IInspectableVtbl
	GetAt    uintptr
	Get_Size uintptr
	IndexOf  uintptr
	GetMany  uintptr
}

type IVectorView[T any] struct {
	win32.IInspectable
}

func (this *IVectorView[T]) Vtbl() *IVectorViewVtbl {
	return (*IVectorViewVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVectorView[T]) GetAt(index uint32) T {
	var _result T
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return PostProcessGenericResult(_result)
}

func (this *IVectorView[T]) Get_Size() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVectorView[T]) IndexOf(value T, index uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IndexOf, uintptr(unsafe.Pointer(this)), uintptr(CastArgToPointer(value)), uintptr(index), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVectorView[T]) GetMany(startIndex uint32, itemsLength uint32, items *T) uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMany, uintptr(unsafe.Pointer(this)), uintptr(startIndex), uintptr(itemsLength), uintptr(unsafe.Pointer(items)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 913337E9-11A1-4345-A3A2-4E7F956E222D
var IID_IVector = syscall.GUID{0x913337E9, 0x11A1, 0x4345,
	[8]byte{0xA3, 0xA2, 0x4E, 0x7F, 0x95, 0x6E, 0x22, 0x2D}}

type IVectorInterface[T any] interface {
	win32.IInspectableInterface
	GetAt(index uint32) T
	Get_Size() uint32
	GetView() *IVectorView[T]
	IndexOf(value T, index uint32) bool
	SetAt(index uint32, value T)
	InsertAt(index uint32, value T)
	RemoveAt(index uint32)
	Append(value T)
	RemoveAtEnd()
	Clear()
	GetMany(startIndex uint32, itemsLength uint32, items *T) uint32
	ReplaceAll(itemsLength uint32, items *T)
}

type IVectorVtbl struct {
	win32.IInspectableVtbl
	GetAt       uintptr
	Get_Size    uintptr
	GetView     uintptr
	IndexOf     uintptr
	SetAt       uintptr
	InsertAt    uintptr
	RemoveAt    uintptr
	Append      uintptr
	RemoveAtEnd uintptr
	Clear       uintptr
	GetMany     uintptr
	ReplaceAll  uintptr
}

type IVector[T any] struct {
	win32.IInspectable
}

func (this *IVector[T]) Vtbl() *IVectorVtbl {
	return (*IVectorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVector[T]) GetAt(index uint32) T {
	var _result T
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return PostProcessGenericResult(_result)
}

func (this *IVector[T]) Get_Size() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVector[T]) GetView() *IVectorView[T] {
	var _result *IVectorView[T]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVector[T]) IndexOf(value T, index uint32) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IndexOf, uintptr(unsafe.Pointer(this)), uintptr(CastArgToPointer(value)), uintptr(index), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVector[T]) SetAt(index uint32, value T) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(CastArgToPointer(value)))
	_ = _hr
}

func (this *IVector[T]) InsertAt(index uint32, value T) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(CastArgToPointer(value)))
	_ = _hr
}

func (this *IVector[T]) RemoveAt(index uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	_ = _hr
}

func (this *IVector[T]) Append(value T) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(CastArgToPointer(value)))
	_ = _hr
}

func (this *IVector[T]) RemoveAtEnd() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemoveAtEnd, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IVector[T]) Clear() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IVector[T]) GetMany(startIndex uint32, itemsLength uint32, items *T) uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMany, uintptr(unsafe.Pointer(this)), uintptr(startIndex), uintptr(itemsLength), uintptr(unsafe.Pointer(items)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVector[T]) ReplaceAll(itemsLength uint32, items *T) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReplaceAll, uintptr(unsafe.Pointer(this)), uintptr(itemsLength), uintptr(unsafe.Pointer(items)))
	_ = _hr
}

// classes

type PropertySet struct {
	RtClass
	*IPropertySet
}

func NewPropertySet() *PropertySet {
	hs := NewHStr("Windows.Foundation.Collections.PropertySet")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PropertySet{
		RtClass:      RtClass{PInspect: p},
		IPropertySet: (*IPropertySet)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type StringMap struct {
	RtClass
	*IMap[string, string]
}

func NewStringMap() *StringMap {
	hs := NewHStr("Windows.Foundation.Collections.StringMap")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &StringMap{
		RtClass: RtClass{PInspect: p},
		IMap:    (*IMap[string, string])(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type ValueSet struct {
	RtClass
	*IPropertySet
}

func NewValueSet() *ValueSet {
	hs := NewHStr("Windows.Foundation.Collections.ValueSet")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &ValueSet{
		RtClass:      RtClass{PInspect: p},
		IPropertySet: (*IPropertySet)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
