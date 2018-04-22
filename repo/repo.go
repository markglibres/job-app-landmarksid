package repo

type FileWriter interface {
	Write(record []string)
	NewWriter() FileWriter
}
