package target

type GoogleTargetOptions struct {
	Filename          string   `json:"filename"`
	Bucket            string   `json:"bucket"`
}

func (GoogleTargetOptions) isTargetOptions() {}

func NewGoogleTarget(options *GoogleTargetOptions) *Target {
	return newTarget("org.osbuild.google", options)
}
