import {
  Heading,
  LinkBox,
  LinkOverlay,
  Table,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
  VStack,
} from "@chakra-ui/react";
import { FishPenCategoryBadge } from "./FishPenCategoryBadge";
import { FishPenResponse } from "./types/fishpen";
import { NavLink } from "react-router-dom";

export const FishPensTable: React.FC<{
  farmId: string;
  fishPens: FishPenResponse[];
}> = ({ farmId, fishPens }): JSX.Element => {
  return (
    <Table variant="simple">
      <Thead>
        <Tr>
          <Th>Name</Th>
          <Th>Material</Th>
          <Th>Net Material</Th>
          <Th>Category</Th>
          <Th>Size (w x l x h) cm</Th>
          <Th>
            Volume (m<sup>3</sup>)
          </Th>
        </Tr>
      </Thead>
      <Tbody>
        {fishPens.map((fishPen) => (
          <LinkBox
            as={Tr}
            key={fishPen.id}
            style={{ cursor: "pointer" }}
            _hover={{ bg: "gray.100" }}
          >
            <Td>
              <LinkOverlay
                as={NavLink}
                to={`/${farmId}/fishpens/${fishPen.id}`}
              >
                {fishPen.name}
              </LinkOverlay>
            </Td>
            <Td>{fishPen.material}</Td>
            <Td>{fishPen.netMaterial}</Td>
            <Td>
              <FishPenCategoryBadge category={fishPen.category} />
            </Td>
            <Td>{`${fishPen.widthCm} x ${fishPen.lengthCm} x ${fishPen.heightCm}`}</Td>
            <Td>
              {Math.round(
                (fishPen.widthCm * fishPen.lengthCm * fishPen.heightCm) / 10000
              ) / 100}
            </Td>
          </LinkBox>
        ))}
      </Tbody>
    </Table>
  );
};
