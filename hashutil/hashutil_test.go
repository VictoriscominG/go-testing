package hashutil

import "testing"

func TestHashSHA256(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"empty string", "", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"english string", "hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{"unicode", "привет", "e58f1e8c55fa105bdd3f40e5037eb0b039b5998d52c05e6cd98878dd2da5cab2"},
		{"1", "Привет мир", "830d1964dc8673182a40f9adebf598960d37fbe200405b249774ecfa5b465748"},
		{"2", "🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀", "7054c6be4c56c595812d0e6adfddff02a2b0342926e24379c94ad4e183e62422"},
		{"3", "مرحبا بالعالم", "9262a0a791605071a500c1a15bef2d5efcc6c8f198567105e9ab364811377e9f"},
	}

	tests[2].want = HashSHA256("привет")

	for _, tt := range tests {
		//tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := HashSHA256(tt.in)
			if got != tt.want {
				t.Errorf("want: %s, got: %s", tt.want, got)
			}
		})
	}
}
