package winrt

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

// enums

// enum
// flags
type ArcadeStickButtons uint32

const (
	ArcadeStickButtons_None       ArcadeStickButtons = 0
	ArcadeStickButtons_StickUp    ArcadeStickButtons = 1
	ArcadeStickButtons_StickDown  ArcadeStickButtons = 2
	ArcadeStickButtons_StickLeft  ArcadeStickButtons = 4
	ArcadeStickButtons_StickRight ArcadeStickButtons = 8
	ArcadeStickButtons_Action1    ArcadeStickButtons = 16
	ArcadeStickButtons_Action2    ArcadeStickButtons = 32
	ArcadeStickButtons_Action3    ArcadeStickButtons = 64
	ArcadeStickButtons_Action4    ArcadeStickButtons = 128
	ArcadeStickButtons_Action5    ArcadeStickButtons = 256
	ArcadeStickButtons_Action6    ArcadeStickButtons = 512
	ArcadeStickButtons_Special1   ArcadeStickButtons = 1024
	ArcadeStickButtons_Special2   ArcadeStickButtons = 2048
)

// enum
// flags
type FlightStickButtons uint32

const (
	FlightStickButtons_None          FlightStickButtons = 0
	FlightStickButtons_FirePrimary   FlightStickButtons = 1
	FlightStickButtons_FireSecondary FlightStickButtons = 2
)

// enum
type GameControllerButtonLabel int32

const (
	GameControllerButtonLabel_None                 GameControllerButtonLabel = 0
	GameControllerButtonLabel_XboxBack             GameControllerButtonLabel = 1
	GameControllerButtonLabel_XboxStart            GameControllerButtonLabel = 2
	GameControllerButtonLabel_XboxMenu             GameControllerButtonLabel = 3
	GameControllerButtonLabel_XboxView             GameControllerButtonLabel = 4
	GameControllerButtonLabel_XboxUp               GameControllerButtonLabel = 5
	GameControllerButtonLabel_XboxDown             GameControllerButtonLabel = 6
	GameControllerButtonLabel_XboxLeft             GameControllerButtonLabel = 7
	GameControllerButtonLabel_XboxRight            GameControllerButtonLabel = 8
	GameControllerButtonLabel_XboxA                GameControllerButtonLabel = 9
	GameControllerButtonLabel_XboxB                GameControllerButtonLabel = 10
	GameControllerButtonLabel_XboxX                GameControllerButtonLabel = 11
	GameControllerButtonLabel_XboxY                GameControllerButtonLabel = 12
	GameControllerButtonLabel_XboxLeftBumper       GameControllerButtonLabel = 13
	GameControllerButtonLabel_XboxLeftTrigger      GameControllerButtonLabel = 14
	GameControllerButtonLabel_XboxLeftStickButton  GameControllerButtonLabel = 15
	GameControllerButtonLabel_XboxRightBumper      GameControllerButtonLabel = 16
	GameControllerButtonLabel_XboxRightTrigger     GameControllerButtonLabel = 17
	GameControllerButtonLabel_XboxRightStickButton GameControllerButtonLabel = 18
	GameControllerButtonLabel_XboxPaddle1          GameControllerButtonLabel = 19
	GameControllerButtonLabel_XboxPaddle2          GameControllerButtonLabel = 20
	GameControllerButtonLabel_XboxPaddle3          GameControllerButtonLabel = 21
	GameControllerButtonLabel_XboxPaddle4          GameControllerButtonLabel = 22
	GameControllerButtonLabel_Mode                 GameControllerButtonLabel = 23
	GameControllerButtonLabel_Select               GameControllerButtonLabel = 24
	GameControllerButtonLabel_Menu                 GameControllerButtonLabel = 25
	GameControllerButtonLabel_View                 GameControllerButtonLabel = 26
	GameControllerButtonLabel_Back                 GameControllerButtonLabel = 27
	GameControllerButtonLabel_Start                GameControllerButtonLabel = 28
	GameControllerButtonLabel_Options              GameControllerButtonLabel = 29
	GameControllerButtonLabel_Share                GameControllerButtonLabel = 30
	GameControllerButtonLabel_Up                   GameControllerButtonLabel = 31
	GameControllerButtonLabel_Down                 GameControllerButtonLabel = 32
	GameControllerButtonLabel_Left                 GameControllerButtonLabel = 33
	GameControllerButtonLabel_Right                GameControllerButtonLabel = 34
	GameControllerButtonLabel_LetterA              GameControllerButtonLabel = 35
	GameControllerButtonLabel_LetterB              GameControllerButtonLabel = 36
	GameControllerButtonLabel_LetterC              GameControllerButtonLabel = 37
	GameControllerButtonLabel_LetterL              GameControllerButtonLabel = 38
	GameControllerButtonLabel_LetterR              GameControllerButtonLabel = 39
	GameControllerButtonLabel_LetterX              GameControllerButtonLabel = 40
	GameControllerButtonLabel_LetterY              GameControllerButtonLabel = 41
	GameControllerButtonLabel_LetterZ              GameControllerButtonLabel = 42
	GameControllerButtonLabel_Cross                GameControllerButtonLabel = 43
	GameControllerButtonLabel_Circle               GameControllerButtonLabel = 44
	GameControllerButtonLabel_Square               GameControllerButtonLabel = 45
	GameControllerButtonLabel_Triangle             GameControllerButtonLabel = 46
	GameControllerButtonLabel_LeftBumper           GameControllerButtonLabel = 47
	GameControllerButtonLabel_LeftTrigger          GameControllerButtonLabel = 48
	GameControllerButtonLabel_LeftStickButton      GameControllerButtonLabel = 49
	GameControllerButtonLabel_Left1                GameControllerButtonLabel = 50
	GameControllerButtonLabel_Left2                GameControllerButtonLabel = 51
	GameControllerButtonLabel_Left3                GameControllerButtonLabel = 52
	GameControllerButtonLabel_RightBumper          GameControllerButtonLabel = 53
	GameControllerButtonLabel_RightTrigger         GameControllerButtonLabel = 54
	GameControllerButtonLabel_RightStickButton     GameControllerButtonLabel = 55
	GameControllerButtonLabel_Right1               GameControllerButtonLabel = 56
	GameControllerButtonLabel_Right2               GameControllerButtonLabel = 57
	GameControllerButtonLabel_Right3               GameControllerButtonLabel = 58
	GameControllerButtonLabel_Paddle1              GameControllerButtonLabel = 59
	GameControllerButtonLabel_Paddle2              GameControllerButtonLabel = 60
	GameControllerButtonLabel_Paddle3              GameControllerButtonLabel = 61
	GameControllerButtonLabel_Paddle4              GameControllerButtonLabel = 62
	GameControllerButtonLabel_Plus                 GameControllerButtonLabel = 63
	GameControllerButtonLabel_Minus                GameControllerButtonLabel = 64
	GameControllerButtonLabel_DownLeftArrow        GameControllerButtonLabel = 65
	GameControllerButtonLabel_DialLeft             GameControllerButtonLabel = 66
	GameControllerButtonLabel_DialRight            GameControllerButtonLabel = 67
	GameControllerButtonLabel_Suspension           GameControllerButtonLabel = 68
)

