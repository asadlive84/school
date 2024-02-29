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
