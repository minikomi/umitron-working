import {
  Alert,
  AlertDescription,
  AlertIcon,
  AlertTitle,
  Button,
  Text,
} from "@chakra-ui/react";

export const LoadingError: React.FC<{
  resourceName: string;
  error: Error;
  retry: () => void;
}> = ({ resourceName, error, retry }): JSX.Element => (
  <Alert status="error">
    <AlertIcon />
    <AlertTitle>Error Loading {resourceName}</AlertTitle>
    <AlertDescription>
      <Text>{error.message}</Text>
      <Button colorScheme="red" onClick={retry}>
        Reload
      </Button>
    </AlertDescription>
  </Alert>
);
