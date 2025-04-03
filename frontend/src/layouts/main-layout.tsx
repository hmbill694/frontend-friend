import NavBar from "../ui/nav/nav";
import React from "react";

export type MainLayoutProps = {
    children: React.JSX.Element
    title?: string
}

export default function MainLayout(props: MainLayoutProps) {
  return (
      <body className="bg-gray-100 min-h-screen flex flex-col">
        <NavBar />

        <main className="flex-grow py-8 px-6">
          <div className="container mx-auto">
            {props.children}
          </div>
        </main>

        <footer className="bg-gray-800 text-white py-4 px-6">
          <div className="container mx-auto text-center">
            <p>&copy; 2024 T0. All rights reserved.</p>
          </div>
        </footer>
      </body>
  );
}
