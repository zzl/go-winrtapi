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
type ByteOrder int32

const (
	ByteOrder_LittleEndian ByteOrder = 0
	ByteOrder_BigEndian    ByteOrder = 1
)

// enum
type FileOpenDisposition int32

const (
	FileOpenDisposition_OpenExisting     FileOpenDisposition = 0
	FileOpenDisposition_OpenAlways       FileOpenDisposition = 1
	FileOpenDisposition_CreateNew        FileOpenDisposition = 2
	FileOpenDisposition_CreateAlways     FileOpenDisposition = 3
	FileOpenDisposition_TruncateExisting FileOpenDisposition = 4
)

// enum
// flags
type InputStreamOptions uint32

const (
	InputStreamOptions_None      InputStreamOptions = 0
	InputStreamOptions_Partial   InputStreamOptions = 1
	InputStreamOptions_ReadAhead InputStreamOptions = 2
)

// enum
type UnicodeEncoding int32

const (
	UnicodeEncoding_Utf8    UnicodeEncoding = 0
	UnicodeEncoding_Utf16LE UnicodeEncoding = 1
	UnicodeEncoding_Utf16BE UnicodeEncoding = 2
)

// interfaces

// 905A0FE0-BC53-11DF-8C49-001E4FC686DA
var IID_IBuffer = syscall.GUID{0x905A0FE0, 0xBC53, 0x11DF,
	[8]byte{0x8C, 0x49, 0x00, 0x1E, 0x4F, 0xC6, 0x86, 0xDA}}

type IBufferInterface interface {
	win32.IInspectableInterface
	Get_Capacity() uint32
	Get_Length() uint32
	Put_Length(value uint32)
}

type IBufferVtbl struct {
	win32.IInspectableVtbl
	Get_Capacity uintptr
	Get_Length   uintptr
	Put_Length   uintptr
}

type IBuffer struct {
	win32.IInspectable
}

