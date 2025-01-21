import { useQuery } from "@tanstack/react-query";

export const usePermission = (permission: string) => {
  const { data, isLoading } = useQuery({
    queryKey: ["auth", permission],
    queryFn: async () => {
      const urlParams = new URLSearchParams({
        permission
      });
      const resp = await fetch(`/api/user/can?${urlParams.toString()}`);
      return resp.status === 200;
    }
  });

  return { isAllowed: data ?? false, isLoading };
};
