import { createTheme } from '@mui/material/styles';

const theme = createTheme({
  palette: {
    mode: 'dark', // or 'light', or detect from system
    background: {
      default: '#0a0a0a',
    },
  },
  typography: {
    fontFamily: 'Arial, Helvetica, sans-serif',
  },
});
export default theme;