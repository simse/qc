package source

// File represents a file that may be worked on by qc
type File struct {
	Key       string
	Extension string
	Format    string
	Path      string
}