func (this *IBuffer) Vtbl() *IBufferVtbl {
	return (*IBufferVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBuffer) Get_Capacity() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Capacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBuffer) Get_Length() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Length, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBuffer) Put_Length(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Length, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 71AF914D-C10F-484B-BC50-14BC623B3A27
var IID_IBufferFactory = syscall.GUID{0x71AF914D, 0xC10F, 0x484B,
	[8]byte{0xBC, 0x50, 0x14, 0xBC, 0x62, 0x3B, 0x3A, 0x27}}

type IBufferFactoryInterface interface {
	win32.IInspectableInterface
	Create(capacity uint32) *IBuffer
}

type IBufferFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IBufferFactory struct {
	win32.IInspectable
}

func (this *IBufferFactory) Vtbl() *IBufferFactoryVtbl {
	return (*IBufferFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBufferFactory) Create(capacity uint32) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(capacity), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E901E65B-D716-475A-A90A-AF7229B1E741
var IID_IBufferStatics = syscall.GUID{0xE901E65B, 0xD716, 0x475A,
	[8]byte{0xA9, 0x0A, 0xAF, 0x72, 0x29, 0xB1, 0xE7, 0x41}}

type IBufferStaticsInterface interface {
	win32.IInspectableInterface
	CreateCopyFromMemoryBuffer(input *IMemoryBuffer) *IBuffer
	CreateMemoryBufferOverIBuffer(input *IBuffer) *IMemoryBuffer
}

type IBufferStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateCopyFromMemoryBuffer    uintptr
	CreateMemoryBufferOverIBuffer uintptr
}

type IBufferStatics struct {
	win32.IInspectable
}

func (this *IBufferStatics) Vtbl() *IBufferStaticsVtbl {
	return (*IBufferStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBufferStatics) CreateCopyFromMemoryBuffer(input *IMemoryBuffer) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCopyFromMemoryBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(input)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBufferStatics) CreateMemoryBufferOverIBuffer(input *IBuffer) *IMemoryBuffer {
	var _result *IMemoryBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateMemoryBufferOverIBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(input)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 97D098A5-3B99-4DE9-88A5-E11D2F50C795
var IID_IContentTypeProvider = syscall.GUID{0x97D098A5, 0x3B99, 0x4DE9,
	[8]byte{0x88, 0xA5, 0xE1, 0x1D, 0x2F, 0x50, 0xC7, 0x95}}

type IContentTypeProviderInterface interface {
	win32.IInspectableInterface
	Get_ContentType() string
}

type IContentTypeProviderVtbl struct {
	win32.IInspectableVtbl
	Get_ContentType uintptr
}

type IContentTypeProvider struct {
	win32.IInspectable
}

func (this *IContentTypeProvider) Vtbl() *IContentTypeProviderVtbl {
	return (*IContentTypeProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IContentTypeProvider) Get_ContentType() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// E2B50029-B4C1-4314-A4B8-FB813A2F275E
var IID_IDataReader = syscall.GUID{0xE2B50029, 0xB4C1, 0x4314,
	[8]byte{0xA4, 0xB8, 0xFB, 0x81, 0x3A, 0x2F, 0x27, 0x5E}}

type IDataReaderInterface interface {
	win32.IInspectableInterface
	Get_UnconsumedBufferLength() uint32
	Get_UnicodeEncoding() UnicodeEncoding
	Put_UnicodeEncoding(value UnicodeEncoding)
	Get_ByteOrder() ByteOrder
	Put_ByteOrder(value ByteOrder)
	Get_InputStreamOptions() InputStreamOptions
	Put_InputStreamOptions(value InputStreamOptions)
	ReadByte() byte
	ReadBytes(valueLength uint32, value *byte)
	ReadBuffer(length uint32) *IBuffer
	ReadBoolean() bool
	ReadGuid() syscall.GUID
	ReadInt16() int16
	ReadInt32() int32
	ReadInt64() int64
	ReadUInt16() uint16
	ReadUInt32() uint32
	ReadUInt64() uint64
	ReadSingle() float32
	ReadDouble() float64
	ReadString(codeUnitCount uint32) string
	ReadDateTime() DateTime
	ReadTimeSpan() TimeSpan
	LoadAsync(count uint32) *IAsyncOperation[uint32]
	DetachBuffer() *IBuffer
	DetachStream() *IInputStream
}

type IDataReaderVtbl struct {
	win32.IInspectableVtbl
	Get_UnconsumedBufferLength uintptr
	Get_UnicodeEncoding        uintptr
	Put_UnicodeEncoding        uintptr
	Get_ByteOrder              uintptr
	Put_ByteOrder              uintptr
	Get_InputStreamOptions     uintptr
	Put_InputStreamOptions     uintptr
	ReadByte                   uintptr
	ReadBytes                  uintptr
	ReadBuffer                 uintptr
	ReadBoolean                uintptr
	ReadGuid                   uintptr
	ReadInt16                  uintptr
	ReadInt32                  uintptr
	ReadInt64                  uintptr
	ReadUInt16                 uintptr
	ReadUInt32                 uintptr
	ReadUInt64                 uintptr
	ReadSingle                 uintptr
	ReadDouble                 uintptr
	ReadString                 uintptr
	ReadDateTime               uintptr
	ReadTimeSpan               uintptr
	LoadAsync                  uintptr
	DetachBuffer               uintptr
	DetachStream               uintptr
}

type IDataReader struct {
	win32.IInspectable
}

func (this *IDataReader) Vtbl() *IDataReaderVtbl {
	return (*IDataReaderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataReader) Get_UnconsumedBufferLength() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UnconsumedBufferLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) Get_UnicodeEncoding() UnicodeEncoding {
	var _result UnicodeEncoding
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UnicodeEncoding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) Put_UnicodeEncoding(value UnicodeEncoding) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UnicodeEncoding, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataReader) Get_ByteOrder() ByteOrder {
	var _result ByteOrder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ByteOrder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) Put_ByteOrder(value ByteOrder) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ByteOrder, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataReader) Get_InputStreamOptions() InputStreamOptions {
	var _result InputStreamOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InputStreamOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) Put_InputStreamOptions(value InputStreamOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InputStreamOptions, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataReader) ReadByte() byte {
	var _result byte
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadByte, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) ReadBytes(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadBytes, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDataReader) ReadBuffer(length uint32) *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadBuffer, uintptr(unsafe.Pointer(this)), uintptr(length), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataReader) ReadBoolean() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadBoolean, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) ReadGuid() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadGuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) ReadInt16() int16 {
	var _result int16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadInt16, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) ReadInt32() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadInt32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) ReadInt64() int64 {
	var _result int64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadInt64, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) ReadUInt16() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadUInt16, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) ReadUInt32() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadUInt32, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) ReadUInt64() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadUInt64, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) ReadSingle() float32 {
	var _result float32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadSingle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) ReadDouble() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadDouble, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) ReadString(codeUnitCount uint32) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadString, uintptr(unsafe.Pointer(this)), uintptr(codeUnitCount), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDataReader) ReadDateTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadDateTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) ReadTimeSpan() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadTimeSpan, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataReader) LoadAsync(count uint32) *IAsyncOperation[uint32] {
	var _result *IAsyncOperation[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadAsync, uintptr(unsafe.Pointer(this)), uintptr(count), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataReader) DetachBuffer() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DetachBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataReader) DetachStream() *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DetachStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D7527847-57DA-4E15-914C-06806699A098
