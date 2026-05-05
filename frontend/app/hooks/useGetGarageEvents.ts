'use client';

import { useState, useEffect } from 'react';
import { getGarageEvents, GarageEvent } from '../api/garageApi';
import { mockGarageEvents } from "../api/mockData";

export const useGetGarageEvents = () => {
  const [events, setEvents] = useState<GarageEvent[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetch = async () => {
      try {
        const data = await getGarageEvents();
        setEvents(data);
      } catch (err) {
        setError('Failed to fetch garage events');
      } finally {
        setLoading(false);
      }
    };

    fetch();
  }, []);

  return { events, loading, error };
};
