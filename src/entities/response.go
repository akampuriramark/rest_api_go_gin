package entities

//Response defines the error structure to be returned
type Response struct {
	StatusCode        int
	StatusDescription string
}

//PostResponse defines the error structure to be returned
type PostResponse struct {
	Response
	InsertedID int64
}
