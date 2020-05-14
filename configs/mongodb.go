package configs

import (
	"os"
)

// MongoDBConfig object
type MongoDBConfig struct {
	URI string `env:"MONGO_URI"` // i.e. "mongodb://localhost:27017"
}

// GetMongoDBConfig returns MongoDBConfig object
func GetMongoDBConfig() MongoDBConfig {
	return MongoDBConfig{
		URI: os.Getenv("MONGO_URI"),
	}
}
