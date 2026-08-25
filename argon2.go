package knock

type Argon2Config struct {
	HashRaw   []byte
	Salt      []byte
	Time      uint32
	Memory    uint32
	Threads   uint8
	KeyLength uint32
}
