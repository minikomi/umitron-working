import { Box, Flex, Heading, Link } from '@chakra-ui/react';
import React, { useEffect, useState } from 'react';
import { NavLink } from 'react-router-dom';

import { fetchFarmsApi } from './apis';
import { Farm } from './types/farm';

export const FarmList: React.FC = (): JSX.Element => {
  const [farms, setFarms] = useState<Farm[]>([]);

  useEffect(() => {
    fetchFarmsApi().then((farms) => {
      setFarms(farms);
    }).catch((e) => {
      console.error(e.message);
    });
  }, []);

  return (
    <Flex
      direction="column"
      alignItems="center">
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
            flexDirection="column" w="400px">
            <Link as={NavLink} to={`/${farm.id}`} p={4}>{farm.name}</Link>
          </Box>
        ))}
      </Box>
    </Flex>
  );
};
