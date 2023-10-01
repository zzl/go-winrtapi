package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
type PhotoOrientation int32

const (
	PhotoOrientation_Unspecified    PhotoOrientation = 0
	PhotoOrientation_Normal         PhotoOrientation = 1
	PhotoOrientation_FlipHorizontal PhotoOrientation = 2
	PhotoOrientation_Rotate180      PhotoOrientation = 3
	PhotoOrientation_FlipVertical   PhotoOrientation = 4
	PhotoOrientation_Transpose      PhotoOrientation = 5
	PhotoOrientation_Rotate270      PhotoOrientation = 6
	PhotoOrientation_Transverse     PhotoOrientation = 7
	PhotoOrientation_Rotate90       PhotoOrientation = 8
)

// enum
// flags
type PropertyPrefetchOptions uint32

const (
	PropertyPrefetchOptions_None               PropertyPrefetchOptions = 0
	PropertyPrefetchOptions_MusicProperties    PropertyPrefetchOptions = 1
	PropertyPrefetchOptions_VideoProperties    PropertyPrefetchOptions = 2
	PropertyPrefetchOptions_ImageProperties    PropertyPrefetchOptions = 4
	PropertyPrefetchOptions_DocumentProperties PropertyPrefetchOptions = 8
	PropertyPrefetchOptions_BasicProperties    PropertyPrefetchOptions = 16
)

// enum
type ThumbnailMode int32

const (
	ThumbnailMode_PicturesView  ThumbnailMode = 0
	ThumbnailMode_VideosView    ThumbnailMode = 1
	ThumbnailMode_MusicView     ThumbnailMode = 2
	ThumbnailMode_DocumentsView ThumbnailMode = 3
	ThumbnailMode_ListView      ThumbnailMode = 4
	ThumbnailMode_SingleItem    ThumbnailMode = 5
)

// enum
// flags
type ThumbnailOptions uint32

const (
	ThumbnailOptions_None               ThumbnailOptions = 0
	ThumbnailOptions_ReturnOnlyIfCached ThumbnailOptions = 1
	ThumbnailOptions_ResizeThumbnail    ThumbnailOptions = 2
	ThumbnailOptions_UseCurrentScale    ThumbnailOptions = 4
)

// enum
type ThumbnailType int32

const (
	ThumbnailType_Image ThumbnailType = 0
	ThumbnailType_Icon  ThumbnailType = 1
)

// enum
type VideoOrientation int32

const (
	VideoOrientation_Normal    VideoOrientation = 0
	VideoOrientation_Rotate90  VideoOrientation = 90
	VideoOrientation_Rotate180 VideoOrientation = 180
	VideoOrientation_Rotate270 VideoOrientation = 270
)

// interfaces

// D05D55DB-785E-4A66-BE02-9BEEC58AEA81
var IID_IBasicProperties = syscall.GUID{0xD05D55DB, 0x785E, 0x4A66,
	[8]byte{0xBE, 0x02, 0x9B, 0xEE, 0xC5, 0x8A, 0xEA, 0x81}}

type IBasicPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Size() uint64
	Get_DateModified() DateTime
	Get_ItemDate() DateTime
}

type IBasicPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Size         uintptr
	Get_DateModified uintptr
	Get_ItemDate     uintptr
}

type IBasicProperties struct {
	win32.IInspectable
}

func (this *IBasicProperties) Vtbl() *IBasicPropertiesVtbl {
	return (*IBasicPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBasicProperties) Get_Size() uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Size, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBasicProperties) Get_DateModified() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DateModified, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IBasicProperties) Get_ItemDate() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ItemDate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 7EAB19BC-1821-4923-B4A9-0AEA404D0070
var IID_IDocumentProperties = syscall.GUID{0x7EAB19BC, 0x1821, 0x4923,
	[8]byte{0xB4, 0xA9, 0x0A, 0xEA, 0x40, 0x4D, 0x00, 0x70}}

type IDocumentPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Author() *IVector[string]
	Get_Title() string
	Put_Title(value string)
	Get_Keywords() *IVector[string]
	Get_Comment() string
	Put_Comment(value string)
}

type IDocumentPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Author   uintptr
	Get_Title    uintptr
	Put_Title    uintptr
	Get_Keywords uintptr
	Get_Comment  uintptr
	Put_Comment  uintptr
}

