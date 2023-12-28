package model

type City struct {
	ID            int    `bson:"id,empty"`
	NamaKabupaten string `bson:"nama_kabupaten,empty"`
}
