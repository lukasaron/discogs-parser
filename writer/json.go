package writer

import (
	"bytes"
	"encoding/json"
	"github.com/lukasaron/discogs-parser/decode"
	"os"
)

type JsonWriter struct {
	option Options
	f      *os.File
	buffer bytes.Buffer
	err    error
}

func NewJsonWriter(fileName string, options ...Options) Writer {
	j := &JsonWriter{
		buffer: bytes.Buffer{},
	}

	j.f, j.err = os.Create(fileName)

	if options != nil && len(options) > 0 {
		j.option = options[0]
	}

	return j
}

func (j *JsonWriter) Reset() error {
	j.buffer.Reset()
	return nil
}

func (j *JsonWriter) Close() error {
	return j.f.Close()
}

func (j *JsonWriter) WriteArtist(a decode.Artist) error {
	if j.option.ExcludeImages {
		a.Images = nil
	}

	j.marshalAndWrite(a)
	j.flush()
	j.clean()

	return j.err
}

func (j *JsonWriter) WriteArtists(artists []decode.Artist) error {
	j.writeInitial()

	for _, a := range artists {
		j.writeDelimiter()

		if j.option.ExcludeImages {
			a.Images = nil
		}

		j.marshalAndWrite(a)
		if j.err != nil {
			return j.err
		}
	}

	j.writeClosing()
	j.flush()
	j.clean()

	return j.err
}

func (j *JsonWriter) WriteLabel(label decode.Label) error {
	if j.option.ExcludeImages {
		label.Images = nil
	}

	j.marshalAndWrite(label)
	j.flush()
	j.clean()

	return j.err
}

func (j *JsonWriter) WriteLabels(labels []decode.Label) error {
	j.writeInitial()

	for _, l := range labels {
		j.writeDelimiter()

		if j.option.ExcludeImages {
			l.Images = nil
		}

		j.marshalAndWrite(l)
		if j.err != nil {
			return j.err
		}
	}

	j.writeClosing()
	j.flush()
	j.clean()

	return j.err
}

func (j *JsonWriter) WriteMaster(master decode.Master) error {
	if j.option.ExcludeImages {
		master.Images = nil
	}

	j.marshalAndWrite(master)
	j.flush()
	j.clean()

	return j.err
}

func (j *JsonWriter) WriteMasters(masters []decode.Master) error {
	j.writeInitial()
	for _, m := range masters {
		j.writeDelimiter()

		if j.option.ExcludeImages {
			m.Images = nil
		}

		j.marshalAndWrite(m)
		if j.err != nil {
			return j.err
		}
	}

	j.writeClosing()
	j.flush()
	j.clean()

	return j.err
}
func (j *JsonWriter) WriteRelease(release decode.Release) error {
	if j.option.ExcludeImages {
		release.Images = nil
	}

	j.marshalAndWrite(release)
	j.flush()
	j.clean()
	return j.err
}

func (j *JsonWriter) WriteReleases(releases []decode.Release) error {
	j.writeInitial()

	for _, r := range releases {
		j.writeDelimiter()

		if j.option.ExcludeImages {
			r.Images = nil
		}

		j.marshalAndWrite(r)
		if j.err != nil {
			return j.err
		}
	}

	j.writeClosing()
	j.flush()
	j.clean()

	return j.err
}

func (j *JsonWriter) marshalAndWrite(d interface{}) {
	if j.err != nil {
		return
	}

	b, err := json.Marshal(d)
	if err != nil {
		j.err = err
		return
	}

	_, j.err = j.f.Write(b)
}

func (j *JsonWriter) writeDelimiter() {
	if j.err == nil && j.buffer.Len() > 0 {
		_, j.err = j.buffer.WriteString(",")
	}
}

func (j *JsonWriter) writeInitial() {
	if j.err != nil {
		return
	}
	_, j.err = j.buffer.WriteString("[")
}

func (j *JsonWriter) writeClosing() {
	if j.err != nil {
		return
	}

	_, j.err = j.buffer.WriteString("]")
}

func (j *JsonWriter) flush() {
	if j.err != nil {
		return
	}

	_, j.err = j.f.Write(j.buffer.Bytes())
}

func (j *JsonWriter) clean() {
	j.buffer.Reset()
}
