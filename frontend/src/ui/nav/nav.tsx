import React from "react";

export default function NavBar() {
  return (
    <nav className="bg-blue-500 text-white py-4 px-6 sticky top-0 z-10">
      <div className="container mx-auto flex items-center justify-between">
        <a href="#" className="text-xl font-semibold">
          T0
        </a>
        <div className="space-x-4">
          <a href="/" className="hover:text-gray-200">
            Home
          </a>
          <a href="/generated-docs" className="hover:text-gray-200">
            My Documents
          </a>
        </div>
      </div>
    </nav>
  );
}
