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
type SpeechContinuousRecognitionMode int32

const (
	SpeechContinuousRecognitionMode_Default            SpeechContinuousRecognitionMode = 0
	SpeechContinuousRecognitionMode_PauseOnRecognition SpeechContinuousRecognitionMode = 1
)

// enum
type SpeechRecognitionAudioProblem int32

const (
	SpeechRecognitionAudioProblem_None     SpeechRecognitionAudioProblem = 0
	SpeechRecognitionAudioProblem_TooNoisy SpeechRecognitionAudioProblem = 1
	SpeechRecognitionAudioProblem_NoSignal SpeechRecognitionAudioProblem = 2
	SpeechRecognitionAudioProblem_TooLoud  SpeechRecognitionAudioProblem = 3
	SpeechRecognitionAudioProblem_TooQuiet SpeechRecognitionAudioProblem = 4
	SpeechRecognitionAudioProblem_TooFast  SpeechRecognitionAudioProblem = 5
	SpeechRecognitionAudioProblem_TooSlow  SpeechRecognitionAudioProblem = 6
)

// enum
type SpeechRecognitionConfidence int32

const (
	SpeechRecognitionConfidence_High     SpeechRecognitionConfidence = 0
	SpeechRecognitionConfidence_Medium   SpeechRecognitionConfidence = 1
	SpeechRecognitionConfidence_Low      SpeechRecognitionConfidence = 2
	SpeechRecognitionConfidence_Rejected SpeechRecognitionConfidence = 3
)

// enum
type SpeechRecognitionConstraintProbability int32

const (
	SpeechRecognitionConstraintProbability_Default SpeechRecognitionConstraintProbability = 0
	SpeechRecognitionConstraintProbability_Min     SpeechRecognitionConstraintProbability = 1
	SpeechRecognitionConstraintProbability_Max     SpeechRecognitionConstraintProbability = 2
)

// enum
type SpeechRecognitionConstraintType int32

const (
	SpeechRecognitionConstraintType_Topic                  SpeechRecognitionConstraintType = 0
	SpeechRecognitionConstraintType_List                   SpeechRecognitionConstraintType = 1
	SpeechRecognitionConstraintType_Grammar                SpeechRecognitionConstraintType = 2
	SpeechRecognitionConstraintType_VoiceCommandDefinition SpeechRecognitionConstraintType = 3
)

// enum
type SpeechRecognitionResultStatus int32

const (
	SpeechRecognitionResultStatus_Success                   SpeechRecognitionResultStatus = 0
	SpeechRecognitionResultStatus_TopicLanguageNotSupported SpeechRecognitionResultStatus = 1
	SpeechRecognitionResultStatus_GrammarLanguageMismatch   SpeechRecognitionResultStatus = 2
	SpeechRecognitionResultStatus_GrammarCompilationFailure SpeechRecognitionResultStatus = 3
	SpeechRecognitionResultStatus_AudioQualityFailure       SpeechRecognitionResultStatus = 4
	SpeechRecognitionResultStatus_UserCanceled              SpeechRecognitionResultStatus = 5
	SpeechRecognitionResultStatus_Unknown                   SpeechRecognitionResultStatus = 6
	SpeechRecognitionResultStatus_TimeoutExceeded           SpeechRecognitionResultStatus = 7
	SpeechRecognitionResultStatus_PauseLimitExceeded        SpeechRecognitionResultStatus = 8
	SpeechRecognitionResultStatus_NetworkFailure            SpeechRecognitionResultStatus = 9
	SpeechRecognitionResultStatus_MicrophoneUnavailable     SpeechRecognitionResultStatus = 10
)

// enum
type SpeechRecognitionScenario int32

const (
	SpeechRecognitionScenario_WebSearch   SpeechRecognitionScenario = 0
	SpeechRecognitionScenario_Dictation   SpeechRecognitionScenario = 1
	SpeechRecognitionScenario_FormFilling SpeechRecognitionScenario = 2
)

// enum
type SpeechRecognizerState int32

const (
	SpeechRecognizerState_Idle           SpeechRecognizerState = 0
	SpeechRecognizerState_Capturing      SpeechRecognizerState = 1
	SpeechRecognizerState_Processing     SpeechRecognizerState = 2
	SpeechRecognizerState_SoundStarted   SpeechRecognizerState = 3
	SpeechRecognizerState_SoundEnded     SpeechRecognizerState = 4
	SpeechRecognizerState_SpeechDetected SpeechRecognizerState = 5
	SpeechRecognizerState_Paused         SpeechRecognizerState = 6
)

// interfaces

// E3D069BB-E30C-5E18-424B-7FBE81F8FBD0
var IID_ISpeechContinuousRecognitionCompletedEventArgs = syscall.GUID{0xE3D069BB, 0xE30C, 0x5E18,
	[8]byte{0x42, 0x4B, 0x7F, 0xBE, 0x81, 0xF8, 0xFB, 0xD0}}

type ISpeechContinuousRecognitionCompletedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Status() SpeechRecognitionResultStatus
}

type ISpeechContinuousRecognitionCompletedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type ISpeechContinuousRecognitionCompletedEventArgs struct {
	win32.IInspectable
}

func (this *ISpeechContinuousRecognitionCompletedEventArgs) Vtbl() *ISpeechContinuousRecognitionCompletedEventArgsVtbl {
	return (*ISpeechContinuousRecognitionCompletedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechContinuousRecognitionCompletedEventArgs) Get_Status() SpeechRecognitionResultStatus {
	var _result SpeechRecognitionResultStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 19091E1E-6E7E-5A46-40FB-76594F786504
var IID_ISpeechContinuousRecognitionResultGeneratedEventArgs = syscall.GUID{0x19091E1E, 0x6E7E, 0x5A46,
	[8]byte{0x40, 0xFB, 0x76, 0x59, 0x4F, 0x78, 0x65, 0x04}}

type ISpeechContinuousRecognitionResultGeneratedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Result() *ISpeechRecognitionResult
}

type ISpeechContinuousRecognitionResultGeneratedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Result uintptr
}

type ISpeechContinuousRecognitionResultGeneratedEventArgs struct {
	win32.IInspectable
}