var IID_IDataReaderFactory = syscall.GUID{0xD7527847, 0x57DA, 0x4E15,
	[8]byte{0x91, 0x4C, 0x06, 0x80, 0x66, 0x99, 0xA0, 0x98}}

type IDataReaderFactoryInterface interface {
	win32.IInspectableInterface
	CreateDataReader(inputStream *IInputStream) *IDataReader
}

type IDataReaderFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateDataReader uintptr
}

type IDataReaderFactory struct {
	win32.IInspectable
}

func (this *IDataReaderFactory) Vtbl() *IDataReaderFactoryVtbl {
	return (*IDataReaderFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataReaderFactory) CreateDataReader(inputStream *IInputStream) *IDataReader {
	var _result *IDataReader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDataReader, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inputStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 11FCBFC8-F93A-471B-B121-F379E349313C
var IID_IDataReaderStatics = syscall.GUID{0x11FCBFC8, 0xF93A, 0x471B,
	[8]byte{0xB1, 0x21, 0xF3, 0x79, 0xE3, 0x49, 0x31, 0x3C}}

type IDataReaderStaticsInterface interface {
	win32.IInspectableInterface
	FromBuffer(buffer *IBuffer) *IDataReader
}

type IDataReaderStaticsVtbl struct {
	win32.IInspectableVtbl
	FromBuffer uintptr
}

type IDataReaderStatics struct {
	win32.IInspectable
}

func (this *IDataReaderStatics) Vtbl() *IDataReaderStaticsVtbl {
	return (*IDataReaderStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataReaderStatics) FromBuffer(buffer *IBuffer) *IDataReader {
	var _result *IDataReader
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 64B89265-D341-4922-B38A-DD4AF8808C4E
var IID_IDataWriter = syscall.GUID{0x64B89265, 0xD341, 0x4922,
	[8]byte{0xB3, 0x8A, 0xDD, 0x4A, 0xF8, 0x80, 0x8C, 0x4E}}

type IDataWriterInterface interface {
	win32.IInspectableInterface
	Get_UnstoredBufferLength() uint32
	Get_UnicodeEncoding() UnicodeEncoding
	Put_UnicodeEncoding(value UnicodeEncoding)
	Get_ByteOrder() ByteOrder
	Put_ByteOrder(value ByteOrder)
	WriteByte(value byte)
	WriteBytes(valueLength uint32, value *byte)
	WriteBuffer(buffer *IBuffer)
	WriteBufferRange(buffer *IBuffer, start uint32, count uint32)
	WriteBoolean(value bool)
	WriteGuid(value syscall.GUID)
	WriteInt16(value int16)
	WriteInt32(value int32)
	WriteInt64(value int64)
	WriteUInt16(value uint16)
	WriteUInt32(value uint32)
	WriteUInt64(value uint64)
	WriteSingle(value float32)
	WriteDouble(value float64)
	WriteDateTime(value DateTime)
	WriteTimeSpan(value TimeSpan)
	WriteString(value string) uint32
	MeasureString(value string) uint32
	StoreAsync() *IAsyncOperation[uint32]
	FlushAsync() *IAsyncOperation[bool]
	DetachBuffer() *IBuffer
	DetachStream() *IOutputStream
}

type IDataWriterVtbl struct {
	win32.IInspectableVtbl
	Get_UnstoredBufferLength uintptr
	Get_UnicodeEncoding      uintptr
	Put_UnicodeEncoding      uintptr
	Get_ByteOrder            uintptr
	Put_ByteOrder            uintptr
	WriteByte                uintptr
	WriteBytes               uintptr
	WriteBuffer              uintptr
	WriteBufferRange         uintptr
	WriteBoolean             uintptr
	WriteGuid                uintptr
	WriteInt16               uintptr
	WriteInt32               uintptr
	WriteInt64               uintptr
	WriteUInt16              uintptr
	WriteUInt32              uintptr
	WriteUInt64              uintptr
	WriteSingle              uintptr
	WriteDouble              uintptr
	WriteDateTime            uintptr
	WriteTimeSpan            uintptr
	WriteString              uintptr
	MeasureString            uintptr
	StoreAsync               uintptr
	FlushAsync               uintptr
	DetachBuffer             uintptr
	DetachStream             uintptr
}

type IDataWriter struct {
	win32.IInspectable
}

func (this *IDataWriter) Vtbl() *IDataWriterVtbl {
	return (*IDataWriterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataWriter) Get_UnstoredBufferLength() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UnstoredBufferLength, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataWriter) Get_UnicodeEncoding() UnicodeEncoding {
	var _result UnicodeEncoding
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UnicodeEncoding, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataWriter) Put_UnicodeEncoding(value UnicodeEncoding) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UnicodeEncoding, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataWriter) Get_ByteOrder() ByteOrder {
	var _result ByteOrder
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ByteOrder, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataWriter) Put_ByteOrder(value ByteOrder) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ByteOrder, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataWriter) WriteByte(value byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteByte, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataWriter) WriteBytes(valueLength uint32, value *byte) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteBytes, uintptr(unsafe.Pointer(this)), uintptr(valueLength), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IDataWriter) WriteBuffer(buffer *IBuffer) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)))
	_ = _hr
}

