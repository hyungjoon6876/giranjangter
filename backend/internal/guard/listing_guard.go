package guard

import (
	"fmt"

	"github.com/jym/lincle/internal/domain"
)

// AllowedListingTransitions defines valid state transitions for listings.
var AllowedListingTransitions = map[domain.ListingStatus][]domain.ListingStatus{
	domain.ListingAvailable:    {domain.ListingReserved, domain.ListingCancelled},
	domain.ListingReserved:     {domain.ListingPendingTrade, domain.ListingAvailable, domain.ListingCancelled},
	domain.ListingPendingTrade: {domain.ListingCompleted, domain.ListingAvailable, domain.ListingCancelled},
	// completed and cancelled are terminal states
}

// ValidateListingTransition checks if a listing state transition is allowed.
func ValidateListingTransition(from, to domain.ListingStatus) error {
	allowed, ok := AllowedListingTransitions[from]
	if !ok {
		return fmt.Errorf("INVALID_TRANSITION: no transitions allowed from %q", from)
	}
	for _, s := range allowed {
		if s == to {
			return nil
		}
	}
	return fmt.Errorf("INVALID_TRANSITION: %q -> %q is not allowed", from, to)
}

// AllowedReservationTransitions defines valid state transitions for reservations.
var AllowedReservationTransitions = map[domain.ReservationStatus][]domain.ReservationStatus{
	domain.ReservationProposed:  {domain.ReservationConfirmed, domain.ReservationCancelled, domain.ReservationExpired},
	domain.ReservationConfirmed: {domain.ReservationFulfilled, domain.ReservationCancelled, domain.ReservationNoShowReported},
	// expired, cancelled, fulfilled, no_show_reported are terminal
}

// ValidateReservationTransition checks if a reservation state transition is allowed.
func ValidateReservationTransition(from, to domain.ReservationStatus) error {
	allowed, ok := AllowedReservationTransitions[from]
	if !ok {
		return fmt.Errorf("INVALID_TRANSITION: no transitions allowed from %q", from)
	}
	for _, s := range allowed {
		if s == to {
			return nil
		}
	}
	return fmt.Errorf("INVALID_TRANSITION: reservation %q -> %q is not allowed", from, to)
}