func (this *ISpeechContinuousRecognitionResultGeneratedEventArgs) Vtbl() *ISpeechContinuousRecognitionResultGeneratedEventArgsVtbl {
	return (*ISpeechContinuousRecognitionResultGeneratedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechContinuousRecognitionResultGeneratedEventArgs) Get_Result() *ISpeechRecognitionResult {
	var _result *ISpeechRecognitionResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Result, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 6A213C04-6614-49F8-99A2-B5E9B3A085C8
var IID_ISpeechContinuousRecognitionSession = syscall.GUID{0x6A213C04, 0x6614, 0x49F8,
	[8]byte{0x99, 0xA2, 0xB5, 0xE9, 0xB3, 0xA0, 0x85, 0xC8}}

type ISpeechContinuousRecognitionSessionInterface interface {
	win32.IInspectableInterface
	Get_AutoStopSilenceTimeout() TimeSpan
	Put_AutoStopSilenceTimeout(value TimeSpan)
	StartAsync() *IAsyncAction
	StartWithModeAsync(mode SpeechContinuousRecognitionMode) *IAsyncAction
	StopAsync() *IAsyncAction
	CancelAsync() *IAsyncAction
	PauseAsync() *IAsyncAction
	Resume()
	Add_Completed(value TypedEventHandler[*ISpeechContinuousRecognitionSession, *ISpeechContinuousRecognitionCompletedEventArgs]) EventRegistrationToken
	Remove_Completed(value EventRegistrationToken)
	Add_ResultGenerated(value TypedEventHandler[*ISpeechContinuousRecognitionSession, *ISpeechContinuousRecognitionResultGeneratedEventArgs]) EventRegistrationToken
	Remove_ResultGenerated(value EventRegistrationToken)
}

type ISpeechContinuousRecognitionSessionVtbl struct {
	win32.IInspectableVtbl
	Get_AutoStopSilenceTimeout uintptr
	Put_AutoStopSilenceTimeout uintptr
	StartAsync                 uintptr
	StartWithModeAsync         uintptr
	StopAsync                  uintptr
	CancelAsync                uintptr
	PauseAsync                 uintptr
	Resume                     uintptr
	Add_Completed              uintptr
	Remove_Completed           uintptr
	Add_ResultGenerated        uintptr
	Remove_ResultGenerated     uintptr
}

type ISpeechContinuousRecognitionSession struct {
	win32.IInspectable
}

func (this *ISpeechContinuousRecognitionSession) Vtbl() *ISpeechContinuousRecognitionSessionVtbl {
	return (*ISpeechContinuousRecognitionSessionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechContinuousRecognitionSession) Get_AutoStopSilenceTimeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoStopSilenceTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechContinuousRecognitionSession) Put_AutoStopSilenceTimeout(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoStopSilenceTimeout, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISpeechContinuousRecognitionSession) StartAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechContinuousRecognitionSession) StartWithModeAsync(mode SpeechContinuousRecognitionMode) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StartWithModeAsync, uintptr(unsafe.Pointer(this)), uintptr(mode), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechContinuousRecognitionSession) StopAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechContinuousRecognitionSession) CancelAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CancelAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechContinuousRecognitionSession) PauseAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PauseAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechContinuousRecognitionSession) Resume() {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Resume, uintptr(unsafe.Pointer(this)))
	_ = _hr
}

func (this *ISpeechContinuousRecognitionSession) Add_Completed(value TypedEventHandler[*ISpeechContinuousRecognitionSession, *ISpeechContinuousRecognitionCompletedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_Completed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechContinuousRecognitionSession) Remove_Completed(value EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_Completed, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISpeechContinuousRecognitionSession) Add_ResultGenerated(value TypedEventHandler[*ISpeechContinuousRecognitionSession, *ISpeechContinuousRecognitionResultGeneratedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ResultGenerated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechContinuousRecognitionSession) Remove_ResultGenerated(value EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ResultGenerated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 407E6C5D-6AC7-4DA4-9CC1-2FCE32CF7489
var IID_ISpeechRecognitionCompilationResult = syscall.GUID{0x407E6C5D, 0x6AC7, 0x4DA4,
	[8]byte{0x9C, 0xC1, 0x2F, 0xCE, 0x32, 0xCF, 0x74, 0x89}}

type ISpeechRecognitionCompilationResultInterface interface {
	win32.IInspectableInterface
	Get_Status() SpeechRecognitionResultStatus
}

type ISpeechRecognitionCompilationResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status uintptr
}

type ISpeechRecognitionCompilationResult struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionCompilationResult) Vtbl() *ISpeechRecognitionCompilationResultVtbl {
	return (*ISpeechRecognitionCompilationResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionCompilationResult) Get_Status() SpeechRecognitionResultStatus {
	var _result SpeechRecognitionResultStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 79AC1628-4D68-43C4-8911-40DC4101B55B
var IID_ISpeechRecognitionConstraint = syscall.GUID{0x79AC1628, 0x4D68, 0x43C4,
	[8]byte{0x89, 0x11, 0x40, 0xDC, 0x41, 0x01, 0xB5, 0x5B}}

type ISpeechRecognitionConstraintInterface interface {
	win32.IInspectableInterface
	Get_IsEnabled() bool
	Put_IsEnabled(value bool)
	Get_Tag() string
	Put_Tag(value string)
	Get_Type() SpeechRecognitionConstraintType
	Get_Probability() SpeechRecognitionConstraintProbability
	Put_Probability(value SpeechRecognitionConstraintProbability)
}

type ISpeechRecognitionConstraintVtbl struct {
	win32.IInspectableVtbl
	Get_IsEnabled   uintptr
	Put_IsEnabled   uintptr
	Get_Tag         uintptr
	Put_Tag         uintptr
	Get_Type        uintptr
	Get_Probability uintptr
	Put_Probability uintptr
}

type ISpeechRecognitionConstraint struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionConstraint) Vtbl() *ISpeechRecognitionConstraintVtbl {
	return (*ISpeechRecognitionConstraintVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionConstraint) Get_IsEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognitionConstraint) Put_IsEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISpeechRecognitionConstraint) Get_Tag() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Tag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpeechRecognitionConstraint) Put_Tag(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Tag, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISpeechRecognitionConstraint) Get_Type() SpeechRecognitionConstraintType {
	var _result SpeechRecognitionConstraintType
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Type, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognitionConstraint) Get_Probability() SpeechRecognitionConstraintProbability {
	var _result SpeechRecognitionConstraintProbability
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Probability, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognitionConstraint) Put_Probability(value SpeechRecognitionConstraintProbability) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Probability, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// B5031A8F-85CA-4FA4-B11A-474FC41B3835
var IID_ISpeechRecognitionGrammarFileConstraint = syscall.GUID{0xB5031A8F, 0x85CA, 0x4FA4,
	[8]byte{0xB1, 0x1A, 0x47, 0x4F, 0xC4, 0x1B, 0x38, 0x35}}

type ISpeechRecognitionGrammarFileConstraintInterface interface {
	win32.IInspectableInterface
	Get_GrammarFile() *IStorageFile
}

type ISpeechRecognitionGrammarFileConstraintVtbl struct {
	win32.IInspectableVtbl
	Get_GrammarFile uintptr
}

