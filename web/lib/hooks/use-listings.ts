import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import type { Listing } from "@/lib/types";

export function useListings(params?: {
  serverId?: string;
  categoryId?: string;
  q?: string;
  listingType?: string;
  sort?: string;
}) {
  return useQuery({
    queryKey: ["listings", params],
    queryFn: () => apiClient.getListings(params),
  });
}

export function useListing(id: string) {
  return useQuery({
    queryKey: ["listing", id],
    queryFn: () => apiClient.getListing(id),
    enabled: !!id,
  });
}

export function useCreateListing() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: Record<string, unknown>) =>
      apiClient.createListing(data as Partial<Listing>),
    onSuccess: () => qc.invalidateQueries({ queryKey: ["listings"] }),
  });
}

export function useUpdateListing() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: Partial<Listing> }) =>
      apiClient.updateListing(id, data),
    onSuccess: (_, { id }) => {
      qc.invalidateQueries({ queryKey: ["listing", id] });
      qc.invalidateQueries({ queryKey: ["listings"] });
      qc.invalidateQueries({ queryKey: ["my-listings"] });
    },
  });
}

export function useChangeListingStatus() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, status }: { id: string; status: string }) =>
      apiClient.changeListingStatus(id, status),
    onSuccess: (_, { id }) => {
      qc.invalidateQueries({ queryKey: ["listing", id] });
      qc.invalidateQueries({ queryKey: ["listings"] });
      qc.invalidateQueries({ queryKey: ["my-listings"] });
    },
  });
}

export function useToggleFavorite() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, isFavorited }: { id: string; isFavorited: boolean }) =>
      isFavorited
        ? apiClient.unfavoriteListing(id)
        : apiClient.favoriteListing(id),
    onSuccess: (_, { id }) => {
      qc.invalidateQueries({ queryKey: ["listing", id] });
      qc.invalidateQueries({ queryKey: ["listings"] });
    },
  });
}
