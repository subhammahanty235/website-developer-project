import React, { useState, useRef } from 'react'
import './createProfilepage.scss'
import UploadImage from '../../assets/icons/user-svgrepo-com.svg'
import UploadLogo from '../../assets/icons/upload-minimalistic-svgrepo-com.svg'
import PlusIcon from "../../assets/icons/plus-svgrepo-com.svg"
import {
    Box,
    TextField,
    MenuItem,
    Grid,
    IconButton,
} from '@mui/material';

import FacebookIcon from '@mui/icons-material/Facebook';
import TwitterIcon from '@mui/icons-material/Twitter';
import InstagramIcon from '@mui/icons-material/Instagram';
import LinkedInIcon from '@mui/icons-material/LinkedIn';
import YouTubeIcon from '@mui/icons-material/YouTube';
import { useDispatch, useSelector } from 'react-redux'
import { AddDataToProfile } from '../../redux/actions/profileActions'
import cloudinaryUpload from '../../utils/functions/cloudinary_upload'

const CreateProfilePage = () => {
    // const [comp, setComp] = useState(1)
    const dispatch = useDispatch()
    const { requiredData, createProfileComponent } = useSelector((state) => state.global)
    const setComp = (val) => {
        dispatch({
            type: "SET_CREATE_PROFILE_COMPONENT",
            payload: val
        })
    }
    return (
        <div className='create_profile_page'>
            <PostLoginNavbar />


            <div className="create_profile_page_inner">
                <p className="create_profile_header">Create Your Profile</p>
                <div className="profile_page_buttons">
                    <ProfileButtonComponentChip value={"Personal Details"} active={true} onClick={() => { setComp(1) }} />
                    <ProfileButtonComponentChip value={"Education"} active={true} onClick={() => { setComp(2) }} />
                    <ProfileButtonComponentChip value={"Experience"} active={true} onClick={() => { setComp(3) }} />
                    <ProfileButtonComponentChip value={"What You are Doing"} active={false} onClick={() => { setComp(4) }} />
                    <ProfileButtonComponentChip value={"Testimonials"} active={true} onClick={() => { setComp(5) }} />
                    <ProfileButtonComponentChip value={"Skills"} active={true} onClick={() => { setComp(6) }} />
                    <ProfileButtonComponentChip value={"Portfolio Projects"} active={true} onClick={() => { setComp(7) }} />
                    <ProfileButtonComponentChip value={"Social Media"} active={true} onClick={() => { setComp(8) }} />
                </div>

                <div className="forms">
                    {
                        createProfileComponent === 1 ? <PersonalDetailsInputForm personalDetailsData={requiredData?.personalDetails[0] || {}} /> :
                            createProfileComponent === 2 ? <EducationForm educationData={requiredData?.educationDetail} /> :
                                createProfileComponent === 3 ? <ExperienceForm experienceData={requiredData?.experienceDetails} /> :
                                    createProfileComponent === 43 ? <WhatYouAreDoing /> :
                                        createProfileComponent === 5 ? <TestimonialForm testimonialsData={requiredData?.testimonialsDetails} /> :
                                            createProfileComponent === 6 ? <SkillsForm skillData={requiredData?.skillsDetails} /> :
                                                createProfileComponent === 7 ? <ProjectForm portfolioDetails={requiredData?.portfolioDetails} /> :
                                                    createProfileComponent === 8 ? <SocialMediaForm socialMediaDetails={requiredData?.socialMediaDetails} /> :
                                                        <PersonalDetailsInputForm />
                    }


                </div>

            </div>
        </div>
    )
}




