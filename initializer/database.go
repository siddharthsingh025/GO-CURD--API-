package initializer

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
)

var DB *gorm.DB

func ConnectToDB() {
	ctx := context.Background()
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	if err := DB.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}

	tracer := otel.Tracer("posts-service")

	ctx, span := tracer.Start(ctx, "root")
	defer span.End()

	var num int
	if err := DB.WithContext(ctx).Raw("SELECT 42").Scan(&num).Error; err != nil {
		panic(err)
	}
}
