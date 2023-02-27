package sssa

import (
	"testing"
)

func TestCreateCombine(t *testing.T) {
	// Short, medium, and long tests
	strings := []string{
		"N17FigASkL6p1EOgJhRaIquQLGvYV0",
		"0y10VAfmyH7GLQY6QccCSLKJi8iFgpcSBTLyYOGbiYPqOpStAf1OYuzEBzZR",
		"KjRHO1nHmIDidf6fKvsiXWcTqNYo2U9U8juO94EHXVqgearRISTQe0zAjkeUYYBvtcB8VWzZHYm6ktMlhOXXCfRFhbJzBUsXaHb5UDQAvs2GKy6yq0mnp8gCj98ksDlUultqygybYyHvjqR7D7EAWIKPKUVz4of8OzSjZlYg7YtCUMYhwQDryESiYabFID1PKBfKn5WSGgJBIsDw5g2HB2AqC1r3K8GboDN616Swo6qjvSFbseeETCYDB3ikS7uiK67ErIULNqVjf7IKoOaooEhQACmZ5HdWpr34tstg18rO",
	}

	minimum := []int{4, 6, 20}
	shares := []int{5, 100, 100}

	for i := range strings {
		created, err := Create(minimum[i], shares[i], strings[i])
		if err != nil {
			t.Fatal("Fatal: creating: ", err)
		}
		combined, err := Combine(created)
		if err != nil {
			t.Fatal("Fatal: combining: ", err)
		}
		if combined != strings[i] {
			t.Fatal("Fatal: combining returned invalid data")
		}
	}
}

func TestCreateCombinePublicKey(t *testing.T) {
	someRandomKey := `-----BEGIN RSA PUBLIC KEY-----
	MIIEhgIBAAKB/DFnL5O2LCGJQJ/6W299AsrXsHU3nsGVTbjoDqXjdHboSqAuv0ap
	oyTPQuBVNff1X0AdVDwjat2vSAukST/3PmRX4TNU4jV0rog/z6grexOCSl3oatJO
	i80t+F6uuTD6XTh5C5yDQNI/sTyaPpydbI+P87UuY4UapZaei7fwc3MfurJ+jwEJ
	c+jOWbll2YhIgCOuIe0GRX4e4CDC2KiO/BqAWCPQNjk0Y0iC2+J+2Qy3QBOJTVO8
	E2DzIhIe4VjKK6OVVesYmJWSXX/Jx382CvUDv5ss8mxGEs3yge4zeQ0GPPDaqTFw
	OJ1uppsdj10ZiW92E8v/fYwlBNGfrQIDAQABAoH8C2OCMEcavVBquXZ5haYH8sLu
	RtdfnbjRhgLY/Z0FyDOcoHimV5/boCy3egeqvVKvdpRMSuDPTfOOZECnMjvJAlDP
	9Yln7HLNmVM8h8QeR00N38Aof/rjd5VVYF5fCs9slgwxhQ8s7ksIjLPyIyCXWjER
	OX9MKe8OpT4/b1Pa1X6I28PaC3LVjDHEkigPd705i8VuF2nvZ3+Kb5uQHeczsq6f
	LfJTME0uewB2UhKokJRUlqNpRMp+N4DUhChwC9yN0EXlxzbTUsW8ouLFNKLCqlxZ
	YAJqWfCFFe14f++Ie4tfyFI0e1VP7Gge+5vxiIkuPapj5MMg47bVQHtxAn5olgFT
	/RWHAGsGPS8cyFZ3sh01qmfaPl4mgthTQdSKj/YKUXy8/u9Wqdk7DK5wfTxD6/Go
	HUpNtXkEhL5HmWxq85lhXzos98C/hNlo2GIP1X3D7IgGTwe3z2nlKUCAXLdc3Wwf
	GmpBsg3HXZ7ZBL2WwH6JNn/KgbjSioZChvMCfnjtHZMoYfgoD3ncUey4Db5K2/cu
	mWPGosswouUOnEFjYHKxo2glxxPgtOK2EgURxE/t/Tz+WPak4gl8fMjg4v8cjuxo
	kZdH6Izm0wqMrPOdKFFwFFSdwhl5DG9Ikg/UTznAprvkfmtF32/sX1Ux4NBD6paq
	XvMSLPz67VIm3wJ+JAvvkT8deFZQjOnxnv39r2uYXbLJ8JKmaKeYX7nEw60ypAPJ
	9mn3m+sWkB+iz+qaJt7ff4342ie9+iy2WH8suwAS0Vi8+Fq7+EaVmGlcAxEWM70G
	dQYwJs46NV2ueY97M2qtpVq5XMM9tIU0BqB3p8nY0voRuX5UcVyFQdC5An493VDc
	EDTOt+/y7/wZlq+xQqr18ikXGm/+c4tik+7spOKayrZGec03JiZkNbFSVpyQJ7j+
	k0EALapWIBHW0vZOfVXBLF4PfwJB03T0WLPCjgwqXaSJBYxfa8YoyH+xCXTenuiu
	B1+FkeGVaN/8vd+9rIE/QzoAMLRDWDxBYxECfkAVgjWSiLoK3fJcdwsh6e9bxgK7
	EI8yFnWyFhymHTLjACcw9DIZyiaFkpjkXjB7NX0EzWtM5FPUmfrFFLlCpHAzZ8P6
	FjXyOcfVlE9IF4gZHNUXHj8R0HflPWg9K9pfAxBhmc5+GJ6aL4smjvpp05fwPD6u
	0yYyxcpe8iznsQ==
	-----END RSA PUBLIC KEY-----`

	created,err := Create(3, 5, someRandomKey)
	if err != nil {
		t.Fatal("Fatal: creating: ", err)
	}

	if len(created) != 5 {
		t.Fatal("Fatal: number of shares created is wrong")
	}

	combined,err2 := Combine(created[1:4])
	if err2 != nil {
		t.Fatal("Fatal: combined: ", err)
	}
	if combined != someRandomKey {
		t.Fatal("Cannot decipher combined message")
	}
}

