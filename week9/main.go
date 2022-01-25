package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filepath")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	rr := bufio.NewReader(f)
	p, err := GoIMDecoder(rr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)
}

const (
	// MaxBodySize max proto body size
	MaxBodySize = uint32(1 << 12)
)

const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_heartSize     = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	_maxPackSize   = MaxBodySize + uint32(_rawHeaderSize)
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	_heartOffset  = _seqOffset + _seqSize
)

var (
	// ErrProtoPackLen proto packet len error
	ErrProtoPackLen = errors.New("default server codec pack length error")
	// ErrProtoHeaderLen proto header len error
	ErrProtoHeaderLen = errors.New("default server codec header length error")
)

type Header struct {
	Ver uint32
	Op  uint32
	Seq uint32
}

type Proto struct {
	Header
	Body []byte
}

func GoIMDecoder(rr *bufio.Reader) (p *Proto, err error) {
	p = &Proto{}
	var (
		bodyLen   int
		headerLen uint16
		packLen   uint32
		buf       []byte
	)
	//假设一个包一行: delimiter based
	if buf, _, err = rr.ReadLine(); err != nil {
		return nil, err
	}

	packLen = binary.BigEndian.Uint32(buf[_packOffset:_headerOffset])
	headerLen = binary.BigEndian.Uint16(buf[_headerOffset:_verOffset])
	p.Ver = uint32(binary.BigEndian.Uint16(buf[_verOffset:_opOffset]))
	p.Op = binary.BigEndian.Uint32(buf[_opOffset:_seqOffset])
	p.Seq = binary.BigEndian.Uint32(buf[_seqOffset:])
	if packLen > _maxPackSize {
		return nil, ErrProtoPackLen
	}
	if headerLen != _rawHeaderSize {
		return nil, ErrProtoHeaderLen
	}
	if bodyLen = int(packLen - uint32(headerLen)); bodyLen > 0 {
		p.Body = buf[headerLen:]
	} else {
		p.Body = nil
	}
	return p, nil
}
