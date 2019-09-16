package figma

// Image format
const (
	JPG ImageFormat = "JPG"
	PNG ImageFormat = "PNG"
	SVG ImageFormat = "SVG"
)

// EntryExportSetting represents the export setting entry
type EntryExportSetting struct {
	// Suffix is the File suffix to append to all filenames
	Suffix string `json:"suffix"`
	// Format is the image type
	Format ImageFormat `json:"format"`
	// Constraint determines sizing of the exported asset
	Constraint EntryConstraint `json:"constraint"`
}

// ImageFormat represents the various types of image formats
type ImageFormat string
