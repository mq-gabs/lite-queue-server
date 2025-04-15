package handler

import (
	"bufio"
	"errors"
	"fmt"
	protocol "lite_queue_server/protocol"
	"lite_queue_server/utils"
	"net"
)

type Handler struct {
	conn   net.Conn
	reader *bufio.Reader
}

func New(conn net.Conn) *Handler {
	return &Handler{
		conn:   conn,
		reader: bufio.NewReader(conn),
	}
}

func (h *Handler) readNext() ([]byte, error) {
	if h.conn == nil {
		return nil, errors.New("conn cannot be nil")
	}

	return h.reader.ReadBytes(protocol.Seperator)
}

func (h *Handler) responseSuccess(data []byte) {
	var bytes [][]byte

	if data == nil {
		bytes = append(bytes, []byte{protocol.ResponseSuccessEmpty, protocol.Seperator})
	} else {
		bytes = append(bytes, []byte{protocol.ResponseSuccessContent, protocol.Seperator}, data)
	}

	flattenBytes := utils.FlattenBytes(bytes)

	h.conn.Write(flattenBytes)
}

func (h *Handler) responseError(data []byte) {
	var bytes [][]byte

	if data == nil {
		bytes = append(bytes, []byte{protocol.ResponseErrorEmpty, protocol.Seperator})
	} else {
		bytes = append(bytes, []byte{protocol.ResponseErrorContent, protocol.Seperator}, data)
	}

	flattenBytes := utils.FlattenBytes(bytes)

	h.conn.Write(flattenBytes)
}

func (h *Handler) Handle() {
	defer h.conn.Close()

	action, err := h.readNext()

	if err != nil {
		h.responseError([]byte(fmt.Errorf("cannot read action: %v", err).Error()))
	}

	if len(action) != 2 {
		h.responseError([]byte(fmt.Errorf("invalid action: %v", action).Error()))
	}

	switch action[0] {
	case protocol.RequestNewQueue:
		h.newQueue()
	case protocol.RequestPush:
		h.push()
	case protocol.RequestPop:
		h.pop()
	default:
		h.responseError([]byte(fmt.Errorf("action not mapped: %v", action[0]).Error()))
	}
}

func (h *Handler) newQueue() {

}

func (h *Handler) push() {

}

func (h *Handler) pop() {

}
