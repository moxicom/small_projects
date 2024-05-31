import { useDispatch } from "react-redux";
import { Link, useNavigate } from "react-router-dom";
import useLogOut from "../hooks/useLogout";
import avatar from '../assets/avatar.jpg'; // Import the avatar image

export default function Header() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const logOutClick = () => {
    console.log("Clicked - logOut");
    useLogOut(dispatch)
    navigate("/login");
  };
  return (
    <div
      className="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8"
      aria-label="Global"
    >
      <Link to={"/"}>
        <h1 className="text-5xl">Home</h1>
      </Link>
      <div className="flex items-center">
        <button className="text-sm font-semibold mr-4" onClick={logOutClick}>
          Log out
        </button>
        <img src={avatar} alt="User Avatar" className="w-10 h-10 rounded-full" />
      </div>
    </div>
  );
}