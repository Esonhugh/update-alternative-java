package define

type Java struct {
	Version       int    // Version of java Head
	DetailVersion string // more Detail Version of Java
	JavaHome      string //	JavaHome location
}

const (
	JAVA_8  = 8
	JAVA_11 = 11
	JAVA_12 = 12
	JAVA_14 = 14
	JAVA_16 = 16
	JAVA_17 = 17
	JAVA_18 = 18
)

var JAVA_VERSION_MAP = map[string]int{
	"JAVA_8":  JAVA_8,
	"JAVA_11": JAVA_11,
	"JAVA_12": JAVA_12,
	"JAVA_14": JAVA_14,
	"JAVA_16": JAVA_16,
	"JAVA_17": JAVA_17,
	"JAVA_18": JAVA_18,
}

func GetJavaVersionList() []string {
	keys := make([]string, 0, len(JAVA_VERSION_MAP))
	for k := range JAVA_VERSION_MAP {
		keys = append(keys, k)
	}
	return keys
}
