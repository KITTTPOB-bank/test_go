package initializers

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	_, b, _, _ := runtime.Caller(0)
	base := filepath.Dir(b)

	for _, p := range []string{
		filepath.Join(base, "..", ".env"),
		filepath.Join(base, ".env"),
	} {
		if err := godotenv.Load(p); err == nil {
			log.Printf("✅ loaded .env from: %s", p)
			return
		}
	}
	log.Println("⚠️ no .env file loaded")
}
