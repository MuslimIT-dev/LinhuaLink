import { Suspense, lazy } from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { AuthProvider } from './context/AuthContext';

import MainLayout from './layouts/mainLayout';
import DashboardLayout from './layouts/DashboardLayout';
 
// const Home = lazy(() => import('./pages/Home'));
const Soon = lazy(() => import('./pages/Soon'));
const Login = lazy(() => import('./pages/login'));
const Signup = lazy(() => import('./pages/signup'));
const NotFound = lazy(() => import('./pages/NotFound'));

function Loader() {
  return (
    <div style={{ padding: 24, textAlign: 'center' }}>
      Загрузка...
    </div>
  );
}

export default function App() {
  return (
      <AuthProvider>
        <Suspense fallback={<Loader />}>
          <Routes>
            <Route path="/" element={<MainLayout />}>
              <Route index element={<Soon />} />
              <Route path="login" element={<Login />} />
              <Route path="signup" element={<Signup />} />
              <Route path="dashboard" element={<DashboardLayout />}>
                <Route index element={<Soon />} />
                <Route path="settings" element={<Soon />} />
              </Route>

              <Route path="*" element={<NotFound />} />
            </Route>
          </Routes>
        </Suspense>
      </AuthProvider>
  );
}