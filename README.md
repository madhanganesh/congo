# A Config Reader for Golang applications

## Usage

```
	config := New()
	err := config.LoadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.GetString("key1")) // prints "value-1"
```

## Sample JSON file
```
	{
		"key1": "value-1",
		"key2": "value-2"
	}
```
