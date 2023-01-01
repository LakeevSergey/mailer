package json

import (
	"math"
	"net/http"
	"strconv"

	responsejson "github.com/LakeevSergey/mailer/internal/application/api/response/json"
	"github.com/LakeevSergey/mailer/internal/application/dto"
	templatemanagerdto "github.com/LakeevSergey/mailer/internal/domain/templatemanager/dto"
)

func (a *JSONApi) SearchTemplates() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var page, perPage int64
		var err error

		pageStr := r.URL.Query().Get("page")

		if pageStr == "" {
			page = 1
		} else {
			page, err = strconv.ParseInt(pageStr, 10, 64)
		}
		if err != nil {
			responsejson.ErrorResponse(err.Error(), http.StatusBadRequest).Write(rw)
			return
		} else if page < 1 {
			responsejson.ErrorResponse("page should be positive number", http.StatusBadRequest).Write(rw)
			return
		}

		perPageStr := r.URL.Query().Get("per_page")

		if perPageStr == "" {
			perPage = 10
		} else {
			perPage, err = strconv.ParseInt(perPageStr, 10, 64)
		}
		if err != nil {
			responsejson.ErrorResponse(err.Error(), http.StatusBadRequest).Write(rw)
			return
		} else if perPage < 1 {
			responsejson.ErrorResponse("per_page should be positive number", http.StatusBadRequest).Write(rw)
			return
		}

		items, totalCount, err := a.templateManager.Search(r.Context(), templatemanagerdto.Search{
			Limit:  perPage,
			Offset: perPage * (page - 1),
		})
		if err != nil {
			responsejson.ErrorResponse(err.Error(), http.StatusInternalServerError).Write(rw)
			return
		}
		pagesCount := int(math.Ceil(float64(totalCount) / float64(perPage)))

		result := dto.SearchResult{
			Items: make([]dto.Template, 0, len(items)),
			Pages: pagesCount,
			Total: totalCount,
		}
		for _, template := range items {
			result.Items = append(result.Items, dto.Template{
				Id:     template.Id,
				Active: template.Active,
				Code:   template.Code,
				Name:   template.Name,
				Body:   template.Body,
				Title:  template.Title,
			})
		}

		responsejson.SuccessResponse(result, 200).Write(rw)
	}
}
