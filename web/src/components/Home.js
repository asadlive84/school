import React from "react";

function Home() {
  return (
    <div className="bg-gray-100 min-h-screen">
      {/* Navbar */}
      <nav className="bg-blue-500 p-4">
        <div className="container mx-auto">
          <h1 className="text-white text-xl font-semibold">My Website</h1>
        </div>
      </nav>

      {/* Hero Section */}
      <header className="py-16 bg-gradient-to-b from-blue-700 to-blue-500 text-center text-white">
        <div className="container mx-auto">
          <h2 className="text-4xl font-bold mb-4">Welcome to Our Website</h2>
          <p className="text-lg">A place where you find everything you need.</p>
        </div>
      </header>

      {/* Features Section */}
      <section className="py-16">
        <div className="container mx-auto grid grid-cols-1 md:grid-cols-3 gap-8">
          {/* Feature 1 */}
          <div className="bg-white p-6 rounded-lg shadow-md">
            <h3 className="text-xl font-semibold mb-2">Feature 1</h3>
            <p className="text-gray-700">Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque tincidunt ullamcorper mi, ac bibendum libero.</p>
          </div>
          {/* Feature 2 */}
          <div className="bg-white p-6 rounded-lg shadow-md">
            <h3 className="text-xl font-semibold mb-2">Feature 2</h3>
            <p className="text-gray-700">Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.</p>
          </div>
          {/* Feature 3 */}
          <div className="bg-white p-6 rounded-lg shadow-md">
            <h3 className="text-xl font-semibold mb-2">Feature 3</h3>
            <p className="text-gray-700">Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
          </div>
        </div>
      </section>

      {/* Call to Action */}
      <section className="bg-blue-500 py-16 text-center text-white">
        <div className="container mx-auto">
          <h2 className="text-3xl font-semibold mb-4">Ready to get started?</h2>
          <p className="text-lg mb-8">Sign up now and start exploring!</p>
          <button className="bg-white text-blue-500 hover:bg-blue-100 text-lg font-semibold py-2 px-8 rounded-full shadow-md transition duration-300 ease-in-out">Sign Up</button>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-gray-800 py-8 text-white">
        <div className="container mx-auto text-center">
          <p>&copy; 2024 My Website. All rights reserved.</p>
        </div>
      </footer>
    </div>
  );
}

export default Home;
