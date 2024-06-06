package entities

type Formats struct {
	Id         int
	Format     string
	StreamType string //Audio or Video
}

type VideoFormat struct {
	Id       int
	VideoId  int
	FormatId int
}

//format operations will go inside videos.go
//operations include READ, INSERT, DELETE
