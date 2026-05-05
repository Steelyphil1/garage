import axios from 'axios';

export interface GarageEvent {
  event_time: string;
  garage_state: string;
}

export const getGarageEvents = async (): Promise<GarageEvent[]> => {
  const response = await axios.get<GarageEvent[]>('https://80c09o5vek.execute-api.us-west-2.amazonaws.com/dev/garage/status');
  return response.data;
};