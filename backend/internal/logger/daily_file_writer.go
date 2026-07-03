package logger

import (
	"os"
	"path/filepath"
	"sync"
	"time"
)

type DailyFileWriter struct {
	mu         sync.Mutex
	logDir     string
	prefix     string
	currentDay string
	file       *os.File
}

func NewDailyFileWriter(logDir string, prefix string) (*DailyFileWriter, error) {
	w := &DailyFileWriter{
		logDir: logDir,
		prefix: prefix,
	}

	if err := w.rotateIfNeeded(); err != nil {
		return nil, err
	}

	return w, nil
}

func (w *DailyFileWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if err := w.rotateIfNeeded(); err != nil {
		return 0, err
	}

	return w.file.Write(p)
}

func (w *DailyFileWriter) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.file != nil {
		return w.file.Close()
	}

	return nil
}

func (w *DailyFileWriter) rotateIfNeeded() error {
	day := time.Now().Format("20060102")

	if w.file != nil && w.currentDay == day {
		return nil
	}

	if w.file != nil {
		_ = w.file.Close()
	}

	filename := w.prefix + "_" + day + ".log"
	path := filepath.Join(w.logDir, filename)

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	w.file = file
	w.currentDay = day

	return nil
}