const PersonalDetailsInputForm = ({ personalDetailsData }) => {
    const [data, setData] = useState({
        firstname: ('' || personalDetailsData.name?.split(" ")[0]),
        lastname: ('' || personalDetailsData.name?.split(" ")[1]),
        email: ('' || personalDetailsData.email),
        mobileNumber: ('' || personalDetailsData.mobilenum),
        dateOfBirth: ('' || personalDetailsData.birthdate),
        location: ('' || personalDetailsData.location),
        bio: ('' || personalDetailsData.bio),
        profilePic: ('' || personalDetailsData.profilePic),
    })

    const onChange = (e) => {
        setData({ ...data, [e.target.name]: e.target.value })
    }
    return (
        <div className='personal_details_form'>

            <div className="personal_details_form_left">
                <div className="upload_image_button">
                    <img src={UploadImage} alt="" />

                    <button> <img src={UploadLogo} alt="" /> Upload Image</button>
                </div>
            </div>
            <div className="personal_details_form_right">
                <Box component="form" noValidate autoComplete="off" sx={{ mt: 3 }}>

                    <Grid container spacing={3}>
                        <Grid item xs={12} sm={4}>
                            <TextField required id="firstname" name="firstname" label="First Name" fullWidth variant="outlined" value={data.firstname} onChange={onChange} />
                        </Grid>
                        <Grid item xs={12} sm={4}>
                            <TextField required id="lastname" name="lastname" label="Last Name" fullWidth variant="outlined" value={data.lastname} onChange={onChange} />
                        </Grid>
                        <Grid item xs={12} sm={4}>
                            <TextField required id="email" name="email" label="Email" type="email" fullWidth variant="outlined" value={data.email} onChange={onChange} />
                        </Grid>
                        <Grid item xs={12} sm={6}>
                            <TextField required id="mobileNumber" name="mobileNumber" label="Mobile Number" fullWidth variant="outlined" value={data.mobileNumber} onChange={onChange} />
                        </Grid>
                        <Grid item xs={12} sm={6}>
                            <TextField required id="dateOfBirth" name="dateOfBirth" label="Date Of Birth" fullWidth variant="outlined" value={data.dateOfBirth} onChange={onChange} />
                        </Grid>
                        <Grid item xs={12} sm={12}>
                            <TextField required id="location" name="location" label="Location" fullWidth variant="outlined" value={data.location} onChange={onChange} />
                        </Grid>
                        <Grid item xs={12} sm={12}>
                            <TextField id="bio" name="bio" label="Bio" multiline rows={6} fullWidth variant="outlined" value={data.bio} onChange={onChange} />
                        </Grid>


                        <Grid item xs={12}>
                            <div className="submit_button">
                                <button>Submit and Go next</button>
                            </div>
                        </Grid>
                    </Grid>
                </Box>
            </div>



        </div>
    )
}

