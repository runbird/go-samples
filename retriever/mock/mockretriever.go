package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	//panic("implement me")
	return r.Contents
}
