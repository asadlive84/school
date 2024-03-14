// Layout.js
import { Outlet, Link } from "react-router-dom";
import "./Layout.css"; // Import your CSS file

const Layout = () => {
  return (
    <>
      <header className="header">
        <nav>
          <ul>
            <li>
              <Link to="/home">Home</Link>
            </li>
            <li>
              <Link to="/upload">Upload</Link>
            </li>
            <li>
              <Link to="/students">Students</Link>
            </li>
            <li>
              <Link to="/stdlistsessionid">StudentListBySessionId</Link>
            </li>
            <li>
              <Link to="/contact">Contact</Link>
            </li>
            <li>
              <Link to="/student/6d1dcbb7-e472-4b46-8748-2d291e7f6f51/profile">Profile</Link>
            </li>
          </ul>
        </nav>
      </header>

      <div className="content">
        <Outlet />
      </div>

      <footer className="footer">
        <p>© 2024 . All rights reserved.</p>
      </footer>
    </>
  );
};

export default Layout;