const EducationForm = ({ educationData }) => {

    const [educationDetails, setEducationDetails] = useState(educationData || [
        {
            institutename: '',
            course: '',
            startyear: '',
            endyear: '',
            description: '',
        },
    ]);

    const handleAddMore = () => {
        setEducationDetails([...educationDetails, {
            institutename: '',
            course: '',
            startyear: '',
            endyear: '',
            description: '',
        }]);
    };

    const handleInputChange = (index, e) => {
        const { name, value } = e.target;
        const updatedDetails = [...educationDetails];
        updatedDetails[index] = { ...updatedDetails[index], [name]: value };
        setEducationDetails(updatedDetails);
    };

    const handleSubmit = () => {
        dispatch(AddDataToProfile(2, educationDetails))
    }

    return (
        <div className="education_form">
            <Box noValidate autoComplete="off" sx={{ mt: 3 }}>

                {educationDetails.map((education, index) => (
                    <Box key={index} sx={{ mb: 3, mt: 5 }} >
                        <Grid container spacing={3}>
                            <Grid item xs={12}>
                                <TextField required name="institutename" label="Institute Name" fullWidth variant="outlined" value={education.institutename} onChange={(e) => handleInputChange(index, e)} />
                            </Grid>
                            <Grid item xs={12}>
                                <TextField required name="course" label="Course" fullWidth variant="outlined" value={education.course} onChange={(e) => handleInputChange(index, e)} />
                            </Grid>
                            <Grid item xs={6}>
                                <TextField required name="startyear" label="Start Year" type="number" fullWidth variant="outlined" value={education.startyear} onChange={(e) => handleInputChange(index, e)} />
                            </Grid>
                            <Grid item xs={6}>
                                <TextField required name="endyear" label="End Year" type="number" fullWidth variant="outlined" value={education.endyear} onChange={(e) => handleInputChange(index, e)} />
                            </Grid>
                            <Grid item xs={12}>
                                <TextField name="description" label="Description" multiline rows={4} fullWidth variant="outlined" value={education.description} onChange={(e) => handleInputChange(index, e)} />
                            </Grid>
                        </Grid>
                    </Box>
                ))}
                <Grid item xs={12}>
                    <div className="submit_button">
                        <button className="addmore" onClick={handleAddMore}>+ Add More</button>
                        <button onClick={() => handleSubmit}>Submit and Go next</button>

                    </div>
                </Grid>
            </Box>
        </div>

    );
};
const ExperienceForm = ({ experienceData }) => {

    const dispatch = useDispatch()
    const [experienceDetails, setExperienceDetails] = useState(experienceData || [
        {
            companyname: '',
            designation: '',
            startyear: '',
            endyear: '',
            description: '',
        },
    ]);

    const handleAddMore = () => {
        setExperienceDetails([...experienceDetails, {
            companyname: '',
            designation: '',
            startyear: '',
            endyear: '',
            description: '',
        }]);
    };

    const handleInputChange = (index, e) => {
        const { name, value } = e.target;
        const updatedDetails = [...experienceDetails];
        updatedDetails[index] = { ...updatedDetails[index], [name]: value };
        setExperienceDetails(updatedDetails);
    };


    const handleSubmit = () => {
        dispatch(AddDataToProfile(3, experienceDetails))
    }



    return (
        <div className="education_form">
            <Box noValidate autoComplete="off" sx={{ mt: 3 }}>

                {experienceDetails.map((experience, index) => (
                    <Box key={index} sx={{ mb: 3, mt: 5 }} >
                        <Grid container spacing={3}>
                            <Grid item xs={12}>
                                <TextField required name="companyname" label="Company Name" fullWidth variant="outlined" value={experience.companyname} onChange={(e) => handleInputChange(index, e)} />
                            </Grid>
                            <Grid item xs={12}>
                                <TextField required name="designation" label="Designation" fullWidth variant="outlined" value={experience.designation} onChange={(e) => handleInputChange(index, e)} />
                            </Grid>
                            <Grid item xs={6}>
                                <TextField required name="startyear" label="Start Year" type="number" fullWidth variant="outlined" value={experience.startyear} onChange={(e) => handleInputChange(index, e)} />
                            </Grid>
                            <Grid item xs={6}>
                                <TextField required name="endyear" label="End Year" type="number" fullWidth variant="outlined" value={experience.endyear} onChange={(e) => handleInputChange(index, e)} />
                            </Grid>
                            <Grid item xs={12}>
                                <TextField name="description" label="Description" multiline rows={4} fullWidth variant="outlined" value={experience.description} onChange={(e) => handleInputChange(index, e)} />
                            </Grid>
                        </Grid>
                    </Box>
                ))}
                <Grid item xs={12}>
                    <div className="submit_button">
                        <button className="addmore" onClick={handleAddMore}>+ Add More</button>
                        <button onClick={() => handleSubmit()}>Submit and Go next</button>

                    </div>
                </Grid>
            </Box>
        </div>

    );
};