type IDocumentProperties struct {
	win32.IInspectable
}

func (this *IDocumentProperties) Vtbl() *IDocumentPropertiesVtbl {
	return (*IDocumentPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDocumentProperties) Get_Author() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Author, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDocumentProperties) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDocumentProperties) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IDocumentProperties) Get_Keywords() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Keywords, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDocumentProperties) Get_Comment() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Comment, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IDocumentProperties) Put_Comment(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Comment, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

// 41493244-2524-4655-86A6-ED16F5FC716B
var IID_IGeotagHelperStatics = syscall.GUID{0x41493244, 0x2524, 0x4655,
	[8]byte{0x86, 0xA6, 0xED, 0x16, 0xF5, 0xFC, 0x71, 0x6B}}

type IGeotagHelperStaticsInterface interface {
	win32.IInspectableInterface
	GetGeotagAsync(file *IStorageFile) *IAsyncOperation[*IGeopoint]
	SetGeotagFromGeolocatorAsync(file *IStorageFile, geolocator *IGeolocator) *IAsyncAction
	SetGeotagAsync(file *IStorageFile, geopoint *IGeopoint) *IAsyncAction
}

type IGeotagHelperStaticsVtbl struct {
	win32.IInspectableVtbl
	GetGeotagAsync               uintptr
	SetGeotagFromGeolocatorAsync uintptr
	SetGeotagAsync               uintptr
}

type IGeotagHelperStatics struct {
	win32.IInspectable
}

func (this *IGeotagHelperStatics) Vtbl() *IGeotagHelperStaticsVtbl {
	return (*IGeotagHelperStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGeotagHelperStatics) GetGeotagAsync(file *IStorageFile) *IAsyncOperation[*IGeopoint] {
	var _result *IAsyncOperation[*IGeopoint]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetGeotagAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeotagHelperStatics) SetGeotagFromGeolocatorAsync(file *IStorageFile, geolocator *IGeolocator) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetGeotagFromGeolocatorAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(geolocator)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGeotagHelperStatics) SetGeotagAsync(file *IStorageFile, geopoint *IGeopoint) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetGeotagAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(geopoint)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 523C9424-FCFF-4275-AFEE-ECDB9AB47973
var IID_IImageProperties = syscall.GUID{0x523C9424, 0xFCFF, 0x4275,
	[8]byte{0xAF, 0xEE, 0xEC, 0xDB, 0x9A, 0xB4, 0x79, 0x73}}

type IImagePropertiesInterface interface {
	win32.IInspectableInterface
	Get_Rating() uint32
	Put_Rating(value uint32)
	Get_Keywords() *IVector[string]
	Get_DateTaken() DateTime
	Put_DateTaken(value DateTime)
	Get_Width() uint32
	Get_Height() uint32
	Get_Title() string
	Put_Title(value string)
	Get_Latitude() *IReference[float64]
	Get_Longitude() *IReference[float64]
	Get_CameraManufacturer() string
	Put_CameraManufacturer(value string)
	Get_CameraModel() string
	Put_CameraModel(value string)
	Get_Orientation() PhotoOrientation
	Get_PeopleNames() *IVectorView[string]
}

type IImagePropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Rating             uintptr
	Put_Rating             uintptr
	Get_Keywords           uintptr
	Get_DateTaken          uintptr
	Put_DateTaken          uintptr
	Get_Width              uintptr
	Get_Height             uintptr
	Get_Title              uintptr
	Put_Title              uintptr
	Get_Latitude           uintptr
	Get_Longitude          uintptr
	Get_CameraManufacturer uintptr
	Put_CameraManufacturer uintptr
	Get_CameraModel        uintptr
	Put_CameraModel        uintptr
	Get_Orientation        uintptr
	Get_PeopleNames        uintptr
}

type IImageProperties struct {
	win32.IInspectable
}

