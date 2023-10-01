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
type DataProtectionStatus int32

const (
	DataProtectionStatus_ProtectedToOtherIdentity DataProtectionStatus = 0
	DataProtectionStatus_Protected                DataProtectionStatus = 1
	DataProtectionStatus_Revoked                  DataProtectionStatus = 2
	DataProtectionStatus_Unprotected              DataProtectionStatus = 3
	DataProtectionStatus_LicenseExpired           DataProtectionStatus = 4
	DataProtectionStatus_AccessSuspended          DataProtectionStatus = 5
)

// enum
type EnforcementLevel int32

const (
	EnforcementLevel_NoProtection EnforcementLevel = 0
	EnforcementLevel_Silent       EnforcementLevel = 1
	EnforcementLevel_Override     EnforcementLevel = 2
	EnforcementLevel_Block        EnforcementLevel = 3
)

// enum
type FileProtectionStatus int32

const (
	FileProtectionStatus_Undetermined               FileProtectionStatus = 0
	FileProtectionStatus_Unknown                    FileProtectionStatus = 0
	FileProtectionStatus_Unprotected                FileProtectionStatus = 1
	FileProtectionStatus_Revoked                    FileProtectionStatus = 2
	FileProtectionStatus_Protected                  FileProtectionStatus = 3
	FileProtectionStatus_ProtectedByOtherUser       FileProtectionStatus = 4
	FileProtectionStatus_ProtectedToOtherEnterprise FileProtectionStatus = 5
	FileProtectionStatus_NotProtectable             FileProtectionStatus = 6
	FileProtectionStatus_ProtectedToOtherIdentity   FileProtectionStatus = 7
	FileProtectionStatus_LicenseExpired             FileProtectionStatus = 8
	FileProtectionStatus_AccessSuspended            FileProtectionStatus = 9
	FileProtectionStatus_FileInUse                  FileProtectionStatus = 10
)

// enum
type ProtectedImportExportStatus int32

const (
	ProtectedImportExportStatus_Ok                       ProtectedImportExportStatus = 0
	ProtectedImportExportStatus_Undetermined             ProtectedImportExportStatus = 1
	ProtectedImportExportStatus_Unprotected              ProtectedImportExportStatus = 2
	ProtectedImportExportStatus_Revoked                  ProtectedImportExportStatus = 3
	ProtectedImportExportStatus_NotRoamable              ProtectedImportExportStatus = 4
	ProtectedImportExportStatus_ProtectedToOtherIdentity ProtectedImportExportStatus = 5
	ProtectedImportExportStatus_LicenseExpired           ProtectedImportExportStatus = 6
	ProtectedImportExportStatus_AccessSuspended          ProtectedImportExportStatus = 7
)

// enum
type ProtectionPolicyAuditAction int32

const (
	ProtectionPolicyAuditAction_Decrypt         ProtectionPolicyAuditAction = 0
	ProtectionPolicyAuditAction_CopyToLocation  ProtectionPolicyAuditAction = 1
	ProtectionPolicyAuditAction_SendToRecipient ProtectionPolicyAuditAction = 2
	ProtectionPolicyAuditAction_Other           ProtectionPolicyAuditAction = 3
)

// enum
type ProtectionPolicyEvaluationResult int32

const (
	ProtectionPolicyEvaluationResult_Allowed         ProtectionPolicyEvaluationResult = 0
	ProtectionPolicyEvaluationResult_Blocked         ProtectionPolicyEvaluationResult = 1
	ProtectionPolicyEvaluationResult_ConsentRequired ProtectionPolicyEvaluationResult = 2
)

// enum
type ProtectionPolicyRequestAccessBehavior int32

const (
	ProtectionPolicyRequestAccessBehavior_Decrypt                    ProtectionPolicyRequestAccessBehavior = 0
	ProtectionPolicyRequestAccessBehavior_TreatOverridePolicyAsBlock ProtectionPolicyRequestAccessBehavior = 1
)

// structs

type EnterpriseDataContract struct {
}

// interfaces

// 47995EDC-6CEC-4E3A-B251-9E7485D79E7A
var IID_IBufferProtectUnprotectResult = syscall.GUID{0x47995EDC, 0x6CEC, 0x4E3A,
	[8]byte{0xB2, 0x51, 0x9E, 0x74, 0x85, 0xD7, 0x9E, 0x7A}}

type IBufferProtectUnprotectResultInterface interface {
	win32.IInspectableInterface
	Get_Buffer() *IBuffer
	Get_ProtectionInfo() *IDataProtectionInfo
}

type IBufferProtectUnprotectResultVtbl struct {
	win32.IInspectableVtbl
	Get_Buffer         uintptr
	Get_ProtectionInfo uintptr
}

type IBufferProtectUnprotectResult struct {
	win32.IInspectable
}