const TestimonialForm = ({ testimonialsData }) => {
    const dispatch = useDispatch()
    const [testimonialDetails, setTestimonialDetails] = useState(testimonialsData?.length > 0 ? testimonialsData : [
        {
            header: '',
            subHeader: '',
            date: '',
            image: '',

        },
    ]);

    const handleAddMore = () => {
        setTestimonialDetails([...testimonialDetails, {
            header: '',
            subHeader: '',
            date: '',
            image: '',
        }]);
    };

    const handleInputChange = (index, e) => {
        const { name, value } = e.target;
        const updatedDetails = [...testimonialDetails];
        updatedDetails[index] = { ...updatedDetails[index], [name]: value };
        setTestimonialDetails(updatedDetails);
    };

    const handleSubmit = () => {
        dispatch(AddDataToProfile(5, testimonialDetails))
    }

    return (
        <div className="testimonial_form">
            <Box noValidate autoComplete="off" sx={{ mt: 3 }}>

                {testimonialDetails.map((testimonialDetail, index) => (

                    <Box key={index} sx={{ mb: 3, mt: 5 }} >
                        <Grid container spacing={3} mt={5}>
                            <Grid container xs={6} >
                                <div className="upload_image_section">
                                    <div className="upload_image_section_inner">
                                        <div className="upload_image_section_inner_left">
                                            <div className="upload_image_button">
                                                <img src={UploadImage} alt="" />

                                                <button> <img src={UploadLogo} alt="" /> Upload Image</button>
                                                {/* <p>Or Choose from the list</p> */}
                                            </div>
                                        </div>
                                        <p>or</p>
                                        <div className="upload_image_section_inner_right">
                                            <div className="profile_pic_box1">
                                                <img src="https://codewithsadee.github.io/vcard-personal-portfolio/assets/images/avatar-1.png" alt="" />
                                            </div>
                                            <div className="profile_pic_box2">
                                                <img src="https://codewithsadee.github.io/vcard-personal-portfolio/assets/images/avatar-2.png" alt="" />
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </Grid>
                            <Grid container xs={6} spacing={3}>
                                <Grid item xs={12}>
                                    <TextField required name="header" label="Testimonial By" fullWidth variant="outlined" value={testimonialDetail.header} onChange={(e) => handleInputChange(index, e)} />
                                </Grid>
                                <Grid item xs={12}>
                                    <TextField id="subHeader" name="subHeader" label="Testimonial" multiline rows={6} value={testimonialDetail.subHeader} fullWidth variant="outlined" />
                                </Grid>
                                <Grid item xs={12}>
                                    <TextField required name="date" label="Date" type="number" fullWidth variant="outlined" value={testimonialDetail.date} onChange={(e) => handleInputChange(index, e)} />
                                </Grid>
                            </Grid>

                        </Grid>
                    </Box>
                ))}
                <Grid item xs={12}>
                    <div className="submit_button">
                        <button className="addmore" onClick={handleAddMore}>+ Add More</button>
                        <div className="submit_button_inner">
                            <button>Skip and Go next</button>
                            <button onClick={() => handleSubmit()}>Submit and Go next</button>
                        </div>

                    </div>
                </Grid>
            </Box>
        </div>

    );
};



const SkillsForm = ({ skillData }) => {
    const dispatch = useDispatch()
    const [skills, setSkills] = useState(skillData?.length > 0 ? skillData : [{ skillName: '', knowledge: '' }]);

    const skillLevels = [
        { value: 'entry', label: 'Entry Level' },
        { value: 'intermediate', label: 'Intermediate' },
        { value: 'pro', label: 'Pro' }
    ];


    const handleAddSkill = () => {
        setSkills([...skills, { skillName: '', knowledge: '' }]);
    };

    const handleChange = (index, event) => {
        const newSkills = skills.map((skill, i) => {
            if (i === index) {
                return { ...skill, [event.target.name]: event.target.value };
            }
            return skill;
        });
        setSkills(newSkills);
    };

    const handleSubmit = () => {
        dispatch(AddDataToProfile(6, skills))
    }

    return (
        <div className="skills_form">
            <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 2 }}>
                {skills.map((skill, index) => (
                    <Box key={index} sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                        <TextField
                            name="skillName"
                            label="Skill Name"
                            value={skill.skillname}
                            onChange={(e) => handleChange(index, e)}
                            fullWidth
                            variant="outlined"
                        />
                        <TextField
                            select
                            name="knowledge"
                            label="Level"
                            value={skill.knowledge}
                            onChange={(e) => handleChange(index, e)}
                            fullWidth
                            variant="outlined"
                        >
                            {skillLevels.map((option) => (
                                <MenuItem key={option.value} value={option.value}>
                                    {option.label}
                                </MenuItem>
                            ))}
                        </TextField>
                    </Box>
                ))}
                <div className="add_more_skill_button" onClick={handleAddSkill}>
                    <img src={PlusIcon} alt="" />
                </div>
            </Box>
            <Grid item xs={12}>
                <div className="submit_button">
                    <button onClick={() => handleSubmit()}>Submit and Go next</button>

                </div>
            </Grid>
        </div>

    );
};

