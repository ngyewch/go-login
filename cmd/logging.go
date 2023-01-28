package cmd

import (
	slog "github.com/go-eden/slf4go"
	"github.com/ngyewch/slf4go-zerolog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"os"
)

func InitLogging() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr}

	var writers []io.Writer
	writers = append(writers, consoleWriter)

	zerologLogger := zerolog.New(io.MultiWriter(writers...)).With().Timestamp().Logger()

	slog.SetDriver(slf4go_zerolog.NewZerologDriver(&zerologLogger))
}
