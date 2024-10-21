import { createReducer } from "@reduxjs/toolkit";
const initialState = {}

export const profileReducer = createReducer(initialState, (builder) => {
    builder
        .addCase("CREATE_INITIAL_PROFILE", state => {
            state.createProfileLoading = true;
        })
        .addCase("CREATE_INITIAL_PROFILE_SUCCESS", state => {
            state.createProfileLoading = false;
        })
        .addCase("CREATE_INITIAL_PROFILE_FAILED", state => {
            state.createProfileLoading = false;
        })

        .addCase("FETCH_PROFILE_DETAILS", (state) => {
            state.fetchProfileLoading = true;
        })
        .addCase("FETCH_PROFILE_DETAILS_SUCCESS", (state, action) => {
            state.fetchProfileLoading = false;
            state.profiles = action.payload.profiles
            state.emptyProfiles = action.payload.emptyProfiles
        })
        .addCase("FETCH_PROFILE_DETAILS_FAILED", (state , action) => {
            state.fetchProfileLoading = false;
            state.error = action.payload;
        })

        .addCase("ADD_DATA_TO_PROFILE" , (state) =>{
            state.addDataToProfileLoading = true;
        })
        .addCase("ADD_DATA_TO_PROFILE_SUCCESS" , (state) =>{
            state.addDataToProfileLoading = false;
        })
        .addCase("ADD_DATA_TO_PROFILE_FAILED" , (state) =>{
            state.addDataToProfileLoading = false;
        })
        .addCase("CHECK_WEBSITE_DEPLOYMENT_ELIGIBILITY", (state)=>{
            state.checkingElegiblity = true
        })
        .addCase("CHECK_WEBSITE_DEPLOYMENT_ELIGIBILITY_SUCCESS", (state, action)=>{
            state.checkingElegiblity = false;
            state.deploymentEligible = action.payload.eligible,
            state.alreadyDepoyed = action.payload?.alreadyDeployed || false
        })
        .addCase("CHECK_WEBSITE_DEPLOYMENT_ELIGIBILITY_FAILED", (state,action)=>{
            state.checkingElegiblity = false;
            state.eligible = false;
            state.deploymentElErr = action.payload
        })

})



/*
"profile": [
        {
            "_id": "6689895378d01fd0ef7da672",
            "educationDetaiil": [
                "668a3e2b29596e0f5f82aba3",
                "668a3e2c29596e0f5f82aba4"
            ],
            "educationDetail": [
                {
                    "_id": "668a3e2c29596e0f5f82aba4",
                    "course": "c2",
                    "description": "Hello i am doing it 2",
                    "endyear": "2024",
                    "institutename": "institude 2 ",
                    "startyear": "2022"
                },
                {
                    "_id": "668a3e2b29596e0f5f82aba3",
                    "course": "c1",
                    "description": "Hello i am doing it",
                    "endyear": "2024",
                    "institutename": "institude 1 ",
                    "startyear": "2022"
                }
            ],
            "experienceDetails": [],
            "personalDetails": [
                {
                    "_id": "66899d049b28b61278496e9e",
                    "bio": "Hello i am here",
                    "birthdate": "1/11/2002",
                    "email": "subhammahantyyyy351@gmail.com",
                    "location": "Kolkata",
                    "mobilenum": "9932237171"
                }
            ],
            "personalDetailsId": "66899d049b28b61278496e9e",
            "portfolioDetails": [],
            "skillsDetails": [],
            "socialMediaDetails": [],
            "testimonialsDetails": [],
            "userId": "66900013bf41b4a7e57d1ff8",
            "whatAreYouDoingDetails": []
        }
    ],
    "success": true
*/