const ProjectForm = ({ portfolioDetails }) => {
    const fileInputRef = useRef(null);
    const dispatch = useDispatch()
    const [imageUploading, setImageUploading] = useState(false)
    const [projectDetails, setProjectDetails] = useState(portfolioDetails.length > 0 ? portfolioDetails : [
        {
            homepageImage: '',
            projectName: '',
            category: '',
            liveLink: '',
            githubLink: ''

        },
    ]);



    const handleAddMore = () => {
        setProjectDetails([...projectDetails, {
            homepageImage: '',
            projectName: '',
            category: '',
            liveLink: '',
            githubLink: ''
        }]);
    };

    const handleInputChange = (index, e) => {
        const { name, value } = e.target;
        const updatedDetails = [...projectDetails];
        updatedDetails[index] = { ...updatedDetails[index], [name]: value };
        setProjectDetails(updatedDetails);
    };

    const projectCategories = [
        { value: 'webdevelopment', label: 'Web Development' },
        { value: 'webdesign', label: 'Web Design' },
        { value: 'applications', label: 'Applications' },
        { value: 'uiux', label: 'UI/UX' }
    ];

    const handleChange = (index, event) => {
        const newProject = projectDetails.map((skill, i) => {
            if (i === index) {
                return { ...skill, [event.target.name]: event.target.value };
            }
            return skill;
        });
        setProjectDetails(newProject)
    };

    const handleImageUpload = async (event, index) => {
        const file = event.target.files[0];
        if (!file) return;

        try {
            setImageUploading(true)
            const url = await cloudinaryUpload(file);
            setImageUploading(false)

            const updatedDetails = [...projectDetails];
            updatedDetails[index] = { ...updatedDetails[index], homepageImage: url };
            setProjectDetails(updatedDetails);

        } catch (error) {
            console.error('Error uploading image:', error);
        }

    };

    const handleClick = () => {
        fileInputRef.current.click();
    };

    const handleSubmit = () => {
        dispatch(AddDataToProfile(7, projectDetails))
    }


    return (
        <>
            <div className="project_form">
                <Box noValidate autoComplete="off" sx={{ mt: 3 }}>

                    {projectDetails.map((experience, index) => (

                        <Box key={index} sx={{ mb: 3, mt: 5 }}>
                            <Grid container spacing={3} mt={5}>
                                <Grid container xs={6} >
                                    {
                                        (experience.homepageImage === '') || (experience.homepageImage === undefined) ?
                                        <>
                                                <div className="image_upload" onClick={handleClick}>
                                                    <div className="image_upload_inner">
                                                        <img src={UploadLogo} alt="" />
                                                        {
                                                            imageUploading === true ? <p>Uploading Image</p> : <p>Upload Project Image</p>
                                                        }

                                                    </div>
                                                </div>
                                                <input
                                                    ref={fileInputRef}
                                                    type="file"
                                                    onChange={(e) => handleImageUpload(e, index)}
                                                    style={{ display: 'none' }}
                                                />
                                            </>
                                            :
                                            <div className="image_uploaded">
                                                 
                                                    <img className='uploaded_image' src={experience.homepageImage} alt="" />

                                            </div>
                                            
                                            
                                    }

                                </Grid>
                                <Grid container xs={6} spacing={3}>
                                    <Grid item xs={12}>
                                        <TextField required name="projectName" label="Project Name" fullWidth variant="outlined" value={experience.projectName} onChange={(e) => handleInputChange(index, e)} />
                                    </Grid>
                                    <Grid item xs={12}>
                                        <TextField
                                            select
                                            name="category"
                                            label="Category"
                                            value={experience.category}
                                            onChange={(e) => handleChange(index, e)}
                                            fullWidth
                                            variant="outlined"
                                        >
                                            {projectCategories.map((option) => (
                                                <MenuItem key={option.value} value={option.value}>
                                                    {option.label}
                                                </MenuItem>
                                            ))}
                                        </TextField>
                                    </Grid>
                                    <Grid item xs={12}>
                                        <TextField required name="liveLink" label="Live Link" fullWidth variant="outlined" value={experience.liveLink} onChange={(e) => handleInputChange(index, e)} />
                                    </Grid>
                                    <Grid item xs={12}>
                                        <TextField required name="githubLink" label="Github Link" fullWidth variant="outlined" value={experience.githubLink} onChange={(e) => handleInputChange(index, e)} />
                                    </Grid>
                                </Grid>

                            </Grid>
                        </Box>
                    ))}
                    <Grid item xs={12}>
                        <div className="submit_button">
                            <button className="addmore" onClick={handleAddMore}>+ Add More</button>

                            {/* <button>Skip and Go next</button> */}
                            <button onClick={handleSubmit}>Submit and Go next</button>


                        </div>
                    </Grid>
                </Box>
            </div>
        </>
    )
}

