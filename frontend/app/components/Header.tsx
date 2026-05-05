'use client';

import { AppBar, Toolbar, Typography } from '@mui/material';

export default function Header() {
  return (
    <AppBar position="static">
      <Toolbar sx={{ backgroundColor: "black"}}>
        <Typography variant="h6" component="h1">
          Phillips Garage
        </Typography>
      </Toolbar>
    </AppBar>
  );
}