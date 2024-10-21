import axios from 'axios'
export const createInitialProfile = () => async (dispatch) => {
    try {
        console.log("Creating new profile")
        dispatch({ type: "CREATE_INITIAL_PROFILE" })
        const response = await axios.post(`${import.meta.env.VITE_REACT_BACKEND_HOSTED_URL}/user/create`,
            {},
            {
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("token")}`
                },
            }
        )

        if (response.data.success) {
            dispatch({
                type: "CREATE_INITIAL_PROFILE_SUCCESS"
            })
            dispatch(fetchProfileOfUsers())
        } else {
            dispatch({
                type: "CREATE_INITIAL_PROFILE_FAILED",
                payload: response.data.error
            })
        }
    } catch (error) {
        dispatch({
            type: "CREATE_INITIAL_PROFILE_FAILED",
            payload: error
        })
    }
}

export const fetchProfileOfUsers = () => async (dispatch) => {
    try {
        dispatch({ type: "FETCH_PROFILE_DETAILS" })
        const response = await axios.get(`${import.meta.env.VITE_REACT_BACKEND_HOSTED_URL}/user/getprofiles`,
            {
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("token")}`
                },
            }
        )

        if (response.data.success) {
            console.log(response.data.profile)
            dispatch({
                type: "FETCH_PROFILE_DETAILS_SUCCESS",
                payload: {
                    profiles: response.data.profile,
                    emptyProfiles: response.data.profile.length > 0 ? false : true
                }
            })
        } else {
            dispatch({
                type: "FETCH_PROFILE_DETAILS_FAILED",
                payload: response.data.error
            })
        }
    } catch (error) {
        dispatch({
            type: "FETCH_PROFILE_DETAILS_FAILED",
            payload: error
        })
    }
}

export const AddDataToProfile = (seg, data) => async (dispatch) => {
    try {
        dispatch({ type: "ADD_DATA_TO_PROFILE" })
        const response = await axios.post(`${import.meta.env.VITE_REACT_BACKEND_HOSTED_URL}/user/adddata/${seg}`,

            data
            ,
            {
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("token")}`
                },
            }
        )

        if (response.data.success) {
            console.log(response.data.profile)
            dispatch({
                type: "ADD_DATA_TO_PROFILE_SUCCESS",

            })
        } else {
            dispatch({
                type: "ADD_DATA_TO_PROFILE_FAILED",
                payload: response.data.error
            })
        }
    } catch (error) {
        dispatch({
            type: "ADD_DATA_TO_PROFILE_FAILED",
            payload: error
        })
    }
}

export const checkDeploymentEligiblity = (seg, data) => async (dispatch) => {
    try {
        dispatch({ type: "CHECK_WEBSITE_DEPLOYMENT_ELIGIBILITY" })
        const response = await axios.post(`${import.meta.env.VITE_REACT_BACKEND_HOSTED_URL}/deploy/checkeligiblity/`,
            {
                profileId: data.profileId
            }
            ,
            {
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("token")}`
                },
            }
        )

        if (response.data.success) {
            console.log(response.data.profile)
            dispatch({
                type: "CHECK_WEBSITE_DEPLOYMENT_ELIGIBILITY_SUCCESS",
                payload:{
                    eligible:response.data.eligible,
                    alreadyDeployed:response.data.alreadyDeployed
                }

            })
        } else {
            dispatch({
                type: "CHECK_WEBSITE_DEPLOYMENT_ELIGIBILITYE_FAILED",
                payload: response.data.error
            })
        }
    } catch (error) {
        dispatch({
            type: "CHECK_WEBSITE_DEPLOYMENT_ELIGIBILITY_FAILED",
            payload: error
        })
    }
}


