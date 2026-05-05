'use client';

import { Box, Typography } from '@mui/material';

export default function Footer() {
  return (
    <Box component="footer" sx={{ py: 2, textAlign: 'center', mt: 'auto' }}>
      <Typography variant="body2" suppressHydrationWarning>
        © {new Date().getFullYear()} Phillips Garage
      </Typography>
    </Box>
  );
}