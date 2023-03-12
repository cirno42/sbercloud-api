package ims

import (
	"fmt"
	"sbercloud-api/api/endpoints"
	"sbercloud-api/api/models/imsModels"
	"sbercloud-api/internal/handlers/requestMakers"
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
