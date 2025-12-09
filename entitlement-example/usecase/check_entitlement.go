package usecase

import (
	"entitlement-example/domain"
	"fmt"
)

// CheckEntitlementUsecase represents the application specific business rules.
type CheckEntitlementUsecase struct {
	repo domain.EntitlementRepository
}

func NewCheckEntitlementUsecase(repo domain.EntitlementRepository) *CheckEntitlementUsecase {
	return &CheckEntitlementUsecase{
		repo: repo,
	}
}

// VerifyUserAccess checks if a user is allowed to access a resource protected by a group.
func (u *CheckEntitlementUsecase) VerifyUserAccess(group string, user string) bool {
	fmt.Printf("\n--- Checking Access: User[%s] -> Group[%s] ---\n", user, group)

	// Usecase simply asks the repository. It doesn't know about cache or AD.
	isMember, err := u.repo.IsMember(group, user)

	if err != nil {
		fmt.Printf("Error checking membership: %v\n", err)
		return false
	}

	if isMember {
		fmt.Printf("Result: ALLOWED (User is a member)\n")
	} else {
		fmt.Printf("Result: DENIED (User is NOT a member)\n")
	}

	return isMember
}