type ISpeechRecognitionGrammarFileConstraint struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionGrammarFileConstraint) Vtbl() *ISpeechRecognitionGrammarFileConstraintVtbl {
	return (*ISpeechRecognitionGrammarFileConstraintVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionGrammarFileConstraint) Get_GrammarFile() *IStorageFile {
	var _result *IStorageFile
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_GrammarFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3DA770EB-C479-4C27-9F19-89974EF392D1
var IID_ISpeechRecognitionGrammarFileConstraintFactory = syscall.GUID{0x3DA770EB, 0xC479, 0x4C27,
	[8]byte{0x9F, 0x19, 0x89, 0x97, 0x4E, 0xF3, 0x92, 0xD1}}

type ISpeechRecognitionGrammarFileConstraintFactoryInterface interface {
	win32.IInspectableInterface
	Create(file *IStorageFile) *ISpeechRecognitionGrammarFileConstraint
	CreateWithTag(file *IStorageFile, tag string) *ISpeechRecognitionGrammarFileConstraint
}

type ISpeechRecognitionGrammarFileConstraintFactoryVtbl struct {
	win32.IInspectableVtbl
	Create        uintptr
	CreateWithTag uintptr
}

type ISpeechRecognitionGrammarFileConstraintFactory struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionGrammarFileConstraintFactory) Vtbl() *ISpeechRecognitionGrammarFileConstraintFactoryVtbl {
	return (*ISpeechRecognitionGrammarFileConstraintFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionGrammarFileConstraintFactory) Create(file *IStorageFile) *ISpeechRecognitionGrammarFileConstraint {
	var _result *ISpeechRecognitionGrammarFileConstraint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognitionGrammarFileConstraintFactory) CreateWithTag(file *IStorageFile, tag string) *ISpeechRecognitionGrammarFileConstraint {
	var _result *ISpeechRecognitionGrammarFileConstraint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithTag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), NewHStr(tag).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7A7B25B0-99C5-4F7D-BF84-10AA1302B634
var IID_ISpeechRecognitionHypothesis = syscall.GUID{0x7A7B25B0, 0x99C5, 0x4F7D,
	[8]byte{0xBF, 0x84, 0x10, 0xAA, 0x13, 0x02, 0xB6, 0x34}}

type ISpeechRecognitionHypothesisInterface interface {
	win32.IInspectableInterface
	Get_Text() string
}

type ISpeechRecognitionHypothesisVtbl struct {
	win32.IInspectableVtbl
	Get_Text uintptr
}

type ISpeechRecognitionHypothesis struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionHypothesis) Vtbl() *ISpeechRecognitionHypothesisVtbl {
	return (*ISpeechRecognitionHypothesisVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionHypothesis) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 55161A7A-8023-5866-411D-1213BB271476
var IID_ISpeechRecognitionHypothesisGeneratedEventArgs = syscall.GUID{0x55161A7A, 0x8023, 0x5866,
	[8]byte{0x41, 0x1D, 0x12, 0x13, 0xBB, 0x27, 0x14, 0x76}}

type ISpeechRecognitionHypothesisGeneratedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Hypothesis() *ISpeechRecognitionHypothesis
}

type ISpeechRecognitionHypothesisGeneratedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Hypothesis uintptr
}