func TestLibraryCombine(t *testing.T) {
	shares := []string{
		"U1k9koNN67-og3ZY3Mmikeyj4gEFwK4HXDSglM8i_xc=yA3eU4_XYcJP0ijD63Tvqu1gklhBV32tu8cHPZXP-bk=",
		"O7c_iMBaGmQQE_uU0XRCPQwhfLBdlc6jseTzK_qN-1s=ICDGdloemG50X5GxteWWVZD3EGuxXST4UfZcek_teng=",
		"8qzYpjk7lmB7cRkOl6-7srVTKNYHuqUO2WO31Y0j1Tw=-g6srNoWkZTBqrKA2cMCA-6jxZiZv25rvbrCUWVHb5g=",
		"wGXxa_7FPFSVqdo26VKdgFxqVVWXNfwSDQyFmCh2e5w=8bTrIEs0e5FeiaXcIBaGwtGFxeyNtCG4R883tS3MsZ0=",
		"j8-Y4_7CJvL8aHxc8WMMhP_K2TEsOkxIHb7hBcwIBOo=T5-EOvAlzGMogdPawv3oK88rrygYFza3KSki2q8WEgs=",
	}

	combined, err := Combine(shares)
	if err != nil {
		t.Fatal("Fatal: combining: ", err)
	}
	if combined != "test-pass" {
		t.Fatal("Failed library cross-language check")
	}
}

func TestIsValidShare(t *testing.T) {
	shares := []string{
		"U1k9koNN67-og3ZY3Mmikeyj4gEFwK4HXDSglM8i_xc=yA3eU4_XYcJP0ijD63Tvqu1gklhBV32tu8cHPZXP-bk=",
		"O7c_iMBaGmQQE_uU0XRCPQwhfLBdlc6jseTzK_qN-1s=ICDGdloemG50X5GxteWWVZD3EGuxXST4UfZcek_teng=",
		"8qzYpjk7lmB7cRkOl6-7srVTKNYHuqUO2WO31Y0j1Tw=-g6srNoWkZTBqrKA2cMCA-6jxZiZv25rvbrCUWVHb5g=",
		"wGXxa_7FPFSVqdo26VKdgFxqVVWXNfwSDQyFmCh2e5w=8bTrIEs0e5FeiaXcIBaGwtGFxeyNtCG4R883tS3MsZ0=",
		"j8-Y4_7CJvL8aHxc8WMMhP_K2TEsOkxIHb7hBcwIBOo=T5-EOvAlzGMogdPawv3oK88rrygYFza3KSki2q8WEgs=",
		"Hello world",
	}

	results := []bool{
		true,
		true,
		true,
		true,
		true,
		false,
	}

	for i := range shares {
		if IsValidShare(shares[i]) != results[i] {
			t.Fatal("Checking for valid shares failed:", i)
		}
	}
}
