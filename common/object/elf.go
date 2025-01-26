package object

import (
	"bytes"
)

type ELF struct {
	Header  *Header
	Tracker *Tracker
}

func (e *ELF) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Encode Header
	err := e.Header.Encode(buf)
	if err != nil {
		return nil, err
	}

	// Encode Tracker
	trackerData, err := e.Tracker.Encode()
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(trackerData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (e *ELF) Decode(data []byte) error {
	buf := bytes.NewReader(data)

	// Decode Header
	e.Header = &Header{}
	err := e.Header.Decode(buf)
	if err != nil {
		return err
	}

	// Decode Tracker
	trackerData := make([]byte, buf.Len())
	_, err = buf.Read(trackerData)
	if err != nil {
		return err
	}
	e.Tracker = &Tracker{}
	err = e.Tracker.Decode(trackerData)
	if err != nil {
		return err
	}

	return nil
}
