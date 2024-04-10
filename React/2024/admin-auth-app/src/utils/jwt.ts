import { jwtDecode } from "jwt-decode";

interface DecodedToken {
    exp: number;
}

export function getExp(token: string): Date {
    const decodedExp = jwtDecode<DecodedToken>(token)
    const date = new Date(decodedExp.exp * 1000)
    console.log("date exp = " + date)
    return date
}

export function isJWTValid(timeNow: Date, token: string): boolean {
    const expTime = getExp(token);
    return timeNow <= expTime;
}