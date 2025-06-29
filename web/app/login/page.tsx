'use client'
import { useState } from 'react'
import axios from 'axios'
export default function Login() {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const submit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    try { await axios.post(`${process.env.NEXT_PUBLIC_VAULT_API_URL}/auth/login`, { email, password }) } catch (e) { console.error(e) }
  }
  return (
    <div className="min-h-screen flex items-center justify-center bg-black">
      <form onSubmit={submit} className="w-80 p-8 rounded-lg shadow-xl bg-opacity-40 backdrop-blur border border-green-500">
        <h1 className="text-green-400 text-2xl font-mono mb-6 text-center">Tamaguru Login</h1>
        <input value={email} onChange={e=>setEmail(e.target.value)} type="email" placeholder="Email" className="w-full mb-4 px-3 py-2 bg-black border border-green-500 text-green-300 focus:outline-none focus:ring-2 focus:ring-green-500 font-mono"/>
        <input value={password} onChange={e=>setPassword(e.target.value)} type="password" placeholder="Password" className="w-full mb-6 px-3 py-2 bg-black border border-green-500 text-green-300 focus:outline-none focus:ring-2 focus:ring-green-500 font-mono"/>
        <button type="submit" className="w-full py-2 bg-green-500 text-black font-bold font-mono">ENTER</button>
      </form>
    </div>
  )
}
