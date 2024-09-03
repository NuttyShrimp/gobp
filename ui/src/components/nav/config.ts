import { To } from "react-router-dom"

declare type NavRoute = {
  title: string
  href: To
  children?: Omit<NavRoute, "children">[]
}

export const navRoutes: NavRoute[] = [];
