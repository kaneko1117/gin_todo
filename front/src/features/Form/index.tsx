"use client"
import { useActionState } from 'react'
import { Button } from "@/components/button";
import { Input } from "@/components/input";
import { loginUser } from '@/app/action';

const initialState = {
    message:"",
}

export const From = () => {
    const [state, formAction, pending] = useActionState(loginUser,initialState)
    return (
        <form className="flex flex-col items-center justify-center gap-5 h-screen" action={formAction}>
            <Input
                type="text"
                placeholder="ユーザー名"
                className="w-1/2"
                id='username'
                name='username'
            />
            <Input
                type="password"
                placeholder="パスワード"
                className="w-1/2"
                id='password'
                name='password'
            />
            <Button type="submit" disabled={pending} className="cursor-pointer">ログイン</Button>
            <p className="text-red-500 h-6">{state.message}</p>
        </form>
    )
}