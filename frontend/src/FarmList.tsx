import { Box, Flex, Heading, Link, Spinner } from "@chakra-ui/react";
import { NavLink } from "react-router-dom";

import { useGetFarms } from "./apis";
import { LoadingError } from "./LoadingError";

export const FarmList: React.FC = (): JSX.Element => {
  const { data: farms, isLoading, isError, error, refetch } = useGetFarms();

  if (isLoading) {
    return <Spinner />;
  }

  if (isError) {
    return (
      <LoadingError
        resourceName="Farms"
        error={error}
        retry={() => refetch()}
      />
    );
  }

  return (
    <Flex direction="column" alignItems="center">
      <Box mt={6}>
        <Heading>Farms</Heading>
      </Box>
      <Box mt={2}>
        {(farms || []).map((farm) => (
          <Box
            key={farm.id}
            mt={4}
            p={0}
            display="flex"
            flexDirection="column"
            w="400px"
          >
            <Link as={NavLink} to={`/${farm.id}`} p={4}>
              {farm.name}
            </Link>
          </Box>
        ))}
      </Box>
    </Flex>
  );
};
