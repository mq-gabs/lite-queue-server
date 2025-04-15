package handler

import (
	"bufio"
	"errors"
	"fmt"
	manager "lite_queue_server/manager"
	protocol "lite_queue_server/protocol"
	"lite_queue_server/utils"
	"net"
)

type Handler struct {
	conn    net.Conn
	reader  *bufio.Reader
	manager *manager.QueueManager
}

func New(conn net.Conn, manager *manager.QueueManager) *Handler {
	return &Handler{
		conn:    conn,
		reader:  bufio.NewReader(conn),
		manager: manager,
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

func (h *Handler) responseError(err error) {
	var bytes [][]byte

	data := []byte(err.Error())

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
		h.responseError(fmt.Errorf("cannot read action: %v", err))
		return
	}

	if len(action) != 2 {
		h.responseError(fmt.Errorf("invalid action: %v", action))
		return
	}

	switch action[0] {
	case protocol.RequestNewQueue:
		h.newQueue()
	case protocol.RequestPush:
		h.push()
	case protocol.RequestPop:
		h.pop()
	default:
		h.responseError(fmt.Errorf("action not mapped: %v", action[0]))
	}
}

func (h *Handler) newQueue() {
	name, err := h.readNext()

	if err != nil {
		h.responseError(fmt.Errorf("cannot read name: %v", err))
		return
	}

	h.manager.NewQueue(string(name))

	h.responseSuccess(nil)
}

func (h *Handler) push() {
	name, err := h.readNext()

	if err != nil {
		h.responseError(fmt.Errorf("cannot read name: %v", err))
		return
	}

	data, err := h.readNext()

	if err != nil {
		h.responseError(fmt.Errorf("cannot read data: %v", err))
	}

	err = h.manager.Push(string(name), data)

	if err != nil {
		h.responseError(fmt.Errorf("cannot push: %v", err))
		return
	}

	h.responseSuccess(nil)
}

func (h *Handler) pop() {
	name, err := h.readNext()

	if err != nil {
		h.responseError(fmt.Errorf("cannot read name: %v", err))
		return
	}

	res, err := h.manager.Pop(string(name))

	if err != nil {
		h.responseError(fmt.Errorf("cannot pop: %v", err))
		return
	}

	h.responseSuccess(res)
}
