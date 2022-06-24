package log

func Initialize() {
	initZap()
}

func Sync() {
	syncZap()
}
