package uploader

type Uploader interface {
	Upload(string, []byte) (any, error)
}

func NewUploader() {}