func (this *IDataWriter) WriteBufferRange(buffer *IBuffer, start uint32, count uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteBufferRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)), uintptr(start), uintptr(count))
	_ = _hr
}

func (this *IDataWriter) WriteBoolean(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteBoolean, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IDataWriter) WriteGuid(value syscall.GUID) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteGuid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IDataWriter) WriteInt16(value int16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteInt16, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataWriter) WriteInt32(value int32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteInt32, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataWriter) WriteInt64(value int64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteInt64, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataWriter) WriteUInt16(value uint16) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteUInt16, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataWriter) WriteUInt32(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteUInt32, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataWriter) WriteUInt64(value uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteUInt64, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataWriter) WriteSingle(value float32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteSingle, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataWriter) WriteDouble(value float64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteDouble, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IDataWriter) WriteDateTime(value DateTime) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteDateTime, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IDataWriter) WriteTimeSpan(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteTimeSpan, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IDataWriter) WriteString(value string) uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteString, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataWriter) MeasureString(value string) uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MeasureString, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataWriter) StoreAsync() *IAsyncOperation[uint32] {
	var _result *IAsyncOperation[uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StoreAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataWriter) FlushAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FlushAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataWriter) DetachBuffer() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DetachBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataWriter) DetachStream() *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DetachStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 338C67C2-8B84-4C2B-9C50-7B8767847A1F
var IID_IDataWriterFactory = syscall.GUID{0x338C67C2, 0x8B84, 0x4C2B,
	[8]byte{0x9C, 0x50, 0x7B, 0x87, 0x67, 0x84, 0x7A, 0x1F}}

type IDataWriterFactoryInterface interface {
	win32.IInspectableInterface
	CreateDataWriter(outputStream *IOutputStream) *IDataWriter
}

type IDataWriterFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateDataWriter uintptr
}

type IDataWriterFactory struct {
	win32.IInspectable
}

