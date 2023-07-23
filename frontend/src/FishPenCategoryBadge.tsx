import { Badge } from "@chakra-ui/react";
import { FishPenCategory } from "./types/fishpen";

export const FishPenCategoryBadge: React.FC<{ category: FishPenCategory }> = ({
  category,
}): JSX.Element => {
  const colorScheme = {
    [FishPenCategory.fixed]: "green",
    [FishPenCategory.floating]: "blue",
    [FishPenCategory.submersible]: "orange",
    [FishPenCategory.submersed]: "teal",
    [FishPenCategory.other]: "gray",
  }[category];

  return <Badge colorScheme={colorScheme}>{category}</Badge>;
};
