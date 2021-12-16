package helper

import (
	"wilayah_indonesia_service/model/domain"
	"wilayah_indonesia_service/model/web"
)

func ToRegionResponse(region domain.Region) web.RegionResponse {
	return web.RegionResponse{
		Kode: region.Kode,
		Name: region.Name,
	}
}

func ToRegionResponses(category []domain.Region) []web.RegionResponse {
	var response []web.RegionResponse
	for _, item_data := range category {
		response = append(response, ToRegionResponse(item_data))
	}

	return response
}
