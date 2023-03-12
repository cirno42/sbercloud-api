package ims

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/models/imsModels"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
)

type getImagesListResponse struct {
	Images []imsModels.ImageModel `json:"images"`
}

func GetImagesList(platform string) ([]imsModels.ImageModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.ImsEndpoint) + "/v2/cloudimages?")
	if platform != "" {
		endpoint += "__platform=" + platform
	}
	var imgs getImagesListResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &imgs, nil)
	return imgs.Images, err
}
