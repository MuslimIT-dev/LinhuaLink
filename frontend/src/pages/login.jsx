import { Link, useNavigate, useLocation } from "react-router-dom";
import { useState } from "react";
import { useAuth } from "../context/AuthContext";
import mail_ic from "../assets/mail.png";
import lock_ic from "../assets/lock.png";

export default function Login() {
  const { login } = useAuth(); // <-- берём login из контекста!!!!!!!
  const navigate = useNavigate();
  const location = useLocation();
  const from = location.state?.from?.pathname || "/dashboard";

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState(null);
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);
    setSubmitting(true);
    try {
      // Вызов контекстного login — он делает POST /api/auth/login (credentials: 'include')
      await login({ email, password });
      // После успешного логина — редирект обратно (или на /dashboard)
      navigate(from, { replace: true });
    } catch (err) {
      // err — то, что бросил AuthContext.login (обычно объект с message)
      setError(err?.message || "Ошибка при входе");
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-white px-4">
      <div className="w-full max-w-md mx-auto">
        <div className="bg-white rounded-2xl border border-slate-200 p-8 shadow-lg">
          <div className="flex items-center gap-2 mb-8 justify-center">
            <span className="font-bold text-3xl text-slate-900">LinhuaLink</span>
          </div>
          <h1 className="text-2xl font-bold text-slate-900 mb-2">Sign In</h1>
          <p className="text-slate-600 mb-8">Welcome back! Sign in to your account to continue.</p>

          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-slate-900 mb-2">Email</label>
              <div className="relative">
                <img src={mail_ic} alt="icon" className="w-5 h-5 absolute left-3 top-1/2 -translate-y-1/2 opacity-70" />
                <input
                  type="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  placeholder="you@example.com"
                  required
                  className="w-full pl-10 pr-4 py-2 border border-slate-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
                />
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-slate-900 mb-2">Password</label>
              <div className="relative">
                <img src={lock_ic} alt="icon" className="w-5 h-5 absolute left-3 top-1/2 -translate-y-1/2 opacity-70" />
                <input
                  type="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  placeholder="••••••••"
                  required
                  className="w-full pl-10 pr-4 py-2 border border-slate-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
                />
              </div>
            </div>

            {error && <div className="text-red-600 text-sm">{error}</div>}

            <button
              type="submit"
              disabled={submitting}
              className="w-full pt-2 pb-3 rounded-lg bg-blue-600 hover:bg-blue-700 text-white gap-2 mt-6 disabled:opacity-60"
            >
              {submitting ? "Signing in..." : "Sign In"}
            </button>
          </form>

          <div className="mt-6 text-center">
            <p className="text-slate-600">
              Don't have an account?{" "}
              <Link to="/signup" className="text-blue-600 hover:text-blue-700 font-medium">Sign Up</Link>
            </p>
          </div>

          <div className="mt-6 pt-6 border-t border-slate-200">
            <p className="text-xs text-slate-500 text-center">Demo credentials: test@example.com / password123</p>
          </div>
        </div>
      </div>
    </div>
  );
}