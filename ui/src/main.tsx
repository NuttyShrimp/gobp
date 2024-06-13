import React from 'react'
import ReactDOM from 'react-dom/client'
import { router } from './router';
import { RouterProvider } from 'react-router-dom';
import { Toaster } from 'sonner';

import './index.css'
import { QueryClientProvider } from '@tanstack/react-query';
import { queryClient } from './lib/query';
import { AuthProvider } from './lib/context/auth';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <Toaster richColors />
        <RouterProvider router={router} />
      </AuthProvider>
    </QueryClientProvider>
  </React.StrictMode>,
)
