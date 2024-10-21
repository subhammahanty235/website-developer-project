import { createReducer } from "@reduxjs/toolkit";
const initialState = {
    homepage:1,
    createProfileComponent: 1
}

export const globalReducer = createReducer(initialState , (builder) => {
    builder
    .addCase("OPEN_LOGIN_PANEL" , (state) => {
        state.openLoginPanel = true;
        state.openSignUpPanel = false;

    })
    .addCase("OPEN_SIGNUP_PANEL" , (state) => {
        state.openLoginPanel = false;
        state.openSignUpPanel = true;
        
    })
    .addCase("CLOSE_AUTH_PANELS" , (state) =>{
        state.openLoginPanel = false;
        state.openSignUpPanel = false;
    }) 
    .addCase("SET_HOMEPAGE", (state , action)=>{
        state.homepage = action.payload.homepage; //1 == prelogin , 2 == postlogin , 3 === CreateProfile
        state.requiredData = action.payload.data;  // for postlogin and prelogin it will be null
    })
    .addCase("SET_CREATE_PROFILE_COMPONENT" , (state , action) => {
        state.createProfileComponent = action.payload
    })
})
