package encrypt

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"
)

type Vault struct {
	enckey   string
	filepath string
	mux      sync.Mutex
	mapping  map[string]string
}

func NewVault(ekey, fp string) *Vault {
	return &Vault{
		enckey: ekey,

		filepath: fp,
	}
}

func (v *Vault) LoadMapping() error {
	fptr, err := os.Open(v.filepath)
	if err != nil {
		v.mapping = make(map[string]string)
		return nil
	}

	reader, err := DecReader(v.enckey, fptr)
	if err != nil {
		return err
	}
	return v.readmapping(reader)
}

func (v *Vault) readmapping(r io.Reader) error {
	return json.NewDecoder(r).Decode(&v.mapping)
}

func (v *Vault) writemapping(w io.Writer) error {
	return json.NewEncoder(w).Encode(v.mapping)
}

func (v *Vault) savemapping() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	writer, err := EncWriter(v.enckey, f) //rgenerates cipher.stream based on key,returns struct streamwriter{writer,stream}  and passes filewriter to
	if err != nil {                       // to writer and cipher.stream to stream.
		return err // internally streamwriter uses XORKeyStream() on stream and passes it to writer on write
	} //				.|.
	return v.writemapping(writer) // cipher.StreamWriter writes to json.NewEncoder(<writer>) (interface chaining)
}

func (v *Vault) Set(key, value string) error {
	v.mux.Lock()
	defer v.mux.Unlock()
	err := v.LoadMapping()
	if err != nil {
		return err
	}
	v.mapping[key] = value
	return v.savemapping()
}

func (v *Vault) Get(key string) (string, error) {
	v.mux.Lock()
	defer v.mux.Unlock()
	err := v.LoadMapping()
	if err != nil {
		//fmt.Println(err.Error() + "1")
		return "", err
	}
	val, bol := v.mapping[key]
	if !bol {
		return "", errors.New("Value does not exists")
	}

	return val, nil
}
