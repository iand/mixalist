package imagegen

import (
	"github.com/iand/mixalist/pkg/blobstore"
	"image"
	"image/png"
	
	_ "image/gif"
	_ "image/jpeg"
)

func GetImageFromBlobstore(id blobstore.ID) (img image.Image, err error) {
	r, err := blobstore.DefaultStore.Open(id)
	if err != nil {
		return nil, err
	}

	img, _, err = image.Decode(r)
	r.Close()
	if err != nil {
		return nil, err
	}
	
	return img, nil
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

func GeneratePlaylistImage(entryBlobIDs []blobstore.ID) (playlistBlobID blobstore.ID, err error) {
	imgs := make([]image.Image, len(entryBlobIDs))
	for i, entryBlobID := range entryBlobIDs {
		imgs[i], err = GetImageFromBlobstore(entryBlobID)
		if err != nil {
			return "", err
		}
	}
	
	c := &Compositer{
		DestWidth: 100,
		TilesPerSide: 3,
		LugDepthRatio: 0.25,
		LugWidthRatio: 0.33,
	}
	
	img := c.Composite(imgs)
	return SaveImageToBlobstore(img)
}
