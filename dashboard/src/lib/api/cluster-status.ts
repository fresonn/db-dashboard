import { queryOptions } from '@tanstack/react-query'
import { getStatusQueryOptions } from './gen'

export const clusterStatusQuery = queryOptions({
  ...getStatusQueryOptions(),
  retry: false,
  staleTime: 0,
  refetchOnWindowFocus: false,
  refetchOnReconnect: true
})
