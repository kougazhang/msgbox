package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Lark struct {
	URL   string `json:"url"`
	Retry int    `json:"retry"`
}

type Text struct {
	MsgType string  `json:"msg_type"`
	Content Content `json:"content"`
}

type Content struct {
	Text string `json:"text"`
}

type Response struct {
	StatusMessage string `json:"StatusMessage"`
	StatusCode    int    `json:"StatusCode"`
}

func (f Lark) Send(msg any) error {
	if f.Retry == 0 {
		f.Retry = 1
	}
	var err error
	for i := 0; i < f.Retry; i++ {
		if err = f.run(msg); err == nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return err
}

func (f Lark) run(s any) error {
	switch val := s.(type) {
	case Text:
		data, err := json.Marshal(val)
		if err != nil {
			return err
		}
		resp, err := http.Post(f.URL, "application/json", bytes.NewReader(data))
		if err != nil {
			return err
		}
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var response Response
		if err = json.Unmarshal(data, &response); err != nil {
			return err
		}
		if response.StatusCode == 0 {
			return err
		}
		return fmt.Errorf("resp:%v", response)
	default:
		return fmt.Errorf("unsupport type %v", val)
	}
}
