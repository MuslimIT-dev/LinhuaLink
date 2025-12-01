import { Outlet, Link } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

export default function MainLayout() {
  const { user, loading, logout } = useAuth();

  return (
    <div>
      <header>
        <h1>LinguaLink</h1>
        <nav>
          <Link to="/">Home</Link>
          <Link to="/dashboard">Dashboard</Link>
          {!loading && user ? (
            <>
              <span>Привет, {user.name}</span>
              <button onClick={logout}>Выйти</button>
            </>
          ) : (
            <Link to="/login">Войти</Link>
          )}
        </nav>
      </header>

      <main><Outlet /></main>
      <footer><p>© 2025.</p></footer>
    </div>
  );
}
