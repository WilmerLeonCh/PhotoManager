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

type MPhotosFilter struct {
	Title string
	Start int
	Limit int
}

type MPhotosFilterOption func(*MPhotosFilter)

func WithTitle(title string) MPhotosFilterOption {
	return func(f *MPhotosFilter) {
		f.Title = title
	}
}

func WithStart(start int) MPhotosFilterOption {
	return func(f *MPhotosFilter) {
		f.Start = start
	}
}

func WithLimit(limit int) MPhotosFilterOption {
	return func(f *MPhotosFilter) {
		f.Limit = limit
	}
}

func NewMPhotosFilter(fns ...MPhotosFilterOption) MPhotosFilter {
	f := MPhotosFilter{
		Limit: 10,
	}
	for _, fn := range fns {
		fn(&f)
	}
	return f
}

type MPhotoReadResponse struct {
	Photos     []MPhoto
	TotalCount int
}

func Read(mPhotoFilter MPhotosFilter) (*MPhotoReadResponse, error) {
	req, errNewReq := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s?title_like=%s", BaseUrl, mPhotoFilter.Title),
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
	var slicePhotos []MPhoto
	totalCount := len(resPhotos)
	for i := 0; i < totalCount; i++ {
		if i >= mPhotoFilter.Start && i < mPhotoFilter.Start+mPhotoFilter.Limit {
			slicePhotos = append(slicePhotos, resPhotos[i])
		}
	}
	return &MPhotoReadResponse{
		Photos:     slicePhotos,
		TotalCount: totalCount,
	}, nil
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

func Delete(id int) error {
	req, errNewReq := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/%d", BaseUrl, id),
		nil)
	if errNewReq != nil {
		return fmt.Errorf("err | building http request: %v", errNewReq)
	}
	res, errDelete := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if errDelete != nil {
		return fmt.Errorf("err | deleting photo: %v", errDelete)
	}
	return nil
}
