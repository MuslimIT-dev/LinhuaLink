import { Outlet, Navigate, useLocation } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

export default function DashboardLayout() {
  const { user, loading } = useAuth();
  const location = useLocation();

  if (loading) return <div style={{ padding: 16 }}>Загрузка...</div>;
  if (!user) return <Navigate to="/login" replace state={{ from: location }} />;

  return (
    <div>
      <header style={{ padding: 12, background: '#fafafa', borderBottom: '1px solid #ddd' }}>
        <h2>Safe zone</h2>
      </header>

      <main style={{ padding: 16 }}>
        <Outlet />
      </main>
    </div>
  );
}
