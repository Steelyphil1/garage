'use client';

import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Typography,
  useMediaQuery,
  useTheme,
  Box,
} from '@mui/material';
import { GarageEvent } from '../api/garageApi';

interface Props {
  events: GarageEvent[];
}

const formatLocalTime = (utcString: string): string => {
  const date = new Date(utcString);
  return date.toLocaleString(undefined, {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
  });
};

export default function GarageEventTable({ events }: Props) {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));

  if (!events.length) return null;

  if (isMobile) {
    return (
      <Box sx={{ width: '100%' }}>
        {events.map((event, index) => (
          <Box
            key={index}
            sx={{
              width: '100%',
              p: 2,
              borderBottom: '1px solid',
              borderColor: 'divider',
              display: 'flex',
              justifyContent: 'space-between',
              alignItems: 'center',
            }}
          >
            <Typography variant="body2" color="text.secondary">
              {formatLocalTime(event.event_time)}
            </Typography>
            <Typography
              variant="body1"
              sx={{ fontWeight: "bold"}}
              color={event.garage_state === 'Open' ? 'success.main' : 'error.main'}
            >
              {event.garage_state}
            </Typography>
          </Box>
        ))}
      </Box>
    );
  }

  return (
    <TableContainer component={Paper} sx={{ width: '100%', mt: 4, mb: 10, maxWidth: 800 }}>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell><strong>Time</strong></TableCell>
            <TableCell><strong>Status</strong></TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {events.map((event, index) => (
            <TableRow key={index}>
              <TableCell>{formatLocalTime(event.event_time)}</TableCell>
              <TableCell>
                <Typography
                  color={event.garage_state === 'Open' ? 'success.main' : 'error.main'}
                  sx={{ fontWeight: "bold"}}
                >
                  {event.garage_state}
                </Typography>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}