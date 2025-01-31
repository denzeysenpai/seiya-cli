package utils

const (
	STANDARD_FORMAT  = "2006-01-02 15:04:05" // FORMAT FOR LOGS
	TIMESTAMP_FORMAT = "200601021504"        // TIMESTAMP
	DATE_FORMAT      = "2006-01-02"          // DATE
	TIME_FORMAT      = "15:04:05"            // TIME
	Reset            = "\033[0m"             // Reset Colors
	Red              = "\033[31m"            // Color
	Green            = "\033[32m"            // Color
	Yellow           = "\033[33m"            // Color
	Blue             = "\033[34m"            // Color
	Magenta          = "\033[35m"            // Color
	Cyan             = "\033[36m"            // Color
	Gray             = "\033[37m"            // Color
	White            = "\033[97m"            //
	HEADER           = "header"              // header task
	TASK             = "task"                // normal task
)

func CheckEror(err error) {
	if err != nil {

	}
}
