import React, { useState } from "react";
import {
  Box,
  Button,
  HStack,
  Heading,
  Spinner,
  Text,
  VStack,
} from "@chakra-ui/react";
import { useGetFishPensForFarm } from "./apis";
import { LoadingError } from "./LoadingError";
import { useParams, NavLink } from "react-router-dom";
import { DragHandleIcon, HamburgerIcon } from "@chakra-ui/icons";
import { FishPensTable } from "./FishPensTable";
import { FishPensGrid } from "./FishPensGrid";
import { FarmLayoutWrapper } from "./FarmLayoutWrapper";

type PageParams = { farmId: string };

export const FarmHome: React.FC = (): JSX.Element => {
  const { farmId } = useParams<keyof PageParams>() as PageParams;

  const [gridView, setGridView] = useState<boolean>(false);

  const {
    data: fishPens,
    isLoading: fishPensIsLoading,
    isError: fishPensIsError,
    error: fishPensError,
    refetch: fishPensRefetch,
  } = useGetFishPensForFarm(farmId);

  if (fishPensIsLoading) {
    return <Spinner />;
  }

  if (fishPensIsError) {
    return (
      <LoadingError
        resourceName="Fish Pens"
        error={fishPensError}
        retry={fishPensRefetch}
      />
    );
  }

  return (
    <FarmLayoutWrapper farmId={farmId}>
      <Box mt={2} w="100%" p={6}>
        <VStack spacing={4} align="center">
          <Heading>Fish Pens</Heading>
          <Button
            as={NavLink}
            to={`/${farmId}/fishpens/create`}
            colorScheme="blue"
          >
            Add New Fish Pen
          </Button>
          {fishPens.length === 0 ? (
            <Text>No fish pens yet.</Text>
          ) : (
            <>
              <HStack mt={4} spacing={2} alignItems="center">
                <Text>View:</Text>
                <Button
                  onClick={() => setGridView(false)}
                  isActive={!gridView}
                  variant="link"
                  leftIcon={<HamburgerIcon />}
                >
                  Table
                </Button>
                <Button
                  onClick={() => setGridView(true)}
                  isActive={gridView}
                  variant="link"
                  leftIcon={<DragHandleIcon />}
                >
                  Grid
                </Button>
              </HStack>
              <Box mt={2}>
                {gridView ? (
                  <FishPensGrid farmId={farmId} fishPens={fishPens} />
                ) : (
                  <FishPensTable farmId={farmId} fishPens={fishPens} />
                )}
              </Box>
            </>
          )}
        </VStack>
      </Box>
    </FarmLayoutWrapper>
  );
};
