package confload

// Globals : Partie globale du fichier de conf
type Globals struct {
	StartLogging bool
	FileLog      string
	Disassamble  bool
	LogLevel     int
}

// ConfigData : Data structure du fichier de conf
type ConfigData struct {
	Globals
}
