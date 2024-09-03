import { Button, Center, Group, Menu, UnstyledButton } from "@mantine/core"
import { navRoutes } from "./config";
import { HiChevronDown } from 'react-icons/hi';
import { Link } from "react-router-dom";

export const NavRow = () => {
  const items = navRoutes.map((route) => {
    const menuItems = route.children?.map((item) => (
      <Menu.Item component={Link} to={item.href} key={item.href.toString()}>{item.title}</Menu.Item>
    ));

    if (menuItems) {
      return (
        <Menu key={route.title} trigger="hover" transitionProps={{ exitDuration: 0 }} withinPortal>
          <Menu.Target>
            <Button
              component={Link}
              to={route.href}
              onClick={(event) => event.preventDefault()}
              variant="subtle"
              color="red"
              className="text-gray-800 hover:text-red-500"
            >
              <Center>
                <span className={"mr-1"}>{route.title}</span>
                <HiChevronDown size="0.9rem" />
              </Center>
            </Button>
          </Menu.Target>
          <Menu.Dropdown>{menuItems}</Menu.Dropdown>
        </Menu>
      );
    }

    return (
      <UnstyledButton component={Link} to={route.href} className={`rounded-md px-2 py-4 font-medium`}>{route.title}</UnstyledButton>
    );
  });
  return (
    <Group gap={"md"} visibleFrom="sm">
      {items}
    </Group>
  )
}
