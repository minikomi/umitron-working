import {
  Box,
  Heading,
} from "@chakra-ui/react";
import { FarmResponse } from "./types/farm";

const FarmInfo: React.FC<{ farm: FarmResponse }> = ({ farm }): JSX.Element => {
  return (
    <Box mt={6}>
      <Heading>{farm.name}</Heading>
    </Box>
  );
};
