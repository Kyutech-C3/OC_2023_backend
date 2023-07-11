//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/sivchari/chat-answer" mock_$GOPACKAGE/mock_$GOFILE
package comment

import (
	"context"
	"oc-2023/pkg/domain/entity"

	"github.com/google/uuid"
)

type Repository interface {
	Insert(ctx context.Context, comment *entity.PostComments) error
	Delete(ctx context.Context, comment *entity.DeleteComments) error
	SelectByID(ctx context.Context, ID uuid.UUID) ([]entity.Comment, error)
}
