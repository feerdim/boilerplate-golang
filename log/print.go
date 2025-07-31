package log

import "fmt"

func PrintJSON(data ...interface{}) {
	for i := range data {
		fmt.Println(ParseJSON(data[i]))
	}
}

func PrettyPrint(data ...interface{}) {
	for i := range data {
		fmt.Println(ParsePrettyJSON(data[i]))
	}
}

func PrintDebug(msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)
	defaultLogger.log.Debug().Msg(msg)
}

func PrintInfo(msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)
	defaultLogger.log.Info().Msg(msg)
}

func PrintWarn(msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)
	defaultLogger.log.Warn().Msg(msg)
}

func PrintError(err error, msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)
	defaultLogger.log.Error().Err(err).Msg(msg)
}

func PrintNewError(err, newError error, fields ...interface{}) error {
	msg := generateMessage(newError.Error(), fields)
	defaultLogger.log.Error().Err(err).Msg(msg)

	return newError
}

func PrintFatal(err, newError error, fields ...interface{}) {
	msg := generateMessage(newError.Error(), fields)
	defaultLogger.log.Fatal().Err(err).Msg(msg)
}
