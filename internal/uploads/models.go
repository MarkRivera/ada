package uploads

type UploadRequestBody struct {
	Title string 			`json:"title"`
	Description string 		`json:"description"`
	FileType string 		`json:"fileType"`
	TotalChunks int 		`json:"totalChunks"`
}