import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'

export interface SearchHistory {
    ID: number
    Content: string
    Date: string
}

export const searchHistoryApi = createApi({
    tagTypes: ['SearchHistory'],
    reducerPath: 'searchHistoryApi',
    baseQuery: fetchBaseQuery({ baseUrl: 'http://localhost:8080/' }),
    endpoints: (builder) => ({
        getSearchHistories: builder.query<SearchHistory[], void>({
            query: () => `searchHistories`,
            providesTags: ['SearchHistory']
        }),
        deleteSearchHistory: builder.mutation<void, number>({
            query: (id) => ({
                url: `searchHistory/${id}`,
                method: 'DELETE',
            }),
            invalidatesTags: ['SearchHistory']
        })
    }),
})


export const { useGetSearchHistoriesQuery, useDeleteSearchHistoryMutation } = searchHistoryApi