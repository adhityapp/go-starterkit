package bootstrap

import (
	"context"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type Container struct {
	ctx   context.Context
	db    *sqlx.DB
	trace *sdktrace.TracerProvider
}

func Init() *Container {
	c := &Container{
		ctx: context.Background(),
	}

	logrus.Debug("init config")
	c.initConfig()

	logrus.Debug("init telemetery")
	c.initTracer()

	return c
}

func (c *Container) initConfig() {
	var err error

	godotenv.Load()
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")

	if os.Getenv("CONSUL_HTTP_ADDR") == "" {
		viper.AddConfigPath(".")
		err = viper.ReadInConfig()
		// } else {
		// 	check := func() error {
		// 		val, err := os.ReadFile(os.Getenv("CONSUL_HTTP_TOKEN_FILE"))
		// 		if err != nil {
		// 			logrus.Warn(err)
		// 		}
		// 		if string(val) == "" && os.Getenv("CONSUL_HTTP_TOKEN") == "" {
		// 			return fmt.Errorf("http token empty ")
		// 		}
		// 		return nil
		// 	}

		// 	notify := func(err error, t time.Duration) {
		// 		logrus.Info(err.Error(), t)
		// 	}

		// 	b := backoff.NewExponentialBackOff()
		// 	b.MaxElapsedTime = 2 * time.Minute
		// 	err = backoff.RetryNotify(check, b, notify)
		// 	if err != nil {
		// 		logrus.Info("http token can't be retrieved")
		// 		panic(err)
		// 	}
		// 	viper.AddRemoteProvider("consul", os.Getenv("CONSUL_HTTP_ADDR"), os.Getenv("CONSUL_FILENAME"))
		// 	err = viper.ReadRemoteConfig()
	}

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Info("config file is not found in directory")
		} else {
			logrus.Fatal(err)
		}
	}
}

func (c *Container) Close() error {
	// close database
	logrus.Info("closing database connection")
	if c.db != nil {
		c.db.Close()
	}

	// logrus.Info("closing telemetry connection")
	// if c.trace != nil {
	// 	c.trace.Shutdown(c.ctx)
	// }
	return nil
}
