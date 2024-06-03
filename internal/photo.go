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
	photoBody, errMarshal := json.Marshal(reqPhoto)
	if errMarshal != nil {
		return nil, fmt.Errorf("err | marshal body :%v", errMarshal)
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

func Read() ([]MPhoto, error) {
	req, errNewReq := http.NewRequest(
		http.MethodGet,
		BaseUrl,
		nil,
	)
	if errNewReq != nil {
		return nil, fmt.Errorf("err | creating request: %v", errNewReq)
	}

	res, errDo := http.DefaultClient.Do(req)
	if errDo != nil {
		return nil, fmt.Errorf("err | doing request: %v", errDo)
	}
	defer res.Body.Close()

	var resPhotos []MPhoto
	if err := json.NewDecoder(res.Body).Decode(&resPhotos); err != nil {
		return nil, fmt.Errorf("err | decoding response: %v", err)
	}
	return resPhotos, nil
}

func Update(photo MPhoto) error {
	photoBody, errMarshal := json.Marshal(photo)
	if errMarshal != nil {
		return fmt.Errorf("err | marshal body :%v", errMarshal)
	}
	req, errNewReq := http.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("%s/%d", BaseUrl, photo.Id),
		bytes.NewBuffer(photoBody),
	)
	if errNewReq != nil {
		return fmt.Errorf("err | creating request: %v", errNewReq)
	}
	res, errDo := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if errDo != nil {
		return fmt.Errorf("err | doing request: %v", errDo)
	}
	return nil
}
