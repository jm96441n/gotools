package writer

import "os"

func WriteToFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0600)
	if err != nil {
		return err
	}

	return os.Chmod(path, 0600)
}

func WriteZerosToFile(path string, count int) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	for i := 0; i < count; i++ {
		_, err := file.Write([]byte{byte(0)})
		if err != nil {
			return err
		}
	}
	return nil
}
