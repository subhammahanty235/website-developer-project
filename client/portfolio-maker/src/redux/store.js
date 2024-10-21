import {configureStore} from '@reduxjs/toolkit'
import { authReducer } from './reducers/authReducer';
import {globalReducer} from './reducers/globalReducer'
import { profileReducer } from './reducers/profileReducer';

const store = configureStore({
    reducer:{
        authReducer:authReducer,
        global:globalReducer,
        profileReducer: profileReducer
    }
})


export default store;