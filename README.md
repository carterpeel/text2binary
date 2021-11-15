# text2binary
Converts []byte slices to a binary format.

## Build Instructions
```
git clone https://github.com/carterpeel/text2binary
cd text2binary
go build
```

## Documentation
```
./text2binary
  -buffersize int
    	the size of the buffer to be used (only applies to data piped through stdin)
  -delim string
    	Set the delimiter between binary values
  -help
    	Displays this help
  -text string
    	The string to be converted to its' binary format
```

