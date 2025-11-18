import { Routes, Route } from "react-router";
import LogIn from "./pages/login.jsx";
import SignUp from "./pages/signup.jsx";

function AppRouter() {
  return (
    <div>
      <Routes>
        <Route path="/" element={<LogIn/>}/>
        <Route path="/signup" element={<SignUp/>} />
        <Route path="/login" element={<LogIn/>} />
        <Route path="/about" element={<h1>О сайте</h1>} />
      </Routes>
    </div>
  );
}

export default AppRouter;
