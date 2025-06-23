import { Button, Paper, Text, Title } from "@mantine/core";
import { FaMicrosoft } from "react-icons/fa6";

export const LoginPage = () => {
  return (
    <div className="flex justify-center items-center h-screen">
      <Paper shadow="sm" p="xl" className="w-96 h-80 flex flex-col text-center">
        <Title>Login</Title>
        <Text c="gray">Access Management</Text>
        <a href={"/api/auth/login/azureadv2"} className="mt-12">
          <Button color="red" size="lg" >
            <FaMicrosoft size={'1.7rem'} className="mr-2" />
            Office 365
          </Button>
        </a>
      </Paper>
    </div>
  )
}
