import { CgSpinner } from 'react-icons/cg'

export const LoadingSpinner = ({ thing }: { thing?: string }) => {
  return (
    <div className="flex justify-center items-center gap-2 h-full">
      <CgSpinner className="animate-spin h-8 w-8" />
      <p>Loading{thing ? ` ${thing}` : ""}...</p>
    </div>
  )
}
