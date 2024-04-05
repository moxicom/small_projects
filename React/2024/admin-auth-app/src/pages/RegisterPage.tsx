import { useEffect, useState } from "react";
import { useForm, FieldError } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { FormData, registerRequest } from "../restRequests/register";

export default function RegisterPage() {
  const fieldStyle = "flex flex-col mb-2";
  const [loading, setLoading] = useState(false);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    mode: "onBlur",
    reValidateMode: "onBlur",
  });

  const navigate = useNavigate();

  useEffect(() => {}, []);

  function onSubmit(data: FormData) {
    console.log("Submitted details:", data);
    registerRequest(data);
    setLoading(true);
    // navigate(`/lists`);
  }

  function getEditorStyle(fieldError: FieldError | undefined) {
    return fieldError ? "border-red-500" : "";
  }
  if (loading) {
    return <>Loading...</>;
  } else {
    return (
      <>
        <div className="flex flex-col py-10 max-w-md mx-auto">
          <h2 className="text-3xl font-bold mb-3">Register page</h2>
          <form noValidate onSubmit={handleSubmit(onSubmit)}>
            <div className={fieldStyle}>
              <label htmlFor="name">Your name</label>
              <input
                type="text"
                id="name"
                {...register("name", {
                  required: "You must enter your name",
                })}
                className={getEditorStyle(errors.name)}
              />
            </div>
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
              <label htmlFor="password">Reason you need to contact us</label>
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
