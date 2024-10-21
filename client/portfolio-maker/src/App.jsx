import { useEffect, useState } from 'react'

import './App.css'
import Homepage from './pages/homepage/prelogin/Homepage'
import SignUp from './pages/authentication/SignUp'
import Login from './pages/authentication/Login'
import HomepagePostLogin from './pages/homepage/postlogin/HomepagePostLogin'
import CreateProfilePage from './pages/createProfilePage/CreateProfilePage'
import { ThemeProvider } from '@emotion/react'
import { darkTheme, lightTheme } from './design/inputs/themes.js'
import { useMediaQuery } from '@mui/material'
import { useSelector, useDispatch } from 'react-redux'
import { getUserDetails } from './redux/actions/authActions.js'



function App() {
  const prefersDarkMode = useMediaQuery('(prefers-color-scheme: dark)')
  const { homepage } = useSelector((state) => state.global)
  const dispatch = useDispatch()
  useEffect(() => {
    if (!localStorage.getItem("token")) {
      dispatch(
        {
          type: "SET_HOMEPAGE",
          payload: {
            homepage: 1
          }
        }
      )
    } else {
      dispatch(getUserDetails())
      dispatch(
        {
          type: "SET_HOMEPAGE",
          payload: {
            homepage: 2,
          }
        }
      )
    }
  }, [dispatch])
  return (
    <>
      <ThemeProvider theme={prefersDarkMode ? darkTheme : lightTheme}>
        {
          homepage === 1 ? <Homepage /> : homepage === 2 ? <HomepagePostLogin /> : homepage === 3 ? <CreateProfilePage /> : <></>

        }
        {/* <Homepage/> */}
        <SignUp />
        <Login />
      </ThemeProvider>

    </>
  )
}

export default App
