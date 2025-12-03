import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { RouterProvider, createRouter } from '@tanstack/react-router'
import { routeTree } from '@/routeTree.gen'
import { ThemeProvider } from '@/providers/theme-provider'
import { TooltipProvider } from '@/components/ui/shadcn/tooltip'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

import '@/styles/index.css'

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}

declare global {
  interface Window {
    __TANSTACK_QUERY_CLIENT__: import('@tanstack/react-query').QueryClient
  }
}

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 1,
      refetchOnWindowFocus: false,
      staleTime: 1000 * 60
    }
  }
})

const router = createRouter({
  routeTree,
  defaultPreload: 'intent',
  context: {
    queryClient
  }
})

if (import.meta.env.DEV) {
  window.__TANSTACK_QUERY_CLIENT__ = queryClient
}

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <ThemeProvider>
        <TooltipProvider>
          <RouterProvider router={router} />
        </TooltipProvider>
      </ThemeProvider>
    </QueryClientProvider>
  </StrictMode>
)
