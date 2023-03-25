package schedule_logger

import (
	"errors"
	"log"
	"os"
)

type Log_Type byte

const (
	Log_OK Log_Type = iota
	Log_WARNING
	Log_ERROR
)

func (l *Log_Type) To_String() (log_str string, err error) {
	switch *l {
	case Log_OK:
		return "OK", nil
	case Log_WARNING:
		return "WARN", nil
	case Log_ERROR:
		return "ERROR", nil
	default:
		return "", errors.New("invalid value of log type")
	}
}

type Log_Struct struct {
	Log_Type    Log_Type
	Log_Message string
}

func Write_Log(current_log Log_Struct) (err error) {

	if _, err := os.Stat("backend_logs/"); os.IsNotExist(err) {

		// Если директория не существует - создаём
		err := os.Mkdir("backend_logs/", 0777)
		if err != nil {
			log.Fatal(err)
		}

	}

	file, err := os.OpenFile("./backend_logs/backend.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	log.SetOutput(file)

	l_type, err := current_log.Log_Type.To_String()
	if err != nil {
		return err
	}

	log.Println(l_type + " " + current_log.Log_Message)

	return nil
}
