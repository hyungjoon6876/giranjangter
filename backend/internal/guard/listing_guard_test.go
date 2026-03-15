package guard

import (
	"testing"

	"github.com/jym/lincle/internal/domain"
)

func TestValidateListingTransition_Allowed(t *testing.T) {
	tests := []struct {
		from domain.ListingStatus
		to   domain.ListingStatus
	}{
		{domain.ListingAvailable, domain.ListingReserved},
		{domain.ListingAvailable, domain.ListingCancelled},
		{domain.ListingReserved, domain.ListingPendingTrade},
		{domain.ListingReserved, domain.ListingAvailable},
		{domain.ListingReserved, domain.ListingCancelled},
		{domain.ListingPendingTrade, domain.ListingCompleted},
		{domain.ListingPendingTrade, domain.ListingAvailable},
		{domain.ListingPendingTrade, domain.ListingCancelled},
	}
	for _, tt := range tests {
		if err := ValidateListingTransition(tt.from, tt.to); err != nil {
			t.Errorf("expected %s -> %s to be allowed, got error: %v", tt.from, tt.to, err)
		}
	}
}

func TestValidateListingTransition_Forbidden(t *testing.T) {
	tests := []struct {
		from domain.ListingStatus
		to   domain.ListingStatus
	}{
		{domain.ListingCompleted, domain.ListingAvailable},
		{domain.ListingCompleted, domain.ListingReserved},
		{domain.ListingCancelled, domain.ListingCompleted},
		{domain.ListingCancelled, domain.ListingAvailable},
		{domain.ListingAvailable, domain.ListingCompleted},
		{domain.ListingAvailable, domain.ListingPendingTrade},
	}
	for _, tt := range tests {
		if err := ValidateListingTransition(tt.from, tt.to); err == nil {
			t.Errorf("expected %s -> %s to be forbidden, but got nil error", tt.from, tt.to)
		}
	}
}

func TestValidateReservationTransition_Allowed(t *testing.T) {
	tests := []struct {
		from domain.ReservationStatus
		to   domain.ReservationStatus
	}{
		{domain.ReservationProposed, domain.ReservationConfirmed},
		{domain.ReservationProposed, domain.ReservationCancelled},
		{domain.ReservationProposed, domain.ReservationExpired},
		{domain.ReservationConfirmed, domain.ReservationFulfilled},
		{domain.ReservationConfirmed, domain.ReservationCancelled},
		{domain.ReservationConfirmed, domain.ReservationNoShowReported},
	}
	for _, tt := range tests {
		if err := ValidateReservationTransition(tt.from, tt.to); err != nil {
			t.Errorf("expected %s -> %s to be allowed, got error: %v", tt.from, tt.to, err)
		}
	}
}

func TestValidateReservationTransition_Forbidden(t *testing.T) {
	tests := []struct {
		from domain.ReservationStatus
		to   domain.ReservationStatus
	}{
		{domain.ReservationExpired, domain.ReservationConfirmed},
		{domain.ReservationCancelled, domain.ReservationConfirmed},
		{domain.ReservationFulfilled, domain.ReservationCancelled},
		{domain.ReservationProposed, domain.ReservationFulfilled},
	}
	for _, tt := range tests {
		if err := ValidateReservationTransition(tt.from, tt.to); err == nil {
			t.Errorf("expected %s -> %s to be forbidden, but got nil error", tt.from, tt.to)
		}
	}
}