// enum
type GameControllerSwitchKind int32

const (
	GameControllerSwitchKind_TwoWay   GameControllerSwitchKind = 0
	GameControllerSwitchKind_FourWay  GameControllerSwitchKind = 1
	GameControllerSwitchKind_EightWay GameControllerSwitchKind = 2
)

// enum
type GameControllerSwitchPosition int32

const (
	GameControllerSwitchPosition_Center    GameControllerSwitchPosition = 0
	GameControllerSwitchPosition_Up        GameControllerSwitchPosition = 1
	GameControllerSwitchPosition_UpRight   GameControllerSwitchPosition = 2
	GameControllerSwitchPosition_Right     GameControllerSwitchPosition = 3
	GameControllerSwitchPosition_DownRight GameControllerSwitchPosition = 4
	GameControllerSwitchPosition_Down      GameControllerSwitchPosition = 5
	GameControllerSwitchPosition_DownLeft  GameControllerSwitchPosition = 6
	GameControllerSwitchPosition_Left      GameControllerSwitchPosition = 7
	GameControllerSwitchPosition_UpLeft    GameControllerSwitchPosition = 8
)

// enum
// flags
type GamepadButtons uint32

const (
	GamepadButtons_None            GamepadButtons = 0
	GamepadButtons_Menu            GamepadButtons = 1
	GamepadButtons_View            GamepadButtons = 2
	GamepadButtons_A               GamepadButtons = 4
	GamepadButtons_B               GamepadButtons = 8
	GamepadButtons_X               GamepadButtons = 16
	GamepadButtons_Y               GamepadButtons = 32
	GamepadButtons_DPadUp          GamepadButtons = 64
	GamepadButtons_DPadDown        GamepadButtons = 128
	GamepadButtons_DPadLeft        GamepadButtons = 256
	GamepadButtons_DPadRight       GamepadButtons = 512
	GamepadButtons_LeftShoulder    GamepadButtons = 1024
	GamepadButtons_RightShoulder   GamepadButtons = 2048
	GamepadButtons_LeftThumbstick  GamepadButtons = 4096
	GamepadButtons_RightThumbstick GamepadButtons = 8192
	GamepadButtons_Paddle1         GamepadButtons = 16384
	GamepadButtons_Paddle2         GamepadButtons = 32768
	GamepadButtons_Paddle3         GamepadButtons = 65536
	GamepadButtons_Paddle4         GamepadButtons = 131072
)

// enum
// flags
type OptionalUINavigationButtons uint32

const (
	OptionalUINavigationButtons_None        OptionalUINavigationButtons = 0
	OptionalUINavigationButtons_Context1    OptionalUINavigationButtons = 1
	OptionalUINavigationButtons_Context2    OptionalUINavigationButtons = 2
	OptionalUINavigationButtons_Context3    OptionalUINavigationButtons = 4
	OptionalUINavigationButtons_Context4    OptionalUINavigationButtons = 8
	OptionalUINavigationButtons_PageUp      OptionalUINavigationButtons = 16
	OptionalUINavigationButtons_PageDown    OptionalUINavigationButtons = 32
	OptionalUINavigationButtons_PageLeft    OptionalUINavigationButtons = 64
	OptionalUINavigationButtons_PageRight   OptionalUINavigationButtons = 128
	OptionalUINavigationButtons_ScrollUp    OptionalUINavigationButtons = 256
	OptionalUINavigationButtons_ScrollDown  OptionalUINavigationButtons = 512
	OptionalUINavigationButtons_ScrollLeft  OptionalUINavigationButtons = 1024
	OptionalUINavigationButtons_ScrollRight OptionalUINavigationButtons = 2048
)

// enum
// flags
type RacingWheelButtons uint32

const (
	RacingWheelButtons_None         RacingWheelButtons = 0
	RacingWheelButtons_PreviousGear RacingWheelButtons = 1
	RacingWheelButtons_NextGear     RacingWheelButtons = 2
	RacingWheelButtons_DPadUp       RacingWheelButtons = 4
	RacingWheelButtons_DPadDown     RacingWheelButtons = 8
	RacingWheelButtons_DPadLeft     RacingWheelButtons = 16
	RacingWheelButtons_DPadRight    RacingWheelButtons = 32
	RacingWheelButtons_Button1      RacingWheelButtons = 64
	RacingWheelButtons_Button2      RacingWheelButtons = 128
	RacingWheelButtons_Button3      RacingWheelButtons = 256
	RacingWheelButtons_Button4      RacingWheelButtons = 512
	RacingWheelButtons_Button5      RacingWheelButtons = 1024
	RacingWheelButtons_Button6      RacingWheelButtons = 2048
	RacingWheelButtons_Button7      RacingWheelButtons = 4096
	RacingWheelButtons_Button8      RacingWheelButtons = 8192
	RacingWheelButtons_Button9      RacingWheelButtons = 16384
	RacingWheelButtons_Button10     RacingWheelButtons = 32768
	RacingWheelButtons_Button11     RacingWheelButtons = 65536
	RacingWheelButtons_Button12     RacingWheelButtons = 131072
	RacingWheelButtons_Button13     RacingWheelButtons = 262144
	RacingWheelButtons_Button14     RacingWheelButtons = 524288
	RacingWheelButtons_Button15     RacingWheelButtons = 1048576
	RacingWheelButtons_Button16     RacingWheelButtons = 2097152
)

// enum
// flags
type RequiredUINavigationButtons uint32

