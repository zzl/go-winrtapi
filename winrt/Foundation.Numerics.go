package winrt

// structs

type Matrix3x2 struct {
	M11 float32
	M12 float32
	M21 float32
	M22 float32
	M31 float32
	M32 float32
}

type Matrix4x4 struct {
	M11 float32
	M12 float32
	M13 float32
	M14 float32
	M21 float32
	M22 float32
	M23 float32
	M24 float32
	M31 float32
	M32 float32
	M33 float32
	M34 float32
	M41 float32
	M42 float32
	M43 float32
	M44 float32
}

type Plane struct {
	Normal Vector3
	D      float32
}

type Quaternion struct {
	X float32
	Y float32
	Z float32
	W float32
}

type Rational struct {
	Numerator   uint32
	Denominator uint32
}

type Vector2 struct {
	X float32
	Y float32
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

type Vector4 struct {
	X float32
	Y float32
	Z float32
	W float32
}
