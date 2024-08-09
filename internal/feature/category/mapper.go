package featureCategory

import (
	"github.com/Sanchir01/microservice_sandjma_category/internal/modain/models"
	sandjmav1 "github.com/Sanchir01/protos_files_job/pkg/gen/golang/category"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapCategoryToGRPCModel(categories []models.Category) (item []*sandjmav1.Category, err error) {
	categoriesChan := make(chan *sandjmav1.Category, len(categories))
	var categoryPtrs []*sandjmav1.Category

	go func() {
		for i := range categories {
			newCategory := &sandjmav1.Category{
				Id:          categories[i].ID.String(),
				Name:        categories[i].Name,
				Slug:        categories[i].Slug,
				Description: categories[i].Description,
				Version:     int64(categories[i].Version),
				UpdatedAt:   timestamppb.New(categories[i].UpdatedAt),
				CreatedAt:   timestamppb.New(categories[i].CreatedAt),
			}
			categoriesChan <- newCategory
		}
		close(categoriesChan)
	}()

	for category := range categoriesChan {
		categoryPtrs = append(categoryPtrs, category)
	}
	return categoryPtrs, nil

}
