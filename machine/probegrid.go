package machine

// ProbeGrid holds information about the configured probe grid (see M557)
type ProbeGrid struct {
	// XMin is the X start coordinate of the heightmap
	XMin float64
	// XMax is the X end coordinate of the heightmap
	XMax float64
	// XSpacing is the spacing between probe points in X direction
	XSpacing float64
	// YMin is the Y start coordinate of the heightmap
	YMin float64
	// YMax is the Y end coordinate of the heightmap
	YMax float64
	// YSpacing is the spacing between probe points in Y direction
	YSpacing float64
	// Radius is the probing radius on delta kinematics
	Radius float64
}
