import React from "react";
import {
  Alert,
  AlertDescription,
  AlertIcon,
  AlertTitle,
  Heading,
  Spinner,
  VStack,
} from "@chakra-ui/react";
import { useGetFarmById } from "./apis";
import { LoadingError } from "./LoadingError";
import { Link } from "react-router-dom";



export const FarmLayoutWrapper: React.FC<{
  farmId: string;
  children?: React.ReactNode;
}> = ({ farmId, children }) => {
  const {
    data: farm,
    isLoading: farmLoading,
    isError: farmIsError,
    error: farmError,
    refetch: farmRefetch,
  } = useGetFarmById(farmId);

  if (!farmId) {
    return (
      <Alert status="error">
        <AlertIcon />
        <AlertTitle>Error Loading Farm</AlertTitle>
        <AlertDescription>Invalid farm ID</AlertDescription>
      </Alert>
    );
  }

  if (farmLoading) {
    return <Spinner />;
  }

  if (farmIsError) {
    return (
      <LoadingError
        resourceName="Farm"
        error={farmError}
        retry={() => {
          farmRefetch();
        }}
      />
    );
  }

  return (
    <VStack p={4} spacing={4} align="stretch">
      <Heading>
        <Link to={`/${farm.id}`}>{farm.name}</Link>
      </Heading>
      {children}
    </VStack>
  );
};
