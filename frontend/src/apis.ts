import { UseQueryResult, useQuery, useMutation } from "@tanstack/react-query";
import { FarmResponse } from "./types/farm";
import { FishPenRequest, FishPenResponse } from "./types/fishpen";

const baseHeaders = {
  Accept: "application/json",
  "Content-Type": "application/json",
};

const parseResponse = async (response: Response) => {
  const data = await response.json();
  if (!response.ok) {
    throw new Error(data.error);
  }
  return data;
};

export const useGetFarms = (): UseQueryResult<FarmResponse[], Error> => {
  return useQuery<FarmResponse[], Error>(["farms"], async () => {
    const response = await fetch(`/api/farms`, {
      method: "GET",
      headers: baseHeaders,
    });
    return parseResponse(response);
  });
};

export const useGetFarmById = (
  farmId: string 
): UseQueryResult<FarmResponse, Error> => {
  return useQuery<FarmResponse, Error>(
    ["farms", farmId],
    async () => {
      const response = await fetch(`/api/farms/${farmId}`, {
        method: "GET",
        headers: baseHeaders,
      });
      return parseResponse(response);
    },
  );
};

export function useCreateFishPen(farmId: string) {
  return useMutation<FishPenResponse, Error, {fishPen: FishPenRequest}>(
    async ({fishPen}) => {
      const response = await fetch(`/api/farms/${farmId}/fishpens`, {
        method: "POST",
        headers: baseHeaders,
        body: JSON.stringify(fishPen),
      });
      return parseResponse(response);
    }
  );
}

export function useGetFishPensForFarm(farmId: string) {
  return useQuery<FishPenResponse[], Error>(["fishPens", farmId], async () => {
    const response = await fetch(`/api/farms/${farmId}/fishpens`, {
      method: "GET",
      headers: baseHeaders,
    });
    return parseResponse(response);
  });
}

export function useGetFishPen(farmId: string, fishPenId: string) {
  return useQuery<FishPenResponse, Error>(["fishPens", farmId, fishPenId], async () => {
    const response = await fetch(`/api/farms/${farmId}/fishpens/${fishPenId}`, {
      method: "GET",
      headers: baseHeaders,
    });
    return parseResponse(response);
  });
}

export function useUpdateFishPen(farmId: string, fishPenId: string) {
  return useMutation<FishPenResponse, Error, {fishPen: FishPenRequest}>(
    async ({fishPen}) => {
      const response = await fetch(`/api/farms/${farmId}/fishpens/${fishPenId}`, {
        method: "PUT",
        headers: baseHeaders,
        body: JSON.stringify(fishPen),
      });
      return parseResponse(response);
    }
  );
}

export function useDeleteFishPen(farmId: string, fishPenId: string) {
  return useMutation<boolean, Error>(
    async () => {
      const response = await fetch(`/api/farms/${farmId}/fishpens/${fishPenId}`, {
        method: "DELETE",
        headers: baseHeaders,
      });
      return response.ok;
    }
  );
}
