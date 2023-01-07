import { Navigate } from "react-router-dom";
import { LocalStorage } from "../Auth/LocalStorage";
import { useNavigate } from "react-router-dom";




export const LoginPage = () => {
    const navigate = useNavigate();
    const clickLogin = () => {
        LocalStorage.set("user", {user: "anshuman"})
        navigate('/')
    }

    const { user } = LocalStorage.get("user");
        if (user) {
            return <Navigate to="/home" />;
        }
    return (
        <div>
            <h1>This is the Login Page</h1>
            <button onClick={clickLogin}>Login</button>
        </div>
    );
}

