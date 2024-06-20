import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { FaMicrosoft } from "react-icons/fa6";

export const LoginPage = () => {
  return (
    <div className="flex justify-center items-center h-screen">
      <Card className="w-96 h-80 flex flex-col">
        <CardHeader>
          <CardTitle className="text-center text-2xl">Login</CardTitle>
          <CardDescription className="text-center text-lg">MyApp</CardDescription>
        </CardHeader>
        <CardContent className="flex justify-center items-center flex-1">
          <a href={"/api/auth/login/microsoftonline"}>
            <Button className="w-52 h-16 flex justify-center items-center text-xl" >
              <FaMicrosoft size={'1.7rem'} className="mr-2" />
              Office 365
            </Button>
          </a>
        </CardContent>
      </Card>
    </div>
  )
}