func (this *IImageProperties) Vtbl() *IImagePropertiesVtbl {
	return (*IImagePropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IImageProperties) Get_Rating() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rating, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageProperties) Put_Rating(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Rating, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IImageProperties) Get_Keywords() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Keywords, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IImageProperties) Get_DateTaken() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DateTaken, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageProperties) Put_DateTaken(value DateTime) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DateTaken, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IImageProperties) Get_Width() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Width, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageProperties) Get_Height() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageProperties) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IImageProperties) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IImageProperties) Get_Latitude() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Latitude, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IImageProperties) Get_Longitude() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Longitude, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IImageProperties) Get_CameraManufacturer() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraManufacturer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IImageProperties) Put_CameraManufacturer(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CameraManufacturer, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IImageProperties) Get_CameraModel() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CameraModel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IImageProperties) Put_CameraModel(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CameraModel, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IImageProperties) Get_Orientation() PhotoOrientation {
	var _result PhotoOrientation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IImageProperties) Get_PeopleNames() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PeopleNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BC8AAB62-66EC-419A-BC5D-CA65A4CB46DA
var IID_IMusicProperties = syscall.GUID{0xBC8AAB62, 0x66EC, 0x419A,
	[8]byte{0xBC, 0x5D, 0xCA, 0x65, 0xA4, 0xCB, 0x46, 0xDA}}

type IMusicPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Album() string
	Put_Album(value string)
	Get_Artist() string
	Put_Artist(value string)
	Get_Genre() *IVector[string]
	Get_TrackNumber() uint32
	Put_TrackNumber(value uint32)
	Get_Title() string
	Put_Title(value string)
	Get_Rating() uint32
	Put_Rating(value uint32)
	Get_Duration() TimeSpan
	Get_Bitrate() uint32
	Get_AlbumArtist() string
	Put_AlbumArtist(value string)
	Get_Composers() *IVector[string]
	Get_Conductors() *IVector[string]
	Get_Subtitle() string
	Put_Subtitle(value string)
	Get_Producers() *IVector[string]
	Get_Publisher() string
	Put_Publisher(value string)
	Get_Writers() *IVector[string]
	Get_Year() uint32
	Put_Year(value uint32)
}

type IMusicPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Album       uintptr
	Put_Album       uintptr
	Get_Artist      uintptr
	Put_Artist      uintptr
	Get_Genre       uintptr
	Get_TrackNumber uintptr
	Put_TrackNumber uintptr
	Get_Title       uintptr
	Put_Title       uintptr
	Get_Rating      uintptr
	Put_Rating      uintptr
	Get_Duration    uintptr
	Get_Bitrate     uintptr
	Get_AlbumArtist uintptr
	Put_AlbumArtist uintptr
	Get_Composers   uintptr
	Get_Conductors  uintptr
	Get_Subtitle    uintptr
	Put_Subtitle    uintptr
	Get_Producers   uintptr
	Get_Publisher   uintptr
	Put_Publisher   uintptr
	Get_Writers     uintptr
	Get_Year        uintptr
	Put_Year        uintptr
}

type IMusicProperties struct {
	win32.IInspectable
}