func (this *IDataWriterFactory) Vtbl() *IDataWriterFactoryVtbl {
	return (*IDataWriterFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataWriterFactory) CreateDataWriter(outputStream *IOutputStream) *IDataWriter {
	var _result *IDataWriter
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateDataWriter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(outputStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 73550107-3B57-4B5D-8345-554D2FC621F0
var IID_IFileRandomAccessStreamStatics = syscall.GUID{0x73550107, 0x3B57, 0x4B5D,
	[8]byte{0x83, 0x45, 0x55, 0x4D, 0x2F, 0xC6, 0x21, 0xF0}}

type IFileRandomAccessStreamStaticsInterface interface {
	win32.IInspectableInterface
	OpenAsync(filePath string, accessMode FileAccessMode) *IAsyncOperation[*IRandomAccessStream]
	OpenWithOptionsAsync(filePath string, accessMode FileAccessMode, sharingOptions StorageOpenOptions, openDisposition FileOpenDisposition) *IAsyncOperation[*IRandomAccessStream]
	OpenTransactedWriteAsync(filePath string) *IAsyncOperation[*IStorageStreamTransaction]
	OpenTransactedWriteWithOptionsAsync(filePath string, openOptions StorageOpenOptions, openDisposition FileOpenDisposition) *IAsyncOperation[*IStorageStreamTransaction]
	OpenForUserAsync(user *IUser, filePath string, accessMode FileAccessMode) *IAsyncOperation[*IRandomAccessStream]
	OpenForUserWithOptionsAsync(user *IUser, filePath string, accessMode FileAccessMode, sharingOptions StorageOpenOptions, openDisposition FileOpenDisposition) *IAsyncOperation[*IRandomAccessStream]
	OpenTransactedWriteForUserAsync(user *IUser, filePath string) *IAsyncOperation[*IStorageStreamTransaction]
	OpenTransactedWriteForUserWithOptionsAsync(user *IUser, filePath string, openOptions StorageOpenOptions, openDisposition FileOpenDisposition) *IAsyncOperation[*IStorageStreamTransaction]
}

type IFileRandomAccessStreamStaticsVtbl struct {
	win32.IInspectableVtbl
	OpenAsync                                  uintptr
	OpenWithOptionsAsync                       uintptr
	OpenTransactedWriteAsync                   uintptr
	OpenTransactedWriteWithOptionsAsync        uintptr
	OpenForUserAsync                           uintptr
	OpenForUserWithOptionsAsync                uintptr
	OpenTransactedWriteForUserAsync            uintptr
	OpenTransactedWriteForUserWithOptionsAsync uintptr
}

type IFileRandomAccessStreamStatics struct {
	win32.IInspectable
}

func (this *IFileRandomAccessStreamStatics) Vtbl() *IFileRandomAccessStreamStaticsVtbl {
	return (*IFileRandomAccessStreamStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileRandomAccessStreamStatics) OpenAsync(filePath string, accessMode FileAccessMode) *IAsyncOperation[*IRandomAccessStream] {
	var _result *IAsyncOperation[*IRandomAccessStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenAsync, uintptr(unsafe.Pointer(this)), NewHStr(filePath).Ptr, uintptr(accessMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileRandomAccessStreamStatics) OpenWithOptionsAsync(filePath string, accessMode FileAccessMode, sharingOptions StorageOpenOptions, openDisposition FileOpenDisposition) *IAsyncOperation[*IRandomAccessStream] {
	var _result *IAsyncOperation[*IRandomAccessStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenWithOptionsAsync, uintptr(unsafe.Pointer(this)), NewHStr(filePath).Ptr, uintptr(accessMode), uintptr(sharingOptions), uintptr(openDisposition), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileRandomAccessStreamStatics) OpenTransactedWriteAsync(filePath string) *IAsyncOperation[*IStorageStreamTransaction] {
	var _result *IAsyncOperation[*IStorageStreamTransaction]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenTransactedWriteAsync, uintptr(unsafe.Pointer(this)), NewHStr(filePath).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileRandomAccessStreamStatics) OpenTransactedWriteWithOptionsAsync(filePath string, openOptions StorageOpenOptions, openDisposition FileOpenDisposition) *IAsyncOperation[*IStorageStreamTransaction] {
	var _result *IAsyncOperation[*IStorageStreamTransaction]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenTransactedWriteWithOptionsAsync, uintptr(unsafe.Pointer(this)), NewHStr(filePath).Ptr, uintptr(openOptions), uintptr(openDisposition), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileRandomAccessStreamStatics) OpenForUserAsync(user *IUser, filePath string, accessMode FileAccessMode) *IAsyncOperation[*IRandomAccessStream] {
	var _result *IAsyncOperation[*IRandomAccessStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(filePath).Ptr, uintptr(accessMode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileRandomAccessStreamStatics) OpenForUserWithOptionsAsync(user *IUser, filePath string, accessMode FileAccessMode, sharingOptions StorageOpenOptions, openDisposition FileOpenDisposition) *IAsyncOperation[*IRandomAccessStream] {
	var _result *IAsyncOperation[*IRandomAccessStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenForUserWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(filePath).Ptr, uintptr(accessMode), uintptr(sharingOptions), uintptr(openDisposition), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileRandomAccessStreamStatics) OpenTransactedWriteForUserAsync(user *IUser, filePath string) *IAsyncOperation[*IStorageStreamTransaction] {
	var _result *IAsyncOperation[*IStorageStreamTransaction]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenTransactedWriteForUserAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(filePath).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileRandomAccessStreamStatics) OpenTransactedWriteForUserWithOptionsAsync(user *IUser, filePath string, openOptions StorageOpenOptions, openDisposition FileOpenDisposition) *IAsyncOperation[*IStorageStreamTransaction] {
	var _result *IAsyncOperation[*IStorageStreamTransaction]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenTransactedWriteForUserWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(user)), NewHStr(filePath).Ptr, uintptr(openOptions), uintptr(openDisposition), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 905A0FE2-BC53-11DF-8C49-001E4FC686DA
var IID_IInputStream = syscall.GUID{0x905A0FE2, 0xBC53, 0x11DF,
	[8]byte{0x8C, 0x49, 0x00, 0x1E, 0x4F, 0xC6, 0x86, 0xDA}}

type IInputStreamInterface interface {
	win32.IInspectableInterface
	ReadAsync(buffer *IBuffer, count uint32, options InputStreamOptions) *IAsyncOperationWithProgress[*IBuffer, uint32]
}

type IInputStreamVtbl struct {
	win32.IInspectableVtbl
	ReadAsync uintptr
}

type IInputStream struct {
	win32.IInspectable
}

func (this *IInputStream) Vtbl() *IInputStreamVtbl {
	return (*IInputStreamVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInputStream) ReadAsync(buffer *IBuffer, count uint32, options InputStreamOptions) *IAsyncOperationWithProgress[*IBuffer, uint32] {
	var _result *IAsyncOperationWithProgress[*IBuffer, uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ReadAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)), uintptr(count), uintptr(options), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 43929D18-5EC9-4B5A-919C-4205B0C804B6
var IID_IInputStreamReference = syscall.GUID{0x43929D18, 0x5EC9, 0x4B5A,
	[8]byte{0x91, 0x9C, 0x42, 0x05, 0xB0, 0xC8, 0x04, 0xB6}}

type IInputStreamReferenceInterface interface {
	win32.IInspectableInterface
	OpenSequentialReadAsync() *IAsyncOperation[*IInputStream]
}

type IInputStreamReferenceVtbl struct {
	win32.IInspectableVtbl
	OpenSequentialReadAsync uintptr
}

type IInputStreamReference struct {
	win32.IInspectable
}

func (this *IInputStreamReference) Vtbl() *IInputStreamReferenceVtbl {
	return (*IInputStreamReferenceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInputStreamReference) OpenSequentialReadAsync() *IAsyncOperation[*IInputStream] {
	var _result *IAsyncOperation[*IInputStream]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenSequentialReadAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 905A0FE6-BC53-11DF-8C49-001E4FC686DA
var IID_IOutputStream = syscall.GUID{0x905A0FE6, 0xBC53, 0x11DF,
	[8]byte{0x8C, 0x49, 0x00, 0x1E, 0x4F, 0xC6, 0x86, 0xDA}}

type IOutputStreamInterface interface {
	win32.IInspectableInterface
	WriteAsync(buffer *IBuffer) *IAsyncOperationWithProgress[uint32, uint32]
	FlushAsync() *IAsyncOperation[bool]
}

type IOutputStreamVtbl struct {
	win32.IInspectableVtbl
	WriteAsync uintptr
	FlushAsync uintptr
}

type IOutputStream struct {
	win32.IInspectable
}

func (this *IOutputStream) Vtbl() *IOutputStreamVtbl {
	return (*IOutputStreamVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IOutputStream) WriteAsync(buffer *IBuffer) *IAsyncOperationWithProgress[uint32, uint32] {
	var _result *IAsyncOperationWithProgress[uint32, uint32]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().WriteAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IOutputStream) FlushAsync() *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FlushAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 905A0FE1-BC53-11DF-8C49-001E4FC686DA
var IID_IRandomAccessStream = syscall.GUID{0x905A0FE1, 0xBC53, 0x11DF,
	[8]byte{0x8C, 0x49, 0x00, 0x1E, 0x4F, 0xC6, 0x86, 0xDA}}

type IRandomAccessStreamInterface interface {
	win32.IInspectableInterface
	Get_Size() uint64
	Put_Size(value uint64)
	GetInputStreamAt(position uint64) *IInputStream
	GetOutputStreamAt(position uint64) *IOutputStream
	Get_Position() uint64
	Seek(position uint64)
	CloneStream() *IRandomAccessStream
	Get_CanRead() bool
	Get_CanWrite() bool
}

type IRandomAccessStreamVtbl struct {
	win32.IInspectableVtbl
	Get_Size          uintptr
	Put_Size          uintptr
	GetInputStreamAt  uintptr
	GetOutputStreamAt uintptr
	Get_Position      uintptr
	Seek              uintptr
	CloneStream       uintptr
	Get_CanRead       uintptr
	Get_CanWrite      uintptr
}

type IRandomAccessStream struct {
	win32.IInspectable
}

func (this *IRandomAccessStream) Vtbl() *IRandomAccessStreamVtbl {
	return (*IRandomAccessStreamVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRandomAccessStream) Get_Size() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRandomAccessStream) Put_Size(value uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Size, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IRandomAccessStream) GetInputStreamAt(position uint64) *IInputStream {
	var _result *IInputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetInputStreamAt, uintptr(unsafe.Pointer(this)), uintptr(position), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRandomAccessStream) GetOutputStreamAt(position uint64) *IOutputStream {
	var _result *IOutputStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetOutputStreamAt, uintptr(unsafe.Pointer(this)), uintptr(position), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRandomAccessStream) Get_Position() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Position, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRandomAccessStream) Seek(position uint64) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Seek, uintptr(unsafe.Pointer(this)), uintptr(position))
	_ = _hr
}

func (this *IRandomAccessStream) CloneStream() *IRandomAccessStream {
	var _result *IRandomAccessStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CloneStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRandomAccessStream) Get_CanRead() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanRead, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRandomAccessStream) Get_CanWrite() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CanWrite, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 33EE3134-1DD6-4E3A-8067-D1C162E8642B
var IID_IRandomAccessStreamReference = syscall.GUID{0x33EE3134, 0x1DD6, 0x4E3A,
	[8]byte{0x80, 0x67, 0xD1, 0xC1, 0x62, 0xE8, 0x64, 0x2B}}

type IRandomAccessStreamReferenceInterface interface {
	win32.IInspectableInterface
	OpenReadAsync() *IAsyncOperation[*IRandomAccessStreamWithContentType]
}

type IRandomAccessStreamReferenceVtbl struct {
	win32.IInspectableVtbl
	OpenReadAsync uintptr
}

type IRandomAccessStreamReference struct {
	win32.IInspectable
}

func (this *IRandomAccessStreamReference) Vtbl() *IRandomAccessStreamReferenceVtbl {
	return (*IRandomAccessStreamReferenceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRandomAccessStreamReference) OpenReadAsync() *IAsyncOperation[*IRandomAccessStreamWithContentType] {
	var _result *IAsyncOperation[*IRandomAccessStreamWithContentType]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().OpenReadAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 857309DC-3FBF-4E7D-986F-EF3B1A07A964
var IID_IRandomAccessStreamReferenceStatics = syscall.GUID{0x857309DC, 0x3FBF, 0x4E7D,
	[8]byte{0x98, 0x6F, 0xEF, 0x3B, 0x1A, 0x07, 0xA9, 0x64}}

type IRandomAccessStreamReferenceStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromFile(file *IStorageFile) *IRandomAccessStreamReference
	CreateFromUri(uri *IUriRuntimeClass) *IRandomAccessStreamReference
	CreateFromStream(stream *IRandomAccessStream) *IRandomAccessStreamReference
}

type IRandomAccessStreamReferenceStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromFile   uintptr
	CreateFromUri    uintptr
	CreateFromStream uintptr
}

type IRandomAccessStreamReferenceStatics struct {
	win32.IInspectable
}

func (this *IRandomAccessStreamReferenceStatics) Vtbl() *IRandomAccessStreamReferenceStaticsVtbl {
	return (*IRandomAccessStreamReferenceStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRandomAccessStreamReferenceStatics) CreateFromFile(file *IStorageFile) *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRandomAccessStreamReferenceStatics) CreateFromUri(uri *IUriRuntimeClass) *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRandomAccessStreamReferenceStatics) CreateFromStream(stream *IRandomAccessStream) *IRandomAccessStreamReference {
	var _result *IRandomAccessStreamReference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 524CEDCF-6E29-4CE5-9573-6B753DB66C3A
var IID_IRandomAccessStreamStatics = syscall.GUID{0x524CEDCF, 0x6E29, 0x4CE5,
	[8]byte{0x95, 0x73, 0x6B, 0x75, 0x3D, 0xB6, 0x6C, 0x3A}}

type IRandomAccessStreamStaticsInterface interface {
	win32.IInspectableInterface
	CopyAsync(source *IInputStream, destination *IOutputStream) *IAsyncOperationWithProgress[uint64, uint64]
	CopySizeAsync(source *IInputStream, destination *IOutputStream, bytesToCopy uint64) *IAsyncOperationWithProgress[uint64, uint64]
	CopyAndCloseAsync(source *IInputStream, destination *IOutputStream) *IAsyncOperationWithProgress[uint64, uint64]
}

type IRandomAccessStreamStaticsVtbl struct {
	win32.IInspectableVtbl
	CopyAsync         uintptr
	CopySizeAsync     uintptr
	CopyAndCloseAsync uintptr
}

type IRandomAccessStreamStatics struct {
	win32.IInspectable
}

func (this *IRandomAccessStreamStatics) Vtbl() *IRandomAccessStreamStaticsVtbl {
	return (*IRandomAccessStreamStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRandomAccessStreamStatics) CopyAsync(source *IInputStream, destination *IOutputStream) *IAsyncOperationWithProgress[uint64, uint64] {
	var _result *IAsyncOperationWithProgress[uint64, uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(destination)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRandomAccessStreamStatics) CopySizeAsync(source *IInputStream, destination *IOutputStream, bytesToCopy uint64) *IAsyncOperationWithProgress[uint64, uint64] {
	var _result *IAsyncOperationWithProgress[uint64, uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopySizeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(destination)), uintptr(bytesToCopy), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRandomAccessStreamStatics) CopyAndCloseAsync(source *IInputStream, destination *IOutputStream) *IAsyncOperationWithProgress[uint64, uint64] {
	var _result *IAsyncOperationWithProgress[uint64, uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyAndCloseAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(destination)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CC254827-4B3D-438F-9232-10C76BC7E038
var IID_IRandomAccessStreamWithContentType = syscall.GUID{0xCC254827, 0x4B3D, 0x438F,
	[8]byte{0x92, 0x32, 0x10, 0xC7, 0x6B, 0xC7, 0xE0, 0x38}}

type IRandomAccessStreamWithContentTypeInterface interface {
	win32.IInspectableInterface
}

type IRandomAccessStreamWithContentTypeVtbl struct {
	win32.IInspectableVtbl
}

type IRandomAccessStreamWithContentType struct {
	win32.IInspectable
}

func (this *IRandomAccessStreamWithContentType) Vtbl() *IRandomAccessStreamWithContentTypeVtbl {
	return (*IRandomAccessStreamWithContentTypeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// classes

type Buffer struct {
	RtClass
	*IBuffer
}

func NewBuffer_Create(capacity uint32) *Buffer {
	hs := NewHStr("Windows.Storage.Streams.Buffer")
	var pFac *IBufferFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBufferFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IBuffer
	p = pFac.Create(capacity)
	result := &Buffer{
		RtClass: RtClass{PInspect: &p.IInspectable},
		IBuffer: p,
	}
	com.AddToScope(result)
	return result
}

func NewIBufferStatics() *IBufferStatics {
	var p *IBufferStatics
	hs := NewHStr("Windows.Storage.Streams.Buffer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IBufferStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type DataReader struct {
	RtClass
	*IDataReader
}

func NewDataReader_CreateDataReader(inputStream *IInputStream) *DataReader {
	hs := NewHStr("Windows.Storage.Streams.DataReader")
	var pFac *IDataReaderFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDataReaderFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IDataReader
	p = pFac.CreateDataReader(inputStream)
	result := &DataReader{
		RtClass:     RtClass{PInspect: &p.IInspectable},
		IDataReader: p,
	}
	com.AddToScope(result)
	return result
}

func NewIDataReaderStatics() *IDataReaderStatics {
	var p *IDataReaderStatics
	hs := NewHStr("Windows.Storage.Streams.DataReader")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDataReaderStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type DataWriter struct {
	RtClass
	*IDataWriter
}

func NewDataWriter() *DataWriter {
	hs := NewHStr("Windows.Storage.Streams.DataWriter")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &DataWriter{
		RtClass:     RtClass{PInspect: p},
		IDataWriter: (*IDataWriter)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewDataWriter_CreateDataWriter(outputStream *IOutputStream) *DataWriter {
	hs := NewHStr("Windows.Storage.Streams.DataWriter")
	var pFac *IDataWriterFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDataWriterFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IDataWriter
	p = pFac.CreateDataWriter(outputStream)
	result := &DataWriter{
		RtClass:     RtClass{PInspect: &p.IInspectable},
		IDataWriter: p,
	}
	com.AddToScope(result)
	return result
}
