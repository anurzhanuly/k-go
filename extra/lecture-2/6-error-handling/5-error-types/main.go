package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

// Использование errors.Is и errors.As для определения типа ошибок

// Ошибка чтения файла
type FileReadErr struct {
	// По какому пути пытались прочитать файл
	Path string
	// Исходная ошибка, возвращённая функцией os.ReadFile
	OriginalErr error
}

// Возвращает описание ошибки.
// Благодаря этой функции FileReadErr реализует интерфейс error.
func (f FileReadErr) Error() string {
	return fmt.Sprintf("не удалось прочитать файл %s: %s", f.Path, f.OriginalErr)
}

// Отдельный тип ошибки о невалидном формате файла
var InvalidFileType = errors.New("invalid file type")

func main() {
	file, err := getFile("main.go1")

	// пример использования errors.Is: отдельно обрабатываем ошибку о невалидном типе файла
	if errors.Is(err, InvalidFileType) {
		fmt.Println("Чтение файлов данного типа запрещено")
		os.Exit(1)
	}

	// пример использования errors.As: превращаем "абстрактную" ошибку err
	// в конкретную ошибку readErr, из которой можно прочитать кастомные свойства
	var readErr *FileReadErr
	if errors.As(err, &readErr) {
		fmt.Println("Ошибка при чтении файла по пути " + readErr.Path)
		os.Exit(2)
	}

	fmt.Println(file)
}

func getFile(path string) (string, error) {
	// валидируем тип файла
	if strings.HasSuffix(path, ".md") {
		return "", InvalidFileType
	}

	file, err := os.ReadFile(path)
	if err != nil {
		// логируем ошибку
		logrus.WithError(err).WithField("path", path).Error("не удалось прочитать файл")

		// возвращаем ошибку кастомного типа
		return "", &FileReadErr{OriginalErr: err, Path: path}
	}

	return string(file), nil
}
