package machine

const (
	DefaultFilamentsPath = "0:/filaments"
	DefaultGCodesPath    = "0:/gcodes"
	DefaultMacrosPath    = "0:/macros"
	DefaultSystemPath    = "0:/sys"
	DefaultWWWPath       = "0:/www"
)

// Directories holds information about the directory structure
type Directories struct {
	// Filaments is the path to filaments directory
	Filaments string
	// GCodes is the path to the gcodes directory
	GCodes string
	// Macros is the path to the macros directory
	Macros string
	// System is the path to the sys directory
	System string
	// WWW is the path to the www directory
	WWW string
}

// NewDirectories returns an instance with all paths set to their defaults
func NewDirectories() *Directories {
	return &Directories{
		Filaments: DefaultFilamentsPath,
		GCodes:    DefaultGCodesPath,
		Macros:    DefaultMacrosPath,
		System:    DefaultSystemPath,
		WWW:       DefaultWWWPath,
	}
}
