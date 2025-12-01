import { useState } from "react";
import { Link } from "react-router-dom";
import { useAuth } from "../context/AuthContext";

import mail_ic from "../assets/mail.png";
import lock_ic from "../assets/lock.png";
import profile_ic from "../assets/people.png";

export default function Signup() {
  const { signup } = useAuth();

  const months = [
    "January", "February", "March", "April",
    "May", "June", "July", "August",
    "September", "October", "November", "December",
  ];

  const [formData, setFormData] = useState({
    firstName: "",
    lastName: "",
    email: "",
    birthdayDay: "",
    birthdayMonth: "",
    birthdayYear: "",
    password: "",
    confirmPassword: "",
  });

  const [error, setError] = useState("");

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    if (formData.password !== formData.confirmPassword) {
      setError("Passwords do not match!");
      return;
    }

    const birthDate = `${formData.birthdayDay} ${formData.birthdayMonth} ${formData.birthdayYear}`;

    signup({
      firstName: formData.firstName,
      lastName: formData.lastName,
      email: formData.email,
      birthDate,
      password: formData.password,
    });
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-white px-4 py-8">
      <div className="w-full max-w-md">
        <div className="bg-white rounded-2xl border border-slate-200 p-8 shadow-lg">
          <div className="flex items-center gap-2 mb-8 justify-center">
            <span className="font-bold text-3xl text-slate-900">LinguaLink</span>
          </div>

          <h1 className="text-2xl font-bold text-slate-900 mb-2">
            Create Account
          </h1>
          <p className="text-slate-600 mb-8">
            Join to LinhuaLink and start your plan.
          </p>

          {error && (
            <p className="text-red-500 text-center mb-3">{error}</p>
          )}

          <form onSubmit={handleSubmit} className="space-y-4">

            {/* FIRST + LAST NAME */}
            <div className="grid grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-slate-900 mb-2">
                  First Name
                </label>
                <div className="relative">
                  <img src={profile_ic} alt="icon" className="w-4 h-4 absolute left-3 top-1/2 -translate-y-1/2 opacity-70" />
                  <input
                    name="firstName"
                    value={formData.firstName}
                    onChange={handleChange}
                    placeholder="Tom"
                    required
                    className="w-full pl-10 pr-4 py-2 border border-slate-300 rounded-lg focus:ring-blue-600"
                  />
                </div>
              </div>

              <div>
                <label className="block text-sm font-medium text-slate-900 mb-2">
                  Last Name
                </label>
                <div className="relative">
                  <img src={profile_ic} alt="icon" className="w-4 h-4 absolute left-3 top-1/2 -translate-y-1/2 opacity-70" />
                  <input
                    name="lastName"
                    value={formData.lastName}
                    onChange={handleChange}
                    placeholder="Cruise"
                    required
                    className="w-full pl-10 pr-4 py-2 border border-slate-300 rounded-lg focus:ring-blue-600"
                  />
                </div>
              </div>
            </div>

            {/* EMAIL */}
            <div>
              <label className="block text-sm font-medium text-slate-900 mb-2">
                Email
              </label>
              <div className="relative">
                <img src={mail_ic} alt="icon" className="w-5 h-5 absolute left-3 top-1/2 -translate-y-1/2 opacity-70" />
                <input
                  type="email"
                  name="email"
                  value={formData.email}
                  onChange={handleChange}
                  placeholder="you@example.com"
                  required
                  className="w-full pl-10 pr-4 py-2 border border-slate-300 rounded-lg focus:ring-blue-600"
                />
              </div>
            </div>

            {/* DATE */}
            <div>
              <label className="block text-sm font-medium text-slate-900 mb-2">
                Birthdate
              </label>

              <div className="flex gap-3">
                <select
                  name="birthdayDay"
                  value={formData.birthdayDay}
                  onChange={handleChange}
                  className="border rounded-lg p-2 flex-1 focus:ring-blue-600"
                >
                  <option>Day</option>
                  {Array.from({ length: 31 }, (_, i) => (
                    <option key={i + 1}>{i + 1}</option>
                  ))}
                </select>

                <select
                  name="birthdayMonth"
                  value={formData.birthdayMonth}
                  onChange={handleChange}
                  className="border rounded-lg p-2 flex-1 focus:ring-blue-600"
                >
                  <option>Month</option>
                  {months.map((month, index) => (
                    <option key={index}>{month}</option>
                  ))}
                </select>

                <select
                  name="birthdayYear"
                  value={formData.birthdayYear}
                  onChange={handleChange}
                  className="border rounded-lg p-2 flex-1 focus:ring-blue-600"
                >
                  <option>Year</option>
                  {Array.from({ length: 100 }, (_, i) => (
                    <option key={i}>{2025 - i}</option>
                  ))}
                </select>
              </div>
            </div>

            {/* PASSWORD */}
            <div>
              <label className="block text-sm font-medium text-slate-900 mb-2">
                Password
              </label>
              <div className="relative">
                <img src={lock_ic} alt="icon" className="w-5 h-5 absolute left-3 top-1/2 -translate-y-1/2 opacity-70" />
                <input
                  type="password"
                  name="password"
                  value={formData.password}
                  onChange={handleChange}
                  placeholder="••••••••"
                  required
                  className="w-full pl-10 pr-4 py-2 border border-slate-300 rounded-lg focus:ring-blue-600"
                />
              </div>
            </div>

            {/* CONFIRM PASSWORD */}
            <div>
              <label className="block text-sm font-medium text-slate-900 mb-2">
                Confirm Password
              </label>
              <div className="relative">
                <img src={lock_ic} alt="icon" className="w-5 h-5 absolute left-3 top-1/2 -translate-y-1/2 opacity-70" />
                <input
                  type="password"
                  name="confirmPassword"
                  value={formData.confirmPassword}
                  onChange={handleChange}
                  placeholder="••••••••"
                  required
                  className="w-full pl-10 pr-4 py-2 border border-slate-300 rounded-lg focus:ring-blue-600"
                />
              </div>
            </div>

            <button type="submit" className="w-full pt-2 pb-3 rounded-lg bg-blue-600 hover:bg-blue-700 text-white gap-2 mt-6">
              Create account
            </button>
          </form>

          <div className="mt-6 text-center">
            <p className="text-slate-600">
              Already have an account?{" "}
              <Link to="/login" className="text-blue-600 hover:text-blue-700 font-medium">
                Sign In
              </Link>
            </p>
          </div>

        </div>
      </div>
    </div>
  );
}
