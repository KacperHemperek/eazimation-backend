package services

type SqlScanner interface {
	Scan(dest ...any) error
}
