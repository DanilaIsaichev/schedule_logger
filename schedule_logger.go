package schedule_logger

import (
	"errors"
	"log"
	"os"
)

type log_type byte

const (
	log_ok log_type = iota
	log_warning
	log_err
)

func (l *log_type) to_string() (log_str string, err error) {
	switch *l {
	case log_ok:
		return "OK", nil
	case log_warning:
		return "WARN", nil
	case log_err:
		return "ERROR", nil
	default:
		return "", errors.New("invalid value of log type")
	}
}

type log_struct struct {
	log_type    log_type
	log_message string
}

func Write_Log(current_log log_struct) (err error) {

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

	l_type, err := current_log.log_type.to_string()
	if err != nil {
		return err
	}

	log.Println(l_type + " " + current_log.log_message)

	return nil
}