type ISpeechRecognitionHypothesisGeneratedEventArgs struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionHypothesisGeneratedEventArgs) Vtbl() *ISpeechRecognitionHypothesisGeneratedEventArgsVtbl {
	return (*ISpeechRecognitionHypothesisGeneratedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionHypothesisGeneratedEventArgs) Get_Hypothesis() *ISpeechRecognitionHypothesis {
	var _result *ISpeechRecognitionHypothesis
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Hypothesis, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 09C487E9-E4AD-4526-81F2-4946FB481D98
var IID_ISpeechRecognitionListConstraint = syscall.GUID{0x09C487E9, 0xE4AD, 0x4526,
	[8]byte{0x81, 0xF2, 0x49, 0x46, 0xFB, 0x48, 0x1D, 0x98}}

type ISpeechRecognitionListConstraintInterface interface {
	win32.IInspectableInterface
	Get_Commands() *IVector[string]
}

type ISpeechRecognitionListConstraintVtbl struct {
	win32.IInspectableVtbl
	Get_Commands uintptr
}

type ISpeechRecognitionListConstraint struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionListConstraint) Vtbl() *ISpeechRecognitionListConstraintVtbl {
	return (*ISpeechRecognitionListConstraintVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionListConstraint) Get_Commands() *IVector[string] {
	var _result *IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Commands, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 40F3CDC7-562A-426A-9F3B-3B4E282BE1D5
var IID_ISpeechRecognitionListConstraintFactory = syscall.GUID{0x40F3CDC7, 0x562A, 0x426A,
	[8]byte{0x9F, 0x3B, 0x3B, 0x4E, 0x28, 0x2B, 0xE1, 0xD5}}

type ISpeechRecognitionListConstraintFactoryInterface interface {
	win32.IInspectableInterface
	Create(commands *IIterable[string]) *ISpeechRecognitionListConstraint
	CreateWithTag(commands *IIterable[string], tag string) *ISpeechRecognitionListConstraint
}

type ISpeechRecognitionListConstraintFactoryVtbl struct {
	win32.IInspectableVtbl
	Create        uintptr
	CreateWithTag uintptr
}

type ISpeechRecognitionListConstraintFactory struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionListConstraintFactory) Vtbl() *ISpeechRecognitionListConstraintFactoryVtbl {
	return (*ISpeechRecognitionListConstraintFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionListConstraintFactory) Create(commands *IIterable[string]) *ISpeechRecognitionListConstraint {
	var _result *ISpeechRecognitionListConstraint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(commands)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognitionListConstraintFactory) CreateWithTag(commands *IIterable[string], tag string) *ISpeechRecognitionListConstraint {
	var _result *ISpeechRecognitionListConstraint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithTag, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(commands)), NewHStr(tag).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 4FE24105-8C3A-4C7E-8D0A-5BD4F5B14AD8
var IID_ISpeechRecognitionQualityDegradingEventArgs = syscall.GUID{0x4FE24105, 0x8C3A, 0x4C7E,
	[8]byte{0x8D, 0x0A, 0x5B, 0xD4, 0xF5, 0xB1, 0x4A, 0xD8}}

type ISpeechRecognitionQualityDegradingEventArgsInterface interface {
	win32.IInspectableInterface
	Get_Problem() SpeechRecognitionAudioProblem
}

type ISpeechRecognitionQualityDegradingEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_Problem uintptr
}

type ISpeechRecognitionQualityDegradingEventArgs struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionQualityDegradingEventArgs) Vtbl() *ISpeechRecognitionQualityDegradingEventArgsVtbl {
	return (*ISpeechRecognitionQualityDegradingEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionQualityDegradingEventArgs) Get_Problem() SpeechRecognitionAudioProblem {
	var _result SpeechRecognitionAudioProblem
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Problem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 4E303157-034E-4652-857E-D0454CC4BEEC
var IID_ISpeechRecognitionResult = syscall.GUID{0x4E303157, 0x034E, 0x4652,
	[8]byte{0x85, 0x7E, 0xD0, 0x45, 0x4C, 0xC4, 0xBE, 0xEC}}

type ISpeechRecognitionResultInterface interface {
	win32.IInspectableInterface
	Get_Status() SpeechRecognitionResultStatus
	Get_Text() string
	Get_Confidence() SpeechRecognitionConfidence
	Get_SemanticInterpretation() *ISpeechRecognitionSemanticInterpretation
	GetAlternates(maxAlternates uint32) *IVectorView[*ISpeechRecognitionResult]
	Get_Constraint() *ISpeechRecognitionConstraint
	Get_RulePath() *IVectorView[string]
	Get_RawConfidence() float64
}

type ISpeechRecognitionResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status                 uintptr
	Get_Text                   uintptr
	Get_Confidence             uintptr
	Get_SemanticInterpretation uintptr
	GetAlternates              uintptr
	Get_Constraint             uintptr
	Get_RulePath               uintptr
	Get_RawConfidence          uintptr
}

type ISpeechRecognitionResult struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionResult) Vtbl() *ISpeechRecognitionResultVtbl {
	return (*ISpeechRecognitionResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionResult) Get_Status() SpeechRecognitionResultStatus {
	var _result SpeechRecognitionResultStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognitionResult) Get_Text() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Text, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpeechRecognitionResult) Get_Confidence() SpeechRecognitionConfidence {
	var _result SpeechRecognitionConfidence
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Confidence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognitionResult) Get_SemanticInterpretation() *ISpeechRecognitionSemanticInterpretation {
	var _result *ISpeechRecognitionSemanticInterpretation
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SemanticInterpretation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognitionResult) GetAlternates(maxAlternates uint32) *IVectorView[*ISpeechRecognitionResult] {
	var _result *IVectorView[*ISpeechRecognitionResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAlternates, uintptr(unsafe.Pointer(this)), uintptr(maxAlternates), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognitionResult) Get_Constraint() *ISpeechRecognitionConstraint {
	var _result *ISpeechRecognitionConstraint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Constraint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognitionResult) Get_RulePath() *IVectorView[string] {
	var _result *IVectorView[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RulePath, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognitionResult) Get_RawConfidence() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RawConfidence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AF7ED1BA-451B-4166-A0C1-1FFE84032D03
var IID_ISpeechRecognitionResult2 = syscall.GUID{0xAF7ED1BA, 0x451B, 0x4166,
	[8]byte{0xA0, 0xC1, 0x1F, 0xFE, 0x84, 0x03, 0x2D, 0x03}}

type ISpeechRecognitionResult2Interface interface {
	win32.IInspectableInterface
	Get_PhraseStartTime() DateTime
	Get_PhraseDuration() TimeSpan
}

type ISpeechRecognitionResult2Vtbl struct {
	win32.IInspectableVtbl
	Get_PhraseStartTime uintptr
	Get_PhraseDuration  uintptr
}

type ISpeechRecognitionResult2 struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionResult2) Vtbl() *ISpeechRecognitionResult2Vtbl {
	return (*ISpeechRecognitionResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionResult2) Get_PhraseStartTime() DateTime {
	var _result DateTime
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhraseStartTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognitionResult2) Get_PhraseDuration() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PhraseDuration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// AAE1DA9B-7E32-4C1F-89FE-0C65F486F52E
var IID_ISpeechRecognitionSemanticInterpretation = syscall.GUID{0xAAE1DA9B, 0x7E32, 0x4C1F,
	[8]byte{0x89, 0xFE, 0x0C, 0x65, 0xF4, 0x86, 0xF5, 0x2E}}

type ISpeechRecognitionSemanticInterpretationInterface interface {
	win32.IInspectableInterface
	Get_Properties() *IMapView[string, *IVectorView[string]]
}

type ISpeechRecognitionSemanticInterpretationVtbl struct {
	win32.IInspectableVtbl
	Get_Properties uintptr
}

type ISpeechRecognitionSemanticInterpretation struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionSemanticInterpretation) Vtbl() *ISpeechRecognitionSemanticInterpretationVtbl {
	return (*ISpeechRecognitionSemanticInterpretationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionSemanticInterpretation) Get_Properties() *IMapView[string, *IVectorView[string]] {
	var _result *IMapView[string, *IVectorView[string]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Properties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BF6FDF19-825D-4E69-A681-36E48CF1C93E
var IID_ISpeechRecognitionTopicConstraint = syscall.GUID{0xBF6FDF19, 0x825D, 0x4E69,
	[8]byte{0xA6, 0x81, 0x36, 0xE4, 0x8C, 0xF1, 0xC9, 0x3E}}

type ISpeechRecognitionTopicConstraintInterface interface {
	win32.IInspectableInterface
	Get_Scenario() SpeechRecognitionScenario
	Get_TopicHint() string
}

type ISpeechRecognitionTopicConstraintVtbl struct {
	win32.IInspectableVtbl
	Get_Scenario  uintptr
	Get_TopicHint uintptr
}

type ISpeechRecognitionTopicConstraint struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionTopicConstraint) Vtbl() *ISpeechRecognitionTopicConstraintVtbl {
	return (*ISpeechRecognitionTopicConstraintVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionTopicConstraint) Get_Scenario() SpeechRecognitionScenario {
	var _result SpeechRecognitionScenario
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Scenario, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognitionTopicConstraint) Get_TopicHint() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TopicHint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// 6E6863DF-EC05-47D7-A5DF-56A3431E58D2
var IID_ISpeechRecognitionTopicConstraintFactory = syscall.GUID{0x6E6863DF, 0xEC05, 0x47D7,
	[8]byte{0xA5, 0xDF, 0x56, 0xA3, 0x43, 0x1E, 0x58, 0xD2}}

type ISpeechRecognitionTopicConstraintFactoryInterface interface {
	win32.IInspectableInterface
	Create(scenario SpeechRecognitionScenario, topicHint string) *ISpeechRecognitionTopicConstraint
	CreateWithTag(scenario SpeechRecognitionScenario, topicHint string, tag string) *ISpeechRecognitionTopicConstraint
}

type ISpeechRecognitionTopicConstraintFactoryVtbl struct {
	win32.IInspectableVtbl
	Create        uintptr
	CreateWithTag uintptr
}

type ISpeechRecognitionTopicConstraintFactory struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionTopicConstraintFactory) Vtbl() *ISpeechRecognitionTopicConstraintFactoryVtbl {
	return (*ISpeechRecognitionTopicConstraintFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognitionTopicConstraintFactory) Create(scenario SpeechRecognitionScenario, topicHint string) *ISpeechRecognitionTopicConstraint {
	var _result *ISpeechRecognitionTopicConstraint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(scenario), NewHStr(topicHint).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognitionTopicConstraintFactory) CreateWithTag(scenario SpeechRecognitionScenario, topicHint string, tag string) *ISpeechRecognitionTopicConstraint {
	var _result *ISpeechRecognitionTopicConstraint
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateWithTag, uintptr(unsafe.Pointer(this)), uintptr(scenario), NewHStr(topicHint).Ptr, NewHStr(tag).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F2791C2B-1EF4-4AE7-9D77-B6FF10B8A3C2
var IID_ISpeechRecognitionVoiceCommandDefinitionConstraint = syscall.GUID{0xF2791C2B, 0x1EF4, 0x4AE7,
	[8]byte{0x9D, 0x77, 0xB6, 0xFF, 0x10, 0xB8, 0xA3, 0xC2}}

type ISpeechRecognitionVoiceCommandDefinitionConstraintInterface interface {
	win32.IInspectableInterface
}

type ISpeechRecognitionVoiceCommandDefinitionConstraintVtbl struct {
	win32.IInspectableVtbl
}

type ISpeechRecognitionVoiceCommandDefinitionConstraint struct {
	win32.IInspectable
}

func (this *ISpeechRecognitionVoiceCommandDefinitionConstraint) Vtbl() *ISpeechRecognitionVoiceCommandDefinitionConstraintVtbl {
	return (*ISpeechRecognitionVoiceCommandDefinitionConstraintVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 0BC3C9CB-C26A-40F2-AEB5-8096B2E48073
var IID_ISpeechRecognizer = syscall.GUID{0x0BC3C9CB, 0xC26A, 0x40F2,
	[8]byte{0xAE, 0xB5, 0x80, 0x96, 0xB2, 0xE4, 0x80, 0x73}}

type ISpeechRecognizerInterface interface {
	win32.IInspectableInterface
	Get_CurrentLanguage() *ILanguage
	Get_Constraints() *IVector[*ISpeechRecognitionConstraint]
	Get_Timeouts() *ISpeechRecognizerTimeouts
	Get_UIOptions() *ISpeechRecognizerUIOptions
	CompileConstraintsAsync() *IAsyncOperation[*ISpeechRecognitionCompilationResult]
	RecognizeAsync() *IAsyncOperation[*ISpeechRecognitionResult]
	RecognizeWithUIAsync() *IAsyncOperation[*ISpeechRecognitionResult]
	Add_RecognitionQualityDegrading(speechRecognitionQualityDegradingHandler TypedEventHandler[*ISpeechRecognizer, *ISpeechRecognitionQualityDegradingEventArgs]) EventRegistrationToken
	Remove_RecognitionQualityDegrading(cookie EventRegistrationToken)
	Add_StateChanged(stateChangedHandler TypedEventHandler[*ISpeechRecognizer, *ISpeechRecognizerStateChangedEventArgs]) EventRegistrationToken
	Remove_StateChanged(cookie EventRegistrationToken)
}

type ISpeechRecognizerVtbl struct {
	win32.IInspectableVtbl
	Get_CurrentLanguage                uintptr
	Get_Constraints                    uintptr
	Get_Timeouts                       uintptr
	Get_UIOptions                      uintptr
	CompileConstraintsAsync            uintptr
	RecognizeAsync                     uintptr
	RecognizeWithUIAsync               uintptr
	Add_RecognitionQualityDegrading    uintptr
	Remove_RecognitionQualityDegrading uintptr
	Add_StateChanged                   uintptr
	Remove_StateChanged                uintptr
}

type ISpeechRecognizer struct {
	win32.IInspectable
}

func (this *ISpeechRecognizer) Vtbl() *ISpeechRecognizerVtbl {
	return (*ISpeechRecognizerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognizer) Get_CurrentLanguage() *ILanguage {
	var _result *ILanguage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognizer) Get_Constraints() *IVector[*ISpeechRecognitionConstraint] {
	var _result *IVector[*ISpeechRecognitionConstraint]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Constraints, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognizer) Get_Timeouts() *ISpeechRecognizerTimeouts {
	var _result *ISpeechRecognizerTimeouts
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Timeouts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognizer) Get_UIOptions() *ISpeechRecognizerUIOptions {
	var _result *ISpeechRecognizerUIOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UIOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognizer) CompileConstraintsAsync() *IAsyncOperation[*ISpeechRecognitionCompilationResult] {
	var _result *IAsyncOperation[*ISpeechRecognitionCompilationResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CompileConstraintsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognizer) RecognizeAsync() *IAsyncOperation[*ISpeechRecognitionResult] {
	var _result *IAsyncOperation[*ISpeechRecognitionResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RecognizeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognizer) RecognizeWithUIAsync() *IAsyncOperation[*ISpeechRecognitionResult] {
	var _result *IAsyncOperation[*ISpeechRecognitionResult]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RecognizeWithUIAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognizer) Add_RecognitionQualityDegrading(speechRecognitionQualityDegradingHandler TypedEventHandler[*ISpeechRecognizer, *ISpeechRecognitionQualityDegradingEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RecognitionQualityDegrading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(speechRecognitionQualityDegradingHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognizer) Remove_RecognitionQualityDegrading(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RecognitionQualityDegrading, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

func (this *ISpeechRecognizer) Add_StateChanged(stateChangedHandler TypedEventHandler[*ISpeechRecognizer, *ISpeechRecognizerStateChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_StateChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(stateChangedHandler))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognizer) Remove_StateChanged(cookie EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_StateChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&cookie)))
	_ = _hr
}

// 63C9BAF1-91E3-4EA4-86A1-7C3867D084A6
var IID_ISpeechRecognizer2 = syscall.GUID{0x63C9BAF1, 0x91E3, 0x4EA4,
	[8]byte{0x86, 0xA1, 0x7C, 0x38, 0x67, 0xD0, 0x84, 0xA6}}

type ISpeechRecognizer2Interface interface {
	win32.IInspectableInterface
	Get_ContinuousRecognitionSession() *ISpeechContinuousRecognitionSession
	Get_State() SpeechRecognizerState
	StopRecognitionAsync() *IAsyncAction
	Add_HypothesisGenerated(value TypedEventHandler[*ISpeechRecognizer, *ISpeechRecognitionHypothesisGeneratedEventArgs]) EventRegistrationToken
	Remove_HypothesisGenerated(value EventRegistrationToken)
}

type ISpeechRecognizer2Vtbl struct {
	win32.IInspectableVtbl
	Get_ContinuousRecognitionSession uintptr
	Get_State                        uintptr
	StopRecognitionAsync             uintptr
	Add_HypothesisGenerated          uintptr
	Remove_HypothesisGenerated       uintptr
}

type ISpeechRecognizer2 struct {
	win32.IInspectable
}

func (this *ISpeechRecognizer2) Vtbl() *ISpeechRecognizer2Vtbl {
	return (*ISpeechRecognizer2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognizer2) Get_ContinuousRecognitionSession() *ISpeechContinuousRecognitionSession {
	var _result *ISpeechContinuousRecognitionSession
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ContinuousRecognitionSession, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognizer2) Get_State() SpeechRecognizerState {
	var _result SpeechRecognizerState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognizer2) StopRecognitionAsync() *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StopRecognitionAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognizer2) Add_HypothesisGenerated(value TypedEventHandler[*ISpeechRecognizer, *ISpeechRecognitionHypothesisGeneratedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_HypothesisGenerated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognizer2) Remove_HypothesisGenerated(value EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_HypothesisGenerated, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 60C488DD-7FB8-4033-AC70-D046F64818E1
var IID_ISpeechRecognizerFactory = syscall.GUID{0x60C488DD, 0x7FB8, 0x4033,
	[8]byte{0xAC, 0x70, 0xD0, 0x46, 0xF6, 0x48, 0x18, 0xE1}}

type ISpeechRecognizerFactoryInterface interface {
	win32.IInspectableInterface
	Create(language *ILanguage) *ISpeechRecognizer
}

type ISpeechRecognizerFactoryVtbl struct {
	win32.IInspectableVtbl
	Create uintptr
}

type ISpeechRecognizerFactory struct {
	win32.IInspectable
}

func (this *ISpeechRecognizerFactory) Vtbl() *ISpeechRecognizerFactoryVtbl {
	return (*ISpeechRecognizerFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognizerFactory) Create(language *ILanguage) *ISpeechRecognizer {
	var _result *ISpeechRecognizer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(language)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 563D4F09-BA03-4BAD-AD81-DDC6C4DAB0C3
var IID_ISpeechRecognizerStateChangedEventArgs = syscall.GUID{0x563D4F09, 0xBA03, 0x4BAD,
	[8]byte{0xAD, 0x81, 0xDD, 0xC6, 0xC4, 0xDA, 0xB0, 0xC3}}

type ISpeechRecognizerStateChangedEventArgsInterface interface {
	win32.IInspectableInterface
	Get_State() SpeechRecognizerState
}

type ISpeechRecognizerStateChangedEventArgsVtbl struct {
	win32.IInspectableVtbl
	Get_State uintptr
}

type ISpeechRecognizerStateChangedEventArgs struct {
	win32.IInspectable
}

func (this *ISpeechRecognizerStateChangedEventArgs) Vtbl() *ISpeechRecognizerStateChangedEventArgsVtbl {
	return (*ISpeechRecognizerStateChangedEventArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognizerStateChangedEventArgs) Get_State() SpeechRecognizerState {
	var _result SpeechRecognizerState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 87A35EAC-A7DC-4B0B-BCC9-24F47C0B7EBF
var IID_ISpeechRecognizerStatics = syscall.GUID{0x87A35EAC, 0xA7DC, 0x4B0B,
	[8]byte{0xBC, 0xC9, 0x24, 0xF4, 0x7C, 0x0B, 0x7E, 0xBF}}

type ISpeechRecognizerStaticsInterface interface {
	win32.IInspectableInterface
	Get_SystemSpeechLanguage() *ILanguage
	Get_SupportedTopicLanguages() *IVectorView[*ILanguage]
	Get_SupportedGrammarLanguages() *IVectorView[*ILanguage]
}

type ISpeechRecognizerStaticsVtbl struct {
	win32.IInspectableVtbl
	Get_SystemSpeechLanguage      uintptr
	Get_SupportedTopicLanguages   uintptr
	Get_SupportedGrammarLanguages uintptr
}

type ISpeechRecognizerStatics struct {
	win32.IInspectable
}

func (this *ISpeechRecognizerStatics) Vtbl() *ISpeechRecognizerStaticsVtbl {
	return (*ISpeechRecognizerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognizerStatics) Get_SystemSpeechLanguage() *ILanguage {
	var _result *ILanguage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SystemSpeechLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognizerStatics) Get_SupportedTopicLanguages() *IVectorView[*ILanguage] {
	var _result *IVectorView[*ILanguage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedTopicLanguages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISpeechRecognizerStatics) Get_SupportedGrammarLanguages() *IVectorView[*ILanguage] {
	var _result *IVectorView[*ILanguage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedGrammarLanguages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1D1B0D95-7565-4EF9-A2F3-BA15162A96CF
var IID_ISpeechRecognizerStatics2 = syscall.GUID{0x1D1B0D95, 0x7565, 0x4EF9,
	[8]byte{0xA2, 0xF3, 0xBA, 0x15, 0x16, 0x2A, 0x96, 0xCF}}

type ISpeechRecognizerStatics2Interface interface {
	win32.IInspectableInterface
	TrySetSystemSpeechLanguageAsync(speechLanguage *ILanguage) *IAsyncOperation[bool]
}

type ISpeechRecognizerStatics2Vtbl struct {
	win32.IInspectableVtbl
	TrySetSystemSpeechLanguageAsync uintptr
}

type ISpeechRecognizerStatics2 struct {
	win32.IInspectable
}

func (this *ISpeechRecognizerStatics2) Vtbl() *ISpeechRecognizerStatics2Vtbl {
	return (*ISpeechRecognizerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognizerStatics2) TrySetSystemSpeechLanguageAsync(speechLanguage *ILanguage) *IAsyncOperation[bool] {
	var _result *IAsyncOperation[bool]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TrySetSystemSpeechLanguageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(speechLanguage)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2EF76FCA-6A3C-4DCA-A153-DF1BC88A79AF
var IID_ISpeechRecognizerTimeouts = syscall.GUID{0x2EF76FCA, 0x6A3C, 0x4DCA,
	[8]byte{0xA1, 0x53, 0xDF, 0x1B, 0xC8, 0x8A, 0x79, 0xAF}}

type ISpeechRecognizerTimeoutsInterface interface {
	win32.IInspectableInterface
	Get_InitialSilenceTimeout() TimeSpan
	Put_InitialSilenceTimeout(value TimeSpan)
	Get_EndSilenceTimeout() TimeSpan
	Put_EndSilenceTimeout(value TimeSpan)
	Get_BabbleTimeout() TimeSpan
	Put_BabbleTimeout(value TimeSpan)
}

type ISpeechRecognizerTimeoutsVtbl struct {
	win32.IInspectableVtbl
	Get_InitialSilenceTimeout uintptr
	Put_InitialSilenceTimeout uintptr
	Get_EndSilenceTimeout     uintptr
	Put_EndSilenceTimeout     uintptr
	Get_BabbleTimeout         uintptr
	Put_BabbleTimeout         uintptr
}

type ISpeechRecognizerTimeouts struct {
	win32.IInspectable
}

func (this *ISpeechRecognizerTimeouts) Vtbl() *ISpeechRecognizerTimeoutsVtbl {
	return (*ISpeechRecognizerTimeoutsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognizerTimeouts) Get_InitialSilenceTimeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InitialSilenceTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognizerTimeouts) Put_InitialSilenceTimeout(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InitialSilenceTimeout, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISpeechRecognizerTimeouts) Get_EndSilenceTimeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_EndSilenceTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognizerTimeouts) Put_EndSilenceTimeout(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_EndSilenceTimeout, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *ISpeechRecognizerTimeouts) Get_BabbleTimeout() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_BabbleTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognizerTimeouts) Put_BabbleTimeout(value TimeSpan) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_BabbleTimeout, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

// 7888D641-B92B-44BA-A25F-D1864630641F
var IID_ISpeechRecognizerUIOptions = syscall.GUID{0x7888D641, 0xB92B, 0x44BA,
	[8]byte{0xA2, 0x5F, 0xD1, 0x86, 0x46, 0x30, 0x64, 0x1F}}

type ISpeechRecognizerUIOptionsInterface interface {
	win32.IInspectableInterface
	Get_ExampleText() string
	Put_ExampleText(value string)
	Get_AudiblePrompt() string
	Put_AudiblePrompt(value string)
	Get_IsReadBackEnabled() bool
	Put_IsReadBackEnabled(value bool)
	Get_ShowConfirmation() bool
	Put_ShowConfirmation(value bool)
}

type ISpeechRecognizerUIOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_ExampleText       uintptr
	Put_ExampleText       uintptr
	Get_AudiblePrompt     uintptr
	Put_AudiblePrompt     uintptr
	Get_IsReadBackEnabled uintptr
	Put_IsReadBackEnabled uintptr
	Get_ShowConfirmation  uintptr
	Put_ShowConfirmation  uintptr
}

type ISpeechRecognizerUIOptions struct {
	win32.IInspectable
}

func (this *ISpeechRecognizerUIOptions) Vtbl() *ISpeechRecognizerUIOptionsVtbl {
	return (*ISpeechRecognizerUIOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpeechRecognizerUIOptions) Get_ExampleText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExampleText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpeechRecognizerUIOptions) Put_ExampleText(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExampleText, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISpeechRecognizerUIOptions) Get_AudiblePrompt() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AudiblePrompt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *ISpeechRecognizerUIOptions) Put_AudiblePrompt(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AudiblePrompt, uintptr(unsafe.Pointer(this)), NewHStr(value).Ptr)
	_ = _hr
}

func (this *ISpeechRecognizerUIOptions) Get_IsReadBackEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReadBackEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognizerUIOptions) Put_IsReadBackEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsReadBackEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ISpeechRecognizerUIOptions) Get_ShowConfirmation() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShowConfirmation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ISpeechRecognizerUIOptions) Put_ShowConfirmation(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ShowConfirmation, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// AA3A8DD5-B6E7-4EE2-BAA9-DD6BACED0A2B
var IID_IVoiceCommandManager = syscall.GUID{0xAA3A8DD5, 0xB6E7, 0x4EE2,
	[8]byte{0xBA, 0xA9, 0xDD, 0x6B, 0xAC, 0xED, 0x0A, 0x2B}}

type IVoiceCommandManagerInterface interface {
	win32.IInspectableInterface
	InstallCommandSetsFromStorageFileAsync(file *IStorageFile) *IAsyncAction
	Get_InstalledCommandSets() *IMapView[string, *IVoiceCommandSet]
}

type IVoiceCommandManagerVtbl struct {
	win32.IInspectableVtbl
	InstallCommandSetsFromStorageFileAsync uintptr
	Get_InstalledCommandSets               uintptr
}

type IVoiceCommandManager struct {
	win32.IInspectable
}

func (this *IVoiceCommandManager) Vtbl() *IVoiceCommandManagerVtbl {
	return (*IVoiceCommandManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVoiceCommandManager) InstallCommandSetsFromStorageFileAsync(file *IStorageFile) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().InstallCommandSetsFromStorageFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IVoiceCommandManager) Get_InstalledCommandSets() *IMapView[string, *IVoiceCommandSet] {
	var _result *IMapView[string, *IVoiceCommandSet]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstalledCommandSets, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0BEDDA75-46E6-4B11-A088-5C68632899B5
var IID_IVoiceCommandSet = syscall.GUID{0x0BEDDA75, 0x46E6, 0x4B11,
	[8]byte{0xA0, 0x88, 0x5C, 0x68, 0x63, 0x28, 0x99, 0xB5}}

type IVoiceCommandSetInterface interface {
	win32.IInspectableInterface
	Get_Language() string
	Get_Name() string
	SetPhraseListAsync(phraseListName string, phraseList *IIterable[string]) *IAsyncAction
}

type IVoiceCommandSetVtbl struct {
	win32.IInspectableVtbl
	Get_Language       uintptr
	Get_Name           uintptr
	SetPhraseListAsync uintptr
}

type IVoiceCommandSet struct {
	win32.IInspectable
}

func (this *IVoiceCommandSet) Vtbl() *IVoiceCommandSetVtbl {
	return (*IVoiceCommandSetVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVoiceCommandSet) Get_Language() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Language, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVoiceCommandSet) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IVoiceCommandSet) SetPhraseListAsync(phraseListName string, phraseList *IIterable[string]) *IAsyncAction {
	var _result *IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPhraseListAsync, uintptr(unsafe.Pointer(this)), NewHStr(phraseListName).Ptr, uintptr(unsafe.Pointer(phraseList)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type SpeechContinuousRecognitionCompletedEventArgs struct {
	RtClass
	*ISpeechContinuousRecognitionCompletedEventArgs
}

type SpeechContinuousRecognitionResultGeneratedEventArgs struct {
	RtClass
	*ISpeechContinuousRecognitionResultGeneratedEventArgs
}

type SpeechContinuousRecognitionSession struct {
	RtClass
	*ISpeechContinuousRecognitionSession
}

type SpeechRecognitionCompilationResult struct {
	RtClass
	*ISpeechRecognitionCompilationResult
}

type SpeechRecognitionGrammarFileConstraint struct {
	RtClass
	*ISpeechRecognitionGrammarFileConstraint
}

func NewSpeechRecognitionGrammarFileConstraint_Create(file *IStorageFile) *SpeechRecognitionGrammarFileConstraint {
	hs := NewHStr("Windows.Media.SpeechRecognition.SpeechRecognitionGrammarFileConstraint")
	var pFac *ISpeechRecognitionGrammarFileConstraintFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpeechRecognitionGrammarFileConstraintFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISpeechRecognitionGrammarFileConstraint
	p = pFac.Create(file)
	result := &SpeechRecognitionGrammarFileConstraint{
		RtClass:                                 RtClass{PInspect: &p.IInspectable},
		ISpeechRecognitionGrammarFileConstraint: p,
	}
	com.AddToScope(result)
	return result
}

func NewSpeechRecognitionGrammarFileConstraint_CreateWithTag(file *IStorageFile, tag string) *SpeechRecognitionGrammarFileConstraint {
	hs := NewHStr("Windows.Media.SpeechRecognition.SpeechRecognitionGrammarFileConstraint")
	var pFac *ISpeechRecognitionGrammarFileConstraintFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpeechRecognitionGrammarFileConstraintFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISpeechRecognitionGrammarFileConstraint
	p = pFac.CreateWithTag(file, tag)
	result := &SpeechRecognitionGrammarFileConstraint{
		RtClass:                                 RtClass{PInspect: &p.IInspectable},
		ISpeechRecognitionGrammarFileConstraint: p,
	}
	com.AddToScope(result)
	return result
}

type SpeechRecognitionHypothesis struct {
	RtClass
	*ISpeechRecognitionHypothesis
}

type SpeechRecognitionHypothesisGeneratedEventArgs struct {
	RtClass
	*ISpeechRecognitionHypothesisGeneratedEventArgs
}

type SpeechRecognitionListConstraint struct {
	RtClass
	*ISpeechRecognitionListConstraint
}

func NewSpeechRecognitionListConstraint_Create(commands *IIterable[string]) *SpeechRecognitionListConstraint {
	hs := NewHStr("Windows.Media.SpeechRecognition.SpeechRecognitionListConstraint")
	var pFac *ISpeechRecognitionListConstraintFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpeechRecognitionListConstraintFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISpeechRecognitionListConstraint
	p = pFac.Create(commands)
	result := &SpeechRecognitionListConstraint{
		RtClass:                          RtClass{PInspect: &p.IInspectable},
		ISpeechRecognitionListConstraint: p,
	}
	com.AddToScope(result)
	return result
}

func NewSpeechRecognitionListConstraint_CreateWithTag(commands *IIterable[string], tag string) *SpeechRecognitionListConstraint {
	hs := NewHStr("Windows.Media.SpeechRecognition.SpeechRecognitionListConstraint")
	var pFac *ISpeechRecognitionListConstraintFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpeechRecognitionListConstraintFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISpeechRecognitionListConstraint
	p = pFac.CreateWithTag(commands, tag)
	result := &SpeechRecognitionListConstraint{
		RtClass:                          RtClass{PInspect: &p.IInspectable},
		ISpeechRecognitionListConstraint: p,
	}
	com.AddToScope(result)
	return result
}

type SpeechRecognitionQualityDegradingEventArgs struct {
	RtClass
	*ISpeechRecognitionQualityDegradingEventArgs
}

type SpeechRecognitionResult struct {
	RtClass
	*ISpeechRecognitionResult
}

type SpeechRecognitionSemanticInterpretation struct {
	RtClass
	*ISpeechRecognitionSemanticInterpretation
}

type SpeechRecognitionTopicConstraint struct {
	RtClass
	*ISpeechRecognitionTopicConstraint
}

func NewSpeechRecognitionTopicConstraint_Create(scenario SpeechRecognitionScenario, topicHint string) *SpeechRecognitionTopicConstraint {
	hs := NewHStr("Windows.Media.SpeechRecognition.SpeechRecognitionTopicConstraint")
	var pFac *ISpeechRecognitionTopicConstraintFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpeechRecognitionTopicConstraintFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISpeechRecognitionTopicConstraint
	p = pFac.Create(scenario, topicHint)
	result := &SpeechRecognitionTopicConstraint{
		RtClass:                           RtClass{PInspect: &p.IInspectable},
		ISpeechRecognitionTopicConstraint: p,
	}
	com.AddToScope(result)
	return result
}

func NewSpeechRecognitionTopicConstraint_CreateWithTag(scenario SpeechRecognitionScenario, topicHint string, tag string) *SpeechRecognitionTopicConstraint {
	hs := NewHStr("Windows.Media.SpeechRecognition.SpeechRecognitionTopicConstraint")
	var pFac *ISpeechRecognitionTopicConstraintFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpeechRecognitionTopicConstraintFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISpeechRecognitionTopicConstraint
	p = pFac.CreateWithTag(scenario, topicHint, tag)
	result := &SpeechRecognitionTopicConstraint{
		RtClass:                           RtClass{PInspect: &p.IInspectable},
		ISpeechRecognitionTopicConstraint: p,
	}
	com.AddToScope(result)
	return result
}

type SpeechRecognizer struct {
	RtClass
	*ISpeechRecognizer
}

func NewSpeechRecognizer() *SpeechRecognizer {
	hs := NewHStr("Windows.Media.SpeechRecognition.SpeechRecognizer")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &SpeechRecognizer{
		RtClass:           RtClass{PInspect: p},
		ISpeechRecognizer: (*ISpeechRecognizer)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

func NewSpeechRecognizer_Create(language *ILanguage) *SpeechRecognizer {
	hs := NewHStr("Windows.Media.SpeechRecognition.SpeechRecognizer")
	var pFac *ISpeechRecognizerFactory
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpeechRecognizerFactory, unsafe.Pointer(&pFac))
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	var p *ISpeechRecognizer
	p = pFac.Create(language)
	result := &SpeechRecognizer{
		RtClass:           RtClass{PInspect: &p.IInspectable},
		ISpeechRecognizer: p,
	}
	com.AddToScope(result)
	return result
}

func NewISpeechRecognizerStatics() *ISpeechRecognizerStatics {
	var p *ISpeechRecognizerStatics
	hs := NewHStr("Windows.Media.SpeechRecognition.SpeechRecognizer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpeechRecognizerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewISpeechRecognizerStatics2() *ISpeechRecognizerStatics2 {
	var p *ISpeechRecognizerStatics2
	hs := NewHStr("Windows.Media.SpeechRecognition.SpeechRecognizer")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_ISpeechRecognizerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type SpeechRecognizerStateChangedEventArgs struct {
	RtClass
	*ISpeechRecognizerStateChangedEventArgs
}

type SpeechRecognizerTimeouts struct {
	RtClass
	*ISpeechRecognizerTimeouts
}

type SpeechRecognizerUIOptions struct {
	RtClass
	*ISpeechRecognizerUIOptions
}
