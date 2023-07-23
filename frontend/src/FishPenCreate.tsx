import React, { useState } from "react";
import { FishPenCategory, FishPenRequest } from "./types/fishpen";
import { Heading, VStack } from "@chakra-ui/react";
import { Navigate, useParams } from "react-router-dom";
import { FishPenForm } from "./FishPenForm";
import { useCreateFishPen } from "./apis";
import { FarmLayoutWrapper } from "./FarmLayoutWrapper";

type PageParams = {
  farmId: string;
};

export const FishPenCreate: React.FC = (): JSX.Element => {
  const { farmId } = useParams<keyof PageParams>() as PageParams;
  // from query param "name" else empty
  const initialName =
    new URLSearchParams(window.location.search).get("name") ?? "";

  const [fishPen, setFishPen] = useState<FishPenRequest>({
    name: initialName,
    makerModelName: "",
    description: "",
    material: "",
    netMaterial: "",
    category: FishPenCategory.fixed,
    widthCm: 0,
    lengthCm: 0,
    heightCm: 0,
  });

  const { mutate, isLoading, isSuccess, isError, error } =
    useCreateFishPen(farmId);

  const onSubmit = () => {
    // TODO: prevent this by using zod as parser within form.
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

  const onChange = (fishPen: FishPenRequest) => {
    setFishPen(fishPen);
  };

  if (isSuccess) {
    return <Navigate to={`/${farmId}`} />;
  }

  return (
    <FarmLayoutWrapper farmId={farmId}>
      <VStack p={4} spacing={4} alignItems="center">
        <Heading>Add New Fish Pen</Heading>
        <FishPenForm
          farmId={farmId}
          fishPen={fishPen}
          onChange={onChange}
          isError={isError}
          error={error}
          onSubmit={onSubmit}
          isLoading={isLoading}
        />
      </VStack>
    </FarmLayoutWrapper>
  );
};
