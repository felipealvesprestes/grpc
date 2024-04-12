package service

import (
	"context"

	"github.com/felipealvesprestes/grpc/internal/database"
	"github.com/felipealvesprestes/grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	categoty, err := c.CategoryDB.Create(in.Name, in.Description)

	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          categoty.Id,
		Name:        categoty.Name,
		Description: categoty.Description,
	}

	return &pb.CategoryResponse{Category: categoryResponse}, nil
}
