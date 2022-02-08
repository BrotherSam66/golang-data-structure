package encodedecode

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode 消息编码包
func Encode(mes string) (retByte []byte, err error) {
	// 消息长度，以int32表示
	length := int32(len(mes))
	var pkg = new(bytes.Buffer)
	// 写入4byte消息头（消息长度）
	err = binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息本身
	err = binary.Write(pkg, binary.LittleEndian, []byte(mes))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

// Decode 消息解码包
func Decode(reader *bufio.Reader) (retStr string, err error) {
	lenByte, _ := reader.Peek(4)          // 读取表示长度的4个字节
	lenBuffer := bytes.NewBuffer(lenByte) // 缓冲byte类型的缓冲器
	var length int32
	err = binary.Read(lenBuffer, binary.LittleEndian, &length) // ?????
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < (length + 4) {
		return "", err
	}
	//读取真正的消息
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
