package main

import (
	"context"
	"log"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/startup"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/startup/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func main() {
	var err error
	tp, err = initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
