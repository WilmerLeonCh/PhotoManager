package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MPhoto struct {
	AlbumId      int    `json:"albumId,omitempty"`
	Id           int    `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Url          string `json:"url,omitempty"`
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`
}

const BaseUrl = "https://jsonplaceholder.typicode.com/photos"

func Create(reqPhoto MPhoto) (*MPhoto, error) {
	photoBody, errMarshaling := json.Marshal(reqPhoto)
	if errMarshaling != nil {
		return nil, fmt.Errorf("err | marshalling :%v", errMarshaling)
	}
	req, errNewReq := http.NewRequest(
		http.MethodPost,
		BaseUrl,
		bytes.NewBuffer(photoBody),
	)
	if errNewReq != nil {
		return nil, fmt.Errorf("err | creating request: %v", errNewReq)
	}
	req.Header.Add("Content-Type", "application/json")

	res, errDo := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if errDo != nil {
		return nil, fmt.Errorf("err | doing request: %v", errDo)
	}

	var resPhoto MPhoto
	if errDecode := json.NewDecoder(res.Body).Decode(&resPhoto); errDecode != nil {
		return nil, fmt.Errorf("err | decoding response: %v", errDecode)
	}
	return &resPhoto, nil
}
