package helpers

import "testing"

func TestIsValidUUID(t *testing.T) {
	tests := map[string]bool{
		"":                                     false,
		"beikon":                               false,
		"Robert'); DROP TABLE Students; --   ": false,
		"000000000000000000000000000000000000": false,
		"00000000-0000-0000-0000-000000000000": false,
		"00000000-000-04000-8000-000000000000": false,
		"00000000-000-04000-8000-00000-000000": false,

		"00000000-0000-4000-0000-000000000000": false,
		"00000000-0000-4000-1000-000000000000": false,
		"00000000-0000-4000-2000-000000000000": false,
		"00000000-0000-4000-3000-000000000000": false,
		"00000000-0000-4000-4000-000000000000": false,
		"00000000-0000-4000-5000-000000000000": false,
		"00000000-0000-4000-6000-000000000000": false,
		"00000000-0000-4000-7000-000000000000": false,

		"00000000-0000-4000-8000-000000000000": true,
		"00000000-0000-4000-9000-000000000000": true,
		"00000000-0000-4000-a000-000000000000": true,
		"00000000-0000-4000-b000-000000000000": true,

		"00000000-0000-4000-c000-000000000000": false,
		"00000000-0000-4000-d000-000000000000": false,
		"00000000-0000-4000-e000-000000000000": false,
		"00000000-0000-4000-f000-000000000000": false,
		"00000000-0000-4000-g000-000000000000": false,

		"00000000-0000-5000-8000-000000000000": false,
		"00000000-0000-5000-9000-000000000000": false,
		"00000000-0000-5000-a000-000000000000": false,
		"00000000-0000-5000-b000-000000000000": false,

		"f1fa89fa-6c5e-46de-be66-26452f4a6355": true,

		"f1fa89fa-6c5e-46de-be66-26452f4a635":   false,
		"f1fa89fa-6c5e-46de-be66-26452f4a63555": false,
	}

	for uuid, expected := range tests {
		if result := IsValidUUID(uuid); result != expected {
			t.Fatalf("IsValidUUID(\"%s\") should be %t but is %t", uuid, expected, result)
		}
	}
}

// 2018-09-19 / GG -- BenchmarkIsValidUUID-8   	 3000000	       574 ns/op
func BenchmarkIsValidUUID(b *testing.B) {
	uuid := "13379001-1337-4000-8000-800851337000"

	for n := 0; n < b.N; n++ {
		IsValidUUID(uuid)
	}
}