const (
	RequiredUINavigationButtons_None   RequiredUINavigationButtons = 0
	RequiredUINavigationButtons_Menu   RequiredUINavigationButtons = 1
	RequiredUINavigationButtons_View   RequiredUINavigationButtons = 2
	RequiredUINavigationButtons_Accept RequiredUINavigationButtons = 4
	RequiredUINavigationButtons_Cancel RequiredUINavigationButtons = 8
	RequiredUINavigationButtons_Up     RequiredUINavigationButtons = 16
	RequiredUINavigationButtons_Down   RequiredUINavigationButtons = 32
	RequiredUINavigationButtons_Left   RequiredUINavigationButtons = 64
	RequiredUINavigationButtons_Right  RequiredUINavigationButtons = 128
)

// structs

type ArcadeStickReading struct {
	Timestamp uint64
	Buttons   ArcadeStickButtons
}

type FlightStickReading struct {
	Timestamp uint64
	Buttons   FlightStickButtons
	HatSwitch GameControllerSwitchPosition
	Roll      float64
	Pitch     float64
	Yaw       float64
	Throttle  float64
}

type GamepadReading struct {
	Timestamp        uint64
	Buttons          GamepadButtons
	LeftTrigger      float64
	RightTrigger     float64
	LeftThumbstickX  float64
	LeftThumbstickY  float64
	RightThumbstickX float64
	RightThumbstickY float64
}

type GamepadVibration struct {
	LeftMotor    float64
	RightMotor   float64
	LeftTrigger  float64
	RightTrigger float64
}

type GamingInputPreviewContract struct {
}

type RacingWheelReading struct {
	Timestamp          uint64
	Buttons            RacingWheelButtons
	PatternShifterGear int32
	Wheel              float64
	Throttle           float64
	Brake              float64
	Clutch             float64
	Handbrake          float64
}

type UINavigationReading struct {
	Timestamp       uint64
	RequiredButtons RequiredUINavigationButtons
	OptionalButtons OptionalUINavigationButtons
}

// interfaces

// B14A539D-BEFB-4C81-8051-15ECF3B13036
var IID_IArcadeStick = syscall.GUID{0xB14A539D, 0xBEFB, 0x4C81,
	[8]byte{0x80, 0x51, 0x15, 0xEC, 0xF3, 0xB1, 0x30, 0x36}}

type IArcadeStickInterface interface {
	win32.IInspectableInterface
	GetButtonLabel(button ArcadeStickButtons) GameControllerButtonLabel
	GetCurrentReading() ArcadeStickReading
}

type IArcadeStickVtbl struct {
	win32.IInspectableVtbl
	GetButtonLabel    uintptr
	GetCurrentReading uintptr
}

type IArcadeStick struct {
	win32.IInspectable
}

