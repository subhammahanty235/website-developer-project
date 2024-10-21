// themes.js
import { createTheme } from '@mui/material/styles';

const lightTheme = createTheme({
  palette: {
    mode: 'light',
  },
  components: {
    MuiTextField: {
      styleOverrides: {
        root: {
          '& .MuiOutlinedInput-root': {
            '& fieldset': {
              borderColor: '#333', // Border color for light mode
            },
            '&:hover fieldset': {
              borderColor: '#333', // Border color on hover
            },
            '&.Mui-focused fieldset': {
              borderColor: '#333', // Border color when focused
            },
          },
          '& .MuiInputLabel-root': {
            color: '#333', // Label color for light mode
          },
          '& .MuiInputBase-input': {
            color: '#333', // Input text color for light mode
          },
        },
      },
    },
    
  },
});

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
  components: {
    MuiTextField: {
      styleOverrides: {
        root: {
          '& .MuiOutlinedInput-root': {
            '& fieldset': {
              borderColor: 'white', // Border color for dark mode
            },
            '&:hover fieldset': {
              borderColor: 'white', // Border color on hover
            },
            '&.Mui-focused fieldset': {
              borderColor: 'white', // Border color when focused
            },
          },
          '& .MuiInputLabel-root': {
            color: 'white', // Label color for dark mode
          },
          '& .MuiInputBase-input': {
            color: 'white', // Input text color for dark mode
          },
        },
      },
    },
    
  },
});

export { lightTheme, darkTheme };
