import React, { useEffect, useState } from "react";
import { FishPenRequest, FishPenResponse } from "./types/fishpen";
import { Heading, Spinner, VStack } from "@chakra-ui/react";
import { Navigate, useParams } from "react-router-dom";
import { FishPenForm } from "./FishPenForm";
import { useDeleteFishPen, useGetFishPen, useUpdateFishPen } from "./apis";
import { FarmLayoutWrapper } from "./FarmLayoutWrapper";
import { LoadingError } from "./LoadingError";

type PageParams = {
  farmId: string;
  fishPenId: string;
};

export const FishPenEditLoaded: React.FC<{
  farmId: string;
  initialFishPen: FishPenResponse;
}> = ({ farmId, initialFishPen }): JSX.Element => {
  const [fishPen, setFishPen] = useState<FishPenRequest>({ ...initialFishPen });

  const { mutate, data, isLoading, isSuccess, isError, error } =
    useUpdateFishPen(farmId, "" + initialFishPen.id);

  const {
    mutate: mutateDelete,
    isLoading: isLoadingDelete,
    isSuccess: isSuccessDelete,
    isError: isErrorDelete,
    error: errorDelete,
  } = useDeleteFishPen(farmId, "" + initialFishPen.id);

  useEffect(() => {
    setFishPen(initialFishPen);
  }, [initialFishPen]);

  useEffect(() => {
    if (isSuccess) {
      setFishPen(data);
    }
  }, [isSuccess, data]);

  const onSubmit = () => {
    if (fishPen.description && fishPen.description.trim() === "") {
      fishPen.description = undefined;
    }
    if (fishPen.netMaterial && fishPen.netMaterial.trim() === "") {
      fishPen.netMaterial = undefined;
    }
    if (fishPen.makerModelName && fishPen.makerModelName.trim() === "") {
      fishPen.makerModelName = undefined;
    }
    if (fishPen.material && fishPen.material.trim() === "") {
      fishPen.material = undefined;
    }

    mutate({ fishPen });
  };

  const onChange = (newFishPen: FishPenRequest) => {
    setFishPen(newFishPen);
  };

  if (isSuccess || isSuccessDelete) {
    return <Navigate to={`/${farmId}`} />;
  }

  return (
    <VStack p={4} spacing={4} alignItems="center">
      <Heading>Edit Fish Pen {`${initialFishPen.id}`}</Heading>
      <FishPenForm
        farmId={farmId}
        fishPen={fishPen}
        onChange={onChange}
        isError={isError}
        error={error}
        onSubmit={onSubmit}
        isLoading={isLoading}
        onDelete={() => {
          mutateDelete();
        }}
        isErrorDelete={isErrorDelete}
        errorDelete={errorDelete}
        isLoadingDelete={isLoadingDelete}
      />
    </VStack>
  );
};

export const FishPenEdit: React.FC = (): JSX.Element => {
  const { farmId, fishPenId } = useParams<keyof PageParams>() as PageParams;

  const { data, isLoading, isError, error, refetch } = useGetFishPen(
    farmId,
    fishPenId
  );

  if (isLoading) {
    return <Spinner />;
  }

  if (isError) {
    return (
      <LoadingError
        resourceName={`fish pen ${fishPenId}`}
        error={error}
        retry={() => {
          refetch();
        }}
      />
    );
  }

  return (
    <FarmLayoutWrapper farmId={farmId}>
      <FishPenEditLoaded farmId={farmId} initialFishPen={data} />;
    </FarmLayoutWrapper>
  );
};