func (this *IArcadeStick) Vtbl() *IArcadeStickVtbl {
	return (*IArcadeStickVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IArcadeStick) GetButtonLabel(button ArcadeStickButtons) GameControllerButtonLabel {
	var _result GameControllerButtonLabel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetButtonLabel, uintptr(unsafe.Pointer(this)), uintptr(button), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IArcadeStick) GetCurrentReading() ArcadeStickReading {
	var _result ArcadeStickReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5C37B8C8-37B1-4AD8-9458-200F1A30018E
var IID_IArcadeStickStatics = syscall.GUID{0x5C37B8C8, 0x37B1, 0x4AD8,
	[8]byte{0x94, 0x58, 0x20, 0x0F, 0x1A, 0x30, 0x01, 0x8E}}

type IArcadeStickStaticsInterface interface {
	win32.IInspectableInterface
	Add_ArcadeStickAdded(value EventHandler[*IArcadeStick]) EventRegistrationToken
	Remove_ArcadeStickAdded(token EventRegistrationToken)
	Add_ArcadeStickRemoved(value EventHandler[*IArcadeStick]) EventRegistrationToken
	Remove_ArcadeStickRemoved(token EventRegistrationToken)
	Get_ArcadeSticks() *IVectorView[*IArcadeStick]
}

type IArcadeStickStaticsVtbl struct {
	win32.IInspectableVtbl
	Add_ArcadeStickAdded      uintptr
	Remove_ArcadeStickAdded   uintptr
	Add_ArcadeStickRemoved    uintptr
	Remove_ArcadeStickRemoved uintptr
	Get_ArcadeSticks          uintptr
}

type IArcadeStickStatics struct {
	win32.IInspectable
}

func (this *IArcadeStickStatics) Vtbl() *IArcadeStickStaticsVtbl {
	return (*IArcadeStickStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IArcadeStickStatics) Add_ArcadeStickAdded(value EventHandler[*IArcadeStick]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ArcadeStickAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IArcadeStickStatics) Remove_ArcadeStickAdded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ArcadeStickAdded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IArcadeStickStatics) Add_ArcadeStickRemoved(value EventHandler[*IArcadeStick]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_ArcadeStickRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IArcadeStickStatics) Remove_ArcadeStickRemoved(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_ArcadeStickRemoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IArcadeStickStatics) Get_ArcadeSticks() *IVectorView[*IArcadeStick] {
	var _result *IVectorView[*IArcadeStick]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ArcadeSticks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 52B5D744-BB86-445A-B59C-596F0E2A49DF
var IID_IArcadeStickStatics2 = syscall.GUID{0x52B5D744, 0xBB86, 0x445A,
	[8]byte{0xB5, 0x9C, 0x59, 0x6F, 0x0E, 0x2A, 0x49, 0xDF}}

type IArcadeStickStatics2Interface interface {
	win32.IInspectableInterface
	FromGameController(gameController *IGameController) *IArcadeStick
}

type IArcadeStickStatics2Vtbl struct {
	win32.IInspectableVtbl
	FromGameController uintptr
}

type IArcadeStickStatics2 struct {
	win32.IInspectable
}

func (this *IArcadeStickStatics2) Vtbl() *IArcadeStickStatics2Vtbl {
	return (*IArcadeStickStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IArcadeStickStatics2) FromGameController(gameController *IGameController) *IArcadeStick {
	var _result *IArcadeStick
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromGameController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(gameController)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B4A2C01C-B83B-4459-A1A9-97B03C33DA7C
var IID_IFlightStick = syscall.GUID{0xB4A2C01C, 0xB83B, 0x4459,
	[8]byte{0xA1, 0xA9, 0x97, 0xB0, 0x3C, 0x33, 0xDA, 0x7C}}

type IFlightStickInterface interface {
	win32.IInspectableInterface
	Get_HatSwitchKind() GameControllerSwitchKind
	GetButtonLabel(button FlightStickButtons) GameControllerButtonLabel
	GetCurrentReading() FlightStickReading
}

type IFlightStickVtbl struct {
	win32.IInspectableVtbl
	Get_HatSwitchKind uintptr
	GetButtonLabel    uintptr
	GetCurrentReading uintptr
}

type IFlightStick struct {
	win32.IInspectable
}

func (this *IFlightStick) Vtbl() *IFlightStickVtbl {
	return (*IFlightStickVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFlightStick) Get_HatSwitchKind() GameControllerSwitchKind {
	var _result GameControllerSwitchKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HatSwitchKind, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFlightStick) GetButtonLabel(button FlightStickButtons) GameControllerButtonLabel {
	var _result GameControllerButtonLabel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetButtonLabel, uintptr(unsafe.Pointer(this)), uintptr(button), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFlightStick) GetCurrentReading() FlightStickReading {
	var _result FlightStickReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 5514924A-FECC-435E-83DC-5CEC8A18A520
var IID_IFlightStickStatics = syscall.GUID{0x5514924A, 0xFECC, 0x435E,
	[8]byte{0x83, 0xDC, 0x5C, 0xEC, 0x8A, 0x18, 0xA5, 0x20}}

type IFlightStickStaticsInterface interface {
	win32.IInspectableInterface
	Add_FlightStickAdded(value EventHandler[*IFlightStick]) EventRegistrationToken
	Remove_FlightStickAdded(token EventRegistrationToken)
	Add_FlightStickRemoved(value EventHandler[*IFlightStick]) EventRegistrationToken
	Remove_FlightStickRemoved(token EventRegistrationToken)
	Get_FlightSticks() *IVectorView[*IFlightStick]
	FromGameController(gameController *IGameController) *IFlightStick
}

type IFlightStickStaticsVtbl struct {
	win32.IInspectableVtbl
	Add_FlightStickAdded      uintptr
	Remove_FlightStickAdded   uintptr
	Add_FlightStickRemoved    uintptr
	Remove_FlightStickRemoved uintptr
	Get_FlightSticks          uintptr
	FromGameController        uintptr
}

type IFlightStickStatics struct {
	win32.IInspectable
}

func (this *IFlightStickStatics) Vtbl() *IFlightStickStaticsVtbl {
	return (*IFlightStickStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFlightStickStatics) Add_FlightStickAdded(value EventHandler[*IFlightStick]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FlightStickAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFlightStickStatics) Remove_FlightStickAdded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FlightStickAdded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IFlightStickStatics) Add_FlightStickRemoved(value EventHandler[*IFlightStick]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_FlightStickRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IFlightStickStatics) Remove_FlightStickRemoved(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_FlightStickRemoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IFlightStickStatics) Get_FlightSticks() *IVectorView[*IFlightStick] {
	var _result *IVectorView[*IFlightStick]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_FlightSticks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IFlightStickStatics) FromGameController(gameController *IGameController) *IFlightStick {
	var _result *IFlightStick
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromGameController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(gameController)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1BAF6522-5F64-42C5-8267-B9FE2215BFBD
var IID_IGameController = syscall.GUID{0x1BAF6522, 0x5F64, 0x42C5,
	[8]byte{0x82, 0x67, 0xB9, 0xFE, 0x22, 0x15, 0xBF, 0xBD}}

type IGameControllerInterface interface {
	win32.IInspectableInterface
	Add_HeadsetConnected(value TypedEventHandler[*IGameController, *IHeadset]) EventRegistrationToken
	Remove_HeadsetConnected(token EventRegistrationToken)
	Add_HeadsetDisconnected(value TypedEventHandler[*IGameController, *IHeadset]) EventRegistrationToken
	Remove_HeadsetDisconnected(token EventRegistrationToken)
	Add_UserChanged(value TypedEventHandler[*IGameController, *IUserChangedEventArgs]) EventRegistrationToken
	Remove_UserChanged(token EventRegistrationToken)
	Get_Headset() *IHeadset
	Get_IsWireless() bool
	Get_User() *IUser
}

type IGameControllerVtbl struct {
	win32.IInspectableVtbl
	Add_HeadsetConnected       uintptr
	Remove_HeadsetConnected    uintptr
	Add_HeadsetDisconnected    uintptr
	Remove_HeadsetDisconnected uintptr
	Add_UserChanged            uintptr
	Remove_UserChanged         uintptr
	Get_Headset                uintptr
	Get_IsWireless             uintptr
	Get_User                   uintptr
}

type IGameController struct {
	win32.IInspectable
}

func (this *IGameController) Vtbl() *IGameControllerVtbl {
	return (*IGameControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameController) Add_HeadsetConnected(value TypedEventHandler[*IGameController, *IHeadset]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_HeadsetConnected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameController) Remove_HeadsetConnected(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_HeadsetConnected, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGameController) Add_HeadsetDisconnected(value TypedEventHandler[*IGameController, *IHeadset]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_HeadsetDisconnected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameController) Remove_HeadsetDisconnected(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_HeadsetDisconnected, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGameController) Add_UserChanged(value TypedEventHandler[*IGameController, *IUserChangedEventArgs]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_UserChanged, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameController) Remove_UserChanged(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_UserChanged, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGameController) Get_Headset() *IHeadset {
	var _result *IHeadset
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Headset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IGameController) Get_IsWireless() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsWireless, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGameController) Get_User() *IUser {
	var _result *IUser
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_User, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DCECC681-3963-4DA6-955D-553F3B6F6161
var IID_IGameControllerBatteryInfo = syscall.GUID{0xDCECC681, 0x3963, 0x4DA6,
	[8]byte{0x95, 0x5D, 0x55, 0x3F, 0x3B, 0x6F, 0x61, 0x61}}

type IGameControllerBatteryInfoInterface interface {
	win32.IInspectableInterface
	TryGetBatteryReport() *IBatteryReport
}

type IGameControllerBatteryInfoVtbl struct {
	win32.IInspectableVtbl
	TryGetBatteryReport uintptr
}

type IGameControllerBatteryInfo struct {
	win32.IInspectable
}

func (this *IGameControllerBatteryInfo) Vtbl() *IGameControllerBatteryInfoVtbl {
	return (*IGameControllerBatteryInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGameControllerBatteryInfo) TryGetBatteryReport() *IBatteryReport {
	var _result *IBatteryReport
	_hr, _, _ := syscall.SyscallN(this.Vtbl().TryGetBatteryReport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BC7BB43C-0A69-3903-9E9D-A50F86A45DE5
var IID_IGamepad = syscall.GUID{0xBC7BB43C, 0x0A69, 0x3903,
	[8]byte{0x9E, 0x9D, 0xA5, 0x0F, 0x86, 0xA4, 0x5D, 0xE5}}

type IGamepadInterface interface {
	win32.IInspectableInterface
	Get_Vibration() GamepadVibration
	Put_Vibration(value GamepadVibration)
	GetCurrentReading() GamepadReading
}

type IGamepadVtbl struct {
	win32.IInspectableVtbl
	Get_Vibration     uintptr
	Put_Vibration     uintptr
	GetCurrentReading uintptr
}

type IGamepad struct {
	win32.IInspectable
}

func (this *IGamepad) Vtbl() *IGamepadVtbl {
	return (*IGamepadVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGamepad) Get_Vibration() GamepadVibration {
	var _result GamepadVibration
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Vibration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGamepad) Put_Vibration(value GamepadVibration) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Vibration, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IGamepad) GetCurrentReading() GamepadReading {
	var _result GamepadReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3C1689BD-5915-4245-B0C0-C89FAE0308FF
var IID_IGamepad2 = syscall.GUID{0x3C1689BD, 0x5915, 0x4245,
	[8]byte{0xB0, 0xC0, 0xC8, 0x9F, 0xAE, 0x03, 0x08, 0xFF}}

type IGamepad2Interface interface {
	win32.IInspectableInterface
	GetButtonLabel(button GamepadButtons) GameControllerButtonLabel
}

type IGamepad2Vtbl struct {
	win32.IInspectableVtbl
	GetButtonLabel uintptr
}

type IGamepad2 struct {
	win32.IInspectable
}

func (this *IGamepad2) Vtbl() *IGamepad2Vtbl {
	return (*IGamepad2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGamepad2) GetButtonLabel(button GamepadButtons) GameControllerButtonLabel {
	var _result GameControllerButtonLabel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetButtonLabel, uintptr(unsafe.Pointer(this)), uintptr(button), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 8BBCE529-D49C-39E9-9560-E47DDE96B7C8
var IID_IGamepadStatics = syscall.GUID{0x8BBCE529, 0xD49C, 0x39E9,
	[8]byte{0x95, 0x60, 0xE4, 0x7D, 0xDE, 0x96, 0xB7, 0xC8}}

type IGamepadStaticsInterface interface {
	win32.IInspectableInterface
	Add_GamepadAdded(value EventHandler[*IGamepad]) EventRegistrationToken
	Remove_GamepadAdded(token EventRegistrationToken)
	Add_GamepadRemoved(value EventHandler[*IGamepad]) EventRegistrationToken
	Remove_GamepadRemoved(token EventRegistrationToken)
	Get_Gamepads() *IVectorView[*IGamepad]
}

type IGamepadStaticsVtbl struct {
	win32.IInspectableVtbl
	Add_GamepadAdded      uintptr
	Remove_GamepadAdded   uintptr
	Add_GamepadRemoved    uintptr
	Remove_GamepadRemoved uintptr
	Get_Gamepads          uintptr
}

type IGamepadStatics struct {
	win32.IInspectable
}

func (this *IGamepadStatics) Vtbl() *IGamepadStaticsVtbl {
	return (*IGamepadStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGamepadStatics) Add_GamepadAdded(value EventHandler[*IGamepad]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_GamepadAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGamepadStatics) Remove_GamepadAdded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_GamepadAdded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGamepadStatics) Add_GamepadRemoved(value EventHandler[*IGamepad]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_GamepadRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IGamepadStatics) Remove_GamepadRemoved(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_GamepadRemoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IGamepadStatics) Get_Gamepads() *IVectorView[*IGamepad] {
	var _result *IVectorView[*IGamepad]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Gamepads, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 42676DC5-0856-47C4-9213-B395504C3A3C
var IID_IGamepadStatics2 = syscall.GUID{0x42676DC5, 0x0856, 0x47C4,
	[8]byte{0x92, 0x13, 0xB3, 0x95, 0x50, 0x4C, 0x3A, 0x3C}}

type IGamepadStatics2Interface interface {
	win32.IInspectableInterface
	FromGameController(gameController *IGameController) *IGamepad
}

type IGamepadStatics2Vtbl struct {
	win32.IInspectableVtbl
	FromGameController uintptr
}

type IGamepadStatics2 struct {
	win32.IInspectable
}

func (this *IGamepadStatics2) Vtbl() *IGamepadStatics2Vtbl {
	return (*IGamepadStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGamepadStatics2) FromGameController(gameController *IGameController) *IGamepad {
	var _result *IGamepad
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromGameController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(gameController)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3FD156EF-6925-3FA8-9181-029C5223AE3B
var IID_IHeadset = syscall.GUID{0x3FD156EF, 0x6925, 0x3FA8,
	[8]byte{0x91, 0x81, 0x02, 0x9C, 0x52, 0x23, 0xAE, 0x3B}}

type IHeadsetInterface interface {
	win32.IInspectableInterface
	Get_CaptureDeviceId() string
	Get_RenderDeviceId() string
}

type IHeadsetVtbl struct {
	win32.IInspectableVtbl
	Get_CaptureDeviceId uintptr
	Get_RenderDeviceId  uintptr
}

type IHeadset struct {
	win32.IInspectable
}

func (this *IHeadset) Vtbl() *IHeadsetVtbl {
	return (*IHeadsetVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IHeadset) Get_CaptureDeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CaptureDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IHeadset) Get_RenderDeviceId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RenderDeviceId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// F546656F-E106-4C82-A90F-554012904B85
var IID_IRacingWheel = syscall.GUID{0xF546656F, 0xE106, 0x4C82,
	[8]byte{0xA9, 0x0F, 0x55, 0x40, 0x12, 0x90, 0x4B, 0x85}}

type IRacingWheelInterface interface {
	win32.IInspectableInterface
	Get_HasClutch() bool
	Get_HasHandbrake() bool
	Get_HasPatternShifter() bool
	Get_MaxPatternShifterGear() int32
	Get_MaxWheelAngle() float64
	Get_WheelMotor() *IForceFeedbackMotor
	GetButtonLabel(button RacingWheelButtons) GameControllerButtonLabel
	GetCurrentReading() RacingWheelReading
}

type IRacingWheelVtbl struct {
	win32.IInspectableVtbl
	Get_HasClutch             uintptr
	Get_HasHandbrake          uintptr
	Get_HasPatternShifter     uintptr
	Get_MaxPatternShifterGear uintptr
	Get_MaxWheelAngle         uintptr
	Get_WheelMotor            uintptr
	GetButtonLabel            uintptr
	GetCurrentReading         uintptr
}

type IRacingWheel struct {
	win32.IInspectable
}

func (this *IRacingWheel) Vtbl() *IRacingWheelVtbl {
	return (*IRacingWheelVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRacingWheel) Get_HasClutch() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasClutch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRacingWheel) Get_HasHandbrake() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasHandbrake, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRacingWheel) Get_HasPatternShifter() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HasPatternShifter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRacingWheel) Get_MaxPatternShifterGear() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxPatternShifterGear, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRacingWheel) Get_MaxWheelAngle() float64 {
	var _result float64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MaxWheelAngle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRacingWheel) Get_WheelMotor() *IForceFeedbackMotor {
	var _result *IForceFeedbackMotor
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_WheelMotor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRacingWheel) GetButtonLabel(button RacingWheelButtons) GameControllerButtonLabel {
	var _result GameControllerButtonLabel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetButtonLabel, uintptr(unsafe.Pointer(this)), uintptr(button), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRacingWheel) GetCurrentReading() RacingWheelReading {
	var _result RacingWheelReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 3AC12CD5-581B-4936-9F94-69F1E6514C7D
var IID_IRacingWheelStatics = syscall.GUID{0x3AC12CD5, 0x581B, 0x4936,
	[8]byte{0x9F, 0x94, 0x69, 0xF1, 0xE6, 0x51, 0x4C, 0x7D}}

type IRacingWheelStaticsInterface interface {
	win32.IInspectableInterface
	Add_RacingWheelAdded(value EventHandler[*IRacingWheel]) EventRegistrationToken
	Remove_RacingWheelAdded(token EventRegistrationToken)
	Add_RacingWheelRemoved(value EventHandler[*IRacingWheel]) EventRegistrationToken
	Remove_RacingWheelRemoved(token EventRegistrationToken)
	Get_RacingWheels() *IVectorView[*IRacingWheel]
}

type IRacingWheelStaticsVtbl struct {
	win32.IInspectableVtbl
	Add_RacingWheelAdded      uintptr
	Remove_RacingWheelAdded   uintptr
	Add_RacingWheelRemoved    uintptr
	Remove_RacingWheelRemoved uintptr
	Get_RacingWheels          uintptr
}

type IRacingWheelStatics struct {
	win32.IInspectable
}

func (this *IRacingWheelStatics) Vtbl() *IRacingWheelStaticsVtbl {
	return (*IRacingWheelStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRacingWheelStatics) Add_RacingWheelAdded(value EventHandler[*IRacingWheel]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RacingWheelAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRacingWheelStatics) Remove_RacingWheelAdded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RacingWheelAdded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IRacingWheelStatics) Add_RacingWheelRemoved(value EventHandler[*IRacingWheel]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RacingWheelRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRacingWheelStatics) Remove_RacingWheelRemoved(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RacingWheelRemoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IRacingWheelStatics) Get_RacingWheels() *IVectorView[*IRacingWheel] {
	var _result *IVectorView[*IRacingWheel]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RacingWheels, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E666BCAA-EDFD-4323-A9F6-3C384048D1ED
var IID_IRacingWheelStatics2 = syscall.GUID{0xE666BCAA, 0xEDFD, 0x4323,
	[8]byte{0xA9, 0xF6, 0x3C, 0x38, 0x40, 0x48, 0xD1, 0xED}}

type IRacingWheelStatics2Interface interface {
	win32.IInspectableInterface
	FromGameController(gameController *IGameController) *IRacingWheel
}

type IRacingWheelStatics2Vtbl struct {
	win32.IInspectableVtbl
	FromGameController uintptr
}

type IRacingWheelStatics2 struct {
	win32.IInspectable
}

func (this *IRacingWheelStatics2) Vtbl() *IRacingWheelStatics2Vtbl {
	return (*IRacingWheelStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRacingWheelStatics2) FromGameController(gameController *IGameController) *IRacingWheel {
	var _result *IRacingWheel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromGameController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(gameController)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 7CAD6D91-A7E1-4F71-9A78-33E9C5DFEA62
var IID_IRawGameController = syscall.GUID{0x7CAD6D91, 0xA7E1, 0x4F71,
	[8]byte{0x9A, 0x78, 0x33, 0xE9, 0xC5, 0xDF, 0xEA, 0x62}}

type IRawGameControllerInterface interface {
	win32.IInspectableInterface
	Get_AxisCount() int32
	Get_ButtonCount() int32
	Get_ForceFeedbackMotors() *IVectorView[*IForceFeedbackMotor]
	Get_HardwareProductId() uint16
	Get_HardwareVendorId() uint16
	Get_SwitchCount() int32
	GetButtonLabel(buttonIndex int32) GameControllerButtonLabel
	GetCurrentReading(buttonArrayLength uint32, buttonArray *bool, switchArrayLength uint32, switchArray *GameControllerSwitchPosition, axisArrayLength uint32, axisArray *float64) uint64
	GetSwitchKind(switchIndex int32) GameControllerSwitchKind
}

type IRawGameControllerVtbl struct {
	win32.IInspectableVtbl
	Get_AxisCount           uintptr
	Get_ButtonCount         uintptr
	Get_ForceFeedbackMotors uintptr
	Get_HardwareProductId   uintptr
	Get_HardwareVendorId    uintptr
	Get_SwitchCount         uintptr
	GetButtonLabel          uintptr
	GetCurrentReading       uintptr
	GetSwitchKind           uintptr
}

type IRawGameController struct {
	win32.IInspectable
}

func (this *IRawGameController) Vtbl() *IRawGameControllerVtbl {
	return (*IRawGameControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRawGameController) Get_AxisCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AxisCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRawGameController) Get_ButtonCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ButtonCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRawGameController) Get_ForceFeedbackMotors() *IVectorView[*IForceFeedbackMotor] {
	var _result *IVectorView[*IForceFeedbackMotor]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForceFeedbackMotors, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRawGameController) Get_HardwareProductId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HardwareProductId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRawGameController) Get_HardwareVendorId() uint16 {
	var _result uint16
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HardwareVendorId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRawGameController) Get_SwitchCount() int32 {
	var _result int32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SwitchCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRawGameController) GetButtonLabel(buttonIndex int32) GameControllerButtonLabel {
	var _result GameControllerButtonLabel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetButtonLabel, uintptr(unsafe.Pointer(this)), uintptr(buttonIndex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRawGameController) GetCurrentReading(buttonArrayLength uint32, buttonArray *bool, switchArrayLength uint32, switchArray *GameControllerSwitchPosition, axisArrayLength uint32, axisArray *float64) uint64 {
	var _result uint64
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(buttonArrayLength), uintptr(unsafe.Pointer(buttonArray)), uintptr(switchArrayLength), uintptr(unsafe.Pointer(switchArray)), uintptr(axisArrayLength), uintptr(unsafe.Pointer(axisArray)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRawGameController) GetSwitchKind(switchIndex int32) GameControllerSwitchKind {
	var _result GameControllerSwitchKind
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetSwitchKind, uintptr(unsafe.Pointer(this)), uintptr(switchIndex), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 43C0C035-BB73-4756-A787-3ED6BEA617BD
var IID_IRawGameController2 = syscall.GUID{0x43C0C035, 0xBB73, 0x4756,
	[8]byte{0xA7, 0x87, 0x3E, 0xD6, 0xBE, 0xA6, 0x17, 0xBD}}

type IRawGameController2Interface interface {
	win32.IInspectableInterface
	Get_SimpleHapticsControllers() *IVectorView[*ISimpleHapticsController]
	Get_NonRoamableId() string
	Get_DisplayName() string
}

type IRawGameController2Vtbl struct {
	win32.IInspectableVtbl
	Get_SimpleHapticsControllers uintptr
	Get_NonRoamableId            uintptr
	Get_DisplayName              uintptr
}

type IRawGameController2 struct {
	win32.IInspectable
}

func (this *IRawGameController2) Vtbl() *IRawGameController2Vtbl {
	return (*IRawGameController2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRawGameController2) Get_SimpleHapticsControllers() *IVectorView[*ISimpleHapticsController] {
	var _result *IVectorView[*ISimpleHapticsController]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SimpleHapticsControllers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRawGameController2) Get_NonRoamableId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_NonRoamableId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

func (this *IRawGameController2) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return HStringToStrAndFree(_result)
}

// EB8D0792-E95A-4B19-AFC7-0A59F8BF759E
var IID_IRawGameControllerStatics = syscall.GUID{0xEB8D0792, 0xE95A, 0x4B19,
	[8]byte{0xAF, 0xC7, 0x0A, 0x59, 0xF8, 0xBF, 0x75, 0x9E}}

type IRawGameControllerStaticsInterface interface {
	win32.IInspectableInterface
	Add_RawGameControllerAdded(value EventHandler[*IRawGameController]) EventRegistrationToken
	Remove_RawGameControllerAdded(token EventRegistrationToken)
	Add_RawGameControllerRemoved(value EventHandler[*IRawGameController]) EventRegistrationToken
	Remove_RawGameControllerRemoved(token EventRegistrationToken)
	Get_RawGameControllers() *IVectorView[*IRawGameController]
	FromGameController(gameController *IGameController) *IRawGameController
}

type IRawGameControllerStaticsVtbl struct {
	win32.IInspectableVtbl
	Add_RawGameControllerAdded      uintptr
	Remove_RawGameControllerAdded   uintptr
	Add_RawGameControllerRemoved    uintptr
	Remove_RawGameControllerRemoved uintptr
	Get_RawGameControllers          uintptr
	FromGameController              uintptr
}

type IRawGameControllerStatics struct {
	win32.IInspectable
}

func (this *IRawGameControllerStatics) Vtbl() *IRawGameControllerStaticsVtbl {
	return (*IRawGameControllerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRawGameControllerStatics) Add_RawGameControllerAdded(value EventHandler[*IRawGameController]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RawGameControllerAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRawGameControllerStatics) Remove_RawGameControllerAdded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RawGameControllerAdded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IRawGameControllerStatics) Add_RawGameControllerRemoved(value EventHandler[*IRawGameController]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_RawGameControllerRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRawGameControllerStatics) Remove_RawGameControllerRemoved(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_RawGameControllerRemoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IRawGameControllerStatics) Get_RawGameControllers() *IVectorView[*IRawGameController] {
	var _result *IVectorView[*IRawGameController]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RawGameControllers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRawGameControllerStatics) FromGameController(gameController *IGameController) *IRawGameController {
	var _result *IRawGameController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromGameController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(gameController)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E5AEEFDD-F50E-4A55-8CDC-D33229548175
var IID_IUINavigationController = syscall.GUID{0xE5AEEFDD, 0xF50E, 0x4A55,
	[8]byte{0x8C, 0xDC, 0xD3, 0x32, 0x29, 0x54, 0x81, 0x75}}

type IUINavigationControllerInterface interface {
	win32.IInspectableInterface
	GetCurrentReading() UINavigationReading
	GetOptionalButtonLabel(button OptionalUINavigationButtons) GameControllerButtonLabel
	GetRequiredButtonLabel(button RequiredUINavigationButtons) GameControllerButtonLabel
}

type IUINavigationControllerVtbl struct {
	win32.IInspectableVtbl
	GetCurrentReading      uintptr
	GetOptionalButtonLabel uintptr
	GetRequiredButtonLabel uintptr
}

type IUINavigationController struct {
	win32.IInspectable
}

func (this *IUINavigationController) Vtbl() *IUINavigationControllerVtbl {
	return (*IUINavigationControllerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUINavigationController) GetCurrentReading() UINavigationReading {
	var _result UINavigationReading
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentReading, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUINavigationController) GetOptionalButtonLabel(button OptionalUINavigationButtons) GameControllerButtonLabel {
	var _result GameControllerButtonLabel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetOptionalButtonLabel, uintptr(unsafe.Pointer(this)), uintptr(button), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUINavigationController) GetRequiredButtonLabel(button RequiredUINavigationButtons) GameControllerButtonLabel {
	var _result GameControllerButtonLabel
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetRequiredButtonLabel, uintptr(unsafe.Pointer(this)), uintptr(button), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2F14930A-F6F8-4A48-8D89-94786CCA0C2E
var IID_IUINavigationControllerStatics = syscall.GUID{0x2F14930A, 0xF6F8, 0x4A48,
	[8]byte{0x8D, 0x89, 0x94, 0x78, 0x6C, 0xCA, 0x0C, 0x2E}}

type IUINavigationControllerStaticsInterface interface {
	win32.IInspectableInterface
	Add_UINavigationControllerAdded(value EventHandler[*IUINavigationController]) EventRegistrationToken
	Remove_UINavigationControllerAdded(token EventRegistrationToken)
	Add_UINavigationControllerRemoved(value EventHandler[*IUINavigationController]) EventRegistrationToken
	Remove_UINavigationControllerRemoved(token EventRegistrationToken)
	Get_UINavigationControllers() *IVectorView[*IUINavigationController]
}

type IUINavigationControllerStaticsVtbl struct {
	win32.IInspectableVtbl
	Add_UINavigationControllerAdded      uintptr
	Remove_UINavigationControllerAdded   uintptr
	Add_UINavigationControllerRemoved    uintptr
	Remove_UINavigationControllerRemoved uintptr
	Get_UINavigationControllers          uintptr
}

type IUINavigationControllerStatics struct {
	win32.IInspectable
}

func (this *IUINavigationControllerStatics) Vtbl() *IUINavigationControllerStaticsVtbl {
	return (*IUINavigationControllerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUINavigationControllerStatics) Add_UINavigationControllerAdded(value EventHandler[*IUINavigationController]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_UINavigationControllerAdded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUINavigationControllerStatics) Remove_UINavigationControllerAdded(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_UINavigationControllerAdded, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IUINavigationControllerStatics) Add_UINavigationControllerRemoved(value EventHandler[*IUINavigationController]) EventRegistrationToken {
	var _result EventRegistrationToken
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Add_UINavigationControllerRemoved, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(NewTwoArgFuncDelegate(value))), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUINavigationControllerStatics) Remove_UINavigationControllerRemoved(token EventRegistrationToken) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Remove_UINavigationControllerRemoved, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	_ = _hr
}

func (this *IUINavigationControllerStatics) Get_UINavigationControllers() *IVectorView[*IUINavigationController] {
	var _result *IVectorView[*IUINavigationController]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UINavigationControllers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// E0CB28E3-B20B-4B0B-9ED4-F3D53CEC0DE4
var IID_IUINavigationControllerStatics2 = syscall.GUID{0xE0CB28E3, 0xB20B, 0x4B0B,
	[8]byte{0x9E, 0xD4, 0xF3, 0xD5, 0x3C, 0xEC, 0x0D, 0xE4}}

type IUINavigationControllerStatics2Interface interface {
	win32.IInspectableInterface
	FromGameController(gameController *IGameController) *IUINavigationController
}

type IUINavigationControllerStatics2Vtbl struct {
	win32.IInspectableVtbl
	FromGameController uintptr
}

type IUINavigationControllerStatics2 struct {
	win32.IInspectable
}

func (this *IUINavigationControllerStatics2) Vtbl() *IUINavigationControllerStatics2Vtbl {
	return (*IUINavigationControllerStatics2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUINavigationControllerStatics2) FromGameController(gameController *IGameController) *IUINavigationController {
	var _result *IUINavigationController
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FromGameController, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(gameController)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// classes

type ArcadeStick struct {
	RtClass
	*IArcadeStick
}

func NewIArcadeStickStatics2() *IArcadeStickStatics2 {
	var p *IArcadeStickStatics2
	hs := NewHStr("Windows.Gaming.Input.ArcadeStick")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IArcadeStickStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIArcadeStickStatics() *IArcadeStickStatics {
	var p *IArcadeStickStatics
	hs := NewHStr("Windows.Gaming.Input.ArcadeStick")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IArcadeStickStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type FlightStick struct {
	RtClass
	*IFlightStick
}

func NewIFlightStickStatics() *IFlightStickStatics {
	var p *IFlightStickStatics
	hs := NewHStr("Windows.Gaming.Input.FlightStick")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IFlightStickStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type Gamepad struct {
	RtClass
	*IGamepad
}

func NewIGamepadStatics2() *IGamepadStatics2 {
	var p *IGamepadStatics2
	hs := NewHStr("Windows.Gaming.Input.Gamepad")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGamepadStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIGamepadStatics() *IGamepadStatics {
	var p *IGamepadStatics
	hs := NewHStr("Windows.Gaming.Input.Gamepad")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IGamepadStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type Headset struct {
	RtClass
	*IHeadset
}

type RacingWheel struct {
	RtClass
	*IRacingWheel
}

func NewIRacingWheelStatics2() *IRacingWheelStatics2 {
	var p *IRacingWheelStatics2
	hs := NewHStr("Windows.Gaming.Input.RacingWheel")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRacingWheelStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIRacingWheelStatics() *IRacingWheelStatics {
	var p *IRacingWheelStatics
	hs := NewHStr("Windows.Gaming.Input.RacingWheel")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRacingWheelStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type RawGameController struct {
	RtClass
	*IRawGameController
}

func NewIRawGameControllerStatics() *IRawGameControllerStatics {
	var p *IRawGameControllerStatics
	hs := NewHStr("Windows.Gaming.Input.RawGameController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IRawGameControllerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type UINavigationController struct {
	RtClass
	*IUINavigationController
}

func NewIUINavigationControllerStatics2() *IUINavigationControllerStatics2 {
	var p *IUINavigationControllerStatics2
	hs := NewHStr("Windows.Gaming.Input.UINavigationController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUINavigationControllerStatics2, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

func NewIUINavigationControllerStatics() *IUINavigationControllerStatics {
	var p *IUINavigationControllerStatics
	hs := NewHStr("Windows.Gaming.Input.UINavigationController")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IUINavigationControllerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}