const WhatYouAreDoing = () => {
    return (
        <div className="notactivemessage">
            Coming Soon
        </div>
    )
}


const SocialMediaForm = ({ socialMediaDetails }) => {
    const active = false;
    // const dispatch = useDispatch()
    // const [projectDetails, setProjectDetails] = useState(socialMediaDetails.length > 0 ? socialMediaDetails : [
    //     {
    //         socialMediaName: '',
    //         icon: '',
    //         profileLink: '',

    //     },
    // ]);

    // const handleAddMore = () => {
    //     setProjectDetails([...projectDetails, {
    //         socialMediaName: '',
    //         icon: '',
    //         profileLink: '',
    //     }]);
    // };

    // const handleInputChange = (index, e) => {
    //     const { name, value } = e.target;
    //     const updatedDetails = [...projectDetails];
    //     updatedDetails[index] = { ...updatedDetails[index], [name]: value };
    //     setProjectDetails(updatedDetails);
    // };

    return (
        <>
            {
                active === true ?
                    <div className="social_media_form">
                        <Box
                            component="form"
                            sx={{
                                display: 'flex',
                                flexDirection: 'column',
                                gap: 3,
                                width: '100%',
                                maxWidth: '600px',

                            }}
                        >
                            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                                <IconButton>
                                    <FacebookIcon />
                                </IconButton>
                                <TextField
                                    fullWidth
                                    label="Facebook URL"
                                    variant="outlined"

                                />
                            </Box>
                            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                                <IconButton>
                                    <TwitterIcon />
                                </IconButton>
                                <TextField
                                    fullWidth
                                    label="Twitter URL"
                                    variant="outlined"
                                />
                            </Box>
                            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                                <IconButton>
                                    <InstagramIcon />
                                </IconButton>
                                <TextField
                                    fullWidth
                                    label="Instagram URL"
                                    variant="outlined"
                                />
                            </Box>
                            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                                <IconButton>
                                    <LinkedInIcon />
                                </IconButton>
                                <TextField
                                    fullWidth
                                    label="LinkedIn URL"
                                    variant="outlined"
                                />
                            </Box>
                            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                                <IconButton>
                                    <YouTubeIcon />
                                </IconButton>
                                <TextField
                                    fullWidth
                                    label="YouTube URL"
                                    variant="outlined"
                                />
                            </Box>
                        </Box>
                        <Grid item xs={6}>
                            <div className="submit_button">
                                <button>Submit and Go next</button>

                            </div>
                        </Grid>
                    </div>
                    :
                    <div className="notactivemessage">
                        Coming Soon
                    </div>
            }
        </>

    )
}


const ProfileButtonComponentChip = ({ value, active, onClick }) => {
    return (
        <div className='profile_button_component_chip' onClick={onClick}>
            {
                active ?
                    <div className="filled"></div> :
                    <div className="notfilled"></div>
            }

            <p>{value}</p>
        </div>
    )
}


const PostLoginNavbar = () => {
    return (
        <>
            <div className='homepage__navbar__desktop'>
                <div className="navbar__logo">
                    Portfolio <span>Get</span>
                </div>

                <div className="navbar__options">
                    <p className='text_style_2' >Subham Mahanty</p>
                    <p className='text_style_2' onClick={() => { }}>Logout</p>
                </div>
            </div>
        </>
    )
}

export default CreateProfilePage