func (this *IMusicProperties) Vtbl() *IMusicPropertiesVtbl {
	return (*IMusicPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMusicProperties) Get_Album() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Album, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMusicProperties) Put_Album(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Album, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMusicProperties) Get_Artist() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Artist, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMusicProperties) Put_Artist(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Artist, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMusicProperties) Get_Genre() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Genre, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMusicProperties) Get_TrackNumber() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TrackNumber, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMusicProperties) Put_TrackNumber(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TrackNumber, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMusicProperties) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMusicProperties) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMusicProperties) Get_Rating() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rating, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMusicProperties) Put_Rating(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Rating, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IMusicProperties) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMusicProperties) Get_Bitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMusicProperties) Get_AlbumArtist() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AlbumArtist, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMusicProperties) Put_AlbumArtist(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AlbumArtist, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMusicProperties) Get_Composers() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Composers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMusicProperties) Get_Conductors() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Conductors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMusicProperties) Get_Subtitle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subtitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMusicProperties) Put_Subtitle(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Subtitle, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMusicProperties) Get_Producers() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Producers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMusicProperties) Get_Publisher() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Publisher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IMusicProperties) Put_Publisher(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Publisher, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IMusicProperties) Get_Writers() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Writers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IMusicProperties) Get_Year() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Year, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IMusicProperties) Put_Year(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Year, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// 05294BAD-BC38-48BF-85D7-770E0E2AE0BA
var IID_IStorageItemContentProperties = syscall.GUID{0x05294BAD, 0xBC38, 0x48BF,
	[8]byte{0x85, 0xD7, 0x77, 0x0E, 0x0E, 0x2A, 0xE0, 0xBA}}

type IStorageItemContentPropertiesInterface interface {
	win32.IInspectableInterface
	GetMusicPropertiesAsync() *IAsyncOperation[*IMusicProperties]
	GetVideoPropertiesAsync() *IAsyncOperation[*IVideoProperties]
	GetImagePropertiesAsync() *IAsyncOperation[*IImageProperties]
	GetDocumentPropertiesAsync() *IAsyncOperation[*IDocumentProperties]
}

type IStorageItemContentPropertiesVtbl struct {
	win32.IInspectableVtbl
	GetMusicPropertiesAsync    uintptr
	GetVideoPropertiesAsync    uintptr
	GetImagePropertiesAsync    uintptr
	GetDocumentPropertiesAsync uintptr
}

type IStorageItemContentProperties struct {
	win32.IInspectable
}

func (this *IStorageItemContentProperties) Vtbl() *IStorageItemContentPropertiesVtbl {
	return (*IStorageItemContentPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageItemContentProperties) GetMusicPropertiesAsync() *IAsyncOperation[*IMusicProperties] {
	var _result *IAsyncOperation[*IMusicProperties]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMusicPropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItemContentProperties) GetVideoPropertiesAsync() *IAsyncOperation[*IVideoProperties] {
	var _result *IAsyncOperation[*IVideoProperties]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetVideoPropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItemContentProperties) GetImagePropertiesAsync() *IAsyncOperation[*IImageProperties] {
	var _result *IAsyncOperation[*IImageProperties]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetImagePropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItemContentProperties) GetDocumentPropertiesAsync() *IAsyncOperation[*IDocumentProperties] {
	var _result *IAsyncOperation[*IDocumentProperties]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDocumentPropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C54361B2-54CD-432B-BDBC-4B19C4B470D7
var IID_IStorageItemExtraProperties = syscall.GUID{0xC54361B2, 0x54CD, 0x432B,
	[8]byte{0xBD, 0xBC, 0x4B, 0x19, 0xC4, 0xB4, 0x70, 0xD7}}

type IStorageItemExtraPropertiesInterface interface {
	win32.IInspectableInterface
	RetrievePropertiesAsync(propertiesToRetrieve *IIterable[string]) *IAsyncOperation[*IMap[string, interface{}]]
	SavePropertiesAsync(propertiesToSave *IIterable[*IKeyValuePair[string, interface{}]]) *IAsyncAction
	SavePropertiesAsyncOverloadDefault() *IAsyncAction
}

type IStorageItemExtraPropertiesVtbl struct {
	win32.IInspectableVtbl
	RetrievePropertiesAsync            uintptr
	SavePropertiesAsync                uintptr
	SavePropertiesAsyncOverloadDefault uintptr
}

type IStorageItemExtraProperties struct {
	win32.IInspectable
}

func (this *IStorageItemExtraProperties) Vtbl() *IStorageItemExtraPropertiesVtbl {
	return (*IStorageItemExtraPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorageItemExtraProperties) RetrievePropertiesAsync(propertiesToRetrieve *IIterable[string]) *IAsyncOperation[*IMap[string, interface{}]] {
	var _result *IAsyncOperation[*IMap[string, interface{}]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RetrievePropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propertiesToRetrieve)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItemExtraProperties) SavePropertiesAsync(propertiesToSave *IIterable[*IKeyValuePair[string, interface{}]]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SavePropertiesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propertiesToSave)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStorageItemExtraProperties) SavePropertiesAsyncOverloadDefault() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SavePropertiesAsyncOverloadDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 693DD42F-DBE7-49B5-B3B3-2893AC5D3423
var IID_IThumbnailProperties = syscall.GUID{0x693DD42F, 0xDBE7, 0x49B5,
	[8]byte{0xB3, 0xB3, 0x28, 0x93, 0xAC, 0x5D, 0x34, 0x23}}

type IThumbnailPropertiesInterface interface {
	win32.IInspectableInterface
	Get_OriginalWidth() uint32
	Get_OriginalHeight() uint32
	Get_ReturnedSmallerCachedSize() bool
	Get_Type() ThumbnailType
}

type IThumbnailPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_OriginalWidth             uintptr
	Get_OriginalHeight            uintptr
	Get_ReturnedSmallerCachedSize uintptr
	Get_Type                      uintptr
}

type IThumbnailProperties struct {
	win32.IInspectable
}

func (this *IThumbnailProperties) Vtbl() *IThumbnailPropertiesVtbl {
	return (*IThumbnailPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IThumbnailProperties) Get_OriginalWidth() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OriginalWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IThumbnailProperties) Get_OriginalHeight() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OriginalHeight, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IThumbnailProperties) Get_ReturnedSmallerCachedSize() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ReturnedSmallerCachedSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IThumbnailProperties) Get_Type() ThumbnailType {
	var _result ThumbnailType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 719AE507-68DE-4DB8-97DE-49998C059F2F
var IID_IVideoProperties = syscall.GUID{0x719AE507, 0x68DE, 0x4DB8,
	[8]byte{0x97, 0xDE, 0x49, 0x99, 0x8C, 0x05, 0x9F, 0x2F}}

type IVideoPropertiesInterface interface {
	win32.IInspectableInterface
	Get_Rating() uint32
	Put_Rating(value uint32)
	Get_Keywords() *IVector[string]
	Get_Width() uint32
	Get_Height() uint32
	Get_Duration() TimeSpan
	Get_Latitude() *IReference[float64]
	Get_Longitude() *IReference[float64]
	Get_Title() string
	Put_Title(value string)
	Get_Subtitle() string
	Put_Subtitle(value string)
	Get_Producers() *IVector[string]
	Get_Publisher() string
	Put_Publisher(value string)
	Get_Writers() *IVector[string]
	Get_Year() uint32
	Put_Year(value uint32)
	Get_Bitrate() uint32
	Get_Directors() *IVector[string]
	Get_Orientation() VideoOrientation
}

type IVideoPropertiesVtbl struct {
	win32.IInspectableVtbl
	Get_Rating      uintptr
	Put_Rating      uintptr
	Get_Keywords    uintptr
	Get_Width       uintptr
	Get_Height      uintptr
	Get_Duration    uintptr
	Get_Latitude    uintptr
	Get_Longitude   uintptr
	Get_Title       uintptr
	Put_Title       uintptr
	Get_Subtitle    uintptr
	Put_Subtitle    uintptr
	Get_Producers   uintptr
	Get_Publisher   uintptr
	Put_Publisher   uintptr
	Get_Writers     uintptr
	Get_Year        uintptr
	Put_Year        uintptr
	Get_Bitrate     uintptr
	Get_Directors   uintptr
	Get_Orientation uintptr
}

type IVideoProperties struct {
	win32.IInspectable
}

func (this *IVideoProperties) Vtbl() *IVideoPropertiesVtbl {
	return (*IVideoPropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVideoProperties) Get_Rating() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Rating, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoProperties) Put_Rating(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Rating, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVideoProperties) Get_Keywords() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Keywords, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoProperties) Get_Width() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Width, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoProperties) Get_Height() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Height, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoProperties) Get_Duration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Duration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoProperties) Get_Latitude() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Latitude, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoProperties) Get_Longitude() *IReference[float64] {
	var _result *IReference[float64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Longitude, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoProperties) Get_Title() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Title, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVideoProperties) Put_Title(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Title, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IVideoProperties) Get_Subtitle() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Subtitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVideoProperties) Put_Subtitle(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Subtitle, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IVideoProperties) Get_Producers() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Producers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoProperties) Get_Publisher() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Publisher, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVideoProperties) Put_Publisher(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Publisher, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IVideoProperties) Get_Writers() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Writers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoProperties) Get_Year() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Year, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoProperties) Put_Year(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Year, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IVideoProperties) Get_Bitrate() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Bitrate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IVideoProperties) Get_Directors() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Directors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVideoProperties) Get_Orientation() VideoOrientation {
	var _result VideoOrientation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Orientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type GeotagHelper struct {
	RtClass
}

func NewIGeotagHelperStatics() *IGeotagHelperStatics {
	var p *IGeotagHelperStatics
	hs := NewHStr("Windows.Storage.FileProperties.GeotagHelper")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGeotagHelperStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
