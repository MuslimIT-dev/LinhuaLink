import React, { createContext, useContext, useEffect, useState } from 'react';

const AuthContext = createContext();

export function AuthProvider({ children }) {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

    // check auth status
    useEffect(() => {
        let mounted = true;
        async function check() {
            try {
                const res = await fetch('/api/auth/me', {
                    method: 'GET',
                    credentials: 'include',
                    headers: { 'Accept': 'application/json' },
                });
                if (!mounted) return;
                if (res.ok) {
                    const data = await res.json();
                    setUser(data.user || null);
                } else {
                    setUser(null);
                }
            } catch (err) {
                console.error('Auth check failed', err);
                setUser(null);
            } finally {
                if (mounted) setLoading(false);
            }
        }
        check();
        return () => { mounted = false; };
    }, []);

    const login = async (credentials) => {
        const res = await fetch('/api/auth/login', {
            method: 'POST',
            credentials: 'include', // сервер выставит httpOnly cookie
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(credentials),
        });

        if (!res.ok) {
            const err = await res.json().catch(() => ({ message: 'Login failed' }));
            throw err;
        }

    const profile = await fetch('/api/auth/me', { method: 'GET', credentials: 'include' });
        if (profile.ok) {
            const { user } = await profile.json();
            setUser(user);
            return user;
        } else {
            setUser(null);
            return null;
        }
    };

    const signup = async (credentials) => {
        const res = await fetch('/api/auth/signup', {
            method: 'POST',
            credentials: 'include',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(credentials),
        });

        if (!res.ok) {
            const err = await res.json().catch(() => ({ message: 'Signup failed' }));
            throw err;
        }

        const profile = await fetch('/api/auth/me', {
            method: 'GET',
            credentials: 'include'
        });

        if (profile.ok) {
            const { user } = await profile.json();
            setUser(user);
            return user;
        } else {
            setUser(null);
            return null;
        }
    };

    const logout = async () => {
        try {
            await fetch('/api/auth/logout', {
                method: 'POST',
                credentials: 'include',
                headers: { 'Accept': 'application/json' },
            });
        } catch (err) {
            console.warn('Logout error', err);
        } finally {
            setUser(null);
        }
    };

    return (
        <AuthContext.Provider value={{ user, loading, login, signup, logout }}>
            {children}
        </AuthContext.Provider>
    );
}

export function useAuth() {
    return useContext(AuthContext);
}
