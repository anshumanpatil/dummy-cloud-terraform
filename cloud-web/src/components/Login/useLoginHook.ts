import {
  useQuery,
  useMutation,
  useQueryClient,
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'

const getInfra = async () => {
    return fetch('/api/infrastructure')
}

const queryClient = useQueryClient()


export const useLoginHook = () => {
  return useMutation({
    mutationFn: getInfra
  })
}
