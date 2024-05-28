package model

type tempat struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Tempat  string             `bson:"nama_tempat,omitempty" json:"nama_tempat,omitempty"`
	Lokasi       string      `bson:"lokasi,omitempty" json:"lokasi,omitempty"`
	Fasilitas    string             `bson:"fasilitas,omitempty" json:"fasilitas,omitempty"`
	Lon         float64            `bson:"lon,omitempty" json:"lon,omitempty"`
	Lat         float64            `bson:"lat,omitempty" json:"lat,omitempty"`

}