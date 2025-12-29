package domain

type Computer interface {
	InsertIntoLightningPort()
}

type Logger interface {
	Log(message string)
}
