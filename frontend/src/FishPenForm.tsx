import React from "react";
import { FishPenCategory, FishPenRequest } from "./types/fishpen";
import { NavLink } from "react-router-dom";
import {
  Alert,
  AlertDescription,
  AlertIcon,
  AlertTitle,
  Box,
  Button,
  FormControl,
  FormLabel,
  HStack,
  Input,
  InputGroup,
  InputLeftAddon,
  NumberInput,
  NumberInputField,
  Radio,
  RadioGroup,
  Spinner,
  Textarea,
  VStack,
} from "@chakra-ui/react";
import { FishPenCategoryBadge } from "./FishPenCategoryBadge";

export const FishPenForm: React.FC<{
  farmId: string;
  fishPen: FishPenRequest;
  onChange: (fishPen: FishPenRequest) => void;
  onSubmit: (fishPen: FishPenRequest) => void;
  isLoading: boolean;
  isError: boolean;
  error: Error | null;
  onDelete?: () => void;
  isErrorDelete?: boolean;
  errorDelete?: Error | null;
  isLoadingDelete?: boolean;
}> = ({
  farmId,
  fishPen,
  onChange,
  onSubmit,
  isLoading,
  isError,
  error,
  onDelete,
  isErrorDelete,
  errorDelete,
  isLoadingDelete,
}) => {
  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    onSubmit(fishPen);
  };

  const handleInputChange = (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    onChange({
      ...fishPen,
      [event.target.name]: event.target.value,
    });
  };

  const handleNumberInputChange = (name: string) => {
    return (_valueAsString: string, valueAsNumber: number) =>
      onChange({
        ...fishPen,
        [name]: valueAsNumber,
      });
  };

  const handleCategoryChange = (category: FishPenCategory) => {
    onChange({
      ...fishPen,
      category,
    });
  };

  if (isLoading || isLoadingDelete) {
    return <Spinner />;
  }

  return (
    <Box maxW="md" w="100%">
      {isError && (
        <Alert status="error">
          <AlertIcon />
          <AlertTitle mr={2}>Error!</AlertTitle>
          <AlertDescription>{error && error.message}</AlertDescription>
        </Alert>
      )}
      {isErrorDelete && (
        <Alert status="error">
          <AlertIcon />
          <AlertTitle mr={2}>Error Deleting!</AlertTitle>
          <AlertDescription>
            {errorDelete && errorDelete.message}
          </AlertDescription>
        </Alert>
      )}

      <form onSubmit={handleSubmit}>
        <FormControl>
          <FormLabel>Name</FormLabel>
          <Input
            required
            name="name"
            value={fishPen.name}
            onChange={handleInputChange}
          />
        </FormControl>

        <FormControl>
          <FormLabel>Maker Model Name</FormLabel>
          <Input
            name="makerModelName"
            value={fishPen.makerModelName}
            onChange={handleInputChange}
          />
        </FormControl>

        <FormControl>
          <FormLabel>Description</FormLabel>
          <Textarea
            name="description"
            value={fishPen.description}
            onChange={handleInputChange}
          />
        </FormControl>

        <FormControl>
          <FormLabel>Material</FormLabel>
          <Input
            name="material"
            value={fishPen.material}
            onChange={handleInputChange}
          />
        </FormControl>

        <FormControl>
          <FormLabel>Net Material</FormLabel>
          <Input
            name="netMaterial"
            value={fishPen.netMaterial}
            onChange={handleInputChange}
          />
        </FormControl>

        <FormControl>
          <FormLabel>Category</FormLabel>
          <RadioGroup
            name="category"
            value={fishPen.category}
            onChange={(nextValue: string) =>
              handleCategoryChange(nextValue as FishPenCategory)
            }
          >
            <VStack alignItems="flex-start">
              {Object.values(FishPenCategory).map((category) => (
                <Radio value={category}>
                  <FishPenCategoryBadge category={category} />
                </Radio>
              ))}
            </VStack>
          </RadioGroup>
        </FormControl>

        <FormControl>
          <FormLabel>Size (cm)</FormLabel>
          <HStack>
            <InputGroup size="sm">
              <InputLeftAddon children="w" />
              <NumberInput
                min={0}
                value={fishPen.widthCm}
                onChange={handleNumberInputChange("widthCm")}
              >
                <NumberInputField />
              </NumberInput>
            </InputGroup>

            <InputGroup size="sm">
              <InputLeftAddon children="l" />
              <NumberInput
                precision={2}
                min={0}
                value={fishPen.lengthCm}
                onChange={handleNumberInputChange("lengthCm")}
              >
                <NumberInputField />
              </NumberInput>
            </InputGroup>

            <InputGroup size="sm">
              <InputLeftAddon children="h" />
              <NumberInput
                precision={2}
                min={0}
                value={fishPen.heightCm}
                onChange={handleNumberInputChange("heightCm")}
              >
                <NumberInputField />
              </NumberInput>
            </InputGroup>
          </HStack>
        </FormControl>

        <Box mt={4}>
          <Button colorScheme="blue" type="submit">
            Submit
          </Button>
          <Button as={NavLink} ml={4} to={`/${farmId}`}>
            Cancel
          </Button>
          {onDelete && (
            <Button ml={4} onClick={onDelete} colorScheme="red">
              Delete
            </Button>
          )}
        </Box>
      </form>
    </Box>
  );
};