func (this *IBufferProtectUnprotectResult) Vtbl() *IBufferProtectUnprotectResultVtbl {
	return (*IBufferProtectUnprotectResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IBufferProtectUnprotectResult) Get_Buffer() *IBuffer {
	var _result *IBuffer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Buffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IBufferProtectUnprotectResult) Get_ProtectionInfo() *IDataProtectionInfo {
	var _result *IDataProtectionInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtectionInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 8420B0C1-5E31-4405-9540-3F943AF0CB26
var IID_IDataProtectionInfo = syscall.GUID{0x8420B0C1, 0x5E31, 0x4405,
	[8]byte{0x95, 0x40, 0x3F, 0x94, 0x3A, 0xF0, 0xCB, 0x26}}

type IDataProtectionInfoInterface interface {
	win32.IInspectableInterface
	Get_Status() DataProtectionStatus
	Get_Identity() string
}

type IDataProtectionInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Status   uintptr
	Get_Identity uintptr
}

type IDataProtectionInfo struct {
	win32.IInspectable
}

func (this *IDataProtectionInfo) Vtbl() *IDataProtectionInfoVtbl {
	return (*IDataProtectionInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataProtectionInfo) Get_Status() DataProtectionStatus {
	var _result DataProtectionStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDataProtectionInfo) Get_Identity() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Identity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// B6149B74-9144-4EE4-8A8A-30B5F361430E
var IID_IDataProtectionManagerStatics = syscall.GUID{0xB6149B74, 0x9144, 0x4EE4,
	[8]byte{0x8A, 0x8A, 0x30, 0xB5, 0xF3, 0x61, 0x43, 0x0E}}

type IDataProtectionManagerStaticsInterface interface {
	win32.IInspectableInterface
	ProtectAsync(data *IBuffer, identity string) *IAsyncOperation[*IBufferProtectUnprotectResult]
	UnprotectAsync(data *IBuffer) *IAsyncOperation[*IBufferProtectUnprotectResult]
	ProtectStreamAsync(unprotectedStream *IInputStream, identity string, protectedStream *IOutputStream) *IAsyncOperation[*IDataProtectionInfo]
	UnprotectStreamAsync(protectedStream *IInputStream, unprotectedStream *IOutputStream) *IAsyncOperation[*IDataProtectionInfo]
	GetProtectionInfoAsync(protectedData *IBuffer) *IAsyncOperation[*IDataProtectionInfo]
	GetStreamProtectionInfoAsync(protectedStream *IInputStream) *IAsyncOperation[*IDataProtectionInfo]
}

type IDataProtectionManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	ProtectAsync                 uintptr
	UnprotectAsync               uintptr
	ProtectStreamAsync           uintptr
	UnprotectStreamAsync         uintptr
	GetProtectionInfoAsync       uintptr
	GetStreamProtectionInfoAsync uintptr
}

type IDataProtectionManagerStatics struct {
	win32.IInspectable
}

func (this *IDataProtectionManagerStatics) Vtbl() *IDataProtectionManagerStaticsVtbl {
	return (*IDataProtectionManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDataProtectionManagerStatics) ProtectAsync(data *IBuffer, identity string) *IAsyncOperation[*IBufferProtectUnprotectResult] {
	var _result *IAsyncOperation[*IBufferProtectUnprotectResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProtectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), NewHStr(identity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataProtectionManagerStatics) UnprotectAsync(data *IBuffer) *IAsyncOperation[*IBufferProtectUnprotectResult] {
	var _result *IAsyncOperation[*IBufferProtectUnprotectResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprotectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataProtectionManagerStatics) ProtectStreamAsync(unprotectedStream *IInputStream, identity string, protectedStream *IOutputStream) *IAsyncOperation[*IDataProtectionInfo] {
	var _result *IAsyncOperation[*IDataProtectionInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProtectStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(unprotectedStream)), NewHStr(identity).Ptr, uintptr(unsafe.Pointer(protectedStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataProtectionManagerStatics) UnprotectStreamAsync(protectedStream *IInputStream, unprotectedStream *IOutputStream) *IAsyncOperation[*IDataProtectionInfo] {
	var _result *IAsyncOperation[*IDataProtectionInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprotectStreamAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(protectedStream)), uintptr(unsafe.Pointer(unprotectedStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataProtectionManagerStatics) GetProtectionInfoAsync(protectedData *IBuffer) *IAsyncOperation[*IDataProtectionInfo] {
	var _result *IAsyncOperation[*IDataProtectionInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetProtectionInfoAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(protectedData)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IDataProtectionManagerStatics) GetStreamProtectionInfoAsync(protectedStream *IInputStream) *IAsyncOperation[*IDataProtectionInfo] {
	var _result *IAsyncOperation[*IDataProtectionInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStreamProtectionInfoAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(protectedStream)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4EE96486-147E-4DD0-8FAF-5253ED91AD0C
var IID_IFileProtectionInfo = syscall.GUID{0x4EE96486, 0x147E, 0x4DD0,
	[8]byte{0x8F, 0xAF, 0x52, 0x53, 0xED, 0x91, 0xAD, 0x0C}}

type IFileProtectionInfoInterface interface {
	win32.IInspectableInterface
	Get_Status() FileProtectionStatus
	Get_IsRoamable() bool
	Get_Identity() string
}

type IFileProtectionInfoVtbl struct {
	win32.IInspectableVtbl
	Get_Status     uintptr
	Get_IsRoamable uintptr
	Get_Identity   uintptr
}

type IFileProtectionInfo struct {
	win32.IInspectable
}

func (this *IFileProtectionInfo) Vtbl() *IFileProtectionInfoVtbl {
	return (*IFileProtectionInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileProtectionInfo) Get_Status() FileProtectionStatus {
	var _result FileProtectionStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFileProtectionInfo) Get_IsRoamable() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRoamable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFileProtectionInfo) Get_Identity() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Identity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 82123A4C-557A-498D-8E94-944CD5836432
var IID_IFileProtectionInfo2 = syscall.GUID{0x82123A4C, 0x557A, 0x498D,
	[8]byte{0x8E, 0x94, 0x94, 0x4C, 0xD5, 0x83, 0x64, 0x32}}

type IFileProtectionInfo2Interface interface {
	win32.IInspectableInterface
	Get_IsProtectWhileOpenSupported() bool
}

type IFileProtectionInfo2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsProtectWhileOpenSupported uintptr
}

type IFileProtectionInfo2 struct {
	win32.IInspectable
}

func (this *IFileProtectionInfo2) Vtbl() *IFileProtectionInfo2Vtbl {
	return (*IFileProtectionInfo2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileProtectionInfo2) Get_IsProtectWhileOpenSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsProtectWhileOpenSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5846FC9B-E613-426B-BB38-88CBA1DC9ADB
var IID_IFileProtectionManagerStatics = syscall.GUID{0x5846FC9B, 0xE613, 0x426B,
	[8]byte{0xBB, 0x38, 0x88, 0xCB, 0xA1, 0xDC, 0x9A, 0xDB}}

type IFileProtectionManagerStaticsInterface interface {
	win32.IInspectableInterface
	ProtectAsync(target *IStorageItem, identity string) *IAsyncOperation[*IFileProtectionInfo]
	CopyProtectionAsync(source *IStorageItem, target *IStorageItem) *IAsyncOperation[bool]
	GetProtectionInfoAsync(source *IStorageItem) *IAsyncOperation[*IFileProtectionInfo]
	SaveFileAsContainerAsync(protectedFile *IStorageFile) *IAsyncOperation[*IProtectedContainerExportResult]
	LoadFileFromContainerAsync(containerFile *IStorageFile) *IAsyncOperation[*IProtectedContainerImportResult]
	LoadFileFromContainerWithTargetAsync(containerFile *IStorageFile, target *IStorageItem) *IAsyncOperation[*IProtectedContainerImportResult]
	CreateProtectedAndOpenAsync(parentFolder *IStorageFolder, desiredName string, identity string, collisionOption CreationCollisionOption) *IAsyncOperation[*IProtectedFileCreateResult]
}

type IFileProtectionManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	ProtectAsync                         uintptr
	CopyProtectionAsync                  uintptr
	GetProtectionInfoAsync               uintptr
	SaveFileAsContainerAsync             uintptr
	LoadFileFromContainerAsync           uintptr
	LoadFileFromContainerWithTargetAsync uintptr
	CreateProtectedAndOpenAsync          uintptr
}

type IFileProtectionManagerStatics struct {
	win32.IInspectable
}

func (this *IFileProtectionManagerStatics) Vtbl() *IFileProtectionManagerStaticsVtbl {
	return (*IFileProtectionManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileProtectionManagerStatics) ProtectAsync(target *IStorageItem, identity string) *IAsyncOperation[*IFileProtectionInfo] {
	var _result *IAsyncOperation[*IFileProtectionInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProtectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(target)), NewHStr(identity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileProtectionManagerStatics) CopyProtectionAsync(source *IStorageItem, target *IStorageItem) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyProtectionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(target)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileProtectionManagerStatics) GetProtectionInfoAsync(source *IStorageItem) *IAsyncOperation[*IFileProtectionInfo] {
	var _result *IAsyncOperation[*IFileProtectionInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetProtectionInfoAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileProtectionManagerStatics) SaveFileAsContainerAsync(protectedFile *IStorageFile) *IAsyncOperation[*IProtectedContainerExportResult] {
	var _result *IAsyncOperation[*IProtectedContainerExportResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveFileAsContainerAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(protectedFile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileProtectionManagerStatics) LoadFileFromContainerAsync(containerFile *IStorageFile) *IAsyncOperation[*IProtectedContainerImportResult] {
	var _result *IAsyncOperation[*IProtectedContainerImportResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadFileFromContainerAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(containerFile)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileProtectionManagerStatics) LoadFileFromContainerWithTargetAsync(containerFile *IStorageFile, target *IStorageItem) *IAsyncOperation[*IProtectedContainerImportResult] {
	var _result *IAsyncOperation[*IProtectedContainerImportResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadFileFromContainerWithTargetAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(containerFile)), uintptr(unsafe.Pointer(target)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileProtectionManagerStatics) CreateProtectedAndOpenAsync(parentFolder *IStorageFolder, desiredName string, identity string, collisionOption CreationCollisionOption) *IAsyncOperation[*IProtectedFileCreateResult] {
	var _result *IAsyncOperation[*IProtectedFileCreateResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateProtectedAndOpenAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(parentFolder)), NewHStr(desiredName).Ptr, NewHStr(identity).Ptr, uintptr(collisionOption), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 83D2A745-0483-41AB-B2D5-BC7F23D74EBB
var IID_IFileProtectionManagerStatics2 = syscall.GUID{0x83D2A745, 0x0483, 0x41AB,
	[8]byte{0xB2, 0xD5, 0xBC, 0x7F, 0x23, 0xD7, 0x4E, 0xBB}}

type IFileProtectionManagerStatics2Interface interface {
	win32.IInspectableInterface
	IsContainerAsync(file *IStorageFile) *IAsyncOperation[bool]
	LoadFileFromContainerWithTargetAndNameCollisionOptionAsync(containerFile *IStorageFile, target *IStorageItem, collisionOption NameCollisionOption) *IAsyncOperation[*IProtectedContainerImportResult]
	SaveFileAsContainerWithSharingAsync(protectedFile *IStorageFile, sharedWithIdentities *IIterable[string]) *IAsyncOperation[*IProtectedContainerExportResult]
}

type IFileProtectionManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	IsContainerAsync                                           uintptr
	LoadFileFromContainerWithTargetAndNameCollisionOptionAsync uintptr
	SaveFileAsContainerWithSharingAsync                        uintptr
}

type IFileProtectionManagerStatics2 struct {
	win32.IInspectable
}

func (this *IFileProtectionManagerStatics2) Vtbl() *IFileProtectionManagerStatics2Vtbl {
	return (*IFileProtectionManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileProtectionManagerStatics2) IsContainerAsync(file *IStorageFile) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsContainerAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileProtectionManagerStatics2) LoadFileFromContainerWithTargetAndNameCollisionOptionAsync(containerFile *IStorageFile, target *IStorageItem, collisionOption NameCollisionOption) *IAsyncOperation[*IProtectedContainerImportResult] {
	var _result *IAsyncOperation[*IProtectedContainerImportResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LoadFileFromContainerWithTargetAndNameCollisionOptionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(containerFile)), uintptr(unsafe.Pointer(target)), uintptr(collisionOption), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileProtectionManagerStatics2) SaveFileAsContainerWithSharingAsync(protectedFile *IStorageFile, sharedWithIdentities *IIterable[string]) *IAsyncOperation[*IProtectedContainerExportResult] {
	var _result *IAsyncOperation[*IProtectedContainerExportResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SaveFileAsContainerWithSharingAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(protectedFile)), uintptr(unsafe.Pointer(sharedWithIdentities)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6918849A-624F-46D6-B241-E9CD5FDF3E3F
var IID_IFileProtectionManagerStatics3 = syscall.GUID{0x6918849A, 0x624F, 0x46D6,
	[8]byte{0xB2, 0x41, 0xE9, 0xCD, 0x5F, 0xDF, 0x3E, 0x3F}}

type IFileProtectionManagerStatics3Interface interface {
	win32.IInspectableInterface
	UnprotectAsync(target *IStorageItem) *IAsyncOperation[*IFileProtectionInfo]
	UnprotectWithOptionsAsync(target *IStorageItem, options *IFileUnprotectOptions) *IAsyncOperation[*IFileProtectionInfo]
}

type IFileProtectionManagerStatics3Vtbl struct {
	win32.IInspectableVtbl
	UnprotectAsync            uintptr
	UnprotectWithOptionsAsync uintptr
}

type IFileProtectionManagerStatics3 struct {
	win32.IInspectable
}

func (this *IFileProtectionManagerStatics3) Vtbl() *IFileProtectionManagerStatics3Vtbl {
	return (*IFileProtectionManagerStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileProtectionManagerStatics3) UnprotectAsync(target *IStorageItem) *IAsyncOperation[*IFileProtectionInfo] {
	var _result *IAsyncOperation[*IFileProtectionInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprotectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(target)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileProtectionManagerStatics3) UnprotectWithOptionsAsync(target *IStorageItem, options *IFileUnprotectOptions) *IAsyncOperation[*IFileProtectionInfo] {
	var _result *IAsyncOperation[*IFileProtectionInfo]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UnprotectWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(target)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 256BBC3D-1C5D-4260-8C75-9144CFB78BA9
var IID_IFileRevocationManagerStatics = syscall.GUID{0x256BBC3D, 0x1C5D, 0x4260,
	[8]byte{0x8C, 0x75, 0x91, 0x44, 0xCF, 0xB7, 0x8B, 0xA9}}

type IFileRevocationManagerStaticsInterface interface {
	win32.IInspectableInterface
	ProtectAsync(storageItem *IStorageItem, enterpriseIdentity string) *IAsyncOperation[FileProtectionStatus]
	CopyProtectionAsync(sourceStorageItem *IStorageItem, targetStorageItem *IStorageItem) *IAsyncOperation[bool]
	Revoke(enterpriseIdentity string)
	GetStatusAsync(storageItem *IStorageItem) *IAsyncOperation[FileProtectionStatus]
}

type IFileRevocationManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	ProtectAsync        uintptr
	CopyProtectionAsync uintptr
	Revoke              uintptr
	GetStatusAsync      uintptr
}

type IFileRevocationManagerStatics struct {
	win32.IInspectable
}

func (this *IFileRevocationManagerStatics) Vtbl() *IFileRevocationManagerStaticsVtbl {
	return (*IFileRevocationManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileRevocationManagerStatics) ProtectAsync(storageItem *IStorageItem, enterpriseIdentity string) *IAsyncOperation[FileProtectionStatus] {
	var _result *IAsyncOperation[FileProtectionStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProtectAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(storageItem)), NewHStr(enterpriseIdentity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileRevocationManagerStatics) CopyProtectionAsync(sourceStorageItem *IStorageItem, targetStorageItem *IStorageItem) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CopyProtectionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceStorageItem)), uintptr(unsafe.Pointer(targetStorageItem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFileRevocationManagerStatics) Revoke(enterpriseIdentity string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Revoke, uintptr(unsafe.Pointer(this)), NewHStr(enterpriseIdentity).Ptr)
	_ = _hr
}

func (this *IFileRevocationManagerStatics) GetStatusAsync(storageItem *IStorageItem) *IAsyncOperation[FileProtectionStatus] {
	var _result *IAsyncOperation[FileProtectionStatus]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetStatusAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(storageItem)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7D1312F1-3B0D-4DD8-A1F8-1EC53822E2F3
var IID_IFileUnprotectOptions = syscall.GUID{0x7D1312F1, 0x3B0D, 0x4DD8,
	[8]byte{0xA1, 0xF8, 0x1E, 0xC5, 0x38, 0x22, 0xE2, 0xF3}}

type IFileUnprotectOptionsInterface interface {
	win32.IInspectableInterface
	Put_Audit(value bool)
	Get_Audit() bool
}

type IFileUnprotectOptionsVtbl struct {
	win32.IInspectableVtbl
	Put_Audit uintptr
	Get_Audit uintptr
}

type IFileUnprotectOptions struct {
	win32.IInspectable
}

func (this *IFileUnprotectOptions) Vtbl() *IFileUnprotectOptionsVtbl {
	return (*IFileUnprotectOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileUnprotectOptions) Put_Audit(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Audit, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IFileUnprotectOptions) Get_Audit() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Audit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 51AEB39C-DA8C-4C3F-9BFB-CB73A7CCE0DD
var IID_IFileUnprotectOptionsFactory = syscall.GUID{0x51AEB39C, 0xDA8C, 0x4C3F,
	[8]byte{0x9B, 0xFB, 0xCB, 0x73, 0xA7, 0xCC, 0xE0, 0xDD}}

type IFileUnprotectOptionsFactoryInterface interface {
	win32.IInspectableInterface
	Create(audit bool) *IFileUnprotectOptions
}

type IFileUnprotectOptionsFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type IFileUnprotectOptionsFactory struct {
	win32.IInspectable
}

func (this *IFileUnprotectOptionsFactory) Vtbl() *IFileUnprotectOptionsFactoryVtbl {
	return (*IFileUnprotectOptionsFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFileUnprotectOptionsFactory) Create(audit bool) *IFileUnprotectOptions {
	var _result *IFileUnprotectOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&audit))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// AC4DCA59-5D80-4E95-8C5F-8539450EEBE0
var IID_IProtectedAccessResumedEventArgs = syscall.GUID{0xAC4DCA59, 0x5D80, 0x4E95,
	[8]byte{0x8C, 0x5F, 0x85, 0x39, 0x45, 0x0E, 0xEB, 0xE0}}

type IProtectedAccessResumedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Identities() *IVectorView[string]
}

type IProtectedAccessResumedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Identities uintptr
}

type IProtectedAccessResumedEventArgs struct {
	win32.IInspectable
}

func (this *IProtectedAccessResumedEventArgs) Vtbl() *IProtectedAccessResumedEventArgsVtbl {
	return (*IProtectedAccessResumedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectedAccessResumedEventArgs) Get_Identities() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Identities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 75A193E0-A344-429F-B975-04FC1F88C185
var IID_IProtectedAccessSuspendingEventArgs = syscall.GUID{0x75A193E0, 0xA344, 0x429F,
	[8]byte{0xB9, 0x75, 0x04, 0xFC, 0x1F, 0x88, 0xC1, 0x85}}

type IProtectedAccessSuspendingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Identities() *IVectorView[string]
	Get_Deadline() DateTime
	GetDeferral() *IDeferral
}

type IProtectedAccessSuspendingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Identities uintptr
	Get_Deadline   uintptr
	GetDeferral    uintptr
}

type IProtectedAccessSuspendingEventArgs struct {
	win32.IInspectable
}

func (this *IProtectedAccessSuspendingEventArgs) Vtbl() *IProtectedAccessSuspendingEventArgsVtbl {
	return (*IProtectedAccessSuspendingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectedAccessSuspendingEventArgs) Get_Identities() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Identities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectedAccessSuspendingEventArgs) Get_Deadline() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Deadline, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectedAccessSuspendingEventArgs) GetDeferral() *IDeferral {
	var _result *IDeferral
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDeferral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3948EF95-F7FB-4B42-AFB0-DF70B41543C1
var IID_IProtectedContainerExportResult = syscall.GUID{0x3948EF95, 0xF7FB, 0x4B42,
	[8]byte{0xAF, 0xB0, 0xDF, 0x70, 0xB4, 0x15, 0x43, 0xC1}}

type IProtectedContainerExportResultInterface interface {
	win32.IInspectableInterface
	Get_Status() ProtectedImportExportStatus
	Get_File() *IStorageFile
}

type IProtectedContainerExportResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_File   uintptr
}

type IProtectedContainerExportResult struct {
	win32.IInspectable
}

func (this *IProtectedContainerExportResult) Vtbl() *IProtectedContainerExportResultVtbl {
	return (*IProtectedContainerExportResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectedContainerExportResult) Get_Status() ProtectedImportExportStatus {
	var _result ProtectedImportExportStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectedContainerExportResult) Get_File() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_File, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// CDB780D1-E7BB-4D1A-9339-34DC41149F9B
var IID_IProtectedContainerImportResult = syscall.GUID{0xCDB780D1, 0xE7BB, 0x4D1A,
	[8]byte{0x93, 0x39, 0x34, 0xDC, 0x41, 0x14, 0x9F, 0x9B}}

type IProtectedContainerImportResultInterface interface {
	win32.IInspectableInterface
	Get_Status() ProtectedImportExportStatus
	Get_File() *IStorageFile
}

type IProtectedContainerImportResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
	Get_File   uintptr
}

type IProtectedContainerImportResult struct {
	win32.IInspectable
}

func (this *IProtectedContainerImportResult) Vtbl() *IProtectedContainerImportResultVtbl {
	return (*IProtectedContainerImportResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectedContainerImportResult) Get_Status() ProtectedImportExportStatus {
	var _result ProtectedImportExportStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectedContainerImportResult) Get_File() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_File, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 63686821-58B9-47EE-93D9-F0F741CF43F0
var IID_IProtectedContentRevokedEventArgs = syscall.GUID{0x63686821, 0x58B9, 0x47EE,
	[8]byte{0x93, 0xD9, 0xF0, 0xF7, 0x41, 0xCF, 0x43, 0xF0}}

type IProtectedContentRevokedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Identities() *IVectorView[string]
}

type IProtectedContentRevokedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Identities uintptr
}

type IProtectedContentRevokedEventArgs struct {
	win32.IInspectable
}

func (this *IProtectedContentRevokedEventArgs) Vtbl() *IProtectedContentRevokedEventArgsVtbl {
	return (*IProtectedContentRevokedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectedContentRevokedEventArgs) Get_Identities() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Identities, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 28E3ED6A-E9E7-4A03-9F53-BDB16172699B
var IID_IProtectedFileCreateResult = syscall.GUID{0x28E3ED6A, 0xE9E7, 0x4A03,
	[8]byte{0x9F, 0x53, 0xBD, 0xB1, 0x61, 0x72, 0x69, 0x9B}}

type IProtectedFileCreateResultInterface interface {
	win32.IInspectableInterface
	Get_File() *IStorageFile
	Get_Stream() *IRandomAccessStream
	Get_ProtectionInfo() *IFileProtectionInfo
}

type IProtectedFileCreateResultVtbl struct {
	win32.IInspectableVtbl
	Get_File           uintptr
	Get_Stream         uintptr
	Get_ProtectionInfo uintptr
}

type IProtectedFileCreateResult struct {
	win32.IInspectable
}

func (this *IProtectedFileCreateResult) Vtbl() *IProtectedFileCreateResultVtbl {
	return (*IProtectedFileCreateResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectedFileCreateResult) Get_File() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_File, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectedFileCreateResult) Get_Stream() *IRandomAccessStream {
	var _result *IRandomAccessStream
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Stream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectedFileCreateResult) Get_ProtectionInfo() *IFileProtectionInfo {
	var _result *IFileProtectionInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProtectionInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 425AB7E4-FEB7-44FC-B3BB-C3C4D7ECBEBB
var IID_IProtectionPolicyAuditInfo = syscall.GUID{0x425AB7E4, 0xFEB7, 0x44FC,
	[8]byte{0xB3, 0xBB, 0xC3, 0xC4, 0xD7, 0xEC, 0xBE, 0xBB}}

type IProtectionPolicyAuditInfoInterface interface {
	win32.IInspectableInterface
	Put_Action(value ProtectionPolicyAuditAction)
	Get_Action() ProtectionPolicyAuditAction
	Put_DataDescription(value string)
	Get_DataDescription() string
	Put_SourceDescription(value string)
	Get_SourceDescription() string
	Put_TargetDescription(value string)
	Get_TargetDescription() string
}

type IProtectionPolicyAuditInfoVtbl struct {
	win32.IInspectableVtbl
	Put_Action            uintptr
	Get_Action            uintptr
	Put_DataDescription   uintptr
	Get_DataDescription   uintptr
	Put_SourceDescription uintptr
	Get_SourceDescription uintptr
	Put_TargetDescription uintptr
	Get_TargetDescription uintptr
}

type IProtectionPolicyAuditInfo struct {
	win32.IInspectable
}

func (this *IProtectionPolicyAuditInfo) Vtbl() *IProtectionPolicyAuditInfoVtbl {
	return (*IProtectionPolicyAuditInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectionPolicyAuditInfo) Put_Action(value ProtectionPolicyAuditAction) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Action, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IProtectionPolicyAuditInfo) Get_Action() ProtectionPolicyAuditAction {
	var _result ProtectionPolicyAuditAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Action, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyAuditInfo) Put_DataDescription(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DataDescription, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IProtectionPolicyAuditInfo) Get_DataDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DataDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IProtectionPolicyAuditInfo) Put_SourceDescription(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_SourceDescription, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IProtectionPolicyAuditInfo) Get_SourceDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SourceDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IProtectionPolicyAuditInfo) Put_TargetDescription(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TargetDescription, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IProtectionPolicyAuditInfo) Get_TargetDescription() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TargetDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 7ED4180B-92E8-42D5-83D4-25440B423549
var IID_IProtectionPolicyAuditInfoFactory = syscall.GUID{0x7ED4180B, 0x92E8, 0x42D5,
	[8]byte{0x83, 0xD4, 0x25, 0x44, 0x0B, 0x42, 0x35, 0x49}}

type IProtectionPolicyAuditInfoFactoryInterface interface {
	win32.IInspectableInterface
	Create(action ProtectionPolicyAuditAction, dataDescription string, sourceDescription string, targetDescription string) *IProtectionPolicyAuditInfo
	CreateWithActionAndDataDescription(action ProtectionPolicyAuditAction, dataDescription string) *IProtectionPolicyAuditInfo
}

type IProtectionPolicyAuditInfoFactoryVtbl struct {
	win32.IInspectableVtbl
	Create                             uintptr
	CreateWithActionAndDataDescription uintptr
}

type IProtectionPolicyAuditInfoFactory struct {
	win32.IInspectable
}

func (this *IProtectionPolicyAuditInfoFactory) Vtbl() *IProtectionPolicyAuditInfoFactoryVtbl {
	return (*IProtectionPolicyAuditInfoFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectionPolicyAuditInfoFactory) Create(action ProtectionPolicyAuditAction, dataDescription string, sourceDescription string, targetDescription string) *IProtectionPolicyAuditInfo {
	var _result *IProtectionPolicyAuditInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(action), NewHStr(dataDescription).Ptr, NewHStr(sourceDescription).Ptr, NewHStr(targetDescription).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyAuditInfoFactory) CreateWithActionAndDataDescription(action ProtectionPolicyAuditAction, dataDescription string) *IProtectionPolicyAuditInfo {
	var _result *IProtectionPolicyAuditInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithActionAndDataDescription, uintptr(unsafe.Pointer(this)), uintptr(action), NewHStr(dataDescription).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// D5703E18-A08D-47E6-A240-9934D7165EB5
var IID_IProtectionPolicyManager = syscall.GUID{0xD5703E18, 0xA08D, 0x47E6,
	[8]byte{0xA2, 0x40, 0x99, 0x34, 0xD7, 0x16, 0x5E, 0xB5}}

type IProtectionPolicyManagerInterface interface {
	win32.IInspectableInterface
	Put_Identity(value string)
	Get_Identity() string
}

type IProtectionPolicyManagerVtbl struct {
	win32.IInspectableVtbl
	Put_Identity uintptr
	Get_Identity uintptr
}

type IProtectionPolicyManager struct {
	win32.IInspectable
}

func (this *IProtectionPolicyManager) Vtbl() *IProtectionPolicyManagerVtbl {
	return (*IProtectionPolicyManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectionPolicyManager) Put_Identity(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Identity, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *IProtectionPolicyManager) Get_Identity() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Identity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// ABF7527A-8435-417F-99B6-51BEAF365888
var IID_IProtectionPolicyManager2 = syscall.GUID{0xABF7527A, 0x8435, 0x417F,
	[8]byte{0x99, 0xB6, 0x51, 0xBE, 0xAF, 0x36, 0x58, 0x88}}

type IProtectionPolicyManager2Interface interface {
	win32.IInspectableInterface
	Put_ShowEnterpriseIndicator(value bool)
	Get_ShowEnterpriseIndicator() bool
}

type IProtectionPolicyManager2Vtbl struct {
	win32.IInspectableVtbl
	Put_ShowEnterpriseIndicator uintptr
	Get_ShowEnterpriseIndicator uintptr
}

type IProtectionPolicyManager2 struct {
	win32.IInspectable
}

func (this *IProtectionPolicyManager2) Vtbl() *IProtectionPolicyManager2Vtbl {
	return (*IProtectionPolicyManager2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectionPolicyManager2) Put_ShowEnterpriseIndicator(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ShowEnterpriseIndicator, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IProtectionPolicyManager2) Get_ShowEnterpriseIndicator() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShowEnterpriseIndicator, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// C0BFFC66-8C3D-4D56-8804-C68F0AD32EC5
var IID_IProtectionPolicyManagerStatics = syscall.GUID{0xC0BFFC66, 0x8C3D, 0x4D56,
	[8]byte{0x88, 0x04, 0xC6, 0x8F, 0x0A, 0xD3, 0x2E, 0xC5}}

type IProtectionPolicyManagerStaticsInterface interface {
	win32.IInspectableInterface
	IsIdentityManaged(identity string) bool
	TryApplyProcessUIPolicy(identity string) bool
	ClearProcessUIPolicy()
	CreateCurrentThreadNetworkContext(identity string) *IThreadNetworkContext
	GetPrimaryManagedIdentityForNetworkEndpointAsync(endpointHost *IHostName) *IAsyncOperation[string]
	RevokeContent(identity string)
	GetForCurrentView() *IProtectionPolicyManager
	Add_ProtectedAccessSuspending(handler EventHandler[*IProtectedAccessSuspendingEventArgs]) EventRegistrationToken
	Remove_ProtectedAccessSuspending(token EventRegistrationToken)
	Add_ProtectedAccessResumed(handler EventHandler[*IProtectedAccessResumedEventArgs]) EventRegistrationToken
	Remove_ProtectedAccessResumed(token EventRegistrationToken)
	Add_ProtectedContentRevoked(handler EventHandler[*IProtectedContentRevokedEventArgs]) EventRegistrationToken
	Remove_ProtectedContentRevoked(token EventRegistrationToken)
	CheckAccess(sourceIdentity string, targetIdentity string) ProtectionPolicyEvaluationResult
	RequestAccessAsync(sourceIdentity string, targetIdentity string) *IAsyncOperation[ProtectionPolicyEvaluationResult]
}

type IProtectionPolicyManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	IsIdentityManaged                                uintptr
	TryApplyProcessUIPolicy                          uintptr
	ClearProcessUIPolicy                             uintptr
	CreateCurrentThreadNetworkContext                uintptr
	GetPrimaryManagedIdentityForNetworkEndpointAsync uintptr
	RevokeContent                                    uintptr
	GetForCurrentView                                uintptr
	Add_ProtectedAccessSuspending                    uintptr
	Remove_ProtectedAccessSuspending                 uintptr
	Add_ProtectedAccessResumed                       uintptr
	Remove_ProtectedAccessResumed                    uintptr
	Add_ProtectedContentRevoked                      uintptr
	Remove_ProtectedContentRevoked                   uintptr
	CheckAccess                                      uintptr
	RequestAccessAsync                               uintptr
}

type IProtectionPolicyManagerStatics struct {
	win32.IInspectable
}

func (this *IProtectionPolicyManagerStatics) Vtbl() *IProtectionPolicyManagerStaticsVtbl {
	return (*IProtectionPolicyManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectionPolicyManagerStatics) IsIdentityManaged(identity string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsIdentityManaged, uintptr(unsafe.Pointer(this)), NewHStr(identity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyManagerStatics) TryApplyProcessUIPolicy(identity string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryApplyProcessUIPolicy, uintptr(unsafe.Pointer(this)), NewHStr(identity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyManagerStatics) ClearProcessUIPolicy() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearProcessUIPolicy, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *IProtectionPolicyManagerStatics) CreateCurrentThreadNetworkContext(identity string) *IThreadNetworkContext {
	var _result *IThreadNetworkContext
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateCurrentThreadNetworkContext, uintptr(unsafe.Pointer(this)), NewHStr(identity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics) GetPrimaryManagedIdentityForNetworkEndpointAsync(endpointHost *IHostName) *IAsyncOperation[string] {
	var _result *IAsyncOperation[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPrimaryManagedIdentityForNetworkEndpointAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(endpointHost)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics) RevokeContent(identity string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RevokeContent, uintptr(unsafe.Pointer(this)), NewHStr(identity).Ptr)
	_ = _hr
}

func (this *IProtectionPolicyManagerStatics) GetForCurrentView() *IProtectionPolicyManager {
	var _result *IProtectionPolicyManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics) Add_ProtectedAccessSuspending(handler EventHandler[*IProtectedAccessSuspendingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ProtectedAccessSuspending, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyManagerStatics) Remove_ProtectedAccessSuspending(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ProtectedAccessSuspending, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IProtectionPolicyManagerStatics) Add_ProtectedAccessResumed(handler EventHandler[*IProtectedAccessResumedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ProtectedAccessResumed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyManagerStatics) Remove_ProtectedAccessResumed(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ProtectedAccessResumed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IProtectionPolicyManagerStatics) Add_ProtectedContentRevoked(handler EventHandler[*IProtectedContentRevokedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ProtectedContentRevoked, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyManagerStatics) Remove_ProtectedContentRevoked(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ProtectedContentRevoked, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IProtectionPolicyManagerStatics) CheckAccess(sourceIdentity string, targetIdentity string) ProtectionPolicyEvaluationResult {
	var _result ProtectionPolicyEvaluationResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckAccess, uintptr(unsafe.Pointer(this)), NewHStr(sourceIdentity).Ptr, NewHStr(targetIdentity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyManagerStatics) RequestAccessAsync(sourceIdentity string, targetIdentity string) *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessAsync, uintptr(unsafe.Pointer(this)), NewHStr(sourceIdentity).Ptr, NewHStr(targetIdentity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B68F9A8C-39E0-4649-B2E4-070AB8A579B3
var IID_IProtectionPolicyManagerStatics2 = syscall.GUID{0xB68F9A8C, 0x39E0, 0x4649,
	[8]byte{0xB2, 0xE4, 0x07, 0x0A, 0xB8, 0xA5, 0x79, 0xB3}}

type IProtectionPolicyManagerStatics2Interface interface {
	win32.IInspectableInterface
	HasContentBeenRevokedSince(identity string, since DateTime) bool
	CheckAccessForApp(sourceIdentity string, appPackageFamilyName string) ProtectionPolicyEvaluationResult
	RequestAccessForAppAsync(sourceIdentity string, appPackageFamilyName string) *IAsyncOperation[ProtectionPolicyEvaluationResult]
	GetEnforcementLevel(identity string) EnforcementLevel
	IsUserDecryptionAllowed(identity string) bool
	IsProtectionUnderLockRequired(identity string) bool
	Add_PolicyChanged(handler EventHandler[interface{}]) EventRegistrationToken
	Remove_PolicyChanged(token EventRegistrationToken)
	Get_IsProtectionEnabled() bool
}

type IProtectionPolicyManagerStatics2Vtbl struct {
	win32.IInspectableVtbl
	HasContentBeenRevokedSince    uintptr
	CheckAccessForApp             uintptr
	RequestAccessForAppAsync      uintptr
	GetEnforcementLevel           uintptr
	IsUserDecryptionAllowed       uintptr
	IsProtectionUnderLockRequired uintptr
	Add_PolicyChanged             uintptr
	Remove_PolicyChanged          uintptr
	Get_IsProtectionEnabled       uintptr
}

type IProtectionPolicyManagerStatics2 struct {
	win32.IInspectable
}

func (this *IProtectionPolicyManagerStatics2) Vtbl() *IProtectionPolicyManagerStatics2Vtbl {
	return (*IProtectionPolicyManagerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectionPolicyManagerStatics2) HasContentBeenRevokedSince(identity string, since DateTime) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().HasContentBeenRevokedSince, uintptr(unsafe.Pointer(this)), NewHStr(identity).Ptr, *(*uintptr)(unsafe.Pointer(&since)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyManagerStatics2) CheckAccessForApp(sourceIdentity string, appPackageFamilyName string) ProtectionPolicyEvaluationResult {
	var _result ProtectionPolicyEvaluationResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CheckAccessForApp, uintptr(unsafe.Pointer(this)), NewHStr(sourceIdentity).Ptr, NewHStr(appPackageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyManagerStatics2) RequestAccessForAppAsync(sourceIdentity string, appPackageFamilyName string) *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessForAppAsync, uintptr(unsafe.Pointer(this)), NewHStr(sourceIdentity).Ptr, NewHStr(appPackageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics2) GetEnforcementLevel(identity string) EnforcementLevel {
	var _result EnforcementLevel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetEnforcementLevel, uintptr(unsafe.Pointer(this)), NewHStr(identity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyManagerStatics2) IsUserDecryptionAllowed(identity string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsUserDecryptionAllowed, uintptr(unsafe.Pointer(this)), NewHStr(identity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyManagerStatics2) IsProtectionUnderLockRequired(identity string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsProtectionUnderLockRequired, uintptr(unsafe.Pointer(this)), NewHStr(identity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyManagerStatics2) Add_PolicyChanged(handler EventHandler[interface{}]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_PolicyChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(handler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyManagerStatics2) Remove_PolicyChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_PolicyChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IProtectionPolicyManagerStatics2) Get_IsProtectionEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsProtectionEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 48FF9E8C-6A6F-4D9F-BCED-18AB537AA015
var IID_IProtectionPolicyManagerStatics3 = syscall.GUID{0x48FF9E8C, 0x6A6F, 0x4D9F,
	[8]byte{0xBC, 0xED, 0x18, 0xAB, 0x53, 0x7A, 0xA0, 0x15}}

type IProtectionPolicyManagerStatics3Interface interface {
	win32.IInspectableInterface
	RequestAccessWithAuditingInfoAsync(sourceIdentity string, targetIdentity string, auditInfo *IProtectionPolicyAuditInfo) *IAsyncOperation[ProtectionPolicyEvaluationResult]
	RequestAccessWithMessageAsync(sourceIdentity string, targetIdentity string, auditInfo *IProtectionPolicyAuditInfo, messageFromApp string) *IAsyncOperation[ProtectionPolicyEvaluationResult]
	RequestAccessForAppWithAuditingInfoAsync(sourceIdentity string, appPackageFamilyName string, auditInfo *IProtectionPolicyAuditInfo) *IAsyncOperation[ProtectionPolicyEvaluationResult]
	RequestAccessForAppWithMessageAsync(sourceIdentity string, appPackageFamilyName string, auditInfo *IProtectionPolicyAuditInfo, messageFromApp string) *IAsyncOperation[ProtectionPolicyEvaluationResult]
	LogAuditEvent(sourceIdentity string, targetIdentity string, auditInfo *IProtectionPolicyAuditInfo)
}

type IProtectionPolicyManagerStatics3Vtbl struct {
	win32.IInspectableVtbl
	RequestAccessWithAuditingInfoAsync       uintptr
	RequestAccessWithMessageAsync            uintptr
	RequestAccessForAppWithAuditingInfoAsync uintptr
	RequestAccessForAppWithMessageAsync      uintptr
	LogAuditEvent                            uintptr
}

type IProtectionPolicyManagerStatics3 struct {
	win32.IInspectable
}

func (this *IProtectionPolicyManagerStatics3) Vtbl() *IProtectionPolicyManagerStatics3Vtbl {
	return (*IProtectionPolicyManagerStatics3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectionPolicyManagerStatics3) RequestAccessWithAuditingInfoAsync(sourceIdentity string, targetIdentity string, auditInfo *IProtectionPolicyAuditInfo) *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessWithAuditingInfoAsync, uintptr(unsafe.Pointer(this)), NewHStr(sourceIdentity).Ptr, NewHStr(targetIdentity).Ptr, uintptr(unsafe.Pointer(auditInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics3) RequestAccessWithMessageAsync(sourceIdentity string, targetIdentity string, auditInfo *IProtectionPolicyAuditInfo, messageFromApp string) *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessWithMessageAsync, uintptr(unsafe.Pointer(this)), NewHStr(sourceIdentity).Ptr, NewHStr(targetIdentity).Ptr, uintptr(unsafe.Pointer(auditInfo)), NewHStr(messageFromApp).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics3) RequestAccessForAppWithAuditingInfoAsync(sourceIdentity string, appPackageFamilyName string, auditInfo *IProtectionPolicyAuditInfo) *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessForAppWithAuditingInfoAsync, uintptr(unsafe.Pointer(this)), NewHStr(sourceIdentity).Ptr, NewHStr(appPackageFamilyName).Ptr, uintptr(unsafe.Pointer(auditInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics3) RequestAccessForAppWithMessageAsync(sourceIdentity string, appPackageFamilyName string, auditInfo *IProtectionPolicyAuditInfo, messageFromApp string) *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessForAppWithMessageAsync, uintptr(unsafe.Pointer(this)), NewHStr(sourceIdentity).Ptr, NewHStr(appPackageFamilyName).Ptr, uintptr(unsafe.Pointer(auditInfo)), NewHStr(messageFromApp).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics3) LogAuditEvent(sourceIdentity string, targetIdentity string, auditInfo *IProtectionPolicyAuditInfo) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().LogAuditEvent, uintptr(unsafe.Pointer(this)), NewHStr(sourceIdentity).Ptr, NewHStr(targetIdentity).Ptr, uintptr(unsafe.Pointer(auditInfo)))
	_ = _hr
}

// 20B794DB-CCBD-490F-8C83-49CCB77AEA6C
var IID_IProtectionPolicyManagerStatics4 = syscall.GUID{0x20B794DB, 0xCCBD, 0x490F,
	[8]byte{0x8C, 0x83, 0x49, 0xCC, 0xB7, 0x7A, 0xEA, 0x6C}}

type IProtectionPolicyManagerStatics4Interface interface {
	win32.IInspectableInterface
	IsRoamableProtectionEnabled(identity string) bool
	RequestAccessWithBehaviorAsync(sourceIdentity string, targetIdentity string, auditInfo *IProtectionPolicyAuditInfo, messageFromApp string, behavior ProtectionPolicyRequestAccessBehavior) *IAsyncOperation[ProtectionPolicyEvaluationResult]
	RequestAccessForAppWithBehaviorAsync(sourceIdentity string, appPackageFamilyName string, auditInfo *IProtectionPolicyAuditInfo, messageFromApp string, behavior ProtectionPolicyRequestAccessBehavior) *IAsyncOperation[ProtectionPolicyEvaluationResult]
	RequestAccessToFilesForAppAsync(sourceItemList *IIterable[*IStorageItem], appPackageFamilyName string, auditInfo *IProtectionPolicyAuditInfo) *IAsyncOperation[ProtectionPolicyEvaluationResult]
	RequestAccessToFilesForAppWithMessageAndBehaviorAsync(sourceItemList *IIterable[*IStorageItem], appPackageFamilyName string, auditInfo *IProtectionPolicyAuditInfo, messageFromApp string, behavior ProtectionPolicyRequestAccessBehavior) *IAsyncOperation[ProtectionPolicyEvaluationResult]
	RequestAccessToFilesForProcessAsync(sourceItemList *IIterable[*IStorageItem], processId uint32, auditInfo *IProtectionPolicyAuditInfo) *IAsyncOperation[ProtectionPolicyEvaluationResult]
	RequestAccessToFilesForProcessWithMessageAndBehaviorAsync(sourceItemList *IIterable[*IStorageItem], processId uint32, auditInfo *IProtectionPolicyAuditInfo, messageFromApp string, behavior ProtectionPolicyRequestAccessBehavior) *IAsyncOperation[ProtectionPolicyEvaluationResult]
	IsFileProtectionRequiredAsync(target *IStorageItem, identity string) *IAsyncOperation[bool]
	IsFileProtectionRequiredForNewFileAsync(parentFolder *IStorageFolder, identity string, desiredName string) *IAsyncOperation[bool]
	Get_PrimaryManagedIdentity() string
	GetPrimaryManagedIdentityForIdentity(identity string) string
}

type IProtectionPolicyManagerStatics4Vtbl struct {
	win32.IInspectableVtbl
	IsRoamableProtectionEnabled                               uintptr
	RequestAccessWithBehaviorAsync                            uintptr
	RequestAccessForAppWithBehaviorAsync                      uintptr
	RequestAccessToFilesForAppAsync                           uintptr
	RequestAccessToFilesForAppWithMessageAndBehaviorAsync     uintptr
	RequestAccessToFilesForProcessAsync                       uintptr
	RequestAccessToFilesForProcessWithMessageAndBehaviorAsync uintptr
	IsFileProtectionRequiredAsync                             uintptr
	IsFileProtectionRequiredForNewFileAsync                   uintptr
	Get_PrimaryManagedIdentity                                uintptr
	GetPrimaryManagedIdentityForIdentity                      uintptr
}

type IProtectionPolicyManagerStatics4 struct {
	win32.IInspectable
}

func (this *IProtectionPolicyManagerStatics4) Vtbl() *IProtectionPolicyManagerStatics4Vtbl {
	return (*IProtectionPolicyManagerStatics4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProtectionPolicyManagerStatics4) IsRoamableProtectionEnabled(identity string) bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsRoamableProtectionEnabled, uintptr(unsafe.Pointer(this)), NewHStr(identity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IProtectionPolicyManagerStatics4) RequestAccessWithBehaviorAsync(sourceIdentity string, targetIdentity string, auditInfo *IProtectionPolicyAuditInfo, messageFromApp string, behavior ProtectionPolicyRequestAccessBehavior) *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessWithBehaviorAsync, uintptr(unsafe.Pointer(this)), NewHStr(sourceIdentity).Ptr, NewHStr(targetIdentity).Ptr, uintptr(unsafe.Pointer(auditInfo)), NewHStr(messageFromApp).Ptr, uintptr(behavior), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics4) RequestAccessForAppWithBehaviorAsync(sourceIdentity string, appPackageFamilyName string, auditInfo *IProtectionPolicyAuditInfo, messageFromApp string, behavior ProtectionPolicyRequestAccessBehavior) *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessForAppWithBehaviorAsync, uintptr(unsafe.Pointer(this)), NewHStr(sourceIdentity).Ptr, NewHStr(appPackageFamilyName).Ptr, uintptr(unsafe.Pointer(auditInfo)), NewHStr(messageFromApp).Ptr, uintptr(behavior), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics4) RequestAccessToFilesForAppAsync(sourceItemList *IIterable[*IStorageItem], appPackageFamilyName string, auditInfo *IProtectionPolicyAuditInfo) *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessToFilesForAppAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceItemList)), NewHStr(appPackageFamilyName).Ptr, uintptr(unsafe.Pointer(auditInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics4) RequestAccessToFilesForAppWithMessageAndBehaviorAsync(sourceItemList *IIterable[*IStorageItem], appPackageFamilyName string, auditInfo *IProtectionPolicyAuditInfo, messageFromApp string, behavior ProtectionPolicyRequestAccessBehavior) *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessToFilesForAppWithMessageAndBehaviorAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceItemList)), NewHStr(appPackageFamilyName).Ptr, uintptr(unsafe.Pointer(auditInfo)), NewHStr(messageFromApp).Ptr, uintptr(behavior), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics4) RequestAccessToFilesForProcessAsync(sourceItemList *IIterable[*IStorageItem], processId uint32, auditInfo *IProtectionPolicyAuditInfo) *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessToFilesForProcessAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceItemList)), uintptr(processId), uintptr(unsafe.Pointer(auditInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics4) RequestAccessToFilesForProcessWithMessageAndBehaviorAsync(sourceItemList *IIterable[*IStorageItem], processId uint32, auditInfo *IProtectionPolicyAuditInfo, messageFromApp string, behavior ProtectionPolicyRequestAccessBehavior) *IAsyncOperation[ProtectionPolicyEvaluationResult] {
	var _result *IAsyncOperation[ProtectionPolicyEvaluationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAccessToFilesForProcessWithMessageAndBehaviorAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceItemList)), uintptr(processId), uintptr(unsafe.Pointer(auditInfo)), NewHStr(messageFromApp).Ptr, uintptr(behavior), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics4) IsFileProtectionRequiredAsync(target *IStorageItem, identity string) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsFileProtectionRequiredAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(target)), NewHStr(identity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics4) IsFileProtectionRequiredForNewFileAsync(parentFolder *IStorageFolder, identity string, desiredName string) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().IsFileProtectionRequiredForNewFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(parentFolder)), NewHStr(identity).Ptr, NewHStr(desiredName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IProtectionPolicyManagerStatics4) Get_PrimaryManagedIdentity() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PrimaryManagedIdentity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IProtectionPolicyManagerStatics4) GetPrimaryManagedIdentityForIdentity(identity string) string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPrimaryManagedIdentityForIdentity, uintptr(unsafe.Pointer(this)), NewHStr(identity).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// FA4EA8E9-EF13-405A-B12C-D7348C6F41FC
var IID_IThreadNetworkContext = syscall.GUID{0xFA4EA8E9, 0xEF13, 0x405A,
	[8]byte{0xB1, 0x2C, 0xD7, 0x34, 0x8C, 0x6F, 0x41, 0xFC}}

type IThreadNetworkContextInterface interface {
	win32.IInspectableInterface
}

type IThreadNetworkContextVtbl struct {
	win32.IInspectableVtbl
}

type IThreadNetworkContext struct {
	win32.IInspectable
}

func (this *IThreadNetworkContext) Vtbl() *IThreadNetworkContextVtbl {
	return (*IThreadNetworkContextVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// classes

type BufferProtectUnprotectResult struct {
	RtClass
	*IBufferProtectUnprotectResult
}

type DataProtectionInfo struct {
	RtClass
	*IDataProtectionInfo
}

type DataProtectionManager struct {
	RtClass
}

func NewIDataProtectionManagerStatics() *IDataProtectionManagerStatics {
	var p *IDataProtectionManagerStatics
	hs := NewHStr("Windows.Security.EnterpriseData.DataProtectionManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IDataProtectionManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type FileProtectionInfo struct {
	RtClass
	*IFileProtectionInfo
}

type FileProtectionManager struct {
	RtClass
}

func NewIFileProtectionManagerStatics3() *IFileProtectionManagerStatics3 {
	var p *IFileProtectionManagerStatics3
	hs := NewHStr("Windows.Security.EnterpriseData.FileProtectionManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IFileProtectionManagerStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIFileProtectionManagerStatics2() *IFileProtectionManagerStatics2 {
	var p *IFileProtectionManagerStatics2
	hs := NewHStr("Windows.Security.EnterpriseData.FileProtectionManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IFileProtectionManagerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIFileProtectionManagerStatics() *IFileProtectionManagerStatics {
	var p *IFileProtectionManagerStatics
	hs := NewHStr("Windows.Security.EnterpriseData.FileProtectionManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IFileProtectionManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type FileRevocationManager struct {
	RtClass
}

func NewIFileRevocationManagerStatics() *IFileRevocationManagerStatics {
	var p *IFileRevocationManagerStatics
	hs := NewHStr("Windows.Security.EnterpriseData.FileRevocationManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IFileRevocationManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type FileUnprotectOptions struct {
	RtClass
	*IFileUnprotectOptions
}

func NewFileUnprotectOptions_Create(audit bool) *FileUnprotectOptions {
	hs := NewHStr("Windows.Security.EnterpriseData.FileUnprotectOptions")
	var pFac *IFileUnprotectOptionsFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IFileUnprotectOptionsFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IFileUnprotectOptions
	p = pFac.Create(audit)
	result := &FileUnprotectOptions{
		RtClass:               RtClass{PInspect: &p.IInspectable},
		IFileUnprotectOptions: p,
	}
	com.AddToScope(result)
	return result
}

type ProtectedAccessResumedEventArgs struct {
	RtClass
	*IProtectedAccessResumedEventArgs
}

type ProtectedAccessSuspendingEventArgs struct {
	RtClass
	*IProtectedAccessSuspendingEventArgs
}

type ProtectedContainerExportResult struct {
	RtClass
	*IProtectedContainerExportResult
}

type ProtectedContainerImportResult struct {
	RtClass
	*IProtectedContainerImportResult
}

type ProtectedContentRevokedEventArgs struct {
	RtClass
	*IProtectedContentRevokedEventArgs
}

type ProtectedFileCreateResult struct {
	RtClass
	*IProtectedFileCreateResult
}

type ProtectionPolicyAuditInfo struct {
	RtClass
	*IProtectionPolicyAuditInfo
}

func NewProtectionPolicyAuditInfo_Create(action ProtectionPolicyAuditAction, dataDescription string, sourceDescription string, targetDescription string) *ProtectionPolicyAuditInfo {
	hs := NewHStr("Windows.Security.EnterpriseData.ProtectionPolicyAuditInfo")
	var pFac *IProtectionPolicyAuditInfoFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProtectionPolicyAuditInfoFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IProtectionPolicyAuditInfo
	p = pFac.Create(action, dataDescription, sourceDescription, targetDescription)
	result := &ProtectionPolicyAuditInfo{
		RtClass:                    RtClass{PInspect: &p.IInspectable},
		IProtectionPolicyAuditInfo: p,
	}
	com.AddToScope(result)
	return result
}

func NewProtectionPolicyAuditInfo_CreateWithActionAndDataDescription(action ProtectionPolicyAuditAction, dataDescription string) *ProtectionPolicyAuditInfo {
	hs := NewHStr("Windows.Security.EnterpriseData.ProtectionPolicyAuditInfo")
	var pFac *IProtectionPolicyAuditInfoFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProtectionPolicyAuditInfoFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *IProtectionPolicyAuditInfo
	p = pFac.CreateWithActionAndDataDescription(action, dataDescription)
	result := &ProtectionPolicyAuditInfo{
		RtClass:                    RtClass{PInspect: &p.IInspectable},
		IProtectionPolicyAuditInfo: p,
	}
	com.AddToScope(result)
	return result
}

type ProtectionPolicyManager struct {
	RtClass
	*IProtectionPolicyManager
}

func NewIProtectionPolicyManagerStatics4() *IProtectionPolicyManagerStatics4 {
	var p *IProtectionPolicyManagerStatics4
	hs := NewHStr("Windows.Security.EnterpriseData.ProtectionPolicyManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProtectionPolicyManagerStatics4, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIProtectionPolicyManagerStatics() *IProtectionPolicyManagerStatics {
	var p *IProtectionPolicyManagerStatics
	hs := NewHStr("Windows.Security.EnterpriseData.ProtectionPolicyManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProtectionPolicyManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIProtectionPolicyManagerStatics2() *IProtectionPolicyManagerStatics2 {
	var p *IProtectionPolicyManagerStatics2
	hs := NewHStr("Windows.Security.EnterpriseData.ProtectionPolicyManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProtectionPolicyManagerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIProtectionPolicyManagerStatics3() *IProtectionPolicyManagerStatics3 {
	var p *IProtectionPolicyManagerStatics3
	hs := NewHStr("Windows.Security.EnterpriseData.ProtectionPolicyManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IProtectionPolicyManagerStatics3, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type ThreadNetworkContext struct {
	RtClass
	*IThreadNetworkContext
}
