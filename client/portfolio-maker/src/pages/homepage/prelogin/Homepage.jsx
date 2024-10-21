import React, { useEffect, useState } from 'react';
import './homepage.scss'
// import { openLoginPanel } from "../../redux/reducers/globalReducer"
import { useDispatch } from 'react-redux';

const TOTAL_STEPS = 3;
const ANIMATION_DURATION = 5000;

const Homepage = () => {
    const [currentStep, setCurrentStep] = useState(0);

    useEffect(() => {
        const handleStepChange = () => {
            setCurrentStep(prevStep => (prevStep + 1) % TOTAL_STEPS);
        };

        const timeout = setTimeout(handleStepChange, ANIMATION_DURATION);

        return () => clearTimeout(timeout);
    }, [currentStep]);

    const handleStepClick = (step) => {
        setCurrentStep(step);
    };
    return (
        <div className='homepage'>
            <NavigationBar />
            <div className="homepage_inner">
                <p className="header_text">Get Your Portfolio Website in <span>THREE</span> Steps</p>
                <div className="container">
                    <div className="progress">
                        <div
                            className="percent"
                            style={{
                                width: `${(currentStep * 100) / (TOTAL_STEPS - 1)}%`,
                                transition: `width ${ANIMATION_DURATION / 1000}s linear, height ${ANIMATION_DURATION / 1000}s linear`,
                            }}
                        ></div>
                    </div>
                    <div className="steps">
                        {[0, 1, 2].map(step => (
                            <div
                                key={step}
                                className={`step ${currentStep === step ? 'selected' : ''} ${currentStep >= step ? 'completed' : ''}`}
                                onClick={() => handleStepClick(step)}
                            ></div>
                        ))}
                    </div>
                    <div className="steps-mob">
                        {[{ key: "Create an Account", value: "create a new account or login into your existing account " },
                        { key: "Fill The Details", value: "Fill all the details to get the best looking portfolio{" },
                        { key: "  Save and Deploy", value: "Just save & automatically deploy your portfolio" }].map((step, index) => (
                            <div
                                key={index}
                                className={`step-mob ${currentStep === index ? 'selected' : ''} ${currentStep >= index ? 'completed' : ''}`}
                                onClick={() => handleStepClick(index)}
                            >{step.key}</div>
                        ))}
                    </div>

                </div>
                <div className="progress_texts">
                    {[{ key: "Create an Account", value: "create a new account or login into your existing account" },
                    { key: "Fill The Details", value: "Fill all the details to get the best looking portfolio{" },
                    { key: "  Save and Deploy", value: "Just save & automatically deploy your portfolio" }].map((step, index) => {
                        console.log(`Current Step is : ${currentStep} and index is ${index}`)
                        return <div
                            key={index}
                            className={`texts ${currentStep === index ? 'selected' : ''} ${currentStep > index ? 'completed' : ''}`}
                            onClick={() => handleStepClick(index)}
                        >{step.key} <br />
                        </div>

                    })}
                </div>
                <div className="progress_texts-sec">
                    {[{ key: "Create an Account", value: "create a new account or \nlogin into your existing account" }, { key: "Fill The Details", value: "Fill all the details to get \n  the best looking portfolio" }, { key: "Save and Deploy", value: "Just save and we \nwill automatically deploy your \n Portfolio Website." }].map((step, index) => (
                        <div
                            key={index}
                            className={`texts-sec ${"index" + index} ${currentStep === index ? 'selected' : ''} ${currentStep >= index ? 'completed' : ''}`}
                            onClick={() => handleStepClick(index)}
                        >{step.value.split('\n').map((line, i) => (
                            <React.Fragment key={i}>
                                {line}
                                <br />
                            </React.Fragment>
                        ))} <br />
                        </div>
                    ))}
                </div>


            </div>
        </div>
    )
}

const NavigationBar = () => {
    const dispatch = useDispatch()
    const loginButtonHandler = () => {
        dispatch(
            {
              type:"OPEN_LOGIN_PANEL"
            }
          )
    };

    return (
        <>
            <div className='homepage__navbar__desktop'>
                <div className="navbar__logo">
                    Portfolio <span>Get</span>
                </div>

                <div className="navbar__options">
                    <p className='text_style_2' onClick={() => { loginButtonHandler()}}>Log In</p>
                    <p className='text_style_2' onClick={() => { }}>Create new Account</p>
                </div>
            </div>
        </>
    )
}

export default Homepage