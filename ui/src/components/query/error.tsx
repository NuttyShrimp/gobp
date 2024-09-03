import { useEffect } from 'react';
import { FaExclamationTriangle } from 'react-icons/fa';

export const ErrorView = ({ error }: { error: Error }) => {
  useEffect(() => {
    console.error(error)
  }, [error]);

  return (
    <div className='w-full flex flex-col items-center'>
      <FaExclamationTriangle className="w-8 h-8 text-orange-500" />
      {error.message}
    </div>
  )
}
