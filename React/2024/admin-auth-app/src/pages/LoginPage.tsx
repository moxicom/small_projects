import { useEffect, useState } from "react";
import { useForm, FieldError } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";
import { loginRequest } from "../restRequests/login";
import { useDispatch } from "react-redux";
import { setUser } from "../store/userSlice";

type FormData = {
  username: string;
  password: string;
};

export default function LoginPage() {
  const navigate = useNavigate();
  const dispatch = useDispatch();
  const fieldStyle = "flex flex-col mb-2";

  const [isLoading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    mode: "onBlur",
    reValidateMode: "onBlur",
  });

  async function onSubmit(data: FormData) {
    console.log("Submitted details:", data);
    setLoading(true);
    try {
      const response = await loginRequest(data);
      setError(null);
      dispatch(setUser({ token: response.token, username: data.username }));
      navigate("/lists");
    } catch (error) {
      console.log(error);
      if (error === 401) {
        setError("Incorrect username or password");
      }
    } finally {
      setLoading(false);
    }
  }

  function getEditorStyle(fieldError: FieldError | undefined) {
    return fieldError ? "border-red-500 bg-black border-2" : "bg-black";
  }
  if (isLoading) {
    return <>Loading...</>;
  } else {
    return (
      <>
        <div className="flex flex-col py-10 max-w-md mx-auto">
          <h2 className="text-3xl font-bold mb-3">Login page</h2>
          <form noValidate onSubmit={handleSubmit(onSubmit)}>
            <div className={fieldStyle}>
              <label htmlFor="username">Your username</label>
              <input
                type="text"
                id="username"
                {...register("username", {
                  required: "You must enter your username",
                })}
                className={getEditorStyle(errors.username)}
              />
            </div>
            <div className={fieldStyle}>
              <label htmlFor="password">Password</label>
              <input
                type="password"
                id="password"
                {...register("password", {
                  required: "You must enter the passowrd",
                })}
                className={getEditorStyle(errors.password)}
              />
            </div>
            <div>
              <div>
                Don't have an account?
                <Link to={"/register"}> Go register page!</Link>
              </div>
              {error && (
                <div className="w-full text-red-600 bg-red-950 border-2 rounded-md p-1 mt-3 border-red-500">
                  {error}
                </div>
              )}
              <button
                type="submit"
                className="mt-2 h-10 px-6 font-semibold bg-black text-white"
              >
                Submit
              </button>
            </div>
          </form>
        </div>
      </>
    );
  }
}
