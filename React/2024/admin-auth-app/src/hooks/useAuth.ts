import { useSelector } from "react-redux";

export function useAuth() {
  const { username, token } = useSelector((state: any) => state.user);
  return {
    isAuth: !!username,
    token,
  };
}
