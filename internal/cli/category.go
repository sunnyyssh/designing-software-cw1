package cli

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/sunnyyssh/designing-software-cw1/internal/application/dto"
)

type CategoryService interface {
	Get(ctx context.Context, id uuid.UUID) (*dto.CategoryDTO, error)
	Create(ctx context.Context, typ string, name string) (*dto.CategoryDTO, error)
	List(ctx context.Context) ([]dto.CategoryDTO, error)
	Delete(ctx context.Context, id uuid.UUID) (*dto.CategoryDTO, error)
}

func Category(svc CategoryService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "category",
		Short: "Operations connected to categories",
	}
	cmd.AddCommand(
		getCategory(svc),
		createCategory(svc),
		listCategories(svc),
		deleteCategory(svc),
	)
	return cmd
}

func getCategory(svc CategoryService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a category by its ID",
	}

	var categoryIDStr string
	cmd.Flags().StringVarP(&categoryIDStr, "id", "i", "", "Category ID")
	cmd.MarkFlagRequired("id")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		categoryID, err := uuid.Parse(categoryIDStr)
		if err != nil {
			return fmt.Errorf("invalid category ID: %w", err)
		}

		category, err := svc.Get(cmd.Context(), categoryID)
		if err != nil {
			return fmt.Errorf("failed to get category: %w", err)
		}

		cmd.Println("Category details:")
		PrettyJSON(cmd, category)
		return nil
	}

	return cmd
}

func createCategory(svc CategoryService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
	}

	var (
		typ  string
		name string
	)
	cmd.Flags().StringVarP(&typ, "type", "t", "", "Category type (income/outcome)")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Category name")
	cmd.MarkFlagRequired("type")
	cmd.MarkFlagRequired("name")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		category, err := svc.Create(cmd.Context(), typ, name)
		if err != nil {
			return fmt.Errorf("failed to create category: %w", err)
		}

		cmd.Println("Created category:")
		PrettyJSON(cmd, category)
		return nil
	}

	return cmd
}

func listCategories(svc CategoryService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all categories",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		categories, err := svc.List(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to list categories: %w", err)
		}

		cmd.Println("Categories:")
		PrettyJSON(cmd, categories)
		return nil
	}

	return cmd
}

func deleteCategory(svc CategoryService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a category by its ID",
	}

	var categoryIDStr string
	cmd.Flags().StringVarP(&categoryIDStr, "id", "i", "", "Category ID")
	cmd.MarkFlagRequired("id")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		categoryID, err := uuid.Parse(categoryIDStr)
		if err != nil {
			return fmt.Errorf("invalid category ID: %w", err)
		}

		category, err := svc.Delete(cmd.Context(), categoryID)
		if err != nil {
			return fmt.Errorf("failed to delete category: %w", err)
		}

		cmd.Println("Deleted category:")
		PrettyJSON(cmd, category)
		return nil
	}

	return cmd
}
