import { useState } from "react";

export function AuthForm({
  title,
  handleClick,
  handleReset,
}: {
  title: string;
  handleClick: (email: string, password: string) => void;
  handleReset: (email: string) => void
}) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  return (
    <div className="flex flex-col py-10 max-w-md mx-auto bg-zinc-400 pt-10 pb-10  pr-5 pl-5">
      <div className="flex flex-col mb-2">
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
      </div>
      <div className="flex flex-col mb-2">
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
      </div>
      <div className="flex flex-col mb-2">
        <button onClick={() => {
          if (email != "" && password != ""){
            handleClick(email, password)
          } else {
            alert("EMPTY FIELDS")
          }
          }}>{title}</button>
      </div>
      <div className="flex flex-col mb-2">
        <button onClick={() => {
          if (email != "") {
            handleReset(email)
          } else {
            alert("Error :(");
          }
        }}>RESET PASSWORD USING SUBMITTED EMAIL </button>
      </div>
    </div>
  );
}
