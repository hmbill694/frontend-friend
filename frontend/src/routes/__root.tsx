import * as React from "react";
import { createRootRoute, Link, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import MainLayout from "../layouts/main-layout";

export const Route = createRootRoute({
  component: () => (
    <>
      <MainLayout>
        <>
          <Outlet />
          <TanStackRouterDevtools />
        </>
      </MainLayout>
    </>
  ),
});
