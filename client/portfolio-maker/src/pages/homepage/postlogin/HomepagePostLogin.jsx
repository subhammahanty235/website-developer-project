import React, { useEffect } from 'react'
import './homepage.scss'
import PlusIcon from "../../../assets/icons/plus-svgrepo-com.svg"
import { useDispatch, useSelector } from 'react-redux'
import { createInitialProfile, fetchProfileOfUsers } from '../../../redux/actions/profileActions'
import UploadLogo from '../../../assets/icons/upload-minimalistic-svgrepo-com.svg'
const HomepagePostLogin = () => {
    const dispatch = useDispatch();
    const { profiles, fetchProfileLoading, createProfileLoading } = useSelector((state) => state.profileReducer)

    useEffect(() => {
        dispatch(fetchProfileOfUsers())
    }, [dispatch])

    return (
        <div className='homepage-postlogin'>
            <PostLoginNavbar />
            {
                (createProfileLoading || fetchProfileLoading) ?
                    <h1>Loading............</h1> :
                    <div className="homepage-postlogin-inner">
                        {
                            profiles?.map((profile) => {
                                return <ProfileButtonComponent profile={profile} />
                            })
                        }
                        <AddProfileButtonComponent />
                        {/* <AddProfileButtonComponent /> */}
                    </div>
            }


        </div>
    )
}


const AddProfileButtonComponent = () => {
    const dispatch = useDispatch()
    const createNewProfileHandler = () => {
        dispatch(createInitialProfile())
    }
    return (
        <>
            <div className="add_profile_button_component" onClick={()=>{dispatch(createInitialProfile())}}>
                <img src={PlusIcon} alt="" />
            </div>
        </>
    )
}

const ProfileButtonComponent = ({ profile }) => {
    const dispatch = useDispatch()
    const goToProfileHandler = () =>{
        console.log("heteeeeee")
        dispatch(
            {
              type: "SET_HOMEPAGE",
              payload: {
                homepage: 3,
                data:profile
              }
            }
          )
    }
    return (
        <>
            <div className="profile_button_component" onClick={goToProfileHandler}>
                <div className="top_part">
                    <div className="top_part_left">
                        <img src={profile.personalDetails[0]?.profilePic || UploadLogo} alt="" />
                    </div>
                    <div className="top_part_right">

                        <p className="profile_button_header">{profile.personalDetails[0]?.name || "New Profile"}</p>
                        <p className="profile_deployment_status"><div></div> Active</p>
                    </div>
                </div>

                <div className="bottom_part">
                    <div className="bottom_part_inner">
                        <ProfileButtonComponentChip value={"Personal Details"} active={profile.personalDetails?.length > 0} />
                        <ProfileButtonComponentChip value={"Education"} active={profile.educationDetail?.length > 0} />
                        <ProfileButtonComponentChip value={"Experience"} active={profile.experienceDetails?.length > 0} />
                        <ProfileButtonComponentChip value={"What You are Doing"} active={profile.whatAreYouDoingDetails?.length > 0} />
                        <ProfileButtonComponentChip value={"Testimonials"} active={profile.testimonialsDetails?.length > 0} />
                        <ProfileButtonComponentChip value={"Skills"} active={profile.skillsDetails?.length > 0} />
                        <ProfileButtonComponentChip value={"Portfolio Projects"} active={profile.portfolioDetails?.length > 0} />
                        <ProfileButtonComponentChip value={"Social Media"} active={profile.ocialMediaDetails?.length > 0} />

                    </div>
                </div>
            </div></>
    )
}

const ProfileButtonComponentChip = ({ value, active }) => {
    return (
        <div className='profile_button_component_chip'>
            {
                active ?
                    <div className="active"></div> :
                    <div className="inactive"></div>
            }

            <p>{value}</p>
        </div>
    )
}



const PostLoginNavbar = () => {
    const { user } = useSelector((state) => state.authReducer)
    const logOutUser = () => {
        localStorage.clear();

    }
    return (
        <>
            <div className='homepage__navbar__desktop'>
                <div className="navbar__logo">
                    Portfolio <span>Get</span>
                </div>

                <div className="navbar__options">
                    <p className='text_style_2' >{user?.name}</p>
                    <p className='text_style_2' onClick={() => { logOutUser() }}>Logout</p>
                </div>
            </div>
        </>
    )
}

export default HomepagePostLogin