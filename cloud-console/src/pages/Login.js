import { Navigate } from "react-router-dom";
import { LocalStorage } from "../Auth/LocalStorage";
import { useNavigate } from "react-router-dom";
import HTTPCalls from "../http/calls";



export const LoginPage = () => {
    const navigate = useNavigate();

    const userValues = {
        username: "",
        password: ""
    }

    const clickLogin = () => {
        console.log(userValues);
        HTTPCalls.requestLogin(userValues).then(loginresponse => {
            if(loginresponse.success){
                console.log(" loginresponse ", loginresponse);
                LocalStorage.set("user", { user: userValues.username, "token" :  loginresponse['token']})
                navigate('/')
            }
        }).catch(error => {
            console.log(" loginresponse error ", error, userValues);
        });
        
    }
    const updateUNameValue = (evt) => {
        const val = evt.target.value;
        userValues.username = (val);
    }
    const updatePwdValue = (evt) => {
        const val = evt.target.value;
        userValues.password = (val);
    }

    const { user } = LocalStorage.get("user");
    if (user) {
        return <Navigate to="/home" />;
    }

    return (
        <section className="vh-100">
            <div className="container-fluid h-custom">
                <div className="row mt-2"></div>
                <div className="row mt-2"></div>
                <div className="row mt-2"></div>
                <div className="row mt-2"></div>
                <div className="row mt-2"></div>
                <div className="row d-flex justify-content-center align-items-center h-100">
                    <div className="col-md-9 col-lg-6 col-xl-5">
                        <img src="https://img.freepik.com/free-vector/isolated-simple-cloud-sticker-template_1308-68055.jpg?w=2000"
                            className="img-fluid" alt="Sample image" />
                    </div>
                    <div className="col-md-8 col-lg-6 col-xl-4 offset-xl-1">
                        <div className="form-outline mb-4">
                            <input type="text" id="form3Example3" className="form-control form-control-lg"
                                placeholder="Enter a valid Username" onChange={evt => updateUNameValue(evt)}/>
                            <label className="form-label" htmlFor="form3Example3">Username</label>
                        </div>

                        <div className="form-outline mb-3">
                            <input type="text" id="form3Example4" className="form-control form-control-lg"
                                placeholder="Enter password" onChange={evt => updatePwdValue(evt)}/>
                            <label className="form-label" htmlFor="form3Example4">Password</label>
                        </div>


                        <div className="text-center text-lg-start mt-4 pt-2">
                            <button type="button" className="btn btn-primary btn-lg"
                                style={{ paddingLeft: '2.5rem', paddingRight: '2.5rem' }} onClick={clickLogin}>Login</button>
                            <p className="small fw-bold mt-2 pt-1 mb-0">Don't have an account?  Register with Postman..</p>
                        </div>

                    </div>
                </div>
                <div className="row mt-2"></div>
                <div className="row mt-2"></div>
                <div className="row mt-2"></div>
                <div className="row mt-2"></div>
                <div className="row mt-2"></div>
            </div>
            <div
                className="d-flex flex-column flex-md-row text-center text-md-start justify-content-between py-4 px-4 px-xl-5 bg-primary">
                <div className="text-white mb-3 mb-md-0">
                    Copyright © DummyCloud LLP. 2020. All rights reserved .
                </div>
            </div>
        </section>
    );
}

