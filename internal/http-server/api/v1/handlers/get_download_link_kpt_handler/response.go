package get_download_link_kpt_handler

import (
	"github.com/ShevelevEvgeniy/app/internal/dto"
	"github.com/ShevelevEvgeniy/app/lib/api/response"
)

type Response struct {
	Response response.Response
	Result   *dto.KptInfo `json:"info"`
}
