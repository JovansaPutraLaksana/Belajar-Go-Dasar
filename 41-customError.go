package main

type validationError struct {
	Msg string
}

func (e *validationError) Error() string {
	return e.Msg
}

type notFoundError struct {
	Msg string
}

func (e *notFoundError) Error() string {
	return e.Msg
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Msg: "ID tidak boleh kosong"}
	}

	if id != "123" {
		return &notFoundError{Msg: "Data dengan ID tersebut tidak ditemukan"}
	}

	// oke
	return nil
}

func main() {
	err := SaveData("", nil)
	// Menggunakan if statement
	if err != nil {
		if validationError, ok := err.(*validationError); ok {
			println("Terjadi kesalahan validasi:", validationError.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			println("Terjadi kesalahan data tidak ditemukan:", notFoundError.Error())
		} else {
			println("Terjadi kesalahan lainnya:", err.Error())
		}
	} else {
		println("Data berhasil disimpan")
	}

	// Menggunakan type switch
	switch e := err.(type) {
	case *validationError:
		println("Terjadi kesalahan validasi (dengan switch):", e.Error())
	case *notFoundError:
		println("Terjadi kesalahan data tidak ditemukan (dengan switch):", e.Error())
	default:
		println("Terjadi kesalahan lainnya (dengan switch):", err.Error())
	}
}
