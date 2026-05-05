'use client';

import { Box, Typography, CircularProgress } from '@mui/material';
import { useGetGarageEvents } from './hooks/useGetGarageEvents';
import GarageEventTable from './components/GarageEventTable';

export default function Home() {
  const { events, loading, error } = useGetGarageEvents();

  const latestState = events?.[0]?.garage_state ?? 'Unknown';

  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100%',
        px: 2,
      }}
    >
      {loading ? (
        <CircularProgress />
      ) : error ? (
        <Typography color="error">{error}</Typography>
      ) : (
        <>
          <Typography variant="h4" sx={{ mb: 4 }}>
            The garage is&nbsp;
            <Typography
              component="span"
              variant="h4"
              sx={{ fontWeight: "bold", color: latestState === 'Open' ? 'red' : 'green'}}
            >
              {latestState}
            </Typography>
          </Typography>
          <GarageEventTable events={events} />
        </>
      )}
    </Box>
  );
}