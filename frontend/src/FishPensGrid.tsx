import {
  Box,
  Heading,
  SimpleGrid,
  VStack,
  Text,
  LinkBox,
  LinkOverlay,
} from "@chakra-ui/react";
import { FishPenResponse } from "./types/fishpen";
import React from "react";
import { FishPenCategoryBadge } from "./FishPenCategoryBadge";
import { FishPensTable } from "./FishPensTable";
import { NavLink } from "react-router-dom";

const CharCodeA = "A".charCodeAt(0);

const separateKnownAndUnknownLocation = (
  fishPens: FishPenResponse[]
): {
  locationKnown: FishPenResponse[];
  locationUnknown: FishPenResponse[];
} => {
  // Separate into:
  //   - known location: matches pattern {A-Z}{positive integer}
  //   - unknown location: does not match pattern
  return fishPens.reduce(
    (acc, fishPen) => {
      if (fishPen.name.match(/[A-Z][1-9][0-9]*/)) {
        acc.locationKnown.push(fishPen);
      } else {
        acc.locationUnknown.push(fishPen);
      }
      return acc;
    },
    {
      locationKnown: [] as FishPenResponse[],
      locationUnknown: [] as FishPenResponse[],
    }
  );
};

const getColumnsAndRows = (
  fishPens: FishPenResponse[]
): { rows: string[]; cols: number[] } => {
  // Builds:
  // rows: List of upper case letters from `A` to max existing `[A-Z]` prefix
  // columns: List of numbers from `1` to max existing numeric suffix
  const numberOfRows =
    Math.max(
      ...fishPens.map((fishPen) => fishPen.name.charCodeAt(0) - CharCodeA)
    ) + 1;

  const numberOfColumns = Math.max(
    ...fishPens.map((fishPen) => parseInt(fishPen.name.slice(1), 10))
  );

  const rowArray = Array.from({ length: numberOfRows }, (_, i) =>
    String.fromCharCode(CharCodeA + i)
  );
  const colArray = Array.from({ length: numberOfColumns }, (_, i) => i + 1);

  return { rows: rowArray, cols: colArray };
};

const FishPensGridCell: React.FC<{
  farmId: string;
  row: string;
  column: number;
  fishPen?: FishPenResponse;
}> = ({ farmId, row, column, fishPen }) => (
  <LinkBox
    p={4}
    display="flex"
    flexDirection="column"
    alignItems="center"
    justifyContent="center"
    borderWidth="1px"
    textAlign="center"
    borderRadius="lg"
    bg="gray.50"
    style={{ cursor: "pointer" }}
    _hover={{ border: "1px solid gray.500", bg: "gray.100" }}
  >
    <LinkOverlay
      as={NavLink}
      to={
        fishPen
          ? `/${farmId}/fishpens/${fishPen.id}`
          : `/${farmId}/fishpens/create?name=${row}${column}`
      }
    >
      <Text>{`${row}${column}`}</Text>
      {fishPen && <FishPenCategoryBadge category={fishPen.category} />}
    </LinkOverlay>
  </LinkBox>
);

const FishPensGridKnownLocation: React.FC<{
  farmId: string;
  fishPens: FishPenResponse[];
}> = ({ farmId, fishPens }) => {
  const { rows, cols } = getColumnsAndRows(fishPens);

  return (
    <SimpleGrid columns={cols.length} spacing={2} w={"100%"}>
      {rows.map((row) =>
        cols.map((column) => {
          const fishPen = fishPens.find(
            (fishPen) => fishPen.name === `${row}${column}`
          );
          return (
            <FishPensGridCell
              key={`${row}-${column}`}
              farmId={farmId}
              fishPen={fishPen}
              row={row}
              column={column}
            />
          );
        })
      )}
    </SimpleGrid>
  );
};

export const FishPensGrid: React.FC<{
  farmId: string;
  fishPens: FishPenResponse[];
}> = ({ farmId, fishPens }) => {
  const { locationKnown, locationUnknown } =
    separateKnownAndUnknownLocation(fishPens);

  return (
    <VStack spacing={6}>
      {locationKnown.length > 0 && (
        <VStack spacing={2}>
          <Heading>Location Known</Heading>
          <FishPensGridKnownLocation farmId={farmId} fishPens={locationKnown} />
        </VStack>
      )}
      {locationUnknown.length > 0 && (
        <VStack spacing={2}>
          <Heading>Location Unknown</Heading>
          <FishPensTable farmId={farmId} fishPens={locationUnknown} />
        </VStack>
      )}
    </VStack>
  );
};
