package domain

type QueueService interface {
	Enqueue(name string, data []byte, retry int) error // sebuah proses untuk memasukkan kedalam antrian
}
