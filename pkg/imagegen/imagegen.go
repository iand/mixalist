package imagegen

import (
	"github.com/iand/mixalist/pkg/blobstore"
	"image"
	"image/png"
)

func GetImageFromBlobstore(id blobstore.ID) (img image.Image, err error) {
	r, err := blobstore.DefaultStore.Open(id)
	if err != nil {
		return nil, err
	}

	img, _, err = image.Decode(r)
	r.Close()
	return img, err
}

func SaveImageToBlobstore(img image.Image) (id blobstore.ID, err error) {
	id, w, err := blobstore.DefaultStore.Create()
	if err != nil {
		return "", err
	}

	err = png.Encode(w, img)
	w.Close()
	if err != nil {
		blobstore.DefaultStore.Delete(id)
		return "", err
	}

	return id, nil
}
