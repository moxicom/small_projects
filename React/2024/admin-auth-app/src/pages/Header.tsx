import { useDispatch } from "react-redux";
import { Link, useNavigate } from "react-router-dom";
import { removeUser } from "../store/userSlice";

export default function Header() {
  const dispatcher = useDispatch();
  const navigate = useNavigate();
  const logOutClick = () => {
    console.log("Clicked - logOut");
    dispatcher(removeUser());
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
      <button className="text-sm font-semibold" onClick={logOutClick}>
        Log out
      </button>
    </div>
  );
}
