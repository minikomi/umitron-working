import { Farm } from './types/farm';

const baseHeaders = {
  'Accept': 'application/json',
  'Content-Type': 'application/json',
};

const parseResponse = async (response: Response) => {
  const data = await response.json();
  if (!response.ok) {
    throw new Error(data.error);
  }
  return data;
};

export const fetchFarmsApi = async (): Promise<Farm[]> => {
  const response = await fetch(`/api/farms`, {
    method: 'GET',
    headers: baseHeaders,
  });
  return parseResponse(response);
};
