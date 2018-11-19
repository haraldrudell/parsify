# parsify
Declarative Extensible JSON parsing by Harald Rudell

© 2018-present Harald Rudell <harald.rudell@gmail.com> (http://www.haraldrudell.com)

# Benefits
* JSON field verification and value extraction declarative
* In the declaration it is easy to see what values are relevant
* Extensible parsing functions, provided to **parsify.New**<br/>
see **parsifiers.go** for examples

# Adding to a Project
* Install to your local Go on the commad-line:<br/>
**go get -u github.com/haraldrudell/parsify**
* Import in a .go file:<br/>
**import (… "github.com/haraldrudell/parsify"…**
* Use in that same .go file:<br/>
**parsify.New(…)**

# Example Usage
* First pass checks that the JSON-RPC version is correct and that the response id matches the request id. Then returns the result object value
* Second pass extracts data from the result value: TimeStamp and BlockNumber both numeric strings (hexadecimal)
* A reader parses JSON as a stream
<pre>
	steps := New("Verify JSON-RPC response", []Step{
		{"VerifyStringProperty", "jsonrpc", aJSONRpcVersion, nil},
		{"VerifyNumberProperty", "id", float64(rqID), nil},
		{"EnterKey", "result", nil, nil},
	}, nil, nil)
	data, e := steps.ParsifyReader(strings.NewReader(string(jsonString)))
	if e != nil {
		panic(e)
	}

	type Block struct {
		TimeStamp   uint64
		BlockNumber uint64
	}
	block := Block{}
	ps := &Steps{
		"Parsing result",
		data,
		[]Step{
			{"StoreNumber", "timestamp", nil, &block.TimeStamp},
			{"StoreNumber", "number", nil, &block.BlockNumber},
		},
		0,
		nil,
	}
</pre>
The JSON input:
<pre>
{
	"jsonrpc":"2.0",
	"id":0,
	"result": {
		"timestamp":"0x5bf1ec6e",
		"number":"0x66b15f"
	}
}</pre>
© 2018-present Harald Rudell <harald.rudell@gmail.com> (http://www.haraldrudell.com)
