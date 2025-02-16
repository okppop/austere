# austere

austere is a mini TCP Application layer protocol, only have a body length field in protocol header.

This package aims to solve TCP sticky packet. But it's not production ready since it don't have any data verification means.

The checksum field in TCP header may not guarantee data integrity in all scenarios. Do not trust it too much.

## Usage

### Encode

#### Example code

```go
func main() {
	s1 := "test data1"
	s2 := "test data2"
	encoder := austere.NewEncoderWithBuffer(os.Stdout)

	// each call will create a new message
	err := encoder.EncodeAndWrite([]byte(s1))
	if err != nil {
		log.Fatal(err)
	}

	err = encoder.EncodeAndWrite([]byte(s2))
	if err != nil {
		log.Fatal(err)
	}

	// Don't forget to write buffer
	encoder.Flush()
}
```

#### Run

```sh
go run . > data
```

### Decode

#### Example code

```go
func main() {
	file, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := austere.NewDecoderWithBuffer(file)

	for {
		msg, err := decoder.ReadAndDecode()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println("message:", string(msg))
	}
}
```

#### Run

```sh
go run .
```

### Use custom buffer size

```go
func main() {
	encoder := austere.NewEncoderWithBufferSize(os.Stdout, 8192)
	decoder := austere.NewDecoderWithBufferSize(os.Stdin, 8192)
	
	// ...
}
```