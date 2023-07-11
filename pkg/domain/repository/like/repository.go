//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/sivchari/chat-answer" mock_$GOPACKAGE/mock_$GOFILE
package like

import (
	"context"
	"oc-2023/pkg/domain/entity"

	"github.com/google/uuid"
)

type Repository interface {
	Insert(ctx context.Context, like *entity.PostLikes) error
	Delete(ctx context.Context, like *entity.DeleteLikes) error
	SelectByID(ctx context.Context, ID uuid.UUID) ([]entity.Like, error)
}
