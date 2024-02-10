package argon2id

import "github.com/alexedwards/argon2id"

type PasswordService struct {
	// The amount of memory used by the algorithm (in kibibytes).
	// 64 << 10 (64mb) or more is recommended
	Memory uint32

	// The number of iterations over the memory.
	Iterations uint32

	// The number of threads (or lanes) used by the algorithm.
	// Recommended value is between 1 and runtime.NumCPU().
	Parallelism uint8

	// Length of the random salt. 16 bytes is recommended for password hashing.
	SaltLength uint32

	// Length of the generated key. 32 bytes or more is recommended.
	KeyLength uint32
}

func (ps *PasswordService) Hash(password string) (string, error) {
	return argon2id.CreateHash(password, &argon2id.Params{
		Memory:      ps.Memory,
		Iterations:  ps.Iterations,
		Parallelism: ps.Parallelism,
		SaltLength:  ps.SaltLength,
		KeyLength:   ps.KeyLength,
	})
}

func (ps *PasswordService) Verify(password string, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}
